package controllers

import (
	"carina/utils/log"
	"context"
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// PersistentVolumeClaimReconciler reconciles a PersistentVolumeClaim object
type PersistentVolumeReconciler struct {
	client.Client
	APIReader client.Reader
	Log       logr.Logger
}

// +kubebuilder:rbac:groups="",resources=persistentvolumes,verbs=get;list;watch;update
// +kubebuilder:rbac:groups="",resources=pods,verbs=get;list;watch;delete

// Reconcile finalize PVC
func (r *PersistentVolumeReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	_ = r.Log.WithValues("persistentvolume", req.NamespacedName)
	// your logic here
	pv := &corev1.PersistentVolume{}
	err := r.Get(ctx, req.NamespacedName, pv)
	switch {
	case err == nil:
	case apierrors.IsNotFound(err):
		return ctrl.Result{}, nil
	default:
		return ctrl.Result{}, err
	}
	log.Info(pv.Spec.HostPath.String())

	//if pv.DeletionTimestamp == nil {
	//	return ctrl.Result{}, nil
	//}

	//needFinalize := false
	//for _, fin := range pvc.Finalizers {
	//	if fin == utils.PVCFinalizer {
	//		needFinalize = true
	//		break
	//	}
	//}
	//if !needFinalize {
	//	return ctrl.Result{}, nil
	//}
	//
	//// Requeue until other finalizers complete their jobs.
	//if len(pvc.Finalizers) != 1 {
	//	return ctrl.Result{
	//		Requeue:      true,
	//		RequeueAfter: 10 * time.Second,
	//	}, nil
	//}
	//
	//pvc.Finalizers = nil
	//if err := r.Update(ctx, pvc); err != nil {
	//	log.Error(err, "failed to remove finalizer", "name", pvc.Name)
	//	return ctrl.Result{}, err
	//}
	//
	//// sleep shortly to wait StatefulSet controller notices PVC deletion
	//time.Sleep(100 * time.Millisecond)
	//
	//pods, err := r.getPodsByPVC(ctx, pvc)
	//if err != nil {
	//	log.Error(err, "unable to fetch PodList for a PVC", "pvc", pvc.Name, "namespace", pvc.Namespace)
	//	return ctrl.Result{}, err
	//}
	//for _, pod := range pods {
	//	err := r.Delete(ctx, &pod, client.GracePeriodSeconds(1))
	//	if err != nil {
	//		log.Error(err, "unable to delete Pod", "name", pod.Name, "namespace", pod.Namespace)
	//		return ctrl.Result{}, err
	//	}
	//	log.Info("deleted Pod", "name", pod.Name, "namespace", pod.Namespace)
	//}

	return ctrl.Result{}, nil
}

func (r *PersistentVolumeReconciler) getPodsByPVC(ctx context.Context, pvc *corev1.PersistentVolumeClaim) ([]corev1.Pod, error) {
	var pods corev1.PodList
	// query directly to API server to avoid latency for cache updates
	err := r.APIReader.List(ctx, &pods, client.InNamespace(pvc.Namespace))
	if err != nil {
		return nil, err
	}

	var result []corev1.Pod
OUTER:
	for _, pod := range pods.Items {
		for _, volume := range pod.Spec.Volumes {
			if volume.PersistentVolumeClaim == nil {
				continue
			}
			if volume.PersistentVolumeClaim.ClaimName == pvc.Name {
				result = append(result, pod)
				continue OUTER
			}
		}
	}

	return result, nil
}

// SetupWithManager sets up Reconciler with Manager.
func (r *PersistentVolumeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	pred := predicate.Funcs{
		CreateFunc:  func(event.CreateEvent) bool { return true },
		DeleteFunc:  func(event.DeleteEvent) bool { return false },
		UpdateFunc:  func(event.UpdateEvent) bool { return false },
		GenericFunc: func(event.GenericEvent) bool { return false },
	}
	return ctrl.NewControllerManagedBy(mgr).
		WithEventFilter(pred).
		For(&corev1.PersistentVolumeClaim{}).
		Complete(r)
}
