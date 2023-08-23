/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="outline_options_data.go">
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

// Container class for outline options.

type IOutlineOptionsData interface {
    IsOutlineOptionsData() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetBookmarksOutlineLevels() []IBookmarksOutlineLevelData
    SetBookmarksOutlineLevels(value []IBookmarksOutlineLevelData)
    GetCreateMissingOutlineLevels() *bool
    SetCreateMissingOutlineLevels(value *bool)
    GetCreateOutlinesForHeadingsInTables() *bool
    SetCreateOutlinesForHeadingsInTables(value *bool)
    GetDefaultBookmarksOutlineLevel() *int32
    SetDefaultBookmarksOutlineLevel(value *int32)
    GetExpandedOutlineLevels() *int32
    SetExpandedOutlineLevels(value *int32)
    GetHeadingsOutlineLevels() *int32
    SetHeadingsOutlineLevels(value *int32)
}

type OutlineOptionsData struct {
    // Container class for outline options.
    BookmarksOutlineLevels []IBookmarksOutlineLevelData

    // Container class for outline options.
    CreateMissingOutlineLevels *bool

    // Container class for outline options.
    CreateOutlinesForHeadingsInTables *bool

    // Container class for outline options.
    DefaultBookmarksOutlineLevel *int32

    // Container class for outline options.
    ExpandedOutlineLevels *int32

    // Container class for outline options.
    HeadingsOutlineLevels *int32
}

func (OutlineOptionsData) IsOutlineOptionsData() bool {
    return true
}


func (obj *OutlineOptionsData) Initialize() {
    if (obj.BookmarksOutlineLevels != nil) {
        for _, objElementBookmarksOutlineLevels := range obj.BookmarksOutlineLevels {
            objElementBookmarksOutlineLevels.Initialize()
        }
    }

}

func (obj *OutlineOptionsData) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["BookmarksOutlineLevels"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IBookmarksOutlineLevelData = new(BookmarksOutlineLevelData)
                    modelElementInstance.Deserialize(elementValue)
                    obj.BookmarksOutlineLevels = append(obj.BookmarksOutlineLevels, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["bookmarksOutlineLevels"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IBookmarksOutlineLevelData = new(BookmarksOutlineLevelData)
                    modelElementInstance.Deserialize(elementValue)
                    obj.BookmarksOutlineLevels = append(obj.BookmarksOutlineLevels, modelElementInstance)
                }

            }
        }

    }

    if jsonValue, exists := json["CreateMissingOutlineLevels"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.CreateMissingOutlineLevels = &parsedValue
        }

    } else if jsonValue, exists := json["createMissingOutlineLevels"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.CreateMissingOutlineLevels = &parsedValue
        }

    }

    if jsonValue, exists := json["CreateOutlinesForHeadingsInTables"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.CreateOutlinesForHeadingsInTables = &parsedValue
        }

    } else if jsonValue, exists := json["createOutlinesForHeadingsInTables"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.CreateOutlinesForHeadingsInTables = &parsedValue
        }

    }

    if jsonValue, exists := json["DefaultBookmarksOutlineLevel"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.DefaultBookmarksOutlineLevel = new(int32)
            *obj.DefaultBookmarksOutlineLevel = int32(parsedValue)
        }

    } else if jsonValue, exists := json["defaultBookmarksOutlineLevel"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.DefaultBookmarksOutlineLevel = new(int32)
            *obj.DefaultBookmarksOutlineLevel = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["ExpandedOutlineLevels"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ExpandedOutlineLevels = new(int32)
            *obj.ExpandedOutlineLevels = int32(parsedValue)
        }

    } else if jsonValue, exists := json["expandedOutlineLevels"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ExpandedOutlineLevels = new(int32)
            *obj.ExpandedOutlineLevels = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["HeadingsOutlineLevels"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.HeadingsOutlineLevels = new(int32)
            *obj.HeadingsOutlineLevels = int32(parsedValue)
        }

    } else if jsonValue, exists := json["headingsOutlineLevels"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.HeadingsOutlineLevels = new(int32)
            *obj.HeadingsOutlineLevels = int32(parsedValue)
        }

    }
}

func (obj *OutlineOptionsData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *OutlineOptionsData) GetBookmarksOutlineLevels() []IBookmarksOutlineLevelData {
    return obj.BookmarksOutlineLevels
}

func (obj *OutlineOptionsData) SetBookmarksOutlineLevels(value []IBookmarksOutlineLevelData) {
    obj.BookmarksOutlineLevels = value
}

func (obj *OutlineOptionsData) GetCreateMissingOutlineLevels() *bool {
    return obj.CreateMissingOutlineLevels
}

func (obj *OutlineOptionsData) SetCreateMissingOutlineLevels(value *bool) {
    obj.CreateMissingOutlineLevels = value
}

func (obj *OutlineOptionsData) GetCreateOutlinesForHeadingsInTables() *bool {
    return obj.CreateOutlinesForHeadingsInTables
}

func (obj *OutlineOptionsData) SetCreateOutlinesForHeadingsInTables(value *bool) {
    obj.CreateOutlinesForHeadingsInTables = value
}

func (obj *OutlineOptionsData) GetDefaultBookmarksOutlineLevel() *int32 {
    return obj.DefaultBookmarksOutlineLevel
}

func (obj *OutlineOptionsData) SetDefaultBookmarksOutlineLevel(value *int32) {
    obj.DefaultBookmarksOutlineLevel = value
}

func (obj *OutlineOptionsData) GetExpandedOutlineLevels() *int32 {
    return obj.ExpandedOutlineLevels
}

func (obj *OutlineOptionsData) SetExpandedOutlineLevels(value *int32) {
    obj.ExpandedOutlineLevels = value
}

func (obj *OutlineOptionsData) GetHeadingsOutlineLevels() *int32 {
    return obj.HeadingsOutlineLevels
}

func (obj *OutlineOptionsData) SetHeadingsOutlineLevels(value *int32) {
    obj.HeadingsOutlineLevels = value
}

