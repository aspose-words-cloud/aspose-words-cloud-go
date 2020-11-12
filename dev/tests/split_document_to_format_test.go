/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="split_document_to_format_test.go">
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

// Example of how to split document and return result with specified format and page range.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

// Test for document splitting.
func Test_SplitDocumentToFormat_SplitDocument(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentActions/SplitDocument"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestSplitDocument.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/TestSplitDocument.text",
        "from": int32(1),
        "to": int32(2),
    }
    actual, _, err := client.WordsApi.SplitDocument(ctx, remoteFileName, "text", options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.SplitResult, "Validate SplitDocument response.");
    assert.NotNil(t, actual.SplitResult.Pages, "Validate SplitDocument response.");
    assert.Equal(t, 2, len(actual.SplitResult.Pages), "Validate SplitDocument response.");
}
