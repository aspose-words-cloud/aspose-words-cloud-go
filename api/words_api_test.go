//
// MIT License

// Copyright (c) 2019 Aspose Pty Ltd

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

package api_test

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"testing"

	"github.com/aspose-words-cloud/aspose-words-cloud-go/api"
)

var remoteBaseTestDataFolder string = "Temp/SdkTests/TestData"

func GetLocalPath(folderPath string, filename string) string {
	_, currentFilePath, _, _ := runtime.Caller(0)
	currentFolder := filepath.Dir(currentFilePath)
	return filepath.Join(currentFolder, "testdata", folderPath, filename)
}

func ReadConfiguration(t *testing.T) (cfg *api.Configuration) {
	_, filename, _, _ := runtime.Caller(0)
	configFile := filepath.Join(filepath.Dir(filename), "../config.json")
	configuration, err := api.NewConfiguration(configFile)

	if err != nil {
		t.Error(err)
	}

	return configuration
}

func PrepareTest(t *testing.T, config *api.Configuration) (apiClient *api.APIClient, ctx context.Context) {
	client, err := api.NewAPIClient(config)

	if err != nil {
		t.Error(err)
	}

	context, err := client.NewContextWithToken(nil)

	if err != nil {
		t.Error(err)
	}

	return client, context
}

func UploadFileToStorage(t *testing.T, fileName string, path string) (*api.APIClient, context.Context) {
	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	file, fileErr := os.Open(fileName)
	if fileErr != nil {
		t.Error(fileErr)
	}

	_, _, err := client.WordsApi.UploadFile(ctx, file, path, nil)
	if err != nil {
		t.Error(err)
	}

	return client, ctx
}

func TestDebugMode(t *testing.T) {
	config := ReadConfiguration(t)
	config.DebugMode = true
	client, ctx := PrepareTest(t, config)

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	_, _, err := client.WordsApi.CreateDocument(ctx, map[string]interface{}{
		"fileName": "test.docx",
	})

	if err != nil {
		t.Error(err)
	}

	logData := buf.String()

	url := "PUT /v4.0/words/create?FileName=test.docx"
	success := " 200 OK"

	if !strings.Contains(logData, url) {
		t.Errorf("Output doesn't contain the string '%s'", url)
	}

	if !strings.Contains(logData, success) {
		t.Errorf("Output doesn't contain the string '%s'", success)
	}
}

func TestTestCoverage(t *testing.T) {
	_, testFileName, _, _ := runtime.Caller(0)
	apiFileName := filepath.Join(filepath.Dir(testFileName), "words_api.go")
	api, err := ioutil.ReadFile(apiFileName)

	if err != nil {
		t.Error(err)
	}

	declareFuncRegex := regexp.MustCompile(`func \(a \*WordsApiService\) (\w+)`)
	allApiFuncNames := declareFuncRegex.FindAllSubmatch(api, -1)

	var exists = struct{}{}

	notTestedApiFuncNames := make(map[string]struct{})

	for _, name := range allApiFuncNames {
		notTestedApiFuncNames[string(name[1])] = exists
	}

	tests, err := ioutil.ReadFile(testFileName)
	if err != nil {
		t.Error(err)
	}

	callFuncRegex := regexp.MustCompile(`\.WordsApiService\.(\w+)`)
	testedlApiFuncNames := callFuncRegex.FindAllSubmatch(tests, -1)

	for _, name := range testedlApiFuncNames {
		delete(notTestedApiFuncNames, string(name[1]))
	}

	if len(notTestedApiFuncNames) > 0 {
		report := "Not covered methods:\n"

		for name := range notTestedApiFuncNames {
			report = report + fmt.Sprintf("\t%s\n", name)
		}

		t.Error(report)
	}
}

func TestFieldDelete(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Fields"), "GetField.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Fields")
	remoteName := "TestFieldDelete.docx"
	nodePath := "sections/0/paragraphs/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteField(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}
