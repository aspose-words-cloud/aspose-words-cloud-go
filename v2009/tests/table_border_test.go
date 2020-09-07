/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="table_border_test.go">
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

// Example of how to work with table borders.
package api_test

import (
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2009/api/models"
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
    _, _, err := client.WordsApi.GetBorders(ctx, remoteFileName, options)

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
    _, _, err := client.WordsApi.GetBorder(ctx, remoteFileName, "left", options)

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
    _, _, err := client.WordsApi.DeleteBorders(ctx, remoteFileName, options)

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
    _, _, err := client.WordsApi.DeleteBorder(ctx, remoteFileName, "left", options)

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
        Alpha: int32(2),
    }
    requestBorderProperties := models.Border{
        BorderType: "Left",
        Color: &requestBorderPropertiesColor,
        DistanceFromText: 6,
        LineStyle: "DashDotStroker",
        LineWidth: 2,
        Shadow: true,
    }

    options := map[string]interface{}{
        "nodePath": "tables/1/rows/0/cells/0",
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.UpdateBorder(ctx, remoteFileName, requestBorderProperties, "left", options)

    if err != nil {
        t.Error(err)
    }

}
