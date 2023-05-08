// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220701storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=network.azure.com,resources=privateendpoints,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={privateendpoints/status,privateendpoints/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20220701.PrivateEndpoint
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2022-07-01/privateEndpoint.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}
type PrivateEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateEndpoint_Spec                                       `json:"spec,omitempty"`
	Status            PrivateEndpoint_STATUS_PrivateEndpoint_SubResourceEmbedded `json:"status,omitempty"`
}

var _ conditions.Conditioner = &PrivateEndpoint{}

// GetConditions returns the conditions of the resource
func (endpoint *PrivateEndpoint) GetConditions() conditions.Conditions {
	return endpoint.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (endpoint *PrivateEndpoint) SetConditions(conditions conditions.Conditions) {
	endpoint.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &PrivateEndpoint{}

// AzureName returns the Azure name of the resource
func (endpoint *PrivateEndpoint) AzureName() string {
	return endpoint.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-07-01"
func (endpoint PrivateEndpoint) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (endpoint *PrivateEndpoint) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (endpoint *PrivateEndpoint) GetSpec() genruntime.ConvertibleSpec {
	return &endpoint.Spec
}

// GetStatus returns the status of this resource
func (endpoint *PrivateEndpoint) GetStatus() genruntime.ConvertibleStatus {
	return &endpoint.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/privateEndpoints"
func (endpoint *PrivateEndpoint) GetType() string {
	return "Microsoft.Network/privateEndpoints"
}

// NewEmptyStatus returns a new empty (blank) status
func (endpoint *PrivateEndpoint) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &PrivateEndpoint_STATUS_PrivateEndpoint_SubResourceEmbedded{}
}

// Owner returns the ResourceReference of the owner
func (endpoint *PrivateEndpoint) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(endpoint.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  endpoint.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (endpoint *PrivateEndpoint) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*PrivateEndpoint_STATUS_PrivateEndpoint_SubResourceEmbedded); ok {
		endpoint.Status = *st
		return nil
	}

	// Convert status to required version
	var st PrivateEndpoint_STATUS_PrivateEndpoint_SubResourceEmbedded
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	endpoint.Status = st
	return nil
}

// Hub marks that this PrivateEndpoint is the hub type for conversion
func (endpoint *PrivateEndpoint) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (endpoint *PrivateEndpoint) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: endpoint.Spec.OriginalVersion,
		Kind:    "PrivateEndpoint",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220701.PrivateEndpoint
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2022-07-01/privateEndpoint.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}
type PrivateEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateEndpoint `json:"items"`
}

// Storage version of v1api20220701.PrivateEndpoint_Spec
type PrivateEndpoint_Spec struct {
	ApplicationSecurityGroups []ApplicationSecurityGroupSpec_PrivateEndpoint_SubResourceEmbedded `json:"applicationSecurityGroups,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                           string                           `json:"azureName,omitempty"`
	CustomNetworkInterfaceName          *string                          `json:"customNetworkInterfaceName,omitempty"`
	ExtendedLocation                    *ExtendedLocation                `json:"extendedLocation,omitempty"`
	IpConfigurations                    []PrivateEndpointIPConfiguration `json:"ipConfigurations,omitempty"`
	Location                            *string                          `json:"location,omitempty"`
	ManualPrivateLinkServiceConnections []PrivateLinkServiceConnection   `json:"manualPrivateLinkServiceConnections,omitempty"`
	OriginalVersion                     string                           `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner                         *genruntime.KnownResourceReference          `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PrivateLinkServiceConnections []PrivateLinkServiceConnection              `json:"privateLinkServiceConnections,omitempty"`
	PropertyBag                   genruntime.PropertyBag                      `json:"$propertyBag,omitempty"`
	Subnet                        *Subnet_PrivateEndpoint_SubResourceEmbedded `json:"subnet,omitempty"`
	Tags                          map[string]string                           `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &PrivateEndpoint_Spec{}

// ConvertSpecFrom populates our PrivateEndpoint_Spec from the provided source
func (endpoint *PrivateEndpoint_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == endpoint {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(endpoint)
}

// ConvertSpecTo populates the provided destination from our PrivateEndpoint_Spec
func (endpoint *PrivateEndpoint_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == endpoint {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(endpoint)
}

// Storage version of v1api20220701.PrivateEndpoint_STATUS_PrivateEndpoint_SubResourceEmbedded
// Private endpoint resource.
type PrivateEndpoint_STATUS_PrivateEndpoint_SubResourceEmbedded struct {
	ApplicationSecurityGroups           []ApplicationSecurityGroup_STATUS_PrivateEndpoint_SubResourceEmbedded `json:"applicationSecurityGroups,omitempty"`
	Conditions                          []conditions.Condition                                                `json:"conditions,omitempty"`
	CustomDnsConfigs                    []CustomDnsConfigPropertiesFormat_STATUS                              `json:"customDnsConfigs,omitempty"`
	CustomNetworkInterfaceName          *string                                                               `json:"customNetworkInterfaceName,omitempty"`
	Etag                                *string                                                               `json:"etag,omitempty"`
	ExtendedLocation                    *ExtendedLocation_STATUS                                              `json:"extendedLocation,omitempty"`
	Id                                  *string                                                               `json:"id,omitempty"`
	IpConfigurations                    []PrivateEndpointIPConfiguration_STATUS                               `json:"ipConfigurations,omitempty"`
	Location                            *string                                                               `json:"location,omitempty"`
	ManualPrivateLinkServiceConnections []PrivateLinkServiceConnection_STATUS                                 `json:"manualPrivateLinkServiceConnections,omitempty"`
	Name                                *string                                                               `json:"name,omitempty"`
	NetworkInterfaces                   []NetworkInterface_STATUS_PrivateEndpoint_SubResourceEmbedded         `json:"networkInterfaces,omitempty"`
	PrivateLinkServiceConnections       []PrivateLinkServiceConnection_STATUS                                 `json:"privateLinkServiceConnections,omitempty"`
	PropertyBag                         genruntime.PropertyBag                                                `json:"$propertyBag,omitempty"`
	ProvisioningState                   *string                                                               `json:"provisioningState,omitempty"`
	Subnet                              *Subnet_STATUS_PrivateEndpoint_SubResourceEmbedded                    `json:"subnet,omitempty"`
	Tags                                map[string]string                                                     `json:"tags,omitempty"`
	Type                                *string                                                               `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &PrivateEndpoint_STATUS_PrivateEndpoint_SubResourceEmbedded{}

// ConvertStatusFrom populates our PrivateEndpoint_STATUS_PrivateEndpoint_SubResourceEmbedded from the provided source
func (embedded *PrivateEndpoint_STATUS_PrivateEndpoint_SubResourceEmbedded) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == embedded {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(embedded)
}

// ConvertStatusTo populates the provided destination from our PrivateEndpoint_STATUS_PrivateEndpoint_SubResourceEmbedded
func (embedded *PrivateEndpoint_STATUS_PrivateEndpoint_SubResourceEmbedded) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == embedded {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(embedded)
}

// Storage version of v1api20220701.ApplicationSecurityGroup_STATUS_PrivateEndpoint_SubResourceEmbedded
// An application security group in a resource group.
type ApplicationSecurityGroup_STATUS_PrivateEndpoint_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220701.ApplicationSecurityGroupSpec_PrivateEndpoint_SubResourceEmbedded
// An application security group in a resource group.
type ApplicationSecurityGroupSpec_PrivateEndpoint_SubResourceEmbedded struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// Reference: Resource ID.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

// Storage version of v1api20220701.CustomDnsConfigPropertiesFormat_STATUS
// Contains custom Dns resolution configuration from customer.
type CustomDnsConfigPropertiesFormat_STATUS struct {
	Fqdn        *string                `json:"fqdn,omitempty"`
	IpAddresses []string               `json:"ipAddresses,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220701.ExtendedLocation
// ExtendedLocation complex type.
type ExtendedLocation struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1api20220701.ExtendedLocation_STATUS
// ExtendedLocation complex type.
type ExtendedLocation_STATUS struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1api20220701.NetworkInterface_STATUS_PrivateEndpoint_SubResourceEmbedded
// A network interface in a resource group.
type NetworkInterface_STATUS_PrivateEndpoint_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220701.PrivateEndpointIPConfiguration
// An IP Configuration of the private endpoint.
type PrivateEndpointIPConfiguration struct {
	GroupId          *string                `json:"groupId,omitempty"`
	MemberName       *string                `json:"memberName,omitempty"`
	Name             *string                `json:"name,omitempty"`
	PrivateIPAddress *string                `json:"privateIPAddress,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220701.PrivateEndpointIPConfiguration_STATUS
// An IP Configuration of the private endpoint.
type PrivateEndpointIPConfiguration_STATUS struct {
	Etag             *string                `json:"etag,omitempty"`
	GroupId          *string                `json:"groupId,omitempty"`
	MemberName       *string                `json:"memberName,omitempty"`
	Name             *string                `json:"name,omitempty"`
	PrivateIPAddress *string                `json:"privateIPAddress,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type             *string                `json:"type,omitempty"`
}

// Storage version of v1api20220701.PrivateLinkServiceConnection
// PrivateLinkServiceConnection resource.
type PrivateLinkServiceConnection struct {
	GroupIds                          []string                           `json:"groupIds,omitempty"`
	Name                              *string                            `json:"name,omitempty"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState `json:"privateLinkServiceConnectionState,omitempty"`

	// PrivateLinkServiceReference: The resource id of private link service.
	PrivateLinkServiceReference *genruntime.ResourceReference `armReference:"PrivateLinkServiceId" json:"privateLinkServiceReference,omitempty"`
	PropertyBag                 genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	RequestMessage              *string                       `json:"requestMessage,omitempty"`
}

// Storage version of v1api20220701.PrivateLinkServiceConnection_STATUS
// PrivateLinkServiceConnection resource.
type PrivateLinkServiceConnection_STATUS struct {
	Etag                              *string                                   `json:"etag,omitempty"`
	GroupIds                          []string                                  `json:"groupIds,omitempty"`
	Id                                *string                                   `json:"id,omitempty"`
	Name                              *string                                   `json:"name,omitempty"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState_STATUS `json:"privateLinkServiceConnectionState,omitempty"`
	PrivateLinkServiceId              *string                                   `json:"privateLinkServiceId,omitempty"`
	PropertyBag                       genruntime.PropertyBag                    `json:"$propertyBag,omitempty"`
	ProvisioningState                 *string                                   `json:"provisioningState,omitempty"`
	RequestMessage                    *string                                   `json:"requestMessage,omitempty"`
	Type                              *string                                   `json:"type,omitempty"`
}

// Storage version of v1api20220701.Subnet_PrivateEndpoint_SubResourceEmbedded
// Subnet in a virtual network resource.
type Subnet_PrivateEndpoint_SubResourceEmbedded struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// Reference: Resource ID.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

// Storage version of v1api20220701.Subnet_STATUS_PrivateEndpoint_SubResourceEmbedded
// Subnet in a virtual network resource.
type Subnet_STATUS_PrivateEndpoint_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220701.PrivateLinkServiceConnectionState
// A collection of information about the state of the connection between service consumer and provider.
type PrivateLinkServiceConnectionState struct {
	ActionsRequired *string                `json:"actionsRequired,omitempty"`
	Description     *string                `json:"description,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Status          *string                `json:"status,omitempty"`
}

// Storage version of v1api20220701.PrivateLinkServiceConnectionState_STATUS
// A collection of information about the state of the connection between service consumer and provider.
type PrivateLinkServiceConnectionState_STATUS struct {
	ActionsRequired *string                `json:"actionsRequired,omitempty"`
	Description     *string                `json:"description,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Status          *string                `json:"status,omitempty"`
}

func init() {
	SchemeBuilder.Register(&PrivateEndpoint{}, &PrivateEndpointList{})
}