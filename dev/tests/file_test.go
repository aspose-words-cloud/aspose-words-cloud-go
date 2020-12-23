/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="file_test.go">
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

// Example of how to work with files.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/dev/api/models"
)

// Test for uploading file.
func Test_File_UploadFile(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/Storage"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestUploadFile.docx"


    options := map[string]interface{}{
    }

    request := &models.UploadFileRequest{
        FileContent: OpenFile(t, localFile),
        Path: ToStringPointer(remoteDataFolder + "/" + remoteFileName),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.UploadFile(ctx, request)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Uploaded, "Validate UploadFile response.");
    assert.Equal(t, 1, len(actual.Uploaded), "Validate UploadFile response.");
    assert.Equal(t, "TestUploadFile.docx", actual.Uploaded[0], "Validate UploadFile response.");
}

// Test for copy file.
func Test_File_CopyFile(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/Storage"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestCopyFileSrc.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
    }

    request := &models.CopyFileRequest{
        DestPath: ToStringPointer(remoteDataFolder + "/TestCopyFileDest.docx"),
        SrcPath: ToStringPointer(remoteDataFolder + "/" + remoteFileName),
        Optionals: options,
    }

    _, err := client.WordsApi.CopyFile(ctx, request)

    if err != nil {
        t.Error(err)
    }

}

// Test for move file.
func Test_File_MoveFile(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/Storage"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestMoveFileSrc.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
    }

    request := &models.MoveFileRequest{
        DestPath: ToStringPointer(baseTestOutPath + "/TestMoveFileDest_" + CreateRandomGuid() + ".docx"),
        SrcPath: ToStringPointer(remoteDataFolder + "/" + remoteFileName),
        Optionals: options,
    }

    _, err := client.WordsApi.MoveFile(ctx, request)

    if err != nil {
        t.Error(err)
    }

}

// Test for delete file.
func Test_File_DeleteFile(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/Storage"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestDeleteFile.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
    }

    request := &models.DeleteFileRequest{
        Path: ToStringPointer(remoteDataFolder + "/" + remoteFileName),
        Optionals: options,
    }

    _, err := client.WordsApi.DeleteFile(ctx, request)

    if err != nil {
        t.Error(err)
    }

}

// Test for download file.
func Test_File_DownloadFile(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/Storage"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestDownloadFile.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
    }

    request := &models.DownloadFileRequest{
        Path: ToStringPointer(remoteDataFolder + "/" + remoteFileName),
        Optionals: options,
    }

    _, , _, err := client.WordsApi.DownloadFile(ctx, request)

    if err != nil {
        t.Error(err)
    }

}
