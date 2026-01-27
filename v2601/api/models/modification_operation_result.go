/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="modification_operation_result.go">
 *   Copyright (c) 2026 Aspose.Words for Cloud
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

// result of the operation which modifies the original document and saves the result.

type IModificationOperationResult interface {
    IsModificationOperationResult() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetDest() IFileLink
    SetDest(value IFileLink)
    GetSource() IFileLink
    SetSource(value IFileLink)
}

type ModificationOperationResult struct {
    // result of the operation which modifies the original document and saves the result.
    Dest IFileLink `json:"Dest,omitempty"`

    // result of the operation which modifies the original document and saves the result.
    Source IFileLink `json:"Source,omitempty"`
}

func (ModificationOperationResult) IsModificationOperationResult() bool {
    return true
}


func (obj *ModificationOperationResult) Initialize() {
    if (obj.Dest != nil) {
        obj.Dest.Initialize()
    }

    if (obj.Source != nil) {
        obj.Source.Initialize()
    }


}

func (obj *ModificationOperationResult) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Dest"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFileLink = new(FileLink)
            modelInstance.Deserialize(parsedValue)
            obj.Dest = modelInstance
        }

    } else if jsonValue, exists := json["dest"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFileLink = new(FileLink)
            modelInstance.Deserialize(parsedValue)
            obj.Dest = modelInstance
        }

    }

    if jsonValue, exists := json["Source"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFileLink = new(FileLink)
            modelInstance.Deserialize(parsedValue)
            obj.Source = modelInstance
        }

    } else if jsonValue, exists := json["source"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFileLink = new(FileLink)
            modelInstance.Deserialize(parsedValue)
            obj.Source = modelInstance
        }

    }
}

func (obj *ModificationOperationResult) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *ModificationOperationResult) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Dest != nil {
        if err := obj.Dest.Validate(); err != nil {
            return err
        }
    }
    if obj.Source != nil {
        if err := obj.Source.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *ModificationOperationResult) GetDest() IFileLink {
    return obj.Dest
}

func (obj *ModificationOperationResult) SetDest(value IFileLink) {
    obj.Dest = value
}

func (obj *ModificationOperationResult) GetSource() IFileLink {
    return obj.Source
}

func (obj *ModificationOperationResult) SetSource(value IFileLink) {
    obj.Source = value
}

