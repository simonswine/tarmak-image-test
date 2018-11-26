/*
Copyright 2018 The Kubernetes Authors.

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

// TarmakVersion defines the Tarmak version as tag or branch name (e.g master, release-0.5, 0.5.0)
type TarmakVersion string

// TestState returns the overall result of the test
type TestState string

// ConditionStatus returns if a condition is active
type ConditionStatus string

const (
	// ConditionTrue condition is active
	ConditionTrue ConditionStatus = "True"
	// ConditionFalse condition is not active
	ConditionFalse ConditionStatus = "False"
	// ConditionUnknown condition status is not known
	ConditionUnknown ConditionStatus = "Unknown"
)

// TestConditionType specifies what problems have occured during tests
type TestConditionType string

const (
	// TestPrometheusAlerts means that there are alerts in prometheus after deploy
	TestPrometheusAlerts TestConditionType = "PrometheusAlerts"
	// TestOverlayNetworkBroken means that there are problems with the overlay network
	TestOverlayNetworkBroken TestConditionType = "OverlayNetworkBroken"
	// TestTerraformError means that there were errors applying the terraform code
	TestTerraformError TestConditionType = "TerraformError"
)

// TestCondition shows conditions that are currently active
type TestCondition struct {
	// Type of test condition.
	Type TestConditionType `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status ConditionStatus `json:"status"`
	// (brief) reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty"`
	// Human readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty"`
}

const (
	// TestStateSuccess overall test result is successful
	TestStateSuccess = TestState("success")
	// TestStateFailed overall test result is failed
	TestStateFailed = TestState("failed")
	// TestStatePending test is still running
	TestStatePending = TestState("pending")
)

// ImageTestSpec defines the desired state of ImageTest
type ImageTestSpec struct {
	Destination ImageTarmakSpec  `json:",inline"`
	Source      *ImageTarmakSpec `json:"source,omitempty"`
}

// ImageTarmakSpec defines the combination of base image and Tarmak version
type ImageTarmakSpec struct {
	Image         Image         `json:"image,omitempty"`
	TarmakVersion TarmakVersion `json:"tarmakVersion,omitempty"`
}

// ImageTestStatus defines the observed state of ImageTest
type ImageTestStatus struct {
	State          TestState       `json:"state,omitempty"`
	TestConditions []TestCondition `json:"testConditions,omitempty"`
}

// Image defines the base image to test TODO: Should be imported from "github.com/jetstack/tarmak/pkg/apis/tarmak/v1alpha1"
type Image struct {
	Name     string `json:"name,omitempty"`
	Location string `json:"location,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ImageTest is the Schema for the imagetests API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type ImageTest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ImageTestSpec   `json:"spec,omitempty"`
	Status ImageTestStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ImageTestList contains a list of ImageTest
type ImageTestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImageTest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ImageTest{}, &ImageTestList{})
}
