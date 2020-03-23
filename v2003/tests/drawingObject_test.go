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
	"os"
	"path"
	"path/filepath"
	"testing"
)

func TestDeleteDrawingObject(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DrawingObjects")
	remoteName := "TestDeleteDrawingObject.docx"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteDrawingObject(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteDrawingObjectWithoutNodePath(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DrawingObjects")
	remoteName := "TestDeleteDrawingObjectWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteDrawingObjectWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocumentDrawingObjectByIndex(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DrawingObjects")
	remoteName := "TestGetDocumentDrawingObjectByIndex.docx"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetDocumentDrawingObjectByIndex(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocumentDrawingObjectByIndexWithoutNodePath(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DrawingObjects")
	remoteName := "TestGetDocumentDrawingObjectByIndexWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetDocumentDrawingObjectByIndexWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocumentDrawingObjectImageData(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DrawingObjects")
	remoteName := "TestGetDocumentDrawingObjectImageData.docx"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	output, err := client.WordsApi.GetDocumentDrawingObjectImageData(ctx, remoteName, nodePath, int32(index), options)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocumentDrawingObjectImageDataWithoutNodePath(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DrawingObjects")
	remoteName := "TestGetDocumentDrawingObjectImageDataWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	output, err := client.WordsApi.GetDocumentDrawingObjectImageDataWithoutNodePath(ctx, remoteName, int32(index), options)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocumentDrawingObjectOleData(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "DrawingObjects"), "sample_EmbeddedOLE.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DrawingObjects")
	remoteName := "TestGetDocumentDrawingObjectOleData.docx"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	output, err := client.WordsApi.GetDocumentDrawingObjectOleData(ctx, remoteName, nodePath, int32(index), options)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocumentDrawingObjectOleDataWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "DrawingObjects"), "sample_EmbeddedOLE.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DrawingObjects")
	remoteName := "TestGetDocumentDrawingObjectOleDataWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	output, err := client.WordsApi.GetDocumentDrawingObjectOleDataWithoutNodePath(ctx, remoteName, int32(index), options)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocumentDrawingObjects(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DrawingObjects")
	remoteName := "TestGetDocumentDrawingObjects.docx"
	nodePath := "sections/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetDocumentDrawingObjects(ctx, remoteName, nodePath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocumentDrawingObjectsWithoutNodePath(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DrawingObjects")
	remoteName := "TestGetDocumentDrawingObjectsWithoutNodePath.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetDocumentDrawingObjectsWithoutNodePath(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertDrawingObject(t *testing.T) {

	localFilePath := commonTestFile
	localImagePath := GetLocalPath("Common", "aspose-cloud.png")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DrawingObjects")
	remoteName := "TestInsertDrawingObject.docx"
	imageFile, fileErr := os.Open(localImagePath)
	if fileErr != nil {
		t.Error(fileErr)
	}
	drawingObject := "{\"Left\": 0}"
	nodePath := "sections/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertDrawingObject(ctx, remoteName, drawingObject, imageFile, nodePath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertDrawingObjectWithoutNodePath(t *testing.T) {

	localFilePath := commonTestFile
	localImagePath := GetLocalPath("Common", "aspose-cloud.png")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DrawingObjects")
	remoteName := "TestInsertDrawingObjectWithoutNodePath.docx"
	imageFile, fileErr := os.Open(localImagePath)
	if fileErr != nil {
		t.Error(fileErr)
	}
	drawingObject := "{\"Left\": 0}"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertDrawingObjectWithoutNodePath(ctx, remoteName, drawingObject, imageFile, options)

	if err != nil {
		t.Error(err)
	}
}

func TestRenderDrawingObject(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DrawingObjects")
	remoteName := "TestRenderDrawingObject.docx"
	format := "png"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	output, err := client.WordsApi.RenderDrawingObject(ctx, remoteName, format, nodePath, int32(index), options)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}

func TestRenderDrawingObjectWithoutNodePath(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DrawingObjects")
	remoteName := "TestRenderDrawingObjectWithoutNodePath.docx"
	format := "png"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	output, err := client.WordsApi.RenderDrawingObjectWithoutNodePath(ctx, remoteName, format, int32(index), options)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateDrawingObject(t *testing.T) {

	localFilePath := commonTestFile
	localImagePath := GetLocalPath("Common", "aspose-cloud.png")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DrawingObjects")
	remoteName := "TestUpdateDrawingObject.docx"
	imageFile, fileErr := os.Open(localImagePath)
	if fileErr != nil {
		t.Error(fileErr)
	}
	nodePath := "sections/0"
	index := 0
	drawingObject := "{\"Left\": 0}"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateDrawingObject(ctx, remoteName, drawingObject, imageFile, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateDrawingObjectWithoutNodePath(t *testing.T) {

	localFilePath := commonTestFile
	localImagePath := GetLocalPath("Common", "aspose-cloud.png")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DrawingObjects")
	remoteName := "TestUpdateDrawingObjectWithoutNodePath.docx"
	imageFile, fileErr := os.Open(localImagePath)
	if fileErr != nil {
		t.Error(fileErr)
	}
	index := 0
	drawingObject := "{\"Left\": 0}"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateDrawingObjectWithoutNodePath(ctx, remoteName, drawingObject, imageFile, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}
