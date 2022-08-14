// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210301

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of RedisEnterprise_Spec. Use v1beta20210301.RedisEnterprise_Spec instead
type RedisEnterprise_SpecARM struct {
	Location   *string               `json:"location,omitempty"`
	Name       string                `json:"name,omitempty"`
	Properties *ClusterPropertiesARM `json:"properties,omitempty"`
	Sku        *SkuARM               `json:"sku,omitempty"`
	Tags       map[string]string     `json:"tags,omitempty"`
	Zones      []string              `json:"zones,omitempty"`
}

var _ genruntime.ARMResourceSpec = &RedisEnterprise_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-03-01"
func (enterprise RedisEnterprise_SpecARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (enterprise *RedisEnterprise_SpecARM) GetName() string {
	return enterprise.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redisEnterprise"
func (enterprise *RedisEnterprise_SpecARM) GetType() string {
	return "Microsoft.Cache/redisEnterprise"
}

// Deprecated version of ClusterProperties. Use v1beta20210301.ClusterProperties instead
type ClusterPropertiesARM struct {
	MinimumTlsVersion *ClusterPropertiesMinimumTlsVersion `json:"minimumTlsVersion,omitempty"`
}

// Deprecated version of Sku. Use v1beta20210301.Sku instead
type SkuARM struct {
	Capacity *int     `json:"capacity,omitempty"`
	Name     *SkuName `json:"name,omitempty"`
}

// Deprecated version of SkuName. Use v1beta20210301.SkuName instead
// +kubebuilder:validation:Enum={"Enterprise_E10","Enterprise_E100","Enterprise_E20","Enterprise_E50","EnterpriseFlash_F1500","EnterpriseFlash_F300","EnterpriseFlash_F700"}
type SkuName string

const (
	SkuNameEnterpriseE10        = SkuName("Enterprise_E10")
	SkuNameEnterpriseE100       = SkuName("Enterprise_E100")
	SkuNameEnterpriseE20        = SkuName("Enterprise_E20")
	SkuNameEnterpriseE50        = SkuName("Enterprise_E50")
	SkuNameEnterpriseFlashF1500 = SkuName("EnterpriseFlash_F1500")
	SkuNameEnterpriseFlashF300  = SkuName("EnterpriseFlash_F300")
	SkuNameEnterpriseFlashF700  = SkuName("EnterpriseFlash_F700")
)