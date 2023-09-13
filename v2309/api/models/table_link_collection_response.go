/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="table_link_collection_response.go">
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

// The REST response with a collection of tables.

type ITableLinkCollectionResponse interface {
    IsTableLinkCollectionResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetRequestId() *string
    SetRequestId(value *string)
    GetTables() ITableLinkCollection
    SetTables(value ITableLinkCollection)
}

type TableLinkCollectionResponse struct {
    // The REST response with a collection of tables.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a collection of tables.
    Tables ITableLinkCollection `json:"Tables,omitempty"`
}

func (TableLinkCollectionResponse) IsTableLinkCollectionResponse() bool {
    return true
}

func (TableLinkCollectionResponse) IsWordsResponse() bool {
    return true
}

func (obj *TableLinkCollectionResponse) Initialize() {
    if (obj.Tables != nil) {
        obj.Tables.Initialize()
    }


}

func (obj *TableLinkCollectionResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["Tables"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ITableLinkCollection = new(TableLinkCollection)
            modelInstance.Deserialize(parsedValue)
            obj.Tables = modelInstance
        }

    } else if jsonValue, exists := json["tables"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ITableLinkCollection = new(TableLinkCollection)
            modelInstance.Deserialize(parsedValue)
            obj.Tables = modelInstance
        }

    }
}

func (obj *TableLinkCollectionResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *TableLinkCollectionResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *TableLinkCollectionResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *TableLinkCollectionResponse) GetTables() ITableLinkCollection {
    return obj.Tables
}

func (obj *TableLinkCollectionResponse) SetTables(value ITableLinkCollection) {
    obj.Tables = value
}

