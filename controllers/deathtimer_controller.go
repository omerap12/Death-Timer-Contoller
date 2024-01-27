/*
Copyright 2024 Omer Aplatony.

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
	"log/slog"
	"os"
	"time"

	apiv1alpha1 "github.com/omerap12/death-timer-contoller/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// DeathTimerReconciler reconciles a DeathTimer object
type DeathTimerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=api.omer.aplatony,resources=deathtimers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=api.omer.aplatony,resources=deathtimers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=api.omer.aplatony,resources=deathtimers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the DeathTimer object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.4/pkg/reconcile
func (r *DeathTimerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	operatorLogger := slog.New(jsonHandler)
	DeathTimer := &apiv1alpha1.DeathTimer{}
	err := r.Get(ctx, req.NamespacedName, DeathTimer)
	if err != nil {
		fmt.Println("error getting DeathTimer object")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	namespaces := DeathTimer.Spec.Namespaces
	layout := "2006-01-02T15:04:05"
	for _, namespace := range namespaces {
		namespaceDateString := namespace.Date
		namespaceTtl, err := time.Parse(layout, namespaceDateString)
		if err != nil {
			fmt.Println("error parsing date string of namespace ", namespace.Name)
			return ctrl.Result{RequeueAfter: 2 * time.Second}, nil
		}
		if time.Now().After(namespaceTtl) {
			deleted, err := r.DeleteNamespace(namespace.Name, ctx)
			if err != nil {
				operatorLogger.Error(err.Error())
			}
			if !deleted {
				operatorLogger.Error(err.Error())
			}
			operatorLogger.Info("Deleted", "Namespace", namespace.Name)
		}
	}

	pods := DeathTimer.Spec.Pods
	for _, pod := range pods {
		podDate := pod.Date
		podTtl, err := time.Parse(layout, podDate)
		if err != nil {
			fmt.Println("error parsing date string of pod ", pod.Name)
			return ctrl.Result{RequeueAfter: 2 * time.Second}, nil
		}
		if time.Now().After(podTtl) {
			deleted, err := r.DeletePod(pod.Name, pod.Namespace, ctx)
			if err != nil {
				operatorLogger.Error(err.Error())
			}
			if !deleted {
				operatorLogger.Error(err.Error())
			}
			operatorLogger.Info("Deleted", "Pod", pod.Name, "Namespace", pod.Namespace)
		}
	}

	deployments := DeathTimer.Spec.Deployments
	for _, deployment := range deployments {
		deploymentDate := deployment.Date
		deploymentTTL, err := time.Parse(layout, deploymentDate)
		if err != nil {
			fmt.Println("error parsing date string of deployment ", deployment.Name)
			return ctrl.Result{RequeueAfter: 2 * time.Second}, nil
		}
		if time.Now().After(deploymentTTL) {
			deleted, err := r.DeleteDeployment(deployment.Name, deployment.Namespace, ctx)
			if err != nil {
				operatorLogger.Error(err.Error())
			}
			if !deleted {
				operatorLogger.Error(err.Error())
			}
			operatorLogger.Info("Deleted", "Deployment", deployment.Name, "Namespace", deployment.Namespace)
		}
	}
	return ctrl.Result{RequeueAfter: 2 * time.Second}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DeathTimerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&apiv1alpha1.DeathTimer{}).
		Complete(r)
}

// delete namespace
func (r *DeathTimerReconciler) DeleteNamespace(namespaceName string, ctx context.Context) (bool, error) {
	namespaceDelete := &corev1.Namespace{}
	err := r.Get(ctx, client.ObjectKey{Name: namespaceName}, namespaceDelete)
	if err != nil {
		return false, err
	}
	err = r.Delete(ctx, namespaceDelete)
	if err != nil {
		return false, err
	}
	return true, nil
}

// delete deployment
func (r *DeathTimerReconciler) DeleteDeployment(deploymentName string, namespaceName string, ctx context.Context) (bool, error) {
	deploymentDelete := &appsv1.Deployment{}
	err := r.Get(ctx, client.ObjectKey{Name: deploymentName, Namespace: namespaceName}, deploymentDelete)
	if err != nil {
		return false, err
	}
	err = r.Delete(ctx, deploymentDelete)
	if err != nil {
		return false, err
	}
	return true, nil
}

// delete pod
func (r *DeathTimerReconciler) DeletePod(podName string, namespaceName string, ctx context.Context) (bool, error) {
	podDelete := &corev1.Pod{}
	err := r.Get(ctx, client.ObjectKey{Name: podName, Namespace: namespaceName}, podDelete)
	if err != nil {
		return false, err
	}
	err = r.Delete(ctx, podDelete)
	if err != nil {
		return false, err
	}
	return true, nil
}
