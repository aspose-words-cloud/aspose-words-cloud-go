/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="compare_document_test.go">
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

// Example of document comparison.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2011/api/models"
)

// Test for document comparison.
func Test_CompareDocument_CompareDocument(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteFolder := remoteBaseTestDataFolder + "/DocumentActions/CompareDocument"
    localFolder := "DocumentActions/CompareDocument"
    localName1 := "compareTestDoc1.doc"
    localName2 := "compareTestDoc2.doc"
    remoteName1 := "TestCompareDocument1.doc"
    remoteName2 := "TestCompareDocument2.doc"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFolder + "/" + localName1), remoteFolder + "/" + remoteName1)
    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFolder + "/" + localName2), remoteFolder + "/" + remoteName2)

    requestCompareData := models.CompareData{
        Author: ToStringPointer("author"),
        ComparingWithDocument: ToStringPointer(remoteFolder + "/" + remoteName2),
        DateTime: CreateTime(2015, 10, 26, 0, 0, 0),
    }

    options := map[string]interface{}{
        "folder": remoteFolder,
        "destFileName": baseTestOutPath + "/TestCompareDocumentOut.doc",
    }
    actual, _, err := client.WordsApi.CompareDocument(ctx, remoteName1, requestCompareData, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Document, "Validate CompareDocument response.");
    assert.Equal(t, "TestCompareDocumentOut.doc", *actual.Document.FileName, "Validate CompareDocument response.");
}
