/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="bookmark_test.go">
 *   Copyright (c) 2022 Aspose.Words for Cloud
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

    request := &models.GetBookmarksRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetBookmarks(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for getting bookmarks from document online.
func Test_Bookmark_GetBookmarksOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.GetBookmarksOnlineRequest{
        Document: requestDocument,
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetBookmarksOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for getting bookmark by specified name.
func Test_Bookmark_GetBookmarkByName(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Bookmarks"
    localFile := "Common/test_multi_pages.docx"
    bookmarkName := "aspose"
    remoteFileName := "TestGetDocumentBookmarkByName.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.GetBookmarkByNameRequest{
        Name: ToStringPointer(remoteFileName),
        BookmarkName: ToStringPointer(bookmarkName),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetBookmarkByName(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for getting bookmark by specified name online.
func Test_Bookmark_GetBookmarkByNameOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"
    bookmarkName := "aspose"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.GetBookmarkByNameOnlineRequest{
        Document: requestDocument,
        BookmarkName: ToStringPointer(bookmarkName),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetBookmarkByNameOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for updating existed bookmark.
func Test_Bookmark_UpdateBookmark(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Bookmarks"
    localFile := "Common/test_multi_pages.docx"
    bookmarkName := "aspose"
    remoteFileName := "TestUpdateDocumentBookmark.docx"
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

    request := &models.UpdateBookmarkRequest{
        Name: ToStringPointer(remoteFileName),
        BookmarkName: ToStringPointer(bookmarkName),
        BookmarkData: requestBookmarkData,
        Optionals: options,
    }

    _, _, err := client.WordsApi.UpdateBookmark(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for updating existed bookmark online.
func Test_Bookmark_UpdateBookmarkOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"
    bookmarkName := "aspose"
    remoteFileName := "TestUpdateDocumentBookmark.docx"

    requestDocument := OpenFile(t, localFile)
    requestBookmarkData := models.BookmarkData{
        Name: ToStringPointer(bookmarkName),
        Text: ToStringPointer("This will be the text for Aspose"),
    }

    options := map[string]interface{}{
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }

    request := &models.UpdateBookmarkOnlineRequest{
        Document: requestDocument,
        BookmarkName: ToStringPointer(bookmarkName),
        BookmarkData: requestBookmarkData,
        Optionals: options,
    }

    _, _, err := client.WordsApi.UpdateBookmarkOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}
