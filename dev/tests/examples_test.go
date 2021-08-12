/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="examplestest.go">
 *   Copyright (c) 2021 Aspose.Words for Cloud
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
package api_test

import (
    "os"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/dev/api/models"
)

func Test_Examples_AcceptAllRevisions(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    remoteFileName:= "Sample.docx"

    acceptRequestOptions := map[string]interface{}{"destFileName": remoteFileName,}
    acceptRequest := &models.AcceptAllRevisionsRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: acceptRequestOptions,
    }
    _, _, _ = wordsApi.AcceptAllRevisions(ctx, acceptRequest)
}



// func Test_Examples_AcceptAllRevisionsOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_AppendDocument(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    remoteFileName:= "Sample.docx"

    requestDocumentListDocumentEntries0 := models.DocumentEntry{
        Href: ToStringPointer(remoteFileName),
        ImportFormatMode: ToStringPointer("KeepSourceFormatting"),
    }
    requestDocumentListDocumentEntries := []models.DocumentEntry{
        requestDocumentListDocumentEntries0,
    }
    requestDocumentList := models.DocumentEntryList{
        DocumentEntries: requestDocumentListDocumentEntries,
    }
    appendRequestOptions := map[string]interface{}{"destFileName": remoteFileName,}
    appendRequest := &models.AppendDocumentRequest{
        Name: ToStringPointer(remoteFileName),
        DocumentList: requestDocumentList,
        Optionals: appendRequestOptions,
    }
    _, _, _ = wordsApi.AppendDocument(ctx, appendRequest)
}



// func Test_Examples_AppendDocumentOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_ApplyStyleToDocumentElement(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestStyleApply := models.StyleApply{
        StyleName: ToStringPointer("Heading 1"),
    }
    applyStyleRequestOptions := map[string]interface{}{}
    applyStyleRequest := &models.ApplyStyleToDocumentElementRequest{
        Name: ToStringPointer("Sample.docx"),
        StyledNodePath: ToStringPointer("paragraphs/1/paragraphFormat"),
        StyleApply: requestStyleApply,
        Optionals: applyStyleRequestOptions,
    }
    _, _, _ = wordsApi.ApplyStyleToDocumentElement(ctx, applyStyleRequest)
}



// func Test_Examples_ApplyStyleToDocumentElementOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_BuildReport(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestReportEngineSettingsReportBuildOptions := []string{
        "AllowMissingMembers",
        "RemoveEmptyParagraphs",
    }
    requestReportEngineSettings := models.ReportEngineSettings{
        DataSourceType: ToStringPointer("Json"),
        ReportBuildOptions: requestReportEngineSettingsReportBuildOptions,
    }
    buildReportRequestOptions := map[string]interface{}{}
    buildReportRequest := &models.BuildReportRequest{
        Name: ToStringPointer("Sample.docx"),
        Data: ToStringPointer("Data.json"),
        ReportEngineSettings: requestReportEngineSettings,
        Optionals: buildReportRequestOptions,
    }
    _, _, _ = wordsApi.BuildReport(ctx, buildReportRequest)
}



// func Test_Examples_BuildReportOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_Classify(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    classifyRequestOptions := map[string]interface{}{"bestClassesCount": "3",}
    classifyRequest := &models.ClassifyRequest{
        Text: ToStringPointer("Try text classification"),
        Optionals: classifyRequestOptions,
    }
    _, _, _ = wordsApi.Classify(ctx, classifyRequest)
}



func Test_Examples_ClassifyDocument(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    classifyRequestOptions := map[string]interface{}{"bestClassesCount": "3",}
    classifyRequest := &models.ClassifyDocumentRequest{
        Name: ToStringPointer("Sample.docx"),
        Optionals: classifyRequestOptions,
    }
    _, _, _ = wordsApi.ClassifyDocument(ctx, classifyRequest)
}



// func Test_Examples_ClassifyDocumentOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_CompareDocument(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestCompareData := models.CompareData{
        Author: ToStringPointer("author"),
        ComparingWithDocument: ToStringPointer("TestCompareDocument2.doc"),
        DateTime: ToTimePointer(CreateTime(2015, 10, 26, 0, 0, 0)),
    }
    compareRequestOptions := map[string]interface{}{"destFileName": "/TestCompareDocumentOut.doc",}
    compareRequest := &models.CompareDocumentRequest{
        Name: ToStringPointer("TestCompareDocument1.doc"),
        CompareData: requestCompareData,
        Optionals: compareRequestOptions,
    }
    _, _, _ = wordsApi.CompareDocument(ctx, compareRequest)
}



// func Test_Examples_CompareDocumentOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_ConvertDocument(t *testing.T) {
    documentsDir := GetExamplesDir()
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestDocumentFileData, _ := os.Open(documentsDir + "/" + "Sample.docx")
    convertRequestOptions := map[string]interface{}{}
    convertRequest := &models.ConvertDocumentRequest{
        Document: requestDocumentFileData,
        Format: ToStringPointer("pdf"),
        Optionals: convertRequestOptions,
    }
    _, _ = wordsApi.ConvertDocument(ctx, convertRequest)
}



func Test_Examples_CopyFile(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    copyRequestOptions := map[string]interface{}{}
    copyRequest := &models.CopyFileRequest{
        DestPath: ToStringPointer("/TestCopyFileDest.docx"),
        SrcPath: ToStringPointer("Sample.docx"),
        Optionals: copyRequestOptions,
    }
    _, _ = wordsApi.CopyFile(ctx, copyRequest)
}



func Test_Examples_CopyFolder(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    folderToCopy:= "/TestCopyFolder"

    copyRequestOptions := map[string]interface{}{}
    copyRequest := &models.CopyFolderRequest{
        DestPath: ToStringPointer(folderToCopy + "Dest"),
        SrcPath: ToStringPointer(folderToCopy + "Src"),
        Optionals: copyRequestOptions,
    }
    _, _ = wordsApi.CopyFolder(ctx, copyRequest)
}



func Test_Examples_CopyStyle(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestStyleCopy := models.StyleCopy{
        StyleName: ToStringPointer("Heading 1"),
    }
    copyRequestOptions := map[string]interface{}{}
    copyRequest := &models.CopyStyleRequest{
        Name: ToStringPointer("Sample.docx"),
        StyleCopy: requestStyleCopy,
        Optionals: copyRequestOptions,
    }
    _, _, _ = wordsApi.CopyStyle(ctx, copyRequest)
}



// func Test_Examples_CopyStyleOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_CreateDocument(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    createRequestOptions := map[string]interface{}{"fileName": "Sample.docx",}
    createRequest := &models.CreateDocumentRequest{
        Optionals: createRequestOptions,
    }
    _, _, _ = wordsApi.CreateDocument(ctx, createRequest)
}



func Test_Examples_CreateFolder(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    createRequestOptions := map[string]interface{}{}
    createRequest := &models.CreateFolderRequest{
        Path: ToStringPointer("/TestCreateFolder"),
        Optionals: createRequestOptions,
    }
    _, _ = wordsApi.CreateFolder(ctx, createRequest)
}



func Test_Examples_CreateOrUpdateDocumentProperty(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    remoteFileName:= "Sample.docx"

    requestProperty := models.DocumentPropertyCreateOrUpdate{
        Value: ToStringPointer("Imran Anwar"),
    }
    createRequestOptions := map[string]interface{}{"destFileName": remoteFileName,}
    createRequest := &models.CreateOrUpdateDocumentPropertyRequest{
        Name: ToStringPointer(remoteFileName),
        PropertyName: ToStringPointer("AsposeAuthor"),
        Property: requestProperty,
        Optionals: createRequestOptions,
    }
    _, _, _ = wordsApi.CreateOrUpdateDocumentProperty(ctx, createRequest)
}



// func Test_Examples_CreateOrUpdateDocumentPropertyOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_DeleteAllParagraphTabStops(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    deleteRequestOptions := map[string]interface{}{}
    deleteRequest := &models.DeleteAllParagraphTabStopsRequest{
        Name: ToStringPointer("Sample.docx"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: deleteRequestOptions,
    }
    _, _, _ = wordsApi.DeleteAllParagraphTabStops(ctx, deleteRequest)
}



// func Test_Examples_DeleteAllParagraphTabStopsOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_DeleteBorder(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    deleteRequestOptions := map[string]interface{}{"nodePath": "tables/1/rows/0/cells/0",}
    deleteRequest := &models.DeleteBorderRequest{
        Name: ToStringPointer("Sample.docx"),
        BorderType: ToStringPointer("left"),
        Optionals: deleteRequestOptions,
    }
    _, _, _ = wordsApi.DeleteBorder(ctx, deleteRequest)
}



// func Test_Examples_DeleteBorderOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_DeleteBorders(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    deleteRequestOptions := map[string]interface{}{"nodePath": "tables/1/rows/0/cells/0",}
    deleteRequest := &models.DeleteBordersRequest{
        Name: ToStringPointer("Sample.docx"),
        Optionals: deleteRequestOptions,
    }
    _, _, _ = wordsApi.DeleteBorders(ctx, deleteRequest)
}



// func Test_Examples_DeleteBordersOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_DeleteComment(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    remoteFileName:= "Sample.docx"

    deleteRequestOptions := map[string]interface{}{"destFileName": remoteFileName,}
    deleteRequest := &models.DeleteCommentRequest{
        Name: ToStringPointer(remoteFileName),
        CommentIndex: ToInt32Pointer(int32(0)),
        Optionals: deleteRequestOptions,
    }
    _, _ = wordsApi.DeleteComment(ctx, deleteRequest)
}



// func Test_Examples_DeleteCommentOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_DeleteComments(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    remoteFileName:= "Sample.docx"

    deleteRequestOptions := map[string]interface{}{"destFileName": remoteFileName,}
    deleteRequest := &models.DeleteCommentsRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: deleteRequestOptions,
    }
    _, _ = wordsApi.DeleteComments(ctx, deleteRequest)
}



// func Test_Examples_DeleteCommentsOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_DeleteCustomXmlPart(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    remoteFileName:= "Sample.docx"

    deleteRequestOptions := map[string]interface{}{"destFileName": remoteFileName,}
    deleteRequest := &models.DeleteCustomXmlPartRequest{
        Name: ToStringPointer(remoteFileName),
        CustomXmlPartIndex: ToInt32Pointer(int32(0)),
        Optionals: deleteRequestOptions,
    }
    _, _ = wordsApi.DeleteCustomXmlPart(ctx, deleteRequest)
}



// func Test_Examples_DeleteCustomXmlPartOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_DeleteCustomXmlParts(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    remoteFileName:= "Sample.docx"

    deleteRequestOptions := map[string]interface{}{"destFileName": remoteFileName,}
    deleteRequest := &models.DeleteCustomXmlPartsRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: deleteRequestOptions,
    }
    _, _ = wordsApi.DeleteCustomXmlParts(ctx, deleteRequest)
}



// func Test_Examples_DeleteCustomXmlPartsOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_DeleteDocumentProperty(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    remoteFileName:= "Sample.docx"

    deleteRequestOptions := map[string]interface{}{"destFileName": remoteFileName,}
    deleteRequest := &models.DeleteDocumentPropertyRequest{
        Name: ToStringPointer(remoteFileName),
        PropertyName: ToStringPointer("testProp"),
        Optionals: deleteRequestOptions,
    }
    _, _ = wordsApi.DeleteDocumentProperty(ctx, deleteRequest)
}



// func Test_Examples_DeleteDocumentPropertyOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_DeleteDrawingObject(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    deleteRequestOptions := map[string]interface{}{}
    deleteRequest := &models.DeleteDrawingObjectRequest{
        Name: ToStringPointer("Sample.docx"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: deleteRequestOptions,
    }
    _, _ = wordsApi.DeleteDrawingObject(ctx, deleteRequest)
}



// func Test_Examples_DeleteDrawingObjectOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_DeleteField(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    deleteRequestOptions := map[string]interface{}{}
    deleteRequest := &models.DeleteFieldRequest{
        Name: ToStringPointer("Sample.docx"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: deleteRequestOptions,
    }
    _, _ = wordsApi.DeleteField(ctx, deleteRequest)
}



// func Test_Examples_DeleteFieldOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_DeleteFields(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    deleteRequestOptions := map[string]interface{}{}
    deleteRequest := &models.DeleteFieldsRequest{
        Name: ToStringPointer("Sample.docx"),
        Optionals: deleteRequestOptions,
    }
    _, _ = wordsApi.DeleteFields(ctx, deleteRequest)
}



// func Test_Examples_DeleteFieldsOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_DeleteFile(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    deleteRequestOptions := map[string]interface{}{}
    deleteRequest := &models.DeleteFileRequest{
        Path: ToStringPointer("Sample.docx"),
        Optionals: deleteRequestOptions,
    }
    _, _ = wordsApi.DeleteFile(ctx, deleteRequest)
}



func Test_Examples_DeleteFolder(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    deleteRequestOptions := map[string]interface{}{}
    deleteRequest := &models.DeleteFolderRequest{
        Path: ToStringPointer(""),
        Optionals: deleteRequestOptions,
    }
    _, _ = wordsApi.DeleteFolder(ctx, deleteRequest)
}



func Test_Examples_DeleteFootnote(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    deleteRequestOptions := map[string]interface{}{}
    deleteRequest := &models.DeleteFootnoteRequest{
        Name: ToStringPointer("Sample.docx"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: deleteRequestOptions,
    }
    _, _ = wordsApi.DeleteFootnote(ctx, deleteRequest)
}



// func Test_Examples_DeleteFootnoteOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_DeleteFormField(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    remoteFileName:= "Sample.docx"

    deleteRequestOptions := map[string]interface{}{"destFileName": remoteFileName,}
    deleteRequest := &models.DeleteFormFieldRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: deleteRequestOptions,
    }
    _, _ = wordsApi.DeleteFormField(ctx, deleteRequest)
}



// func Test_Examples_DeleteFormFieldOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_DeleteHeaderFooter(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    deleteRequestOptions := map[string]interface{}{}
    deleteRequest := &models.DeleteHeaderFooterRequest{
        Name: ToStringPointer("Sample.docx"),
        SectionPath: ToStringPointer(""),
        Index: ToInt32Pointer(int32(0)),
        Optionals: deleteRequestOptions,
    }
    _, _ = wordsApi.DeleteHeaderFooter(ctx, deleteRequest)
}



// func Test_Examples_DeleteHeaderFooterOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_DeleteHeadersFooters(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    deleteRequestOptions := map[string]interface{}{}
    deleteRequest := &models.DeleteHeadersFootersRequest{
        Name: ToStringPointer("Sample.docx"),
        SectionPath: ToStringPointer(""),
        Optionals: deleteRequestOptions,
    }
    _, _ = wordsApi.DeleteHeadersFooters(ctx, deleteRequest)
}



// func Test_Examples_DeleteHeadersFootersOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_DeleteMacros(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    deleteRequestOptions := map[string]interface{}{}
    deleteRequest := &models.DeleteMacrosRequest{
        Name: ToStringPointer("Sample.docx"),
        Optionals: deleteRequestOptions,
    }
    _, _ = wordsApi.DeleteMacros(ctx, deleteRequest)
}



// func Test_Examples_DeleteMacrosOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_DeleteOfficeMathObject(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    deleteRequestOptions := map[string]interface{}{}
    deleteRequest := &models.DeleteOfficeMathObjectRequest{
        Name: ToStringPointer("Sample.docx"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: deleteRequestOptions,
    }
    _, _ = wordsApi.DeleteOfficeMathObject(ctx, deleteRequest)
}



// func Test_Examples_DeleteOfficeMathObjectOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_DeleteParagraph(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    deleteRequestOptions := map[string]interface{}{}
    deleteRequest := &models.DeleteParagraphRequest{
        Name: ToStringPointer("Sample.docx"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: deleteRequestOptions,
    }
    _, _ = wordsApi.DeleteParagraph(ctx, deleteRequest)
}



func Test_Examples_DeleteParagraphListFormat(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    deleteRequestOptions := map[string]interface{}{}
    deleteRequest := &models.DeleteParagraphListFormatRequest{
        Name: ToStringPointer("Sample.docx"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: deleteRequestOptions,
    }
    _, _, _ = wordsApi.DeleteParagraphListFormat(ctx, deleteRequest)
}



// func Test_Examples_DeleteParagraphListFormatOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



// func Test_Examples_DeleteParagraphOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_DeleteParagraphTabStop(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    deleteRequestOptions := map[string]interface{}{}
    deleteRequest := &models.DeleteParagraphTabStopRequest{
        Name: ToStringPointer("Sample.docx"),
        Position: ToFloat64Pointer(72.0),
        Index: ToInt32Pointer(int32(0)),
        Optionals: deleteRequestOptions,
    }
    _, _, _ = wordsApi.DeleteParagraphTabStop(ctx, deleteRequest)
}



// func Test_Examples_DeleteParagraphTabStopOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_DeleteRun(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    deleteRequestOptions := map[string]interface{}{}
    deleteRequest := &models.DeleteRunRequest{
        Name: ToStringPointer("Sample.docx"),
        ParagraphPath: ToStringPointer("paragraphs/1"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: deleteRequestOptions,
    }
    _, _ = wordsApi.DeleteRun(ctx, deleteRequest)
}



// func Test_Examples_DeleteRunOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_DeleteSection(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    deleteRequestOptions := map[string]interface{}{}
    deleteRequest := &models.DeleteSectionRequest{
        Name: ToStringPointer("Sample.docx"),
        SectionIndex: ToInt32Pointer(int32(0)),
        Optionals: deleteRequestOptions,
    }
    _, _ = wordsApi.DeleteSection(ctx, deleteRequest)
}



// func Test_Examples_DeleteSectionOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_DeleteTable(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    deleteRequestOptions := map[string]interface{}{}
    deleteRequest := &models.DeleteTableRequest{
        Name: ToStringPointer("Sample.docx"),
        Index: ToInt32Pointer(int32(1)),
        Optionals: deleteRequestOptions,
    }
    _, _ = wordsApi.DeleteTable(ctx, deleteRequest)
}



func Test_Examples_DeleteTableCell(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    deleteRequestOptions := map[string]interface{}{}
    deleteRequest := &models.DeleteTableCellRequest{
        Name: ToStringPointer("Sample.docx"),
        TableRowPath: ToStringPointer("sections/0/tables/2/rows/0"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: deleteRequestOptions,
    }
    _, _ = wordsApi.DeleteTableCell(ctx, deleteRequest)
}



// func Test_Examples_DeleteTableCellOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



// func Test_Examples_DeleteTableOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_DeleteTableRow(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    deleteRequestOptions := map[string]interface{}{}
    deleteRequest := &models.DeleteTableRowRequest{
        Name: ToStringPointer("Sample.docx"),
        TablePath: ToStringPointer("tables/1"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: deleteRequestOptions,
    }
    _, _ = wordsApi.DeleteTableRow(ctx, deleteRequest)
}



// func Test_Examples_DeleteTableRowOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_DeleteWatermark(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    remoteFileName:= "Sample.docx"

    deleteRequestOptions := map[string]interface{}{"destFileName": remoteFileName,}
    deleteRequest := &models.DeleteWatermarkRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: deleteRequestOptions,
    }
    _, _, _ = wordsApi.DeleteWatermark(ctx, deleteRequest)
}



// func Test_Examples_DeleteWatermarkOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_DownloadFile(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    downloadRequestOptions := map[string]interface{}{}
    downloadRequest := &models.DownloadFileRequest{
        Path: ToStringPointer("Sample.docx"),
        Optionals: downloadRequestOptions,
    }
    _, _ = wordsApi.DownloadFile(ctx, downloadRequest)
}



func Test_Examples_ExecuteMailMerge(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    remoteFileName:= "Sample.docx"

    mailMergeRequestOptions := map[string]interface{}{"data": "TestExecuteTemplateData.txt",
    "destFileName": remoteFileName,}
    mailMergeRequest := &models.ExecuteMailMergeRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: mailMergeRequestOptions,
    }
    _, _, _ = wordsApi.ExecuteMailMerge(ctx, mailMergeRequest)
}



// func Test_Examples_ExecuteMailMergeOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetAvailableFonts(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetAvailableFontsRequest{
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetAvailableFonts(ctx, request)
}



func Test_Examples_GetBookmarkByName(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetBookmarkByNameRequest{
        Name: ToStringPointer("Sample.docx"),
        BookmarkName: ToStringPointer("aspose"),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetBookmarkByName(ctx, request)
}



// func Test_Examples_GetBookmarkByNameOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetBookmarks(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetBookmarksRequest{
        Name: ToStringPointer("Sample.docx"),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetBookmarks(ctx, request)
}



// func Test_Examples_GetBookmarksOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetBorder(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{"nodePath": "tables/1/rows/0/cells/0",}
    request := &models.GetBorderRequest{
        Name: ToStringPointer("Sample.docx"),
        BorderType: ToStringPointer("left"),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetBorder(ctx, request)
}



// func Test_Examples_GetBorderOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetBorders(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{"nodePath": "tables/1/rows/0/cells/0",}
    request := &models.GetBordersRequest{
        Name: ToStringPointer("Sample.docx"),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetBorders(ctx, request)
}



// func Test_Examples_GetBordersOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetComment(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetCommentRequest{
        Name: ToStringPointer("Sample.docx"),
        CommentIndex: ToInt32Pointer(int32(0)),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetComment(ctx, request)
}



// func Test_Examples_GetCommentOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetComments(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetCommentsRequest{
        Name: ToStringPointer("Sample.docx"),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetComments(ctx, request)
}



// func Test_Examples_GetCommentsOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetCustomXmlPart(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetCustomXmlPartRequest{
        Name: ToStringPointer("Sample.docx"),
        CustomXmlPartIndex: ToInt32Pointer(int32(0)),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetCustomXmlPart(ctx, request)
}



// func Test_Examples_GetCustomXmlPartOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetCustomXmlParts(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetCustomXmlPartsRequest{
        Name: ToStringPointer("Sample.docx"),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetCustomXmlParts(ctx, request)
}



// func Test_Examples_GetCustomXmlPartsOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetDocument(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetDocumentRequest{
        DocumentName: ToStringPointer("Sample.docx"),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetDocument(ctx, request)
}



func Test_Examples_GetDocumentDrawingObjectByIndex(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetDocumentDrawingObjectByIndexRequest{
        Name: ToStringPointer("Sample.docx"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetDocumentDrawingObjectByIndex(ctx, request)
}



// func Test_Examples_GetDocumentDrawingObjectByIndexOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetDocumentDrawingObjectImageData(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetDocumentDrawingObjectImageDataRequest{
        Name: ToStringPointer("Sample.docx"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: requestOptions,
    }
    _, _ = wordsApi.GetDocumentDrawingObjectImageData(ctx, request)
}



// func Test_Examples_GetDocumentDrawingObjectImageDataOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetDocumentDrawingObjectOleData(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetDocumentDrawingObjectOleDataRequest{
        Name: ToStringPointer("Sample.docx"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: requestOptions,
    }
    _, _ = wordsApi.GetDocumentDrawingObjectOleData(ctx, request)
}



// func Test_Examples_GetDocumentDrawingObjectOleDataOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetDocumentDrawingObjects(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetDocumentDrawingObjectsRequest{
        Name: ToStringPointer("Sample.docx"),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetDocumentDrawingObjects(ctx, request)
}



// func Test_Examples_GetDocumentDrawingObjectsOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetDocumentFieldNames(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetDocumentFieldNamesRequest{
        Name: ToStringPointer("Sample.docx"),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetDocumentFieldNames(ctx, request)
}



// func Test_Examples_GetDocumentFieldNamesOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetDocumentHyperlinkByIndex(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetDocumentHyperlinkByIndexRequest{
        Name: ToStringPointer("Sample.docx"),
        HyperlinkIndex: ToInt32Pointer(int32(0)),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetDocumentHyperlinkByIndex(ctx, request)
}



// func Test_Examples_GetDocumentHyperlinkByIndexOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetDocumentHyperlinks(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetDocumentHyperlinksRequest{
        Name: ToStringPointer("Sample.docx"),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetDocumentHyperlinks(ctx, request)
}



// func Test_Examples_GetDocumentHyperlinksOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetDocumentProperties(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetDocumentPropertiesRequest{
        Name: ToStringPointer("Sample.docx"),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetDocumentProperties(ctx, request)
}



// func Test_Examples_GetDocumentPropertiesOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetDocumentProperty(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetDocumentPropertyRequest{
        Name: ToStringPointer("Sample.docx"),
        PropertyName: ToStringPointer("Author"),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetDocumentProperty(ctx, request)
}



// func Test_Examples_GetDocumentPropertyOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetDocumentProtection(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetDocumentProtectionRequest{
        Name: ToStringPointer("Sample.docx"),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetDocumentProtection(ctx, request)
}



// func Test_Examples_GetDocumentProtectionOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetDocumentStatistics(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetDocumentStatisticsRequest{
        Name: ToStringPointer("Sample.docx"),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetDocumentStatistics(ctx, request)
}



// func Test_Examples_GetDocumentStatisticsOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetDocumentWithFormat(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{"outPath": "/TestGetDocumentWithFormatAndOutPath.text",}
    request := &models.GetDocumentWithFormatRequest{
        Name: ToStringPointer("Sample.docx"),
        Format: ToStringPointer("text"),
        Optionals: requestOptions,
    }
    _, _ = wordsApi.GetDocumentWithFormat(ctx, request)
}



func Test_Examples_GetField(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetFieldRequest{
        Name: ToStringPointer("Sample.docx"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetField(ctx, request)
}



// func Test_Examples_GetFieldOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetFields(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetFieldsRequest{
        Name: ToStringPointer("Sample.docx"),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetFields(ctx, request)
}



// func Test_Examples_GetFieldsOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetFilesList(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetFilesListRequest{
        Path: ToStringPointer(""),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetFilesList(ctx, request)
}



func Test_Examples_GetFootnote(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetFootnoteRequest{
        Name: ToStringPointer("Sample.docx"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetFootnote(ctx, request)
}



// func Test_Examples_GetFootnoteOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetFootnotes(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetFootnotesRequest{
        Name: ToStringPointer("Sample.docx"),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetFootnotes(ctx, request)
}



// func Test_Examples_GetFootnotesOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetFormField(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetFormFieldRequest{
        Name: ToStringPointer("Sample.docx"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetFormField(ctx, request)
}



// func Test_Examples_GetFormFieldOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetFormFields(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetFormFieldsRequest{
        Name: ToStringPointer("Sample.docx"),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetFormFields(ctx, request)
}



// func Test_Examples_GetFormFieldsOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetHeaderFooter(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetHeaderFooterRequest{
        Name: ToStringPointer("Sample.docx"),
        HeaderFooterIndex: ToInt32Pointer(int32(0)),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetHeaderFooter(ctx, request)
}



func Test_Examples_GetHeaderFooterOfSection(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetHeaderFooterOfSectionRequest{
        Name: ToStringPointer("Sample.docx"),
        HeaderFooterIndex: ToInt32Pointer(int32(0)),
        SectionIndex: ToInt32Pointer(int32(0)),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetHeaderFooterOfSection(ctx, request)
}



// func Test_Examples_GetHeaderFooterOfSectionOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



// func Test_Examples_GetHeaderFooterOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetHeaderFooters(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetHeaderFootersRequest{
        Name: ToStringPointer("Sample.docx"),
        SectionPath: ToStringPointer(""),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetHeaderFooters(ctx, request)
}



// func Test_Examples_GetHeaderFootersOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetList(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetListRequest{
        Name: ToStringPointer("TestGetLists.doc"),
        ListId: ToInt32Pointer(int32(1)),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetList(ctx, request)
}



// func Test_Examples_GetListOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetLists(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetListsRequest{
        Name: ToStringPointer("TestGetLists.doc"),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetLists(ctx, request)
}



// func Test_Examples_GetListsOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetOfficeMathObject(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetOfficeMathObjectRequest{
        Name: ToStringPointer("Sample.docx"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetOfficeMathObject(ctx, request)
}



// func Test_Examples_GetOfficeMathObjectOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetOfficeMathObjects(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetOfficeMathObjectsRequest{
        Name: ToStringPointer("Sample.docx"),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetOfficeMathObjects(ctx, request)
}



// func Test_Examples_GetOfficeMathObjectsOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetParagraph(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetParagraphRequest{
        Name: ToStringPointer("Sample.docx"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetParagraph(ctx, request)
}



func Test_Examples_GetParagraphFormat(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetParagraphFormatRequest{
        Name: ToStringPointer("Sample.docx"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetParagraphFormat(ctx, request)
}



// func Test_Examples_GetParagraphFormatOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetParagraphListFormat(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetParagraphListFormatRequest{
        Name: ToStringPointer("Sample.docx"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetParagraphListFormat(ctx, request)
}



// func Test_Examples_GetParagraphListFormatOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



// func Test_Examples_GetParagraphOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetParagraphs(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetParagraphsRequest{
        Name: ToStringPointer("Sample.docx"),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetParagraphs(ctx, request)
}



// func Test_Examples_GetParagraphsOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetParagraphTabStops(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetParagraphTabStopsRequest{
        Name: ToStringPointer("Sample.docx"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetParagraphTabStops(ctx, request)
}



// func Test_Examples_GetParagraphTabStopsOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetPublicKey(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    request := &models.GetPublicKeyRequest{
    }
    _, _, _ = wordsApi.GetPublicKey(ctx, request)
}



func Test_Examples_GetRangeText(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{"rangeEndIdentifier": "id0.0.1",}
    request := &models.GetRangeTextRequest{
        Name: ToStringPointer("Sample.docx"),
        RangeStartIdentifier: ToStringPointer("id0.0.0"),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetRangeText(ctx, request)
}



// func Test_Examples_GetRangeTextOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetRun(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetRunRequest{
        Name: ToStringPointer("Sample.docx"),
        ParagraphPath: ToStringPointer("paragraphs/0"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetRun(ctx, request)
}



func Test_Examples_GetRunFont(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetRunFontRequest{
        Name: ToStringPointer("Sample.docx"),
        ParagraphPath: ToStringPointer("paragraphs/0"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetRunFont(ctx, request)
}



// func Test_Examples_GetRunFontOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



// func Test_Examples_GetRunOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetRuns(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetRunsRequest{
        Name: ToStringPointer("Sample.docx"),
        ParagraphPath: ToStringPointer("sections/0/paragraphs/0"),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetRuns(ctx, request)
}



// func Test_Examples_GetRunsOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetSection(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetSectionRequest{
        Name: ToStringPointer("Sample.docx"),
        SectionIndex: ToInt32Pointer(int32(0)),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetSection(ctx, request)
}



// func Test_Examples_GetSectionOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetSectionPageSetup(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetSectionPageSetupRequest{
        Name: ToStringPointer("Sample.docx"),
        SectionIndex: ToInt32Pointer(int32(0)),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetSectionPageSetup(ctx, request)
}



// func Test_Examples_GetSectionPageSetupOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetSections(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetSectionsRequest{
        Name: ToStringPointer("Sample.docx"),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetSections(ctx, request)
}



// func Test_Examples_GetSectionsOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetStyle(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetStyleRequest{
        Name: ToStringPointer("Sample.docx"),
        StyleName: ToStringPointer("Heading 1"),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetStyle(ctx, request)
}



func Test_Examples_GetStyleFromDocumentElement(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetStyleFromDocumentElementRequest{
        Name: ToStringPointer("Sample.docx"),
        StyledNodePath: ToStringPointer("paragraphs/1/paragraphFormat"),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetStyleFromDocumentElement(ctx, request)
}



// func Test_Examples_GetStyleFromDocumentElementOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



// func Test_Examples_GetStyleOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetStyles(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetStylesRequest{
        Name: ToStringPointer("Sample.docx"),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetStyles(ctx, request)
}



// func Test_Examples_GetStylesOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetTable(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetTableRequest{
        Name: ToStringPointer("Sample.docx"),
        Index: ToInt32Pointer(int32(1)),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetTable(ctx, request)
}



func Test_Examples_GetTableCell(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetTableCellRequest{
        Name: ToStringPointer("Sample.docx"),
        TableRowPath: ToStringPointer("sections/0/tables/2/rows/0"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetTableCell(ctx, request)
}



func Test_Examples_GetTableCellFormat(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetTableCellFormatRequest{
        Name: ToStringPointer("Sample.docx"),
        TableRowPath: ToStringPointer("sections/0/tables/2/rows/0"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetTableCellFormat(ctx, request)
}



// func Test_Examples_GetTableCellFormatOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



// func Test_Examples_GetTableCellOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



// func Test_Examples_GetTableOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetTableProperties(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetTablePropertiesRequest{
        Name: ToStringPointer("Sample.docx"),
        Index: ToInt32Pointer(int32(1)),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetTableProperties(ctx, request)
}



// func Test_Examples_GetTablePropertiesOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetTableRow(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetTableRowRequest{
        Name: ToStringPointer("Sample.docx"),
        TablePath: ToStringPointer("tables/1"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetTableRow(ctx, request)
}



func Test_Examples_GetTableRowFormat(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetTableRowFormatRequest{
        Name: ToStringPointer("Sample.docx"),
        TablePath: ToStringPointer("sections/0/tables/2"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetTableRowFormat(ctx, request)
}



// func Test_Examples_GetTableRowFormatOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



// func Test_Examples_GetTableRowOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_GetTables(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := map[string]interface{}{}
    request := &models.GetTablesRequest{
        Name: ToStringPointer("Sample.docx"),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.GetTables(ctx, request)
}



// func Test_Examples_GetTablesOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_InsertComment(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestCommentRangeStartNode := models.NodeLink{
        NodeId: ToStringPointer("0.3.0.3"),
    }
    requestCommentRangeStart := models.DocumentPosition{
        Node: requestCommentRangeStartNode,
        Offset: ToInt32Pointer(int32(0)),
    }
    requestCommentRangeEndNode := models.NodeLink{
        NodeId: ToStringPointer("0.3.0.3"),
    }
    requestCommentRangeEnd := models.DocumentPosition{
        Node: requestCommentRangeEndNode,
        Offset: ToInt32Pointer(int32(0)),
    }
    requestComment := models.CommentInsert{
        RangeStart: requestCommentRangeStart,
        RangeEnd: requestCommentRangeEnd,
        Initial: ToStringPointer("IA"),
        Author: ToStringPointer("Imran Anwar"),
        Text: ToStringPointer("A new Comment"),
    }
    insertRequestOptions := map[string]interface{}{}
    insertRequest := &models.InsertCommentRequest{
        Name: ToStringPointer("Sample.docx"),
        Comment: requestComment,
        Optionals: insertRequestOptions,
    }
    _, _, _ = wordsApi.InsertComment(ctx, insertRequest)
}



// func Test_Examples_InsertCommentOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_InsertCustomXmlPart(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestCustomXmlPart := models.CustomXmlPartInsert{
        Id: ToStringPointer("hello"),
        Data: ToStringPointer("<data>Hello world</data>"),
    }
    insertRequestOptions := map[string]interface{}{}
    insertRequest := &models.InsertCustomXmlPartRequest{
        Name: ToStringPointer("Sample.docx"),
        CustomXmlPart: requestCustomXmlPart,
        Optionals: insertRequestOptions,
    }
    _, _, _ = wordsApi.InsertCustomXmlPart(ctx, insertRequest)
}



// func Test_Examples_InsertCustomXmlPartOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_InsertDrawingObject(t *testing.T) {
    documentsDir := GetExamplesDir()
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestImageFileFileData, _ := os.Open(documentsDir + "/" + "Common/aspose-cloud.png")
    requestDrawingObject := models.DrawingObjectInsert{
        Height: ToFloat64Pointer(0),
        Left: ToFloat64Pointer(0),
        Top: ToFloat64Pointer(0),
        Width: ToFloat64Pointer(0),
        RelativeHorizontalPosition: ToStringPointer("Margin"),
        RelativeVerticalPosition: ToStringPointer("Margin"),
        WrapType: ToStringPointer("Inline"),
    }
    insertRequestOptions := map[string]interface{}{}
    insertRequest := &models.InsertDrawingObjectRequest{
        Name: ToStringPointer("Sample.docx"),
        DrawingObject: requestDrawingObject,
        ImageFile: requestImageFileFileData,
        Optionals: insertRequestOptions,
    }
    _, _, _ = wordsApi.InsertDrawingObject(ctx, insertRequest)
}



// func Test_Examples_InsertDrawingObjectOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_InsertField(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestField := models.FieldInsert{
        FieldCode: ToStringPointer("{ NUMPAGES }"),
    }
    insertRequestOptions := map[string]interface{}{}
    insertRequest := &models.InsertFieldRequest{
        Name: ToStringPointer("Sample.docx"),
        Field: requestField,
        Optionals: insertRequestOptions,
    }
    _, _, _ = wordsApi.InsertField(ctx, insertRequest)
}



// func Test_Examples_InsertFieldOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_InsertFootnote(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestFootnoteDto := models.FootnoteInsert{
        FootnoteType: ToStringPointer("Endnote"),
        Text: ToStringPointer("test endnote"),
    }
    insertRequestOptions := map[string]interface{}{}
    insertRequest := &models.InsertFootnoteRequest{
        Name: ToStringPointer("Sample.docx"),
        FootnoteDto: requestFootnoteDto,
        Optionals: insertRequestOptions,
    }
    _, _, _ = wordsApi.InsertFootnote(ctx, insertRequest)
}



// func Test_Examples_InsertFootnoteOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_InsertFormField(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    remoteFileName:= "Sample.docx"

    requestFormField := models.FormFieldTextInput{
        Name: ToStringPointer("FullName"),
        Enabled: ToBoolPointer(true),
        CalculateOnExit: ToBoolPointer(true),
        StatusText: ToStringPointer(""),
        TextInputType: ToStringPointer("Regular"),
        TextInputDefault: ToStringPointer("123"),
        TextInputFormat: ToStringPointer("UPPERCASE"),
    }
    insertRequestOptions := map[string]interface{}{"destFileName": remoteFileName,}
    insertRequest := &models.InsertFormFieldRequest{
        Name: ToStringPointer(remoteFileName),
        FormField: requestFormField,
        Optionals: insertRequestOptions,
    }
    _, _, _ = wordsApi.InsertFormField(ctx, insertRequest)
}



// func Test_Examples_InsertFormFieldOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_InsertHeaderFooter(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    insertRequestOptions := map[string]interface{}{}
    insertRequest := &models.InsertHeaderFooterRequest{
        Name: ToStringPointer("Sample.docx"),
        SectionPath: ToStringPointer(""),
        HeaderFooterType: ToStringPointer("FooterEven"),
        Optionals: insertRequestOptions,
    }
    _, _, _ = wordsApi.InsertHeaderFooter(ctx, insertRequest)
}



// func Test_Examples_InsertHeaderFooterOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_InsertList(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestListInsert := models.ListInsert{
        Template: ToStringPointer("OutlineLegal"),
    }
    insertRequestOptions := map[string]interface{}{}
    insertRequest := &models.InsertListRequest{
        Name: ToStringPointer("TestGetLists.doc"),
        ListInsert: requestListInsert,
        Optionals: insertRequestOptions,
    }
    _, _, _ = wordsApi.InsertList(ctx, insertRequest)
}



// func Test_Examples_InsertListOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_InsertOrUpdateParagraphTabStop(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestTabStopInsertDto := models.TabStopInsert{
        Alignment: ToStringPointer("Left"),
        Leader: ToStringPointer("None"),
        Position: ToFloat64Pointer(100.0),
    }
    insertRequestOptions := map[string]interface{}{}
    insertRequest := &models.InsertOrUpdateParagraphTabStopRequest{
        Name: ToStringPointer("Sample.docx"),
        Index: ToInt32Pointer(int32(0)),
        TabStopInsertDto: requestTabStopInsertDto,
        Optionals: insertRequestOptions,
    }
    _, _, _ = wordsApi.InsertOrUpdateParagraphTabStop(ctx, insertRequest)
}



// func Test_Examples_InsertOrUpdateParagraphTabStopOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_InsertPageNumbers(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    remoteFileName:= "Sample.docx"

    requestPageNumber := models.PageNumber{
        Alignment: ToStringPointer("center"),
        Format: ToStringPointer("{PAGE} of {NUMPAGES}"),
    }
    insertRequestOptions := map[string]interface{}{"destFileName": remoteFileName,}
    insertRequest := &models.InsertPageNumbersRequest{
        Name: ToStringPointer(remoteFileName),
        PageNumber: requestPageNumber,
        Optionals: insertRequestOptions,
    }
    _, _, _ = wordsApi.InsertPageNumbers(ctx, insertRequest)
}



// func Test_Examples_InsertPageNumbersOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_InsertParagraph(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestParagraph := models.ParagraphInsert{
        Text: ToStringPointer("This is a new paragraph for your document"),
    }
    insertRequestOptions := map[string]interface{}{}
    insertRequest := &models.InsertParagraphRequest{
        Name: ToStringPointer("Sample.docx"),
        Paragraph: requestParagraph,
        Optionals: insertRequestOptions,
    }
    _, _, _ = wordsApi.InsertParagraph(ctx, insertRequest)
}



// func Test_Examples_InsertParagraphOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_InsertRun(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestRun := models.RunInsert{
        Text: ToStringPointer("run with text"),
    }
    insertRequestOptions := map[string]interface{}{}
    insertRequest := &models.InsertRunRequest{
        Name: ToStringPointer("Sample.docx"),
        ParagraphPath: ToStringPointer("paragraphs/1"),
        Run: requestRun,
        Optionals: insertRequestOptions,
    }
    _, _, _ = wordsApi.InsertRun(ctx, insertRequest)
}



// func Test_Examples_InsertRunOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_InsertStyle(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestStyleInsert := models.StyleInsert{
        StyleName: ToStringPointer("My Style"),
        StyleType: ToStringPointer("Paragraph"),
    }
    insertRequestOptions := map[string]interface{}{}
    insertRequest := &models.InsertStyleRequest{
        Name: ToStringPointer("Sample.docx"),
        StyleInsert: requestStyleInsert,
        Optionals: insertRequestOptions,
    }
    _, _, _ = wordsApi.InsertStyle(ctx, insertRequest)
}



// func Test_Examples_InsertStyleOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_InsertTable(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestTable := models.TableInsert{
        ColumnsCount: ToInt32Pointer(int32(5)),
        RowsCount: ToInt32Pointer(int32(4)),
    }
    insertRequestOptions := map[string]interface{}{}
    insertRequest := &models.InsertTableRequest{
        Name: ToStringPointer("Sample.docx"),
        Table: requestTable,
        Optionals: insertRequestOptions,
    }
    _, _, _ = wordsApi.InsertTable(ctx, insertRequest)
}



func Test_Examples_InsertTableCell(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestCell := models.TableCellInsert{
    }
    insertRequestOptions := map[string]interface{}{}
    insertRequest := &models.InsertTableCellRequest{
        Name: ToStringPointer("Sample.docx"),
        TableRowPath: ToStringPointer("sections/0/tables/2/rows/0"),
        Cell: requestCell,
        Optionals: insertRequestOptions,
    }
    _, _, _ = wordsApi.InsertTableCell(ctx, insertRequest)
}



// func Test_Examples_InsertTableCellOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



// func Test_Examples_InsertTableOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_InsertTableRow(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestRow := models.TableRowInsert{
        ColumnsCount: ToInt32Pointer(int32(5)),
    }
    insertRequestOptions := map[string]interface{}{}
    insertRequest := &models.InsertTableRowRequest{
        Name: ToStringPointer("Sample.docx"),
        TablePath: ToStringPointer("sections/0/tables/2"),
        Row: requestRow,
        Optionals: insertRequestOptions,
    }
    _, _, _ = wordsApi.InsertTableRow(ctx, insertRequest)
}



// func Test_Examples_InsertTableRowOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_InsertWatermarkImage(t *testing.T) {
    documentsDir := GetExamplesDir()
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    remoteFileName:= "Sample.docx"

    insertRequestOptions := map[string]interface{}{"imageFile": nil,
    "destFileName": remoteFileName,
    "image": "Sample.png",}
    insertRequest := &models.InsertWatermarkImageRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: insertRequestOptions,
    }
    _, _, _ = wordsApi.InsertWatermarkImage(ctx, insertRequest)
}



// func Test_Examples_InsertWatermarkImageOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_InsertWatermarkText(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    remoteFileName:= "Sample.docx"

    requestWatermarkText := models.WatermarkText{
        Text: ToStringPointer("This is the text"),
        RotationAngle: ToFloat64Pointer(90.0),
    }
    insertRequestOptions := map[string]interface{}{"destFileName": remoteFileName,}
    insertRequest := &models.InsertWatermarkTextRequest{
        Name: ToStringPointer(remoteFileName),
        WatermarkText: requestWatermarkText,
        Optionals: insertRequestOptions,
    }
    _, _, _ = wordsApi.InsertWatermarkText(ctx, insertRequest)
}



// func Test_Examples_InsertWatermarkTextOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_LoadWebDocument(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestDataSaveOptions := models.SaveOptionsData{
        FileName: ToStringPointer("google.doc"),
        SaveFormat: ToStringPointer("doc"),
        DmlEffectsRenderingMode: ToStringPointer("1"),
        DmlRenderingMode: ToStringPointer("1"),
        UpdateSdtContent: ToBoolPointer(false),
        ZipOutput: ToBoolPointer(false),
    }
    requestData := models.LoadWebDocumentData{
        LoadingDocumentUrl: ToStringPointer("http://google.com"),
        SaveOptions: requestDataSaveOptions,
    }
    loadRequestOptions := map[string]interface{}{}
    loadRequest := &models.LoadWebDocumentRequest{
        Data: requestData,
        Optionals: loadRequestOptions,
    }
    _, _, _ = wordsApi.LoadWebDocument(ctx, loadRequest)
}



func Test_Examples_MoveFile(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    moveRequestOptions := map[string]interface{}{}
    moveRequest := &models.MoveFileRequest{
        DestPath: ToStringPointer("/TestMoveFileDest_Sample.docx"),
        SrcPath: ToStringPointer("Sample.docx"),
        Optionals: moveRequestOptions,
    }
    _, _ = wordsApi.MoveFile(ctx, moveRequest)
}



func Test_Examples_MoveFolder(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    moveRequestOptions := map[string]interface{}{}
    moveRequest := &models.MoveFolderRequest{
        DestPath: ToStringPointer("/TestMoveFolderDest_Sample"),
        SrcPath: ToStringPointer("/TestMoveFolderSrc"),
        Optionals: moveRequestOptions,
    }
    _, _ = wordsApi.MoveFolder(ctx, moveRequest)
}



func Test_Examples_OptimizeDocument(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestOptions := models.OptimizationOptions{
        MsWordVersion: ToStringPointer("Word2002"),
    }
    optimizeRequestOptions := map[string]interface{}{}
    optimizeRequest := &models.OptimizeDocumentRequest{
        Name: ToStringPointer("Sample.docx"),
        Options: requestOptions,
        Optionals: optimizeRequestOptions,
    }
    _, _ = wordsApi.OptimizeDocument(ctx, optimizeRequest)
}



// func Test_Examples_OptimizeDocumentOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_ProtectDocument(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    remoteFileName:= "Sample.docx"

    requestProtectionRequest := models.ProtectionRequest{
        Password: ToStringPointer("123"),
        ProtectionType: ToStringPointer("ReadOnly"),
    }
    protectRequestOptions := map[string]interface{}{"destFileName": remoteFileName,}
    protectRequest := &models.ProtectDocumentRequest{
        Name: ToStringPointer(remoteFileName),
        ProtectionRequest: requestProtectionRequest,
        Optionals: protectRequestOptions,
    }
    _, _, _ = wordsApi.ProtectDocument(ctx, protectRequest)
}



// func Test_Examples_ProtectDocumentOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_RejectAllRevisions(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    remoteFileName:= "Sample.docx"

    rejectRequestOptions := map[string]interface{}{"destFileName": remoteFileName,}
    rejectRequest := &models.RejectAllRevisionsRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: rejectRequestOptions,
    }
    _, _, _ = wordsApi.RejectAllRevisions(ctx, rejectRequest)
}



// func Test_Examples_RejectAllRevisionsOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_RemoveRange(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    removeRequestOptions := map[string]interface{}{"rangeEndIdentifier": "id0.0.1",}
    removeRequest := &models.RemoveRangeRequest{
        Name: ToStringPointer("Sample.docx"),
        RangeStartIdentifier: ToStringPointer("id0.0.0"),
        Optionals: removeRequestOptions,
    }
    _, _, _ = wordsApi.RemoveRange(ctx, removeRequest)
}



// func Test_Examples_RemoveRangeOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_RenderDrawingObject(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    renderRequestOptions := map[string]interface{}{}
    renderRequest := &models.RenderDrawingObjectRequest{
        Name: ToStringPointer("Sample.docx"),
        Format: ToStringPointer("png"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: renderRequestOptions,
    }
    _, _ = wordsApi.RenderDrawingObject(ctx, renderRequest)
}



// func Test_Examples_RenderDrawingObjectOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_RenderMathObject(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    renderRequestOptions := map[string]interface{}{}
    renderRequest := &models.RenderMathObjectRequest{
        Name: ToStringPointer("Sample.docx"),
        Format: ToStringPointer("png"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: renderRequestOptions,
    }
    _, _ = wordsApi.RenderMathObject(ctx, renderRequest)
}



// func Test_Examples_RenderMathObjectOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_RenderPage(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    renderRequestOptions := map[string]interface{}{}
    renderRequest := &models.RenderPageRequest{
        Name: ToStringPointer("Sample.docx"),
        PageIndex: ToInt32Pointer(int32(1)),
        Format: ToStringPointer("bmp"),
        Optionals: renderRequestOptions,
    }
    _, _ = wordsApi.RenderPage(ctx, renderRequest)
}



// func Test_Examples_RenderPageOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_RenderParagraph(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    renderRequestOptions := map[string]interface{}{}
    renderRequest := &models.RenderParagraphRequest{
        Name: ToStringPointer("Sample.docx"),
        Format: ToStringPointer("png"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: renderRequestOptions,
    }
    _, _ = wordsApi.RenderParagraph(ctx, renderRequest)
}



// func Test_Examples_RenderParagraphOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_RenderTable(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    renderRequestOptions := map[string]interface{}{}
    renderRequest := &models.RenderTableRequest{
        Name: ToStringPointer("Sample.docx"),
        Format: ToStringPointer("png"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: renderRequestOptions,
    }
    _, _ = wordsApi.RenderTable(ctx, renderRequest)
}



// func Test_Examples_RenderTableOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_ReplaceText(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    remoteFileName:= "Sample.docx"

    requestReplaceText := models.ReplaceTextParameters{
        OldValue: ToStringPointer("Testing"),
        NewValue: ToStringPointer("Aspose testing"),
    }
    replaceRequestOptions := map[string]interface{}{"destFileName": remoteFileName,}
    replaceRequest := &models.ReplaceTextRequest{
        Name: ToStringPointer(remoteFileName),
        ReplaceText: requestReplaceText,
        Optionals: replaceRequestOptions,
    }
    _, _, _ = wordsApi.ReplaceText(ctx, replaceRequest)
}



// func Test_Examples_ReplaceTextOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_ReplaceWithText(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestRangeText := models.ReplaceRange{
        Text: ToStringPointer("Replaced header"),
    }
    replaceRequestOptions := map[string]interface{}{"rangeEndIdentifier": "id0.0.1",}
    replaceRequest := &models.ReplaceWithTextRequest{
        Name: ToStringPointer("Sample.docx"),
        RangeStartIdentifier: ToStringPointer("id0.0.0"),
        RangeText: requestRangeText,
        Optionals: replaceRequestOptions,
    }
    _, _, _ = wordsApi.ReplaceWithText(ctx, replaceRequest)
}



// func Test_Examples_ReplaceWithTextOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_ResetCache(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    resetRequest := &models.ResetCacheRequest{
    }
    _, _ = wordsApi.ResetCache(ctx, resetRequest)
}



func Test_Examples_SaveAs(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestSaveOptionsData := models.SaveOptionsData{
        SaveFormat: ToStringPointer("docx"),
        FileName: ToStringPointer("/TestSaveAsFromPdfToDoc.docx"),
    }
    saveRequestOptions := map[string]interface{}{}
    saveRequest := &models.SaveAsRequest{
        Name: ToStringPointer("Sample.docx"),
        SaveOptionsData: requestSaveOptionsData,
        Optionals: saveRequestOptions,
    }
    _, _, _ = wordsApi.SaveAs(ctx, saveRequest)
}



// func Test_Examples_SaveAsOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_SaveAsRange(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestDocumentParameters := models.RangeDocument{
        DocumentName: ToStringPointer("/NewDoc.docx"),
    }
    saveRequestOptions := map[string]interface{}{"rangeEndIdentifier": "id0.0.1",}
    saveRequest := &models.SaveAsRangeRequest{
        Name: ToStringPointer("Sample.docx"),
        RangeStartIdentifier: ToStringPointer("id0.0.0"),
        DocumentParameters: requestDocumentParameters,
        Optionals: saveRequestOptions,
    }
    _, _, _ = wordsApi.SaveAsRange(ctx, saveRequest)
}



// func Test_Examples_SaveAsRangeOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_SaveAsTiff(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestSaveOptions := models.TiffSaveOptionsData{
        SaveFormat: ToStringPointer("tiff"),
        FileName: ToStringPointer("/abc.tiff"),
    }
    saveRequestOptions := map[string]interface{}{}
    saveRequest := &models.SaveAsTiffRequest{
        Name: ToStringPointer("Sample.docx"),
        SaveOptions: requestSaveOptions,
        Optionals: saveRequestOptions,
    }
    _, _, _ = wordsApi.SaveAsTiff(ctx, saveRequest)
}



// func Test_Examples_SaveAsTiffOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_Search(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    searchRequestOptions := map[string]interface{}{}
    searchRequest := &models.SearchRequest{
        Name: ToStringPointer("Sample.docx"),
        Pattern: ToStringPointer("aspose"),
        Optionals: searchRequestOptions,
    }
    _, _, _ = wordsApi.Search(ctx, searchRequest)
}



// func Test_Examples_SearchOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_SplitDocument(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    splitRequestOptions := map[string]interface{}{"destFileName": "/TestSplitDocument.text",
    "from": int32(1),
    "to": int32(2),}
    splitRequest := &models.SplitDocumentRequest{
        Name: ToStringPointer("Sample.docx"),
        Format: ToStringPointer("text"),
        Optionals: splitRequestOptions,
    }
    _, _, _ = wordsApi.SplitDocument(ctx, splitRequest)
}



// func Test_Examples_SplitDocumentOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_UnprotectDocument(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestProtectionRequest := models.ProtectionRequest{
        Password: ToStringPointer("aspose"),
    }
    unprotectRequestOptions := map[string]interface{}{}
    unprotectRequest := &models.UnprotectDocumentRequest{
        Name: ToStringPointer("Sample.docx"),
        ProtectionRequest: requestProtectionRequest,
        Optionals: unprotectRequestOptions,
    }
    _, _, _ = wordsApi.UnprotectDocument(ctx, unprotectRequest)
}



// func Test_Examples_UnprotectDocumentOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_UpdateBookmark(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    bookmarkName:= "aspose"
    remoteFileName:= "Sample.docx"

    requestBookmarkData := models.BookmarkData{
        Name: ToStringPointer(bookmarkName),
        Text: ToStringPointer("This will be the text for Aspose"),
    }
    updateRequestOptions := map[string]interface{}{"destFileName": remoteFileName,}
    updateRequest := &models.UpdateBookmarkRequest{
        Name: ToStringPointer(remoteFileName),
        BookmarkName: ToStringPointer(bookmarkName),
        BookmarkData: requestBookmarkData,
        Optionals: updateRequestOptions,
    }
    _, _, _ = wordsApi.UpdateBookmark(ctx, updateRequest)
}



// func Test_Examples_UpdateBookmarkOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_UpdateBorder(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestBorderPropertiesColor := models.XmlColor{
        Web: ToStringPointer("#AABBCC"),
    }
    requestBorderProperties := models.Border{
        BorderType: ToStringPointer("Left"),
        Color: requestBorderPropertiesColor,
        DistanceFromText: ToFloat64Pointer(6.0),
        LineStyle: ToStringPointer("DashDotStroker"),
        LineWidth: ToFloat64Pointer(2.0),
        Shadow: ToBoolPointer(true),
    }
    updateRequestOptions := map[string]interface{}{"nodePath": "tables/1/rows/0/cells/0",}
    updateRequest := &models.UpdateBorderRequest{
        Name: ToStringPointer("Sample.docx"),
        BorderType: ToStringPointer("left"),
        BorderProperties: requestBorderProperties,
        Optionals: updateRequestOptions,
    }
    _, _, _ = wordsApi.UpdateBorder(ctx, updateRequest)
}



// func Test_Examples_UpdateBorderOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_UpdateComment(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestCommentRangeStartNode := models.NodeLink{
        NodeId: ToStringPointer("0.3.0"),
    }
    requestCommentRangeStart := models.DocumentPosition{
        Node: requestCommentRangeStartNode,
        Offset: ToInt32Pointer(int32(0)),
    }
    requestCommentRangeEndNode := models.NodeLink{
        NodeId: ToStringPointer("0.3.0"),
    }
    requestCommentRangeEnd := models.DocumentPosition{
        Node: requestCommentRangeEndNode,
        Offset: ToInt32Pointer(int32(0)),
    }
    requestComment := models.CommentUpdate{
        RangeStart: requestCommentRangeStart,
        RangeEnd: requestCommentRangeEnd,
        Initial: ToStringPointer("IA"),
        Author: ToStringPointer("Imran Anwar"),
        Text: ToStringPointer("A new Comment"),
    }
    updateRequestOptions := map[string]interface{}{}
    updateRequest := &models.UpdateCommentRequest{
        Name: ToStringPointer("Sample.docx"),
        CommentIndex: ToInt32Pointer(int32(0)),
        Comment: requestComment,
        Optionals: updateRequestOptions,
    }
    _, _, _ = wordsApi.UpdateComment(ctx, updateRequest)
}



// func Test_Examples_UpdateCommentOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_UpdateCustomXmlPart(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestCustomXmlPart := models.CustomXmlPartUpdate{
        Data: ToStringPointer("<data>Hello world</data>"),
    }
    updateRequestOptions := map[string]interface{}{}
    updateRequest := &models.UpdateCustomXmlPartRequest{
        Name: ToStringPointer("Sample.docx"),
        CustomXmlPartIndex: ToInt32Pointer(int32(0)),
        CustomXmlPart: requestCustomXmlPart,
        Optionals: updateRequestOptions,
    }
    _, _, _ = wordsApi.UpdateCustomXmlPart(ctx, updateRequest)
}



// func Test_Examples_UpdateCustomXmlPartOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_UpdateDrawingObject(t *testing.T) {
    documentsDir := GetExamplesDir()
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestImageFileFileData, _ := os.Open(documentsDir + "/" + "Common/aspose-cloud.png")
    requestDrawingObject := models.DrawingObjectUpdate{
        Left: ToFloat64Pointer(0),
    }
    updateRequestOptions := map[string]interface{}{}
    updateRequest := &models.UpdateDrawingObjectRequest{
        Name: ToStringPointer("Sample.docx"),
        DrawingObject: requestDrawingObject,
        ImageFile: requestImageFileFileData,
        Index: ToInt32Pointer(int32(0)),
        Optionals: updateRequestOptions,
    }
    _, _, _ = wordsApi.UpdateDrawingObject(ctx, updateRequest)
}



// func Test_Examples_UpdateDrawingObjectOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_UpdateField(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestField := models.FieldUpdate{
        FieldCode: ToStringPointer("{ NUMPAGES }"),
    }
    updateRequestOptions := map[string]interface{}{"nodePath": "sections/0/paragraphs/0",}
    updateRequest := &models.UpdateFieldRequest{
        Name: ToStringPointer("Sample.docx"),
        Index: ToInt32Pointer(int32(0)),
        Field: requestField,
        Optionals: updateRequestOptions,
    }
    _, _, _ = wordsApi.UpdateField(ctx, updateRequest)
}



// func Test_Examples_UpdateFieldOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_UpdateFields(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    updateRequestOptions := map[string]interface{}{}
    updateRequest := &models.UpdateFieldsRequest{
        Name: ToStringPointer("Sample.docx"),
        Optionals: updateRequestOptions,
    }
    _, _, _ = wordsApi.UpdateFields(ctx, updateRequest)
}



// func Test_Examples_UpdateFieldsOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_UpdateFootnote(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestFootnoteDto := models.FootnoteUpdate{
        Text: ToStringPointer("new text is here"),
    }
    updateRequestOptions := map[string]interface{}{}
    updateRequest := &models.UpdateFootnoteRequest{
        Name: ToStringPointer("Sample.docx"),
        Index: ToInt32Pointer(int32(0)),
        FootnoteDto: requestFootnoteDto,
        Optionals: updateRequestOptions,
    }
    _, _, _ = wordsApi.UpdateFootnote(ctx, updateRequest)
}



// func Test_Examples_UpdateFootnoteOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_UpdateFormField(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    remoteFileName:= "Sample.docx"

    requestFormField := models.FormFieldTextInput{
        Name: ToStringPointer("FullName"),
        Enabled: ToBoolPointer(true),
        CalculateOnExit: ToBoolPointer(true),
        StatusText: ToStringPointer(""),
        TextInputType: ToStringPointer("Regular"),
        TextInputDefault: ToStringPointer("No name"),
    }
    updateRequestOptions := map[string]interface{}{"destFileName": remoteFileName,}
    updateRequest := &models.UpdateFormFieldRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        FormField: requestFormField,
        Optionals: updateRequestOptions,
    }
    _, _, _ = wordsApi.UpdateFormField(ctx, updateRequest)
}



// func Test_Examples_UpdateFormFieldOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_UpdateList(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestListUpdate := models.ListUpdate{
        IsRestartAtEachSection: ToBoolPointer(true),
    }
    updateRequestOptions := map[string]interface{}{}
    updateRequest := &models.UpdateListRequest{
        Name: ToStringPointer("TestGetLists.doc"),
        ListId: ToInt32Pointer(int32(1)),
        ListUpdate: requestListUpdate,
        Optionals: updateRequestOptions,
    }
    _, _, _ = wordsApi.UpdateList(ctx, updateRequest)
}



func Test_Examples_UpdateListLevel(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestListUpdate := models.ListLevelUpdate{
        Alignment: ToStringPointer("Right"),
    }
    updateRequestOptions := map[string]interface{}{}
    updateRequest := &models.UpdateListLevelRequest{
        Name: ToStringPointer("TestGetLists.doc"),
        ListId: ToInt32Pointer(int32(1)),
        ListLevel: ToInt32Pointer(int32(1)),
        ListUpdate: requestListUpdate,
        Optionals: updateRequestOptions,
    }
    _, _, _ = wordsApi.UpdateListLevel(ctx, updateRequest)
}



// func Test_Examples_UpdateListLevelOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



// func Test_Examples_UpdateListOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_UpdateParagraphFormat(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestParagraphFormatDto := models.ParagraphFormatUpdate{
        Alignment: ToStringPointer("Right"),
    }
    updateRequestOptions := map[string]interface{}{}
    updateRequest := &models.UpdateParagraphFormatRequest{
        Name: ToStringPointer("Sample.docx"),
        Index: ToInt32Pointer(int32(0)),
        ParagraphFormatDto: requestParagraphFormatDto,
        Optionals: updateRequestOptions,
    }
    _, _, _ = wordsApi.UpdateParagraphFormat(ctx, updateRequest)
}



// func Test_Examples_UpdateParagraphFormatOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_UpdateParagraphListFormat(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestListFormatDto := models.ListFormatUpdate{
        ListId: ToInt32Pointer(int32(2)),
    }
    updateRequestOptions := map[string]interface{}{}
    updateRequest := &models.UpdateParagraphListFormatRequest{
        Name: ToStringPointer("Sample.docx"),
        Index: ToInt32Pointer(int32(0)),
        ListFormatDto: requestListFormatDto,
        Optionals: updateRequestOptions,
    }
    _, _, _ = wordsApi.UpdateParagraphListFormat(ctx, updateRequest)
}



// func Test_Examples_UpdateParagraphListFormatOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_UpdateRun(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestRun := models.RunUpdate{
        Text: ToStringPointer("run with text"),
    }
    updateRequestOptions := map[string]interface{}{}
    updateRequest := &models.UpdateRunRequest{
        Name: ToStringPointer("Sample.docx"),
        ParagraphPath: ToStringPointer("paragraphs/1"),
        Index: ToInt32Pointer(int32(0)),
        Run: requestRun,
        Optionals: updateRequestOptions,
    }
    _, _, _ = wordsApi.UpdateRun(ctx, updateRequest)
}



func Test_Examples_UpdateRunFont(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    remoteFileName:= "Sample.docx"

    requestFontDto := models.Font{
        Bold: ToBoolPointer(true),
    }
    updateRequestOptions := map[string]interface{}{"destFileName": remoteFileName,}
    updateRequest := &models.UpdateRunFontRequest{
        Name: ToStringPointer(remoteFileName),
        ParagraphPath: ToStringPointer("paragraphs/0"),
        Index: ToInt32Pointer(int32(0)),
        FontDto: requestFontDto,
        Optionals: updateRequestOptions,
    }
    _, _, _ = wordsApi.UpdateRunFont(ctx, updateRequest)
}



// func Test_Examples_UpdateRunFontOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



// func Test_Examples_UpdateRunOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_UpdateSectionPageSetup(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestPageSetup := models.PageSetup{
        RtlGutter: ToBoolPointer(true),
        LeftMargin: ToFloat64Pointer(10.0),
        Orientation: ToStringPointer("Landscape"),
        PaperSize: ToStringPointer("A5"),
    }
    updateRequestOptions := map[string]interface{}{}
    updateRequest := &models.UpdateSectionPageSetupRequest{
        Name: ToStringPointer("Sample.docx"),
        SectionIndex: ToInt32Pointer(int32(0)),
        PageSetup: requestPageSetup,
        Optionals: updateRequestOptions,
    }
    _, _, _ = wordsApi.UpdateSectionPageSetup(ctx, updateRequest)
}



// func Test_Examples_UpdateSectionPageSetupOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_UpdateStyle(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestStyleUpdate := models.StyleUpdate{
        Name: ToStringPointer("My Style"),
    }
    updateRequestOptions := map[string]interface{}{}
    updateRequest := &models.UpdateStyleRequest{
        Name: ToStringPointer("Sample.docx"),
        StyleName: ToStringPointer("Heading 1"),
        StyleUpdate: requestStyleUpdate,
        Optionals: updateRequestOptions,
    }
    _, _, _ = wordsApi.UpdateStyle(ctx, updateRequest)
}



// func Test_Examples_UpdateStyleOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_UpdateTableCellFormat(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestFormat := models.TableCellFormat{
        BottomPadding: ToFloat64Pointer(5.0),
        FitText: ToBoolPointer(true),
        HorizontalMerge: ToStringPointer("First"),
        WrapText: ToBoolPointer(true),
    }
    updateRequestOptions := map[string]interface{}{}
    updateRequest := &models.UpdateTableCellFormatRequest{
        Name: ToStringPointer("Sample.docx"),
        TableRowPath: ToStringPointer("sections/0/tables/2/rows/0"),
        Index: ToInt32Pointer(int32(0)),
        Format: requestFormat,
        Optionals: updateRequestOptions,
    }
    _, _, _ = wordsApi.UpdateTableCellFormat(ctx, updateRequest)
}



// func Test_Examples_UpdateTableCellFormatOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_UpdateTableProperties(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestProperties := models.TableProperties{
        Alignment: ToStringPointer("Right"),
        AllowAutoFit: ToBoolPointer(false),
        Bidi: ToBoolPointer(true),
        BottomPadding: ToFloat64Pointer(1.0),
        CellSpacing: ToFloat64Pointer(2.0),
        StyleOptions: ToStringPointer("ColumnBands"),
    }
    updateRequestOptions := map[string]interface{}{}
    updateRequest := &models.UpdateTablePropertiesRequest{
        Name: ToStringPointer("Sample.docx"),
        Index: ToInt32Pointer(int32(1)),
        Properties: requestProperties,
        Optionals: updateRequestOptions,
    }
    _, _, _ = wordsApi.UpdateTableProperties(ctx, updateRequest)
}



// func Test_Examples_UpdateTablePropertiesOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_UpdateTableRowFormat(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestFormat := models.TableRowFormat{
        AllowBreakAcrossPages: ToBoolPointer(true),
        HeadingFormat: ToBoolPointer(true),
        Height: ToFloat64Pointer(10.0),
        HeightRule: ToStringPointer("Exactly"),
    }
    updateRequestOptions := map[string]interface{}{}
    updateRequest := &models.UpdateTableRowFormatRequest{
        Name: ToStringPointer("Sample.docx"),
        TablePath: ToStringPointer("sections/0/tables/2"),
        Index: ToInt32Pointer(int32(0)),
        Format: requestFormat,
        Optionals: updateRequestOptions,
    }
    _, _, _ = wordsApi.UpdateTableRowFormat(ctx, updateRequest)
}



// func Test_Examples_UpdateTableRowFormatOnline(t *testing.T) => ONLINE METHODS NOT SUPPORTED AT THIS MOMENT



func Test_Examples_UploadFile(t *testing.T) {
    documentsDir := GetExamplesDir()
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    requestFileContentFileData, _ := os.Open(documentsDir + "/" + "Sample.docx")
    uploadRequestOptions := map[string]interface{}{}
    uploadRequest := &models.UploadFileRequest{
        FileContent: requestFileContentFileData,
        Path: ToStringPointer("Sample.docx"),
        Optionals: uploadRequestOptions,
    }
    _, _, _ = wordsApi.UploadFile(ctx, uploadRequest)
}
