/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="range_test.go">
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

// Example of how to work with ranges.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2510/api/models"
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

    request := &models.GetRangeTextRequest{
        Name: ToStringPointer(remoteFileName),
        RangeStartIdentifier: ToStringPointer("id0.0.0"),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetRangeText(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.Equal(t, "This is HEADER ", DereferenceValue(actual.GetText()), "Validate GetRangeText response.");
}

// Test for getting the text from range online.
func Test_Range_GetRangeTextOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Range/RangeGet.doc"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
        "rangeEndIdentifier": "id0.0.1",
    }

    request := &models.GetRangeTextOnlineRequest{
        Document: requestDocument,
        RangeStartIdentifier: ToStringPointer("id0.0.0"),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetRangeTextOnline(ctx, request)
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

    request := &models.RemoveRangeRequest{
        Name: ToStringPointer(remoteFileName),
        RangeStartIdentifier: ToStringPointer("id0.0.0"),
        Optionals: options,
    }

    _, _, err := client.WordsApi.RemoveRange(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for removing the text for range online.
func Test_Range_RemoveRangeOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Range/RangeGet.doc"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
        "rangeEndIdentifier": "id0.0.1",
    }

    request := &models.RemoveRangeOnlineRequest{
        Document: requestDocument,
        RangeStartIdentifier: ToStringPointer("id0.0.0"),
        Optionals: options,
    }

    _, _, err := client.WordsApi.RemoveRangeOnline(ctx, request)
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
        DocumentName: ToStringPointer(remoteDataFolder + "/NewDoc.docx"),
    }

    options := map[string]interface{}{
        "rangeEndIdentifier": "id0.0.1",
        "folder": remoteDataFolder,
    }

    request := &models.SaveAsRangeRequest{
        Name: ToStringPointer(remoteFileName),
        RangeStartIdentifier: ToStringPointer("id0.0.0"),
        DocumentParameters: &requestDocumentParameters,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.SaveAsRange(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetDocument(), "Validate SaveAsRange response.");
    assert.Equal(t, "NewDoc.docx", DereferenceValue(actual.GetDocument().GetFileName()), "Validate SaveAsRange response.");
}

// Test for saving a range as a new document online.
func Test_Range_SaveAsRangeOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Range"
    localFile := "DocumentElements/Range/RangeGet.doc"

    requestDocument := OpenFile(t, localFile)
    requestDocumentParameters := models.RangeDocument{
        DocumentName: ToStringPointer(remoteDataFolder + "/NewDoc.docx"),
    }

    options := map[string]interface{}{
        "rangeEndIdentifier": "id0.0.1",
    }

    request := &models.SaveAsRangeOnlineRequest{
        Document: requestDocument,
        RangeStartIdentifier: ToStringPointer("id0.0.0"),
        DocumentParameters: &requestDocumentParameters,
        Optionals: options,
    }

    _, _, err := client.WordsApi.SaveAsRangeOnline(ctx, request)
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
        Text: ToStringPointer("Replaced header"),
    }

    options := map[string]interface{}{
        "rangeEndIdentifier": "id0.0.1",
        "folder": remoteDataFolder,
    }

    request := &models.ReplaceWithTextRequest{
        Name: ToStringPointer(remoteFileName),
        RangeStartIdentifier: ToStringPointer("id0.0.0"),
        RangeText: &requestRangeText,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.ReplaceWithText(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetDocument(), "Validate ReplaceWithText response.");
    assert.Equal(t, "TestReplaceWithText.docx", DereferenceValue(actual.GetDocument().GetFileName()), "Validate ReplaceWithText response.");
}

// Test for replacing text in range online.
func Test_Range_ReplaceWithTextOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Range/RangeGet.doc"

    requestDocument := OpenFile(t, localFile)
    requestRangeText := models.ReplaceRange{
        Text: ToStringPointer("Replaced header"),
    }

    options := map[string]interface{}{
        "rangeEndIdentifier": "id0.0.1",
    }

    request := &models.ReplaceWithTextOnlineRequest{
        Document: requestDocument,
        RangeStartIdentifier: ToStringPointer("id0.0.0"),
        RangeText: &requestRangeText,
        Optionals: options,
    }

    _, _, err := client.WordsApi.ReplaceWithTextOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test to translate node id to node path.
func Test_Range_TranslateNodeId(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Range"
    localFile := "DocumentElements/Range/RangeGet.doc"
    remoteFileName := "TestTranslateNodeId.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.TranslateNodeIdRequest{
        Name: ToStringPointer(remoteFileName),
        NodeId: ToStringPointer("id0.0.0"),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.TranslateNodeId(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.Equal(t, "sections/0/body/paragraphs/0", DereferenceValue(actual.GetPath()), "Validate TranslateNodeId response.");
}

// Test to translate node id to node path online.
func Test_Range_TranslateNodeIdOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Range/RangeGet.doc"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.TranslateNodeIdOnlineRequest{
        Document: requestDocument,
        NodeId: ToStringPointer("id0.0.0"),
        Optionals: options,
    }

    _, _, err := client.WordsApi.TranslateNodeIdOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}
