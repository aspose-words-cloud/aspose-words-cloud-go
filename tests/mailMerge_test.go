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
	//"os"
	"path"
	"path/filepath"
	"testing"
)

func TestExecuteMailMerge(t *testing.T) {

	localFolder := filepath.Join("DocumentActions", "MailMerge")
	localFilePath := GetLocalPath(localFolder, "SampleMailMergeTemplate.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "MailMerge")
	remoteName := "TestExecuteMailMerge.docx"
	data, fileErr := ioutil.ReadFile(GetLocalPath(localFolder, "SampleMailMergeTemplateData.txt"))
	if fileErr != nil {
		t.Error(fileErr)
	}
	options := map[string]interface{}{
		"data":        string(data),
		"folder":      remoteFolder,
		"withRegions": false,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.ExecuteMailMerge(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

/*func TestExecuteMailMergeOnline(t *testing.T) {

	localFolder := filepath.Join("DocumentActions", "MailMerge")
	localTemplatePath := GetLocalPath(localFolder, "SampleMailMergeTemplate.docx")
	localDataPath := GetLocalPath(localFolder, "SampleMailMergeTemplateData.txt")

	template, err := os.Open(localTemplatePath)
	if err != nil {
		t.Error(err)
	}
	data, err := os.Open(localDataPath)
	if err != nil {
		t.Error(err)
	}

	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	output, err := client.WordsApi.ExecuteMailMergeOnline(ctx, template, data, nil)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}*/
