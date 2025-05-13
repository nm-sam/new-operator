/*
Copyright 2025.

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

package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	flashbotv1 "github.com/nm-sam/my-operator/api/v1"
)

// MyCustomResourceReconciler reconciles a MyCustomResource object
type MyCustomResourceReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// +kubebuilder:rbac:groups=flashbot.example.com,resources=mycustomresources,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=flashbot.example.com,resources=mycustomresources/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=flashbot.example.com,resources=mycustomresources/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the MyCustomResource object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.20.4/pkg/reconcile
func (r *MyCustomResourceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	// _ = logf.FromContext(ctx)
	log := logf.FromContext(ctx)


	// TODO(user): your logic here
	var myResource flashbotv1.MyCustomResource
	if err := r.Get(ctx, req.NamespacedName, &myResource); err != nil {
	  log.Error(err, "Failed to fetch MyResource")
	  return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Define the desired Deployment name
	deploymentName := myResource.Name + "-deployment"

	// Check if Deployment already exists
	var deployment appsv1.Deployment
	err := r.Get(ctx, types.NamespacedName{Name: deploymentName, Namespace: myResource.Namespace}, &deployment)
	if err == nil {
	  log.Info("Deployment already exists, skipping creation", "Deployment.Namespace", deployment.Namespace, "Deployment.Name", deployment.Name)
	  return ctrl.Result{}, nil
	}

	// Define a new Deployment
	newDeployment := r.createDeployment(&myResource)

	// Create the Deployment
	if err := r.Create(ctx, newDeployment); err != nil {
	  log.Error(err, "Failed to create Deployment")
	  return ctrl.Result{}, err
	}

	log.Info("Deployment created successfully", "Deployment.Namespace", newDeployment.Namespace, "Deployment.Name", newDeployment.Name)

	return ctrl.Result{}, nil
}


func (r *MyCustomResourceReconciler) createDeployment(myResource *flashbotv1.MyCustomResource) *appsv1.Deployment {
 labels := map[string]string{"app": myResource.Name}
 replicas := int32(myResource.Spec.Replicas)
 return &appsv1.Deployment{
  ObjectMeta: metav1.ObjectMeta{
   Name:      myResource.Name + "-deployment",
   Namespace: myResource.Namespace,
  },
  Spec: appsv1.DeploymentSpec{
  
   Replicas: &replicas, // Use replica count from CRD
   Selector: &metav1.LabelSelector{
    MatchLabels: labels,
   },
   Template: corev1.PodTemplateSpec{
    ObjectMeta: metav1.ObjectMeta{
     Labels: labels,
    },
    Spec: corev1.PodSpec{
     Containers: []corev1.Container{
      {
       Name:  "my-app",
       Image: myResource.Spec.Image, // Use image from CRD
       Ports: []corev1.ContainerPort{
        {
         ContainerPort: 80,
        },
       },
      },
     },
    },
   },
  },
 }
}


// SetupWithManager sets up the controller with the Manager.
func (r *MyCustomResourceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&flashbotv1.MyCustomResource{}).
		Owns(&appsv1.Deployment{}).
		// Named("mycustomresource").
		Complete(r)
}
