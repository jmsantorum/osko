package controller

import (
	"context"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	openslov1 "github.com/oskoperator/osko/apis/openslo/v1"
)

// SLIReconciler reconciles a SLI object
type SLIReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=openslo.com,resources=slis,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=openslo.com,resources=slis/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=openslo.com,resources=slis/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the SLI object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.15.0/pkg/reconcile
func (r *SLIReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	log := log.FromContext(ctx)

	var sli openslov1.SLI

	err := r.Get(ctx, req.NamespacedName, &sli)
	if err != nil {
		if apierrors.IsNotFound(err) {
			log.Info("SLI deleted")
			return ctrl.Result{}, nil
		}

		log.Error(err, errGetDS)
		return ctrl.Result{}, nil
	}

	log.Info("SLI reconciled", "SLI Name", sli.Name, "SLI Namespace", sli.Namespace)
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *SLIReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&openslov1.SLI{}).
		Complete(r)
}
