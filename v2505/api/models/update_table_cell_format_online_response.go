/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="update_table_cell_format_online_response.go">
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

package models

import "io"


// UpdateTableCellFormatOnlineResponse struct
// Updates the formatting properties of a cell in the table row.
type UpdateTableCellFormatOnlineResponse struct {
    // The REST response with the formatting properties of a table cell.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/tables/{0}/rows/{1}/cells/{2}/cellformat" REST API requests.
    Model ITableCellFormatResponse


    // The document after modification.
    Document map[string]io.Reader
}

func (obj *UpdateTableCellFormatOnlineResponse) GetModel() ITableCellFormatResponse {
    return obj.Model
}

func (obj *UpdateTableCellFormatOnlineResponse) SetModel(value ITableCellFormatResponse) {
    obj.Model = value
}

func (obj *UpdateTableCellFormatOnlineResponse) GetDocument() map[string]io.Reader {
    return obj.Document
}

func (obj *UpdateTableCellFormatOnlineResponse) SetDocument(value map[string]io.Reader) {
    obj.Document = value
}