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
	"github.com/aspose-words-cloud/aspose-words-cloud-go/v2007/api/models"
	"path"
    "path/filepath"
	"testing"
)

func TestGetLists(t *testing.T) {
	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Lists"), "ListsGet.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, filepath.Join("DocumentElements", "Lists"))
	remoteName := "TestGetLists.doc"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetLists(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetList(t *testing.T) {
	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Lists"), "ListsGet.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, filepath.Join("DocumentElements", "Lists"))
	remoteName := "TestGetList.doc"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetList(ctx, remoteName, 1, options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateList(t *testing.T) {
	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Lists"), "ListsGet.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, filepath.Join("DocumentElements", "Lists"))
	remoteName := "TestUpdateList.doc"
	listUpdate := models.ListUpdate{IsRestartAtEachSection:true}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateList(ctx, remoteName, listUpdate,1, options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateListLevel(t *testing.T) {
	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Lists"), "ListsGet.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, filepath.Join("DocumentElements", "Lists"))
	remoteName := "TestUpdateListLevel.doc"
	listUpdate := models.ListLevelUpdate{Alignment:"Right"}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateListLevel(ctx, remoteName, listUpdate,1, 1, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertList(t *testing.T) {
	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Lists"), "ListsGet.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, filepath.Join("DocumentElements", "Lists"))
	remoteName := "TestInsertList.doc"
	listInsert := models.ListInsert{Template:"OutlineLegal"}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertList(ctx, remoteName, listInsert, options)

	if err != nil {
		t.Error(err)
	}
}
