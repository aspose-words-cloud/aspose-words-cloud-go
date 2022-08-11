/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="section_test.go">
 *   Copyright (c) 2022 Aspose.Words for Cloud
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

// Example of how to work with sections.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2208/api/models"
)

// Test for getting section by index.
func Test_Section_GetSection(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Section"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetSection.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.GetSectionRequest{
        Name: ToStringPointer(remoteFileName),
        SectionIndex: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetSection(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Section, "Validate GetSection response.");
    assert.NotNil(t, actual.Section.ChildNodes, "Validate GetSection response.");
    assert.Equal(t, 13, len(actual.Section.ChildNodes), "Validate GetSection response.");
    assert.Equal(t, "0.3.0", actual.Section.ChildNodes[0].NodeId, "Validate GetSection response.");
}

// Test for getting section by index online.
func Test_Section_GetSectionOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.GetSectionOnlineRequest{
        Document: requestDocument,
        SectionIndex: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetSectionOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for getting sections.
func Test_Section_GetSections(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Section"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetSections.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.GetSectionsRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetSections(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Sections, "Validate GetSections response.");
    assert.NotNil(t, actual.Sections.SectionLinkList, "Validate GetSections response.");
    assert.Equal(t, 1, len(actual.Sections.SectionLinkList), "Validate GetSections response.");
    assert.Equal(t, "0", actual.Sections.SectionLinkList[0].NodeId, "Validate GetSections response.");
}

// Test for getting sections online.
func Test_Section_GetSectionsOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.GetSectionsOnlineRequest{
        Document: requestDocument,
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetSectionsOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for delete a section.
func Test_Section_DeleteSection(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Section"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestDeleteSection.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.DeleteSectionRequest{
        Name: ToStringPointer(remoteFileName),
        SectionIndex: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, err := client.WordsApi.DeleteSection(ctx, request)

    if err != nil {
        t.Error(err)
    }

}

// Test for delete a section online.
func Test_Section_DeleteSectionOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.DeleteSectionOnlineRequest{
        Document: requestDocument,
        SectionIndex: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.DeleteSectionOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for linking headers and footers to previous section.
func Test_Section_LinkHeaderFootersToPrevious(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Section"
    remoteFileName := "TestLinkHeaderFootersToPrevious.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile("DocumentElements/Sections/Source.docx"), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.LinkHeaderFootersToPreviousRequest{
        Name: ToStringPointer(remoteFileName),
        SectionIndex: ToInt32Pointer(int32(1)),
        Optionals: options,
    }

    _, err := client.WordsApi.LinkHeaderFootersToPrevious(ctx, request)

    if err != nil {
        t.Error(err)
    }

}
