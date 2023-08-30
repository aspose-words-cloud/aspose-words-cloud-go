/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="paragraph_format_response.go">
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

// The REST response with the formatting properties of a paragraph.

type IParagraphFormatResponse interface {
    IsParagraphFormatResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetRequestId() *string
    SetRequestId(value *string)
    GetParagraphFormat() IParagraphFormat
    SetParagraphFormat(value IParagraphFormat)
}

type ParagraphFormatResponse struct {
    // The REST response with the formatting properties of a paragraph.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with the formatting properties of a paragraph.
    ParagraphFormat IParagraphFormat `json:"ParagraphFormat,omitempty"`
}

func (ParagraphFormatResponse) IsParagraphFormatResponse() bool {
    return true
}

func (ParagraphFormatResponse) IsWordsResponse() bool {
    return true
}

func (obj *ParagraphFormatResponse) Initialize() {
    if (obj.ParagraphFormat != nil) {
        obj.ParagraphFormat.Initialize()
    }


}

func (obj *ParagraphFormatResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["ParagraphFormat"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IParagraphFormat = new(ParagraphFormat)
            modelInstance.Deserialize(parsedValue)
            obj.ParagraphFormat = modelInstance
        }

    } else if jsonValue, exists := json["paragraphFormat"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IParagraphFormat = new(ParagraphFormat)
            modelInstance.Deserialize(parsedValue)
            obj.ParagraphFormat = modelInstance
        }

    }
}

func (obj *ParagraphFormatResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *ParagraphFormatResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *ParagraphFormatResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *ParagraphFormatResponse) GetParagraphFormat() IParagraphFormat {
    return obj.ParagraphFormat
}

func (obj *ParagraphFormatResponse) SetParagraphFormat(value IParagraphFormat) {
    obj.ParagraphFormat = value
}

