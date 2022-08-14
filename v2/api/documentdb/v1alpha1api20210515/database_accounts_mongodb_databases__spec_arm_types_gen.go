// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of DatabaseAccountsMongodbDatabases_Spec. Use v1beta20210515.DatabaseAccountsMongodbDatabases_Spec instead
type DatabaseAccountsMongodbDatabases_SpecARM struct {
	Location   *string                                   `json:"location,omitempty"`
	Name       string                                    `json:"name,omitempty"`
	Properties *MongoDBDatabaseCreateUpdatePropertiesARM `json:"properties,omitempty"`
	Tags       map[string]string                         `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccountsMongodbDatabases_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (databases DatabaseAccountsMongodbDatabases_SpecARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (databases *DatabaseAccountsMongodbDatabases_SpecARM) GetName() string {
	return databases.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases"
func (databases *DatabaseAccountsMongodbDatabases_SpecARM) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases"
}

// Deprecated version of MongoDBDatabaseCreateUpdateProperties. Use v1beta20210515.MongoDBDatabaseCreateUpdateProperties instead
type MongoDBDatabaseCreateUpdatePropertiesARM struct {
	Options  *CreateUpdateOptionsARM     `json:"options,omitempty"`
	Resource *MongoDBDatabaseResourceARM `json:"resource,omitempty"`
}

// Deprecated version of MongoDBDatabaseResource. Use v1beta20210515.MongoDBDatabaseResource instead
type MongoDBDatabaseResourceARM struct {
	Id *string `json:"id,omitempty"`
}