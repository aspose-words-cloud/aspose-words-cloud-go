/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="comment.go">
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

// DTO container with a comment.

type IComment interface {
    IsComment() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetRangeStart() IDocumentPosition
    SetRangeStart(value IDocumentPosition)
    GetRangeEnd() IDocumentPosition
    SetRangeEnd(value IDocumentPosition)
    GetAuthor() *string
    SetAuthor(value *string)
    GetInitial() *string
    SetInitial(value *string)
    GetDateTime() *Time
    SetDateTime(value *Time)
    GetText() *string
    SetText(value *string)
    GetContent() IStoryChildNodes
    SetContent(value IStoryChildNodes)
}

type Comment struct {
    // DTO container with a comment.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with a comment.
    RangeStart IDocumentPosition `json:"RangeStart,omitempty"`

    // DTO container with a comment.
    RangeEnd IDocumentPosition `json:"RangeEnd,omitempty"`

    // DTO container with a comment.
    Author *string `json:"Author,omitempty"`

    // DTO container with a comment.
    Initial *string `json:"Initial,omitempty"`

    // DTO container with a comment.
    DateTime *Time `json:"DateTime,omitempty"`

    // DTO container with a comment.
    Text *string `json:"Text,omitempty"`

    // DTO container with a comment.
    Content IStoryChildNodes `json:"Content,omitempty"`
}

func (Comment) IsComment() bool {
    return true
}

func (Comment) IsCommentLink() bool {
    return true
}

func (Comment) IsLinkElement() bool {
    return true
}

func (obj *Comment) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.RangeStart != nil) {
        obj.RangeStart.Initialize()
    }

    if (obj.RangeEnd != nil) {
        obj.RangeEnd.Initialize()
    }

    if (obj.Content != nil) {
        obj.Content.Initialize()
    }


}

func (obj *Comment) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Link"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IWordsApiLink = new(WordsApiLink)
            modelInstance.Deserialize(parsedValue)
            obj.Link = modelInstance
        }

    } else if jsonValue, exists := json["link"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IWordsApiLink = new(WordsApiLink)
            modelInstance.Deserialize(parsedValue)
            obj.Link = modelInstance
        }

    }

    if jsonValue, exists := json["RangeStart"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IDocumentPosition = new(DocumentPosition)
            modelInstance.Deserialize(parsedValue)
            obj.RangeStart = modelInstance
        }

    } else if jsonValue, exists := json["rangeStart"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IDocumentPosition = new(DocumentPosition)
            modelInstance.Deserialize(parsedValue)
            obj.RangeStart = modelInstance
        }

    }

    if jsonValue, exists := json["RangeEnd"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IDocumentPosition = new(DocumentPosition)
            modelInstance.Deserialize(parsedValue)
            obj.RangeEnd = modelInstance
        }

    } else if jsonValue, exists := json["rangeEnd"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IDocumentPosition = new(DocumentPosition)
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

    if jsonValue, exists := json["Content"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IStoryChildNodes = new(StoryChildNodes)
            modelInstance.Deserialize(parsedValue)
            obj.Content = modelInstance
        }

    } else if jsonValue, exists := json["content"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IStoryChildNodes = new(StoryChildNodes)
            modelInstance.Deserialize(parsedValue)
            obj.Content = modelInstance
        }

    }
}

func (obj *Comment) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *Comment) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Link != nil {
        if err := obj.Link.Validate(); err != nil {
            return err
        }
    }
    if obj.RangeStart != nil {
        if err := obj.RangeStart.Validate(); err != nil {
            return err
        }
    }
    if obj.RangeEnd != nil {
        if err := obj.RangeEnd.Validate(); err != nil {
            return err
        }
    }
    if obj.Content != nil {
        if err := obj.Content.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *Comment) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *Comment) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *Comment) GetRangeStart() IDocumentPosition {
    return obj.RangeStart
}

func (obj *Comment) SetRangeStart(value IDocumentPosition) {
    obj.RangeStart = value
}

func (obj *Comment) GetRangeEnd() IDocumentPosition {
    return obj.RangeEnd
}

func (obj *Comment) SetRangeEnd(value IDocumentPosition) {
    obj.RangeEnd = value
}

func (obj *Comment) GetAuthor() *string {
    return obj.Author
}

func (obj *Comment) SetAuthor(value *string) {
    obj.Author = value
}

func (obj *Comment) GetInitial() *string {
    return obj.Initial
}

func (obj *Comment) SetInitial(value *string) {
    obj.Initial = value
}

func (obj *Comment) GetDateTime() *Time {
    return obj.DateTime
}

func (obj *Comment) SetDateTime(value *Time) {
    obj.DateTime = value
}

func (obj *Comment) GetText() *string {
    return obj.Text
}

func (obj *Comment) SetText(value *string) {
    obj.Text = value
}

func (obj *Comment) GetContent() IStoryChildNodes {
    return obj.Content
}

func (obj *Comment) SetContent(value IStoryChildNodes) {
    obj.Content = value
}

