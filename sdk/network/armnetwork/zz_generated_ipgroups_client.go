//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// IPGroupsClient contains the methods for the IPGroups group.
// Don't use this type directly, use NewIPGroupsClient() instead.
type IPGroupsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewIPGroupsClient creates a new instance of IPGroupsClient with the specified values.
func NewIPGroupsClient(con *arm.Connection, subscriptionID string) *IPGroupsClient {
	return &IPGroupsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates an ipGroups in a specified resource group.
// If the operation fails it returns the *Error error type.
func (client *IPGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, ipGroupsName string, parameters IPGroup, options *IPGroupsBeginCreateOrUpdateOptions) (IPGroupsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, ipGroupsName, parameters, options)
	if err != nil {
		return IPGroupsCreateOrUpdatePollerResponse{}, err
	}
	result := IPGroupsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("IPGroupsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return IPGroupsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &IPGroupsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates an ipGroups in a specified resource group.
// If the operation fails it returns the *Error error type.
func (client *IPGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, ipGroupsName string, parameters IPGroup, options *IPGroupsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, ipGroupsName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IPGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, ipGroupsName string, parameters IPGroup, options *IPGroupsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ipGroups/{ipGroupsName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ipGroupsName == "" {
		return nil, errors.New("parameter ipGroupsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipGroupsName}", url.PathEscape(ipGroupsName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *IPGroupsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes the specified ipGroups.
// If the operation fails it returns the *Error error type.
func (client *IPGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, ipGroupsName string, options *IPGroupsBeginDeleteOptions) (IPGroupsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, ipGroupsName, options)
	if err != nil {
		return IPGroupsDeletePollerResponse{}, err
	}
	result := IPGroupsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("IPGroupsClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return IPGroupsDeletePollerResponse{}, err
	}
	result.Poller = &IPGroupsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified ipGroups.
// If the operation fails it returns the *Error error type.
func (client *IPGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, ipGroupsName string, options *IPGroupsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, ipGroupsName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IPGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, ipGroupsName string, options *IPGroupsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ipGroups/{ipGroupsName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ipGroupsName == "" {
		return nil, errors.New("parameter ipGroupsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipGroupsName}", url.PathEscape(ipGroupsName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *IPGroupsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets the specified ipGroups.
// If the operation fails it returns the *Error error type.
func (client *IPGroupsClient) Get(ctx context.Context, resourceGroupName string, ipGroupsName string, options *IPGroupsGetOptions) (IPGroupsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, ipGroupsName, options)
	if err != nil {
		return IPGroupsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IPGroupsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IPGroupsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IPGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, ipGroupsName string, options *IPGroupsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ipGroups/{ipGroupsName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ipGroupsName == "" {
		return nil, errors.New("parameter ipGroupsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipGroupsName}", url.PathEscape(ipGroupsName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IPGroupsClient) getHandleResponse(resp *http.Response) (IPGroupsGetResponse, error) {
	result := IPGroupsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPGroup); err != nil {
		return IPGroupsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *IPGroupsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Gets all IpGroups in a subscription.
// If the operation fails it returns the *Error error type.
func (client *IPGroupsClient) List(options *IPGroupsListOptions) *IPGroupsListPager {
	return &IPGroupsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp IPGroupsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.IPGroupListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *IPGroupsClient) listCreateRequest(ctx context.Context, options *IPGroupsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ipGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IPGroupsClient) listHandleResponse(resp *http.Response) (IPGroupsListResponse, error) {
	result := IPGroupsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPGroupListResult); err != nil {
		return IPGroupsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *IPGroupsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - Gets all IpGroups in a resource group.
// If the operation fails it returns the *Error error type.
func (client *IPGroupsClient) ListByResourceGroup(resourceGroupName string, options *IPGroupsListByResourceGroupOptions) *IPGroupsListByResourceGroupPager {
	return &IPGroupsListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp IPGroupsListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.IPGroupListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *IPGroupsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *IPGroupsListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ipGroups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *IPGroupsClient) listByResourceGroupHandleResponse(resp *http.Response) (IPGroupsListByResourceGroupResponse, error) {
	result := IPGroupsListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPGroupListResult); err != nil {
		return IPGroupsListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *IPGroupsClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// UpdateGroups - Updates tags of an IpGroups resource.
// If the operation fails it returns the *Error error type.
func (client *IPGroupsClient) UpdateGroups(ctx context.Context, resourceGroupName string, ipGroupsName string, parameters TagsObject, options *IPGroupsUpdateGroupsOptions) (IPGroupsUpdateGroupsResponse, error) {
	req, err := client.updateGroupsCreateRequest(ctx, resourceGroupName, ipGroupsName, parameters, options)
	if err != nil {
		return IPGroupsUpdateGroupsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IPGroupsUpdateGroupsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IPGroupsUpdateGroupsResponse{}, client.updateGroupsHandleError(resp)
	}
	return client.updateGroupsHandleResponse(resp)
}

// updateGroupsCreateRequest creates the UpdateGroups request.
func (client *IPGroupsClient) updateGroupsCreateRequest(ctx context.Context, resourceGroupName string, ipGroupsName string, parameters TagsObject, options *IPGroupsUpdateGroupsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ipGroups/{ipGroupsName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ipGroupsName == "" {
		return nil, errors.New("parameter ipGroupsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipGroupsName}", url.PathEscape(ipGroupsName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateGroupsHandleResponse handles the UpdateGroups response.
func (client *IPGroupsClient) updateGroupsHandleResponse(resp *http.Response) (IPGroupsUpdateGroupsResponse, error) {
	result := IPGroupsUpdateGroupsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPGroup); err != nil {
		return IPGroupsUpdateGroupsResponse{}, err
	}
	return result, nil
}

// updateGroupsHandleError handles the UpdateGroups error response.
func (client *IPGroupsClient) updateGroupsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
