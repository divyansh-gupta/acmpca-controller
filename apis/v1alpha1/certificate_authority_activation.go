// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CertificateAuthorityActivationSpec defines the desired state of CertificateAuthorityActivation.
type CertificateAuthorityActivationSpec struct {

	// +kubebuilder:validation:Required
	Certificate *ackv1alpha1.SecretKeyReference `json:"certificate"`
	// The Amazon Resource Name (ARN) that was returned when you called CreateCertificateAuthority
	// (https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html).
	// This must be of the form:
	//
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	CertificateAuthorityARN *string                                  `json:"certificateAuthorityARN,omitempty"`
	CertificateAuthorityRef *ackv1alpha1.AWSResourceReferenceWrapper `json:"certificateAuthorityRef,omitempty"`
	// A PEM-encoded file that contains all of your certificates, other than the
	// certificate you're importing, chaining up to your root CA. Your Amazon Web
	// Services Private CA-hosted or on-premises root certificate is the last in
	// the chain, and each certificate in the chain signs the one preceding.
	//
	// This parameter must be supplied when you import a subordinate CA. When you
	// import a root CA, there is no chain.
	CertificateChain []byte `json:"certificateChain,omitempty"`
}

// CertificateAuthorityActivationStatus defines the observed state of CertificateAuthorityActivation
type CertificateAuthorityActivationStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty"`
}

// CertificateAuthorityActivation is the Schema for the CertificateAuthorityActivations API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type CertificateAuthorityActivation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CertificateAuthorityActivationSpec   `json:"spec,omitempty"`
	Status            CertificateAuthorityActivationStatus `json:"status,omitempty"`
}

// CertificateAuthorityActivationList contains a list of CertificateAuthorityActivation
// +kubebuilder:object:root=true
type CertificateAuthorityActivationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CertificateAuthorityActivation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CertificateAuthorityActivation{}, &CertificateAuthorityActivationList{})
}
