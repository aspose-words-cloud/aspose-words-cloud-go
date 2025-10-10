/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="styles_test.go">
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

// Example of how to work with styles.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2510/api/models"
)

// Test for getting styles from document.
func Test_Styles_GetStyles(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Styles"
    localFile := "DocumentElements/Styles/GetStyles.docx"
    remoteFileName := "TestGetStyles.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.GetStylesRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetStyles(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetStyles(), "Validate GetStyles response.");
    assert.Equal(t, 22, len(actual.GetStyles()), "Validate GetStyles response.");
    assert.Equal(t, "Default Paragraph Font", DereferenceValue(actual.GetStyles()[0].GetName()), "Validate GetStyles response.");
}

// Test for getting styles from document online.
func Test_Styles_GetStylesOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Styles/GetStyles.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.GetStylesOnlineRequest{
        Document: requestDocument,
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetStylesOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for getting style from document.
func Test_Styles_GetStyle(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Styles"
    localFile := "DocumentElements/Styles/GetStyles.docx"
    remoteFileName := "TestGetStyle.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.GetStyleRequest{
        Name: ToStringPointer(remoteFileName),
        StyleName: ToStringPointer("Heading 1"),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetStyle(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetStyle(), "Validate GetStyle response.");
    assert.Equal(t, "Heading 1", DereferenceValue(actual.GetStyle().GetName()), "Validate GetStyle response.");
}

// Test for getting style from document online.
func Test_Styles_GetStyleOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Styles/GetStyles.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.GetStyleOnlineRequest{
        Document: requestDocument,
        StyleName: ToStringPointer("Heading 1"),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetStyleOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for updating style from document.
func Test_Styles_UpdateStyle(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Styles"
    localFile := "DocumentElements/Styles/GetStyles.docx"
    remoteFileName := "TestUpdateStyle.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestStyleUpdate := models.StyleUpdate{
        Name: ToStringPointer("My Style"),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.UpdateStyleRequest{
        Name: ToStringPointer(remoteFileName),
        StyleName: ToStringPointer("Heading 1"),
        StyleUpdate: &requestStyleUpdate,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.UpdateStyle(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetStyle(), "Validate UpdateStyle response.");
    assert.Equal(t, "My Style", DereferenceValue(actual.GetStyle().GetName()), "Validate UpdateStyle response.");
}

// Test for updating style from document online.
func Test_Styles_UpdateStyleOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Styles/GetStyles.docx"

    requestDocument := OpenFile(t, localFile)
    requestStyleUpdate := models.StyleUpdate{
        Name: ToStringPointer("My Style"),
    }

    options := map[string]interface{}{
    }

    request := &models.UpdateStyleOnlineRequest{
        Document: requestDocument,
        StyleName: ToStringPointer("Heading 1"),
        StyleUpdate: &requestStyleUpdate,
        Optionals: options,
    }

    _, _, err := client.WordsApi.UpdateStyleOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for inserting style from document.
func Test_Styles_InsertStyle(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Styles"
    localFile := "DocumentElements/Styles/GetStyles.docx"
    remoteFileName := "TestInsertStyle.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestStyleInsert := models.StyleInsert{
        StyleName: ToStringPointer("My Style"),
        StyleType: ToStringPointer("Paragraph"),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.InsertStyleRequest{
        Name: ToStringPointer(remoteFileName),
        StyleInsert: &requestStyleInsert,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.InsertStyle(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetStyle(), "Validate InsertStyle response.");
    assert.Equal(t, "My Style", DereferenceValue(actual.GetStyle().GetName()), "Validate InsertStyle response.");
}

// Test for inserting style from document online.
func Test_Styles_InsertStyleOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Styles/GetStyles.docx"

    requestDocument := OpenFile(t, localFile)
    requestStyleInsert := models.StyleInsert{
        StyleName: ToStringPointer("My Style"),
        StyleType: ToStringPointer("Paragraph"),
    }

    options := map[string]interface{}{
    }

    request := &models.InsertStyleOnlineRequest{
        Document: requestDocument,
        StyleInsert: &requestStyleInsert,
        Optionals: options,
    }

    _, _, err := client.WordsApi.InsertStyleOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for coping style from document.
func Test_Styles_CopyStyle(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Styles"
    localFile := "DocumentElements/Styles/GetStyles.docx"
    remoteFileName := "TestCopyStyle.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestStyleCopy := models.StyleCopy{
        StyleName: ToStringPointer("Heading 1"),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.CopyStyleRequest{
        Name: ToStringPointer(remoteFileName),
        StyleCopy: &requestStyleCopy,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.CopyStyle(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetStyle(), "Validate CopyStyle response.");
    assert.Equal(t, "Heading 1_0", DereferenceValue(actual.GetStyle().GetName()), "Validate CopyStyle response.");
}

// Test for coping style from document online.
func Test_Styles_CopyStyleOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Styles/GetStyles.docx"

    requestDocument := OpenFile(t, localFile)
    requestStyleCopy := models.StyleCopy{
        StyleName: ToStringPointer("Heading 1"),
    }

    options := map[string]interface{}{
    }

    request := &models.CopyStyleOnlineRequest{
        Document: requestDocument,
        StyleCopy: &requestStyleCopy,
        Optionals: options,
    }

    _, _, err := client.WordsApi.CopyStyleOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for getting style from document element.
func Test_Styles_GetStyleFromDocumentElement(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Styles"
    localFile := "DocumentElements/Styles/GetStyles.docx"
    remoteFileName := "TestGetStyleFromDocumentElement.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.GetStyleFromDocumentElementRequest{
        Name: ToStringPointer(remoteFileName),
        StyledNodePath: ToStringPointer("paragraphs/1/paragraphFormat"),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetStyleFromDocumentElement(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetStyle(), "Validate GetStyleFromDocumentElement response.");
    assert.Equal(t, "TOC 1", DereferenceValue(actual.GetStyle().GetName()), "Validate GetStyleFromDocumentElement response.");
}

// Test for getting style from document element online.
func Test_Styles_GetStyleFromDocumentElementOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Styles/GetStyles.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.GetStyleFromDocumentElementOnlineRequest{
        Document: requestDocument,
        StyledNodePath: ToStringPointer("paragraphs/1/paragraphFormat"),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetStyleFromDocumentElementOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for applying style to document element.
func Test_Styles_ApplyStyleToDocumentElement(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Styles"
    localFile := "DocumentElements/Styles/GetStyles.docx"
    remoteFileName := "TestApplyStyleToDocumentElement.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestStyleApply := models.StyleApply{
        StyleName: ToStringPointer("Heading 1"),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.ApplyStyleToDocumentElementRequest{
        Name: ToStringPointer(remoteFileName),
        StyledNodePath: ToStringPointer("paragraphs/1/paragraphFormat"),
        StyleApply: &requestStyleApply,
        Optionals: options,
    }

    _, _, err := client.WordsApi.ApplyStyleToDocumentElement(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for applying style to document element online.
func Test_Styles_ApplyStyleToDocumentElementOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Styles/GetStyles.docx"

    requestDocument := OpenFile(t, localFile)
    requestStyleApply := models.StyleApply{
        StyleName: ToStringPointer("Heading 1"),
    }

    options := map[string]interface{}{
    }

    request := &models.ApplyStyleToDocumentElementOnlineRequest{
        Document: requestDocument,
        StyledNodePath: ToStringPointer("paragraphs/1/paragraphFormat"),
        StyleApply: &requestStyleApply,
        Optionals: options,
    }

    _, _, err := client.WordsApi.ApplyStyleToDocumentElementOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for copying styles from a template.
func Test_Styles_CopyStylesFromTemplate(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Styles"
    localFile := "DocumentElements/Styles/GetStyles.docx"
    remoteFileName := "TestCopyStylesFromTemplate.docx"
    templateFolder := "DocumentElements/Styles"
    templateName := "StyleTemplate.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)
    UploadNextFileToStorage(t, ctx, client, GetLocalFile(templateFolder + "/" + templateName), remoteDataFolder + "/" + templateName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.CopyStylesFromTemplateRequest{
        Name: ToStringPointer(remoteFileName),
        TemplateName: ToStringPointer(templateName),
        Optionals: options,
    }

    _, _, err := client.WordsApi.CopyStylesFromTemplate(ctx, request)
    if err != nil {
        t.Error(err)
    }

}
