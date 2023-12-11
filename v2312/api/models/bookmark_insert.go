/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="bookmark_insert.go">
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

// Represents a bookmark to insert.

type IBookmarkInsert interface {
    IsBookmarkInsert() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetName() *string
    SetName(value *string)
    GetText() *string
    SetText(value *string)
    GetStartRange() INewDocumentPosition
    SetStartRange(value INewDocumentPosition)
    GetEndRange() INewDocumentPosition
    SetEndRange(value INewDocumentPosition)
}

type BookmarkInsert struct {
    // Represents a bookmark to insert.
    Name *string `json:"Name,omitempty"`

    // Represents a bookmark to insert.
    Text *string `json:"Text,omitempty"`

    // Represents a bookmark to insert.
    StartRange INewDocumentPosition `json:"StartRange,omitempty"`

    // Represents a bookmark to insert.
    EndRange INewDocumentPosition `json:"EndRange,omitempty"`
}

func (BookmarkInsert) IsBookmarkInsert() bool {
    return true
}


func (obj *BookmarkInsert) Initialize() {
    if (obj.StartRange != nil) {
        obj.StartRange.Initialize()
    }

    if (obj.EndRange != nil) {
        obj.EndRange.Initialize()
    }


}

func (obj *BookmarkInsert) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Name"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Name = &parsedValue
        }

    } else if jsonValue, exists := json["name"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Name = &parsedValue
        }

    }

    if jsonValue, exists := json["Text"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Text = &parsedValue
        }

    } else if jsonValue, exists := json["text"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Text = &parsedValue
        }

    }

    if jsonValue, exists := json["StartRange"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance INewDocumentPosition = new(NewDocumentPosition)
            modelInstance.Deserialize(parsedValue)
            obj.StartRange = modelInstance
        }

    } else if jsonValue, exists := json["startRange"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance INewDocumentPosition = new(NewDocumentPosition)
            modelInstance.Deserialize(parsedValue)
            obj.StartRange = modelInstance
        }

    }

    if jsonValue, exists := json["EndRange"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance INewDocumentPosition = new(NewDocumentPosition)
            modelInstance.Deserialize(parsedValue)
            obj.EndRange = modelInstance
        }

    } else if jsonValue, exists := json["endRange"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance INewDocumentPosition = new(NewDocumentPosition)
            modelInstance.Deserialize(parsedValue)
            obj.EndRange = modelInstance
        }

    }
}

func (obj *BookmarkInsert) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *BookmarkInsert) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Name == nil {
        return errors.New("Property Name in BookmarkInsert is required.")
    }
    if obj.Text == nil {
        return errors.New("Property Text in BookmarkInsert is required.")
    }
    if obj.StartRange == nil {
        return errors.New("Property StartRange in BookmarkInsert is required.")
    }
    if obj.EndRange == nil {
        return errors.New("Property EndRange in BookmarkInsert is required.")
    }
    if obj.StartRange != nil {
        if err := obj.StartRange.Validate(); err != nil {
            return err
        }
    }
    if obj.EndRange != nil {
        if err := obj.EndRange.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *BookmarkInsert) GetName() *string {
    return obj.Name
}

func (obj *BookmarkInsert) SetName(value *string) {
    obj.Name = value
}

func (obj *BookmarkInsert) GetText() *string {
    return obj.Text
}

func (obj *BookmarkInsert) SetText(value *string) {
    obj.Text = value
}

func (obj *BookmarkInsert) GetStartRange() INewDocumentPosition {
    return obj.StartRange
}

func (obj *BookmarkInsert) SetStartRange(value INewDocumentPosition) {
    obj.StartRange = value
}

func (obj *BookmarkInsert) GetEndRange() INewDocumentPosition {
    return obj.EndRange
}

func (obj *BookmarkInsert) SetEndRange(value INewDocumentPosition) {
    obj.EndRange = value
}

