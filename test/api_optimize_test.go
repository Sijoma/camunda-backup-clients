/*
Backup Management API

Testing OptimizeApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package camunda_backup_clients

import (
	"context"
	"testing"

	openapiclient "github.com/sijoma/camunda-backup-clients"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_camunda_backup_clients_OptimizeApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OptimizeApiService BackupIdDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var backupId int64

		httpRes, err := apiClient.OptimizeApi.BackupIdDelete(context.Background(), backupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OptimizeApiService BackupIdGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var backupId int64

		resp, httpRes, err := apiClient.OptimizeApi.BackupIdGet(context.Background(), backupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OptimizeApiService RootGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OptimizeApi.RootGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OptimizeApiService RootPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OptimizeApi.RootPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
