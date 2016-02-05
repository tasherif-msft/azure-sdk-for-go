// Package search implements the Azure ARM search service API version
// 2015-02-28.
//
// client that can be used to manage Azure Search services and API keys..
package search

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.14.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
)

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Failed specifies the failed state for provisioning state.
	Failed ProvisioningState = "failed"
	// Provisioning specifies the provisioning state for provisioning state.
	Provisioning ProvisioningState = "provisioning"
	// Succeeded specifies the succeeded state for provisioning state.
	Succeeded ProvisioningState = "succeeded"
)

// ServiceStatus enumerates the values for service status.
type ServiceStatus string

const (
	// ServiceStatusDegraded specifies the service status degraded state for
	// service status.
	ServiceStatusDegraded ServiceStatus = "degraded"
	// ServiceStatusDeleting specifies the service status deleting state for
	// service status.
	ServiceStatusDeleting ServiceStatus = "deleting"
	// ServiceStatusDisabled specifies the service status disabled state for
	// service status.
	ServiceStatusDisabled ServiceStatus = "disabled"
	// ServiceStatusError specifies the service status error state for service
	// status.
	ServiceStatusError ServiceStatus = "error"
	// ServiceStatusProvisioning specifies the service status provisioning
	// state for service status.
	ServiceStatusProvisioning ServiceStatus = "provisioning"
	// ServiceStatusRunning specifies the service status running state for
	// service status.
	ServiceStatusRunning ServiceStatus = "running"
)

// SkuType enumerates the values for sku type.
type SkuType string

const (
	// Free specifies the free state for sku type.
	Free SkuType = "free"
	// Standard specifies the standard state for sku type.
	Standard SkuType = "standard"
	// Standard2 specifies the standard 2 state for sku type.
	Standard2 SkuType = "standard2"
)

// AdminKeyResult is response containing the primary and secondary API keys
// for a given Azure Search service.
type AdminKeyResult struct {
	autorest.Response `json:"-"`
	PrimaryKey        *string `json:"primaryKey,omitempty"`
	SecondaryKey      *string `json:"secondaryKey,omitempty"`
}

// ListQueryKeysResult is response containing the query API keys for a given
// Azure Search service.
type ListQueryKeysResult struct {
	autorest.Response `json:"-"`
	Value             *[]QueryKey `json:"value,omitempty"`
}

// QueryKey is describes an API key for a given Azure Search service that has
// permissions for query operations only.
type QueryKey struct {
	Name *string `json:"name,omitempty"`
	Key  *string `json:"key,omitempty"`
}

// Resource is
type Resource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}

// ServiceCreateOrUpdateParameters is properties that describe an Azure Search
// service.
type ServiceCreateOrUpdateParameters struct {
	Location   *string             `json:"location,omitempty"`
	Tags       *map[string]*string `json:"tags,omitempty"`
	Properties *ServiceProperties  `json:"properties,omitempty"`
}

// ServiceListResult is response containing a list of Azure Search services
// for a given resource group.
type ServiceListResult struct {
	autorest.Response `json:"-"`
	Value             *[]ServiceResource `json:"value,omitempty"`
}

// ServiceProperties is defines properties of an Azure Search service that can
// be modified.
type ServiceProperties struct {
	Sku            *Sku `json:"sku,omitempty"`
	ReplicaCount   *int `json:"replicaCount,omitempty"`
	PartitionCount *int `json:"partitionCount,omitempty"`
}

// ServiceReadableProperties is defines all the properties of an Azure Search
// service.
type ServiceReadableProperties struct {
	Status            ServiceStatus     `json:"status,omitempty"`
	StatusDetails     *string           `json:"statusDetails,omitempty"`
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
	Sku               *Sku              `json:"sku,omitempty"`
	ReplicaCount      *int              `json:"replicaCount,omitempty"`
	PartitionCount    *int              `json:"partitionCount,omitempty"`
}

// ServiceResource is describes an Azure Search service and its current state.
type ServiceResource struct {
	autorest.Response `json:"-"`
	ID                *string                    `json:"id,omitempty"`
	Name              *string                    `json:"name,omitempty"`
	Location          *string                    `json:"location,omitempty"`
	Tags              *map[string]*string        `json:"tags,omitempty"`
	Properties        *ServiceReadableProperties `json:"properties,omitempty"`
}

// Sku is defines the SKU of an Azure Search Service, which determines price
// tier and capacity limits.
type Sku struct {
	Name SkuType `json:"name,omitempty"`
}

// SubResource is
type SubResource struct {
	ID *string `json:"id,omitempty"`
}
