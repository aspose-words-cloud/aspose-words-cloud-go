/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="classification_response.go">
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

// The REST response with data on multi-class text classification.

type IClassificationResponse interface {
    IsClassificationResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetRequestId() *string
    SetRequestId(value *string)
    GetBestClassName() *string
    SetBestClassName(value *string)
    GetBestClassProbability() *float64
    SetBestClassProbability(value *float64)
    GetBestResults() []IClassificationResult
    SetBestResults(value []IClassificationResult)
}

type ClassificationResponse struct {
    // The REST response with data on multi-class text classification.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with data on multi-class text classification.
    BestClassName *string `json:"BestClassName,omitempty"`

    // The REST response with data on multi-class text classification.
    BestClassProbability *float64 `json:"BestClassProbability,omitempty"`

    // The REST response with data on multi-class text classification.
    BestResults []IClassificationResult `json:"BestResults,omitempty"`
}

func (ClassificationResponse) IsClassificationResponse() bool {
    return true
}

func (ClassificationResponse) IsWordsResponse() bool {
    return true
}

func (obj *ClassificationResponse) Initialize() {
    if (obj.BestResults != nil) {
        for _, objElementBestResults := range obj.BestResults {
            objElementBestResults.Initialize()
        }
    }

}

func (obj *ClassificationResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["BestClassName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.BestClassName = &parsedValue
        }

    } else if jsonValue, exists := json["bestClassName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.BestClassName = &parsedValue
        }

    }

    if jsonValue, exists := json["BestClassProbability"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.BestClassProbability = &parsedValue
        }

    } else if jsonValue, exists := json["bestClassProbability"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.BestClassProbability = &parsedValue
        }

    }

    if jsonValue, exists := json["BestResults"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.BestResults = make([]IClassificationResult, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IClassificationResult = new(ClassificationResult)
                    modelElementInstance.Deserialize(elementValue)
                    obj.BestResults = append(obj.BestResults, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["bestResults"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.BestResults = make([]IClassificationResult, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IClassificationResult = new(ClassificationResult)
                    modelElementInstance.Deserialize(elementValue)
                    obj.BestResults = append(obj.BestResults, modelElementInstance)
                }

            }
        }

    }
}

func (obj *ClassificationResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *ClassificationResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *ClassificationResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *ClassificationResponse) GetBestClassName() *string {
    return obj.BestClassName
}

func (obj *ClassificationResponse) SetBestClassName(value *string) {
    obj.BestClassName = value
}

func (obj *ClassificationResponse) GetBestClassProbability() *float64 {
    return obj.BestClassProbability
}

func (obj *ClassificationResponse) SetBestClassProbability(value *float64) {
    obj.BestClassProbability = value
}

func (obj *ClassificationResponse) GetBestResults() []IClassificationResult {
    return obj.BestResults
}

func (obj *ClassificationResponse) SetBestResults(value []IClassificationResult) {
    obj.BestResults = value
}

