/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="comment_insert.go">
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

// Comment insert.

type ICommentInsert interface {
    IsCommentInsert() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetRangeStart() INewDocumentPosition
    SetRangeStart(value INewDocumentPosition)
    GetRangeEnd() INewDocumentPosition
    SetRangeEnd(value INewDocumentPosition)
    GetAuthor() *string
    SetAuthor(value *string)
    GetInitial() *string
    SetInitial(value *string)
    GetDateTime() *Time
    SetDateTime(value *Time)
    GetText() *string
    SetText(value *string)
}

type CommentInsert struct {
    // Comment insert.
    RangeStart INewDocumentPosition `json:"RangeStart,omitempty"`

    // Comment insert.
    RangeEnd INewDocumentPosition `json:"RangeEnd,omitempty"`

    // Comment insert.
    Author *string `json:"Author,omitempty"`

    // Comment insert.
    Initial *string `json:"Initial,omitempty"`

    // Comment insert.
    DateTime *Time `json:"DateTime,omitempty"`

    // Comment insert.
    Text *string `json:"Text,omitempty"`
}

func (CommentInsert) IsCommentInsert() bool {
    return true
}

func (CommentInsert) IsCommentBase() bool {
    return true
}

func (obj *CommentInsert) Initialize() {
    if (obj.RangeStart != nil) {
        obj.RangeStart.Initialize()
    }

    if (obj.RangeEnd != nil) {
        obj.RangeEnd.Initialize()
    }


}

func (obj *CommentInsert) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RangeStart"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance INewDocumentPosition = new(NewDocumentPosition)
            modelInstance.Deserialize(parsedValue)
            obj.RangeStart = modelInstance
        }

    } else if jsonValue, exists := json["rangeStart"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance INewDocumentPosition = new(NewDocumentPosition)
            modelInstance.Deserialize(parsedValue)
            obj.RangeStart = modelInstance
        }

    }

    if jsonValue, exists := json["RangeEnd"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance INewDocumentPosition = new(NewDocumentPosition)
            modelInstance.Deserialize(parsedValue)
            obj.RangeEnd = modelInstance
        }

    } else if jsonValue, exists := json["rangeEnd"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance INewDocumentPosition = new(NewDocumentPosition)
            modelInstance.Deserialize(parsedValue)
            obj.RangeEnd = modelInstance
        }

    }

    if jsonValue, exists := json["Author"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Author = &parsedValue
        }

    } else if jsonValue, exists := json["author"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Author = &parsedValue
        }

    }

    if jsonValue, exists := json["Initial"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Initial = &parsedValue
        }

    } else if jsonValue, exists := json["initial"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Initial = &parsedValue
        }

    }

    if jsonValue, exists := json["DateTime"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DateTime = new(Time)
            obj.DateTime.Parse(parsedValue)
        }

    } else if jsonValue, exists := json["dateTime"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DateTime = new(Time)
            obj.DateTime.Parse(parsedValue)
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
}

func (obj *CommentInsert) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *CommentInsert) GetRangeStart() INewDocumentPosition {
    return obj.RangeStart
}

func (obj *CommentInsert) SetRangeStart(value INewDocumentPosition) {
    obj.RangeStart = value
}

func (obj *CommentInsert) GetRangeEnd() INewDocumentPosition {
    return obj.RangeEnd
}

func (obj *CommentInsert) SetRangeEnd(value INewDocumentPosition) {
    obj.RangeEnd = value
}

func (obj *CommentInsert) GetAuthor() *string {
    return obj.Author
}

func (obj *CommentInsert) SetAuthor(value *string) {
    obj.Author = value
}

func (obj *CommentInsert) GetInitial() *string {
    return obj.Initial
}

func (obj *CommentInsert) SetInitial(value *string) {
    obj.Initial = value
}

func (obj *CommentInsert) GetDateTime() *Time {
    return obj.DateTime
}

func (obj *CommentInsert) SetDateTime(value *Time) {
    obj.DateTime = value
}

func (obj *CommentInsert) GetText() *string {
    return obj.Text
}

func (obj *CommentInsert) SetText(value *string) {
    obj.Text = value
}

