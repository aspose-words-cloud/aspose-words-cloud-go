/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="bookmarks_outline_level_data.go">
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

// Container class for individual bookmarks outline level.

type IBookmarksOutlineLevelData interface {
    IsBookmarksOutlineLevelData() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetBookmarksOutlineLevel() *int32
    SetBookmarksOutlineLevel(value *int32)
    GetName() *string
    SetName(value *string)
}

type BookmarksOutlineLevelData struct {
    // Container class for individual bookmarks outline level.
    BookmarksOutlineLevel *int32 `json:"BookmarksOutlineLevel,omitempty"`

    // Container class for individual bookmarks outline level.
    Name *string `json:"Name,omitempty"`
}

func (BookmarksOutlineLevelData) IsBookmarksOutlineLevelData() bool {
    return true
}


func (obj *BookmarksOutlineLevelData) Initialize() {
}

func (obj *BookmarksOutlineLevelData) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["BookmarksOutlineLevel"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.BookmarksOutlineLevel = new(int32)
            *obj.BookmarksOutlineLevel = int32(parsedValue)
        }

    } else if jsonValue, exists := json["bookmarksOutlineLevel"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.BookmarksOutlineLevel = new(int32)
            *obj.BookmarksOutlineLevel = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["Name"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Name = &parsedValue
        }

    } else if jsonValue, exists := json["name"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Name = &parsedValue
        }

    }
}

func (obj *BookmarksOutlineLevelData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *BookmarksOutlineLevelData) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.BookmarksOutlineLevel == nil {
        return errors.New("Property BookmarksOutlineLevel in BookmarksOutlineLevelData is required.")
    }
    if obj.BookmarksOutlineLevel == nil {
        return errors.New("Property BookmarksOutlineLevel in BookmarksOutlineLevelData is required.")
    }

    if obj.Name == nil {
        return errors.New("Property Name in BookmarksOutlineLevelData is required.")
    }

    return nil;
}

func (obj *BookmarksOutlineLevelData) GetBookmarksOutlineLevel() *int32 {
    return obj.BookmarksOutlineLevel
}

func (obj *BookmarksOutlineLevelData) SetBookmarksOutlineLevel(value *int32) {
    obj.BookmarksOutlineLevel = value
}

func (obj *BookmarksOutlineLevelData) GetName() *string {
    return obj.Name
}

func (obj *BookmarksOutlineLevelData) SetName(value *string) {
    obj.Name = value
}

