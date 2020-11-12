/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="hyperlink_test.go">
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

// Example of how to work with hyperlinks.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

// Test for getting hyperlink by specified index.
func Test_Hyperlink_GetDocumentHyperlinkByIndex(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Hyperlink"
    localFile := "Common/test_doc.docx"
    remoteFileName := "TestGetDocumentHyperlinkByIndex.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.GetDocumentHyperlinkByIndex(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Hyperlink, "Validate GetDocumentHyperlinkByIndex response.");
    assert.Equal(t, "Aspose", *actual.Hyperlink.DisplayText, "Validate GetDocumentHyperlinkByIndex response.");
}

// Test for getting hyperlinks.
func Test_Hyperlink_GetDocumentHyperlinks(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Hyperlink"
    localFile := "Common/test_doc.docx"
    remoteFileName := "TestGetDocumentHyperlinks.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.GetDocumentHyperlinks(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Hyperlinks, "Validate GetDocumentHyperlinks response.");
    assert.NotNil(t, actual.Hyperlinks.HyperlinkList, "Validate GetDocumentHyperlinks response.");
    assert.Equal(t, 2, len(actual.Hyperlinks.HyperlinkList), "Validate GetDocumentHyperlinks response.");
    assert.Equal(t, "Aspose", *actual.Hyperlinks.HyperlinkList[0].DisplayText, "Validate GetDocumentHyperlinks response.");
}
