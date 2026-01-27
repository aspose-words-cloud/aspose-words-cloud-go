/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="text_test.go">
 *   Copyright (c) 2026 Aspose.Words for Cloud
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

// Example of how to work with macros.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2601/api/models"
)

// Test for replacing text.
func Test_Text_ReplaceText(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Text"
    remoteFileName := "TestReplaceText.docx"
    localFile := "Common/test_multi_pages.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestReplaceText := models.ReplaceTextParameters{
        OldValue: ToStringPointer("Testing"),
        NewValue: ToStringPointer("Aspose testing"),
        IsMatchCase: ToBoolPointer(true),
        IsMatchWholeWord: ToBoolPointer(false),
        IsOldValueRegex: ToBoolPointer(false),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }

    request := &models.ReplaceTextRequest{
        Name: ToStringPointer(remoteFileName),
        ReplaceText: &requestReplaceText,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.ReplaceText(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.Equal(t, int32(3), DereferenceValue(actual.GetMatches()), "Validate ReplaceText response.");
}

// Test for replacing text online.
func Test_Text_ReplaceTextOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)
    requestReplaceText := models.ReplaceTextParameters{
        OldValue: ToStringPointer("aspose"),
        NewValue: ToStringPointer("aspose new"),
        IsMatchCase: ToBoolPointer(true),
        IsMatchWholeWord: ToBoolPointer(false),
        IsOldValueRegex: ToBoolPointer(false),
    }

    options := map[string]interface{}{
    }

    request := &models.ReplaceTextOnlineRequest{
        Document: requestDocument,
        ReplaceText: &requestReplaceText,
        Optionals: options,
    }

    _, _, err := client.WordsApi.ReplaceTextOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for searching.
func Test_Text_Search(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Text"
    remoteFileName := "TestSearch.docx"
    localFile := "DocumentElements/Text/SampleWordDocument.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.SearchRequest{
        Name: ToStringPointer(remoteFileName),
        Pattern: ToStringPointer("aspose"),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.Search(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetSearchResults(), "Validate Search response.");
    assert.NotNil(t, actual.GetSearchResults().GetResultsList(), "Validate Search response.");
    assert.Equal(t, 23, len(actual.GetSearchResults().GetResultsList()), "Validate Search response.");
    assert.NotNil(t, actual.GetSearchResults().GetResultsList()[0].GetRangeStart(), "Validate Search response.");
    assert.Equal(t, int32(65), DereferenceValue(actual.GetSearchResults().GetResultsList()[0].GetRangeStart().GetOffset()), "Validate Search response.");
}

// Test for searching online.
func Test_Text_SearchOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Text/SampleWordDocument.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.SearchOnlineRequest{
        Document: requestDocument,
        Pattern: ToStringPointer("aspose"),
        Optionals: options,
    }

    _, _, err := client.WordsApi.SearchOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}
