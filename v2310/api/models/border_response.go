/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="border_response.go">
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

// The REST response with a border.
// This response is returned by the Service when handling "GET {nodeWithBorders}/borders" REST API requests.

type IBorderResponse interface {
    IsBorderResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetRequestId() *string
    SetRequestId(value *string)
    GetBorder() IBorder
    SetBorder(value IBorder)
}

type BorderResponse struct {
    // The REST response with a border.
    // This response is returned by the Service when handling "GET {nodeWithBorders}/borders" REST API requests.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a border.
    // This response is returned by the Service when handling "GET {nodeWithBorders}/borders" REST API requests.
    Border IBorder `json:"Border,omitempty"`
}

func (BorderResponse) IsBorderResponse() bool {
    return true
}

func (BorderResponse) IsWordsResponse() bool {
    return true
}

func (obj *BorderResponse) Initialize() {
    if (obj.Border != nil) {
        obj.Border.Initialize()
    }


}

func (obj *BorderResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["Border"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IBorder = new(Border)
            modelInstance.Deserialize(parsedValue)
            obj.Border = modelInstance
        }

    } else if jsonValue, exists := json["border"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IBorder = new(Border)
            modelInstance.Deserialize(parsedValue)
            obj.Border = modelInstance
        }

    }
}

func (obj *BorderResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *BorderResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *BorderResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *BorderResponse) GetBorder() IBorder {
    return obj.Border
}

func (obj *BorderResponse) SetBorder(value IBorder) {
    obj.Border = value
}

