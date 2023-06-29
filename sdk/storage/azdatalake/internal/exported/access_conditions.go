package exported

import (
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azdatalake/internal/generated"
)

// AccessConditions identifies container-specific access conditions which you optionally set.
type AccessConditions struct {
	ModifiedAccessConditions *ModifiedAccessConditions
	LeaseAccessConditions    *LeaseAccessConditions
}

// LeaseAccessConditions contains optional parameters to access leased entity.
type LeaseAccessConditions = generated.LeaseAccessConditions

// ModifiedAccessConditions contains a group of parameters for specifying access conditions.
type ModifiedAccessConditions = generated.ModifiedAccessConditions

// FormatContainerAccessConditions formats FilesystemAccessConditions into container's LeaseAccessConditions and ModifiedAccessConditions.
func FormatContainerAccessConditions(b *AccessConditions) *container.AccessConditions {
	if b == nil {
		return nil
	}
	if b.LeaseAccessConditions == nil {
		b.LeaseAccessConditions = &LeaseAccessConditions{}
	}
	if b.ModifiedAccessConditions == nil {
		b.ModifiedAccessConditions = &ModifiedAccessConditions{}
	}
	return &container.AccessConditions{
		LeaseAccessConditions: &container.LeaseAccessConditions{
			LeaseID: b.LeaseAccessConditions.LeaseID,
		},
		ModifiedAccessConditions: &container.ModifiedAccessConditions{
			IfMatch:           b.ModifiedAccessConditions.IfMatch,
			IfNoneMatch:       b.ModifiedAccessConditions.IfNoneMatch,
			IfModifiedSince:   b.ModifiedAccessConditions.IfModifiedSince,
			IfUnmodifiedSince: b.ModifiedAccessConditions.IfUnmodifiedSince,
		},
	}
}

// FormatPathAccessConditions formats PathAccessConditions into path's LeaseAccessConditions and ModifiedAccessConditions.
func FormatPathAccessConditions(p *AccessConditions) (*generated.LeaseAccessConditions, *generated.ModifiedAccessConditions) {
	if p == nil {
		return nil, nil
	}
	if p.LeaseAccessConditions == nil {
		p.LeaseAccessConditions = &LeaseAccessConditions{}
	}
	if p.ModifiedAccessConditions == nil {
		p.ModifiedAccessConditions = &ModifiedAccessConditions{}
	}
	return &generated.LeaseAccessConditions{
			LeaseID: p.LeaseAccessConditions.LeaseID,
		}, &generated.ModifiedAccessConditions{
			IfMatch:           p.ModifiedAccessConditions.IfMatch,
			IfNoneMatch:       p.ModifiedAccessConditions.IfNoneMatch,
			IfModifiedSince:   p.ModifiedAccessConditions.IfModifiedSince,
			IfUnmodifiedSince: p.ModifiedAccessConditions.IfUnmodifiedSince,
		}
}

// FormatBlobAccessConditions formats PathAccessConditions into blob's LeaseAccessConditions and ModifiedAccessConditions.
func FormatBlobAccessConditions(p *AccessConditions) *blob.AccessConditions {
	if p == nil {
		return nil
	}
	if p.LeaseAccessConditions == nil {
		p.LeaseAccessConditions = &LeaseAccessConditions{}
	}
	if p.ModifiedAccessConditions == nil {
		p.ModifiedAccessConditions = &ModifiedAccessConditions{}
	}
	return &blob.AccessConditions{LeaseAccessConditions: &blob.LeaseAccessConditions{
		LeaseID: p.LeaseAccessConditions.LeaseID,
	}, ModifiedAccessConditions: &blob.ModifiedAccessConditions{
		IfMatch:           p.ModifiedAccessConditions.IfMatch,
		IfNoneMatch:       p.ModifiedAccessConditions.IfNoneMatch,
		IfModifiedSince:   p.ModifiedAccessConditions.IfModifiedSince,
		IfUnmodifiedSince: p.ModifiedAccessConditions.IfUnmodifiedSince,
	}}
}
