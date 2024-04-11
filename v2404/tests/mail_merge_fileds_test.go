/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="mail_merge_fileds_test.go">
 *   Copyright (c) 2024 Aspose.Words for Cloud
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

// Example of how to work with merge fields.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2404/api/models"
)

// Test for putting new fields.
func Test_MailMergeFileds_GetDocumentFieldNamesOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    mailMergeFolder := "DocumentActions/MailMerge"
    localDocumentFile := "SampleExecuteTemplate.docx"

    requestTemplate := OpenFile(t, mailMergeFolder + "/" + localDocumentFile)

    options := map[string]interface{}{
        "useNonMergeFields": true,
    }

    request := &models.GetDocumentFieldNamesOnlineRequest{
        Template: requestTemplate,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetDocumentFieldNamesOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetFieldNames(), "Validate GetDocumentFieldNamesOnline response.");
    assert.NotNil(t, actual.GetFieldNames().GetNames(), "Validate GetDocumentFieldNamesOnline response.");
    assert.Equal(t, 15, len(actual.GetFieldNames().GetNames()), "Validate GetDocumentFieldNamesOnline response.");
    assert.Equal(t, "TableStart:Order", DereferenceValue(actual.GetFieldNames().GetNames()[0]), "Validate GetDocumentFieldNamesOnline response.");
}

// Test for getting mailmerge fields.
func Test_MailMergeFileds_GetDocumentFieldNames(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentActions/MailMerge"
    remoteFileName := "TestGetDocumentFieldNames.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile("Common/test_multi_pages.docx"), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.GetDocumentFieldNamesRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetDocumentFieldNames(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetFieldNames(), "Validate GetDocumentFieldNames response.");
    assert.NotNil(t, actual.GetFieldNames().GetNames(), "Validate GetDocumentFieldNames response.");
    assert.Equal(t, 0, len(actual.GetFieldNames().GetNames()), "Validate GetDocumentFieldNames response.");
}
