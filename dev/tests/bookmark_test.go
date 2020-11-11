/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="bookmark_test.go">
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

// Example of how to get all bookmarks from document.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/dev/api/models"
)

// Test for getting bookmarks from document.
func Test_Bookmark_GetBookmarks(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Bookmarks"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetDocumentBookmarks.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.GetBookmarks(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Bookmarks, "Validate GetBookmarks response.");
    assert.Equal(t, 3, len(actual.Bookmarks.BookmarkList), "Validate GetBookmarks response.");
    assert.Equal(t, "aspose", *actual.Bookmarks.BookmarkList[1].Name, "Validate GetBookmarks response.");
}

// Test for getting bookmark by specified name.
func Test_Bookmark_GetBookmarkByName(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Bookmarks"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetDocumentBookmarkByName.docx"
    bookmarkName := "aspose"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.GetBookmarkByName(ctx, remoteFileName, bookmarkName, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Bookmark, "Validate GetBookmarkByName response.");
    assert.Equal(t, bookmarkName, *actual.Bookmark.Name, "Validate GetBookmarkByName response.");
}

// Test for updating existed bookmark.
func Test_Bookmark_UpdateBookmark(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Bookmarks"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestUpdateDocumentBookmark.docx"
    bookmarkName := "aspose"
    bookmarkText := "This will be the text for Aspose"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestBookmarkData := models.BookmarkData{
        Name: ToStringPointer(bookmarkName),
        Text: ToStringPointer(bookmarkText),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }
    actual, _, err := client.WordsApi.UpdateBookmark(ctx, remoteFileName, requestBookmarkData, bookmarkName, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Bookmark, "Validate UpdateBookmark response.");
    assert.Equal(t, bookmarkName, *actual.Bookmark.Name, "Validate UpdateBookmark response.");
    assert.Equal(t, bookmarkText, *actual.Bookmark.Text, "Validate UpdateBookmark response.");
}
