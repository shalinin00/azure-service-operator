// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

type VirtualNetworkPeering_StatusARM struct {
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the virtual network peering.
	Properties *VirtualNetworkPeeringPropertiesFormat_StatusARM `json:"properties,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

type VirtualNetworkPeeringPropertiesFormat_StatusARM struct {
	// AllowForwardedTraffic: Whether the forwarded traffic from the VMs in the local virtual network will be
	// allowed/disallowed in remote virtual network.
	AllowForwardedTraffic *bool `json:"allowForwardedTraffic,omitempty"`

	// AllowGatewayTransit: If gateway links can be used in remote virtual networking to link to this virtual network.
	AllowGatewayTransit *bool `json:"allowGatewayTransit,omitempty"`

	// AllowVirtualNetworkAccess: Whether the VMs in the local virtual network space would be able to access the VMs in remote
	// virtual network space.
	AllowVirtualNetworkAccess *bool `json:"allowVirtualNetworkAccess,omitempty"`

	// DoNotVerifyRemoteGateways: If we need to verify the provisioning state of the remote gateway.
	DoNotVerifyRemoteGateways *bool `json:"doNotVerifyRemoteGateways,omitempty"`

	// PeeringState: The status of the virtual network peering.
	PeeringState *VirtualNetworkPeeringPropertiesFormatStatusPeeringState `json:"peeringState,omitempty"`

	// ProvisioningState: The provisioning state of the virtual network peering resource.
	ProvisioningState *ProvisioningState_Status `json:"provisioningState,omitempty"`

	// RemoteAddressSpace: The reference to the remote virtual network address space.
	RemoteAddressSpace *AddressSpace_StatusARM `json:"remoteAddressSpace,omitempty"`

	// RemoteBgpCommunities: The reference to the remote virtual network's Bgp Communities.
	RemoteBgpCommunities *VirtualNetworkBgpCommunities_StatusARM `json:"remoteBgpCommunities,omitempty"`

	// RemoteVirtualNetwork: The reference to the remote virtual network. The remote virtual network can be in the same or
	// different region (preview). See here to register for the preview and learn more
	// (https://docs.microsoft.com/en-us/azure/virtual-network/virtual-network-create-peering).
	RemoteVirtualNetwork *SubResource_StatusARM `json:"remoteVirtualNetwork,omitempty"`

	// ResourceGuid: The resourceGuid property of the Virtual Network peering resource.
	ResourceGuid *string `json:"resourceGuid,omitempty"`

	// UseRemoteGateways: If remote gateways can be used on this virtual network. If the flag is set to true, and
	// allowGatewayTransit on remote peering is also true, virtual network will use gateways of remote virtual network for
	// transit. Only one peering can have this flag set to true. This flag cannot be set if virtual network already has a
	// gateway.
	UseRemoteGateways *bool `json:"useRemoteGateways,omitempty"`
}

type VirtualNetworkBgpCommunities_StatusARM struct {
	// RegionalCommunity: The BGP community associated with the region of the virtual network.
	RegionalCommunity *string `json:"regionalCommunity,omitempty"`

	// VirtualNetworkCommunity: The BGP community associated with the virtual network.
	VirtualNetworkCommunity *string `json:"virtualNetworkCommunity,omitempty"`
}

type VirtualNetworkPeeringPropertiesFormatStatusPeeringState string

const (
	VirtualNetworkPeeringPropertiesFormatStatusPeeringStateConnected    = VirtualNetworkPeeringPropertiesFormatStatusPeeringState("Connected")
	VirtualNetworkPeeringPropertiesFormatStatusPeeringStateDisconnected = VirtualNetworkPeeringPropertiesFormatStatusPeeringState("Disconnected")
	VirtualNetworkPeeringPropertiesFormatStatusPeeringStateInitiated    = VirtualNetworkPeeringPropertiesFormatStatusPeeringState("Initiated")
)