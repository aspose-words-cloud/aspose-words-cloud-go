/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="table.go">
 *   Copyright (c) 2026 Aspose.Words for Cloud
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

// DTO container with a table element.

type ITable interface {
    IsTable() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetNodeId() *string
    SetNodeId(value *string)
    GetTableRowList() []ITableRow
    SetTableRowList(value []ITableRow)
    GetTableProperties() ITableProperties
    SetTableProperties(value ITableProperties)
}

type Table struct {
    // DTO container with a table element.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with a table element.
    NodeId *string `json:"NodeId,omitempty"`

    // DTO container with a table element.
    TableRowList []ITableRow `json:"TableRowList,omitempty"`

    // DTO container with a table element.
    TableProperties ITableProperties `json:"TableProperties,omitempty"`
}

func (Table) IsTable() bool {
    return true
}

func (Table) IsNodeLink() bool {
    return true
}

func (Table) IsLinkElement() bool {
    return true
}

func (obj *Table) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.TableRowList != nil) {
        for _, objElementTableRowList := range obj.TableRowList {
            objElementTableRowList.Initialize()
        }
    }
    if (obj.TableProperties != nil) {
        obj.TableProperties.Initialize()
    }


}

func (obj *Table) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["TableRowList"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.TableRowList = make([]ITableRow, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance ITableRow = new(TableRow)
                    modelElementInstance.Deserialize(elementValue)
                    obj.TableRowList = append(obj.TableRowList, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["tableRowList"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.TableRowList = make([]ITableRow, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance ITableRow = new(TableRow)
                    modelElementInstance.Deserialize(elementValue)
                    obj.TableRowList = append(obj.TableRowList, modelElementInstance)
                }

            }
        }

    }

    if jsonValue, exists := json["TableProperties"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ITableProperties = new(TableProperties)
            modelInstance.Deserialize(parsedValue)
            obj.TableProperties = modelInstance
        }

    } else if jsonValue, exists := json["tableProperties"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ITableProperties = new(TableProperties)
            modelInstance.Deserialize(parsedValue)
            obj.TableProperties = modelInstance
        }

    }
}

func (obj *Table) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *Table) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Link != nil {
        if err := obj.Link.Validate(); err != nil {
            return err
        }
    }
    if obj.TableRowList != nil {
        for _, elementTableRowList := range obj.TableRowList {
            if elementTableRowList != nil {
                if err := elementTableRowList.Validate(); err != nil {
                    return err
                }
            }
        }
    }
    if obj.TableProperties != nil {
        if err := obj.TableProperties.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *Table) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *Table) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *Table) GetNodeId() *string {
    return obj.NodeId
}

func (obj *Table) SetNodeId(value *string) {
    obj.NodeId = value
}

func (obj *Table) GetTableRowList() []ITableRow {
    return obj.TableRowList
}

func (obj *Table) SetTableRowList(value []ITableRow) {
    obj.TableRowList = value
}

func (obj *Table) GetTableProperties() ITableProperties {
    return obj.TableProperties
}

func (obj *Table) SetTableProperties(value ITableProperties) {
    obj.TableProperties = value
}

