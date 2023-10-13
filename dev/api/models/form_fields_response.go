/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="form_fields_response.go">
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

// The REST response with a collection of form fields.
// This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/{nodePath}/form fields" REST API requests.

type IFormFieldsResponse interface {
    IsFormFieldsResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetFormFields() IFormFieldCollection
    SetFormFields(value IFormFieldCollection)
}

type FormFieldsResponse struct {
    // The REST response with a collection of form fields.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/{nodePath}/form fields" REST API requests.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a collection of form fields.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/{nodePath}/form fields" REST API requests.
    FormFields IFormFieldCollection `json:"FormFields,omitempty"`
}

func (FormFieldsResponse) IsFormFieldsResponse() bool {
    return true
}

func (FormFieldsResponse) IsWordsResponse() bool {
    return true
}

func (obj *FormFieldsResponse) Initialize() {
    if (obj.FormFields != nil) {
        obj.FormFields.Initialize()
    }


}

func (obj *FormFieldsResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["FormFields"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFormFieldCollection = new(FormFieldCollection)
            modelInstance.Deserialize(parsedValue)
            obj.FormFields = modelInstance
        }

    } else if jsonValue, exists := json["formFields"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFormFieldCollection = new(FormFieldCollection)
            modelInstance.Deserialize(parsedValue)
            obj.FormFields = modelInstance
        }

    }
}

func (obj *FormFieldsResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *FormFieldsResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    return nil;
}

func (obj *FormFieldsResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *FormFieldsResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *FormFieldsResponse) GetFormFields() IFormFieldCollection {
    return obj.FormFields
}

func (obj *FormFieldsResponse) SetFormFields(value IFormFieldCollection) {
    obj.FormFields = value
}

