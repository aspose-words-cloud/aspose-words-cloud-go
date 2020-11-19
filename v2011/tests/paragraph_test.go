/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="paragraph_test.go">
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

// Example of how to work with paragraph.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2011/api/models"
)

// Test for getting paragraph.
func Test_Paragraph_GetDocumentParagraphByIndex(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetDocumentParagraphByIndex.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "sections/0",
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.GetParagraph(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Paragraph, "Validate GetDocumentParagraphByIndex response.");
    assert.Equal(t, "0.0.0", *actual.Paragraph.NodeId, "Validate GetDocumentParagraphByIndex response.");
}

// Test for getting paragraph without node path.
func Test_Paragraph_GetDocumentParagraphByIndexWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetDocumentParagraphByIndexWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.GetParagraph(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Paragraph, "Validate GetDocumentParagraphByIndexWithoutNodePath response.");
    assert.Equal(t, "0.0.0", *actual.Paragraph.NodeId, "Validate GetDocumentParagraphByIndexWithoutNodePath response.");
}

// Test for getting all paragraphs.
func Test_Paragraph_GetDocumentParagraphs(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetDocumentParagraphs.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "sections/0",
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.GetParagraphs(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Paragraphs, "Validate GetDocumentParagraphs response.");
    assert.NotNil(t, actual.Paragraphs.ParagraphLinkList, "Validate GetDocumentParagraphs response.");
    assert.Equal(t, 15, len(actual.Paragraphs.ParagraphLinkList), "Validate GetDocumentParagraphs response.");
    assert.Equal(t, "Page 1 of 3", *actual.Paragraphs.ParagraphLinkList[0].Text, "Validate GetDocumentParagraphs response.");
}

// Test for getting all paragraphs without node path.
func Test_Paragraph_GetDocumentParagraphsWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetDocumentParagraphsWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.GetParagraphs(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Paragraphs, "Validate GetDocumentParagraphsWithoutNodePath response.");
    assert.NotNil(t, actual.Paragraphs.ParagraphLinkList, "Validate GetDocumentParagraphsWithoutNodePath response.");
    assert.Equal(t, 15, len(actual.Paragraphs.ParagraphLinkList), "Validate GetDocumentParagraphsWithoutNodePath response.");
    assert.Equal(t, "Page 1 of 3", *actual.Paragraphs.ParagraphLinkList[0].Text, "Validate GetDocumentParagraphsWithoutNodePath response.");
}

// Test for getting paragraph run.
func Test_Paragraph_GetDocumentParagraphRun(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetDocumentParagraphRun.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.GetRun(ctx, remoteFileName, "paragraphs/0", int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Run, "Validate GetDocumentParagraphRun response.");
    assert.Equal(t, "Page ", *actual.Run.Text, "Validate GetDocumentParagraphRun response.");
}

// Test for getting paragraph run font.
func Test_Paragraph_GetDocumentParagraphRunFont(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetDocumentParagraphRunFont.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.GetRunFont(ctx, remoteFileName, "paragraphs/0", int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Font, "Validate GetDocumentParagraphRunFont response.");
    assert.Equal(t, "Times New Roman", *actual.Font.Name, "Validate GetDocumentParagraphRunFont response.");
}

// Test for getting paragraph runs.
func Test_Paragraph_GetParagraphRuns(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetParagraphRuns.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.GetRuns(ctx, remoteFileName, "sections/0/paragraphs/0", options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Runs, "Validate GetParagraphRuns response.");
    assert.NotNil(t, actual.Runs.List, "Validate GetParagraphRuns response.");
    assert.Equal(t, 6, len(actual.Runs.List), "Validate GetParagraphRuns response.");
    assert.Equal(t, "Page ", *actual.Runs.List[0].Text, "Validate GetParagraphRuns response.");
}

// Test for updating paragraph run font.
func Test_Paragraph_UpdateRunFont(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestUpdateRunFont.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestFontDto := models.Font{
        Bold: ToBoolPointer(true),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }
    actual, _, err := client.WordsApi.UpdateRunFont(ctx, remoteFileName, requestFontDto, "paragraphs/0", int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Font, "Validate UpdateRunFont response.");
    assert.True(t, *actual.Font.Bold, "Validate UpdateRunFont response.");
}

// Test for adding paragraph.
func Test_Paragraph_InsertParagraph(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestInsertParagraph.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestParagraph := models.ParagraphInsert{
        Text: ToStringPointer("This is a new paragraph for your document"),
    }

    options := map[string]interface{}{
        "nodePath": "sections/0",
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.InsertParagraph(ctx, remoteFileName, requestParagraph, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Paragraph, "Validate InsertParagraph response.");
    assert.Equal(t, "0.3.8", *actual.Paragraph.NodeId, "Validate InsertParagraph response.");
}

// Test for adding paragraph without node path.
func Test_Paragraph_InsertParagraphWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestInsertParagraphWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestParagraph := models.ParagraphInsert{
        Text: ToStringPointer("This is a new paragraph for your document"),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.InsertParagraph(ctx, remoteFileName, requestParagraph, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Paragraph, "Validate InsertParagraphWithoutNodePath response.");
    assert.Equal(t, "0.3.8", *actual.Paragraph.NodeId, "Validate InsertParagraphWithoutNodePath response.");
}

// Test for paragraph rendering.
func Test_Paragraph_RenderParagraph(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestRenderParagraph.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }
    _, err := client.WordsApi.RenderParagraph(ctx, remoteFileName, "png", int32(0), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for paragraph rendering without node path.
func Test_Paragraph_RenderParagraphWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestRenderParagraphWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, err := client.WordsApi.RenderParagraph(ctx, remoteFileName, "png", int32(0), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for getting paragraph format settings.
func Test_Paragraph_GetParagraphFormat(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetDocumentParagraphs.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.GetParagraphFormat(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.ParagraphFormat, "Validate GetParagraphFormat response.");
    assert.Equal(t, "Normal", *actual.ParagraphFormat.StyleName, "Validate GetParagraphFormat response.");
}

// Test for getting paragraph format settings without node path.
func Test_Paragraph_GetParagraphFormatWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetDocumentParagraphsWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.GetParagraphFormat(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.ParagraphFormat, "Validate GetParagraphFormatWithoutNodePath response.");
    assert.Equal(t, "Normal", *actual.ParagraphFormat.StyleName, "Validate GetParagraphFormatWithoutNodePath response.");
}

// Test for updating  paragraph format settings.
func Test_Paragraph_UpdateParagraphFormat(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetDocumentParagraphs.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestDto := models.ParagraphFormatUpdate{
        Alignment: "Right",
    }

    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.UpdateParagraphFormat(ctx, remoteFileName, requestDto, int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.ParagraphFormat, "Validate UpdateParagraphFormat response.");

}

// Test for deleting  a paragraph.
func Test_Paragraph_DeleteParagraph(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestDeleteParagraph.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }
    _, err := client.WordsApi.DeleteParagraph(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for deleting  a paragraph without node path.
func Test_Paragraph_DeleteParagraphWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestDeleteParagraphWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, err := client.WordsApi.DeleteParagraph(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for getting paragraph list format.
func Test_Paragraph_GetParagraphListFormat(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    listFolder := "DocumentElements/ParagraphListFormat"
    remoteFileName := "TestParagraphGetListFormat.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(listFolder + "/ParagraphGetListFormat.doc"), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.GetParagraphListFormat(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.ListFormat, "Validate GetParagraphListFormat response.");
    assert.Equal(t, int32(1), *actual.ListFormat.ListId, "Validate GetParagraphListFormat response.");
}

// Test for getting paragraph list format without node path.
func Test_Paragraph_GetParagraphListFormatWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    listFolder := "DocumentElements/ParagraphListFormat"
    remoteFileName := "TestParagraphGetListFormatWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(listFolder + "/ParagraphGetListFormat.doc"), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.GetParagraphListFormat(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.ListFormat, "Validate GetParagraphListFormatWithoutNodePath response.");
    assert.Equal(t, int32(1), *actual.ListFormat.ListId, "Validate GetParagraphListFormatWithoutNodePath response.");
}

// Test for updating paragraph list format.
func Test_Paragraph_UpdateParagraphListFormat(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    listFolder := "DocumentElements/ParagraphListFormat"
    remoteFileName := "TestUpdateParagraphListFormat.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(listFolder + "/ParagraphUpdateListFormat.doc"), remoteDataFolder + "/" + remoteFileName)

    requestDto := models.ListFormatUpdate{
        ListId: ToIn32Pointer(int32(2)),
    }

    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.UpdateParagraphListFormat(ctx, remoteFileName, requestDto, int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.ListFormat, "Validate UpdateParagraphListFormat response.");
    assert.Equal(t, int32(2), *actual.ListFormat.ListId, "Validate UpdateParagraphListFormat response.");
}

// Test for updating paragraph list format without node path.
func Test_Paragraph_UpdateParagraphListFormatWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    listFolder := "DocumentElements/ParagraphListFormat"
    remoteFileName := "TestUpdateParagraphListFormatWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(listFolder + "/ParagraphUpdateListFormat.doc"), remoteDataFolder + "/" + remoteFileName)

    requestDto := models.ListFormatUpdate{
        ListId: ToIn32Pointer(int32(2)),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.UpdateParagraphListFormat(ctx, remoteFileName, requestDto, int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.ListFormat, "Validate UpdateParagraphListFormatWithoutNodePath response.");
    assert.Equal(t, int32(2), *actual.ListFormat.ListId, "Validate UpdateParagraphListFormatWithoutNodePath response.");
}

// Test for deleting paragraph list format.
func Test_Paragraph_DeleteParagraphListFormat(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    listFolder := "DocumentElements/ParagraphListFormat"
    remoteFileName := "TestDeleteParagraphListFormat.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(listFolder + "/ParagraphDeleteListFormat.doc"), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.DeleteParagraphListFormat(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for deleting paragraph list format without node path.
func Test_Paragraph_DeleteParagraphListFormatWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    listFolder := "DocumentElements/ParagraphListFormat"
    remoteFileName := "TestDeleteParagraphListFormatWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(listFolder + "/ParagraphDeleteListFormat.doc"), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.DeleteParagraphListFormat(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for getting paragraph tab stops.
func Test_Paragraph_GetParagraphTabStops(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    tabStopFolder := "DocumentElements/Paragraphs"
    remoteFileName := "TestGetParagraphTabStops.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(tabStopFolder + "/ParagraphTabStops.docx"), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.GetParagraphTabStops(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.TabStops, "Validate GetParagraphTabStops response.");
    assert.Equal(t, 2, len(actual.TabStops), "Validate GetParagraphTabStops response.");
    assert.Equal(t, 72.0, *actual.TabStops[0].Position, "Validate GetParagraphTabStops response.");
}

// Test for getting paragraph tab stops without node path.
func Test_Paragraph_GetParagraphTabStopsWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    tabStopFolder := "DocumentElements/Paragraphs"
    remoteFileName := "TestGetParagraphTabStopsWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(tabStopFolder + "/ParagraphTabStops.docx"), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.GetParagraphTabStops(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.TabStops, "Validate GetParagraphTabStopsWithoutNodePath response.");
    assert.Equal(t, 2, len(actual.TabStops), "Validate GetParagraphTabStopsWithoutNodePath response.");
    assert.Equal(t, 72.0, *actual.TabStops[0].Position, "Validate GetParagraphTabStopsWithoutNodePath response.");
}

// Test for inserting paragraph tab stop.
func Test_Paragraph_InsertParagraphTabStops(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    tabStopFolder := "DocumentElements/Paragraphs"
    remoteFileName := "TestInsertOrUpdateParagraphTabStop.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(tabStopFolder + "/ParagraphTabStops.docx"), remoteDataFolder + "/" + remoteFileName)

    requestDto := models.TabStopInsert{
        Alignment: "Left",
        Leader: "None",
        Position: ToFloat64Pointer(100.0),
    }

    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.InsertOrUpdateParagraphTabStop(ctx, remoteFileName, requestDto, int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.TabStops, "Validate InsertParagraphTabStops response.");
    assert.Equal(t, 3, len(actual.TabStops), "Validate InsertParagraphTabStops response.");
    assert.Equal(t, 100.0, *actual.TabStops[1].Position, "Validate InsertParagraphTabStops response.");


}

// Test for inserting paragraph tab stop without node path.
func Test_Paragraph_InsertParagraphTabStopsWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    tabStopFolder := "DocumentElements/Paragraphs"
    remoteFileName := "TestInsertOrUpdateParagraphTabStopWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(tabStopFolder + "/ParagraphTabStops.docx"), remoteDataFolder + "/" + remoteFileName)

    requestDto := models.TabStopInsert{
        Alignment: "Left",
        Leader: "None",
        Position: ToFloat64Pointer(100.0),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.InsertOrUpdateParagraphTabStop(ctx, remoteFileName, requestDto, int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.TabStops, "Validate InsertParagraphTabStopsWithoutNodePath response.");
    assert.Equal(t, 3, len(actual.TabStops), "Validate InsertParagraphTabStopsWithoutNodePath response.");
    assert.Equal(t, 100.0, *actual.TabStops[1].Position, "Validate InsertParagraphTabStopsWithoutNodePath response.");


}

// Test for deleting all paragraph tab stops.
func Test_Paragraph_DeleteAllParagraphTabStops(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    tabStopFolder := "DocumentElements/Paragraphs"
    remoteFileName := "TestDeleteAllParagraphTabStops.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(tabStopFolder + "/ParagraphTabStops.docx"), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.DeleteAllParagraphTabStops(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.TabStops, "Validate DeleteAllParagraphTabStops response.");
    assert.Equal(t, 0, len(actual.TabStops), "Validate DeleteAllParagraphTabStops response.");
}

// Test for deleting all paragraph tab stops without node path.
func Test_Paragraph_DeleteAllParagraphTabStopsWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    tabStopFolder := "DocumentElements/Paragraphs"
    remoteFileName := "TestDeleteAllParagraphTabStopsWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(tabStopFolder + "/ParagraphTabStops.docx"), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.DeleteAllParagraphTabStops(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.TabStops, "Validate DeleteAllParagraphTabStopsWithoutNodePath response.");
    assert.Equal(t, 0, len(actual.TabStops), "Validate DeleteAllParagraphTabStopsWithoutNodePath response.");
}

// Test for deleting a tab stops.
func Test_Paragraph_DeleteParagraphTabStop(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    tabStopFolder := "DocumentElements/Paragraphs"
    remoteFileName := "TestDeleteParagraphTabStop.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(tabStopFolder + "/ParagraphTabStops.docx"), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.DeleteParagraphTabStop(ctx, remoteFileName, 72.0, int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.TabStops, "Validate DeleteParagraphTabStop response.");
    assert.Equal(t, 1, len(actual.TabStops), "Validate DeleteParagraphTabStop response.");
}

// Test for deleting a tab stops without node path.
func Test_Paragraph_DeleteParagraphTabStopWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    tabStopFolder := "DocumentElements/Paragraphs"
    remoteFileName := "TestDeleteParagraphTabStopWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(tabStopFolder + "/ParagraphTabStops.docx"), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.DeleteParagraphTabStop(ctx, remoteFileName, 72.0, int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.TabStops, "Validate DeleteParagraphTabStopWithoutNodePath response.");
    assert.Equal(t, 1, len(actual.TabStops), "Validate DeleteParagraphTabStopWithoutNodePath response.");
}
