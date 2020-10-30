/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="document_protection_test.go">
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

// Example of how to set document protection.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/dev/api/models"
)

// Test for setting document protection.
func Test_DocumentProtection_ProtectDocument(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/DocumentProtection"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestProtectDocument.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestProtectionRequest := models.ProtectionRequest{
        Password: "123",
        ProtectionType: "ReadOnly",
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }
    actual, _, err := client.WordsApi.ProtectDocument(ctx, remoteFileName, requestProtectionRequest, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.ProtectionData, "Validate ProtectDocument response.");
    assert.True(t, strings.HasPrefix(actual.ProtectionData.ProtectionType, "ReadOnly"), "Validate ProtectDocument response.");
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
    actual, _, err := client.WordsApi.GetDocumentProtection(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.ProtectionData, "Validate GetDocumentProtection response.");
    assert.True(t, strings.HasPrefix(actual.ProtectionData.ProtectionType, "ReadOnly"), "Validate GetDocumentProtection response.");
}

// Test for deleting unprotect document.
func Test_DocumentProtection_DeleteUnprotectDocument(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/DocumentProtection"
    localFilePath := "DocumentActions/DocumentProtection/SampleProtectedBlankWordDocument.docx"
    remoteFileName := "TestDeleteUnprotectDocument.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFilePath), remoteDataFolder + "/" + remoteFileName)

    requestProtectionRequest := models.ProtectionRequest{
        Password: "aspose",
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.UnprotectDocument(ctx, remoteFileName, requestProtectionRequest, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.ProtectionData, "Validate DeleteUnprotectDocument response.");
    assert.True(t, strings.HasPrefix(actual.ProtectionData.ProtectionType, "NoProtection"), "Validate DeleteUnprotectDocument response.");
}
