/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="table_row_insert_dto.go">
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

package models

// DTO container with a table row element.
type TableRowInsertDtoResult struct {
    // DTO container with a table row element.
    ColumnsCount int32 `json:"ColumnsCount,omitempty"`

    // DTO container with a table row element.
    InsertAfter int32 `json:"InsertAfter,omitempty"`
}

type TableRowInsertDto struct {
    // DTO container with a table row element.
    ColumnsCount *int32 `json:"ColumnsCount,omitempty"`

    // DTO container with a table row element.
    InsertAfter *int32 `json:"InsertAfter,omitempty"`
}

type ITableRowInsertDto interface {
    IsTableRowInsertDto() bool
}
func (TableRowInsertDto) IsTableRowInsertDto() bool {
    return true
}


