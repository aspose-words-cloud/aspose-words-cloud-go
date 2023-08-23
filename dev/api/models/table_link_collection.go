/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="table_link_collection.go">
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

// The collection of table's links.

type ITableLinkCollection interface {
    IsTableLinkCollection() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetTableLinkList() []ITableLink
    SetTableLinkList(value []ITableLink)
}

type TableLinkCollection struct {
    // The collection of table's links.
    Link IWordsApiLink

    // The collection of table's links.
    TableLinkList []ITableLink
}

func (TableLinkCollection) IsTableLinkCollection() bool {
    return true
}

func (TableLinkCollection) IsLinkElement() bool {
    return true
}

func (obj *TableLinkCollection) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.TableLinkList != nil) {
        for _, objElementTableLinkList := range obj.TableLinkList {
            objElementTableLinkList.Initialize()
        }
    }

}

func (obj *TableLinkCollection) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["TableLinkList"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance ITableLink = new(TableLink)
                    modelElementInstance.Deserialize(elementValue)
                    obj.TableLinkList = append(obj.TableLinkList, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["tableLinkList"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance ITableLink = new(TableLink)
                    modelElementInstance.Deserialize(elementValue)
                    obj.TableLinkList = append(obj.TableLinkList, modelElementInstance)
                }

            }
        }

    }
}

func (obj *TableLinkCollection) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *TableLinkCollection) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *TableLinkCollection) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *TableLinkCollection) GetTableLinkList() []ITableLink {
    return obj.TableLinkList
}

func (obj *TableLinkCollection) SetTableLinkList(value []ITableLink) {
    obj.TableLinkList = value
}

