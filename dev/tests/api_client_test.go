/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="api_client_test.go">
 *   Copyright (c) 2020 Aspose.Words for Cloud
 * </copyright>
 * <summary>
 *   Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the "Software"), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 * 
 *  The above copyright notice and this permission notice shall be included in all
 *  copies or substantial portions of the Software.
 * 
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *  SOFTWARE.
 * </summary>
 * --------------------------------------------------------------------------------
 */

package api_test

import (
    "context"
    "os"
    "path/filepath"
    "strings"
    "testing"

    "github.com/aspose-words-cloud/aspose-words-cloud-go/dev/api"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/dev/api/models"
    "github.com/stretchr/testify/assert"
)

func TestClientIdAndClientSecret(t *testing.T) {
    p := []struct {
        clientId string
        clientSecret string
    }{
        {"", "x"},
        {"x", ""},
    }
    for _, cp := range p {
        config := models.Configuration{
            ClientId: cp.clientId,
            ClientSecret: cp.clientSecret,
        }
        _, err := api.NewAPIClient(&config)

        assert.Error(t, err)
    }
}

func TestBaseUrl(t *testing.T) {
    config := models.Configuration{
        ClientId:  "x",
        ClientSecret:  "x",
        BaseUrl: "x",
    }
    _, err := api.NewAPIClient(&config)

    assert.Error(t, err)
}

func TestWrongClientSecret(t *testing.T) {
    config := ReadConfiguration(t)
    config.ClientSecret = "x"
    _, _, err := api.CreateWordsApi(config)

    assert.Error(t, err)
}

func TestWrongClientId(t *testing.T) {
    config := ReadConfiguration(t)
    config.ClientId = "x"
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

	request := &models.ConvertDocumentRequest{
		Document: document,
		Format:   ToStringPointer(format),
	}

	_, err = client.WordsApi.ConvertDocument(context.Background(), request)

    assert.Error(t, err)
}

