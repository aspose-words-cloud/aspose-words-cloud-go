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
	"path"
	"path/filepath"
	"testing"
)

func TestDeleteOfficeMathObject(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "MathObjects"), "MathObjects.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "MathObjects")
	remoteName := "TestDeleteOfficeMathObject.docx"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteOfficeMathObject(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteOfficeMathObjectWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "MathObjects"), "MathObjects.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "MathObjects")
	remoteName := "TestDeleteOfficeMathObjectWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteOfficeMathObjectWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetOfficeMathObject(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "MathObjects"), "MathObjects.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "MathObjects")
	remoteName := "TestGetOfficeMathObject.docx"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetOfficeMathObject(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetOfficeMathObjectWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "MathObjects"), "MathObjects.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "MathObjects")
	remoteName := "TestGetOfficeMathObjectWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetOfficeMathObjectWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetOfficeMathObjects(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "MathObjects"), "MathObjects.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "MathObjects")
	remoteName := "TestGetOfficeMathObjects.docx"
	nodePath := "sections/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetOfficeMathObjects(ctx, remoteName, nodePath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetOfficeMathObjectsWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "MathObjects"), "MathObjects.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "MathObjects")
	remoteName := "TestGetOfficeMathObjectsWithoutNodePath.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetOfficeMathObjectsWithoutNodePath(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestRenderMathObject(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "MathObjects"), "MathObjects.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "MathObjects")
	remoteName := "TestRenderMathObject.docx"
	format := "png"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	output, err := client.WordsApi.RenderMathObject(ctx, remoteName, format, nodePath, int32(index), options)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}

func TestRenderMathObjectWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "MathObjects"), "MathObjects.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "MathObjects")
	remoteName := "TestRenderMathObjectWithoutNodePath.docx"
	format := "png"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	output, err := client.WordsApi.RenderMathObjectWithoutNodePath(ctx, remoteName, format, int32(index), options)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}
