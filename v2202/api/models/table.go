/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="table.go">
 *   Copyright (c) 2022 Aspose.Words for Cloud
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

// DTO container with a table element.
type TableResult struct {
    // DTO container with a table element.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // DTO container with a table element.
    NodeId string `json:"NodeId,omitempty"`

    // DTO container with a table element.
    TableProperties TablePropertiesResult `json:"TableProperties,omitempty"`

    // DTO container with a table element.
    TableRowList []TableRowResult `json:"TableRowList,omitempty"`
}

type Table struct {
    // DTO container with a table element.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with a table element.
    NodeId *string `json:"NodeId,omitempty"`

    // DTO container with a table element.
    TableProperties ITableProperties `json:"TableProperties,omitempty"`

    // DTO container with a table element.
    TableRowList []TableRow `json:"TableRowList,omitempty"`
}

type ITable interface {
    IsTable() bool
    Initialize()
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





    if (obj.TableProperties != nil) {
        obj.TableProperties.Initialize()
    }



}


