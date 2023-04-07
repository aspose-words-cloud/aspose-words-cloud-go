/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="paragraph_test.go">
 *   Copyright (c) 2023 Aspose.Words for Cloud
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
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2303/api/models"
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

    request := &models.GetParagraphRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetParagraph(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Paragraph, "Validate GetDocumentParagraphByIndex response.");
    assert.Equal(t, "0.0.0", actual.Paragraph.NodeId, "Validate GetDocumentParagraphByIndex response.");
}

// Test for getting paragraph online.
func Test_Paragraph_GetDocumentParagraphOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
        "nodePath": "sections/0",
    }

    request := &models.GetParagraphOnlineRequest{
        Document: requestDocument,
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetParagraphOnline(ctx, request)
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

    request := &models.GetParagraphRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetParagraph(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Paragraph, "Validate GetDocumentParagraphByIndexWithoutNodePath response.");
    assert.Equal(t, "0.0.0", actual.Paragraph.NodeId, "Validate GetDocumentParagraphByIndexWithoutNodePath response.");
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

    request := &models.GetParagraphsRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetParagraphs(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Paragraphs, "Validate GetDocumentParagraphs response.");
    assert.NotNil(t, actual.Paragraphs.ParagraphLinkList, "Validate GetDocumentParagraphs response.");
    assert.Equal(t, 15, len(actual.Paragraphs.ParagraphLinkList), "Validate GetDocumentParagraphs response.");
    assert.Equal(t, "Page 1 of 3", actual.Paragraphs.ParagraphLinkList[0].Text, "Validate GetDocumentParagraphs response.");
}

// Test for getting all paragraphs online.
func Test_Paragraph_GetDocumentParagraphsOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
        "nodePath": "sections/0",
    }

    request := &models.GetParagraphsOnlineRequest{
        Document: requestDocument,
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetParagraphsOnline(ctx, request)
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

    request := &models.GetParagraphsRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetParagraphs(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Paragraphs, "Validate GetDocumentParagraphsWithoutNodePath response.");
    assert.NotNil(t, actual.Paragraphs.ParagraphLinkList, "Validate GetDocumentParagraphsWithoutNodePath response.");
    assert.Equal(t, 15, len(actual.Paragraphs.ParagraphLinkList), "Validate GetDocumentParagraphsWithoutNodePath response.");
    assert.Equal(t, "Page 1 of 3", actual.Paragraphs.ParagraphLinkList[0].Text, "Validate GetDocumentParagraphsWithoutNodePath response.");
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

    request := &models.GetRunRequest{
        Name: ToStringPointer(remoteFileName),
        ParagraphPath: ToStringPointer("paragraphs/0"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetRun(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Run, "Validate GetDocumentParagraphRun response.");
    assert.Equal(t, "Page ", actual.Run.Text, "Validate GetDocumentParagraphRun response.");
}

// Test for getting paragraph run online.
func Test_Paragraph_GetDocumentParagraphRunOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.GetRunOnlineRequest{
        Document: requestDocument,
        ParagraphPath: ToStringPointer("paragraphs/0"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetRunOnline(ctx, request)
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

    request := &models.GetRunFontRequest{
        Name: ToStringPointer(remoteFileName),
        ParagraphPath: ToStringPointer("paragraphs/0"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetRunFont(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Font, "Validate GetDocumentParagraphRunFont response.");
    assert.Equal(t, "Times New Roman", actual.Font.Name, "Validate GetDocumentParagraphRunFont response.");
}

// Test for getting paragraph run font online.
func Test_Paragraph_GetDocumentParagraphRunFontOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.GetRunFontOnlineRequest{
        Document: requestDocument,
        ParagraphPath: ToStringPointer("paragraphs/0"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetRunFontOnline(ctx, request)
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

    request := &models.GetRunsRequest{
        Name: ToStringPointer(remoteFileName),
        ParagraphPath: ToStringPointer("sections/0/paragraphs/0"),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetRuns(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Runs, "Validate GetParagraphRuns response.");
    assert.NotNil(t, actual.Runs.List, "Validate GetParagraphRuns response.");
    assert.Equal(t, 6, len(actual.Runs.List), "Validate GetParagraphRuns response.");
    assert.Equal(t, "Page ", actual.Runs.List[0].Text, "Validate GetParagraphRuns response.");
}

// Test for getting paragraph runs online.
func Test_Paragraph_GetParagraphRunsOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.GetRunsOnlineRequest{
        Document: requestDocument,
        ParagraphPath: ToStringPointer("sections/0/paragraphs/0"),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetRunsOnline(ctx, request)
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
        Bold: ToBoolPointer(true),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }

    request := &models.UpdateRunFontRequest{
        Name: ToStringPointer(remoteFileName),
        ParagraphPath: ToStringPointer("paragraphs/0"),
        Index: ToInt32Pointer(int32(0)),
        FontDto: &requestFontDto,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.UpdateRunFont(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Font, "Validate UpdateRunFont response.");
    assert.True(t, actual.Font.Bold, "Validate UpdateRunFont response.");
}

// Test for updating paragraph run font online.
func Test_Paragraph_UpdateRunFontOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)
    requestFontDto := models.Font{
        Bold: ToBoolPointer(true),
    }

    options := map[string]interface{}{
    }

    request := &models.UpdateRunFontOnlineRequest{
        Document: requestDocument,
        ParagraphPath: ToStringPointer("paragraphs/0"),
        FontDto: &requestFontDto,
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.UpdateRunFontOnline(ctx, request)
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
        Text: ToStringPointer("This is a new paragraph for your document"),
    }

    options := map[string]interface{}{
        "nodePath": "sections/0",
        "folder": remoteDataFolder,
    }

    request := &models.InsertParagraphRequest{
        Name: ToStringPointer(remoteFileName),
        Paragraph: &requestParagraph,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.InsertParagraph(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Paragraph, "Validate InsertParagraph response.");
    assert.Equal(t, "0.3.8", actual.Paragraph.NodeId, "Validate InsertParagraph response.");
}

// Test for adding paragraph online.
func Test_Paragraph_InsertParagraphOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)
    requestParagraph := models.ParagraphInsert{
        Text: ToStringPointer("This is a new paragraph for your document"),
    }

    options := map[string]interface{}{
        "nodePath": "sections/0",
    }

    request := &models.InsertParagraphOnlineRequest{
        Document: requestDocument,
        Paragraph: &requestParagraph,
        Optionals: options,
    }

    _, _, err := client.WordsApi.InsertParagraphOnline(ctx, request)
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
        Text: ToStringPointer("This is a new paragraph for your document"),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.InsertParagraphRequest{
        Name: ToStringPointer(remoteFileName),
        Paragraph: &requestParagraph,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.InsertParagraph(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Paragraph, "Validate InsertParagraphWithoutNodePath response.");
    assert.Equal(t, "0.3.8", actual.Paragraph.NodeId, "Validate InsertParagraphWithoutNodePath response.");
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

    request := &models.RenderParagraphRequest{
        Name: ToStringPointer(remoteFileName),
        Format: ToStringPointer("png"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, err := client.WordsApi.RenderParagraph(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for paragraph rendering.
func Test_Paragraph_RenderParagraphOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
        "nodePath": "",
    }

    request := &models.RenderParagraphOnlineRequest{
        Document: requestDocument,
        Format: ToStringPointer("png"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, err := client.WordsApi.RenderParagraphOnline(ctx, request)
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

    request := &models.RenderParagraphRequest{
        Name: ToStringPointer(remoteFileName),
        Format: ToStringPointer("png"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, err := client.WordsApi.RenderParagraph(ctx, request)
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

    request := &models.GetParagraphFormatRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetParagraphFormat(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.ParagraphFormat, "Validate GetParagraphFormat response.");
    assert.Equal(t, "Normal", actual.ParagraphFormat.StyleName, "Validate GetParagraphFormat response.");
}

// Test for getting paragraph format settings online.
func Test_Paragraph_GetParagraphFormatOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
        "nodePath": "",
    }

    request := &models.GetParagraphFormatOnlineRequest{
        Document: requestDocument,
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetParagraphFormatOnline(ctx, request)
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

    request := &models.GetParagraphFormatRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetParagraphFormat(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.ParagraphFormat, "Validate GetParagraphFormatWithoutNodePath response.");
    assert.Equal(t, "Normal", actual.ParagraphFormat.StyleName, "Validate GetParagraphFormatWithoutNodePath response.");
}

// Test for updating  paragraph format settings.
func Test_Paragraph_UpdateParagraphFormat(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetDocumentParagraphs.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestParagraphFormatDto := models.ParagraphFormatUpdate{
        Alignment: ToStringPointer("Right"),
    }

    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }

    request := &models.UpdateParagraphFormatRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        ParagraphFormatDto: &requestParagraphFormatDto,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.UpdateParagraphFormat(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.ParagraphFormat, "Validate UpdateParagraphFormat response.");

}

// Test for updating  paragraph format settings online.
func Test_Paragraph_UpdateParagraphFormatOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)
    requestParagraphFormatDto := models.ParagraphFormatUpdate{
        Alignment: ToStringPointer("Right"),
    }

    options := map[string]interface{}{
        "nodePath": "",
    }

    request := &models.UpdateParagraphFormatOnlineRequest{
        Document: requestDocument,
        ParagraphFormatDto: &requestParagraphFormatDto,
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.UpdateParagraphFormatOnline(ctx, request)
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
        "nodePath": "",
        "folder": remoteDataFolder,
    }

    request := &models.DeleteParagraphRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, err := client.WordsApi.DeleteParagraph(ctx, request)

    if err != nil {
        t.Error(err)
    }

}

// Test for deleting  a paragraph online.
func Test_Paragraph_DeleteParagraphOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
        "nodePath": "",
    }

    request := &models.DeleteParagraphOnlineRequest{
        Document: requestDocument,
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.DeleteParagraphOnline(ctx, request)
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

    request := &models.DeleteParagraphRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, err := client.WordsApi.DeleteParagraph(ctx, request)

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

    request := &models.GetParagraphListFormatRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetParagraphListFormat(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.ListFormat, "Validate GetParagraphListFormat response.");
    assert.Equal(t, int32(1), actual.ListFormat.ListId, "Validate GetParagraphListFormat response.");
}

// Test for getting paragraph list format online.
func Test_Paragraph_GetParagraphListFormatOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    listFolder := "DocumentElements/ParagraphListFormat"

    requestDocument := OpenFile(t, listFolder + "/ParagraphGetListFormat.doc")

    options := map[string]interface{}{
        "nodePath": "",
    }

    request := &models.GetParagraphListFormatOnlineRequest{
        Document: requestDocument,
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetParagraphListFormatOnline(ctx, request)
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

    request := &models.GetParagraphListFormatRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetParagraphListFormat(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.ListFormat, "Validate GetParagraphListFormatWithoutNodePath response.");
    assert.Equal(t, int32(1), actual.ListFormat.ListId, "Validate GetParagraphListFormatWithoutNodePath response.");
}

// Test for updating paragraph list format.
func Test_Paragraph_UpdateParagraphListFormat(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    listFolder := "DocumentElements/ParagraphListFormat"
    remoteFileName := "TestUpdateParagraphListFormat.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(listFolder + "/ParagraphUpdateListFormat.doc"), remoteDataFolder + "/" + remoteFileName)

    requestListFormatDto := models.ListFormatUpdate{
        ListId: ToInt32Pointer(int32(2)),
    }

    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }

    request := &models.UpdateParagraphListFormatRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        ListFormatDto: &requestListFormatDto,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.UpdateParagraphListFormat(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.ListFormat, "Validate UpdateParagraphListFormat response.");
    assert.Equal(t, int32(2), actual.ListFormat.ListId, "Validate UpdateParagraphListFormat response.");
}

// Test for updating paragraph list format online.
func Test_Paragraph_UpdateParagraphListFormatOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    listFolder := "DocumentElements/ParagraphListFormat"

    requestDocument := OpenFile(t, listFolder + "/ParagraphUpdateListFormat.doc")
    requestListFormatDto := models.ListFormatUpdate{
        ListId: ToInt32Pointer(int32(2)),
    }

    options := map[string]interface{}{
        "nodePath": "",
    }

    request := &models.UpdateParagraphListFormatOnlineRequest{
        Document: requestDocument,
        ListFormatDto: &requestListFormatDto,
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.UpdateParagraphListFormatOnline(ctx, request)
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

    requestListFormatDto := models.ListFormatUpdate{
        ListId: ToInt32Pointer(int32(2)),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.UpdateParagraphListFormatRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        ListFormatDto: &requestListFormatDto,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.UpdateParagraphListFormat(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.ListFormat, "Validate UpdateParagraphListFormatWithoutNodePath response.");
    assert.Equal(t, int32(2), actual.ListFormat.ListId, "Validate UpdateParagraphListFormatWithoutNodePath response.");
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

    request := &models.DeleteParagraphListFormatRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.DeleteParagraphListFormat(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for deleting paragraph list format online.
func Test_Paragraph_DeleteParagraphListFormatOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    listFolder := "DocumentElements/ParagraphListFormat"

    requestDocument := OpenFile(t, listFolder + "/ParagraphDeleteListFormat.doc")

    options := map[string]interface{}{
        "nodePath": "",
    }

    request := &models.DeleteParagraphListFormatOnlineRequest{
        Document: requestDocument,
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.DeleteParagraphListFormatOnline(ctx, request)
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

    request := &models.DeleteParagraphListFormatRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.DeleteParagraphListFormat(ctx, request)
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

    request := &models.GetParagraphTabStopsRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetParagraphTabStops(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.TabStops, "Validate GetParagraphTabStops response.");
    assert.Equal(t, 2, len(actual.TabStops), "Validate GetParagraphTabStops response.");
    assert.Equal(t, 72.0, actual.TabStops[0].Position, "Validate GetParagraphTabStops response.");
}

// Test for getting paragraph tab stops online.
func Test_Paragraph_GetParagraphTabStopsOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    tabStopFolder := "DocumentElements/Paragraphs"

    requestDocument := OpenFile(t, tabStopFolder + "/ParagraphTabStops.docx")

    options := map[string]interface{}{
        "nodePath": "",
    }

    request := &models.GetParagraphTabStopsOnlineRequest{
        Document: requestDocument,
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetParagraphTabStopsOnline(ctx, request)
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

    request := &models.GetParagraphTabStopsRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetParagraphTabStops(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.TabStops, "Validate GetParagraphTabStopsWithoutNodePath response.");
    assert.Equal(t, 2, len(actual.TabStops), "Validate GetParagraphTabStopsWithoutNodePath response.");
    assert.Equal(t, 72.0, actual.TabStops[0].Position, "Validate GetParagraphTabStopsWithoutNodePath response.");
}

// Test for inserting paragraph tab stop.
func Test_Paragraph_InsertParagraphTabStops(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
    tabStopFolder := "DocumentElements/Paragraphs"
    remoteFileName := "TestInsertOrUpdateParagraphTabStop.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(tabStopFolder + "/ParagraphTabStops.docx"), remoteDataFolder + "/" + remoteFileName)

    requestTabStopInsertDto := models.TabStopInsert{
        Alignment: ToStringPointer("Left"),
        Leader: ToStringPointer("None"),
        Position: ToFloat64Pointer(100.0),
    }

    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }

    request := &models.InsertOrUpdateParagraphTabStopRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        TabStopInsertDto: &requestTabStopInsertDto,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.InsertOrUpdateParagraphTabStop(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.TabStops, "Validate InsertParagraphTabStops response.");
    assert.Equal(t, 3, len(actual.TabStops), "Validate InsertParagraphTabStops response.");
    assert.Equal(t, 100.0, actual.TabStops[1].Position, "Validate InsertParagraphTabStops response.");


}

// Test for inserting paragraph tab stop online.
func Test_Paragraph_InsertParagraphTabStopsOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    tabStopFolder := "DocumentElements/Paragraphs"

    requestDocument := OpenFile(t, tabStopFolder + "/ParagraphTabStops.docx")
    requestTabStopInsertDto := models.TabStopInsert{
        Alignment: ToStringPointer("Left"),
        Leader: ToStringPointer("None"),
        Position: ToFloat64Pointer(72),
    }

    options := map[string]interface{}{
        "nodePath": "",
    }

    request := &models.InsertOrUpdateParagraphTabStopOnlineRequest{
        Document: requestDocument,
        TabStopInsertDto: &requestTabStopInsertDto,
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.InsertOrUpdateParagraphTabStopOnline(ctx, request)
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

    requestTabStopInsertDto := models.TabStopInsert{
        Alignment: ToStringPointer("Left"),
        Leader: ToStringPointer("None"),
        Position: ToFloat64Pointer(100.0),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.InsertOrUpdateParagraphTabStopRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        TabStopInsertDto: &requestTabStopInsertDto,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.InsertOrUpdateParagraphTabStop(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.TabStops, "Validate InsertParagraphTabStopsWithoutNodePath response.");
    assert.Equal(t, 3, len(actual.TabStops), "Validate InsertParagraphTabStopsWithoutNodePath response.");
    assert.Equal(t, 100.0, actual.TabStops[1].Position, "Validate InsertParagraphTabStopsWithoutNodePath response.");


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

    request := &models.DeleteAllParagraphTabStopsRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.DeleteAllParagraphTabStops(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.TabStops, "Validate DeleteAllParagraphTabStops response.");
    assert.Equal(t, 0, len(actual.TabStops), "Validate DeleteAllParagraphTabStops response.");
}

// Test for deleting all paragraph tab stops online.
func Test_Paragraph_DeleteAllParagraphTabStopsOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    tabStopFolder := "DocumentElements/Paragraphs"

    requestDocument := OpenFile(t, tabStopFolder + "/ParagraphTabStops.docx")

    options := map[string]interface{}{
        "nodePath": "",
    }

    request := &models.DeleteAllParagraphTabStopsOnlineRequest{
        Document: requestDocument,
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.DeleteAllParagraphTabStopsOnline(ctx, request)
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

    request := &models.DeleteAllParagraphTabStopsRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.DeleteAllParagraphTabStops(ctx, request)
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

    request := &models.DeleteParagraphTabStopRequest{
        Name: ToStringPointer(remoteFileName),
        Position: ToFloat64Pointer(72.0),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.DeleteParagraphTabStop(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.TabStops, "Validate DeleteParagraphTabStop response.");
    assert.Equal(t, 1, len(actual.TabStops), "Validate DeleteParagraphTabStop response.");
}

// Test for deleting a tab stops online.
func Test_Paragraph_DeleteParagraphTabStopOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    tabStopFolder := "DocumentElements/Paragraphs"

    requestDocument := OpenFile(t, tabStopFolder + "/ParagraphTabStops.docx")

    options := map[string]interface{}{
        "nodePath": "",
    }

    request := &models.DeleteParagraphTabStopOnlineRequest{
        Document: requestDocument,
        Position: ToFloat64Pointer(72.0),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.DeleteParagraphTabStopOnline(ctx, request)
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

    request := &models.DeleteParagraphTabStopRequest{
        Name: ToStringPointer(remoteFileName),
        Position: ToFloat64Pointer(72.0),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.DeleteParagraphTabStop(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.TabStops, "Validate DeleteParagraphTabStopWithoutNodePath response.");
    assert.Equal(t, 1, len(actual.TabStops), "Validate DeleteParagraphTabStopWithoutNodePath response.");
}
