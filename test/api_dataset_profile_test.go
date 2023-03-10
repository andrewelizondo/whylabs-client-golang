/*
WhyLabs Songbird

Testing DatasetProfileApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_DatasetProfileApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DatasetProfileApiService DeleteAnalyzerResults", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var datasetId string

		resp, httpRes, err := apiClient.DatasetProfileApi.DeleteAnalyzerResults(context.Background(), orgId, datasetId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatasetProfileApiService DeleteDatasetProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var datasetId string

		resp, httpRes, err := apiClient.DatasetProfileApi.DeleteDatasetProfiles(context.Background(), orgId, datasetId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatasetProfileApiService DeleteReferenceProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var modelId string
		var referenceId string

		resp, httpRes, err := apiClient.DatasetProfileApi.DeleteReferenceProfile(context.Background(), orgId, modelId, referenceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatasetProfileApiService GetReferenceProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var modelId string
		var referenceId string

		resp, httpRes, err := apiClient.DatasetProfileApi.GetReferenceProfile(context.Background(), orgId, modelId, referenceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatasetProfileApiService ListReferenceProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var modelId string

		resp, httpRes, err := apiClient.DatasetProfileApi.ListReferenceProfiles(context.Background(), orgId, modelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
