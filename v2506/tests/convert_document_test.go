/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="convert_document_test.go">
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

// Example of how to convert document to one of the available formats.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2506/api/models"
)

// Test for converting document to one of the available formats.
func Test_ConvertDocument_SaveAs(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteFolder := remoteBaseTestDataFolder + "/DocumentActions/ConvertDocument"
    localName := "test_multi_pages.docx"
    remoteName := "TestSaveAs.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile("Common/" + localName), remoteFolder + "/" + remoteName)

    requestSaveOptionsData := models.PdfSaveOptionsData{
        FileName: ToStringPointer(baseTestOutPath + "/TestSaveAs.pdf"),
    }

    options := map[string]interface{}{
        "folder": remoteFolder,
    }

    request := &models.SaveAsRequest{
        Name: ToStringPointer(remoteName),
        SaveOptionsData: &requestSaveOptionsData,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.SaveAs(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetSaveResult(), "Validate SaveAs response.");
    assert.NotNil(t, actual.GetSaveResult().GetDestDocument(), "Validate SaveAs response.");
}

// Test for converting document online to one of the available formats.
func Test_ConvertDocument_SaveAsOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localName := "test_multi_pages.docx"

    requestDocument := OpenFile(t, "Common/" + localName)
    requestSaveOptionsData := models.PdfSaveOptionsData{
        FileName: ToStringPointer(baseTestOutPath + "/TestSaveAs.pdf"),
    }

    options := map[string]interface{}{
    }

    request := &models.SaveAsOnlineRequest{
        Document: requestDocument,
        SaveOptionsData: &requestSaveOptionsData,
        Optionals: options,
    }

    _, _, err := client.WordsApi.SaveAsOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for converting document online to html with additional files like css and images.
func Test_ConvertDocument_SaveAsOnlineHtmlMultifile(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localName := "test_multi_pages.docx"

    requestDocument := OpenFile(t, "Common/" + localName)
    requestSaveOptionsData := models.HtmlSaveOptionsData{
        FileName: ToStringPointer(baseTestOutPath + "/TestSaveAsHtml.html"),
        CssStyleSheetType: ToStringPointer("External"),
        CssStyleSheetFileName: ToStringPointer(baseTestOutPath + "/TestSaveAsHtml.css"),
    }

    options := map[string]interface{}{
    }

    request := &models.SaveAsOnlineRequest{
        Document: requestDocument,
        SaveOptionsData: &requestSaveOptionsData,
        Optionals: options,
    }

    _, _, err := client.WordsApi.SaveAsOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for converting document to one of the available formats.
func Test_ConvertDocument_SaveAsDocx(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteFolder := remoteBaseTestDataFolder + "/DocumentActions/ConvertDocument"
    localFolder := "DocumentActions/ConvertDocument"
    localName := "45.pdf"
    remoteName := "TestSaveAsFromPdfToDoc.pdf"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFolder + "/" + localName), remoteFolder + "/" + remoteName)

    requestSaveOptionsData := models.DocxSaveOptionsData{
        FileName: ToStringPointer(baseTestOutPath + "/TestSaveAsFromPdfToDoc.docx"),
    }

    options := map[string]interface{}{
        "folder": remoteFolder,
    }

    request := &models.SaveAsRequest{
        Name: ToStringPointer(remoteName),
        SaveOptionsData: &requestSaveOptionsData,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.SaveAs(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetSaveResult(), "Validate SaveAsDocx response.");
    assert.NotNil(t, actual.GetSaveResult().GetDestDocument(), "Validate SaveAsDocx response.");
}

// Test for converting document to one of the available formats.
func Test_ConvertDocument_SaveAsTiff(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteFolder := remoteBaseTestDataFolder + "/DocumentActions/ConvertDocument"
    localName := "test_multi_pages.docx"
    remoteName := "TestSaveAsTiff.pdf"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile("Common/" + localName), remoteFolder + "/" + remoteName)

    requestSaveOptions := models.TiffSaveOptionsData{
        FileName: ToStringPointer(baseTestOutPath + "/abc.tiff"),
    }

    options := map[string]interface{}{
        "folder": remoteFolder,
    }

    request := &models.SaveAsTiffRequest{
        Name: ToStringPointer(remoteName),
        SaveOptions: &requestSaveOptions,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.SaveAsTiff(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetSaveResult(), "Validate SaveAsTiff response.");
    assert.NotNil(t, actual.GetSaveResult().GetDestDocument(), "Validate SaveAsTiff response.");
}

// Test for converting document to one of the available formats.
func Test_ConvertDocument_SaveAsTiffOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localName := "test_multi_pages.docx"

    requestDocument := OpenFile(t, "Common/" + localName)
    requestSaveOptions := models.TiffSaveOptionsData{
        FileName: ToStringPointer(baseTestOutPath + "/abc.tiff"),
    }

    options := map[string]interface{}{
    }

    request := &models.SaveAsTiffOnlineRequest{
        Document: requestDocument,
        SaveOptions: &requestSaveOptions,
        Optionals: options,
    }

    _, _, err := client.WordsApi.SaveAsTiffOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// A test for ConvertDocument.
func Test_ConvertDocument_ConvertDocument(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFolder := "DocumentActions/ConvertDocument"

    requestDocument := OpenFile(t, localFolder + "/test_uploadfile.docx")

    options := map[string]interface{}{
    }

    request := &models.ConvertDocumentRequest{
        Document: requestDocument,
        Format: ToStringPointer("pdf"),
        Optionals: options,
    }

    _, err := client.WordsApi.ConvertDocument(ctx, request)
    if err != nil {
        t.Error(err)
    }

}
