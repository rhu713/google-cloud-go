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

package dataform_test

import (
	"context"

	dataform "cloud.google.com/go/dataform/apiv1alpha2"
	"google.golang.org/api/iterator"
	dataformpb "google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
)

func ExampleNewClient() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleNewRESTClient() {
	ctx := context.Background()
	c, err := dataform.NewRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleClient_ListRepositories() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.ListRepositoriesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#ListRepositoriesRequest.
	}
	it := c.ListRepositories(ctx, req)
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

func ExampleClient_GetRepository() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.GetRepositoryRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#GetRepositoryRequest.
	}
	resp, err := c.GetRepository(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CreateRepository() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.CreateRepositoryRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#CreateRepositoryRequest.
	}
	resp, err := c.CreateRepository(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateRepository() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.UpdateRepositoryRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#UpdateRepositoryRequest.
	}
	resp, err := c.UpdateRepository(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteRepository() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.DeleteRepositoryRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#DeleteRepositoryRequest.
	}
	err = c.DeleteRepository(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_FetchRemoteBranches() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.FetchRemoteBranchesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#FetchRemoteBranchesRequest.
	}
	resp, err := c.FetchRemoteBranches(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListWorkspaces() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.ListWorkspacesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#ListWorkspacesRequest.
	}
	it := c.ListWorkspaces(ctx, req)
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

func ExampleClient_GetWorkspace() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.GetWorkspaceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#GetWorkspaceRequest.
	}
	resp, err := c.GetWorkspace(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CreateWorkspace() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.CreateWorkspaceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#CreateWorkspaceRequest.
	}
	resp, err := c.CreateWorkspace(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteWorkspace() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.DeleteWorkspaceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#DeleteWorkspaceRequest.
	}
	err = c.DeleteWorkspace(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_InstallNpmPackages() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.InstallNpmPackagesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#InstallNpmPackagesRequest.
	}
	resp, err := c.InstallNpmPackages(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_PullGitCommits() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.PullGitCommitsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#PullGitCommitsRequest.
	}
	err = c.PullGitCommits(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_PushGitCommits() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.PushGitCommitsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#PushGitCommitsRequest.
	}
	err = c.PushGitCommits(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_FetchFileGitStatuses() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.FetchFileGitStatusesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#FetchFileGitStatusesRequest.
	}
	resp, err := c.FetchFileGitStatuses(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_FetchGitAheadBehind() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.FetchGitAheadBehindRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#FetchGitAheadBehindRequest.
	}
	resp, err := c.FetchGitAheadBehind(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CommitWorkspaceChanges() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.CommitWorkspaceChangesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#CommitWorkspaceChangesRequest.
	}
	err = c.CommitWorkspaceChanges(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_ResetWorkspaceChanges() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.ResetWorkspaceChangesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#ResetWorkspaceChangesRequest.
	}
	err = c.ResetWorkspaceChanges(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_FetchFileDiff() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.FetchFileDiffRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#FetchFileDiffRequest.
	}
	resp, err := c.FetchFileDiff(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_QueryDirectoryContents() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.QueryDirectoryContentsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#QueryDirectoryContentsRequest.
	}
	it := c.QueryDirectoryContents(ctx, req)
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

func ExampleClient_MakeDirectory() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.MakeDirectoryRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#MakeDirectoryRequest.
	}
	resp, err := c.MakeDirectory(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_RemoveDirectory() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.RemoveDirectoryRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#RemoveDirectoryRequest.
	}
	err = c.RemoveDirectory(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_MoveDirectory() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.MoveDirectoryRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#MoveDirectoryRequest.
	}
	resp, err := c.MoveDirectory(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ReadFile() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.ReadFileRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#ReadFileRequest.
	}
	resp, err := c.ReadFile(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_RemoveFile() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.RemoveFileRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#RemoveFileRequest.
	}
	err = c.RemoveFile(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_MoveFile() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.MoveFileRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#MoveFileRequest.
	}
	resp, err := c.MoveFile(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_WriteFile() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.WriteFileRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#WriteFileRequest.
	}
	resp, err := c.WriteFile(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListCompilationResults() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.ListCompilationResultsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#ListCompilationResultsRequest.
	}
	it := c.ListCompilationResults(ctx, req)
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

func ExampleClient_GetCompilationResult() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.GetCompilationResultRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#GetCompilationResultRequest.
	}
	resp, err := c.GetCompilationResult(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CreateCompilationResult() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.CreateCompilationResultRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#CreateCompilationResultRequest.
	}
	resp, err := c.CreateCompilationResult(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_QueryCompilationResultActions() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.QueryCompilationResultActionsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#QueryCompilationResultActionsRequest.
	}
	it := c.QueryCompilationResultActions(ctx, req)
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

func ExampleClient_ListWorkflowInvocations() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.ListWorkflowInvocationsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#ListWorkflowInvocationsRequest.
	}
	it := c.ListWorkflowInvocations(ctx, req)
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

func ExampleClient_GetWorkflowInvocation() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.GetWorkflowInvocationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#GetWorkflowInvocationRequest.
	}
	resp, err := c.GetWorkflowInvocation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CreateWorkflowInvocation() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.CreateWorkflowInvocationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#CreateWorkflowInvocationRequest.
	}
	resp, err := c.CreateWorkflowInvocation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteWorkflowInvocation() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.DeleteWorkflowInvocationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#DeleteWorkflowInvocationRequest.
	}
	err = c.DeleteWorkflowInvocation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_CancelWorkflowInvocation() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.CancelWorkflowInvocationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#CancelWorkflowInvocationRequest.
	}
	err = c.CancelWorkflowInvocation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_QueryWorkflowInvocationActions() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.QueryWorkflowInvocationActionsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dataform/v1alpha2#QueryWorkflowInvocationActionsRequest.
	}
	it := c.QueryWorkflowInvocationActions(ctx, req)
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

func ExampleClient_GetLocation() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &locationpb.GetLocationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/location#GetLocationRequest.
	}
	resp, err := c.GetLocation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListLocations() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &locationpb.ListLocationsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/location#ListLocationsRequest.
	}
	it := c.ListLocations(ctx, req)
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

func ExampleClient_GetIamPolicy() {
	ctx := context.Background()
	c, err := dataform.NewClient(ctx)
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
	c, err := dataform.NewClient(ctx)
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
	c, err := dataform.NewClient(ctx)
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
