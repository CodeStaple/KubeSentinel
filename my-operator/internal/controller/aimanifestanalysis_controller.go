package controller

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	aiv1alpha1 "github.com/example/my-operator/api/v1alpha1" // Adjust if your module path is different
)

// AIManifestAnalysisReconciler reconciles a AIManifestAnalysis object
type AIManifestAnalysisReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=ai.example.com,resources=aimanifestanalyses,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=ai.example.com,resources=aimanifestanalyses/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=ai.example.com,resources=aimanifestanalyses/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *AIManifestAnalysisReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx).WithName("aimanifestanalysis_controller")
	logger.Info("Reconciling AIManifestAnalysis", "Request.Namespace", req.Namespace, "Request.Name", req.Name)

	// Fetch the AIManifestAnalysis instance
	instance := &aiv1alpha1.AIManifestAnalysis{}
	err := r.Get(ctx, req.NamespacedName, instance)
	if err != nil {
		if client.IgnoreNotFound(err) != nil {
			logger.Error(err, "Failed to get AIManifestAnalysis")
			return ctrl.Result{}, err
		}
		// Resource not found. Could have been deleted after reconcile request.
		// Return and don't requeue
		logger.Info("AIManifestAnalysis resource not found. Ignoring since object must be deleted.")
		return ctrl.Result{}, nil
	}

	// Log the Manifest spec
	logger.Info("AIManifestAnalysis found", "Manifest Spec", instance.Spec.Manifest)

	// For now, do nothing else. Future logic will go here.

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *AIManifestAnalysisReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&aiv1alpha1.AIManifestAnalysis{}).
		Complete(r)
}
