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

	"github.com/aspose-words-cloud/aspose-words-cloud-go/v2004/api/models"
)

func TestDeleteBorder(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestDeleteBorder.docx"
	nodePath := "tables/1/rows/0/cells/0"
	borderType := "Left"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.DeleteBorder(ctx, remoteName, nodePath, borderType, options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteBorders(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestDeleteBorders.docx"
	nodePath := "tables/1/rows/0/cells/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.DeleteBorders(ctx, remoteName, nodePath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetBorder(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestGetBorder.docx"
	nodePath := "tables/1/rows/0/cells/0"
	borderType := "Left"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetBorder(ctx, remoteName, nodePath, borderType, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetBorders(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestGetBorders.docx"
	nodePath := "tables/1/rows/0/cells/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetBorders(ctx, remoteName, nodePath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateBorder(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestUpdateBorder.docx"
	nodePath := "tables/1/rows/0/cells/0"
	borderType := "Left"
	borderProperties := models.Border{
		BorderType: "Left",
		Color: &models.XmlColor{
			Alpha: 2,
		},
		DistanceFromText: 6.0,
		LineStyle:        "DashDotStroker",
		LineWidth:        2.0,
		Shadow:           true,
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateBorder(ctx, remoteName, borderProperties, nodePath, borderType, options)

	if err != nil {
		t.Error(err)
	}
}
