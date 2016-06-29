package sql

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
// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// SecurityAlertPolicyClient is the the Azure SQL Database management API
// provides a RESTful set of web services that interact with Azure SQL
// Database services to manage your databases. The API enables users to
// create, retrieve, update, and delete Databases, Servers and related
// resources.
type SecurityAlertPolicyClient struct {
	ManagementClient
}

// NewSecurityAlertPolicyClient creates an instance of the
// SecurityAlertPolicyClient client.
func NewSecurityAlertPolicyClient(subscriptionID string) SecurityAlertPolicyClient {
	return SecurityAlertPolicyClient{New(subscriptionID)}
}

// CreateOrUpdate creates or updates an Azure SQL Server security alert
// policy. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel
// polling and any outstanding HTTP requests.
//
// parameters is the required parameters for creating or updating a Azure SQL
// Server security alert policy. resourceGroupName is the name of the
// Resource Group to which the server belongs. serverName is the name of the
// Azure SQL Server.
func (client SecurityAlertPolicyClient) CreateOrUpdate(parameters ServerSecurityAlertPolicyCreateOrUpdateParameters, resourceGroupName string, serverName string, cancel <-chan struct{}) (result autorest.Response, err error) {
	req, err := client.CreateOrUpdatePreparer(parameters, resourceGroupName, serverName, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.SecurityAlertPolicyClient", "CreateOrUpdate", nil, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "sql.SecurityAlertPolicyClient", "CreateOrUpdate", resp, "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SecurityAlertPolicyClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SecurityAlertPolicyClient) CreateOrUpdatePreparer(parameters ServerSecurityAlertPolicyCreateOrUpdateParameters, resourceGroupName string, serverName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/securityAlertPolicies/Default", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityAlertPolicyClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SecurityAlertPolicyClient) CreateOrUpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get returns an Azure SQL Server security alert policy.
//
// resourceGroupName is the name of the Resource Group to which the server
// belongs. serverName is the name of the Azure SQL Server.
func (client SecurityAlertPolicyClient) Get(resourceGroupName string, serverName string) (result ServerSecurityAlertPolicy, err error) {
	req, err := client.GetPreparer(resourceGroupName, serverName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.SecurityAlertPolicyClient", "Get", nil, "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.SecurityAlertPolicyClient", "Get", resp, "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SecurityAlertPolicyClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SecurityAlertPolicyClient) GetPreparer(resourceGroupName string, serverName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/securityAlertPolicies/Default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityAlertPolicyClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SecurityAlertPolicyClient) GetResponder(resp *http.Response) (result ServerSecurityAlertPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
