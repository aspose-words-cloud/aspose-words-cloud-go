/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="range_test.go">
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

// Example of how to work with ranges.
package api_test

import (
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2007/api/models"
)

// Test for getting the text from range.
func Test_Range_GetRangeText(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Range"
    localFile := "DocumentElements/Range/RangeGet.doc"
    remoteFileName := "TestGetRangeText.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "rangeEndIdentifier": "id0.0.1",
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetRangeText(ctx, remoteFileName, "id0.0.0", options)

    if err != nil {
        t.Error(err)
    }
}

// Test for removing the text for range.
func Test_Range_RemoveRange(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Range"
    localFile := "DocumentElements/Range/RangeGet.doc"
    remoteFileName := "TestRemoveRange.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "rangeEndIdentifier": "id0.0.1",
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.RemoveRange(ctx, remoteFileName, "id0.0.0", options)

    if err != nil {
        t.Error(err)
    }
}

// Test for saving a range as a new document.
func Test_Range_SaveAsRange(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Range"
    localFile := "DocumentElements/Range/RangeGet.doc"
    remoteFileName := "TestSaveAsRange.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestDocumentParameters := models.RangeDocument{
        DocumentName: remoteDataFolder + "/NewDoc.docx",
    }

    options := map[string]interface{}{
        "rangeEndIdentifier": "id0.0.1",
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.SaveAsRange(ctx, remoteFileName, "id0.0.0", requestDocumentParameters, options)

    if err != nil {
        t.Error(err)
    }
}

// Test for replacing text in range.
func Test_Range_ReplaceWithText(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Range"
    localFile := "DocumentElements/Range/RangeGet.doc"
    remoteFileName := "TestReplaceWithText.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestRangeText := models.ReplaceRange{
        Text: "Replaced header",
    }

    options := map[string]interface{}{
        "rangeEndIdentifier": "id0.0.1",
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.ReplaceWithText(ctx, remoteFileName, "id0.0.0", requestRangeText, options)

    if err != nil {
        t.Error(err)
    }
}
