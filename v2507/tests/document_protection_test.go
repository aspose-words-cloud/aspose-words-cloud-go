/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="document_protection_test.go">
 *   Copyright (c) 2025 Aspose.Words for Cloud
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

// Example of how to set document protection.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2507/api/models"
)

// Test for setting document protection.
func Test_DocumentProtection_ProtectDocument(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/DocumentProtection"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestProtectDocument.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestProtectionRequest := models.ProtectionRequestV2{
        ProtectionPassword: ToStringPointer("123"),
        ProtectionType: ToStringPointer("ReadOnly"),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }

    request := &models.ProtectDocumentRequest{
        Name: ToStringPointer(remoteFileName),
        ProtectionRequest: &requestProtectionRequest,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.ProtectDocument(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetProtectionData(), "Validate ProtectDocument response.");

}

// Test for setting document protection.
func Test_DocumentProtection_ProtectDocumentOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)
    requestProtectionRequest := models.ProtectionRequestV2{
        ProtectionPassword: ToStringPointer("123"),
        ProtectionType: ToStringPointer("ReadOnly"),
    }

    options := map[string]interface{}{
    }

    request := &models.ProtectDocumentOnlineRequest{
        Document: requestDocument,
        ProtectionRequest: &requestProtectionRequest,
        Optionals: options,
    }

    _, _, err := client.WordsApi.ProtectDocumentOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for getting document protection.
func Test_DocumentProtection_GetDocumentProtection(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/DocumentProtection"
    localFilePath := "DocumentActions/DocumentProtection/SampleProtectedBlankWordDocument.docx"
    remoteFileName := "TestGetDocumentProtection.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFilePath), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.GetDocumentProtectionRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetDocumentProtection(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for getting document protection.
func Test_DocumentProtection_GetDocumentProtectionOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.GetDocumentProtectionOnlineRequest{
        Document: requestDocument,
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetDocumentProtectionOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for deleting unprotect document.
func Test_DocumentProtection_DeleteUnprotectDocument(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/DocumentProtection"
    localFilePath := "DocumentActions/DocumentProtection/SampleProtectedBlankWordDocument.docx"
    remoteFileName := "TestDeleteUnprotectDocument.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFilePath), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.UnprotectDocumentRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.UnprotectDocument(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetProtectionData(), "Validate DeleteUnprotectDocument response.");

}

// Test for deleting unprotect document.
func Test_DocumentProtection_DeleteUnprotectDocumentOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFilePath := "DocumentActions/DocumentProtection/SampleProtectedBlankWordDocument.docx"

    requestDocument := OpenFile(t, localFilePath)

    options := map[string]interface{}{
    }

    request := &models.UnprotectDocumentOnlineRequest{
        Document: requestDocument,
        Optionals: options,
    }

    _, _, err := client.WordsApi.UnprotectDocumentOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}
