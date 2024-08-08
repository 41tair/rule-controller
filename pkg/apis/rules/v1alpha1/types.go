/*
Copyright 2024.

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

type Trigger struct {
	PodStatus *PodStatus `json:"podStatus"`
}

type Executor struct {
	Print *Print `json:"print,omitempty"`
	Taint *Taint `json:"taint,omitempty"`
}

type PodStatus struct {
	ExitCode int       `json:"exitCode"`
	Selector *Selector `json:"selector"`
}

type Selector struct {
	Name          string `json:"name"`
	Namespace     string `json:"namespace"`
	LabelSelector string `json:"labelSelector"`
}

type Print struct {
	Message string `json:"message"`
}

type Taint struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Effect string `json:"effect"`
	Node   Node   `json:"node"`
}

type Node struct {
	Name string `json:"name"`
}

// RulesSpec defines the desired state of Rules
type RulesSpec struct {
	Trigger   Trigger    `json:"trigger"`
	Executors []Executor `json:"executors"`
}

// RulesStatus defines the observed state of Rules
type RulesStatus struct {
	LatestTime   string `json:"latest"`
	SucceedCount int    `json:"succeedCount"`
	FailedCount  int    `json:"failedCount"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Rules is the Schema for the rules API
type Rules struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RulesSpec   `json:"spec,omitempty"`
	Status RulesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// RulesList contains a list of Rules
type RulesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Rules `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Rules{}, &RulesList{})
}
