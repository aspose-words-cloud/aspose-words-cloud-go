/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="json_data_load_options.go">
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

// Represents options for parsing JSON data.
// An instance of this class can be passed into constructors of Aspose.Words.Reporting.JsonDataSource.

type IJsonDataLoadOptions interface {
    IsJsonDataLoadOptions() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetAlwaysGenerateRootObject() *bool
    SetAlwaysGenerateRootObject(value *bool)
    GetExactDateTimeParseFormats() []string
    SetExactDateTimeParseFormats(value []string)
    GetSimpleValueParseMode() *string
    SetSimpleValueParseMode(value *string)
}

type JsonDataLoadOptions struct {
    // Represents options for parsing JSON data.
    // An instance of this class can be passed into constructors of Aspose.Words.Reporting.JsonDataSource.
    AlwaysGenerateRootObject *bool `json:"AlwaysGenerateRootObject,omitempty"`

    // Represents options for parsing JSON data.
    // An instance of this class can be passed into constructors of Aspose.Words.Reporting.JsonDataSource.
    ExactDateTimeParseFormats []string `json:"ExactDateTimeParseFormats,omitempty"`

    // Represents options for parsing JSON data.
    // An instance of this class can be passed into constructors of Aspose.Words.Reporting.JsonDataSource.
    SimpleValueParseMode *string `json:"SimpleValueParseMode,omitempty"`
}

func (JsonDataLoadOptions) IsJsonDataLoadOptions() bool {
    return true
}


func (obj *JsonDataLoadOptions) Initialize() {
}

func (obj *JsonDataLoadOptions) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["AlwaysGenerateRootObject"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.AlwaysGenerateRootObject = &parsedValue
        }

    } else if jsonValue, exists := json["alwaysGenerateRootObject"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.AlwaysGenerateRootObject = &parsedValue
        }

    }

    if jsonValue, exists := json["ExactDateTimeParseFormats"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.ExactDateTimeParseFormats = make([]string, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(string); valid {
                    obj.ExactDateTimeParseFormats = append(obj.ExactDateTimeParseFormats, elementValue)
                }

            }
        }

    } else if jsonValue, exists := json["exactDateTimeParseFormats"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.ExactDateTimeParseFormats = make([]string, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(string); valid {
                    obj.ExactDateTimeParseFormats = append(obj.ExactDateTimeParseFormats, elementValue)
                }

            }
        }

    }

    if jsonValue, exists := json["SimpleValueParseMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.SimpleValueParseMode = &parsedValue
        }

    } else if jsonValue, exists := json["simpleValueParseMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.SimpleValueParseMode = &parsedValue
        }

    }
}

func (obj *JsonDataLoadOptions) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *JsonDataLoadOptions) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.AlwaysGenerateRootObject == nil {
        return errors.New("Property AlwaysGenerateRootObject in JsonDataLoadOptions is required.")
    }
    if obj.SimpleValueParseMode == nil {
        return errors.New("Property SimpleValueParseMode in JsonDataLoadOptions is required.")
    }
    return nil;
}

func (obj *JsonDataLoadOptions) GetAlwaysGenerateRootObject() *bool {
    return obj.AlwaysGenerateRootObject
}

func (obj *JsonDataLoadOptions) SetAlwaysGenerateRootObject(value *bool) {
    obj.AlwaysGenerateRootObject = value
}

func (obj *JsonDataLoadOptions) GetExactDateTimeParseFormats() []string {
    return obj.ExactDateTimeParseFormats
}

func (obj *JsonDataLoadOptions) SetExactDateTimeParseFormats(value []string) {
    obj.ExactDateTimeParseFormats = value
}

func (obj *JsonDataLoadOptions) GetSimpleValueParseMode() *string {
    return obj.SimpleValueParseMode
}

func (obj *JsonDataLoadOptions) SetSimpleValueParseMode(value *string) {
    obj.SimpleValueParseMode = value
}

