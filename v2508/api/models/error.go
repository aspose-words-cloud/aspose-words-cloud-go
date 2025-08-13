/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="error.go">
 *   Copyright (c) 2025 Aspose.Words for Cloud
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

// Error.

type IError interface {
    IsError() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetCode() *string
    SetCode(value *string)
    GetDescription() *string
    SetDescription(value *string)
    GetInnerError() IErrorDetails
    SetInnerError(value IErrorDetails)
    GetMessage() *string
    SetMessage(value *string)
}

type Error struct {
    // Error.
    Code *string `json:"Code,omitempty"`

    // Error.
    Description *string `json:"Description,omitempty"`

    // Error.
    InnerError IErrorDetails `json:"InnerError,omitempty"`

    // Error.
    Message *string `json:"Message,omitempty"`
}

func (Error) IsError() bool {
    return true
}


func (obj *Error) Initialize() {
}

func (obj *Error) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Code"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Code = &parsedValue
        }

    } else if jsonValue, exists := json["code"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Code = &parsedValue
        }

    }

    if jsonValue, exists := json["Description"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Description = &parsedValue
        }

    } else if jsonValue, exists := json["description"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Description = &parsedValue
        }

    }

    if jsonValue, exists := json["InnerError"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IErrorDetails = new(ErrorDetails)
            modelInstance.Deserialize(parsedValue)
            obj.InnerError = modelInstance
        }

    } else if jsonValue, exists := json["innerError"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IErrorDetails = new(ErrorDetails)
            modelInstance.Deserialize(parsedValue)
            obj.InnerError = modelInstance
        }

    }

    if jsonValue, exists := json["Message"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Message = &parsedValue
        }

    } else if jsonValue, exists := json["message"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Message = &parsedValue
        }

    }
}

func (obj *Error) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *Error) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.InnerError != nil {
        if err := obj.InnerError.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *Error) GetCode() *string {
    return obj.Code
}

func (obj *Error) SetCode(value *string) {
    obj.Code = value
}

func (obj *Error) GetDescription() *string {
    return obj.Description
}

func (obj *Error) SetDescription(value *string) {
    obj.Description = value
}

func (obj *Error) GetInnerError() IErrorDetails {
    return obj.InnerError
}

func (obj *Error) SetInnerError(value IErrorDetails) {
    obj.InnerError = value
}

func (obj *Error) GetMessage() *string {
    return obj.Message
}

func (obj *Error) SetMessage(value *string) {
    obj.Message = value
}

