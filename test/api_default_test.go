/*
Backup Management API

Testing DefaultApiService

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

func Test_camunda_backup_clients_DefaultApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DefaultApiService BackupIdDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var backupId int64

		httpRes, err := apiClient.DefaultApi.BackupIdDelete(context.Background(), backupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService BackupIdGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var backupId int64

		resp, httpRes, err := apiClient.DefaultApi.BackupIdGet(context.Background(), backupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService RootGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DefaultApi.RootGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService RootPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DefaultApi.RootPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
