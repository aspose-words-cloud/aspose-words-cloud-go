/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="table_test.go">
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

// Example of how to work wtih table.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2104/api/models"
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

    request := &models.GetTablesRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetTables(ctx, request)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Tables, "Validate GetTables response.");
    assert.NotNil(t, actual.Tables.TableLinkList, "Validate GetTables response.");
    assert.Equal(t, 5, len(actual.Tables.TableLinkList), "Validate GetTables response.");
    assert.Equal(t, "0.0.1", actual.Tables.TableLinkList[0].NodeId, "Validate GetTables response.");
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

    request := &models.GetTablesRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetTables(ctx, request)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Tables, "Validate GetTablesWithoutNodePath response.");
    assert.NotNil(t, actual.Tables.TableLinkList, "Validate GetTablesWithoutNodePath response.");
    assert.Equal(t, 5, len(actual.Tables.TableLinkList), "Validate GetTablesWithoutNodePath response.");
    assert.Equal(t, "0.0.1", actual.Tables.TableLinkList[0].NodeId, "Validate GetTablesWithoutNodePath response.");
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

    request := &models.GetTableRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(1)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetTable(ctx, request)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Table, "Validate GetTable response.");
    assert.NotNil(t, actual.Table.TableRowList, "Validate GetTable response.");
    assert.Equal(t, 1, len(actual.Table.TableRowList), "Validate GetTable response.");
    assert.NotNil(t, actual.Table.TableRowList[0].TableCellList, "Validate GetTable response.");
    assert.Equal(t, 2, len(actual.Table.TableRowList[0].TableCellList), "Validate GetTable response.");
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

    request := &models.GetTableRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(1)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetTable(ctx, request)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Table, "Validate GetTableWithoutNodePath response.");
    assert.NotNil(t, actual.Table.TableRowList, "Validate GetTableWithoutNodePath response.");
    assert.Equal(t, 1, len(actual.Table.TableRowList), "Validate GetTableWithoutNodePath response.");
    assert.NotNil(t, actual.Table.TableRowList[0].TableCellList, "Validate GetTableWithoutNodePath response.");
    assert.Equal(t, 2, len(actual.Table.TableRowList[0].TableCellList), "Validate GetTableWithoutNodePath response.");
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

    request := &models.DeleteTableRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(1)),
        Optionals: options,
    }

_, err := client.WordsApi.DeleteTable(ctx, request)

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

    request := &models.DeleteTableRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(1)),
        Optionals: options,
    }

_, err := client.WordsApi.DeleteTable(ctx, request)

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
        ColumnsCount: ToInt32Pointer(int32(5)),
        RowsCount: ToInt32Pointer(int32(4)),
    }

    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }

    request := &models.InsertTableRequest{
        Name: ToStringPointer(remoteFileName),
        Table: requestTable,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.InsertTable(ctx, request)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Table, "Validate InsertTable response.");
    assert.NotNil(t, actual.Table.TableRowList, "Validate InsertTable response.");
    assert.Equal(t, 4, len(actual.Table.TableRowList), "Validate InsertTable response.");
    assert.NotNil(t, actual.Table.TableRowList[0].TableCellList, "Validate InsertTable response.");
    assert.Equal(t, 5, len(actual.Table.TableRowList[0].TableCellList), "Validate InsertTable response.");
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
        ColumnsCount: ToInt32Pointer(int32(5)),
        RowsCount: ToInt32Pointer(int32(4)),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.InsertTableRequest{
        Name: ToStringPointer(remoteFileName),
        Table: requestTable,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.InsertTable(ctx, request)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Table, "Validate InsertTableWithoutNodePath response.");
    assert.NotNil(t, actual.Table.TableRowList, "Validate InsertTableWithoutNodePath response.");
    assert.Equal(t, 4, len(actual.Table.TableRowList), "Validate InsertTableWithoutNodePath response.");
    assert.NotNil(t, actual.Table.TableRowList[0].TableCellList, "Validate InsertTableWithoutNodePath response.");
    assert.Equal(t, 5, len(actual.Table.TableRowList[0].TableCellList), "Validate InsertTableWithoutNodePath response.");
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

    request := &models.GetTablePropertiesRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(1)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetTableProperties(ctx, request)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Properties, "Validate GetTableProperties response.");
    assert.Equal(t, "Table Grid", actual.Properties.StyleName, "Validate GetTableProperties response.");
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

    request := &models.GetTablePropertiesRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(1)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetTableProperties(ctx, request)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Properties, "Validate GetTablePropertiesWithoutNodePath response.");
    assert.Equal(t, "Table Grid", actual.Properties.StyleName, "Validate GetTablePropertiesWithoutNodePath response.");
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
        Alignment: ToStringPointer("Right"),
        AllowAutoFit: ToBoolPointer(false),
        Bidi: ToBoolPointer(true),
        BottomPadding: ToFloat64Pointer(1),
        CellSpacing: ToFloat64Pointer(2.0),
        StyleOptions: ToStringPointer("ColumnBands"),
    }

    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }

    request := &models.UpdateTablePropertiesRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(1)),
        Properties: requestProperties,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.UpdateTableProperties(ctx, request)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Properties, "Validate UpdateTableProperties response.");
    assert.False(t, actual.Properties.AllowAutoFit, "Validate UpdateTableProperties response.");
    assert.True(t, actual.Properties.Bidi, "Validate UpdateTableProperties response.");
    assert.Equal(t, 1.0, actual.Properties.BottomPadding, "Validate UpdateTableProperties response.");
    assert.Equal(t, 2.0, actual.Properties.CellSpacing, "Validate UpdateTableProperties response.");
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
        Alignment: ToStringPointer("Right"),
        AllowAutoFit: ToBoolPointer(false),
        Bidi: ToBoolPointer(true),
        BottomPadding: ToFloat64Pointer(1.0),
        CellSpacing: ToFloat64Pointer(2.0),
        StyleOptions: ToStringPointer("ColumnBands"),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.UpdateTablePropertiesRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(1)),
        Properties: requestProperties,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.UpdateTableProperties(ctx, request)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Properties, "Validate UpdateTablePropertiesWithoutNodePath response.");
    assert.False(t, actual.Properties.AllowAutoFit, "Validate UpdateTablePropertiesWithoutNodePath response.");
    assert.True(t, actual.Properties.Bidi, "Validate UpdateTablePropertiesWithoutNodePath response.");
    assert.Equal(t, 1.0, actual.Properties.BottomPadding, "Validate UpdateTablePropertiesWithoutNodePath response.");
    assert.Equal(t, 2.0, actual.Properties.CellSpacing, "Validate UpdateTablePropertiesWithoutNodePath response.");
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

    request := &models.GetTableRowRequest{
        Name: ToStringPointer(remoteFileName),
        TablePath: ToStringPointer("tables/1"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetTableRow(ctx, request)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Row, "Validate GetTableRow response.");
    assert.NotNil(t, actual.Row.TableCellList, "Validate GetTableRow response.");
    assert.Equal(t, 2, len(actual.Row.TableCellList), "Validate GetTableRow response.");
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

    request := &models.DeleteTableRowRequest{
        Name: ToStringPointer(remoteFileName),
        TablePath: ToStringPointer("tables/1"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

_, err := client.WordsApi.DeleteTableRow(ctx, request)

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
        ColumnsCount: ToInt32Pointer(int32(5)),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.InsertTableRowRequest{
        Name: ToStringPointer(remoteFileName),
        TablePath: ToStringPointer("sections/0/tables/2"),
        Row: requestRow,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.InsertTableRow(ctx, request)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Row, "Validate InsertTableRow response.");
    assert.NotNil(t, actual.Row.TableCellList, "Validate InsertTableRow response.");
    assert.Equal(t, 5, len(actual.Row.TableCellList), "Validate InsertTableRow response.");
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

    request := &models.GetTableRowFormatRequest{
        Name: ToStringPointer(remoteFileName),
        TablePath: ToStringPointer("sections/0/tables/2"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetTableRowFormat(ctx, request)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.RowFormat, "Validate GetTableRowFormat response.");
    assert.True(t, actual.RowFormat.AllowBreakAcrossPages, "Validate GetTableRowFormat response.");
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
        AllowBreakAcrossPages: ToBoolPointer(true),
        HeadingFormat: ToBoolPointer(true),
        Height: ToFloat64Pointer(10.0),
        HeightRule: ToStringPointer("Exactly"),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.UpdateTableRowFormatRequest{
        Name: ToStringPointer(remoteFileName),
        TablePath: ToStringPointer("sections/0/tables/2"),
        Index: ToInt32Pointer(int32(0)),
        Format: requestFormat,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.UpdateTableRowFormat(ctx, request)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.RowFormat, "Validate UpdateTableRowFormat response.");
    assert.True(t, actual.RowFormat.AllowBreakAcrossPages, "Validate UpdateTableRowFormat response.");
    assert.True(t, actual.RowFormat.HeadingFormat, "Validate UpdateTableRowFormat response.");
    assert.Equal(t, 10.0, actual.RowFormat.Height, "Validate UpdateTableRowFormat response.");
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

    request := &models.GetTableCellRequest{
        Name: ToStringPointer(remoteFileName),
        TableRowPath: ToStringPointer("sections/0/tables/2/rows/0"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetTableCell(ctx, request)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Cell, "Validate GetTableCell response.");
    assert.Equal(t, "0.0.5.0.0", actual.Cell.NodeId, "Validate GetTableCell response.");
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

    request := &models.DeleteTableCellRequest{
        Name: ToStringPointer(remoteFileName),
        TableRowPath: ToStringPointer("sections/0/tables/2/rows/0"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

_, err := client.WordsApi.DeleteTableCell(ctx, request)

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

    request := &models.InsertTableCellRequest{
        Name: ToStringPointer(remoteFileName),
        TableRowPath: ToStringPointer("sections/0/tables/2/rows/0"),
        Cell: requestCell,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.InsertTableCell(ctx, request)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Cell, "Validate InsertTableCell response.");
    assert.Equal(t, "0.0.5.0.3", actual.Cell.NodeId, "Validate InsertTableCell response.");
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

    request := &models.GetTableCellFormatRequest{
        Name: ToStringPointer(remoteFileName),
        TableRowPath: ToStringPointer("sections/0/tables/2/rows/0"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetTableCellFormat(ctx, request)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.CellFormat, "Validate GetTableCellFormat response.");
    assert.True(t, actual.CellFormat.WrapText, "Validate GetTableCellFormat response.");
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
        BottomPadding: ToFloat64Pointer(5.0),
        FitText: ToBoolPointer(true),
        HorizontalMerge: ToStringPointer("First"),
        WrapText: ToBoolPointer(true),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.UpdateTableCellFormatRequest{
        Name: ToStringPointer(remoteFileName),
        TableRowPath: ToStringPointer("sections/0/tables/2/rows/0"),
        Index: ToInt32Pointer(int32(0)),
        Format: requestFormat,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.UpdateTableCellFormat(ctx, request)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.CellFormat, "Validate UpdateTableCellFormat response.");
    assert.Equal(t, 5.0, actual.CellFormat.BottomPadding, "Validate UpdateTableCellFormat response.");
    assert.True(t, actual.CellFormat.FitText, "Validate UpdateTableCellFormat response.");
    assert.True(t, actual.CellFormat.WrapText, "Validate UpdateTableCellFormat response.");
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

    request := &models.RenderTableRequest{
        Name: ToStringPointer(remoteFileName),
        Format: ToStringPointer("png"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

_, err := client.WordsApi.RenderTable(ctx, request)

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

    request := &models.RenderTableRequest{
        Name: ToStringPointer(remoteFileName),
        Format: ToStringPointer("png"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

_, err := client.WordsApi.RenderTable(ctx, request)

    if err != nil {
        t.Error(err)
    }

}

