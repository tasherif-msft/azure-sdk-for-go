//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package path

import (
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azdatalake/datalakeerror"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azdatalake/internal/exported"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azdatalake/internal/generated"
	"time"
)

// GetPropertiesOptions contains the optional parameters for the Client.GetProperties method
type GetPropertiesOptions struct {
	AccessConditions *AccessConditions
	CPKInfo          *CPKInfo
}

func FormatGetPropertiesOptions(o *GetPropertiesOptions) *blob.GetPropertiesOptions {
	if o == nil {
		return nil
	}
	accessConditions := exported.FormatBlobAccessConditions(o.AccessConditions)
	return &blob.GetPropertiesOptions{
		AccessConditions: accessConditions,
		CPKInfo: &blob.CPKInfo{
			EncryptionKey:       o.CPKInfo.EncryptionKey,
			EncryptionAlgorithm: o.CPKInfo.EncryptionAlgorithm,
			EncryptionKeySHA256: o.CPKInfo.EncryptionKeySHA256,
		},
	}
}

// ===================================== PATH IMPORTS ===========================================

// SetAccessControlOptions contains the optional parameters when calling the SetAccessControl operation. dfs endpoint
type SetAccessControlOptions struct {
	// Owner is the owner of the path.
	Owner *string
	// Group is the owning group of the path.
	Group *string
	// ACL is the access control list for the path.
	ACL *string
	// Permissions is the octal representation of the permissions for user, group and mask.
	Permissions *string
	// AccessConditions contains parameters for accessing the path.
	AccessConditions *AccessConditions
}

func FormatSetAccessControlOptions(o *SetAccessControlOptions) (*generated.PathClientSetAccessControlOptions, *generated.LeaseAccessConditions, *generated.ModifiedAccessConditions, error) {
	if o == nil {
		return nil, nil, nil, datalakeerror.MissingParameters
	}
	// call path formatter since we're hitting dfs in this operation
	leaseAccessConditions, modifiedAccessConditions := exported.FormatPathAccessConditions(o.AccessConditions)
	if o.Owner == nil && o.Group == nil && o.ACL == nil && o.Permissions == nil {
		return nil, nil, nil, errors.New("at least one parameter should be set for SetAccessControl API")
	}
	return &generated.PathClientSetAccessControlOptions{
		Owner:       o.Owner,
		Group:       o.Group,
		ACL:         o.ACL,
		Permissions: o.Permissions,
	}, leaseAccessConditions, modifiedAccessConditions, nil
}

// GetAccessControlOptions contains the optional parameters when calling the GetAccessControl operation.
type GetAccessControlOptions struct {
	// UPN is the user principal name.
	UPN *bool
	// AccessConditions contains parameters for accessing the path.
	AccessConditions *AccessConditions
}

func FormatGetAccessControlOptions(o *GetAccessControlOptions) (*generated.PathClientGetPropertiesOptions, *generated.LeaseAccessConditions, *generated.ModifiedAccessConditions) {
	action := generated.PathGetPropertiesActionGetAccessControl
	if o == nil {
		return &generated.PathClientGetPropertiesOptions{
			Action: &action,
		}, nil, nil
	}
	// call path formatter since we're hitting dfs in this operation
	leaseAccessConditions, modifiedAccessConditions := exported.FormatPathAccessConditions(o.AccessConditions)
	return &generated.PathClientGetPropertiesOptions{
		Upn:    o.UPN,
		Action: &action,
	}, leaseAccessConditions, modifiedAccessConditions
}

// CPKInfo contains a group of parameters for the PathClient.Download method.
type CPKInfo struct {
	EncryptionAlgorithm *EncryptionAlgorithmType
	EncryptionKey       *string
	EncryptionKeySHA256 *string
}

// GetSASURLOptions contains the optional parameters for the Client.GetSASURL method.
type GetSASURLOptions struct {
	StartTime *time.Time
}

func FormatGetSASURLOptions(o *GetSASURLOptions) time.Time {
	if o == nil {
		return time.Time{}
	}

	var st time.Time
	if o.StartTime != nil {
		st = o.StartTime.UTC()
	} else {
		st = time.Time{}
	}
	return st
}

// SetHTTPHeadersOptions contains the optional parameters for the Client.SetHTTPHeaders method.
type SetHTTPHeadersOptions struct {
	AccessConditions *AccessConditions
}

func FormatSetHTTPHeadersOptions(o *SetHTTPHeadersOptions, httpHeaders HTTPHeaders) (*blob.SetHTTPHeadersOptions, blob.HTTPHeaders) {
	httpHeaderOpts := blob.HTTPHeaders{
		BlobCacheControl:       httpHeaders.CacheControl,
		BlobContentDisposition: httpHeaders.ContentDisposition,
		BlobContentEncoding:    httpHeaders.ContentEncoding,
		BlobContentLanguage:    httpHeaders.ContentLanguage,
		BlobContentMD5:         httpHeaders.ContentMD5,
		BlobContentType:        httpHeaders.ContentType,
	}
	if o == nil {
		return nil, httpHeaderOpts
	}
	accessConditions := exported.FormatBlobAccessConditions(o.AccessConditions)
	return &blob.SetHTTPHeadersOptions{
		AccessConditions: accessConditions,
	}, httpHeaderOpts
}

// HTTPHeaders contains the HTTP headers for path operations.
type HTTPHeaders struct {
	// Sets the path's cache control. If specified, this property is stored with the path and returned with a read request.
	CacheControl *string
	// Sets the path's Content-Disposition header.
	ContentDisposition *string
	// Sets the path's content encoding. If specified, this property is stored with the blobpath and returned with a read
	// request.
	ContentEncoding *string
	// Set the path's content language. If specified, this property is stored with the path and returned with a read
	// request.
	ContentLanguage *string
	// Specify the transactional md5 for the body, to be validated by the service.
	ContentMD5 []byte
	// Sets the path's content type. If specified, this property is stored with the path and returned with a read request.
	ContentType *string
}

//
//func (o HTTPHeaders) formatBlobHTTPHeaders() blob.HTTPHeaders {
//
//	opts := blob.HTTPHeaders{
//		BlobCacheControl:       o.CacheControl,
//		BlobContentDisposition: o.ContentDisposition,
//		BlobContentEncoding:    o.ContentEncoding,
//		BlobContentLanguage:    o.ContentLanguage,
//		BlobContentMD5:         o.ContentMD5,
//		BlobContentType:        o.ContentType,
//	}
//	return opts
//}

func FormatPathHTTPHeaders(o *HTTPHeaders) *generated.PathHTTPHeaders {
	// TODO: will be used for file related ops, like append
	if o == nil {
		return nil
	}
	opts := generated.PathHTTPHeaders{
		CacheControl:             o.CacheControl,
		ContentDisposition:       o.ContentDisposition,
		ContentEncoding:          o.ContentEncoding,
		ContentLanguage:          o.ContentLanguage,
		ContentMD5:               o.ContentMD5,
		ContentType:              o.ContentType,
		TransactionalContentHash: o.ContentMD5,
	}
	return &opts
}

// SetMetadataOptions provides set of configurations for Set Metadata on path operation
type SetMetadataOptions struct {
	Metadata         map[string]*string
	AccessConditions *AccessConditions
	CPKInfo          *CPKInfo
	CPKScopeInfo     *CPKScopeInfo
}

func FormatSetMetadataOptions(o *SetMetadataOptions) (*blob.SetMetadataOptions, map[string]*string) {
	if o == nil {
		return nil, nil
	}
	accessConditions := exported.FormatBlobAccessConditions(o.AccessConditions)
	opts := &blob.SetMetadataOptions{
		AccessConditions: accessConditions,
	}
	if o.CPKInfo != nil {
		opts.CPKInfo = &blob.CPKInfo{
			EncryptionKey:       o.CPKInfo.EncryptionKey,
			EncryptionAlgorithm: o.CPKInfo.EncryptionAlgorithm,
			EncryptionKeySHA256: o.CPKInfo.EncryptionKeySHA256,
		}
	}
	if o.CPKScopeInfo != nil {
		opts.CPKScopeInfo = (*blob.CPKScopeInfo)(o.CPKScopeInfo)
	}
	return opts, o.Metadata
}

// ========================================= constants =========================================

// SharedKeyCredential contains an account's name and its primary or secondary key.
type SharedKeyCredential = exported.SharedKeyCredential

// AccessConditions identifies access conditions which you optionally set.
type AccessConditions = exported.AccessConditions

// SourceAccessConditions identifies source access conditions which you optionally set.
type SourceAccessConditions = exported.SourceAccessConditions

// LeaseAccessConditions contains optional parameters to access leased entity.
type LeaseAccessConditions = exported.LeaseAccessConditions

// ModifiedAccessConditions contains a group of parameters for specifying access conditions.
type ModifiedAccessConditions = exported.ModifiedAccessConditions

// SourceModifiedAccessConditions contains a group of parameters for specifying access conditions.
type SourceModifiedAccessConditions = exported.SourceModifiedAccessConditions

// CPKScopeInfo contains a group of parameters for the Client.SetMetadata() method.
type CPKScopeInfo blob.CPKScopeInfo
