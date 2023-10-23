/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="save_response.go">
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

// The REST response with a save result.

type ISaveResponse interface {
    IsSaveResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetSaveResult() ISaveResult
    SetSaveResult(value ISaveResult)
}

type SaveResponse struct {
    // The REST response with a save result.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a save result.
    SaveResult ISaveResult `json:"SaveResult,omitempty"`
}

func (SaveResponse) IsSaveResponse() bool {
    return true
}

func (SaveResponse) IsWordsResponse() bool {
    return true
}

func (obj *SaveResponse) Initialize() {
    if (obj.SaveResult != nil) {
        obj.SaveResult.Initialize()
    }


}

func (obj *SaveResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["SaveResult"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ISaveResult = new(SaveResult)
            modelInstance.Deserialize(parsedValue)
            obj.SaveResult = modelInstance
        }

    } else if jsonValue, exists := json["saveResult"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ISaveResult = new(SaveResult)
            modelInstance.Deserialize(parsedValue)
            obj.SaveResult = modelInstance
        }

    }
}

func (obj *SaveResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *SaveResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.RequestId == nil {
        return errors.New("Property RequestId in SaveResponse is required.")
    }

    if obj.SaveResult == nil {
        return errors.New("Property SaveResult in SaveResponse is required.")
    }

    if obj.SaveResult != nil {
        if err := obj.SaveResult.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *SaveResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *SaveResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *SaveResponse) GetSaveResult() ISaveResult {
    return obj.SaveResult
}

func (obj *SaveResponse) SetSaveResult(value ISaveResult) {
    obj.SaveResult = value
}

