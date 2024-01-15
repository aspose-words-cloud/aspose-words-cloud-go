/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="protection_data_response.go">
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

// The REST response with data on document's protection.

type IProtectionDataResponse interface {
    IsProtectionDataResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetDocumentLink() IFileLink
    SetDocumentLink(value IFileLink)
    GetProtectionData() IProtectionData
    SetProtectionData(value IProtectionData)
}

type ProtectionDataResponse struct {
    // The REST response with data on document's protection.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with data on document's protection.
    DocumentLink IFileLink `json:"DocumentLink,omitempty"`

    // The REST response with data on document's protection.
    ProtectionData IProtectionData `json:"ProtectionData,omitempty"`
}

func (ProtectionDataResponse) IsProtectionDataResponse() bool {
    return true
}

func (ProtectionDataResponse) IsWordsResponse() bool {
    return true
}

func (obj *ProtectionDataResponse) Initialize() {
    if (obj.DocumentLink != nil) {
        obj.DocumentLink.Initialize()
    }

    if (obj.ProtectionData != nil) {
        obj.ProtectionData.Initialize()
    }


}

func (obj *ProtectionDataResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["DocumentLink"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFileLink = new(FileLink)
            modelInstance.Deserialize(parsedValue)
            obj.DocumentLink = modelInstance
        }

    } else if jsonValue, exists := json["documentLink"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFileLink = new(FileLink)
            modelInstance.Deserialize(parsedValue)
            obj.DocumentLink = modelInstance
        }

    }

    if jsonValue, exists := json["ProtectionData"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IProtectionData = new(ProtectionData)
            modelInstance.Deserialize(parsedValue)
            obj.ProtectionData = modelInstance
        }

    } else if jsonValue, exists := json["protectionData"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IProtectionData = new(ProtectionData)
            modelInstance.Deserialize(parsedValue)
            obj.ProtectionData = modelInstance
        }

    }
}

func (obj *ProtectionDataResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *ProtectionDataResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.DocumentLink != nil {
        if err := obj.DocumentLink.Validate(); err != nil {
            return err
        }
    }
    if obj.ProtectionData != nil {
        if err := obj.ProtectionData.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *ProtectionDataResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *ProtectionDataResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *ProtectionDataResponse) GetDocumentLink() IFileLink {
    return obj.DocumentLink
}

func (obj *ProtectionDataResponse) SetDocumentLink(value IFileLink) {
    obj.DocumentLink = value
}

func (obj *ProtectionDataResponse) GetProtectionData() IProtectionData {
    return obj.ProtectionData
}

func (obj *ProtectionDataResponse) SetProtectionData(value IProtectionData) {
    obj.ProtectionData = value
}

