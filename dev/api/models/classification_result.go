/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="classification_result.go">
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

// Represents a single classification result.

type IClassificationResult interface {
    IsClassificationResult() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetClassName() *string
    SetClassName(value *string)
    GetClassProbability() *float64
    SetClassProbability(value *float64)
}

type ClassificationResult struct {
    // Represents a single classification result.
    ClassName *string `json:"ClassName,omitempty"`

    // Represents a single classification result.
    ClassProbability *float64 `json:"ClassProbability,omitempty"`
}

func (ClassificationResult) IsClassificationResult() bool {
    return true
}


func (obj *ClassificationResult) Initialize() {
}

func (obj *ClassificationResult) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["ClassName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ClassName = &parsedValue
        }

    } else if jsonValue, exists := json["className"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ClassName = &parsedValue
        }

    }

    if jsonValue, exists := json["ClassProbability"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ClassProbability = &parsedValue
        }

    } else if jsonValue, exists := json["classProbability"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ClassProbability = &parsedValue
        }

    }
}

func (obj *ClassificationResult) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *ClassificationResult) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.ClassProbability == nil {
        return errors.New("Property ClassProbability in ClassificationResult is required.")
    }
    if obj.ClassName == nil {
        return errors.New("Property ClassName in ClassificationResult is required.")
    }

    if obj.ClassProbability == nil {
        return errors.New("Property ClassProbability in ClassificationResult is required.")
    }

    return nil;
}

func (obj *ClassificationResult) GetClassName() *string {
    return obj.ClassName
}

func (obj *ClassificationResult) SetClassName(value *string) {
    obj.ClassName = value
}

func (obj *ClassificationResult) GetClassProbability() *float64 {
    return obj.ClassProbability
}

func (obj *ClassificationResult) SetClassProbability(value *float64) {
    obj.ClassProbability = value
}

