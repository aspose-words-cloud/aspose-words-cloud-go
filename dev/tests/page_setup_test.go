/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="page_setup_test.go">
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

// Example of how to work with macros.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/dev/api/models"
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
    actual, _, err := client.WordsApi.GetSectionPageSetup(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.PageSetup, "Validate GetSectionPageSetup response.");
    assert.Equal(t, int32(1), actual.PageSetup.LineStartingNumber, "Validate GetSectionPageSetup response.");
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
        RtlGutter: true,
        LeftMargin: 10.0,
        Orientation: "Landscape",
        PaperSize: "A5",
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.UpdateSectionPageSetup(ctx, remoteFileName, int32(0), requestPageSetup, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.PageSetup, "Validate UpdateSectionPageSetup response.");
    assert.True(t, actual.PageSetup.RtlGutter, "Validate UpdateSectionPageSetup response.");


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
    _, err := client.WordsApi.RenderPage(ctx, remoteFileName, int32(1), "bmp", options)

    if err != nil {
        t.Error(err)
    }

}
