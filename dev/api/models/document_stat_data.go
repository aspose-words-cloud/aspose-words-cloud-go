/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="document_stat_data.go">
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

// Container for the document's statistical data.

type IDocumentStatData interface {
    IsDocumentStatData() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetFootnotesStatData() IFootnotesStatData
    SetFootnotesStatData(value IFootnotesStatData)
    GetPageCount() *int32
    SetPageCount(value *int32)
    GetParagraphCount() *int32
    SetParagraphCount(value *int32)
    GetWordCount() *int32
    SetWordCount(value *int32)
    GetPageStatData() []IPageStatData
    SetPageStatData(value []IPageStatData)
}

type DocumentStatData struct {
    // Container for the document's statistical data.
    FootnotesStatData IFootnotesStatData

    // Container for the document's statistical data.
    PageCount *int32

    // Container for the document's statistical data.
    ParagraphCount *int32

    // Container for the document's statistical data.
    WordCount *int32

    // Container for the document's statistical data.
    PageStatData []IPageStatData
}

func (DocumentStatData) IsDocumentStatData() bool {
    return true
}


func (obj *DocumentStatData) Initialize() {
    if (obj.FootnotesStatData != nil) {
        obj.FootnotesStatData.Initialize()
    }

    if (obj.PageStatData != nil) {
        for _, objElementPageStatData := range obj.PageStatData {
            objElementPageStatData.Initialize()
        }
    }

}

func (obj *DocumentStatData) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["FootnotesStatData"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFootnotesStatData = new(FootnotesStatData)
            modelInstance.Deserialize(parsedValue)
            obj.FootnotesStatData = modelInstance
        }

    } else if jsonValue, exists := json["footnotesStatData"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFootnotesStatData = new(FootnotesStatData)
            modelInstance.Deserialize(parsedValue)
            obj.FootnotesStatData = modelInstance
        }

    }

    if jsonValue, exists := json["PageCount"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.PageCount = new(int32)
            *obj.PageCount = int32(parsedValue)
        }

    } else if jsonValue, exists := json["pageCount"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.PageCount = new(int32)
            *obj.PageCount = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["ParagraphCount"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ParagraphCount = new(int32)
            *obj.ParagraphCount = int32(parsedValue)
        }

    } else if jsonValue, exists := json["paragraphCount"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ParagraphCount = new(int32)
            *obj.ParagraphCount = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["WordCount"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.WordCount = new(int32)
            *obj.WordCount = int32(parsedValue)
        }

    } else if jsonValue, exists := json["wordCount"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.WordCount = new(int32)
            *obj.WordCount = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["PageStatData"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IPageStatData = new(PageStatData)
                    modelElementInstance.Deserialize(elementValue)
                    obj.PageStatData = append(obj.PageStatData, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["pageStatData"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IPageStatData = new(PageStatData)
                    modelElementInstance.Deserialize(elementValue)
                    obj.PageStatData = append(obj.PageStatData, modelElementInstance)
                }

            }
        }

    }
}

func (obj *DocumentStatData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *DocumentStatData) GetFootnotesStatData() IFootnotesStatData {
    return obj.FootnotesStatData
}

func (obj *DocumentStatData) SetFootnotesStatData(value IFootnotesStatData) {
    obj.FootnotesStatData = value
}

func (obj *DocumentStatData) GetPageCount() *int32 {
    return obj.PageCount
}

func (obj *DocumentStatData) SetPageCount(value *int32) {
    obj.PageCount = value
}

func (obj *DocumentStatData) GetParagraphCount() *int32 {
    return obj.ParagraphCount
}

func (obj *DocumentStatData) SetParagraphCount(value *int32) {
    obj.ParagraphCount = value
}

func (obj *DocumentStatData) GetWordCount() *int32 {
    return obj.WordCount
}

func (obj *DocumentStatData) SetWordCount(value *int32) {
    obj.WordCount = value
}

func (obj *DocumentStatData) GetPageStatData() []IPageStatData {
    return obj.PageStatData
}

func (obj *DocumentStatData) SetPageStatData(value []IPageStatData) {
    obj.PageStatData = value
}

