/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="table_test.go">
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

// Example of how to work wtih table.
package api_test

import (
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2009/api/models"
)

// Test for getting tables.
func Test_Table_GetTables(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Tables"
    localFile := "DocumentElements/Tables/TablesGet.docx"
    remoteFileName := "TestGetTables.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetTables(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }

}

// Test for getting tables without node path.
func Test_Table_GetTablesWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Tables"
    localFile := "DocumentElements/Tables/TablesGet.docx"
    remoteFileName := "TestGetTablesWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetTables(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }

}

// Test for getting table.
func Test_Table_GetTable(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Tables"
    localFile := "DocumentElements/Tables/TablesGet.docx"
    remoteFileName := "TestGetTable.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetTable(ctx, remoteFileName, int32(1), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for getting table without node path.
func Test_Table_GetTableWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Tables"
    localFile := "DocumentElements/Tables/TablesGet.docx"
    remoteFileName := "TestGetTableWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetTable(ctx, remoteFileName, int32(1), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for deleting table.
func Test_Table_DeleteTable(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Tables"
    localFile := "DocumentElements/Tables/TablesGet.docx"
    remoteFileName := "TestDeleteTable.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }
    _, err := client.WordsApi.DeleteTable(ctx, remoteFileName, int32(1), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for deleting table without node path.
func Test_Table_DeleteTableWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Tables"
    localFile := "DocumentElements/Tables/TablesGet.docx"
    remoteFileName := "TestDeleteTableWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, err := client.WordsApi.DeleteTable(ctx, remoteFileName, int32(1), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for adding table.
func Test_Table_InsertTable(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Tables"
    localFile := "DocumentElements/Tables/TablesGet.docx"
    remoteFileName := "TestInsertTable.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestTable := models.TableInsert{
        ColumnsCount: int32(5),
        RowsCount: int32(4),
    }

    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.InsertTable(ctx, remoteFileName, requestTable, options)

    if err != nil {
        t.Error(err)
    }

}

// Test for adding table without node path.
func Test_Table_InsertTableWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Tables"
    localFile := "DocumentElements/Tables/TablesGet.docx"
    remoteFileName := "TestInsertTableWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestTable := models.TableInsert{
        ColumnsCount: int32(5),
        RowsCount: int32(4),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.InsertTable(ctx, remoteFileName, requestTable, options)

    if err != nil {
        t.Error(err)
    }

}

// Test for getting document properties.
func Test_Table_GetTableProperties(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Tables"
    localFile := "DocumentElements/Tables/TablesGet.docx"
    remoteFileName := "TestGetTableProperties.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetTableProperties(ctx, remoteFileName, int32(1), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for getting document properties without node path.
func Test_Table_GetTablePropertiesWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Tables"
    localFile := "DocumentElements/Tables/TablesGet.docx"
    remoteFileName := "TestGetTablePropertiesWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetTableProperties(ctx, remoteFileName, int32(1), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for updating table properties.
func Test_Table_UpdateTableProperties(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Tables"
    localFile := "DocumentElements/Tables/TablesGet.docx"
    remoteFileName := "TestUpdateTableProperties.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestProperties := models.TableProperties{
        Alignment: "Right",
        AllowAutoFit: false,
        Bidi: true,
        BottomPadding: 1,
        CellSpacing: 2,
        StyleOptions: "ColumnBands",
    }

    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.UpdateTableProperties(ctx, remoteFileName, requestProperties, int32(1), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for updating table properties without node path.
func Test_Table_UpdateTablePropertiesWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Tables"
    localFile := "DocumentElements/Tables/TablesGet.docx"
    remoteFileName := "TestUpdateTablePropertiesWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestProperties := models.TableProperties{
        Alignment: "Right",
        AllowAutoFit: false,
        Bidi: true,
        BottomPadding: 1,
        CellSpacing: 2,
        StyleOptions: "ColumnBands",
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.UpdateTableProperties(ctx, remoteFileName, requestProperties, int32(1), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for getting table row.
func Test_Table_GetTableRow(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Tables"
    localFile := "DocumentElements/Tables/TablesGet.docx"
    remoteFileName := "TestGetTableRow.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetTableRow(ctx, remoteFileName, "tables/1", int32(0), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for deleting table row.
func Test_Table_DeleteTableRow(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Tables"
    localFile := "DocumentElements/Tables/TablesGet.docx"
    remoteFileName := "TestDeleteTableRow.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, err := client.WordsApi.DeleteTableRow(ctx, remoteFileName, "tables/1", int32(0), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for adding row.
func Test_Table_InsertTableRow(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Tables"
    localFile := "DocumentElements/Tables/TablesGet.docx"
    remoteFileName := "TestInsertTableRow.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestRow := models.TableRowInsert{
        ColumnsCount: int32(5),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.InsertTableRow(ctx, remoteFileName, requestRow, "sections/0/tables/2", options)

    if err != nil {
        t.Error(err)
    }

}

// Test for getting row format.
func Test_Table_GetTableRowFormat(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Tables"
    localFile := "DocumentElements/Tables/TablesGet.docx"
    remoteFileName := "TestGetTableRowFormat.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetTableRowFormat(ctx, remoteFileName, "sections/0/tables/2", int32(0), options)

    if err != nil {
        t.Error(err)
    }

}

// Test updating row format.
func Test_Table_UpdateTableRowFormat(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Tables"
    localFile := "DocumentElements/Tables/TablesGet.docx"
    remoteFileName := "TestUpdateTableRowFormat.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestFormat := models.TableRowFormat{
        AllowBreakAcrossPages: true,
        HeadingFormat: true,
        Height: 10,
        HeightRule: "Auto",
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.UpdateTableRowFormat(ctx, remoteFileName, requestFormat, "sections/0/tables/2", int32(0), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for getting table cell.
func Test_Table_GetTableCell(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Tables"
    localFile := "DocumentElements/Tables/TablesGet.docx"
    remoteFileName := "TestGetTableCell.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetTableCell(ctx, remoteFileName, "sections/0/tables/2/rows/0", int32(0), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for deleting cell.
func Test_Table_DeleteTableCell(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Tables"
    localFile := "DocumentElements/Tables/TablesGet.docx"
    remoteFileName := "TestDeleteTableCell.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, err := client.WordsApi.DeleteTableCell(ctx, remoteFileName, "sections/0/tables/2/rows/0", int32(0), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for adding cell.
func Test_Table_InsertTableCell(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Tables"
    localFile := "DocumentElements/Tables/TablesGet.docx"
    remoteFileName := "TestInsertTableCell.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestCell := models.TableCellInsert{
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.InsertTableCell(ctx, remoteFileName, requestCell, "sections/0/tables/2/rows/0", options)

    if err != nil {
        t.Error(err)
    }

}

// Test for getting cell format.
func Test_Table_GetTableCellFormat(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Tables"
    localFile := "DocumentElements/Tables/TablesGet.docx"
    remoteFileName := "TestGetTableCellFormat.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetTableCellFormat(ctx, remoteFileName, "sections/0/tables/2/rows/0", int32(0), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for updating cell format.
func Test_Table_UpdateTableCellFormat(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Tables"
    localFile := "DocumentElements/Tables/TablesGet.docx"
    remoteFileName := "TestUpdateTableCellFormat.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestFormat := models.TableCellFormat{
        BottomPadding: 5,
        FitText: true,
        HorizontalMerge: "First",
        WrapText: true,
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.UpdateTableCellFormat(ctx, remoteFileName, requestFormat, "sections/0/tables/2/rows/0", int32(0), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for table rendering.
func Test_Table_RenderTable(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Tables"
    localFile := "DocumentElements/Tables/TablesGet.docx"
    remoteFileName := "TestRenderTable.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }
    _, err := client.WordsApi.RenderTable(ctx, remoteFileName, "png", int32(0), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for table rendering without node path.
func Test_Table_RenderTableWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Tables"
    localFile := "DocumentElements/Tables/TablesGet.docx"
    remoteFileName := "TestRenderTableWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, err := client.WordsApi.RenderTable(ctx, remoteFileName, "png", int32(0), options)

    if err != nil {
        t.Error(err)
    }

}
