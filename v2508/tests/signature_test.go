/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="signature_test.go">
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

// Example of how to work with signatures.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2508/api/models"
)

// Test for getting all document signatures.
func Test_Signature_GetSignatures(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteFolder := remoteBaseTestDataFolder + "/DocumentActions/Signature"
    localFolder := "DocumentActions/Signature"
    signedDocument := "signedDocument.docx"
    remoteName := "TestGetSignatures.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFolder + "/" + signedDocument), remoteFolder + "/" + remoteName)


    options := map[string]interface{}{
        "folder": remoteFolder,
    }

    request := &models.GetSignaturesRequest{
        Name: ToStringPointer(remoteName),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetSignatures(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetSignatures(), "Validate GetSignatures response.");
    assert.Equal(t, 1, len(actual.GetSignatures()), "Validate GetSignatures response.");
}

// Test for getting all document signatures online.
func Test_Signature_GetSignaturesOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFolder := "DocumentActions/Signature"
    signedDocument := "signedDocument.docx"

    requestDocument := OpenFile(t, localFolder + "/" + signedDocument)

    options := map[string]interface{}{
    }

    request := &models.GetSignaturesOnlineRequest{
        Document: requestDocument,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetSignaturesOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetSignatures(), "Validate GetSignaturesOnline response.");
    assert.Equal(t, 1, len(actual.GetSignatures()), "Validate GetSignaturesOnline response.");
}

// Test for removing all document signatures.
func Test_Signature_RemoveAllSignatures(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteFolder := remoteBaseTestDataFolder + "/DocumentActions/Signature"
    localFolder := "DocumentActions/Signature"
    signedDocument := "signedDocument.docx"
    remoteName := "TestRemoveAllSignatures.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFolder + "/" + signedDocument), remoteFolder + "/" + remoteName)


    options := map[string]interface{}{
        "folder": remoteFolder,
    }

    request := &models.RemoveAllSignaturesRequest{
        Name: ToStringPointer(remoteName),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.RemoveAllSignatures(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetSignatures(), "Validate RemoveAllSignatures response.");
    assert.Equal(t, 0, len(actual.GetSignatures()), "Validate RemoveAllSignatures response.");
}

// Test for removing all document signatures online.
func Test_Signature_RemoveAllSignaturesOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFolder := "DocumentActions/Signature"
    signedDocument := "signedDocument.docx"

    requestDocument := OpenFile(t, localFolder + "/" + signedDocument)

    options := map[string]interface{}{
    }

    request := &models.RemoveAllSignaturesOnlineRequest{
        Document: requestDocument,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.RemoveAllSignaturesOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetModel().GetSignatures(), "Validate RemoveAllSignaturesOnline response.");
    assert.Equal(t, 0, len(actual.GetModel().GetSignatures()), "Validate RemoveAllSignaturesOnline response.");
}

// Test for signing document.
func Test_Signature_SignDocument(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteFolder := remoteBaseTestDataFolder + "/DocumentActions/Signature"
    localFolder := "DocumentActions/Signature"
    unsignedDocument := "unsignedDocument.docx"
    certificateName := "morzal.pfx"
    certificatePassword := "aw"
    remoteName := "TestSignDocument.docx"
    remoteCertificateName := "TestCertificate.pfx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFolder + "/" + unsignedDocument), remoteFolder + "/" + remoteName)
    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFolder + "/" + certificateName), remoteFolder + "/" + remoteCertificateName)


    options := map[string]interface{}{
        "folder": remoteFolder,
    }

    request := &models.SignDocumentRequest{
        Name: ToStringPointer(remoteName),
        CertificatePath: ToStringPointer(remoteFolder + "/" + remoteCertificateName),
        CertificatePassword: ToStringPointer(certificatePassword),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.SignDocument(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetSignatures(), "Validate SignDocument response.");
    assert.Equal(t, 1, len(actual.GetSignatures()), "Validate SignDocument response.");
}

// Test for signing document online.
func Test_Signature_SignDocumentOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteFolder := remoteBaseTestDataFolder + "/DocumentActions/Signature"
    localFolder := "DocumentActions/Signature"
    unsignedDocument := "unsignedDocument.docx"
    certificateName := "morzal.pfx"
    certificatePassword := "aw"
    remoteCertificateName := "TestCertificateOnline.pfx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFolder + "/" + certificateName), remoteFolder + "/" + remoteCertificateName)

    requestDocument := OpenFile(t, localFolder + "/" + unsignedDocument)

    options := map[string]interface{}{
    }

    request := &models.SignDocumentOnlineRequest{
        Document: requestDocument,
        CertificatePath: ToStringPointer(remoteFolder + "/" + remoteCertificateName),
        CertificatePassword: ToStringPointer(certificatePassword),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.SignDocumentOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetModel().GetSignatures(), "Validate SignDocumentOnline response.");
    assert.Equal(t, 1, len(actual.GetModel().GetSignatures()), "Validate SignDocumentOnline response.");
}
