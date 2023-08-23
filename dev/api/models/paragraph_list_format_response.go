/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="paragraph_list_format_response.go">
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

// The REST response with a list format for a paragraph.

type IParagraphListFormatResponse interface {
    IsParagraphListFormatResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetRequestId() *string
    SetRequestId(value *string)
    GetListFormat() IListFormat
    SetListFormat(value IListFormat)
}

type ParagraphListFormatResponse struct {
    // The REST response with a list format for a paragraph.
    RequestId *string

    // The REST response with a list format for a paragraph.
    ListFormat IListFormat
}

func (ParagraphListFormatResponse) IsParagraphListFormatResponse() bool {
    return true
}

func (ParagraphListFormatResponse) IsWordsResponse() bool {
    return true
}

func (obj *ParagraphListFormatResponse) Initialize() {
    if (obj.ListFormat != nil) {
        obj.ListFormat.Initialize()
    }


}

func (obj *ParagraphListFormatResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["ListFormat"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IListFormat = new(ListFormat)
            modelInstance.Deserialize(parsedValue)
            obj.ListFormat = modelInstance
        }

    } else if jsonValue, exists := json["listFormat"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IListFormat = new(ListFormat)
            modelInstance.Deserialize(parsedValue)
            obj.ListFormat = modelInstance
        }

    }
}

func (obj *ParagraphListFormatResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *ParagraphListFormatResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *ParagraphListFormatResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *ParagraphListFormatResponse) GetListFormat() IListFormat {
    return obj.ListFormat
}

func (obj *ParagraphListFormatResponse) SetListFormat(value IListFormat) {
    obj.ListFormat = value
}

