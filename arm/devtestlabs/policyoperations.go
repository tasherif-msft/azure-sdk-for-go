package devtestlabs

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

// PolicyOperationsClient is the the DevTest Labs Client.
type PolicyOperationsClient struct {
	ManagementClient
}

// NewPolicyOperationsClient creates an instance of the PolicyOperationsClient
// client.
func NewPolicyOperationsClient(subscriptionID string) PolicyOperationsClient {
	return PolicyOperationsClient{New(subscriptionID)}
}

// CreateOrUpdateResource create or replace an existing policy.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. policySetName is the name of the policy set. name is the name of
// the policy.
func (client PolicyOperationsClient) CreateOrUpdateResource(resourceGroupName string, labName string, policySetName string, name string, policy Policy) (result Policy, err error) {
	req, err := client.CreateOrUpdateResourcePreparer(resourceGroupName, labName, policySetName, name, policy)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.PolicyOperationsClient", "CreateOrUpdateResource", nil, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "devtestlabs.PolicyOperationsClient", "CreateOrUpdateResource", resp, "Failure sending request")
	}

	result, err = client.CreateOrUpdateResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.PolicyOperationsClient", "CreateOrUpdateResource", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdateResourcePreparer prepares the CreateOrUpdateResource request.
func (client PolicyOperationsClient) CreateOrUpdateResourcePreparer(resourceGroupName string, labName string, policySetName string, name string, policy Policy) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"policySetName":     autorest.Encode("path", policySetName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/policysets/{policySetName}/policies/{name}", pathParameters),
		autorest.WithJSON(policy),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateResourceSender sends the CreateOrUpdateResource request. The method will close the
// http.Response Body if it receives an error.
func (client PolicyOperationsClient) CreateOrUpdateResourceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResourceResponder handles the response to the CreateOrUpdateResource request. The method always
// closes the http.Response Body.
func (client PolicyOperationsClient) CreateOrUpdateResourceResponder(resp *http.Response) (result Policy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteResource delete policy.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. policySetName is the name of the policy set. name is the name of
// the policy.
func (client PolicyOperationsClient) DeleteResource(resourceGroupName string, labName string, policySetName string, name string) (result autorest.Response, err error) {
	req, err := client.DeleteResourcePreparer(resourceGroupName, labName, policySetName, name)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.PolicyOperationsClient", "DeleteResource", nil, "Failure preparing request")
	}

	resp, err := client.DeleteResourceSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "devtestlabs.PolicyOperationsClient", "DeleteResource", resp, "Failure sending request")
	}

	result, err = client.DeleteResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.PolicyOperationsClient", "DeleteResource", resp, "Failure responding to request")
	}

	return
}

// DeleteResourcePreparer prepares the DeleteResource request.
func (client PolicyOperationsClient) DeleteResourcePreparer(resourceGroupName string, labName string, policySetName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"policySetName":     autorest.Encode("path", policySetName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/policysets/{policySetName}/policies/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteResourceSender sends the DeleteResource request. The method will close the
// http.Response Body if it receives an error.
func (client PolicyOperationsClient) DeleteResourceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResourceResponder handles the response to the DeleteResource request. The method always
// closes the http.Response Body.
func (client PolicyOperationsClient) DeleteResourceResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetResource get policy.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. policySetName is the name of the policy set. name is the name of
// the policy.
func (client PolicyOperationsClient) GetResource(resourceGroupName string, labName string, policySetName string, name string) (result Policy, err error) {
	req, err := client.GetResourcePreparer(resourceGroupName, labName, policySetName, name)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.PolicyOperationsClient", "GetResource", nil, "Failure preparing request")
	}

	resp, err := client.GetResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "devtestlabs.PolicyOperationsClient", "GetResource", resp, "Failure sending request")
	}

	result, err = client.GetResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.PolicyOperationsClient", "GetResource", resp, "Failure responding to request")
	}

	return
}

// GetResourcePreparer prepares the GetResource request.
func (client PolicyOperationsClient) GetResourcePreparer(resourceGroupName string, labName string, policySetName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"policySetName":     autorest.Encode("path", policySetName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/policysets/{policySetName}/policies/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetResourceSender sends the GetResource request. The method will close the
// http.Response Body if it receives an error.
func (client PolicyOperationsClient) GetResourceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResourceResponder handles the response to the GetResource request. The method always
// closes the http.Response Body.
func (client PolicyOperationsClient) GetResourceResponder(resp *http.Response) (result Policy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list policies in a given policy set.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. policySetName is the name of the policy set. filter is the filter
// to apply on the operation. top is the maximum number of resources to
// return from the operation. orderBy is the ordering expression for the
// results, using OData notation.
func (client PolicyOperationsClient) List(resourceGroupName string, labName string, policySetName string, filter string, top *int32, orderBy string) (result ResponseWithContinuationPolicy, err error) {
	req, err := client.ListPreparer(resourceGroupName, labName, policySetName, filter, top, orderBy)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.PolicyOperationsClient", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "devtestlabs.PolicyOperationsClient", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.PolicyOperationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client PolicyOperationsClient) ListPreparer(resourceGroupName string, labName string, policySetName string, filter string, top *int32, orderBy string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"policySetName":     autorest.Encode("path", policySetName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderBy) > 0 {
		queryParameters["$orderBy"] = autorest.Encode("query", orderBy)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/policysets/{policySetName}/policies", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client PolicyOperationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client PolicyOperationsClient) ListResponder(resp *http.Response) (result ResponseWithContinuationPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client PolicyOperationsClient) ListNextResults(lastResults ResponseWithContinuationPolicy) (result ResponseWithContinuationPolicy, err error) {
	req, err := lastResults.ResponseWithContinuationPolicyPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.PolicyOperationsClient", "List", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "devtestlabs.PolicyOperationsClient", "List", resp, "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.PolicyOperationsClient", "List", resp, "Failure responding to next results request request")
	}

	return
}

// PatchResource modify properties of policies.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. policySetName is the name of the policy set. name is the name of
// the policy.
func (client PolicyOperationsClient) PatchResource(resourceGroupName string, labName string, policySetName string, name string, policy Policy) (result Policy, err error) {
	req, err := client.PatchResourcePreparer(resourceGroupName, labName, policySetName, name, policy)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.PolicyOperationsClient", "PatchResource", nil, "Failure preparing request")
	}

	resp, err := client.PatchResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "devtestlabs.PolicyOperationsClient", "PatchResource", resp, "Failure sending request")
	}

	result, err = client.PatchResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.PolicyOperationsClient", "PatchResource", resp, "Failure responding to request")
	}

	return
}

// PatchResourcePreparer prepares the PatchResource request.
func (client PolicyOperationsClient) PatchResourcePreparer(resourceGroupName string, labName string, policySetName string, name string, policy Policy) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"policySetName":     autorest.Encode("path", policySetName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/policysets/{policySetName}/policies/{name}", pathParameters),
		autorest.WithJSON(policy),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// PatchResourceSender sends the PatchResource request. The method will close the
// http.Response Body if it receives an error.
func (client PolicyOperationsClient) PatchResourceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// PatchResourceResponder handles the response to the PatchResource request. The method always
// closes the http.Response Body.
func (client PolicyOperationsClient) PatchResourceResponder(resp *http.Response) (result Policy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
