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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DeathTimerSpec defines the desired state of DeathTimer
type DeathTimerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Namespaces  []NameSpaceName  `json:"namespaces,omitempty"`
	Deployments []DeploymentName `json:"deployments,omitempty"`
	Pods        []PodName        `json:"pods,omitempty"`
}

// DeathTimerStatus defines the observed state of DeathTimer
type DeathTimerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DeathTimer is the Schema for the deathtimers API
type DeathTimer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DeathTimerSpec   `json:"spec,omitempty"`
	Status DeathTimerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DeathTimerList contains a list of DeathTimer
type DeathTimerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeathTimer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DeathTimer{}, &DeathTimerList{})
}

type NameSpaceName struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

type DeploymentName struct {
	Name      string `json:"name"`
	Date      string `json:"date"`
	Namespace string `json:"namespace"`
}

type PodName struct {
	Name      string `json:"name"`
	Date      string `json:"date"`
	Namespace string `json:"namespace"`
}
