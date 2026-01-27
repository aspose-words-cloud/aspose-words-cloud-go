/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="words_api_error_response.go">
 *   Copyright (c) 2026 Aspose.Words for Cloud
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

// The REST response with an API error.

type IWordsApiErrorResponse interface {
    IsWordsApiErrorResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetError() IApiError
    SetError(value IApiError)
}

type WordsApiErrorResponse struct {
    // The REST response with an API error.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with an API error.
    Error_ IApiError `json:"Error,omitempty"`
}

func (WordsApiErrorResponse) IsWordsApiErrorResponse() bool {
    return true
}

func (WordsApiErrorResponse) IsWordsResponse() bool {
    return true
}

func (obj *WordsApiErrorResponse) Initialize() {
}

func (obj *WordsApiErrorResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["Error"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IApiError = new(ApiError)
            modelInstance.Deserialize(parsedValue)
            obj.Error_ = modelInstance
        }

    } else if jsonValue, exists := json["error"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IApiError = new(ApiError)
            modelInstance.Deserialize(parsedValue)
            obj.Error_ = modelInstance
        }

    }
}

func (obj *WordsApiErrorResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *WordsApiErrorResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Error_ != nil {
        if err := obj.Error_.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *WordsApiErrorResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *WordsApiErrorResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *WordsApiErrorResponse) GetError() IApiError {
    return obj.Error_
}

func (obj *WordsApiErrorResponse) SetError(value IApiError) {
    obj.Error_ = value
}

