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

// InstanceProfileParameters defines the desired state of InstanceProfile
type InstanceProfileParameters struct {
	// The name of the instance profile to create.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	// +kubebuilder:validation:Required
	InstanceProfileName *string `json:"instanceProfileName"`
	// The path to the instance profile. For more information about paths, see IAM
	// Identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html)
	// in the IAM User Guide.
	//
	// This parameter is optional. If it is not included, it defaults to a slash
	// (/).
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of either a forward slash (/) by itself
	// or a string that must begin and end with forward slashes. In addition, it
	// can contain any ASCII character from the ! (\u0021) through the DEL character
	// (\u007F), including most punctuation characters, digits, and upper and lowercased
	// letters.
	Path *string `json:"path,omitempty"`
	// A list of tags that you want to attach to the newly created IAM instance
	// profile. Each tag consists of a key name and an associated value. For more
	// information about tagging, see Tagging IAM resources (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html)
	// in the IAM User Guide.
	//
	// If any one of the tags is invalid or if you exceed the allowed maximum number
	// of tags, then the entire request fails and the resource is not created.
	Tags                            []*Tag `json:"tags,omitempty"`
	CustomInstanceProfileParameters `json:",inline"`
}

// InstanceProfileSpec defines the desired state of InstanceProfile
type InstanceProfileSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       InstanceProfileParameters `json:"forProvider"`
}

// InstanceProfileObservation defines the observed state of InstanceProfile
type InstanceProfileObservation struct {
	// The Amazon Resource Name (ARN) specifying the instance profile. For more
	// information about ARNs and how to use them in policies, see IAM identifiers
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html)
	// in the IAM User Guide.
	ARN *string `json:"arn,omitempty"`
	// The date when the instance profile was created.
	CreateDate *metav1.Time `json:"createDate,omitempty"`
	// The stable and unique string identifying the instance profile. For more information
	// about IDs, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html)
	// in the IAM User Guide.
	InstanceProfileID *string `json:"instanceProfileID,omitempty"`
	// The role associated with the instance profile.
	Roles []*Role `json:"roles,omitempty"`
}

// InstanceProfileStatus defines the observed state of InstanceProfile.
type InstanceProfileStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          InstanceProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceProfile is the Schema for the InstanceProfiles API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type InstanceProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceProfileSpec   `json:"spec"`
	Status            InstanceProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceProfileList contains a list of InstanceProfiles
type InstanceProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceProfile `json:"items"`
}

// Repository type metadata.
var (
	InstanceProfileKind             = "InstanceProfile"
	InstanceProfileGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstanceProfileKind}.String()
	InstanceProfileKindAPIVersion   = InstanceProfileKind + "." + GroupVersion.String()
	InstanceProfileGroupVersionKind = GroupVersion.WithKind(InstanceProfileKind)
)

func init() {
	SchemeBuilder.Register(&InstanceProfile{}, &InstanceProfileList{})
}
