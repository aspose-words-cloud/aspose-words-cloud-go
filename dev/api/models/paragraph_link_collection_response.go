/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="paragraph_link_collection_response.go">
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

// The REST response with a collection of paragraphs.
// This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/paragraphs" REST API requests.

type IParagraphLinkCollectionResponse interface {
    IsParagraphLinkCollectionResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetParagraphs() IParagraphLinkCollection
    SetParagraphs(value IParagraphLinkCollection)
}

type ParagraphLinkCollectionResponse struct {
    // The REST response with a collection of paragraphs.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/paragraphs" REST API requests.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a collection of paragraphs.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/paragraphs" REST API requests.
    Paragraphs IParagraphLinkCollection `json:"Paragraphs,omitempty"`
}

func (ParagraphLinkCollectionResponse) IsParagraphLinkCollectionResponse() bool {
    return true
}

func (ParagraphLinkCollectionResponse) IsWordsResponse() bool {
    return true
}

func (obj *ParagraphLinkCollectionResponse) Initialize() {
    if (obj.Paragraphs != nil) {
        obj.Paragraphs.Initialize()
    }


}

func (obj *ParagraphLinkCollectionResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["Paragraphs"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IParagraphLinkCollection = new(ParagraphLinkCollection)
            modelInstance.Deserialize(parsedValue)
            obj.Paragraphs = modelInstance
        }

    } else if jsonValue, exists := json["paragraphs"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IParagraphLinkCollection = new(ParagraphLinkCollection)
            modelInstance.Deserialize(parsedValue)
            obj.Paragraphs = modelInstance
        }

    }
}

func (obj *ParagraphLinkCollectionResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *ParagraphLinkCollectionResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    return nil;
}

func (obj *ParagraphLinkCollectionResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *ParagraphLinkCollectionResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *ParagraphLinkCollectionResponse) GetParagraphs() IParagraphLinkCollection {
    return obj.Paragraphs
}

func (obj *ParagraphLinkCollectionResponse) SetParagraphs(value IParagraphLinkCollection) {
    obj.Paragraphs = value
}

