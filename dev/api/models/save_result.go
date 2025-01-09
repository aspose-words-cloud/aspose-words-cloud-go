/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="save_result.go">
 *   Copyright (c) 2025 Aspose.Words for Cloud
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

// Result of saving.

type ISaveResult interface {
    IsSaveResult() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetDestDocument() IFileLink
    SetDestDocument(value IFileLink)
    GetSourceDocument() IFileLink
    SetSourceDocument(value IFileLink)
    GetAdditionalItems() []IFileLink
    SetAdditionalItems(value []IFileLink)
}

type SaveResult struct {
    // Result of saving.
    DestDocument IFileLink `json:"DestDocument,omitempty"`

    // Result of saving.
    SourceDocument IFileLink `json:"SourceDocument,omitempty"`

    // Result of saving.
    AdditionalItems []IFileLink `json:"AdditionalItems,omitempty"`
}

func (SaveResult) IsSaveResult() bool {
    return true
}


func (obj *SaveResult) Initialize() {
    if (obj.DestDocument != nil) {
        obj.DestDocument.Initialize()
    }

    if (obj.SourceDocument != nil) {
        obj.SourceDocument.Initialize()
    }

    if (obj.AdditionalItems != nil) {
        for _, objElementAdditionalItems := range obj.AdditionalItems {
            objElementAdditionalItems.Initialize()
        }
    }

}

func (obj *SaveResult) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["DestDocument"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFileLink = new(FileLink)
            modelInstance.Deserialize(parsedValue)
            obj.DestDocument = modelInstance
        }

    } else if jsonValue, exists := json["destDocument"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFileLink = new(FileLink)
            modelInstance.Deserialize(parsedValue)
            obj.DestDocument = modelInstance
        }

    }

    if jsonValue, exists := json["SourceDocument"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFileLink = new(FileLink)
            modelInstance.Deserialize(parsedValue)
            obj.SourceDocument = modelInstance
        }

    } else if jsonValue, exists := json["sourceDocument"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFileLink = new(FileLink)
            modelInstance.Deserialize(parsedValue)
            obj.SourceDocument = modelInstance
        }

    }

    if jsonValue, exists := json["AdditionalItems"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.AdditionalItems = make([]IFileLink, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IFileLink = new(FileLink)
                    modelElementInstance.Deserialize(elementValue)
                    obj.AdditionalItems = append(obj.AdditionalItems, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["additionalItems"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.AdditionalItems = make([]IFileLink, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IFileLink = new(FileLink)
                    modelElementInstance.Deserialize(elementValue)
                    obj.AdditionalItems = append(obj.AdditionalItems, modelElementInstance)
                }

            }
        }

    }
}

func (obj *SaveResult) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *SaveResult) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.DestDocument != nil {
        if err := obj.DestDocument.Validate(); err != nil {
            return err
        }
    }
    if obj.SourceDocument != nil {
        if err := obj.SourceDocument.Validate(); err != nil {
            return err
        }
    }
    if obj.AdditionalItems != nil {
        for _, elementAdditionalItems := range obj.AdditionalItems {
            if elementAdditionalItems != nil {
                if err := elementAdditionalItems.Validate(); err != nil {
                    return err
                }
            }
        }
    }

    return nil;
}

func (obj *SaveResult) GetDestDocument() IFileLink {
    return obj.DestDocument
}

func (obj *SaveResult) SetDestDocument(value IFileLink) {
    obj.DestDocument = value
}

func (obj *SaveResult) GetSourceDocument() IFileLink {
    return obj.SourceDocument
}

func (obj *SaveResult) SetSourceDocument(value IFileLink) {
    obj.SourceDocument = value
}

func (obj *SaveResult) GetAdditionalItems() []IFileLink {
    return obj.AdditionalItems
}

func (obj *SaveResult) SetAdditionalItems(value []IFileLink) {
    obj.AdditionalItems = value
}

