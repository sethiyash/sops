// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package timeseriesinsights

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/timeseriesinsights/mgmt/2017-11-15/timeseriesinsights"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AccessPolicyRole = original.AccessPolicyRole

const (
	Contributor AccessPolicyRole = original.Contributor
	Reader      AccessPolicyRole = original.Reader
)

type DataStringComparisonBehavior = original.DataStringComparisonBehavior

const (
	Ordinal           DataStringComparisonBehavior = original.Ordinal
	OrdinalIgnoreCase DataStringComparisonBehavior = original.OrdinalIgnoreCase
)

type IngressState = original.IngressState

const (
	Disabled IngressState = original.Disabled
	Paused   IngressState = original.Paused
	Ready    IngressState = original.Ready
	Running  IngressState = original.Running
	Unknown  IngressState = original.Unknown
)

type Kind = original.Kind

const (
	KindEventSourceCreateOrUpdateParameters Kind = original.KindEventSourceCreateOrUpdateParameters
	KindMicrosoftEventHub                   Kind = original.KindMicrosoftEventHub
	KindMicrosoftIoTHub                     Kind = original.KindMicrosoftIoTHub
)

type KindBasicEventSourceResource = original.KindBasicEventSourceResource

const (
	KindBasicEventSourceResourceKindEventSourceResource KindBasicEventSourceResource = original.KindBasicEventSourceResourceKindEventSourceResource
	KindBasicEventSourceResourceKindMicrosoftEventHub   KindBasicEventSourceResource = original.KindBasicEventSourceResourceKindMicrosoftEventHub
	KindBasicEventSourceResourceKindMicrosoftIotHub     KindBasicEventSourceResource = original.KindBasicEventSourceResourceKindMicrosoftIotHub
)

type LocalTimestampFormat = original.LocalTimestampFormat

const (
	Embedded LocalTimestampFormat = original.Embedded
	Iana     LocalTimestampFormat = original.Iana
	TimeSpan LocalTimestampFormat = original.TimeSpan
)

type PropertyType = original.PropertyType

const (
	String PropertyType = original.String
)

type ProvisioningState = original.ProvisioningState

const (
	Accepted  ProvisioningState = original.Accepted
	Creating  ProvisioningState = original.Creating
	Deleting  ProvisioningState = original.Deleting
	Failed    ProvisioningState = original.Failed
	Succeeded ProvisioningState = original.Succeeded
	Updating  ProvisioningState = original.Updating
)

type ReferenceDataKeyPropertyType = original.ReferenceDataKeyPropertyType

const (
	ReferenceDataKeyPropertyTypeBool     ReferenceDataKeyPropertyType = original.ReferenceDataKeyPropertyTypeBool
	ReferenceDataKeyPropertyTypeDateTime ReferenceDataKeyPropertyType = original.ReferenceDataKeyPropertyTypeDateTime
	ReferenceDataKeyPropertyTypeDouble   ReferenceDataKeyPropertyType = original.ReferenceDataKeyPropertyTypeDouble
	ReferenceDataKeyPropertyTypeString   ReferenceDataKeyPropertyType = original.ReferenceDataKeyPropertyTypeString
)

type SkuName = original.SkuName

const (
	S1 SkuName = original.S1
	S2 SkuName = original.S2
)

type StorageLimitExceededBehavior = original.StorageLimitExceededBehavior

const (
	PauseIngress StorageLimitExceededBehavior = original.PauseIngress
	PurgeOldData StorageLimitExceededBehavior = original.PurgeOldData
)

type AccessPoliciesClient = original.AccessPoliciesClient
type AccessPolicyCreateOrUpdateParameters = original.AccessPolicyCreateOrUpdateParameters
type AccessPolicyListResponse = original.AccessPolicyListResponse
type AccessPolicyMutableProperties = original.AccessPolicyMutableProperties
type AccessPolicyResource = original.AccessPolicyResource
type AccessPolicyResourceProperties = original.AccessPolicyResourceProperties
type AccessPolicyUpdateParameters = original.AccessPolicyUpdateParameters
type AzureEventSourceProperties = original.AzureEventSourceProperties
type BaseClient = original.BaseClient
type BasicEventSourceCreateOrUpdateParameters = original.BasicEventSourceCreateOrUpdateParameters
type BasicEventSourceResource = original.BasicEventSourceResource
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type CreateOrUpdateTrackedResourceProperties = original.CreateOrUpdateTrackedResourceProperties
type EnvironmentCreateOrUpdateParameters = original.EnvironmentCreateOrUpdateParameters
type EnvironmentCreationProperties = original.EnvironmentCreationProperties
type EnvironmentListResponse = original.EnvironmentListResponse
type EnvironmentMutableProperties = original.EnvironmentMutableProperties
type EnvironmentResource = original.EnvironmentResource
type EnvironmentResourceProperties = original.EnvironmentResourceProperties
type EnvironmentStateDetails = original.EnvironmentStateDetails
type EnvironmentStatus = original.EnvironmentStatus
type EnvironmentUpdateParameters = original.EnvironmentUpdateParameters
type EnvironmentsClient = original.EnvironmentsClient
type EnvironmentsCreateOrUpdateFuture = original.EnvironmentsCreateOrUpdateFuture
type EnvironmentsUpdateFuture = original.EnvironmentsUpdateFuture
type EventHubEventSourceCommonProperties = original.EventHubEventSourceCommonProperties
type EventHubEventSourceCreateOrUpdateParameters = original.EventHubEventSourceCreateOrUpdateParameters
type EventHubEventSourceCreationProperties = original.EventHubEventSourceCreationProperties
type EventHubEventSourceMutableProperties = original.EventHubEventSourceMutableProperties
type EventHubEventSourceResource = original.EventHubEventSourceResource
type EventHubEventSourceResourceProperties = original.EventHubEventSourceResourceProperties
type EventHubEventSourceUpdateParameters = original.EventHubEventSourceUpdateParameters
type EventSourceCommonProperties = original.EventSourceCommonProperties
type EventSourceCreateOrUpdateParameters = original.EventSourceCreateOrUpdateParameters
type EventSourceListResponse = original.EventSourceListResponse
type EventSourceMutableProperties = original.EventSourceMutableProperties
type EventSourceResource = original.EventSourceResource
type EventSourceResourceModel = original.EventSourceResourceModel
type EventSourceUpdateParameters = original.EventSourceUpdateParameters
type EventSourcesClient = original.EventSourcesClient
type IngressEnvironmentStatus = original.IngressEnvironmentStatus
type IoTHubEventSourceCommonProperties = original.IoTHubEventSourceCommonProperties
type IoTHubEventSourceCreateOrUpdateParameters = original.IoTHubEventSourceCreateOrUpdateParameters
type IoTHubEventSourceCreationProperties = original.IoTHubEventSourceCreationProperties
type IoTHubEventSourceMutableProperties = original.IoTHubEventSourceMutableProperties
type IoTHubEventSourceResource = original.IoTHubEventSourceResource
type IoTHubEventSourceResourceProperties = original.IoTHubEventSourceResourceProperties
type IoTHubEventSourceUpdateParameters = original.IoTHubEventSourceUpdateParameters
type LocalTimestamp = original.LocalTimestamp
type LocalTimestampTimeZoneOffset = original.LocalTimestampTimeZoneOffset
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type PartitionKeyProperty = original.PartitionKeyProperty
type ReferenceDataSetCreateOrUpdateParameters = original.ReferenceDataSetCreateOrUpdateParameters
type ReferenceDataSetCreationProperties = original.ReferenceDataSetCreationProperties
type ReferenceDataSetKeyProperty = original.ReferenceDataSetKeyProperty
type ReferenceDataSetListResponse = original.ReferenceDataSetListResponse
type ReferenceDataSetResource = original.ReferenceDataSetResource
type ReferenceDataSetResourceProperties = original.ReferenceDataSetResourceProperties
type ReferenceDataSetUpdateParameters = original.ReferenceDataSetUpdateParameters
type ReferenceDataSetsClient = original.ReferenceDataSetsClient
type Resource = original.Resource
type ResourceProperties = original.ResourceProperties
type Sku = original.Sku
type TrackedResource = original.TrackedResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAccessPoliciesClient(subscriptionID string) AccessPoliciesClient {
	return original.NewAccessPoliciesClient(subscriptionID)
}
func NewAccessPoliciesClientWithBaseURI(baseURI string, subscriptionID string) AccessPoliciesClient {
	return original.NewAccessPoliciesClientWithBaseURI(baseURI, subscriptionID)
}
func NewEnvironmentsClient(subscriptionID string) EnvironmentsClient {
	return original.NewEnvironmentsClient(subscriptionID)
}
func NewEnvironmentsClientWithBaseURI(baseURI string, subscriptionID string) EnvironmentsClient {
	return original.NewEnvironmentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewEventSourcesClient(subscriptionID string) EventSourcesClient {
	return original.NewEventSourcesClient(subscriptionID)
}
func NewEventSourcesClientWithBaseURI(baseURI string, subscriptionID string) EventSourcesClient {
	return original.NewEventSourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewReferenceDataSetsClient(subscriptionID string) ReferenceDataSetsClient {
	return original.NewReferenceDataSetsClient(subscriptionID)
}
func NewReferenceDataSetsClientWithBaseURI(baseURI string, subscriptionID string) ReferenceDataSetsClient {
	return original.NewReferenceDataSetsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAccessPolicyRoleValues() []AccessPolicyRole {
	return original.PossibleAccessPolicyRoleValues()
}
func PossibleDataStringComparisonBehaviorValues() []DataStringComparisonBehavior {
	return original.PossibleDataStringComparisonBehaviorValues()
}
func PossibleIngressStateValues() []IngressState {
	return original.PossibleIngressStateValues()
}
func PossibleKindBasicEventSourceResourceValues() []KindBasicEventSourceResource {
	return original.PossibleKindBasicEventSourceResourceValues()
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleLocalTimestampFormatValues() []LocalTimestampFormat {
	return original.PossibleLocalTimestampFormatValues()
}
func PossiblePropertyTypeValues() []PropertyType {
	return original.PossiblePropertyTypeValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleReferenceDataKeyPropertyTypeValues() []ReferenceDataKeyPropertyType {
	return original.PossibleReferenceDataKeyPropertyTypeValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleStorageLimitExceededBehaviorValues() []StorageLimitExceededBehavior {
	return original.PossibleStorageLimitExceededBehaviorValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}