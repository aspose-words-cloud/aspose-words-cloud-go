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
	"testing"

	"github.com/aspose-words-cloud/aspose-words-cloud-go/v2003/api/models"
)

func TestGetBookmarkByName(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Bookmarks")
	remoteName := "TestGetBookmarkByName.docx"
	bookmarkName := "aspose"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetBookmarkByName(ctx, remoteName, bookmarkName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetBookmarks(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Bookmarks")
	remoteName := "TestGetBookmarks.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetBookmarks(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateBookmark(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Bookmarks")
	remoteName := "TestUpdateBookmark.docx"
	bookmarkName := "aspose"
	bookmarkData := models.BookmarkData{
		Name: bookmarkName,
		Text: "This will be the text for Aspose",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateBookmark(ctx, remoteName, bookmarkData, bookmarkName, options)

	if err != nil {
		t.Error(err)
	}
}
