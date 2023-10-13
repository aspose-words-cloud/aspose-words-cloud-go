/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="compare_data.go">
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

// Container class for compare documents.

type ICompareData interface {
    IsCompareData() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetAuthor() *string
    SetAuthor(value *string)
    GetCompareOptions() ICompareOptions
    SetCompareOptions(value ICompareOptions)
    GetComparingWithDocument() *string
    SetComparingWithDocument(value *string)
    GetDateTime() *Time
    SetDateTime(value *Time)
    GetResultDocumentFormat() *string
    SetResultDocumentFormat(value *string)
}

type CompareData struct {
    // Container class for compare documents.
    Author *string `json:"Author,omitempty"`

    // Container class for compare documents.
    CompareOptions ICompareOptions `json:"CompareOptions,omitempty"`

    // Container class for compare documents.
    ComparingWithDocument *string `json:"ComparingWithDocument,omitempty"`

    // Container class for compare documents.
    DateTime *Time `json:"DateTime,omitempty"`

    // Container class for compare documents.
    ResultDocumentFormat *string `json:"ResultDocumentFormat,omitempty"`
}

func (CompareData) IsCompareData() bool {
    return true
}


func (obj *CompareData) Initialize() {
    if (obj.CompareOptions != nil) {
        obj.CompareOptions.Initialize()
    }


}

func (obj *CompareData) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Author"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Author = &parsedValue
        }

    } else if jsonValue, exists := json["author"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Author = &parsedValue
        }

    }

    if jsonValue, exists := json["CompareOptions"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ICompareOptions = new(CompareOptions)
            modelInstance.Deserialize(parsedValue)
            obj.CompareOptions = modelInstance
        }

    } else if jsonValue, exists := json["compareOptions"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ICompareOptions = new(CompareOptions)
            modelInstance.Deserialize(parsedValue)
            obj.CompareOptions = modelInstance
        }

    }

    if jsonValue, exists := json["ComparingWithDocument"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ComparingWithDocument = &parsedValue
        }

    } else if jsonValue, exists := json["comparingWithDocument"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ComparingWithDocument = &parsedValue
        }

    }

    if jsonValue, exists := json["DateTime"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DateTime = new(Time)
            obj.DateTime.Parse(parsedValue)
        }

    } else if jsonValue, exists := json["dateTime"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DateTime = new(Time)
            obj.DateTime.Parse(parsedValue)
        }

    }

    if jsonValue, exists := json["ResultDocumentFormat"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ResultDocumentFormat = &parsedValue
        }

    } else if jsonValue, exists := json["resultDocumentFormat"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ResultDocumentFormat = &parsedValue
        }

    }
}

func (obj *CompareData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *CompareData) GetAuthor() *string {
    return obj.Author
}

func (obj *CompareData) SetAuthor(value *string) {
    obj.Author = value
}

func (obj *CompareData) GetCompareOptions() ICompareOptions {
    return obj.CompareOptions
}

func (obj *CompareData) SetCompareOptions(value ICompareOptions) {
    obj.CompareOptions = value
}

func (obj *CompareData) GetComparingWithDocument() *string {
    return obj.ComparingWithDocument
}

func (obj *CompareData) SetComparingWithDocument(value *string) {
    obj.ComparingWithDocument = value
}

func (obj *CompareData) GetDateTime() *Time {
    return obj.DateTime
}

func (obj *CompareData) SetDateTime(value *Time) {
    obj.DateTime = value
}

func (obj *CompareData) GetResultDocumentFormat() *string {
    return obj.ResultDocumentFormat
}

func (obj *CompareData) SetResultDocumentFormat(value *string) {
    obj.ResultDocumentFormat = value
}

