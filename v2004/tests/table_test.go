//
// MIT License

// Copyright (c) 2019 Aspose Pty Ltd

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

package api_test

import (
	"path"
	"path/filepath"
	"testing"

	"github.com/aspose-words-cloud/aspose-words-cloud-go/v2004/api/models"
)

func TestDeleteTable(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestDeleteTable.docx"
	nodePath := "sections/0"
	index := 1
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteTable(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteTableCell(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestDeleteTableCell.docx"
	tableRowPath := "sections/0/tables/2/rows/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteTableCell(ctx, remoteName, tableRowPath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteTableRow(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestDeleteTableRow.docx"
	tablePath := "tables/1"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteTableRow(ctx, remoteName, tablePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteTableWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestDeleteTableWithoutNodePath.docx"
	index := 1
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteTableWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetTable(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestGetTable.docx"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetTable(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetTableCell(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	tableRowPath := "sections/0/tables/2/rows/0"
	index := 0
	remoteName := "TestGetTableCell.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetTableCell(ctx, remoteName, tableRowPath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetTableCellFormat(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	tableRowPath := "sections/0/tables/2/rows/0"
	index := 0
	remoteName := "TestGetTableCellFormat.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetTableCellFormat(ctx, remoteName, tableRowPath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetTableProperties(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	nodePath := "sections/0"
	index := 0
	remoteName := "TestGetTableProperties.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetTableProperties(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetTablePropertiesWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestGetTablePropertiesWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetTablePropertiesWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetTableRow(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestGetTableRow.docx"
	tablePath := "tables/1"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetTableRow(ctx, remoteName, tablePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetTableRowFormat(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestGetTableRowFormat.docx"
	tablePath := "tables/1"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetTableRowFormat(ctx, remoteName, tablePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetTableWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestGetTableWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetTableWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetTables(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestGetTables.docx"
	nodePath := "sections/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetTables(ctx, remoteName, nodePath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetTablesWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestGetTablesWithoutNodePath.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetTablesWithoutNodePath(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertTable(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestInsertTable.docx"
	nodePath := "sections/0"
	table := models.TableInsert{
		ColumnsCount: 5,
		RowsCount:    4,
	}

	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertTable(ctx, remoteName, nodePath, table, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertTableCell(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestInsertTableCell.docx"
	tableRowPath := "sections/0/tables/2/rows/0"
	cell := models.TableCellInsert{}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertTableCell(ctx, remoteName, tableRowPath, cell, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertTableRow(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	tablePath := "sections/0/tables/2"
	remoteName := "TestInsertTableRow.docx"
	row := models.TableRowInsert{
		ColumnsCount: 5,
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertTableRow(ctx, remoteName, tablePath, row, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertTableWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestInsertTableWithoutNodePath.docx"
	table := models.TableInsert{
		ColumnsCount: 5,
		RowsCount:    4,
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertTableWithoutNodePath(ctx, remoteName, table, options)

	if err != nil {
		t.Error(err)
	}
}

func TestRenderTable(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestRenderTable.docx"
	format := "png"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	output, err := client.WordsApi.RenderTable(ctx, remoteName, format, nodePath, int32(index), options)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}

func TestRenderTableWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestRenderTableWithoutNodePath.docx"
	format := "png"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	output, err := client.WordsApi.RenderTableWithoutNodePath(ctx, remoteName, format, int32(index), options)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateTableCellFormat(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestUpdateTableCellFormat.docx"
	tableRowPath := "sections/0/tables/2/rows/0"
	index := 0
	format := models.TableCellFormat{
		BottomPadding:   5,
		FitText:         true,
		HorizontalMerge: "First",
		WrapText:        true,
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateTableCellFormat(ctx, remoteName, tableRowPath, int32(index), format, options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateTableProperties(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestUpdateTableProperties.docx"
	nodePath := "sections/0"
	index := 0
	tableProperties := models.TableProperties{
		Alignment:     "Right",
		AllowAutoFit:  false,
		Bidi:          true,
		BottomPadding: 1,
		CellSpacing:   2,
		LeftIndent:    3,
		LeftPadding:   4,
		RightPadding:  5,
		StyleOptions:  "ColumnBands",
		TopPadding:    6,
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateTableProperties(ctx, remoteName, nodePath, int32(index), tableProperties, options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateTablePropertiesWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestUpdateTablePropertiesWithoutNodePath.docx"
	index := 0
	properties := models.TableProperties{
		Alignment:     "Right",
		AllowAutoFit:  false,
		Bidi:          true,
		BottomPadding: 1,
		CellSpacing:   2,
		LeftIndent:    3,
		LeftPadding:   4,
		RightPadding:  5,
		StyleOptions:  "ColumnBands",
		TopPadding:    6,
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}
	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateTablePropertiesWithoutNodePath(ctx, remoteName, int32(index), properties, options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateTableRowFormat(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestUpdateTableRowFormat.docx"
	tablePath := "sections/0/tables/2"
	index := 0
	format := models.TableRowFormat{
		AllowBreakAcrossPages: true,
		HeadingFormat:         true,
		Height:                10,
		HeightRule:            "Auto",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateTableRowFormat(ctx, remoteName, tablePath, int32(index), format, options)

	if err != nil {
		t.Error(err)
	}
}
