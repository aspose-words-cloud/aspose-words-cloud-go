/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="file_link.go">
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

// Provides information for the file link.

type IFileLink interface {
    IsFileLink() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetHref() *string
    SetHref(value *string)
    GetRel() *string
    SetRel(value *string)
    GetTitle() *string
    SetTitle(value *string)
    GetType() *string
    SetType(value *string)
}

type FileLink struct {
    // Provides information for the file link.
    Href *string `json:"Href,omitempty"`

    // Provides information for the file link.
    Rel *string `json:"Rel,omitempty"`

    // Provides information for the file link.
    Title *string `json:"Title,omitempty"`

    // Provides information for the file link.
    Type *string `json:"Type,omitempty"`
}

func (FileLink) IsFileLink() bool {
    return true
}

func (FileLink) IsLink() bool {
    return true
}

func (obj *FileLink) Initialize() {
}

func (obj *FileLink) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Href"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Href = &parsedValue
        }

    } else if jsonValue, exists := json["href"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Href = &parsedValue
        }

    }

    if jsonValue, exists := json["Rel"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Rel = &parsedValue
        }

    } else if jsonValue, exists := json["rel"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Rel = &parsedValue
        }

    }

    if jsonValue, exists := json["Title"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Title = &parsedValue
        }

    } else if jsonValue, exists := json["title"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Title = &parsedValue
        }

    }

    if jsonValue, exists := json["Type"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Type = &parsedValue
        }

    } else if jsonValue, exists := json["type"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Type = &parsedValue
        }

    }
}

func (obj *FileLink) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *FileLink) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Href == nil {
        return errors.New("Property Href in FileLink is required.")
    }

    if obj.Rel == nil {
        return errors.New("Property Rel in FileLink is required.")
    }

    if obj.Title == nil {
        return errors.New("Property Title in FileLink is required.")
    }

    if obj.Type == nil {
        return errors.New("Property Type in FileLink is required.")
    }

    return nil;
}

func (obj *FileLink) GetHref() *string {
    return obj.Href
}

func (obj *FileLink) SetHref(value *string) {
    obj.Href = value
}

func (obj *FileLink) GetRel() *string {
    return obj.Rel
}

func (obj *FileLink) SetRel(value *string) {
    obj.Rel = value
}

func (obj *FileLink) GetTitle() *string {
    return obj.Title
}

func (obj *FileLink) SetTitle(value *string) {
    obj.Title = value
}

func (obj *FileLink) GetType() *string {
    return obj.Type
}

func (obj *FileLink) SetType(value *string) {
    obj.Type = value
}

