/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="footnote_test.go">
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

// Example of how to work with footnotes.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2507/api/models"
)

// Test for adding footnote.
func Test_Footnote_InsertFootnote(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Footnotes"
    footnoteFolder := "DocumentElements/Footnotes"
    remoteFileName := "TestInsertFootnote.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(footnoteFolder + "/Footnote.doc"), remoteDataFolder + "/" + remoteFileName)

    requestFootnoteDto := models.FootnoteInsert{
        FootnoteType: ToStringPointer("Endnote"),
        Text: ToStringPointer("test endnote"),
    }

    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }

    request := &models.InsertFootnoteRequest{
        Name: ToStringPointer(remoteFileName),
        FootnoteDto: &requestFootnoteDto,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.InsertFootnote(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetFootnote(), "Validate InsertFootnote response.");
    assert.Equal(t, "0.1.7.1", DereferenceValue(actual.GetFootnote().GetNodeId()), "Validate InsertFootnote response.");
    assert.Equal(t, " test endnote" + "\r\n", DereferenceValue(actual.GetFootnote().GetText()), "Validate InsertFootnote response.");
}

// Test for adding footnote online.
func Test_Footnote_InsertFootnoteOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    footnoteFolder := "DocumentElements/Footnotes"

    requestDocument := OpenFile(t, footnoteFolder + "/Footnote.doc")
    requestFootnoteDto := models.FootnoteInsert{
        FootnoteType: ToStringPointer("Endnote"),
        Text: ToStringPointer("test endnote"),
    }

    options := map[string]interface{}{
        "nodePath": "",
    }

    request := &models.InsertFootnoteOnlineRequest{
        Document: requestDocument,
        FootnoteDto: &requestFootnoteDto,
        Optionals: options,
    }

    _, _, err := client.WordsApi.InsertFootnoteOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for adding footnote without node path.
func Test_Footnote_InsertFootnoteWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Footnotes"
    footnoteFolder := "DocumentElements/Footnotes"
    remoteFileName := "TestInsertFootnoteWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(footnoteFolder + "/Footnote.doc"), remoteDataFolder + "/" + remoteFileName)

    requestFootnoteDto := models.FootnoteInsert{
        FootnoteType: ToStringPointer("Endnote"),
        Text: ToStringPointer("test endnote"),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.InsertFootnoteRequest{
        Name: ToStringPointer(remoteFileName),
        FootnoteDto: &requestFootnoteDto,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.InsertFootnote(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetFootnote(), "Validate InsertFootnoteWithoutNodePath response.");
    assert.Equal(t, "0.1.7.1", DereferenceValue(actual.GetFootnote().GetNodeId()), "Validate InsertFootnoteWithoutNodePath response.");
    assert.Equal(t, " test endnote" + "\r\n", DereferenceValue(actual.GetFootnote().GetText()), "Validate InsertFootnoteWithoutNodePath response.");
}

// Test for deleting footnote.
func Test_Footnote_DeleteFootnote(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Footnotes"
    footnoteFolder := "DocumentElements/Footnotes"
    remoteFileName := "TestDeleteFootnote.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(footnoteFolder + "/Footnote.doc"), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }

    request := &models.DeleteFootnoteRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, err := client.WordsApi.DeleteFootnote(ctx, request)

    if err != nil {
        t.Error(err)
    }

}

// Test for deleting footnote online.
func Test_Footnote_DeleteFootnoteOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    footnoteFolder := "DocumentElements/Footnotes"

    requestDocument := OpenFile(t, footnoteFolder + "/Footnote.doc")

    options := map[string]interface{}{
        "nodePath": "",
    }

    request := &models.DeleteFootnoteOnlineRequest{
        Document: requestDocument,
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.DeleteFootnoteOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for deleting footnote without node path.
func Test_Footnote_DeleteFootnoteWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Footnotes"
    footnoteFolder := "DocumentElements/Footnotes"
    remoteFileName := "TestDeleteFootnoteWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(footnoteFolder + "/Footnote.doc"), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.DeleteFootnoteRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, err := client.WordsApi.DeleteFootnote(ctx, request)

    if err != nil {
        t.Error(err)
    }

}

// Test for getting footnotes.
func Test_Footnote_GetFootnotes(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Footnotes"
    footnoteFolder := "DocumentElements/Footnotes"
    remoteFileName := "TestGetFootnotes.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(footnoteFolder + "/Footnote.doc"), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }

    request := &models.GetFootnotesRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetFootnotes(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetFootnotes(), "Validate GetFootnotes response.");
    assert.NotNil(t, actual.GetFootnotes().GetList(), "Validate GetFootnotes response.");
    assert.Equal(t, 6, len(actual.GetFootnotes().GetList()), "Validate GetFootnotes response.");
    assert.Equal(t, " Footnote 1." + "\r\n", DereferenceValue(actual.GetFootnotes().GetList()[0].GetText()), "Validate GetFootnotes response.");
}

// Test for getting footnotes online.
func Test_Footnote_GetFootnotesOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    footnoteFolder := "DocumentElements/Footnotes"

    requestDocument := OpenFile(t, footnoteFolder + "/Footnote.doc")

    options := map[string]interface{}{
        "nodePath": "",
    }

    request := &models.GetFootnotesOnlineRequest{
        Document: requestDocument,
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetFootnotesOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for getting footnotes without node path.
func Test_Footnote_GetFootnotesWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Footnotes"
    footnoteFolder := "DocumentElements/Footnotes"
    remoteFileName := "TestGetFootnotesWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(footnoteFolder + "/Footnote.doc"), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.GetFootnotesRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetFootnotes(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetFootnotes(), "Validate GetFootnotesWithoutNodePath response.");
    assert.NotNil(t, actual.GetFootnotes().GetList(), "Validate GetFootnotesWithoutNodePath response.");
    assert.Equal(t, 6, len(actual.GetFootnotes().GetList()), "Validate GetFootnotesWithoutNodePath response.");
    assert.Equal(t, " Footnote 1." + "\r\n", DereferenceValue(actual.GetFootnotes().GetList()[0].GetText()), "Validate GetFootnotesWithoutNodePath response.");
}

// Test for getting footnote.
func Test_Footnote_GetFootnote(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Footnotes"
    footnoteFolder := "DocumentElements/Footnotes"
    remoteFileName := "TestGetFootnote.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(footnoteFolder + "/Footnote.doc"), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }

    request := &models.GetFootnoteRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetFootnote(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetFootnote(), "Validate GetFootnote response.");
    assert.Equal(t, " Footnote 1." + "\r\n", DereferenceValue(actual.GetFootnote().GetText()), "Validate GetFootnote response.");
}

// Test for getting footnote online.
func Test_Footnote_GetFootnoteOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    footnoteFolder := "DocumentElements/Footnotes"

    requestDocument := OpenFile(t, footnoteFolder + "/Footnote.doc")

    options := map[string]interface{}{
        "nodePath": "",
    }

    request := &models.GetFootnoteOnlineRequest{
        Document: requestDocument,
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetFootnoteOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for getting footnote without node path.
func Test_Footnote_GetFootnoteWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Footnotes"
    footnoteFolder := "DocumentElements/Footnotes"
    remoteFileName := "TestGetFootnoteWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(footnoteFolder + "/Footnote.doc"), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.GetFootnoteRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetFootnote(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetFootnote(), "Validate GetFootnoteWithoutNodePath response.");
    assert.Equal(t, " Footnote 1." + "\r\n", DereferenceValue(actual.GetFootnote().GetText()), "Validate GetFootnoteWithoutNodePath response.");
}

// Test for updating footnote.
func Test_Footnote_UpdateFootnote(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Footnotes"
    footnoteFolder := "DocumentElements/Footnotes"
    remoteFileName := "TestUpdateFootnote.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(footnoteFolder + "/Footnote.doc"), remoteDataFolder + "/" + remoteFileName)

    requestFootnoteDto := models.FootnoteUpdate{
        Text: ToStringPointer("new text is here"),
    }

    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }

    request := &models.UpdateFootnoteRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        FootnoteDto: &requestFootnoteDto,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.UpdateFootnote(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetFootnote(), "Validate UpdateFootnote response.");
    assert.Equal(t, " new text is here" + "\r\n", DereferenceValue(actual.GetFootnote().GetText()), "Validate UpdateFootnote response.");
}

// Test for updating footnote online.
func Test_Footnote_UpdateFootnoteOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    footnoteFolder := "DocumentElements/Footnotes"

    requestDocument := OpenFile(t, footnoteFolder + "/Footnote.doc")
    requestFootnoteDto := models.FootnoteUpdate{
        Text: ToStringPointer("new text is here"),
    }

    options := map[string]interface{}{
        "nodePath": "",
    }

    request := &models.UpdateFootnoteOnlineRequest{
        Document: requestDocument,
        FootnoteDto: &requestFootnoteDto,
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.UpdateFootnoteOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for updating footnote without node path.
func Test_Footnote_UpdateFootnoteWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Footnotes"
    footnoteFolder := "DocumentElements/Footnotes"
    remoteFileName := "TestUpdateFootnoteWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(footnoteFolder + "/Footnote.doc"), remoteDataFolder + "/" + remoteFileName)

    requestFootnoteDto := models.FootnoteUpdate{
        Text: ToStringPointer("new text is here"),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.UpdateFootnoteRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        FootnoteDto: &requestFootnoteDto,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.UpdateFootnote(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetFootnote(), "Validate UpdateFootnoteWithoutNodePath response.");
    assert.Equal(t, " new text is here" + "\r\n", DereferenceValue(actual.GetFootnote().GetText()), "Validate UpdateFootnoteWithoutNodePath response.");
}
