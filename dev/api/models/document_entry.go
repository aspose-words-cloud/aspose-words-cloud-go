/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="document_entry.go">
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

// Represents a document which will be appended to the original resource document.

type IDocumentEntry interface {
    IsDocumentEntry() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetFileReference() IFileReference
    SetFileReference(value IFileReference)
    GetEncryptedPassword() *string
    SetEncryptedPassword(value *string)
    GetImportFormatMode() *string
    SetImportFormatMode(value *string)
}

type DocumentEntry struct {
    // Represents a document which will be appended to the original resource document.
    FileReference IFileReference `json:"FileReference,omitempty"`

    // Represents a document which will be appended to the original resource document.
    EncryptedPassword *string `json:"EncryptedPassword,omitempty"`

    // Represents a document which will be appended to the original resource document.
    ImportFormatMode *string `json:"ImportFormatMode,omitempty"`
}

func (DocumentEntry) IsDocumentEntry() bool {
    return true
}

func (DocumentEntry) IsBaseEntry() bool {
    return true
}

func (obj *DocumentEntry) Initialize() {
    if (obj.FileReference != nil) {
        obj.FileReference.Initialize()
    }


}

func (obj *DocumentEntry) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["EncryptedPassword"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.EncryptedPassword = &parsedValue
        }

    } else if jsonValue, exists := json["encryptedPassword"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.EncryptedPassword = &parsedValue
        }

    }

    if jsonValue, exists := json["ImportFormatMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ImportFormatMode = &parsedValue
        }

    } else if jsonValue, exists := json["importFormatMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ImportFormatMode = &parsedValue
        }

    }
}

func (obj *DocumentEntry) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    if (obj.FileReference != nil) {
        resultFilesContent = obj.FileReference.CollectFilesContent(resultFilesContent)
    }



    return resultFilesContent
}

func (obj *DocumentEntry) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.FileReference == nil {
        return errors.New("Property FileReference in DocumentEntry is required.")
    }
    if obj.ImportFormatMode == nil {
        return errors.New("Property ImportFormatMode in DocumentEntry is required.")
    }
    if obj.FileReference != nil {
        if err := obj.FileReference.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *DocumentEntry) GetFileReference() IFileReference {
    return obj.FileReference
}

func (obj *DocumentEntry) SetFileReference(value IFileReference) {
    obj.FileReference = value
}

func (obj *DocumentEntry) GetEncryptedPassword() *string {
    return obj.EncryptedPassword
}

func (obj *DocumentEntry) SetEncryptedPassword(value *string) {
    obj.EncryptedPassword = value
}

func (obj *DocumentEntry) GetImportFormatMode() *string {
    return obj.ImportFormatMode
}

func (obj *DocumentEntry) SetImportFormatMode(value *string) {
    obj.ImportFormatMode = value
}

