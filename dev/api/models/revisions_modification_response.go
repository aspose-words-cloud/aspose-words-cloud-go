/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="revisions_modification_response.go">
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

// The REST response with a result of the modification operations for the revisions collection (now these are acceptAll and rejectAll).

type IRevisionsModificationResponse interface {
    IsRevisionsModificationResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetResult() IModificationOperationResult
    SetResult(value IModificationOperationResult)
}

type RevisionsModificationResponse struct {
    // The REST response with a result of the modification operations for the revisions collection (now these are acceptAll and rejectAll).
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a result of the modification operations for the revisions collection (now these are acceptAll and rejectAll).
    Result IModificationOperationResult `json:"Result,omitempty"`
}

func (RevisionsModificationResponse) IsRevisionsModificationResponse() bool {
    return true
}

func (RevisionsModificationResponse) IsWordsResponse() bool {
    return true
}

func (obj *RevisionsModificationResponse) Initialize() {
    if (obj.Result != nil) {
        obj.Result.Initialize()
    }


}

func (obj *RevisionsModificationResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["Result"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IModificationOperationResult = new(ModificationOperationResult)
            modelInstance.Deserialize(parsedValue)
            obj.Result = modelInstance
        }

    } else if jsonValue, exists := json["result"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IModificationOperationResult = new(ModificationOperationResult)
            modelInstance.Deserialize(parsedValue)
            obj.Result = modelInstance
        }

    }
}

func (obj *RevisionsModificationResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *RevisionsModificationResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    return nil;
}

func (obj *RevisionsModificationResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *RevisionsModificationResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *RevisionsModificationResponse) GetResult() IModificationOperationResult {
    return obj.Result
}

func (obj *RevisionsModificationResponse) SetResult(value IModificationOperationResult) {
    obj.Result = value
}

