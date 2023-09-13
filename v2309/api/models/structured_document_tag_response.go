/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="structured_document_tag_response.go">
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

// The REST response with a StructuredDocumentTag.

type IStructuredDocumentTagResponse interface {
    IsStructuredDocumentTagResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetRequestId() *string
    SetRequestId(value *string)
    GetStructuredDocumentTag() IStructuredDocumentTag
    SetStructuredDocumentTag(value IStructuredDocumentTag)
}

type StructuredDocumentTagResponse struct {
    // The REST response with a StructuredDocumentTag.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a StructuredDocumentTag.
    StructuredDocumentTag IStructuredDocumentTag `json:"StructuredDocumentTag,omitempty"`
}

func (StructuredDocumentTagResponse) IsStructuredDocumentTagResponse() bool {
    return true
}

func (StructuredDocumentTagResponse) IsWordsResponse() bool {
    return true
}

func (obj *StructuredDocumentTagResponse) Initialize() {
    if (obj.StructuredDocumentTag != nil) {
        obj.StructuredDocumentTag.Initialize()
    }


}

func (obj *StructuredDocumentTagResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["StructuredDocumentTag"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IStructuredDocumentTag = nil
            if jsonType, found := parsedValue["$type"]; found {
                jsonTypeStr := jsonType.(string)
                if jsonTypeStr == "StructuredDocumentTagInsert, _" { modelInstance = new(StructuredDocumentTagInsert) }
                if jsonTypeStr == "StructuredDocumentTagUpdate, _" { modelInstance = new(StructuredDocumentTagUpdate) }
            }

            if modelInstance == nil { modelInstance = new(StructuredDocumentTag) }
            modelInstance.Deserialize(parsedValue)
            obj.StructuredDocumentTag = modelInstance
        }

    } else if jsonValue, exists := json["structuredDocumentTag"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IStructuredDocumentTag = nil
            if jsonType, found := parsedValue["$type"]; found {
                jsonTypeStr := jsonType.(string)
                if jsonTypeStr == "StructuredDocumentTagInsert, _" { modelInstance = new(StructuredDocumentTagInsert) }
                if jsonTypeStr == "StructuredDocumentTagUpdate, _" { modelInstance = new(StructuredDocumentTagUpdate) }
            }

            if modelInstance == nil { modelInstance = new(StructuredDocumentTag) }
            modelInstance.Deserialize(parsedValue)
            obj.StructuredDocumentTag = modelInstance
        }

    }
}

func (obj *StructuredDocumentTagResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *StructuredDocumentTagResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *StructuredDocumentTagResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *StructuredDocumentTagResponse) GetStructuredDocumentTag() IStructuredDocumentTag {
    return obj.StructuredDocumentTag
}

func (obj *StructuredDocumentTagResponse) SetStructuredDocumentTag(value IStructuredDocumentTag) {
    obj.StructuredDocumentTag = value
}

