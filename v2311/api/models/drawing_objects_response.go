/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="drawing_objects_response.go">
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

// The REST response with a collection of DrawingObjects.
// This response should be returned by the service when handling: GET /drawingObjects.

type IDrawingObjectsResponse interface {
    IsDrawingObjectsResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetDrawingObjects() IDrawingObjectCollection
    SetDrawingObjects(value IDrawingObjectCollection)
}

type DrawingObjectsResponse struct {
    // The REST response with a collection of DrawingObjects.
    // This response should be returned by the service when handling: GET /drawingObjects.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a collection of DrawingObjects.
    // This response should be returned by the service when handling: GET /drawingObjects.
    DrawingObjects IDrawingObjectCollection `json:"DrawingObjects,omitempty"`
}

func (DrawingObjectsResponse) IsDrawingObjectsResponse() bool {
    return true
}

func (DrawingObjectsResponse) IsWordsResponse() bool {
    return true
}

func (obj *DrawingObjectsResponse) Initialize() {
    if (obj.DrawingObjects != nil) {
        obj.DrawingObjects.Initialize()
    }


}

func (obj *DrawingObjectsResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["DrawingObjects"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IDrawingObjectCollection = new(DrawingObjectCollection)
            modelInstance.Deserialize(parsedValue)
            obj.DrawingObjects = modelInstance
        }

    } else if jsonValue, exists := json["drawingObjects"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IDrawingObjectCollection = new(DrawingObjectCollection)
            modelInstance.Deserialize(parsedValue)
            obj.DrawingObjects = modelInstance
        }

    }
}

func (obj *DrawingObjectsResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *DrawingObjectsResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.DrawingObjects != nil {
        if err := obj.DrawingObjects.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *DrawingObjectsResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *DrawingObjectsResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *DrawingObjectsResponse) GetDrawingObjects() IDrawingObjectCollection {
    return obj.DrawingObjects
}

func (obj *DrawingObjectsResponse) SetDrawingObjects(value IDrawingObjectCollection) {
    obj.DrawingObjects = value
}

