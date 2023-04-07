/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="folder_test.go">
 *   Copyright (c) 2023 Aspose.Words for Cloud
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

// Example of how to work with folders.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2303/api/models"
)

// Test for create folder.
func Test_Folder_CreateFolder(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/Storage"


    options := map[string]interface{}{
    }

    request := &models.CreateFolderRequest{
        Path: ToStringPointer(remoteDataFolder + "/TestCreateFolder"),
        Optionals: options,
    }

    _, err := client.WordsApi.CreateFolder(ctx, request)

    if err != nil {
        t.Error(err)
    }

}

// Test for delete folder.
func Test_Folder_DeleteFolder(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/Storage"
    localFile := "Common/test_multi_pages.docx"
    testDeleteFolder := remoteDataFolder + "/TestDeleteFolder"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), testDeleteFolder + "/TestDeleteFolder.docx")


    options := map[string]interface{}{
    }

    request := &models.DeleteFolderRequest{
        Path: ToStringPointer(testDeleteFolder),
        Optionals: options,
    }

    _, err := client.WordsApi.DeleteFolder(ctx, request)

    if err != nil {
        t.Error(err)
    }

}

// Test for get file list of folder.
func Test_Folder_GetFilesList(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/Storage"


    options := map[string]interface{}{
    }

    request := &models.GetFilesListRequest{
        Path: ToStringPointer(remoteDataFolder),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetFilesList(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Value, "Validate GetFilesList response.");
}

// Test for copy folder.
func Test_Folder_CopyFolder(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/Storage"
    localFile := "Common/test_multi_pages.docx"
    folderToCopy := remoteDataFolder + "/TestCopyFolder"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), folderToCopy + "Src/TestCopyFolderSrc.docx")


    options := map[string]interface{}{
    }

    request := &models.CopyFolderRequest{
        DestPath: ToStringPointer(folderToCopy + "Dest"),
        SrcPath: ToStringPointer(folderToCopy + "Src"),
        Optionals: options,
    }

    _, err := client.WordsApi.CopyFolder(ctx, request)

    if err != nil {
        t.Error(err)
    }

}

// Test for move folder.
func Test_Folder_MoveFolder(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/Storage"
    localFile := "Common/test_multi_pages.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/TestMoveFolderSrc/TestMoveFolderSrc.docx")


    options := map[string]interface{}{
    }

    request := &models.MoveFolderRequest{
        DestPath: ToStringPointer(baseTestOutPath + "/TestMoveFolderDest_" + CreateRandomGuid()),
        SrcPath: ToStringPointer(remoteDataFolder + "/TestMoveFolderSrc"),
        Optionals: options,
    }

    _, err := client.WordsApi.MoveFolder(ctx, request)

    if err != nil {
        t.Error(err)
    }

}
