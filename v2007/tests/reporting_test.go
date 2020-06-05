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
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/aspose-words-cloud/aspose-words-cloud-go/v2007/api/models"
)

func TestBuildReport(t *testing.T) {

	localFolder := filepath.Join("DocumentActions", "Reporting")
	localFilePath := GetLocalPath(localFolder, "ReportTemplate.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "Reporting")
	remoteName := "TestBuildReport.docx"
	data, fileErr := ioutil.ReadFile(GetLocalPath(localFolder, "ReportData.json"))
	if fileErr != nil {
		t.Error(fileErr)
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))
	settings := models.ReportEngineSettings{DataSourceType: "Json", ReportBuildOptions: []models.ReportBuildOptions{"AllowMissingMembers", "RemoveEmptyParagraphs"}}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	_, _, err := client.WordsApi.BuildReport(ctx, remoteName, string(data), settings, options)

	if err != nil {
		t.Error(err)
	}
}

func TestBuildReportOnline(t *testing.T) {

	localFolder := filepath.Join("DocumentActions", "Reporting")
	localTemplatePath := GetLocalPath(localFolder, "ReportTemplate.docx")
	localDataPath := GetLocalPath(localFolder, "ReportData.json")

	template, err := os.Open(localTemplatePath)
	if err != nil {
		t.Error(err)
	}
	data, err := ioutil.ReadFile(localDataPath)
	if err != nil {
		t.Error(err)
	}

	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	settings := models.ReportEngineSettings{DataSourceType: "Json", DataSourceName: "persons"}

	_, err2 := client.WordsApi.BuildReportOnline(ctx, template, string(data), settings, nil)

	if err2 != nil {
		t.Error(err2)
	}
}
