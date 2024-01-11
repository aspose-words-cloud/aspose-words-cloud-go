/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="replace_text_response.go">
 *   Copyright (c) 2024 Aspose.Words for Cloud
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

// The REST response with a number of occurrences of the captured text in the document.

type IReplaceTextResponse interface {
    IsReplaceTextResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetDocumentLink() IFileLink
    SetDocumentLink(value IFileLink)
    GetMatches() *int32
    SetMatches(value *int32)
}

type ReplaceTextResponse struct {
    // The REST response with a number of occurrences of the captured text in the document.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a number of occurrences of the captured text in the document.
    DocumentLink IFileLink `json:"DocumentLink,omitempty"`

    // The REST response with a number of occurrences of the captured text in the document.
    Matches *int32 `json:"Matches,omitempty"`
}

func (ReplaceTextResponse) IsReplaceTextResponse() bool {
    return true
}

func (ReplaceTextResponse) IsWordsResponse() bool {
    return true
}

func (obj *ReplaceTextResponse) Initialize() {
    if (obj.DocumentLink != nil) {
        obj.DocumentLink.Initialize()
    }


}

func (obj *ReplaceTextResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["DocumentLink"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFileLink = new(FileLink)
            modelInstance.Deserialize(parsedValue)
            obj.DocumentLink = modelInstance
        }

    } else if jsonValue, exists := json["documentLink"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFileLink = new(FileLink)
            modelInstance.Deserialize(parsedValue)
            obj.DocumentLink = modelInstance
        }

    }

    if jsonValue, exists := json["Matches"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Matches = new(int32)
            *obj.Matches = int32(parsedValue)
        }

    } else if jsonValue, exists := json["matches"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Matches = new(int32)
            *obj.Matches = int32(parsedValue)
        }

    }
}

func (obj *ReplaceTextResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *ReplaceTextResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Matches == nil {
        return errors.New("Property Matches in ReplaceTextResponse is required.")
    }
    if obj.DocumentLink != nil {
        if err := obj.DocumentLink.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *ReplaceTextResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *ReplaceTextResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *ReplaceTextResponse) GetDocumentLink() IFileLink {
    return obj.DocumentLink
}

func (obj *ReplaceTextResponse) SetDocumentLink(value IFileLink) {
    obj.DocumentLink = value
}

func (obj *ReplaceTextResponse) GetMatches() *int32 {
    return obj.Matches
}

func (obj *ReplaceTextResponse) SetMatches(value *int32) {
    obj.Matches = value
}

