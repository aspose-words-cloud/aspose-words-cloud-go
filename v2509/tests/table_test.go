/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="table_test.go">
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

// Example of how to work wtih table.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2509/api/models"
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

    assert.NotNil(t, actual.GetTables(), "Validate GetTables response.");
    assert.NotNil(t, actual.GetTables().GetTableLinkList(), "Validate GetTables response.");
    assert.Equal(t, 5, len(actual.GetTables().GetTableLinkList()), "Validate GetTables response.");
    assert.Equal(t, "0.0.1", DereferenceValue(actual.GetTables().GetTableLinkList()[0].GetNodeId()), "Validate GetTables response.");
}

// Test for getting tables online.
func Test_Table_GetTablesOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Tables/TablesGet.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
        "nodePath": "",
    }

    request := &models.GetTablesOnlineRequest{
        Document: requestDocument,
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetTablesOnline(ctx, request)
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

    request := &models.GetTablesRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetTables(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetTables(), "Validate GetTablesWithoutNodePath response.");
    assert.NotNil(t, actual.GetTables().GetTableLinkList(), "Validate GetTablesWithoutNodePath response.");
    assert.Equal(t, 5, len(actual.GetTables().GetTableLinkList()), "Validate GetTablesWithoutNodePath response.");
    assert.Equal(t, "0.0.1", DereferenceValue(actual.GetTables().GetTableLinkList()[0].GetNodeId()), "Validate GetTablesWithoutNodePath response.");
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

    assert.NotNil(t, actual.GetTable(), "Validate GetTable response.");
    assert.NotNil(t, actual.GetTable().GetTableRowList(), "Validate GetTable response.");
    assert.Equal(t, 1, len(actual.GetTable().GetTableRowList()), "Validate GetTable response.");
    assert.NotNil(t, actual.GetTable().GetTableRowList()[0].GetTableCellList(), "Validate GetTable response.");
    assert.Equal(t, 2, len(actual.GetTable().GetTableRowList()[0].GetTableCellList()), "Validate GetTable response.");
}

// Test for getting table online.
func Test_Table_GetTableOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Tables/TablesGet.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
        "nodePath": "",
    }

    request := &models.GetTableOnlineRequest{
        Document: requestDocument,
        Index: ToInt32Pointer(int32(1)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetTableOnline(ctx, request)
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

    request := &models.GetTableRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(1)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetTable(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetTable(), "Validate GetTableWithoutNodePath response.");
    assert.NotNil(t, actual.GetTable().GetTableRowList(), "Validate GetTableWithoutNodePath response.");
    assert.Equal(t, 1, len(actual.GetTable().GetTableRowList()), "Validate GetTableWithoutNodePath response.");
    assert.NotNil(t, actual.GetTable().GetTableRowList()[0].GetTableCellList(), "Validate GetTableWithoutNodePath response.");
    assert.Equal(t, 2, len(actual.GetTable().GetTableRowList()[0].GetTableCellList()), "Validate GetTableWithoutNodePath response.");
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

// Test for deleting table online.
func Test_Table_DeleteTableOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Tables/TablesGet.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
        "nodePath": "",
    }

    request := &models.DeleteTableOnlineRequest{
        Document: requestDocument,
        Index: ToInt32Pointer(int32(1)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.DeleteTableOnline(ctx, request)
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
        Table: &requestTable,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.InsertTable(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetTable(), "Validate InsertTable response.");
    assert.NotNil(t, actual.GetTable().GetTableRowList(), "Validate InsertTable response.");
    assert.Equal(t, 4, len(actual.GetTable().GetTableRowList()), "Validate InsertTable response.");
    assert.NotNil(t, actual.GetTable().GetTableRowList()[0].GetTableCellList(), "Validate InsertTable response.");
    assert.Equal(t, 5, len(actual.GetTable().GetTableRowList()[0].GetTableCellList()), "Validate InsertTable response.");
}

// Test for adding table online.
func Test_Table_InsertTableOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Tables/TablesGet.docx"

    requestDocument := OpenFile(t, localFile)
    requestTable := models.TableInsert{
        ColumnsCount: ToInt32Pointer(int32(5)),
        RowsCount: ToInt32Pointer(int32(4)),
    }

    options := map[string]interface{}{
        "nodePath": "",
    }

    request := &models.InsertTableOnlineRequest{
        Document: requestDocument,
        Table: &requestTable,
        Optionals: options,
    }

    _, _, err := client.WordsApi.InsertTableOnline(ctx, request)
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
        ColumnsCount: ToInt32Pointer(int32(5)),
        RowsCount: ToInt32Pointer(int32(4)),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.InsertTableRequest{
        Name: ToStringPointer(remoteFileName),
        Table: &requestTable,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.InsertTable(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetTable(), "Validate InsertTableWithoutNodePath response.");
    assert.NotNil(t, actual.GetTable().GetTableRowList(), "Validate InsertTableWithoutNodePath response.");
    assert.Equal(t, 4, len(actual.GetTable().GetTableRowList()), "Validate InsertTableWithoutNodePath response.");
    assert.NotNil(t, actual.GetTable().GetTableRowList()[0].GetTableCellList(), "Validate InsertTableWithoutNodePath response.");
    assert.Equal(t, 5, len(actual.GetTable().GetTableRowList()[0].GetTableCellList()), "Validate InsertTableWithoutNodePath response.");
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

    assert.NotNil(t, actual.GetProperties(), "Validate GetTableProperties response.");
    assert.Equal(t, "Table Grid", DereferenceValue(actual.GetProperties().GetStyleName()), "Validate GetTableProperties response.");
}

// Test for getting document properties online.
func Test_Table_GetTablePropertiesOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Tables/TablesGet.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
        "nodePath": "",
    }

    request := &models.GetTablePropertiesOnlineRequest{
        Document: requestDocument,
        Index: ToInt32Pointer(int32(1)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetTablePropertiesOnline(ctx, request)
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

    request := &models.GetTablePropertiesRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(1)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetTableProperties(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetProperties(), "Validate GetTablePropertiesWithoutNodePath response.");
    assert.Equal(t, "Table Grid", DereferenceValue(actual.GetProperties().GetStyleName()), "Validate GetTablePropertiesWithoutNodePath response.");
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
        Properties: &requestProperties,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.UpdateTableProperties(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetProperties(), "Validate UpdateTableProperties response.");
    assert.Equal(t, false, DereferenceValue(actual.GetProperties().GetAllowAutoFit()), "Validate UpdateTableProperties response.");
    assert.Equal(t, true, DereferenceValue(actual.GetProperties().GetBidi()), "Validate UpdateTableProperties response.");
    assert.Equal(t, 1.0, DereferenceValue(actual.GetProperties().GetBottomPadding()), "Validate UpdateTableProperties response.");
    assert.Equal(t, 2.0, DereferenceValue(actual.GetProperties().GetCellSpacing()), "Validate UpdateTableProperties response.");
}

// Test for updating table properties online.
func Test_Table_UpdateTablePropertiesOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Tables/TablesGet.docx"

    requestDocument := OpenFile(t, localFile)
    requestProperties := models.TableProperties{
        Alignment: ToStringPointer("Right"),
        AllowAutoFit: ToBoolPointer(false),
        Bidi: ToBoolPointer(true),
        BottomPadding: ToFloat64Pointer(1),
        CellSpacing: ToFloat64Pointer(2),
        StyleOptions: ToStringPointer("ColumnBands"),
    }

    options := map[string]interface{}{
        "nodePath": "",
    }

    request := &models.UpdateTablePropertiesOnlineRequest{
        Document: requestDocument,
        Properties: &requestProperties,
        Index: ToInt32Pointer(int32(1)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.UpdateTablePropertiesOnline(ctx, request)
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
        Properties: &requestProperties,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.UpdateTableProperties(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetProperties(), "Validate UpdateTablePropertiesWithoutNodePath response.");
    assert.Equal(t, false, DereferenceValue(actual.GetProperties().GetAllowAutoFit()), "Validate UpdateTablePropertiesWithoutNodePath response.");
    assert.Equal(t, true, DereferenceValue(actual.GetProperties().GetBidi()), "Validate UpdateTablePropertiesWithoutNodePath response.");
    assert.Equal(t, 1.0, DereferenceValue(actual.GetProperties().GetBottomPadding()), "Validate UpdateTablePropertiesWithoutNodePath response.");
    assert.Equal(t, 2.0, DereferenceValue(actual.GetProperties().GetCellSpacing()), "Validate UpdateTablePropertiesWithoutNodePath response.");
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

    assert.NotNil(t, actual.GetRow(), "Validate GetTableRow response.");
    assert.NotNil(t, actual.GetRow().GetTableCellList(), "Validate GetTableRow response.");
    assert.Equal(t, 2, len(actual.GetRow().GetTableCellList()), "Validate GetTableRow response.");
}

// Test for getting table row online.
func Test_Table_GetTableRowOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Tables/TablesGet.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.GetTableRowOnlineRequest{
        Document: requestDocument,
        TablePath: ToStringPointer("tables/1"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetTableRowOnline(ctx, request)
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

// Test for deleting table row online.
func Test_Table_DeleteTableRowOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Tables/TablesGet.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.DeleteTableRowOnlineRequest{
        Document: requestDocument,
        TablePath: ToStringPointer("tables/1"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.DeleteTableRowOnline(ctx, request)
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
        "nodePath": "sections/0/tables/2",
        "folder": remoteDataFolder,
    }

    request := &models.InsertTableRowRequest{
        Name: ToStringPointer(remoteFileName),
        Row: &requestRow,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.InsertTableRow(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetRow(), "Validate InsertTableRow response.");
    assert.NotNil(t, actual.GetRow().GetTableCellList(), "Validate InsertTableRow response.");
    assert.Equal(t, 5, len(actual.GetRow().GetTableCellList()), "Validate InsertTableRow response.");
}

// Test for adding row online.
func Test_Table_InsertTableRowOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Tables/TablesGet.docx"

    requestDocument := OpenFile(t, localFile)
    requestRow := models.TableRowInsert{
        ColumnsCount: ToInt32Pointer(int32(5)),
    }

    options := map[string]interface{}{
        "nodePath": "sections/0/tables/2",
    }

    request := &models.InsertTableRowOnlineRequest{
        Document: requestDocument,
        Row: &requestRow,
        Optionals: options,
    }

    _, _, err := client.WordsApi.InsertTableRowOnline(ctx, request)
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

    assert.NotNil(t, actual.GetRowFormat(), "Validate GetTableRowFormat response.");
    assert.Equal(t, true, DereferenceValue(actual.GetRowFormat().GetAllowBreakAcrossPages()), "Validate GetTableRowFormat response.");
}

// Test for getting row format online.
func Test_Table_GetTableRowFormatOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Tables/TablesGet.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.GetTableRowFormatOnlineRequest{
        Document: requestDocument,
        TablePath: ToStringPointer("sections/0/tables/2"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetTableRowFormatOnline(ctx, request)
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
        Format: &requestFormat,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.UpdateTableRowFormat(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetRowFormat(), "Validate UpdateTableRowFormat response.");
    assert.Equal(t, true, DereferenceValue(actual.GetRowFormat().GetAllowBreakAcrossPages()), "Validate UpdateTableRowFormat response.");
    assert.Equal(t, true, DereferenceValue(actual.GetRowFormat().GetHeadingFormat()), "Validate UpdateTableRowFormat response.");
    assert.Equal(t, 10.0, DereferenceValue(actual.GetRowFormat().GetHeight()), "Validate UpdateTableRowFormat response.");
}

// Test updating row format online.
func Test_Table_UpdateTableRowFormatOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Tables/TablesGet.docx"

    requestDocument := OpenFile(t, localFile)
    requestFormat := models.TableRowFormat{
        AllowBreakAcrossPages: ToBoolPointer(true),
        HeadingFormat: ToBoolPointer(true),
        Height: ToFloat64Pointer(10),
        HeightRule: ToStringPointer("Auto"),
    }

    options := map[string]interface{}{
    }

    request := &models.UpdateTableRowFormatOnlineRequest{
        Document: requestDocument,
        TablePath: ToStringPointer("sections/0/tables/2"),
        Format: &requestFormat,
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.UpdateTableRowFormatOnline(ctx, request)
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

    assert.NotNil(t, actual.GetCell(), "Validate GetTableCell response.");
    assert.Equal(t, "0.0.5.0.0", DereferenceValue(actual.GetCell().GetNodeId()), "Validate GetTableCell response.");
}

// Test for getting table cell online.
func Test_Table_GetTableCellOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Tables/TablesGet.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.GetTableCellOnlineRequest{
        Document: requestDocument,
        TableRowPath: ToStringPointer("sections/0/tables/2/rows/0"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetTableCellOnline(ctx, request)
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

// Test for deleting cell online.
func Test_Table_DeleteTableCellOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Tables/TablesGet.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.DeleteTableCellOnlineRequest{
        Document: requestDocument,
        TableRowPath: ToStringPointer("sections/0/tables/2/rows/0"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.DeleteTableCellOnline(ctx, request)
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
        "tableRowPath": "sections/0/tables/2/rows/0",
        "folder": remoteDataFolder,
    }

    request := &models.InsertTableCellRequest{
        Name: ToStringPointer(remoteFileName),
        Cell: &requestCell,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.InsertTableCell(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetCell(), "Validate InsertTableCell response.");
    assert.Equal(t, "0.0.5.0.3", DereferenceValue(actual.GetCell().GetNodeId()), "Validate InsertTableCell response.");
}

// Test for adding cell online.
func Test_Table_InsertTableCellOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Tables/TablesGet.docx"

    requestDocument := OpenFile(t, localFile)
    requestCell := models.TableCellInsert{
    }

    options := map[string]interface{}{
        "tableRowPath": "sections/0/tables/2/rows/0",
    }

    request := &models.InsertTableCellOnlineRequest{
        Document: requestDocument,
        Cell: &requestCell,
        Optionals: options,
    }

    _, _, err := client.WordsApi.InsertTableCellOnline(ctx, request)
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

    assert.NotNil(t, actual.GetCellFormat(), "Validate GetTableCellFormat response.");
    assert.Equal(t, true, DereferenceValue(actual.GetCellFormat().GetWrapText()), "Validate GetTableCellFormat response.");
}

// Test for getting cell format online.
func Test_Table_GetTableCellFormatOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Tables/TablesGet.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.GetTableCellFormatOnlineRequest{
        Document: requestDocument,
        TableRowPath: ToStringPointer("sections/0/tables/2/rows/0"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetTableCellFormatOnline(ctx, request)
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
        Format: &requestFormat,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.UpdateTableCellFormat(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetCellFormat(), "Validate UpdateTableCellFormat response.");
    assert.Equal(t, 5.0, DereferenceValue(actual.GetCellFormat().GetBottomPadding()), "Validate UpdateTableCellFormat response.");
    assert.Equal(t, true, DereferenceValue(actual.GetCellFormat().GetFitText()), "Validate UpdateTableCellFormat response.");
    assert.Equal(t, true, DereferenceValue(actual.GetCellFormat().GetWrapText()), "Validate UpdateTableCellFormat response.");
}

// Test for updating cell format online.
func Test_Table_UpdateTableCellFormatOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Tables/TablesGet.docx"

    requestDocument := OpenFile(t, localFile)
    requestFormat := models.TableCellFormat{
        BottomPadding: ToFloat64Pointer(5),
        FitText: ToBoolPointer(true),
        HorizontalMerge: ToStringPointer("First"),
        WrapText: ToBoolPointer(true),
    }

    options := map[string]interface{}{
    }

    request := &models.UpdateTableCellFormatOnlineRequest{
        Document: requestDocument,
        TableRowPath: ToStringPointer("sections/0/tables/2/rows/0"),
        Format: &requestFormat,
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.UpdateTableCellFormatOnline(ctx, request)
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

// Test for table rendering.
func Test_Table_RenderTableOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Tables/TablesGet.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
        "nodePath": "",
    }

    request := &models.RenderTableOnlineRequest{
        Document: requestDocument,
        Format: ToStringPointer("png"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, err := client.WordsApi.RenderTableOnline(ctx, request)
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
