/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="table_cell_response.go">
 *   Copyright (c) 2024 Aspose.Words for Cloud
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

import (
    "errors"
)

// The REST response with a table cell.
// This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/tables/{0}" REST API requests.

type ITableCellResponse interface {
    IsTableCellResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetCell() ITableCell
    SetCell(value ITableCell)
}

type TableCellResponse struct {
    // The REST response with a table cell.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/tables/{0}" REST API requests.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a table cell.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/tables/{0}" REST API requests.
    Cell ITableCell `json:"Cell,omitempty"`
}

func (TableCellResponse) IsTableCellResponse() bool {
    return true
}

func (TableCellResponse) IsWordsResponse() bool {
    return true
}

func (obj *TableCellResponse) Initialize() {
    if (obj.Cell != nil) {
        obj.Cell.Initialize()
    }


}

func (obj *TableCellResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["Cell"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ITableCell = new(TableCell)
            modelInstance.Deserialize(parsedValue)
            obj.Cell = modelInstance
        }

    } else if jsonValue, exists := json["cell"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ITableCell = new(TableCell)
            modelInstance.Deserialize(parsedValue)
            obj.Cell = modelInstance
        }

    }
}

func (obj *TableCellResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *TableCellResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Cell != nil {
        if err := obj.Cell.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *TableCellResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *TableCellResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *TableCellResponse) GetCell() ITableCell {
    return obj.Cell
}

func (obj *TableCellResponse) SetCell(value ITableCell) {
    obj.Cell = value
}

