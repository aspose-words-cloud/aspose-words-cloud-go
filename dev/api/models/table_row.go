/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="table_row.go">
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

// DTO container with a table row element.

type ITableRow interface {
    IsTableRow() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetNodeId() *string
    SetNodeId(value *string)
    GetTableCellList() []ITableCell
    SetTableCellList(value []ITableCell)
    GetRowFormat() ITableRowFormat
    SetRowFormat(value ITableRowFormat)
}

type TableRow struct {
    // DTO container with a table row element.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with a table row element.
    NodeId *string `json:"NodeId,omitempty"`

    // DTO container with a table row element.
    TableCellList []ITableCell `json:"TableCellList,omitempty"`

    // DTO container with a table row element.
    RowFormat ITableRowFormat `json:"RowFormat,omitempty"`
}

func (TableRow) IsTableRow() bool {
    return true
}

func (TableRow) IsNodeLink() bool {
    return true
}

func (TableRow) IsLinkElement() bool {
    return true
}

func (obj *TableRow) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.TableCellList != nil) {
        for _, objElementTableCellList := range obj.TableCellList {
            objElementTableCellList.Initialize()
        }
    }
    if (obj.RowFormat != nil) {
        obj.RowFormat.Initialize()
    }


}

func (obj *TableRow) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Link"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IWordsApiLink = new(WordsApiLink)
            modelInstance.Deserialize(parsedValue)
            obj.Link = modelInstance
        }

    } else if jsonValue, exists := json["link"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IWordsApiLink = new(WordsApiLink)
            modelInstance.Deserialize(parsedValue)
            obj.Link = modelInstance
        }

    }

    if jsonValue, exists := json["NodeId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NodeId = &parsedValue
        }

    } else if jsonValue, exists := json["nodeId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NodeId = &parsedValue
        }

    }

    if jsonValue, exists := json["TableCellList"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.TableCellList = make([]ITableCell, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance ITableCell = new(TableCell)
                    modelElementInstance.Deserialize(elementValue)
                    obj.TableCellList = append(obj.TableCellList, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["tableCellList"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.TableCellList = make([]ITableCell, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance ITableCell = new(TableCell)
                    modelElementInstance.Deserialize(elementValue)
                    obj.TableCellList = append(obj.TableCellList, modelElementInstance)
                }

            }
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

func (obj *TableRow) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *TableRow) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Link != nil {
        if err := obj.Link.Validate(); err != nil {
            return err
        }
    }
    if obj.TableCellList != nil {
        for _, elementTableCellList := range obj.TableCellList {
            if elementTableCellList != nil {
                if err := elementTableCellList.Validate(); err != nil {
                    return err
                }
            }
        }
    }
    if obj.RowFormat != nil {
        if err := obj.RowFormat.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *TableRow) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *TableRow) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *TableRow) GetNodeId() *string {
    return obj.NodeId
}

func (obj *TableRow) SetNodeId(value *string) {
    obj.NodeId = value
}

func (obj *TableRow) GetTableCellList() []ITableCell {
    return obj.TableCellList
}

func (obj *TableRow) SetTableCellList(value []ITableCell) {
    obj.TableCellList = value
}

func (obj *TableRow) GetRowFormat() ITableRowFormat {
    return obj.RowFormat
}

func (obj *TableRow) SetRowFormat(value ITableRowFormat) {
    obj.RowFormat = value
}

