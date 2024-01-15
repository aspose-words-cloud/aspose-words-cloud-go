/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="borders_response.go">
 *   Copyright (c) 2024 Aspose.Words for Cloud
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

// The REST response with a collection of borders.
// This response is returned by the Service when handling "GET {nodeWithBorders}/borders" REST API requests.

type IBordersResponse interface {
    IsBordersResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetBorders() IBordersCollection
    SetBorders(value IBordersCollection)
}

type BordersResponse struct {
    // The REST response with a collection of borders.
    // This response is returned by the Service when handling "GET {nodeWithBorders}/borders" REST API requests.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a collection of borders.
    // This response is returned by the Service when handling "GET {nodeWithBorders}/borders" REST API requests.
    Borders IBordersCollection `json:"Borders,omitempty"`
}

func (BordersResponse) IsBordersResponse() bool {
    return true
}

func (BordersResponse) IsWordsResponse() bool {
    return true
}

func (obj *BordersResponse) Initialize() {
    if (obj.Borders != nil) {
        obj.Borders.Initialize()
    }


}

func (obj *BordersResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["Borders"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IBordersCollection = new(BordersCollection)
            modelInstance.Deserialize(parsedValue)
            obj.Borders = modelInstance
        }

    } else if jsonValue, exists := json["borders"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IBordersCollection = new(BordersCollection)
            modelInstance.Deserialize(parsedValue)
            obj.Borders = modelInstance
        }

    }
}

func (obj *BordersResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *BordersResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Borders != nil {
        if err := obj.Borders.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *BordersResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *BordersResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *BordersResponse) GetBorders() IBordersCollection {
    return obj.Borders
}

func (obj *BordersResponse) SetBorders(value IBordersCollection) {
    obj.Borders = value
}

