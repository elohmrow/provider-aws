/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ClusterEndpointObservation struct {

	// The Neptune Cluster Endpoint Amazon Resource Name (ARN).
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The DNS address of the endpoint.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The Neptune Cluster Endpoint Identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ClusterEndpointParameters struct {

	// The DB cluster identifier of the DB cluster associated with the endpoint.
	// +crossplane:generate:reference:type=Cluster
	// +kubebuilder:validation:Optional
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier,omitempty"`

	// Reference to a Cluster to populate clusterIdentifier.
	// +kubebuilder:validation:Optional
	ClusterIdentifierRef *v1.Reference `json:"clusterIdentifierRef,omitempty" tf:"-"`

	// Selector for a Cluster to populate clusterIdentifier.
	// +kubebuilder:validation:Optional
	ClusterIdentifierSelector *v1.Selector `json:"clusterIdentifierSelector,omitempty" tf:"-"`

	// The type of the endpoint. One of: READER, WRITER, ANY.
	// +kubebuilder:validation:Required
	EndpointType *string `json:"endpointType" tf:"endpoint_type,omitempty"`

	// List of DB instance identifiers that aren't part of the custom endpoint group. All other eligible instances are reachable through the custom endpoint. Only relevant if the list of static members is empty.
	// +kubebuilder:validation:Optional
	ExcludedMembers []*string `json:"excludedMembers,omitempty" tf:"excluded_members,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// List of DB instance identifiers that are part of the custom endpoint group.
	// +kubebuilder:validation:Optional
	StaticMembers []*string `json:"staticMembers,omitempty" tf:"static_members,omitempty"`

	// A map of tags to assign to the Neptune cluster. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ClusterEndpointSpec defines the desired state of ClusterEndpoint
type ClusterEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterEndpointParameters `json:"forProvider"`
}

// ClusterEndpointStatus defines the observed state of ClusterEndpoint.
type ClusterEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterEndpoint is the Schema for the ClusterEndpoints API. Provides an Neptune Cluster Endpoint Resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ClusterEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterEndpointSpec   `json:"spec"`
	Status            ClusterEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterEndpointList contains a list of ClusterEndpoints
type ClusterEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterEndpoint `json:"items"`
}

// Repository type metadata.
var (
	ClusterEndpoint_Kind             = "ClusterEndpoint"
	ClusterEndpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClusterEndpoint_Kind}.String()
	ClusterEndpoint_KindAPIVersion   = ClusterEndpoint_Kind + "." + CRDGroupVersion.String()
	ClusterEndpoint_GroupVersionKind = CRDGroupVersion.WithKind(ClusterEndpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&ClusterEndpoint{}, &ClusterEndpointList{})
}
