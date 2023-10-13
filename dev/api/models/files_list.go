/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="files_list.go">
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

// Files list.

type IFilesList interface {
    IsFilesList() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetValue() []IStorageFile
    SetValue(value []IStorageFile)
}

type FilesList struct {
    // Files list.
    Value []IStorageFile `json:"Value,omitempty"`
}

func (FilesList) IsFilesList() bool {
    return true
}


func (obj *FilesList) Initialize() {
    if (obj.Value != nil) {
        for _, objElementValue := range obj.Value {
            objElementValue.Initialize()
        }
    }

}

func (obj *FilesList) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Value"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.Value = make([]IStorageFile, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IStorageFile = new(StorageFile)
                    modelElementInstance.Deserialize(elementValue)
                    obj.Value = append(obj.Value, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["value"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.Value = make([]IStorageFile, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IStorageFile = new(StorageFile)
                    modelElementInstance.Deserialize(elementValue)
                    obj.Value = append(obj.Value, modelElementInstance)
                }

            }
        }

    }
}

func (obj *FilesList) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *FilesList) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    return nil;
}

func (obj *FilesList) GetValue() []IStorageFile {
    return obj.Value
}

func (obj *FilesList) SetValue(value []IStorageFile) {
    obj.Value = value
}

