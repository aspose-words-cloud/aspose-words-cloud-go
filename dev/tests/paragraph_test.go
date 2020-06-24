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
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/dev/api/models"
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
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetParagraph(ctx, remoteFileName, "sections/0", int32(0), options)

    if err != nil {
        t.Error(err)
    }
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
    _, _, err := client.WordsApi.GetParagraphWithoutNodePath(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }
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
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetParagraphs(ctx, remoteFileName, "sections/0", options)

    if err != nil {
        t.Error(err)
    }
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
    _, _, err := client.WordsApi.GetParagraphsWithoutNodePath(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }
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
    _, _, err := client.WordsApi.GetRun(ctx, remoteFileName, "paragraphs/0", int32(0), options)

    if err != nil {
        t.Error(err)
    }
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
    _, _, err := client.WordsApi.GetRunFont(ctx, remoteFileName, "paragraphs/0", int32(0), options)

    if err != nil {
        t.Error(err)
    }
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
    _, _, err := client.WordsApi.GetRuns(ctx, remoteFileName, "sections/0/paragraphs/0", options)

    if err != nil {
        t.Error(err)
    }
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
        Bold: true,
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }
    _, _, err := client.WordsApi.UpdateRunFont(ctx, remoteFileName, requestFontDto, "paragraphs/0", int32(0), options)

    if err != nil {
        t.Error(err)
    }
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
        Text: "This is a new paragraph for your document",
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.InsertParagraph(ctx, remoteFileName, requestParagraph, "sections/0", options)

    if err != nil {
        t.Error(err)
    }
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
        Text: "This is a new paragraph for your document",
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.InsertParagraphWithoutNodePath(ctx, remoteFileName, requestParagraph, options)

    if err != nil {
        t.Error(err)
    }
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
        "folder": remoteDataFolder,
    }
    _, err := client.WordsApi.RenderParagraph(ctx, remoteFileName, "png", "", int32(0), options)

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
    _, err := client.WordsApi.RenderParagraphWithoutNodePath(ctx, remoteFileName, "png", int32(0), options)

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
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetParagraphFormat(ctx, remoteFileName, "", int32(0), options)

    if err != nil {
        t.Error(err)
    }
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
    _, _, err := client.WordsApi.GetParagraphFormatWithoutNodePath(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }
}

// Test for updating  paragraph format settings.
func Test_Paragraph_UpdateParagraphFormat(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetDocumentParagraphs.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestDto := models.ParagraphFormat{
        Alignment: "Right",
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.UpdateParagraphFormat(ctx, remoteFileName, requestDto, "", int32(0), options)

    if err != nil {
        t.Error(err)
    }
}

// Test for updating  paragraph format settings without node path.
func Test_Paragraph_UpdateParagraphFormatWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetDocumentParagraphsWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestDto := models.ParagraphFormat{
        Alignment: "Right",
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.UpdateParagraphFormatWithoutNodePath(ctx, remoteFileName, requestDto, int32(0), options)

    if err != nil {
        t.Error(err)
    }
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
        "folder": remoteDataFolder,
    }
    _, err := client.WordsApi.DeleteParagraph(ctx, remoteFileName, "", int32(0), options)

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
    _, err := client.WordsApi.DeleteParagraphWithoutNodePath(ctx, remoteFileName, int32(0), options)

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
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetParagraphListFormat(ctx, remoteFileName, "", int32(0), options)

    if err != nil {
        t.Error(err)
    }
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
    _, _, err := client.WordsApi.GetParagraphListFormatWithoutNodePath(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }
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
        ListId: int32(2),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.UpdateParagraphListFormat(ctx, remoteFileName, requestDto, "", int32(0), options)

    if err != nil {
        t.Error(err)
    }
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
        ListId: int32(2),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.UpdateParagraphListFormatWithoutNodePath(ctx, remoteFileName, requestDto, int32(0), options)

    if err != nil {
        t.Error(err)
    }
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
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.DeleteParagraphListFormat(ctx, remoteFileName, "", int32(0), options)

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
    _, _, err := client.WordsApi.DeleteParagraphListFormatWithoutNodePath(ctx, remoteFileName, int32(0), options)

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
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetParagraphTabStops(ctx, remoteFileName, "", int32(0), options)

    if err != nil {
        t.Error(err)
    }
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
    _, _, err := client.WordsApi.GetParagraphTabStopsWithoutNodePath(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }
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
        Position: 72,
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.InsertOrUpdateParagraphTabStop(ctx, remoteFileName, requestDto, "", int32(0), options)

    if err != nil {
        t.Error(err)
    }
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
        Position: 72,
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.InsertOrUpdateParagraphTabStopWithoutNodePath(ctx, remoteFileName, requestDto, int32(0), options)

    if err != nil {
        t.Error(err)
    }
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
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.DeleteAllParagraphTabStops(ctx, remoteFileName, "", int32(0), options)

    if err != nil {
        t.Error(err)
    }
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
    _, _, err := client.WordsApi.DeleteAllParagraphTabStopsWithoutNodePath(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }
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
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.DeleteParagraphTabStop(ctx, remoteFileName, 72, "", int32(0), options)

    if err != nil {
        t.Error(err)
    }
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
    _, _, err := client.WordsApi.DeleteParagraphTabStopWithoutNodePath(ctx, remoteFileName, 72, int32(0), options)

    if err != nil {
        t.Error(err)
    }
}
