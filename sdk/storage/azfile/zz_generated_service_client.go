//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azfile

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strconv"
	"strings"
)

type serviceClient struct {
	endpoint string
	pl runtime.Pipeline
}

// newServiceClient creates a new instance of serviceClient with the specified values.
// endpoint - The URL of the service account, share, directory or file that is the target of the desired operation.
// pl - the pipeline used for sending requests and handling responses.
func newServiceClient(endpoint string, pl runtime.Pipeline) *serviceClient {
	client := &serviceClient{
		endpoint: endpoint,
		pl: pl,
	}
	return client
}

// GetProperties - Gets the properties of a storage account's File service, including properties for Storage Analytics metrics
// and CORS (Cross-Origin Resource Sharing) rules.
// If the operation fails it returns an *azcore.ResponseError type.
// options - serviceClientGetPropertiesOptions contains the optional parameters for the serviceClient.GetProperties method.
func (client *serviceClient) GetProperties(ctx context.Context, options *serviceClientGetPropertiesOptions) (serviceClientGetPropertiesResponse, error) {
	req, err := client.getPropertiesCreateRequest(ctx, options)
	if err != nil {
		return serviceClientGetPropertiesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return serviceClientGetPropertiesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return serviceClientGetPropertiesResponse{}, runtime.NewResponseError(resp)
	}
	return client.getPropertiesHandleResponse(resp)
}

// getPropertiesCreateRequest creates the GetProperties request.
func (client *serviceClient) getPropertiesCreateRequest(ctx context.Context, options *serviceClientGetPropertiesOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodGet, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("restype", "service")
	reqQP.Set("comp", "properties")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("x-ms-version", "2020-10-02")
	req.Raw().Header.Set("Accept", "application/xml")
	return req, nil
}

// getPropertiesHandleResponse handles the GetProperties response.
func (client *serviceClient) getPropertiesHandleResponse(resp *http.Response) (serviceClientGetPropertiesResponse, error) {
	result := serviceClientGetPropertiesResponse{RawResponse: resp}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if err := runtime.UnmarshalAsXML(resp, &result.StorageServiceProperties); err != nil {
		return serviceClientGetPropertiesResponse{}, err
	}
	return result, nil
}

// ListSharesSegment - The List Shares Segment operation returns a list of the shares and share snapshots under the specified
// account.
// If the operation fails it returns an *azcore.ResponseError type.
// options - serviceClientListSharesSegmentOptions contains the optional parameters for the serviceClient.ListSharesSegment
// method.
func (client *serviceClient) ListSharesSegment(options *serviceClientListSharesSegmentOptions) (*serviceClientListSharesSegmentPager) {
	return &serviceClientListSharesSegmentPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listSharesSegmentCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp serviceClientListSharesSegmentResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListSharesResponse.NextMarker)
		},
	}
}

// listSharesSegmentCreateRequest creates the ListSharesSegment request.
func (client *serviceClient) listSharesSegmentCreateRequest(ctx context.Context, options *serviceClientListSharesSegmentOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodGet, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("comp", "list")
	if options != nil && options.Prefix != nil {
		reqQP.Set("prefix", *options.Prefix)
	}
	if options != nil && options.Marker != nil {
		reqQP.Set("marker", *options.Marker)
	}
	if options != nil && options.Maxresults != nil {
		reqQP.Set("maxresults", strconv.FormatInt(int64(*options.Maxresults), 10))
	}
	if options != nil && options.Include != nil {
		reqQP.Set("include", strings.Join(strings.Fields(strings.Trim(fmt.Sprint(options.Include), "[]")), ","))
	}
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("x-ms-version", "2020-10-02")
	req.Raw().Header.Set("Accept", "application/xml")
	return req, nil
}

// listSharesSegmentHandleResponse handles the ListSharesSegment response.
func (client *serviceClient) listSharesSegmentHandleResponse(resp *http.Response) (serviceClientListSharesSegmentResponse, error) {
	result := serviceClientListSharesSegmentResponse{RawResponse: resp}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if err := runtime.UnmarshalAsXML(resp, &result.ListSharesResponse); err != nil {
		return serviceClientListSharesSegmentResponse{}, err
	}
	return result, nil
}

// SetProperties - Sets properties for a storage account's File service endpoint, including properties for Storage Analytics
// metrics and CORS (Cross-Origin Resource Sharing) rules.
// If the operation fails it returns an *azcore.ResponseError type.
// storageServiceProperties - The StorageService properties.
// options - serviceClientSetPropertiesOptions contains the optional parameters for the serviceClient.SetProperties method.
func (client *serviceClient) SetProperties(ctx context.Context, storageServiceProperties StorageServiceProperties, options *serviceClientSetPropertiesOptions) (serviceClientSetPropertiesResponse, error) {
	req, err := client.setPropertiesCreateRequest(ctx, storageServiceProperties, options)
	if err != nil {
		return serviceClientSetPropertiesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return serviceClientSetPropertiesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return serviceClientSetPropertiesResponse{}, runtime.NewResponseError(resp)
	}
	return client.setPropertiesHandleResponse(resp)
}

// setPropertiesCreateRequest creates the SetProperties request.
func (client *serviceClient) setPropertiesCreateRequest(ctx context.Context, storageServiceProperties StorageServiceProperties, options *serviceClientSetPropertiesOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodPut, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("restype", "service")
	reqQP.Set("comp", "properties")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("x-ms-version", "2020-10-02")
	req.Raw().Header.Set("Accept", "application/xml")
	return req, runtime.MarshalAsXML(req, storageServiceProperties)
}

// setPropertiesHandleResponse handles the SetProperties response.
func (client *serviceClient) setPropertiesHandleResponse(resp *http.Response) (serviceClientSetPropertiesResponse, error) {
	result := serviceClientSetPropertiesResponse{RawResponse: resp}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	return result, nil
}

