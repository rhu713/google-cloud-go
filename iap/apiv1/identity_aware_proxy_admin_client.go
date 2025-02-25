// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package iap

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	iappb "google.golang.org/genproto/googleapis/cloud/iap/v1"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newIdentityAwareProxyAdminClientHook clientHook

// IdentityAwareProxyAdminCallOptions contains the retry settings for each method of IdentityAwareProxyAdminClient.
type IdentityAwareProxyAdminCallOptions struct {
	SetIamPolicy          []gax.CallOption
	GetIamPolicy          []gax.CallOption
	TestIamPermissions    []gax.CallOption
	GetIapSettings        []gax.CallOption
	UpdateIapSettings     []gax.CallOption
	ListTunnelDestGroups  []gax.CallOption
	CreateTunnelDestGroup []gax.CallOption
	GetTunnelDestGroup    []gax.CallOption
	DeleteTunnelDestGroup []gax.CallOption
	UpdateTunnelDestGroup []gax.CallOption
}

func defaultIdentityAwareProxyAdminGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("iap.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("iap.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://iap.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultIdentityAwareProxyAdminCallOptions() *IdentityAwareProxyAdminCallOptions {
	return &IdentityAwareProxyAdminCallOptions{
		SetIamPolicy:          []gax.CallOption{},
		GetIamPolicy:          []gax.CallOption{},
		TestIamPermissions:    []gax.CallOption{},
		GetIapSettings:        []gax.CallOption{},
		UpdateIapSettings:     []gax.CallOption{},
		ListTunnelDestGroups:  []gax.CallOption{},
		CreateTunnelDestGroup: []gax.CallOption{},
		GetTunnelDestGroup:    []gax.CallOption{},
		DeleteTunnelDestGroup: []gax.CallOption{},
		UpdateTunnelDestGroup: []gax.CallOption{},
	}
}

// internalIdentityAwareProxyAdminClient is an interface that defines the methods available from Cloud Identity-Aware Proxy API.
type internalIdentityAwareProxyAdminClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	SetIamPolicy(context.Context, *iampb.SetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	GetIamPolicy(context.Context, *iampb.GetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	TestIamPermissions(context.Context, *iampb.TestIamPermissionsRequest, ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error)
	GetIapSettings(context.Context, *iappb.GetIapSettingsRequest, ...gax.CallOption) (*iappb.IapSettings, error)
	UpdateIapSettings(context.Context, *iappb.UpdateIapSettingsRequest, ...gax.CallOption) (*iappb.IapSettings, error)
	ListTunnelDestGroups(context.Context, *iappb.ListTunnelDestGroupsRequest, ...gax.CallOption) *TunnelDestGroupIterator
	CreateTunnelDestGroup(context.Context, *iappb.CreateTunnelDestGroupRequest, ...gax.CallOption) (*iappb.TunnelDestGroup, error)
	GetTunnelDestGroup(context.Context, *iappb.GetTunnelDestGroupRequest, ...gax.CallOption) (*iappb.TunnelDestGroup, error)
	DeleteTunnelDestGroup(context.Context, *iappb.DeleteTunnelDestGroupRequest, ...gax.CallOption) error
	UpdateTunnelDestGroup(context.Context, *iappb.UpdateTunnelDestGroupRequest, ...gax.CallOption) (*iappb.TunnelDestGroup, error)
}

// IdentityAwareProxyAdminClient is a client for interacting with Cloud Identity-Aware Proxy API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// APIs for Identity-Aware Proxy Admin configurations.
type IdentityAwareProxyAdminClient struct {
	// The internal transport-dependent client.
	internalClient internalIdentityAwareProxyAdminClient

	// The call options for this service.
	CallOptions *IdentityAwareProxyAdminCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *IdentityAwareProxyAdminClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *IdentityAwareProxyAdminClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *IdentityAwareProxyAdminClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// SetIamPolicy sets the access control policy for an Identity-Aware Proxy protected
// resource. Replaces any existing policy.
// More information about managing access via IAP can be found at:
// https://cloud.google.com/iap/docs/managing-access#managing_access_via_the_api (at https://cloud.google.com/iap/docs/managing-access#managing_access_via_the_api)
func (c *IdentityAwareProxyAdminClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.SetIamPolicy(ctx, req, opts...)
}

// GetIamPolicy gets the access control policy for an Identity-Aware Proxy protected
// resource.
// More information about managing access via IAP can be found at:
// https://cloud.google.com/iap/docs/managing-access#managing_access_via_the_api (at https://cloud.google.com/iap/docs/managing-access#managing_access_via_the_api)
func (c *IdentityAwareProxyAdminClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.GetIamPolicy(ctx, req, opts...)
}

// TestIamPermissions returns permissions that a caller has on the Identity-Aware Proxy protected
// resource.
// More information about managing access via IAP can be found at:
// https://cloud.google.com/iap/docs/managing-access#managing_access_via_the_api (at https://cloud.google.com/iap/docs/managing-access#managing_access_via_the_api)
func (c *IdentityAwareProxyAdminClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	return c.internalClient.TestIamPermissions(ctx, req, opts...)
}

// GetIapSettings gets the IAP settings on a particular IAP protected resource.
func (c *IdentityAwareProxyAdminClient) GetIapSettings(ctx context.Context, req *iappb.GetIapSettingsRequest, opts ...gax.CallOption) (*iappb.IapSettings, error) {
	return c.internalClient.GetIapSettings(ctx, req, opts...)
}

// UpdateIapSettings updates the IAP settings on a particular IAP protected resource. It
// replaces all fields unless the update_mask is set.
func (c *IdentityAwareProxyAdminClient) UpdateIapSettings(ctx context.Context, req *iappb.UpdateIapSettingsRequest, opts ...gax.CallOption) (*iappb.IapSettings, error) {
	return c.internalClient.UpdateIapSettings(ctx, req, opts...)
}

// ListTunnelDestGroups lists the existing TunnelDestGroups. To group across all locations, use a
// - as the location ID. For example:
// /v1/projects/123/iap_tunnel/locations/-/destGroups
func (c *IdentityAwareProxyAdminClient) ListTunnelDestGroups(ctx context.Context, req *iappb.ListTunnelDestGroupsRequest, opts ...gax.CallOption) *TunnelDestGroupIterator {
	return c.internalClient.ListTunnelDestGroups(ctx, req, opts...)
}

// CreateTunnelDestGroup creates a new TunnelDestGroup.
func (c *IdentityAwareProxyAdminClient) CreateTunnelDestGroup(ctx context.Context, req *iappb.CreateTunnelDestGroupRequest, opts ...gax.CallOption) (*iappb.TunnelDestGroup, error) {
	return c.internalClient.CreateTunnelDestGroup(ctx, req, opts...)
}

// GetTunnelDestGroup retrieves an existing TunnelDestGroup.
func (c *IdentityAwareProxyAdminClient) GetTunnelDestGroup(ctx context.Context, req *iappb.GetTunnelDestGroupRequest, opts ...gax.CallOption) (*iappb.TunnelDestGroup, error) {
	return c.internalClient.GetTunnelDestGroup(ctx, req, opts...)
}

// DeleteTunnelDestGroup deletes a TunnelDestGroup.
func (c *IdentityAwareProxyAdminClient) DeleteTunnelDestGroup(ctx context.Context, req *iappb.DeleteTunnelDestGroupRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteTunnelDestGroup(ctx, req, opts...)
}

// UpdateTunnelDestGroup updates a TunnelDestGroup.
func (c *IdentityAwareProxyAdminClient) UpdateTunnelDestGroup(ctx context.Context, req *iappb.UpdateTunnelDestGroupRequest, opts ...gax.CallOption) (*iappb.TunnelDestGroup, error) {
	return c.internalClient.UpdateTunnelDestGroup(ctx, req, opts...)
}

// identityAwareProxyAdminGRPCClient is a client for interacting with Cloud Identity-Aware Proxy API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type identityAwareProxyAdminGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing IdentityAwareProxyAdminClient
	CallOptions **IdentityAwareProxyAdminCallOptions

	// The gRPC API client.
	identityAwareProxyAdminClient iappb.IdentityAwareProxyAdminServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewIdentityAwareProxyAdminClient creates a new identity aware proxy admin service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// APIs for Identity-Aware Proxy Admin configurations.
func NewIdentityAwareProxyAdminClient(ctx context.Context, opts ...option.ClientOption) (*IdentityAwareProxyAdminClient, error) {
	clientOpts := defaultIdentityAwareProxyAdminGRPCClientOptions()
	if newIdentityAwareProxyAdminClientHook != nil {
		hookOpts, err := newIdentityAwareProxyAdminClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := IdentityAwareProxyAdminClient{CallOptions: defaultIdentityAwareProxyAdminCallOptions()}

	c := &identityAwareProxyAdminGRPCClient{
		connPool:                      connPool,
		disableDeadlines:              disableDeadlines,
		identityAwareProxyAdminClient: iappb.NewIdentityAwareProxyAdminServiceClient(connPool),
		CallOptions:                   &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *identityAwareProxyAdminGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *identityAwareProxyAdminGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *identityAwareProxyAdminGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *identityAwareProxyAdminGRPCClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).SetIamPolicy[0:len((*c.CallOptions).SetIamPolicy):len((*c.CallOptions).SetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.identityAwareProxyAdminClient.SetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *identityAwareProxyAdminGRPCClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetIamPolicy[0:len((*c.CallOptions).GetIamPolicy):len((*c.CallOptions).GetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.identityAwareProxyAdminClient.GetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *identityAwareProxyAdminGRPCClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
	var resp *iampb.TestIamPermissionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.identityAwareProxyAdminClient.TestIamPermissions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *identityAwareProxyAdminGRPCClient) GetIapSettings(ctx context.Context, req *iappb.GetIapSettingsRequest, opts ...gax.CallOption) (*iappb.IapSettings, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetIapSettings[0:len((*c.CallOptions).GetIapSettings):len((*c.CallOptions).GetIapSettings)], opts...)
	var resp *iappb.IapSettings
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.identityAwareProxyAdminClient.GetIapSettings(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *identityAwareProxyAdminGRPCClient) UpdateIapSettings(ctx context.Context, req *iappb.UpdateIapSettingsRequest, opts ...gax.CallOption) (*iappb.IapSettings, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "iap_settings.name", url.QueryEscape(req.GetIapSettings().GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateIapSettings[0:len((*c.CallOptions).UpdateIapSettings):len((*c.CallOptions).UpdateIapSettings)], opts...)
	var resp *iappb.IapSettings
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.identityAwareProxyAdminClient.UpdateIapSettings(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *identityAwareProxyAdminGRPCClient) ListTunnelDestGroups(ctx context.Context, req *iappb.ListTunnelDestGroupsRequest, opts ...gax.CallOption) *TunnelDestGroupIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListTunnelDestGroups[0:len((*c.CallOptions).ListTunnelDestGroups):len((*c.CallOptions).ListTunnelDestGroups)], opts...)
	it := &TunnelDestGroupIterator{}
	req = proto.Clone(req).(*iappb.ListTunnelDestGroupsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*iappb.TunnelDestGroup, string, error) {
		resp := &iappb.ListTunnelDestGroupsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.identityAwareProxyAdminClient.ListTunnelDestGroups(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetTunnelDestGroups(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *identityAwareProxyAdminGRPCClient) CreateTunnelDestGroup(ctx context.Context, req *iappb.CreateTunnelDestGroupRequest, opts ...gax.CallOption) (*iappb.TunnelDestGroup, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateTunnelDestGroup[0:len((*c.CallOptions).CreateTunnelDestGroup):len((*c.CallOptions).CreateTunnelDestGroup)], opts...)
	var resp *iappb.TunnelDestGroup
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.identityAwareProxyAdminClient.CreateTunnelDestGroup(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *identityAwareProxyAdminGRPCClient) GetTunnelDestGroup(ctx context.Context, req *iappb.GetTunnelDestGroupRequest, opts ...gax.CallOption) (*iappb.TunnelDestGroup, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetTunnelDestGroup[0:len((*c.CallOptions).GetTunnelDestGroup):len((*c.CallOptions).GetTunnelDestGroup)], opts...)
	var resp *iappb.TunnelDestGroup
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.identityAwareProxyAdminClient.GetTunnelDestGroup(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *identityAwareProxyAdminGRPCClient) DeleteTunnelDestGroup(ctx context.Context, req *iappb.DeleteTunnelDestGroupRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteTunnelDestGroup[0:len((*c.CallOptions).DeleteTunnelDestGroup):len((*c.CallOptions).DeleteTunnelDestGroup)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.identityAwareProxyAdminClient.DeleteTunnelDestGroup(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *identityAwareProxyAdminGRPCClient) UpdateTunnelDestGroup(ctx context.Context, req *iappb.UpdateTunnelDestGroupRequest, opts ...gax.CallOption) (*iappb.TunnelDestGroup, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "tunnel_dest_group.name", url.QueryEscape(req.GetTunnelDestGroup().GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateTunnelDestGroup[0:len((*c.CallOptions).UpdateTunnelDestGroup):len((*c.CallOptions).UpdateTunnelDestGroup)], opts...)
	var resp *iappb.TunnelDestGroup
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.identityAwareProxyAdminClient.UpdateTunnelDestGroup(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// TunnelDestGroupIterator manages a stream of *iappb.TunnelDestGroup.
type TunnelDestGroupIterator struct {
	items    []*iappb.TunnelDestGroup
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*iappb.TunnelDestGroup, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *TunnelDestGroupIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *TunnelDestGroupIterator) Next() (*iappb.TunnelDestGroup, error) {
	var item *iappb.TunnelDestGroup
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *TunnelDestGroupIterator) bufLen() int {
	return len(it.items)
}

func (it *TunnelDestGroupIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
