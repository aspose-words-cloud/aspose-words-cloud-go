/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="classification_test.go">
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

// Example of how to classify text.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2107/api/models"
)

// Test for raw text classification.
func Test_Classification_Classify(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)

    options := map[string]interface{}{
        "bestClassesCount": "3",
    }

    request := &models.ClassifyRequest{
        Text: ToStringPointer("Try text classification"),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.Classify(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.Equal(t, "Science", actual.BestClassName, "Validate Classify response.");
    assert.NotNil(t, actual.BestResults, "Validate Classify response.");
    assert.Equal(t, 3, len(actual.BestResults), "Validate Classify response.");
}

// Test for document classification.
func Test_Classification_ClassifyDocument(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/Common"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestClassifyDocument.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "bestClassesCount": "3",
    }

    request := &models.ClassifyDocumentRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.ClassifyDocument(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.Equal(t, "Hobbies_&_Interests", actual.BestClassName, "Validate ClassifyDocument response.");
    assert.NotNil(t, actual.BestResults, "Validate ClassifyDocument response.");
    assert.Equal(t, 3, len(actual.BestResults), "Validate ClassifyDocument response.");
}

// Test for document classification online.
func Test_Classification_ClassifyDocumentOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"


    options := map[string]interface{}{
        "bestClassesCount": "3",
    }

    request := &models.ClassifyDocumentOnlineRequest{
        Document: OpenFile(t, localFile),
        Optionals: options,
    }

    _, _, err := client.WordsApi.ClassifyDocumentOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}
