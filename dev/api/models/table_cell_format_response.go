/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="table_cell_format_response.go">
 *   Copyright (c) 2023 Aspose.Words for Cloud
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

// The REST response with the formatting properties of a table cell.
// This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/tables/{0}/rows/{1}/cells/{2}/cellformat" REST API requests.

type ITableCellFormatResponse interface {
    IsTableCellFormatResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetCellFormat() ITableCellFormat
    SetCellFormat(value ITableCellFormat)
}

type TableCellFormatResponse struct {
    // The REST response with the formatting properties of a table cell.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/tables/{0}/rows/{1}/cells/{2}/cellformat" REST API requests.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with the formatting properties of a table cell.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/tables/{0}/rows/{1}/cells/{2}/cellformat" REST API requests.
    CellFormat ITableCellFormat `json:"CellFormat,omitempty"`
}

func (TableCellFormatResponse) IsTableCellFormatResponse() bool {
    return true
}

func (TableCellFormatResponse) IsWordsResponse() bool {
    return true
}

func (obj *TableCellFormatResponse) Initialize() {
    if (obj.CellFormat != nil) {
        obj.CellFormat.Initialize()
    }


}

func (obj *TableCellFormatResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["CellFormat"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ITableCellFormat = new(TableCellFormat)
            modelInstance.Deserialize(parsedValue)
            obj.CellFormat = modelInstance
        }

    } else if jsonValue, exists := json["cellFormat"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ITableCellFormat = new(TableCellFormat)
            modelInstance.Deserialize(parsedValue)
            obj.CellFormat = modelInstance
        }

    }
}

func (obj *TableCellFormatResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *TableCellFormatResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.CellFormat != nil {
        if err := obj.CellFormat.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *TableCellFormatResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *TableCellFormatResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *TableCellFormatResponse) GetCellFormat() ITableCellFormat {
    return obj.CellFormat
}

func (obj *TableCellFormatResponse) SetCellFormat(value ITableCellFormat) {
    obj.CellFormat = value
}

