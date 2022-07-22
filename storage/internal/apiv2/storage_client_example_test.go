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

package storage_test

import (
	"context"

	storage "cloud.google.com/go/storage/internal/apiv2"
	storagepb "cloud.google.com/go/storage/internal/apiv2/stubs"
	"google.golang.org/api/iterator"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
)

func ExampleNewClient() {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleClient_DeleteBucket() {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagepb.DeleteBucketRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/internal/apiv2/stubs#DeleteBucketRequest.
	}
	err = c.DeleteBucket(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_GetBucket() {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagepb.GetBucketRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/internal/apiv2/stubs#GetBucketRequest.
	}
	resp, err := c.GetBucket(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CreateBucket() {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagepb.CreateBucketRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/internal/apiv2/stubs#CreateBucketRequest.
	}
	resp, err := c.CreateBucket(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListBuckets() {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagepb.ListBucketsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/internal/apiv2/stubs#ListBucketsRequest.
	}
	it := c.ListBuckets(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_LockBucketRetentionPolicy() {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagepb.LockBucketRetentionPolicyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/internal/apiv2/stubs#LockBucketRetentionPolicyRequest.
	}
	resp, err := c.LockBucketRetentionPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetIamPolicy() {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.GetIamPolicyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/iam/v1#GetIamPolicyRequest.
	}
	resp, err := c.GetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_SetIamPolicy() {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.SetIamPolicyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/iam/v1#SetIamPolicyRequest.
	}
	resp, err := c.SetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_TestIamPermissions() {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.TestIamPermissionsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/iam/v1#TestIamPermissionsRequest.
	}
	resp, err := c.TestIamPermissions(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateBucket() {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagepb.UpdateBucketRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/internal/apiv2/stubs#UpdateBucketRequest.
	}
	resp, err := c.UpdateBucket(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteNotification() {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagepb.DeleteNotificationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/internal/apiv2/stubs#DeleteNotificationRequest.
	}
	err = c.DeleteNotification(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_GetNotification() {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagepb.GetNotificationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/internal/apiv2/stubs#GetNotificationRequest.
	}
	resp, err := c.GetNotification(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CreateNotification() {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagepb.CreateNotificationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/internal/apiv2/stubs#CreateNotificationRequest.
	}
	resp, err := c.CreateNotification(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListNotifications() {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagepb.ListNotificationsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/internal/apiv2/stubs#ListNotificationsRequest.
	}
	it := c.ListNotifications(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_ComposeObject() {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagepb.ComposeObjectRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/internal/apiv2/stubs#ComposeObjectRequest.
	}
	resp, err := c.ComposeObject(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteObject() {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagepb.DeleteObjectRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/internal/apiv2/stubs#DeleteObjectRequest.
	}
	err = c.DeleteObject(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_CancelResumableWrite() {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagepb.CancelResumableWriteRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/internal/apiv2/stubs#CancelResumableWriteRequest.
	}
	resp, err := c.CancelResumableWrite(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetObject() {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagepb.GetObjectRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/internal/apiv2/stubs#GetObjectRequest.
	}
	resp, err := c.GetObject(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateObject() {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagepb.UpdateObjectRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/internal/apiv2/stubs#UpdateObjectRequest.
	}
	resp, err := c.UpdateObject(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListObjects() {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagepb.ListObjectsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/internal/apiv2/stubs#ListObjectsRequest.
	}
	it := c.ListObjects(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_RewriteObject() {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagepb.RewriteObjectRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/internal/apiv2/stubs#RewriteObjectRequest.
	}
	resp, err := c.RewriteObject(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_StartResumableWrite() {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagepb.StartResumableWriteRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/internal/apiv2/stubs#StartResumableWriteRequest.
	}
	resp, err := c.StartResumableWrite(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_QueryWriteStatus() {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagepb.QueryWriteStatusRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/internal/apiv2/stubs#QueryWriteStatusRequest.
	}
	resp, err := c.QueryWriteStatus(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetServiceAccount() {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagepb.GetServiceAccountRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/internal/apiv2/stubs#GetServiceAccountRequest.
	}
	resp, err := c.GetServiceAccount(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CreateHmacKey() {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagepb.CreateHmacKeyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/internal/apiv2/stubs#CreateHmacKeyRequest.
	}
	resp, err := c.CreateHmacKey(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteHmacKey() {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagepb.DeleteHmacKeyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/internal/apiv2/stubs#DeleteHmacKeyRequest.
	}
	err = c.DeleteHmacKey(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_GetHmacKey() {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagepb.GetHmacKeyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/internal/apiv2/stubs#GetHmacKeyRequest.
	}
	resp, err := c.GetHmacKey(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListHmacKeys() {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagepb.ListHmacKeysRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/internal/apiv2/stubs#ListHmacKeysRequest.
	}
	it := c.ListHmacKeys(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_UpdateHmacKey() {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &storagepb.UpdateHmacKeyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/internal/apiv2/stubs#UpdateHmacKeyRequest.
	}
	resp, err := c.UpdateHmacKey(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
