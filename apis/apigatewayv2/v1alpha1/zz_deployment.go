/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// DeploymentParameters defines the desired state of Deployment
type DeploymentParameters struct {
	// Region is which region the Deployment will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`

	Description *string `json:"description,omitempty"`

	StageName                  *string `json:"stageName,omitempty"`
	CustomDeploymentParameters `json:",inline"`
}

// DeploymentSpec defines the desired state of Deployment
type DeploymentSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DeploymentParameters `json:"forProvider"`
}

// DeploymentObservation defines the observed state of Deployment
type DeploymentObservation struct {
	AutoDeployed *bool `json:"autoDeployed,omitempty"`

	CreatedDate *metav1.Time `json:"createdDate,omitempty"`

	DeploymentID *string `json:"deploymentID,omitempty"`

	DeploymentStatus *string `json:"deploymentStatus,omitempty"`

	DeploymentStatusMessage *string `json:"deploymentStatusMessage,omitempty"`
}

// DeploymentStatus defines the observed state of Deployment.
type DeploymentStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DeploymentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Deployment is the Schema for the Deployments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Deployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DeploymentSpec   `json:"spec"`
	Status            DeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeploymentList contains a list of Deployments
type DeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Deployment `json:"items"`
}

// Repository type metadata.
var (
	DeploymentKind             = "Deployment"
	DeploymentGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DeploymentKind}.String()
	DeploymentKindAPIVersion   = DeploymentKind + "." + GroupVersion.String()
	DeploymentGroupVersionKind = GroupVersion.WithKind(DeploymentKind)
)

func init() {
	SchemeBuilder.Register(&Deployment{}, &DeploymentList{})
}
