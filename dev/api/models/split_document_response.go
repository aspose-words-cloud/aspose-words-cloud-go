/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="split_document_response.go">
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

// The REST response with a result of document splitting.
// This response should be returned by the service when handling: POST /{name}/split.

type ISplitDocumentResponse interface {
    IsSplitDocumentResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetSplitResult() ISplitDocumentResult
    SetSplitResult(value ISplitDocumentResult)
}

type SplitDocumentResponse struct {
    // The REST response with a result of document splitting.
    // This response should be returned by the service when handling: POST /{name}/split.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a result of document splitting.
    // This response should be returned by the service when handling: POST /{name}/split.
    SplitResult ISplitDocumentResult `json:"SplitResult,omitempty"`
}

func (SplitDocumentResponse) IsSplitDocumentResponse() bool {
    return true
}

func (SplitDocumentResponse) IsWordsResponse() bool {
    return true
}

func (obj *SplitDocumentResponse) Initialize() {
    if (obj.SplitResult != nil) {
        obj.SplitResult.Initialize()
    }


}

func (obj *SplitDocumentResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["SplitResult"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ISplitDocumentResult = new(SplitDocumentResult)
            modelInstance.Deserialize(parsedValue)
            obj.SplitResult = modelInstance
        }

    } else if jsonValue, exists := json["splitResult"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ISplitDocumentResult = new(SplitDocumentResult)
            modelInstance.Deserialize(parsedValue)
            obj.SplitResult = modelInstance
        }

    }
}

func (obj *SplitDocumentResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *SplitDocumentResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.SplitResult != nil {
        if err := obj.SplitResult.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *SplitDocumentResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *SplitDocumentResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *SplitDocumentResponse) GetSplitResult() ISplitDocumentResult {
    return obj.SplitResult
}

func (obj *SplitDocumentResponse) SetSplitResult(value ISplitDocumentResult) {
    obj.SplitResult = value
}

