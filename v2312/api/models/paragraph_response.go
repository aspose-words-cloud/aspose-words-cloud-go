/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="paragraph_response.go">
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

// The REST response with a paragraph.
// This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/paragraphs/{0}" REST API requests.

type IParagraphResponse interface {
    IsParagraphResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetParagraph() IParagraph
    SetParagraph(value IParagraph)
}

type ParagraphResponse struct {
    // The REST response with a paragraph.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/paragraphs/{0}" REST API requests.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a paragraph.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/paragraphs/{0}" REST API requests.
    Paragraph IParagraph `json:"Paragraph,omitempty"`
}

func (ParagraphResponse) IsParagraphResponse() bool {
    return true
}

func (ParagraphResponse) IsWordsResponse() bool {
    return true
}

func (obj *ParagraphResponse) Initialize() {
    if (obj.Paragraph != nil) {
        obj.Paragraph.Initialize()
    }


}

func (obj *ParagraphResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["Paragraph"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IParagraph = new(Paragraph)
            modelInstance.Deserialize(parsedValue)
            obj.Paragraph = modelInstance
        }

    } else if jsonValue, exists := json["paragraph"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IParagraph = new(Paragraph)
            modelInstance.Deserialize(parsedValue)
            obj.Paragraph = modelInstance
        }

    }
}

func (obj *ParagraphResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *ParagraphResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Paragraph != nil {
        if err := obj.Paragraph.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *ParagraphResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *ParagraphResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *ParagraphResponse) GetParagraph() IParagraph {
    return obj.Paragraph
}

func (obj *ParagraphResponse) SetParagraph(value IParagraph) {
    obj.Paragraph = value
}

