package api_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/aspose-words-cloud/aspose-words-cloud-go/v2004/api"
	"github.com/aspose-words-cloud/aspose-words-cloud-go/v2004/api/models"
	"github.com/stretchr/testify/assert"
)

func TestAppSidAndAppKey(t *testing.T) {
	p := []struct {
		appKey string
		appSid string
	}{
		{"", "x"},
		{"x", ""},
	}
	for _, cp := range p {
		config := models.Configuration{
			AppKey: cp.appKey,
			AppSid: cp.appSid,
		}
		_, err := api.NewAPIClient(&config)

		assert.Error(t, err)
	}
}

func TestBaseUrl(t *testing.T) {
	config := models.Configuration{
		AppKey:  "x",
		AppSid:  "x",
		BaseUrl: "x",
	}
	_, err := api.NewAPIClient(&config)

	assert.Error(t, err)
}

func TestWrongAppKey(t *testing.T) {
	config := ReadConfiguration(t)
	config.AppKey = "x"
	_, _, err := api.CreateWordsApi(config)

	assert.Error(t, err)
}

func TestWrongAppSid(t *testing.T) {
	config := ReadConfiguration(t)
	config.AppSid = "x"
	_, _, err := api.CreateWordsApi(config)

	assert.Error(t, err)
}

func TestUnathorizedAccess(t *testing.T) {
	config := ReadConfiguration(t)
	client, err := api.NewAPIClient(config)

	if err != nil {
		t.Error(err)
	}

	format := "pdf"
	localFilePath := GetLocalPath(filepath.Join("DocumentActions", "ConvertDocument"), "test_uploadfile.docx")
	document, fileErr := os.Open(localFilePath)
	if fileErr != nil {
		t.Error(fileErr)
	}

	_, err = client.WordsApi.ConvertDocument(context.Background(), document, format, nil)

	assert.Error(t, err)
	assert.Equal(t, err.Error(), "Access is denied")
}
