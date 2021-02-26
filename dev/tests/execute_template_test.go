/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="execute_template_test.go">
 *   Copyright (c) 2021 Aspose.Words for Cloud
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
    "github.com/aspose-words-cloud/aspose-words-cloud-go/dev/api/models"
)

// Test for posting execute template.
func Test_ExecuteTemplate_ExecuteTemplate(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentActions/MailMerge"
    mailMergeFolder := "DocumentActions/MailMerge"
    localDocumentFile := "TestExecuteTemplate.doc"
    remoteFileName := "TestExecuteTemplate.docx"
    localDataFile := ReadFile(t, mailMergeFolder + "/TestExecuteTemplateData.txt")

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(mailMergeFolder + "/" + localDocumentFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "data": localDataFile,
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }

    request := &models.ExecuteMailMergeRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: options,
    }

    actual, err := client.WordsApi.ExecuteMailMerge(ctx, request)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Document, "Validate ExecuteTemplate response.");
    assert.Equal(t, "TestExecuteTemplate.docx", actual.Document.FileName, "Validate ExecuteTemplate response.");
}

// Test for execute template online.
func Test_ExecuteTemplate_ExecuteTemplateOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    mailMergeFolder := "DocumentActions/MailMerge"
    localDocumentFile := "SampleMailMergeTemplate.docx"
    localDataFile := "SampleExecuteTemplateData.txt"


    options := map[string]interface{}{
    }

    request := &models.ExecuteMailMergeOnlineRequest{
        Template: OpenFile(t, mailMergeFolder + "/" + localDocumentFile),
        Data: OpenFile(t, mailMergeFolder + "/" + localDataFile),
        Optionals: options,
    }

    _, err := client.WordsApi.ExecuteMailMergeOnline(ctx, request)

    if err != nil {
        t.Error(err)
    }

}
