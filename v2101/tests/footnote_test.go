/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="footnote_test.go">
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

// Example of how to work with footnotes.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2101/api/models"
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
        FootnoteDto: requestFootnoteDto,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.InsertFootnote(ctx, request)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Footnote, "Validate InsertFootnote response.");
    assert.Equal(t, "0.1.7.1", actual.Footnote.NodeId, "Validate InsertFootnote response.");
    assert.Equal(t, " test endnote" + "\r\n", actual.Footnote.Text, "Validate InsertFootnote response.");
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
        FootnoteDto: requestFootnoteDto,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.InsertFootnote(ctx, request)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Footnote, "Validate InsertFootnoteWithoutNodePath response.");
    assert.Equal(t, "0.1.7.1", actual.Footnote.NodeId, "Validate InsertFootnoteWithoutNodePath response.");
    assert.Equal(t, " test endnote" + "\r\n", actual.Footnote.Text, "Validate InsertFootnoteWithoutNodePath response.");
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

    assert.NotNil(t, actual.Footnotes, "Validate GetFootnotes response.");
    assert.NotNil(t, actual.Footnotes.List, "Validate GetFootnotes response.");
    assert.Equal(t, 6, len(actual.Footnotes.List), "Validate GetFootnotes response.");
    assert.Equal(t, " Footnote 1." + "\r\n", actual.Footnotes.List[0].Text, "Validate GetFootnotes response.");
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

    assert.NotNil(t, actual.Footnotes, "Validate GetFootnotesWithoutNodePath response.");
    assert.NotNil(t, actual.Footnotes.List, "Validate GetFootnotesWithoutNodePath response.");
    assert.Equal(t, 6, len(actual.Footnotes.List), "Validate GetFootnotesWithoutNodePath response.");
    assert.Equal(t, " Footnote 1." + "\r\n", actual.Footnotes.List[0].Text, "Validate GetFootnotesWithoutNodePath response.");
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

    assert.NotNil(t, actual.Footnote, "Validate GetFootnote response.");
    assert.Equal(t, " Footnote 1." + "\r\n", actual.Footnote.Text, "Validate GetFootnote response.");
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

    assert.NotNil(t, actual.Footnote, "Validate GetFootnoteWithoutNodePath response.");
    assert.Equal(t, " Footnote 1." + "\r\n", actual.Footnote.Text, "Validate GetFootnoteWithoutNodePath response.");
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
        FootnoteDto: requestFootnoteDto,
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.UpdateFootnote(ctx, request)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Footnote, "Validate UpdateFootnote response.");
    assert.Equal(t, " new text is here" + "\r\n", actual.Footnote.Text, "Validate UpdateFootnote response.");
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
        FootnoteDto: requestFootnoteDto,
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.UpdateFootnote(ctx, request)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Footnote, "Validate UpdateFootnoteWithoutNodePath response.");
    assert.Equal(t, " new text is here" + "\r\n", actual.Footnote.Text, "Validate UpdateFootnoteWithoutNodePath response.");
}
