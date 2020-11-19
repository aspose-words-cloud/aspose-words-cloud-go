/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="styles_test.go">
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

// Example of how to work with styles.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2011/api/models"
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
    actual, _, err := client.WordsApi.GetStyles(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Styles, "Validate GetStyles response.");
    assert.Equal(t, 22, len(actual.Styles), "Validate GetStyles response.");
    assert.Equal(t, "Default Paragraph Font", *actual.Styles[0].Name, "Validate GetStyles response.");
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
    actual, _, err := client.WordsApi.GetStyle(ctx, remoteFileName, "Heading 1", options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Style, "Validate GetStyle response.");
    assert.Equal(t, "Heading 1", *actual.Style.Name, "Validate GetStyle response.");
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
    actual, _, err := client.WordsApi.UpdateStyle(ctx, remoteFileName, requestStyleUpdate, "Heading 1", options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Style, "Validate UpdateStyle response.");
    assert.Equal(t, "My Style", *actual.Style.Name, "Validate UpdateStyle response.");
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
        StyleType: "Paragraph",
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.InsertStyle(ctx, remoteFileName, requestStyleInsert, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Style, "Validate InsertStyle response.");
    assert.Equal(t, "My Style", *actual.Style.Name, "Validate InsertStyle response.");
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
    actual, _, err := client.WordsApi.CopyStyle(ctx, remoteFileName, requestStyleCopy, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Style, "Validate CopyStyle response.");
    assert.Equal(t, "Heading 1_0", *actual.Style.Name, "Validate CopyStyle response.");
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
    actual, _, err := client.WordsApi.GetStyleFromDocumentElement(ctx, remoteFileName, "paragraphs/1/paragraphFormat", options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Style, "Validate GetStyleFromDocumentElement response.");
    assert.Equal(t, "TOC 1", *actual.Style.Name, "Validate GetStyleFromDocumentElement response.");
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
    _, _, err := client.WordsApi.ApplyStyleToDocumentElement(ctx, remoteFileName, requestStyleApply, "paragraphs/1/paragraphFormat", options)

    if err != nil {
        t.Error(err)
    }

}
