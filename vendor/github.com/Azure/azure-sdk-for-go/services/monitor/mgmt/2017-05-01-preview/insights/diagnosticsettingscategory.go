package insights

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
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2017-05-01-preview/insights instead.
// DiagnosticSettingsCategoryClient is the monitor Management Client
type DiagnosticSettingsCategoryClient struct {
	BaseClient
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2017-05-01-preview/insights instead.
// NewDiagnosticSettingsCategoryClient creates an instance of the DiagnosticSettingsCategoryClient client.
func NewDiagnosticSettingsCategoryClient(subscriptionID string) DiagnosticSettingsCategoryClient {
	return NewDiagnosticSettingsCategoryClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2017-05-01-preview/insights instead.
// NewDiagnosticSettingsCategoryClientWithBaseURI creates an instance of the DiagnosticSettingsCategoryClient client.
func NewDiagnosticSettingsCategoryClientWithBaseURI(baseURI string, subscriptionID string) DiagnosticSettingsCategoryClient {
	return DiagnosticSettingsCategoryClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2017-05-01-preview/insights instead.
// Get gets the diagnostic settings category for the specified resource.
//
// resourceURI is the identifier of the resource. name is the name of the diagnostic setting.
func (client DiagnosticSettingsCategoryClient) Get(ctx context.Context, resourceURI string, name string) (result DiagnosticSettingsCategoryResource, err error) {
	req, err := client.GetPreparer(ctx, resourceURI, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.DiagnosticSettingsCategoryClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.DiagnosticSettingsCategoryClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.DiagnosticSettingsCategoryClient", "Get", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2017-05-01-preview/insights instead.
// GetPreparer prepares the Get request.
func (client DiagnosticSettingsCategoryClient) GetPreparer(ctx context.Context, resourceURI string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":        autorest.Encode("path", name),
		"resourceUri": autorest.Encode("path", resourceURI),
	}

	const APIVersion = "2017-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceUri}/providers/microsoft.insights/diagnosticSettingsCategories/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2017-05-01-preview/insights instead.
// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DiagnosticSettingsCategoryClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2017-05-01-preview/insights instead.
// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DiagnosticSettingsCategoryClient) GetResponder(resp *http.Response) (result DiagnosticSettingsCategoryResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2017-05-01-preview/insights instead.
// List lists the diagnostic settings categories for the specified resource.
//
// resourceURI is the identifier of the resource.
func (client DiagnosticSettingsCategoryClient) List(ctx context.Context, resourceURI string) (result DiagnosticSettingsCategoryResourceCollection, err error) {
	req, err := client.ListPreparer(ctx, resourceURI)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.DiagnosticSettingsCategoryClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.DiagnosticSettingsCategoryClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.DiagnosticSettingsCategoryClient", "List", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2017-05-01-preview/insights instead.
// ListPreparer prepares the List request.
func (client DiagnosticSettingsCategoryClient) ListPreparer(ctx context.Context, resourceURI string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceUri": autorest.Encode("path", resourceURI),
	}

	const APIVersion = "2017-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceUri}/providers/microsoft.insights/diagnosticSettingsCategories", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2017-05-01-preview/insights instead.
// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DiagnosticSettingsCategoryClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2017-05-01-preview/insights instead.
// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DiagnosticSettingsCategoryClient) ListResponder(resp *http.Response) (result DiagnosticSettingsCategoryResourceCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}