/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="error_details.go">
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

// The error details.

type IErrorDetails interface {
    IsErrorDetails() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetErrorDateTime() *Time
    SetErrorDateTime(value *Time)
    GetRequestId() *string
    SetRequestId(value *string)
}

type ErrorDetails struct {
    // The error details.
    ErrorDateTime *Time `json:"ErrorDateTime,omitempty"`

    // The error details.
    RequestId *string `json:"RequestId,omitempty"`
}

func (ErrorDetails) IsErrorDetails() bool {
    return true
}


func (obj *ErrorDetails) Initialize() {
}

func (obj *ErrorDetails) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["ErrorDateTime"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ErrorDateTime = new(Time)
            obj.ErrorDateTime.Parse(parsedValue)
        }

    } else if jsonValue, exists := json["errorDateTime"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ErrorDateTime = new(Time)
            obj.ErrorDateTime.Parse(parsedValue)
        }

    }

    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }
}

func (obj *ErrorDetails) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *ErrorDetails) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.ErrorDateTime == nil {
        return errors.New("Property ErrorDateTime in ErrorDetails is required.")
    }

    return nil;
}

func (obj *ErrorDetails) GetErrorDateTime() *Time {
    return obj.ErrorDateTime
}

func (obj *ErrorDetails) SetErrorDateTime(value *Time) {
    obj.ErrorDateTime = value
}

func (obj *ErrorDetails) GetRequestId() *string {
    return obj.RequestId
}

func (obj *ErrorDetails) SetRequestId(value *string) {
    obj.RequestId = value
}

