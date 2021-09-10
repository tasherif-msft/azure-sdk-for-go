//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package videoanalyzer

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/videoanalyzer/mgmt/2021-05-01-preview/videoanalyzer"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AccessPolicyEccAlgo = original.AccessPolicyEccAlgo

const (
	AccessPolicyEccAlgoES256 AccessPolicyEccAlgo = original.AccessPolicyEccAlgoES256
	AccessPolicyEccAlgoES384 AccessPolicyEccAlgo = original.AccessPolicyEccAlgoES384
	AccessPolicyEccAlgoES512 AccessPolicyEccAlgo = original.AccessPolicyEccAlgoES512
)

type AccessPolicyRole = original.AccessPolicyRole

const (
	AccessPolicyRoleReader AccessPolicyRole = original.AccessPolicyRoleReader
)

type AccessPolicyRsaAlgo = original.AccessPolicyRsaAlgo

const (
	AccessPolicyRsaAlgoRS256 AccessPolicyRsaAlgo = original.AccessPolicyRsaAlgoRS256
	AccessPolicyRsaAlgoRS384 AccessPolicyRsaAlgo = original.AccessPolicyRsaAlgoRS384
	AccessPolicyRsaAlgoRS512 AccessPolicyRsaAlgo = original.AccessPolicyRsaAlgoRS512
)

type AccountEncryptionKeyType = original.AccountEncryptionKeyType

const (
	AccountEncryptionKeyTypeCustomerKey AccountEncryptionKeyType = original.AccountEncryptionKeyTypeCustomerKey
	AccountEncryptionKeyTypeSystemKey   AccountEncryptionKeyType = original.AccountEncryptionKeyTypeSystemKey
)

type ActionType = original.ActionType

const (
	ActionTypeInternal ActionType = original.ActionTypeInternal
)

type CheckNameAvailabilityReason = original.CheckNameAvailabilityReason

const (
	CheckNameAvailabilityReasonAlreadyExists CheckNameAvailabilityReason = original.CheckNameAvailabilityReasonAlreadyExists
	CheckNameAvailabilityReasonInvalid       CheckNameAvailabilityReason = original.CheckNameAvailabilityReasonInvalid
)

type CreatedByType = original.CreatedByType

const (
	CreatedByTypeApplication     CreatedByType = original.CreatedByTypeApplication
	CreatedByTypeKey             CreatedByType = original.CreatedByTypeKey
	CreatedByTypeManagedIdentity CreatedByType = original.CreatedByTypeManagedIdentity
	CreatedByTypeUser            CreatedByType = original.CreatedByTypeUser
)

type MetricAggregationType = original.MetricAggregationType

const (
	MetricAggregationTypeAverage MetricAggregationType = original.MetricAggregationTypeAverage
	MetricAggregationTypeCount   MetricAggregationType = original.MetricAggregationTypeCount
	MetricAggregationTypeTotal   MetricAggregationType = original.MetricAggregationTypeTotal
)

type MetricUnit = original.MetricUnit

const (
	MetricUnitBytes        MetricUnit = original.MetricUnitBytes
	MetricUnitCount        MetricUnit = original.MetricUnitCount
	MetricUnitMilliseconds MetricUnit = original.MetricUnitMilliseconds
)

type Type = original.Type

const (
	TypeAuthenticationBase                      Type = original.TypeAuthenticationBase
	TypeMicrosoftVideoAnalyzerJwtAuthentication Type = original.TypeMicrosoftVideoAnalyzerJwtAuthentication
)

type TypeBasicTokenKey = original.TypeBasicTokenKey

const (
	TypeBasicTokenKeyTypeMicrosoftVideoAnalyzerEccTokenKey TypeBasicTokenKey = original.TypeBasicTokenKeyTypeMicrosoftVideoAnalyzerEccTokenKey
	TypeBasicTokenKeyTypeMicrosoftVideoAnalyzerRsaTokenKey TypeBasicTokenKey = original.TypeBasicTokenKeyTypeMicrosoftVideoAnalyzerRsaTokenKey
	TypeBasicTokenKeyTypeTokenKey                          TypeBasicTokenKey = original.TypeBasicTokenKeyTypeTokenKey
)

type VideoType = original.VideoType

const (
	VideoTypeArchive VideoType = original.VideoTypeArchive
)

type AccessPoliciesClient = original.AccessPoliciesClient
type AccessPolicyEntity = original.AccessPolicyEntity
type AccessPolicyEntityCollection = original.AccessPolicyEntityCollection
type AccessPolicyEntityCollectionIterator = original.AccessPolicyEntityCollectionIterator
type AccessPolicyEntityCollectionPage = original.AccessPolicyEntityCollectionPage
type AccessPolicyProperties = original.AccessPolicyProperties
type AccountEncryption = original.AccountEncryption
type AuthenticationBase = original.AuthenticationBase
type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type BasicAuthenticationBase = original.BasicAuthenticationBase
type BasicTokenKey = original.BasicTokenKey
type CheckNameAvailabilityRequest = original.CheckNameAvailabilityRequest
type CheckNameAvailabilityResponse = original.CheckNameAvailabilityResponse
type Collection = original.Collection
type EccTokenKey = original.EccTokenKey
type EdgeModuleEntity = original.EdgeModuleEntity
type EdgeModuleEntityCollection = original.EdgeModuleEntityCollection
type EdgeModuleEntityCollectionIterator = original.EdgeModuleEntityCollectionIterator
type EdgeModuleEntityCollectionPage = original.EdgeModuleEntityCollectionPage
type EdgeModuleProperties = original.EdgeModuleProperties
type EdgeModuleProvisioningToken = original.EdgeModuleProvisioningToken
type EdgeModulesClient = original.EdgeModulesClient
type Endpoint = original.Endpoint
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type Identity = original.Identity
type JwtAuthentication = original.JwtAuthentication
type KeyVaultProperties = original.KeyVaultProperties
type ListProvisioningTokenInput = original.ListProvisioningTokenInput
type LocationsClient = original.LocationsClient
type LogSpecification = original.LogSpecification
type MetricDimension = original.MetricDimension
type MetricSpecification = original.MetricSpecification
type Model = original.Model
type Operation = original.Operation
type OperationCollection = original.OperationCollection
type OperationDisplay = original.OperationDisplay
type OperationsClient = original.OperationsClient
type Properties = original.Properties
type PropertiesType = original.PropertiesType
type PropertiesUpdate = original.PropertiesUpdate
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type ResourceIdentity = original.ResourceIdentity
type RsaTokenKey = original.RsaTokenKey
type ServiceSpecification = original.ServiceSpecification
type StorageAccount = original.StorageAccount
type SyncStorageKeysInput = original.SyncStorageKeysInput
type SystemData = original.SystemData
type TokenClaim = original.TokenClaim
type TokenKey = original.TokenKey
type TrackedResource = original.TrackedResource
type Update = original.Update
type UserAssignedManagedIdentity = original.UserAssignedManagedIdentity
type VideoAnalyzersClient = original.VideoAnalyzersClient
type VideoEntity = original.VideoEntity
type VideoEntityCollection = original.VideoEntityCollection
type VideoEntityCollectionIterator = original.VideoEntityCollectionIterator
type VideoEntityCollectionPage = original.VideoEntityCollectionPage
type VideoFlags = original.VideoFlags
type VideoMediaInfo = original.VideoMediaInfo
type VideoProperties = original.VideoProperties
type VideoStreaming = original.VideoStreaming
type VideoStreamingToken = original.VideoStreamingToken
type VideosClient = original.VideosClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAccessPoliciesClient(subscriptionID string) AccessPoliciesClient {
	return original.NewAccessPoliciesClient(subscriptionID)
}
func NewAccessPoliciesClientWithBaseURI(baseURI string, subscriptionID string) AccessPoliciesClient {
	return original.NewAccessPoliciesClientWithBaseURI(baseURI, subscriptionID)
}
func NewAccessPolicyEntityCollectionIterator(page AccessPolicyEntityCollectionPage) AccessPolicyEntityCollectionIterator {
	return original.NewAccessPolicyEntityCollectionIterator(page)
}
func NewAccessPolicyEntityCollectionPage(cur AccessPolicyEntityCollection, getNextPage func(context.Context, AccessPolicyEntityCollection) (AccessPolicyEntityCollection, error)) AccessPolicyEntityCollectionPage {
	return original.NewAccessPolicyEntityCollectionPage(cur, getNextPage)
}
func NewEdgeModuleEntityCollectionIterator(page EdgeModuleEntityCollectionPage) EdgeModuleEntityCollectionIterator {
	return original.NewEdgeModuleEntityCollectionIterator(page)
}
func NewEdgeModuleEntityCollectionPage(cur EdgeModuleEntityCollection, getNextPage func(context.Context, EdgeModuleEntityCollection) (EdgeModuleEntityCollection, error)) EdgeModuleEntityCollectionPage {
	return original.NewEdgeModuleEntityCollectionPage(cur, getNextPage)
}
func NewEdgeModulesClient(subscriptionID string) EdgeModulesClient {
	return original.NewEdgeModulesClient(subscriptionID)
}
func NewEdgeModulesClientWithBaseURI(baseURI string, subscriptionID string) EdgeModulesClient {
	return original.NewEdgeModulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewLocationsClient(subscriptionID string) LocationsClient {
	return original.NewLocationsClient(subscriptionID)
}
func NewLocationsClientWithBaseURI(baseURI string, subscriptionID string) LocationsClient {
	return original.NewLocationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewVideoAnalyzersClient(subscriptionID string) VideoAnalyzersClient {
	return original.NewVideoAnalyzersClient(subscriptionID)
}
func NewVideoAnalyzersClientWithBaseURI(baseURI string, subscriptionID string) VideoAnalyzersClient {
	return original.NewVideoAnalyzersClientWithBaseURI(baseURI, subscriptionID)
}
func NewVideoEntityCollectionIterator(page VideoEntityCollectionPage) VideoEntityCollectionIterator {
	return original.NewVideoEntityCollectionIterator(page)
}
func NewVideoEntityCollectionPage(cur VideoEntityCollection, getNextPage func(context.Context, VideoEntityCollection) (VideoEntityCollection, error)) VideoEntityCollectionPage {
	return original.NewVideoEntityCollectionPage(cur, getNextPage)
}
func NewVideosClient(subscriptionID string) VideosClient {
	return original.NewVideosClient(subscriptionID)
}
func NewVideosClientWithBaseURI(baseURI string, subscriptionID string) VideosClient {
	return original.NewVideosClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAccessPolicyEccAlgoValues() []AccessPolicyEccAlgo {
	return original.PossibleAccessPolicyEccAlgoValues()
}
func PossibleAccessPolicyRoleValues() []AccessPolicyRole {
	return original.PossibleAccessPolicyRoleValues()
}
func PossibleAccessPolicyRsaAlgoValues() []AccessPolicyRsaAlgo {
	return original.PossibleAccessPolicyRsaAlgoValues()
}
func PossibleAccountEncryptionKeyTypeValues() []AccountEncryptionKeyType {
	return original.PossibleAccountEncryptionKeyTypeValues()
}
func PossibleActionTypeValues() []ActionType {
	return original.PossibleActionTypeValues()
}
func PossibleCheckNameAvailabilityReasonValues() []CheckNameAvailabilityReason {
	return original.PossibleCheckNameAvailabilityReasonValues()
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleMetricAggregationTypeValues() []MetricAggregationType {
	return original.PossibleMetricAggregationTypeValues()
}
func PossibleMetricUnitValues() []MetricUnit {
	return original.PossibleMetricUnitValues()
}
func PossibleTypeBasicTokenKeyValues() []TypeBasicTokenKey {
	return original.PossibleTypeBasicTokenKeyValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func PossibleVideoTypeValues() []VideoType {
	return original.PossibleVideoTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
