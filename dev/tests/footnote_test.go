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
    "github.com/aspose-words-cloud/aspose-words-cloud-go/dev/api/models"
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
        FootnoteType: "Endnote",
        Text: "test endnote",
    }

    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.InsertFootnote(ctx, remoteFileName, requestFootnoteDto, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Footnote, "Validate InsertFootnote response.");
    assert.True(t, strings.HasPrefix(actual.Footnote.NodeId, "0.1.7.1"), "Validate InsertFootnote response.");
    assert.True(t, strings.HasPrefix(actual.Footnote.Text, " test endnote"), "Validate InsertFootnote response.");
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
        FootnoteType: "Endnote",
        Text: "test endnote",
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.InsertFootnote(ctx, remoteFileName, requestFootnoteDto, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Footnote, "Validate InsertFootnoteWithoutNodePath response.");
    assert.True(t, strings.HasPrefix(actual.Footnote.NodeId, "0.1.7.1"), "Validate InsertFootnoteWithoutNodePath response.");
    assert.True(t, strings.HasPrefix(actual.Footnote.Text, " test endnote"), "Validate InsertFootnoteWithoutNodePath response.");
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
    _, err := client.WordsApi.DeleteFootnote(ctx, remoteFileName, int32(0), options)

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
    _, err := client.WordsApi.DeleteFootnote(ctx, remoteFileName, int32(0), options)

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
    actual, _, err := client.WordsApi.GetFootnotes(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Footnotes, "Validate GetFootnotes response.");
    assert.NotNil(t, actual.Footnotes.List, "Validate GetFootnotes response.");
    assert.Equal(t, 6, len(actual.Footnotes.List), "Validate GetFootnotes response.");
    assert.True(t, strings.HasPrefix(actual.Footnotes.List[0].Text, " Footnote 1."), "Validate GetFootnotes response.");
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
    actual, _, err := client.WordsApi.GetFootnotes(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Footnotes, "Validate GetFootnotesWithoutNodePath response.");
    assert.NotNil(t, actual.Footnotes.List, "Validate GetFootnotesWithoutNodePath response.");
    assert.Equal(t, 6, len(actual.Footnotes.List), "Validate GetFootnotesWithoutNodePath response.");
    assert.True(t, strings.HasPrefix(actual.Footnotes.List[0].Text, " Footnote 1."), "Validate GetFootnotesWithoutNodePath response.");
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
    actual, _, err := client.WordsApi.GetFootnote(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Footnote, "Validate GetFootnote response.");
    assert.True(t, strings.HasPrefix(actual.Footnote.Text, " Footnote 1."), "Validate GetFootnote response.");
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
    actual, _, err := client.WordsApi.GetFootnote(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Footnote, "Validate GetFootnoteWithoutNodePath response.");
    assert.True(t, strings.HasPrefix(actual.Footnote.Text, " Footnote 1."), "Validate GetFootnoteWithoutNodePath response.");
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
        Text: "new text is here",
    }

    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.UpdateFootnote(ctx, remoteFileName, requestFootnoteDto, int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Footnote, "Validate UpdateFootnote response.");
    assert.True(t, strings.HasPrefix(actual.Footnote.Text, " new text is here"), "Validate UpdateFootnote response.");
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
        Text: "new text is here",
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.UpdateFootnote(ctx, remoteFileName, requestFootnoteDto, int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Footnote, "Validate UpdateFootnoteWithoutNodePath response.");
    assert.True(t, strings.HasPrefix(actual.Footnote.Text, " new text is here"), "Validate UpdateFootnoteWithoutNodePath response.");
}
