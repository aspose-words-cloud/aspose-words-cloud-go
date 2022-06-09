/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="table_border_test.go">
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

// Example of how to work with table borders.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2206/api/models"
)

// Test for getting borders.
func Test_TableBorder_GetBorders(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Tables"
    localFile := "DocumentElements/Tables/TablesGet.docx"
    remoteFileName := "TestGetBorders.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "tables/1/rows/0/cells/0",
        "folder": remoteDataFolder,
    }

    request := &models.GetBordersRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetBorders(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Borders, "Validate GetBorders response.");
    assert.NotNil(t, actual.Borders.List, "Validate GetBorders response.");
    assert.Equal(t, 6, len(actual.Borders.List), "Validate GetBorders response.");
    assert.NotNil(t, actual.Borders.List[0].Color, "Validate GetBorders response.");
    assert.Equal(t, "#000000", actual.Borders.List[0].Color.Web, "Validate GetBorders response.");
}

// Test for getting borders online.
func Test_TableBorder_GetBordersOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Tables/TablesGet.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
        "nodePath": "tables/1/rows/0/cells/0",
    }

    request := &models.GetBordersOnlineRequest{
        Document: requestDocument,
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetBordersOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for getting border.
func Test_TableBorder_GetBorder(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Tables"
    localFile := "DocumentElements/Tables/TablesGet.docx"
    remoteFileName := "TestGetBorder.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "tables/1/rows/0/cells/0",
        "folder": remoteDataFolder,
    }

    request := &models.GetBorderRequest{
        Name: ToStringPointer(remoteFileName),
        BorderType: ToStringPointer("left"),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetBorder(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Border, "Validate GetBorder response.");
    assert.NotNil(t, actual.Border.Color, "Validate GetBorder response.");
    assert.Equal(t, "#000000", actual.Border.Color.Web, "Validate GetBorder response.");
}

// Test for getting border online.
func Test_TableBorder_GetBorderOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Tables/TablesGet.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
        "nodePath": "tables/1/rows/0/cells/0",
    }

    request := &models.GetBorderOnlineRequest{
        Document: requestDocument,
        BorderType: ToStringPointer("left"),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetBorderOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for deleting borders.
func Test_TableBorder_DeleteBorders(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Tables"
    localFile := "DocumentElements/Tables/TablesGet.docx"
    remoteFileName := "TestDeleteBorders.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "tables/1/rows/0/cells/0",
        "folder": remoteDataFolder,
    }

    request := &models.DeleteBordersRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: options,
    }

    _, _, err := client.WordsApi.DeleteBorders(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for deleting borders online.
func Test_TableBorder_DeleteBordersOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Tables/TablesGet.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
        "nodePath": "tables/1/rows/0/cells/0",
    }

    request := &models.DeleteBordersOnlineRequest{
        Document: requestDocument,
        Optionals: options,
    }

    _, _, err := client.WordsApi.DeleteBordersOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for deleting border.
func Test_TableBorder_DeleteBorder(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Tables"
    localFile := "DocumentElements/Tables/TablesGet.docx"
    remoteFileName := "TestDeleteBorder.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "tables/1/rows/0/cells/0",
        "folder": remoteDataFolder,
    }

    request := &models.DeleteBorderRequest{
        Name: ToStringPointer(remoteFileName),
        BorderType: ToStringPointer("left"),
        Optionals: options,
    }

    _, _, err := client.WordsApi.DeleteBorder(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for deleting border online.
func Test_TableBorder_DeleteBorderOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Tables/TablesGet.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
        "nodePath": "tables/1/rows/0/cells/0",
    }

    request := &models.DeleteBorderOnlineRequest{
        Document: requestDocument,
        BorderType: ToStringPointer("left"),
        Optionals: options,
    }

    _, _, err := client.WordsApi.DeleteBorderOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for updating border.
func Test_TableBorder_UpdateBorder(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Tables"
    localFile := "DocumentElements/Tables/TablesGet.docx"
    remoteFileName := "TestUpdateBorder.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestBorderPropertiesColor := models.XmlColor{
        Web: ToStringPointer("#AABBCC"),
    }
    requestBorderProperties := models.Border{
        BorderType: ToStringPointer("Left"),
        Color: &requestBorderPropertiesColor,
        DistanceFromText: ToFloat64Pointer(6.0),
        LineStyle: ToStringPointer("DashDotStroker"),
        LineWidth: ToFloat64Pointer(2.0),
        Shadow: ToBoolPointer(true),
    }

    options := map[string]interface{}{
        "nodePath": "tables/1/rows/0/cells/0",
        "folder": remoteDataFolder,
    }

    request := &models.UpdateBorderRequest{
        Name: ToStringPointer(remoteFileName),
        BorderType: ToStringPointer("left"),
        BorderProperties: &requestBorderProperties,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.UpdateBorder(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Border, "Validate UpdateBorder response.");
    assert.NotNil(t, actual.Border.Color, "Validate UpdateBorder response.");
    assert.Equal(t, "#AABBCC", actual.Border.Color.Web, "Validate UpdateBorder response.");
    assert.Equal(t, 6.0, actual.Border.DistanceFromText, "Validate UpdateBorder response.");
    assert.Equal(t, 2.0, actual.Border.LineWidth, "Validate UpdateBorder response.");
    assert.True(t, actual.Border.Shadow, "Validate UpdateBorder response.");
}

// Test for updating border online.
func Test_TableBorder_UpdateBorderOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Tables/TablesGet.docx"

    requestDocument := OpenFile(t, localFile)
    requestBorderPropertiesColor := models.XmlColor{
        Web: ToStringPointer("#AABBCC"),
    }
    requestBorderProperties := models.Border{
        BorderType: ToStringPointer("Left"),
        Color: &requestBorderPropertiesColor,
        DistanceFromText: ToFloat64Pointer(6),
        LineStyle: ToStringPointer("DashDotStroker"),
        LineWidth: ToFloat64Pointer(2),
        Shadow: ToBoolPointer(true),
    }

    options := map[string]interface{}{
        "nodePath": "tables/1/rows/0/cells/0",
    }

    request := &models.UpdateBorderOnlineRequest{
        Document: requestDocument,
        BorderProperties: &requestBorderProperties,
        BorderType: ToStringPointer("left"),
        Optionals: options,
    }

    _, _, err := client.WordsApi.UpdateBorderOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}
