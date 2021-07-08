/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="header_footer_test.go">
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

// Example of how to work with headers and footers.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2106/api/models"
)

// Test for getting headers and footers.
func Test_HeaderFooter_GetHeaderFooters(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/HeaderFooters"
    localFile := "DocumentElements/HeaderFooters/HeadersFooters.doc"
    remoteFileName := "TestGetHeadersFooters.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.GetHeaderFootersRequest{
        Name: ToStringPointer(remoteFileName),
        SectionPath: ToStringPointer(""),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetHeaderFooters(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.HeaderFooters, "Validate GetHeaderFooters response.");
    assert.NotNil(t, actual.HeaderFooters.List, "Validate GetHeaderFooters response.");
    assert.Equal(t, 6, len(actual.HeaderFooters.List), "Validate GetHeaderFooters response.");
}

// Test for getting headers and footers online.
func Test_HeaderFooter_GetHeaderFootersOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/HeaderFooters/HeadersFooters.doc"


    options := map[string]interface{}{
    }

    request := &models.GetHeaderFootersOnlineRequest{
        Document: OpenFile(t, localFile),
        SectionPath: ToStringPointer(""),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetHeaderFootersOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for getting headerfooter.
func Test_HeaderFooter_GetHeaderFooter(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/HeaderFooters"
    localFile := "DocumentElements/HeaderFooters/HeadersFooters.doc"
    remoteFileName := "TestGetHeaderFooter.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.GetHeaderFooterRequest{
        Name: ToStringPointer(remoteFileName),
        HeaderFooterIndex: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetHeaderFooter(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.HeaderFooter, "Validate GetHeaderFooter response.");
    assert.NotNil(t, actual.HeaderFooter.ChildNodes, "Validate GetHeaderFooter response.");
    assert.Equal(t, 1, len(actual.HeaderFooter.ChildNodes), "Validate GetHeaderFooter response.");
    assert.Equal(t, "0.0.0", actual.HeaderFooter.ChildNodes[0].NodeId, "Validate GetHeaderFooter response.");
}

// Test for getting headerfooter online.
func Test_HeaderFooter_GetHeaderFooterOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/HeaderFooters/HeadersFooters.doc"


    options := map[string]interface{}{
    }

    request := &models.GetHeaderFooterOnlineRequest{
        Document: OpenFile(t, localFile),
        HeaderFooterIndex: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetHeaderFooterOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for getting headerfooter of section.
func Test_HeaderFooter_GetHeaderFooterOfSection(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/HeaderFooters"
    localFile := "DocumentElements/HeaderFooters/HeadersFooters.doc"
    remoteFileName := "TestGetHeaderFooterOfSection.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.GetHeaderFooterOfSectionRequest{
        Name: ToStringPointer(remoteFileName),
        HeaderFooterIndex: ToInt32Pointer(int32(0)),
        SectionIndex: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetHeaderFooterOfSection(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.HeaderFooter, "Validate GetHeaderFooterOfSection response.");
    assert.NotNil(t, actual.HeaderFooter.ChildNodes, "Validate GetHeaderFooterOfSection response.");
    assert.Equal(t, 1, len(actual.HeaderFooter.ChildNodes), "Validate GetHeaderFooterOfSection response.");
    assert.Equal(t, "0.0.0", actual.HeaderFooter.ChildNodes[0].NodeId, "Validate GetHeaderFooterOfSection response.");
}

// Test for getting headerfooter of section online.
func Test_HeaderFooter_GetHeaderFooterOfSectionOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/HeaderFooters/HeadersFooters.doc"


    options := map[string]interface{}{
    }

    request := &models.GetHeaderFooterOfSectionOnlineRequest{
        Document: OpenFile(t, localFile),
        HeaderFooterIndex: ToInt32Pointer(int32(0)),
        SectionIndex: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetHeaderFooterOfSectionOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for deleting headerfooter.
func Test_HeaderFooter_DeleteHeaderFooter(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/HeaderFooters"
    localFile := "DocumentElements/HeaderFooters/HeadersFooters.doc"
    remoteFileName := "TestDeleteHeaderFooter.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.DeleteHeaderFooterRequest{
        Name: ToStringPointer(remoteFileName),
        SectionPath: ToStringPointer(""),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, err := client.WordsApi.DeleteHeaderFooter(ctx, request)

    if err != nil {
        t.Error(err)
    }

}

// Test for deleting headerfooter online.
func Test_HeaderFooter_DeleteHeaderFooterOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/HeaderFooters/HeadersFooters.doc"


    options := map[string]interface{}{
    }

    request := &models.DeleteHeaderFooterOnlineRequest{
        Document: OpenFile(t, localFile),
        SectionPath: ToStringPointer(""),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, err := client.WordsApi.DeleteHeaderFooterOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for deleting headerfooters.
func Test_HeaderFooter_DeleteHeadersFooters(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/HeaderFooters"
    localFile := "DocumentElements/HeaderFooters/HeadersFooters.doc"
    remoteFileName := "TestDeleteHeadersFooters.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.DeleteHeadersFootersRequest{
        Name: ToStringPointer(remoteFileName),
        SectionPath: ToStringPointer(""),
        Optionals: options,
    }

    _, err := client.WordsApi.DeleteHeadersFooters(ctx, request)

    if err != nil {
        t.Error(err)
    }

}

// Test for deleting headerfooters online.
func Test_HeaderFooter_DeleteHeadersFootersOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/HeaderFooters/HeadersFooters.doc"


    options := map[string]interface{}{
    }

    request := &models.DeleteHeadersFootersOnlineRequest{
        Document: OpenFile(t, localFile),
        SectionPath: ToStringPointer(""),
        Optionals: options,
    }

    _, err := client.WordsApi.DeleteHeadersFootersOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for adding headerfooters.
func Test_HeaderFooter_InsertHeaderFooter(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/HeaderFooters"
    localFile := "DocumentElements/HeaderFooters/HeadersFooters.doc"
    remoteFileName := "TestInsertHeaderFooter.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.InsertHeaderFooterRequest{
        Name: ToStringPointer(remoteFileName),
        SectionPath: ToStringPointer(""),
        HeaderFooterType: ToStringPointer("FooterEven"),
        Optionals: options,
    }

    _, _, err := client.WordsApi.InsertHeaderFooter(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for adding headerfooters online.
func Test_HeaderFooter_InsertHeaderFooterOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/HeaderFooters/HeadersFooters.doc"


    options := map[string]interface{}{
    }

    request := &models.InsertHeaderFooterOnlineRequest{
        Document: OpenFile(t, localFile),
        SectionPath: ToStringPointer(""),
        HeaderFooterType: ToStringPointer("FooterEven"),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.InsertHeaderFooterOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Model.HeaderFooter, "Validate InsertHeaderFooterOnline response.");
    assert.NotNil(t, actual.Model.HeaderFooter.ChildNodes, "Validate InsertHeaderFooterOnline response.");
    assert.Equal(t, 1, len(actual.Model.HeaderFooter.ChildNodes), "Validate InsertHeaderFooterOnline response.");
    assert.Equal(t, "0.2.0", actual.Model.HeaderFooter.ChildNodes[0].NodeId, "Validate InsertHeaderFooterOnline response.");
}
