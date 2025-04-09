/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="execute_template_with_field_options_test.go">
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

// Example of how to perform template execution.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2504/api/models"
)

// Test for posting execute template.
func Test_ExecuteTemplateWithFieldOptions_ExecuteTemplateWithFieldOptions(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentActions/MailMerge"
    mailMergeFolder := "DocumentActions/MailMerge"
    localDocumentFile := "TestMailMergeWithOptions.docx"
    remoteFileName := "TestMailMergeWithOptions.docx"
    localDataFile := ReadFile(t, mailMergeFolder + "/TestMailMergeData.xml")

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(mailMergeFolder + "/" + localDocumentFile), remoteDataFolder + "/" + remoteFileName)

    requestOptionsCurrentUser := models.UserInformation{
        Name: ToStringPointer("SdkTestUser"),
    }
    requestOptions := models.FieldOptions{
        CurrentUser: &requestOptionsCurrentUser,
    }

    options := map[string]interface{}{
        "data": localDataFile,
        "options": &requestOptions,
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }

    request := &models.ExecuteMailMergeRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.ExecuteMailMerge(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetDocument(), "Validate ExecuteTemplateWithFieldOptions response.");
    assert.Equal(t, "TestMailMergeWithOptions.docx", DereferenceValue(actual.GetDocument().GetFileName()), "Validate ExecuteTemplateWithFieldOptions response.");
}

// Test for execute template online.
func Test_ExecuteTemplateWithFieldOptions_ExecuteTemplateOnlineWithFieldOptions(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    mailMergeFolder := "DocumentActions/MailMerge"
    localDocumentFile := "TestMailMergeWithOptions.docx"
    localDataFile := "TestMailMergeData.xml"

    requestTemplate := OpenFile(t, mailMergeFolder + "/" + localDocumentFile)
    requestData := OpenFile(t, mailMergeFolder + "/" + localDataFile)
    requestOptionsCurrentUser := models.UserInformation{
        Name: ToStringPointer("SdkTestUser"),
    }
    requestOptions := models.FieldOptions{
        CurrentUser: &requestOptionsCurrentUser,
    }

    options := map[string]interface{}{
        "options": &requestOptions,
    }

    request := &models.ExecuteMailMergeOnlineRequest{
        Template: requestTemplate,
        Data: requestData,
        Optionals: options,
    }

    _, err := client.WordsApi.ExecuteMailMergeOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}
