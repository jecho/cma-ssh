/*
Copyright 2018 Samsung SDS.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package machine

import (
	"context"
	"fmt"
	"github.com/samsung-cnct/cma-ssh/pkg/util/k8sutil"
	"time"

	"github.com/golang/glog"
	"github.com/masterminds/semver"
	"github.com/samsung-cnct/cma-ssh/pkg/apis/cluster/common"
	clusterv1alpha1 "github.com/samsung-cnct/cma-ssh/pkg/apis/cluster/v1alpha1"
	"github.com/samsung-cnct/cma-ssh/pkg/util"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

type backgroundMachineOp func(r *ReconcileMachine, machineInstance *clusterv1alpha1.CnctMachine, privateKey []byte) (string, error)

// Add creates a new Machine Controller and adds it to the Manager with default RBAC. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileMachine{
		Client:        mgr.GetClient(),
		scheme:        mgr.GetScheme(),
		EventRecorder: mgr.GetRecorder("MachineController"),
	}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("machine-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to Machine
	err = c.Watch(&source.Kind{Type: &clusterv1alpha1.CnctMachine{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &clusterv1alpha1.CnctCluster{}},
		&handler.EnqueueRequestsFromMapFunc{ToRequests: util.ClusterToMachineMapper{Client: mgr.GetClient()}})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileMachine{}

// ReconcileMachine reconciles a Machine object
type ReconcileMachine struct {
	client.Client
	scheme *runtime.Scheme
	record.EventRecorder
}

// Reconcile reads that stamakte of the cluster for a Machine object and makes changes based on the state read
// and what is in the Machine.Spec
// +kubebuilder:rbac:groups=cluster.cnct.sds.samsung.com,resources=cnctmachines;cnctclusters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=events,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=namespaces,verbs=get;list;watch;create;update;patch;delete
func (r *ReconcileMachine) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	// Fetch the Machine machine
	machineInstance := &clusterv1alpha1.CnctMachine{}
	err := r.Get(context.Background(), request.NamespacedName, machineInstance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// log.Error(err, "could not find machine", "machine", request)
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		glog.Errorf("could not read machine %s: %q", request.Name, err)
		return reconcile.Result{}, err
	}

	clusterInstance, err := getCluster(r.Client, machineInstance.GetNamespace(), machineInstance.Spec.ClusterRef)
	if err != nil {
		if errors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			glog.Errorf("could not find cluster %s: %q", machineInstance.Spec.ClusterRef, err)

			machineInstance.Status.Phase = common.ErrorMachinePhase
			err = r.updateStatus(machineInstance, corev1.EventTypeWarning, common.ErrResourceFailed,
				common.MessageResourceFailed, machineInstance.GetName())
			if err != nil {
				glog.Errorf("could not update status of object machine %s: %q", machineInstance.GetName(), err)
				return reconcile.Result{}, err
			}
			return reconcile.Result{}, err
		}
		// Error reading the object - requeue the request.
		glog.Errorf("error reading object machine %s: %q", machineInstance.GetName(), err)
		return reconcile.Result{}, err
	}

	// if we are not being deleted
	if machineInstance.DeletionTimestamp.IsZero() {
		if machineInstance.Status.Phase == common.ReadyMachinePhase {
			// Object is currently ready, so we need to check if
			// the state change came in from cluster version change

			// build semvers
			machineKuberneteVersion, err := semver.NewVersion(machineInstance.Status.KubernetesVersion)
			if err != nil {
				glog.Errorf("could not parse object machine %s kubernetes version: %q",
					machineInstance.GetName(), err)
				return reconcile.Result{}, err
			}

			clusterKuberneteVersion, err := semver.NewVersion(clusterInstance.Spec.KubernetesVersion)
			if err != nil {
				glog.Errorf("could not parse cluster %s kubernetes version: %q",
					clusterInstance.GetName(), err)
				return reconcile.Result{}, err
			}

			// if cluster object kubernetes version is not equal to machine kubernetes version,
			// trigger an upgrade
			glog.Infof("Checking cluster version to see if upgrade is needed for machine %s",
				machineInstance.GetName())
			if !clusterKuberneteVersion.Equal(machineKuberneteVersion) {
				glog.Infof("will upgrade %s to %s", machineInstance.GetName(),
					clusterKuberneteVersion.String())
				return r.handleUpgrade(machineInstance, clusterInstance)
			} else {
				glog.Infof("no upgrade is needed for machine %s", machineInstance.GetName())
			}

		} else if machineInstance.Status.Phase == common.ErrorMachinePhase {
			// if object is in error state, just ignore and move on
			return reconcile.Result{}, nil
		} else if machineInstance.Status.Phase == common.UpgradingMachinePhase {
			// if object is currently updating, ignore and move on
			return reconcile.Result{}, nil
		} else {
			// if not an error, in progress update, or an update request, we must be
			// creating a new machine. Trigger bootstrap
			return r.handleCreate(machineInstance, clusterInstance)
		}
	} else {
		// The object is being deleted, do a node delete
		return r.handleDelete(machineInstance)
	}

	return reconcile.Result{}, nil
}

func (r *ReconcileMachine) handleDelete(machineInstance *clusterv1alpha1.CnctMachine) (reconcile.Result, error) {
	// if already deleting, ignore
	if machineInstance.Status.Phase == common.DeletingMachinePhase {
		glog.Infof("Delete: Already deleting machine %s", machineInstance.GetName())
		return reconcile.Result{}, nil
	}

	// get cluster status to determine whether we should proceed,
	// i.e. if there is a create in progress, we wait for it to either finish or error
	clusterInstance, err := getCluster(r.Client, machineInstance.GetNamespace(), machineInstance.Spec.ClusterRef)
	if err != nil {
		return reconcile.Result{}, err
	}
	machineList, err := util.GetClusterMachineList(r.Client, clusterInstance.GetName())
	if err != nil {
		glog.Errorf("could not list Machines: %q", err)
		return reconcile.Result{}, err
	}
	if !util.IsReadyForDeletion(machineList) {
		glog.Infof("Delete: Waiting for cluster %s to finish reconciling", machineInstance.Spec.ClusterRef)
		return reconcile.Result{Requeue: true}, nil
	}

	if util.ContainsString(machineInstance.Finalizers, clusterv1alpha1.MachineFinalizer) {
		// update status to "deleting"
		machineInstance.Status.Phase = common.DeletingMachinePhase
		err := r.updateStatus(machineInstance, corev1.EventTypeNormal,
			common.ResourceStateChange, common.MessageResourceStateChange,
			machineInstance.GetName(), common.DeletingMachinePhase)
		if err != nil {
			glog.Errorf("could not update status of machine %s: %q", machineInstance.GetName())
			return reconcile.Result{}, err
		}

		// start delete process
		r.backgroundRunner(doDelete, machineInstance, "handleDelete")

	}

	return reconcile.Result{}, nil
}

func (r *ReconcileMachine) handleUpgrade(machineInstance *clusterv1alpha1.CnctMachine, clusterInstance *clusterv1alpha1.CnctCluster) (reconcile.Result, error) {
	// if already upgrading, move on
	if machineInstance.Status.Phase == common.UpgradingMachinePhase {
		return reconcile.Result{}, nil
	}

	// get cluster status to determine whether we should proceed,
	// i.e. if there is a create in progress, we wait for it to either finish or error
	clusterInstance, err := getCluster(r.Client, machineInstance.GetNamespace(), machineInstance.Spec.ClusterRef)
	if err != nil {
		return reconcile.Result{}, err
	}
	machineList, err := util.GetClusterMachineList(r.Client, clusterInstance.GetName())
	if err != nil {
		glog.Errorf("could not list Machines for cluster %s: %q", machineInstance.GetName(), err)
		return reconcile.Result{}, err
	}
	// if not ok to upgrade with error, return and do not requeue
	ok, err := util.IsReadyForUpgrade(machineList)
	if err != nil {
		glog.Errorf("cannot upgrade machine %s: %q", machineInstance.GetName(), err)
		return reconcile.Result{}, nil
	}
	// if not ok to upgrade, try later
	if !ok {
		glog.Infof("Upgrade: Waiting for cluster %s to finish reconciling", machineInstance.Spec.ClusterRef)
		return reconcile.Result{Requeue: true}, nil
	}

	// update status to "upgrading"
	machineInstance.Status.Phase = common.UpgradingMachinePhase
	err = r.updateStatus(machineInstance, corev1.EventTypeNormal,
		common.ResourceStateChange, common.MessageResourceStateChange,
		machineInstance.GetName(), common.UpgradingMachinePhase)
	if err != nil {
		glog.Errorf("could not update status of machine %s: %q", machineInstance.GetName(), err)
		return reconcile.Result{}, err
	}

	// otherwise start upgrade process
	r.backgroundRunner(doUpgrade, machineInstance, "handleUpgrade")

	return reconcile.Result{}, nil
}

func (r *ReconcileMachine) handleCreate(machineInstance *clusterv1alpha1.CnctMachine, clusterInstance *clusterv1alpha1.CnctCluster) (reconcile.Result, error) {
	if machineInstance.Status.Phase == common.ProvisioningMachinePhase {
		return reconcile.Result{}, nil
	}

	// Add the finalizer
	if !util.ContainsString(machineInstance.Finalizers, clusterv1alpha1.MachineFinalizer) {
		machineInstance.Finalizers =
			append(machineInstance.Finalizers, clusterv1alpha1.MachineFinalizer)
	}

	// update status to "creating"
	machineInstance.Status.Phase = common.ProvisioningMachinePhase
	machineInstance.Status.KubernetesVersion = clusterInstance.Spec.KubernetesVersion
	err := r.updateStatus(machineInstance, corev1.EventTypeNormal,
		common.ResourceStateChange, common.MessageResourceStateChange,
		machineInstance.GetName(), common.ProvisioningMachinePhase)
	if err != nil {
		glog.Errorf("could not update status of machine %s: %q", machineInstance.GetName(), err)
		return reconcile.Result{}, err
	}

	// start bootstrap process
	r.backgroundRunner(doBootstrap, machineInstance, "handleCreate")

	return reconcile.Result{}, nil
}

func (r *ReconcileMachine) updateStatus(machineInstance *clusterv1alpha1.CnctMachine, eventType string,
	event common.ControllerEvents, eventMessage common.ControllerEvents, args ...interface{}) error {

	machineFreshInstance := &clusterv1alpha1.CnctMachine{}
	err := r.Get(
		context.Background(),
		client.ObjectKey{
			Namespace: machineInstance.GetNamespace(),
			Name:      machineInstance.GetName(),
		}, machineFreshInstance)
	if err != nil {
		return err
	}

	machineFreshInstance.Finalizers = machineInstance.Finalizers
	machineFreshInstance.Status.Phase = machineInstance.Status.Phase
	machineFreshInstance.Status.KubernetesVersion = machineInstance.Status.KubernetesVersion
	machineFreshInstance.Status.LastUpdated = &metav1.Time{Time: time.Now()}

	err = r.Update(context.Background(), machineFreshInstance)
	if err != nil {
		return err
	}

	r.Eventf(machineFreshInstance, eventType,
		string(event), string(eventMessage), args...)

	return nil
}

func getCluster(c client.Client, namespace string, clusterName string) (*clusterv1alpha1.CnctCluster, error) {

	clusterKey := client.ObjectKey{
		Namespace: namespace,
		Name:      clusterName,
	}

	clusterInstance := &clusterv1alpha1.CnctCluster{}
	err := c.Get(context.Background(), clusterKey, clusterInstance)
	if err != nil {
		return nil, err
	}

	return clusterInstance, nil
}

func doBootstrap(r *ReconcileMachine, machineInstance *clusterv1alpha1.CnctMachine, privateKey []byte) (string, error) {
	// Setup bootstrap repo
	_, cmd, err := RunSshCommand(r.Client, machineInstance, privateKey, InstallBootstrapRepo, make(map[string]string))
	if err != nil {
		return cmd, err
	}

	// install local nginx proxy
	_, cmd, err = RunSshCommand(r.Client, machineInstance, privateKey, InstallNginx, make(map[string]string))
	if err != nil {
		return cmd, err
	}

	// install docker
	_, cmd, err = RunSshCommand(r.Client, machineInstance, privateKey, InstallDocker, make(map[string]string))
	if err != nil {
		return cmd, err
	}

	// install kubernetes components
	_, cmd, err = RunSshCommand(r.Client, machineInstance, privateKey, InstallKubernetes, make(map[string]string))
	if err != nil {
		return cmd, err
	}

	// if this is a master, proceed with bootstrap
	if util.ContainsRole(machineInstance.Spec.Roles, common.MachineRoleMaster) {
		// run kubeadm init
		_, cmd, err = RunSshCommand(r.Client, machineInstance, privateKey, KubeadmInit, make(map[string]string))
		if err != nil {
			return cmd, err
		}

	} else if util.ContainsRole(machineInstance.Spec.Roles, common.MachineRoleWorker) {
		// on worker, see if master is able to run kubeadm
		// if it is, run kubeadm token create and use the token to
		// do kubeadm join.
		// otherwise wait for a bit and try again.

		// get machine list
		machineList := &clusterv1alpha1.CnctMachineList{}
		err := r.List(
			context.Background(),
			&client.ListOptions{LabelSelector: labels.Everything()},
			machineList)
		if err != nil {
			return "get machine list", err
		}

		masterMachine, err := util.GetMaster(machineList.Items)
		if err != nil {
			return "util.GetMaster()", err
		}

		var token []byte
		err = util.Retry(120, 10*time.Second, func() error {
			// run kubeadm create token on master machine, get token back
			glog.Infof("Trying to get kubeadm token from master %s for node %s",
				masterMachine.GetName(), machineInstance.GetName())
			token, cmd, err = RunSshCommand(r.Client, masterMachine, privateKey, KubeadmTokenCreate, make(map[string]string))
			if err != nil {
				glog.Infof("Waiting for kubeadm to be able to create a token on master %s for machine %s",
					masterMachine.GetName(), machineInstance.GetName())
				return err
			}

			glog.Infof("Got master kubeadm token: %s", string(token[:]))
			return nil
		})
		if err != nil {
			glog.Errorf("Failed to get kubeadm token from master %s for machine %s in time: %q",
				masterMachine.GetName(), machineInstance.GetName(), err)
			return "KubeadmTokenCreate", err
		}

		// run kubeadm join on worker machine
		_, cmd, err = RunSshCommand(r.Client, machineInstance, privateKey,
			KubeadmJoin, map[string]string{"token": string(token[:]), "master": masterMachine.Spec.SshConfig.Host})
		if err != nil {
			return cmd, err
		}
	}

	// Set status to ready
	machineInstance.Status.Phase = common.ReadyMachinePhase
	err = r.updateStatus(machineInstance, corev1.EventTypeNormal,
		common.ResourceStateChange, common.MessageResourceStateChange,
		machineInstance.GetName(), common.ReadyMachinePhase)
	if err != nil {
		return "updateStatus()", err
	}
	return "doBootstrap()", err
}

func doUpgrade(r *ReconcileMachine, machineInstance *clusterv1alpha1.CnctMachine, privateKey []byte) (string, error) {
	// get the cluster instance
	clusterInstance, err := getCluster(r.Client, machineInstance.GetNamespace(), machineInstance.Spec.ClusterRef)
	if err != nil {
		glog.Errorf("Could not get cluster %s: %q", clusterInstance.GetName(), err)
		return "getCluster()", err
	}

	// if master, do the upgrade.
	// otherwise wait for master to finish upgrade then
	// proceed with upgrade
	// master is done upgrading when it is in ready status and
	// its status kubernetes version matches clusters kubernetes version
	if util.ContainsRole(machineInstance.Spec.Roles, common.MachineRoleMaster) {
		glog.Infof("running upgrade on master %s for cluster %s",
			machineInstance.GetName(), machineInstance.Spec.ClusterRef)
		_, cmd, err := RunSshCommand(r.Client, machineInstance, privateKey,
			UpgradeMaster, make(map[string]string))
		if err != nil {
			return cmd, err
		}

	} else {
		glog.Infof("running upgrade on worker %s for cluster %s",
			machineInstance.GetName(), machineInstance.Spec.ClusterRef)
		err = util.Retry(120, 10*time.Second, func() error {
			// get list of machines
			machineList, err := util.GetClusterMachineList(r.Client, clusterInstance.GetName())
			if err != nil {
				glog.Errorf("could not list Machines for cluster %s: %q",
					machineInstance.Spec.ClusterRef, err)
				return err
			}

			masterMachine, err := util.GetMaster(machineList)
			if err != nil {
				glog.Errorf("could not get master instance for cluster %s: %q",
					machineInstance.Spec.ClusterRef, err)
				return err
			}

			if masterMachine.Status.Phase != common.ReadyMachinePhase {
				glog.Infof("master %s of cluster %s is not ready, will retry machine %s", masterMachine.GetName(),
					machineInstance.Spec.ClusterRef, machineInstance.GetName())
				return fmt.Errorf("master %s of cluster %s is not ready", masterMachine.GetName(),
					machineInstance.Spec.ClusterRef)
			}

			if masterMachine.Status.KubernetesVersion != clusterInstance.Spec.KubernetesVersion {
				glog.Infof("master %s of cluster %s is not done with upgrade, will retry machine %s",
					masterMachine.GetName(), machineInstance.Spec.ClusterRef, machineInstance.GetName())
				return fmt.Errorf("master %s of cluster %s is not ready", masterMachine.GetName(),
					machineInstance.Spec.ClusterRef)
			}

			glog.Info("master phase: " + string(masterMachine.Status.Phase) +
				" master version: " + masterMachine.Status.KubernetesVersion)
			return nil
		})
		if err != nil {
			glog.Error(err, "Master failed to upgrade in time")
			return "doUpgrade", err
		}

		machineList, err := util.GetClusterMachineList(r.Client, clusterInstance.GetName())
		if err != nil {
			glog.Error("could not list Machines for cluster %s: %q",
				clusterInstance.GetName(), err)
			return "util.GetClusterMachineList", err
		}

		masterMachine, err := util.GetMaster(machineList)
		if err != nil {
			glog.Errorf("could not get master instance for cluster %s: %q",
				clusterInstance.GetName(), err)
			return "util.GetMaster", err
		}

		// get admin kubeconfig
		kubeConfig, cmd, err := RunSshCommand(r.Client, masterMachine, privateKey,
			GetKubeConfig, make(map[string]string))
		if err != nil {
			return cmd, err
		}

		// run node upgrade
		_, cmd, err = RunSshCommand(r.Client, machineInstance, privateKey,
			UpgradeNode, map[string]string{"admin.conf": string(kubeConfig[:])})
		if err != nil {
			return cmd, err
		}
	}

	machineInstance.Status.Phase = common.ReadyMachinePhase
	machineInstance.Status.KubernetesVersion = clusterInstance.Spec.KubernetesVersion

	err = r.updateStatus(machineInstance, corev1.EventTypeNormal,
		common.ResourceStateChange, common.MessageResourceStateChange,
		machineInstance.GetName(), common.ReadyMachinePhase)
	if err != nil {
		return "updateStatus()", err
	}

	return "doUpgrade()", nil
}

func doDelete(r *ReconcileMachine, machineInstance *clusterv1alpha1.CnctMachine, privateKey []byte) (string, error) {
	// there is a few possibilities here:
	// 1. Cluster is being deleted. We can just delete this machine
	// 2. Worker machine is being deleted. We should drain the machine first, and then delete it
	// 3. Master machine is being deleted. We should issue a warning, and then delete it.
	// get the cluster instance
	clusterInstance, err := getCluster(r.Client, machineInstance.GetNamespace(), machineInstance.Spec.ClusterRef)
	if err != nil {
		return "getCluster()", err
	}

	// Cluster is NOT being deleted.
	if clusterInstance.Status.Phase != common.StoppingClusterPhase {
		// get the master machine
		machineList, err := util.GetClusterMachineList(r.Client, clusterInstance.GetName())
		if err != nil {
			glog.Errorf("could not list Machines for cluster %s: %q",
				clusterInstance.GetName(), err)
			return "util.GetClusterMachineList", err
		}
		masterMachine, err := util.GetMaster(machineList)
		if err != nil {
			glog.Errorf("could not get master for cluster %s: %q",
				clusterInstance.GetName(), err)
			return "util.GetMaster", err
		}

		// TODO: this will need to be handled better
		if masterMachine.GetName() == machineInstance.GetName() {
			glog.Infof("WARNING!!! DELETING MASTER, %s"+
				"CLUSTER %s WILL NOT FUNCTION WITHOUT NEW MASTER AND FULL RESET",
				masterMachine.GetName(), clusterInstance.GetName())
		}

		// get admin kubeconfig
		kubeConfig, cmd, err := RunSshCommand(r.Client, masterMachine, privateKey,
			GetKubeConfig, make(map[string]string))
		if err != nil {
			return cmd, err
		}

		// run node drain
		_, cmd, err = RunSshCommand(r.Client, machineInstance, privateKey,
			DrainAndDeleteNode, map[string]string{"admin.conf": string(kubeConfig[:])})
		if err != nil {
			return cmd, err
		}
	}

	// run delete command
	_, cmd, err := RunSshCommand(r.Client, machineInstance, privateKey,
		DeleteNode, make(map[string]string))
	if err != nil {
		glog.Errorf("failed to clean up physical node for machine %s Manual cleanup might be required for %s",
			machineInstance.GetName(), machineInstance.Spec.SshConfig.Host)
	}

	machineInstance.Finalizers =
		util.RemoveString(machineInstance.Finalizers, clusterv1alpha1.MachineFinalizer)
	if err := r.updateStatus(machineInstance, corev1.EventTypeNormal,
		common.ResourceStateChange, common.MessageResourceStateChange,
		machineInstance.GetName(), common.DeletingMachinePhase); err != nil {
		glog.Errorf("failed to update status for machine %s: %q", machineInstance.GetName(), err)
	}

	return cmd, err
}

func (r *ReconcileMachine) backgroundRunner(op backgroundMachineOp,
	machineInstance *clusterv1alpha1.CnctMachine, operationName string) {

	type commandError struct {
		Err error
		Cmd string
	}

	// start bootstrap command (or pre upgrade etc)
	opResult := make(chan commandError)
	timer := time.NewTimer(30 * time.Minute)

	// get the cluster instance
	clusterInstance, err := getCluster(r.Client, machineInstance.GetNamespace(), machineInstance.Spec.ClusterRef)
	if err != nil {
		glog.Errorf("Could not get cluster %s: %q", clusterInstance.GetName(), err)
		return
	}

	privateKeySecret, err := k8sutil.GetSecret(r.Client, clusterInstance.Spec.Secret, clusterInstance.GetNamespace())
	if err != nil {
		glog.Errorf("Could not get cluster %s private key secret %s: %q",
			clusterInstance.GetName(), clusterInstance.Spec.Secret, err)
		return
	}
	privateKey := privateKeySecret.Data["private-key"]

	go func(ch chan<- commandError) {
		cmd, err := op(r, machineInstance, privateKey)
		ch <- commandError{Err: err, Cmd: cmd}
		close(opResult)
	}(opResult)

	go func() {
		timedOut := false
		select {
		// cluster operation timeouts shouldn't take longer than 10 minutes
		case <-timer.C:
			err := fmt.Errorf("operation %s timed out for machine %s",
				operationName, machineInstance.GetNamespace())
			glog.Errorf("Provisioning operation timed out for object machine %s, cluster %s: %q",
				machineInstance.GetName(), machineInstance.Spec.ClusterRef, err)

			timer.Stop()
			timedOut = true
		case result := <-opResult:
			// if finished with error
			if result.Err != nil {
				glog.Errorf("could not complete machine %s pre-start step %s command %s: %q",
					machineInstance.GetName(), operationName, result.Cmd, result.Err)

				// set error status
				machineInstance.Status.Phase = common.ErrorMachinePhase
				err := r.updateStatus(machineInstance, corev1.EventTypeWarning,
					common.ResourceStateChange, common.MessageResourceStateChange,
					machineInstance.GetName(), common.ErrorMachinePhase)
				if err != nil {
					glog.Errorf("could not update status of machine %s of cluster %s: %q",
						machineInstance.GetName(), machineInstance.Spec.ClusterRef, err)
				}
			}
		}

		// if timed out, wait for operation to complete
		if timedOut {
			select {
			case result := <-opResult:
				glog.Errorf("Timed out machine %s pre-start step %s command %s: %q",
					machineInstance.GetName(), operationName, result.Cmd, result.Err)
			}

			// set error status
			machineInstance.Status.Phase = common.ErrorMachinePhase
			err := r.updateStatus(machineInstance, corev1.EventTypeWarning,
				common.ResourceStateChange, common.MessageResourceStateChange,
				machineInstance.GetName(), common.ErrorMachinePhase)
			if err != nil {
				glog.Errorf("could not update status of machine %s cluster %s: %q",
					machineInstance.GetName(), machineInstance.Spec.ClusterRef, err)
			}
		}

	}()
}
