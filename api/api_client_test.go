package api_test

import (
	"testing"

	"github.com/aspose-words-cloud/aspose-words-cloud-go/api"
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
		config := api.Configuration{
			AppKey: cp.appKey,
			AppSid: cp.appSid,
		}
		_, err := api.NewAPIClient(&config)

		if err == nil {
			t.Error(err)
		}
	}
}

func TestBaseUrl(t *testing.T) {
	config := api.Configuration{
		AppKey:  "x",
		AppSid:  "x",
		BaseUrl: "x",
	}
	_, err := api.NewAPIClient(&config)

	if err == nil {
		t.Error(err)
	}
}
