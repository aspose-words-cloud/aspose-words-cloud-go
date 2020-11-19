/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="header_footer_test.go">
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

// Example of how to work with headers and footers.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
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
    actual, _, err := client.WordsApi.GetHeaderFooters(ctx, remoteFileName, "", options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.HeaderFooters, "Validate GetHeaderFooters response.");
    assert.NotNil(t, actual.HeaderFooters.List, "Validate GetHeaderFooters response.");
    assert.Equal(t, 6, len(actual.HeaderFooters.List), "Validate GetHeaderFooters response.");
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
    actual, _, err := client.WordsApi.GetHeaderFooter(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.HeaderFooter, "Validate GetHeaderFooter response.");
    assert.NotNil(t, actual.HeaderFooter.ChildNodes, "Validate GetHeaderFooter response.");
    assert.Equal(t, 1, len(actual.HeaderFooter.ChildNodes), "Validate GetHeaderFooter response.");
    assert.Equal(t, "0.0.0", *actual.HeaderFooter.ChildNodes[0].NodeId, "Validate GetHeaderFooter response.");
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
    actual, _, err := client.WordsApi.GetHeaderFooterOfSection(ctx, remoteFileName, int32(0), int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.HeaderFooter, "Validate GetHeaderFooterOfSection response.");
    assert.NotNil(t, actual.HeaderFooter.ChildNodes, "Validate GetHeaderFooterOfSection response.");
    assert.Equal(t, 1, len(actual.HeaderFooter.ChildNodes), "Validate GetHeaderFooterOfSection response.");
    assert.Equal(t, "0.0.0", *actual.HeaderFooter.ChildNodes[0].NodeId, "Validate GetHeaderFooterOfSection response.");
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
    _, err := client.WordsApi.DeleteHeaderFooter(ctx, remoteFileName, "", int32(0), options)

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
    _, err := client.WordsApi.DeleteHeadersFooters(ctx, remoteFileName, "", options)

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
    actual, _, err := client.WordsApi.InsertHeaderFooter(ctx, remoteFileName, "FooterEven", "", options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.HeaderFooter, "Validate InsertHeaderFooter response.");
    assert.NotNil(t, actual.HeaderFooter.ChildNodes, "Validate InsertHeaderFooter response.");
    assert.Equal(t, 1, len(actual.HeaderFooter.ChildNodes), "Validate InsertHeaderFooter response.");
    assert.Equal(t, "0.2.0", *actual.HeaderFooter.ChildNodes[0].NodeId, "Validate InsertHeaderFooter response.");
}
