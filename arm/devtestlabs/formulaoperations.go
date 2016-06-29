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

// FormulaOperationsClient is the the DevTest Labs Client.
type FormulaOperationsClient struct {
	ManagementClient
}

// NewFormulaOperationsClient creates an instance of the
// FormulaOperationsClient client.
func NewFormulaOperationsClient(subscriptionID string) FormulaOperationsClient {
	return FormulaOperationsClient{New(subscriptionID)}
}

// CreateOrUpdateResource create or replace an existing Formula. This
// operation can take a while to complete. This method may poll for
// completion. Polling can be canceled by passing the cancel channel
// argument. The channel will be used to cancel polling and any outstanding
// HTTP requests.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. name is the name of the formula.
func (client FormulaOperationsClient) CreateOrUpdateResource(resourceGroupName string, labName string, name string, formula Formula, cancel <-chan struct{}) (result autorest.Response, err error) {
	req, err := client.CreateOrUpdateResourcePreparer(resourceGroupName, labName, name, formula, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.FormulaOperationsClient", "CreateOrUpdateResource", nil, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateResourceSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "devtestlabs.FormulaOperationsClient", "CreateOrUpdateResource", resp, "Failure sending request")
	}

	result, err = client.CreateOrUpdateResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.FormulaOperationsClient", "CreateOrUpdateResource", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdateResourcePreparer prepares the CreateOrUpdateResource request.
func (client FormulaOperationsClient) CreateOrUpdateResourcePreparer(resourceGroupName string, labName string, name string, formula Formula, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/formulas/{name}", pathParameters),
		autorest.WithJSON(formula),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateOrUpdateResourceSender sends the CreateOrUpdateResource request. The method will close the
// http.Response Body if it receives an error.
func (client FormulaOperationsClient) CreateOrUpdateResourceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateResourceResponder handles the response to the CreateOrUpdateResource request. The method always
// closes the http.Response Body.
func (client FormulaOperationsClient) CreateOrUpdateResourceResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteResource delete formula.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. name is the name of the formula.
func (client FormulaOperationsClient) DeleteResource(resourceGroupName string, labName string, name string) (result autorest.Response, err error) {
	req, err := client.DeleteResourcePreparer(resourceGroupName, labName, name)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.FormulaOperationsClient", "DeleteResource", nil, "Failure preparing request")
	}

	resp, err := client.DeleteResourceSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "devtestlabs.FormulaOperationsClient", "DeleteResource", resp, "Failure sending request")
	}

	result, err = client.DeleteResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.FormulaOperationsClient", "DeleteResource", resp, "Failure responding to request")
	}

	return
}

// DeleteResourcePreparer prepares the DeleteResource request.
func (client FormulaOperationsClient) DeleteResourcePreparer(resourceGroupName string, labName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/formulas/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteResourceSender sends the DeleteResource request. The method will close the
// http.Response Body if it receives an error.
func (client FormulaOperationsClient) DeleteResourceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResourceResponder handles the response to the DeleteResource request. The method always
// closes the http.Response Body.
func (client FormulaOperationsClient) DeleteResourceResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetResource get formula.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. name is the name of the formula.
func (client FormulaOperationsClient) GetResource(resourceGroupName string, labName string, name string) (result Formula, err error) {
	req, err := client.GetResourcePreparer(resourceGroupName, labName, name)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.FormulaOperationsClient", "GetResource", nil, "Failure preparing request")
	}

	resp, err := client.GetResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "devtestlabs.FormulaOperationsClient", "GetResource", resp, "Failure sending request")
	}

	result, err = client.GetResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.FormulaOperationsClient", "GetResource", resp, "Failure responding to request")
	}

	return
}

// GetResourcePreparer prepares the GetResource request.
func (client FormulaOperationsClient) GetResourcePreparer(resourceGroupName string, labName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/formulas/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetResourceSender sends the GetResource request. The method will close the
// http.Response Body if it receives an error.
func (client FormulaOperationsClient) GetResourceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResourceResponder handles the response to the GetResource request. The method always
// closes the http.Response Body.
func (client FormulaOperationsClient) GetResourceResponder(resp *http.Response) (result Formula, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list formulas in a given lab.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. filter is the filter to apply on the operation. top is the
// maximum number of resources to return from the operation. orderBy is the
// ordering expression for the results, using OData notation.
func (client FormulaOperationsClient) List(resourceGroupName string, labName string, filter string, top *int32, orderBy string) (result ResponseWithContinuationFormula, err error) {
	req, err := client.ListPreparer(resourceGroupName, labName, filter, top, orderBy)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.FormulaOperationsClient", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "devtestlabs.FormulaOperationsClient", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.FormulaOperationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client FormulaOperationsClient) ListPreparer(resourceGroupName string, labName string, filter string, top *int32, orderBy string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/formulas", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client FormulaOperationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client FormulaOperationsClient) ListResponder(resp *http.Response) (result ResponseWithContinuationFormula, err error) {
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
func (client FormulaOperationsClient) ListNextResults(lastResults ResponseWithContinuationFormula) (result ResponseWithContinuationFormula, err error) {
	req, err := lastResults.ResponseWithContinuationFormulaPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.FormulaOperationsClient", "List", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "devtestlabs.FormulaOperationsClient", "List", resp, "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.FormulaOperationsClient", "List", resp, "Failure responding to next results request request")
	}

	return
}
