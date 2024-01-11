/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="image_entry.go">
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

// Represents a image which will be appended to the original resource image or document.

type IImageEntry interface {
    IsImageEntry() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetFileReference() IFileReference
    SetFileReference(value IFileReference)
}

type ImageEntry struct {
    // Represents a image which will be appended to the original resource image or document.
    FileReference IFileReference `json:"FileReference,omitempty"`
}

func (ImageEntry) IsImageEntry() bool {
    return true
}

func (ImageEntry) IsBaseEntry() bool {
    return true
}

func (obj *ImageEntry) Initialize() {
    if (obj.FileReference != nil) {
        obj.FileReference.Initialize()
    }


}

func (obj *ImageEntry) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["FileReference"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFileReference = new(FileReference)
            modelInstance.Deserialize(parsedValue)
            obj.FileReference = modelInstance
        }

    } else if jsonValue, exists := json["fileReference"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFileReference = new(FileReference)
            modelInstance.Deserialize(parsedValue)
            obj.FileReference = modelInstance
        }

    }
}

func (obj *ImageEntry) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    if (obj.FileReference != nil) {
        resultFilesContent = obj.FileReference.CollectFilesContent(resultFilesContent)
    }

    return resultFilesContent
}

func (obj *ImageEntry) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.FileReference == nil {
        return errors.New("Property FileReference in ImageEntry is required.")
    }
    if obj.FileReference != nil {
        if err := obj.FileReference.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *ImageEntry) GetFileReference() IFileReference {
    return obj.FileReference
}

func (obj *ImageEntry) SetFileReference(value IFileReference) {
    obj.FileReference = value
}

