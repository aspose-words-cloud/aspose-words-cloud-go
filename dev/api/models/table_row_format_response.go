/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="table_row_format_response.go">
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

import (
    "errors"
)

// The REST response with the formatting properties of a table row.
// This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/tables/{0}/rows/{1}/rowformat" REST API requests.

type ITableRowFormatResponse interface {
    IsTableRowFormatResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetRowFormat() ITableRowFormat
    SetRowFormat(value ITableRowFormat)
}

type TableRowFormatResponse struct {
    // The REST response with the formatting properties of a table row.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/tables/{0}/rows/{1}/rowformat" REST API requests.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with the formatting properties of a table row.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/tables/{0}/rows/{1}/rowformat" REST API requests.
    RowFormat ITableRowFormat `json:"RowFormat,omitempty"`
}

func (TableRowFormatResponse) IsTableRowFormatResponse() bool {
    return true
}

func (TableRowFormatResponse) IsWordsResponse() bool {
    return true
}

func (obj *TableRowFormatResponse) Initialize() {
    if (obj.RowFormat != nil) {
        obj.RowFormat.Initialize()
    }


}

func (obj *TableRowFormatResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["RowFormat"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ITableRowFormat = new(TableRowFormat)
            modelInstance.Deserialize(parsedValue)
            obj.RowFormat = modelInstance
        }

    } else if jsonValue, exists := json["rowFormat"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ITableRowFormat = new(TableRowFormat)
            modelInstance.Deserialize(parsedValue)
            obj.RowFormat = modelInstance
        }

    }
}

func (obj *TableRowFormatResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *TableRowFormatResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.RowFormat != nil {
        if err := obj.RowFormat.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *TableRowFormatResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *TableRowFormatResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *TableRowFormatResponse) GetRowFormat() ITableRowFormat {
    return obj.RowFormat
}

func (obj *TableRowFormatResponse) SetRowFormat(value ITableRowFormat) {
    obj.RowFormat = value
}

