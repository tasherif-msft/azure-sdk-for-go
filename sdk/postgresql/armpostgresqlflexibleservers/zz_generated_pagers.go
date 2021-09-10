//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresqlflexibleservers

import (
	"context"
	"net/http"
	"reflect"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// ConfigurationsListByServerPager provides operations for iterating over paged responses.
type ConfigurationsListByServerPager struct {
	client    *ConfigurationsClient
	current   ConfigurationsListByServerResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ConfigurationsListByServerResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ConfigurationsListByServerPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ConfigurationsListByServerPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ConfigurationListResult.NextLink == nil || len(*p.current.ConfigurationListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByServerHandleError(resp)
		return false
	}
	result, err := p.client.listByServerHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ConfigurationsListByServerResponse page.
func (p *ConfigurationsListByServerPager) PageResponse() ConfigurationsListByServerResponse {
	return p.current
}

// DatabasesListByServerPager provides operations for iterating over paged responses.
type DatabasesListByServerPager struct {
	client    *DatabasesClient
	current   DatabasesListByServerResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DatabasesListByServerResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DatabasesListByServerPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DatabasesListByServerPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DatabaseListResult.NextLink == nil || len(*p.current.DatabaseListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByServerHandleError(resp)
		return false
	}
	result, err := p.client.listByServerHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current DatabasesListByServerResponse page.
func (p *DatabasesListByServerPager) PageResponse() DatabasesListByServerResponse {
	return p.current
}

// FirewallRulesListByServerPager provides operations for iterating over paged responses.
type FirewallRulesListByServerPager struct {
	client    *FirewallRulesClient
	current   FirewallRulesListByServerResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, FirewallRulesListByServerResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *FirewallRulesListByServerPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *FirewallRulesListByServerPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.FirewallRuleListResult.NextLink == nil || len(*p.current.FirewallRuleListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByServerHandleError(resp)
		return false
	}
	result, err := p.client.listByServerHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current FirewallRulesListByServerResponse page.
func (p *FirewallRulesListByServerPager) PageResponse() FirewallRulesListByServerResponse {
	return p.current
}

// LocationBasedCapabilitiesExecutePager provides operations for iterating over paged responses.
type LocationBasedCapabilitiesExecutePager struct {
	client    *LocationBasedCapabilitiesClient
	current   LocationBasedCapabilitiesExecuteResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, LocationBasedCapabilitiesExecuteResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *LocationBasedCapabilitiesExecutePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *LocationBasedCapabilitiesExecutePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.CapabilitiesListResult.NextLink == nil || len(*p.current.CapabilitiesListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.executeHandleError(resp)
		return false
	}
	result, err := p.client.executeHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current LocationBasedCapabilitiesExecuteResponse page.
func (p *LocationBasedCapabilitiesExecutePager) PageResponse() LocationBasedCapabilitiesExecuteResponse {
	return p.current
}

// ServersListByResourceGroupPager provides operations for iterating over paged responses.
type ServersListByResourceGroupPager struct {
	client    *ServersClient
	current   ServersListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ServersListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ServersListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ServersListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ServerListResult.NextLink == nil || len(*p.current.ServerListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByResourceGroupHandleError(resp)
		return false
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ServersListByResourceGroupResponse page.
func (p *ServersListByResourceGroupPager) PageResponse() ServersListByResourceGroupResponse {
	return p.current
}

// ServersListPager provides operations for iterating over paged responses.
type ServersListPager struct {
	client    *ServersClient
	current   ServersListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ServersListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ServersListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ServersListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ServerListResult.NextLink == nil || len(*p.current.ServerListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ServersListResponse page.
func (p *ServersListPager) PageResponse() ServersListResponse {
	return p.current
}
