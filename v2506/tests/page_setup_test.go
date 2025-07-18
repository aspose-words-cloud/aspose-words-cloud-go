/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="page_setup_test.go">
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

// Example of how to work with macros.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2506/api/models"
)

// Test for getting page settings.
func Test_PageSetup_GetSectionPageSetup(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/PageSetup"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetSectionPageSetup.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.GetSectionPageSetupRequest{
        Name: ToStringPointer(remoteFileName),
        SectionIndex: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetSectionPageSetup(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetPageSetup(), "Validate GetSectionPageSetup response.");
    assert.Equal(t, int32(1), DereferenceValue(actual.GetPageSetup().GetLineStartingNumber()), "Validate GetSectionPageSetup response.");
}

// Test for getting page settings online.
func Test_PageSetup_GetSectionPageSetupOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.GetSectionPageSetupOnlineRequest{
        Document: requestDocument,
        SectionIndex: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetSectionPageSetupOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for updating page settings.
func Test_PageSetup_UpdateSectionPageSetup(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/PageSetup"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestUpdateSectionPageSetup.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestPageSetup := models.PageSetup{
        RtlGutter: ToBoolPointer(true),
        LeftMargin: ToFloat64Pointer(10.0),
        Orientation: ToStringPointer("Landscape"),
        PaperSize: ToStringPointer("A5"),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.UpdateSectionPageSetupRequest{
        Name: ToStringPointer(remoteFileName),
        SectionIndex: ToInt32Pointer(int32(0)),
        PageSetup: &requestPageSetup,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.UpdateSectionPageSetup(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetPageSetup(), "Validate UpdateSectionPageSetup response.");
    assert.Equal(t, true, DereferenceValue(actual.GetPageSetup().GetRtlGutter()), "Validate UpdateSectionPageSetup response.");


}

// Test for updating page settings online.
func Test_PageSetup_UpdateSectionPageSetupOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)
    requestPageSetup := models.PageSetup{
        RtlGutter: ToBoolPointer(true),
        LeftMargin: ToFloat64Pointer(10),
        Orientation: ToStringPointer("Landscape"),
        PaperSize: ToStringPointer("A5"),
    }

    options := map[string]interface{}{
    }

    request := &models.UpdateSectionPageSetupOnlineRequest{
        Document: requestDocument,
        SectionIndex: ToInt32Pointer(int32(0)),
        PageSetup: &requestPageSetup,
        Optionals: options,
    }

    _, _, err := client.WordsApi.UpdateSectionPageSetupOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for page rendering.
func Test_PageSetup_GetRenderPage(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/PageSetup"
    localTextFile := "DocumentElements/Text/SampleWordDocument.docx"
    remoteFileName := "TestGetRenderPage.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localTextFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.RenderPageRequest{
        Name: ToStringPointer(remoteFileName),
        PageIndex: ToInt32Pointer(int32(1)),
        Format: ToStringPointer("jpg"),
        Optionals: options,
    }

    _, err := client.WordsApi.RenderPage(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for page rendering.
func Test_PageSetup_GetRenderPageOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localTextFile := "DocumentElements/Text/SampleWordDocument.docx"

    requestDocument := OpenFile(t, localTextFile)

    options := map[string]interface{}{
    }

    request := &models.RenderPageOnlineRequest{
        Document: requestDocument,
        PageIndex: ToInt32Pointer(int32(1)),
        Format: ToStringPointer("jpg"),
        Optionals: options,
    }

    _, err := client.WordsApi.RenderPageOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}
