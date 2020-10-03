/*


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

package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	v1alpha1 "github.com/kaladaOpuiyo/node-role-controller/api/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var (
	err error
	c   = &v1alpha1.NodeRoleController{}
	n   = &corev1.Node{}
)

const (
	roleLabel = "kubernetes.io/role"
)

// NodeRoleControllerReconciler reconciles a NodeRoleController object
type NodeRoleControllerReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=operators.node.role.controller.io,resources=noderolecontrollers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=operators.node.role.controller.io,resources=noderolecontrollers/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=core,resources=node,verbs=get;list;watch;update;patch;delete

// Reconcile ...
func (r *NodeRoleControllerReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("noderolecontroller", req.Name)

	if req.Namespace != "" {
		err = r.Client.Get(ctx, req.NamespacedName, c)
		if err != nil {
			if errors.IsNotFound(err) {
				// Request object not found, could have been deleted after reconcile req.
				// Owned objects are automatically garbage collected.
				// For additional cleanup logic use finalizers.
				// Return and don't requeue
				log.Info("NodeRoleController resource not found. Ignoring since object must be deleted")
				return ctrl.Result{}, nil
			}
		}
	}

	if req.Namespace == "" {
		err = r.Client.Get(ctx, req.NamespacedName, n)
		if err != nil {
			if errors.IsNotFound(err) {
				// Request object not found, could have been deleted after reconcile req.
				// Owned objects are automatically garbage collected.
				// For additional cleanup logic use finalizers.
				// Return and don't requeue
				log.Info("Could Not find Node")
				return ctrl.Result{}, nil
			}
		}
	}

	if n.GetName() == "" {
		delay := time.Second * time.Duration(5)
		return ctrl.Result{RequeueAfter: delay}, nil
	}

	for _, role := range c.Spec.Roles {

		// Check provided label is set on node
		if _, ok := n.ObjectMeta.Labels[role.Label]; ok {

			// check if role is already set
			if n.ObjectMeta.Labels[roleLabel] == role.Name {
				return reconcile.Result{}, nil
			}

			// Check if label value is correct
			if n.ObjectMeta.Labels[role.Label] == role.Value {

				// Set Role
				n.ObjectMeta.Labels[roleLabel] = role.Name

				// This should be Patch
				err = r.Client.Update(ctx, n)
				if err != nil {
					return reconcile.Result{}, fmt.Errorf("could not write to Node: %+v", err)
				}

				log.Info("Node Role Set", "Role ", n.ObjectMeta.Labels[roleLabel])
			}

		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager ...
func (r *NodeRoleControllerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.NodeRoleController{}).
		Watches(&source.Kind{Type: &corev1.Node{}}, &handler.EnqueueRequestForObject{}).
		Complete(r)
}
