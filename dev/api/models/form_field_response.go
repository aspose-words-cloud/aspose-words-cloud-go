/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="form_field_response.go">
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

// The REST response with a form field.

type IFormFieldResponse interface {
    IsFormFieldResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetRequestId() *string
    SetRequestId(value *string)
    GetFormField() IFormField
    SetFormField(value IFormField)
}

type FormFieldResponse struct {
    // The REST response with a form field.
    RequestId *string

    // The REST response with a form field.
    FormField IFormField
}

func (FormFieldResponse) IsFormFieldResponse() bool {
    return true
}

func (FormFieldResponse) IsWordsResponse() bool {
    return true
}

func (obj *FormFieldResponse) Initialize() {
    if (obj.FormField != nil) {
        obj.FormField.Initialize()
    }


}

func (obj *FormFieldResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["FormField"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFormField = nil
            if jsonType, found := parsedValue["$type"]; found {
                jsonTypeStr := jsonType.(string)
                if jsonTypeStr == "FormFieldCheckbox, _" { modelInstance = new(FormFieldCheckbox) }
                if jsonTypeStr == "FormFieldDropDown, _" { modelInstance = new(FormFieldDropDown) }
                if jsonTypeStr == "FormFieldTextInput, _" { modelInstance = new(FormFieldTextInput) }
            }

            modelInstance.Deserialize(parsedValue)
            obj.FormField = modelInstance
        }

    } else if jsonValue, exists := json["formField"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFormField = nil
            if jsonType, found := parsedValue["$type"]; found {
                jsonTypeStr := jsonType.(string)
                if jsonTypeStr == "FormFieldCheckbox, _" { modelInstance = new(FormFieldCheckbox) }
                if jsonTypeStr == "FormFieldDropDown, _" { modelInstance = new(FormFieldDropDown) }
                if jsonTypeStr == "FormFieldTextInput, _" { modelInstance = new(FormFieldTextInput) }
            }

            modelInstance.Deserialize(parsedValue)
            obj.FormField = modelInstance
        }

    }
}

func (obj *FormFieldResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *FormFieldResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *FormFieldResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *FormFieldResponse) GetFormField() IFormField {
    return obj.FormField
}

func (obj *FormFieldResponse) SetFormField(value IFormField) {
    obj.FormField = value
}

