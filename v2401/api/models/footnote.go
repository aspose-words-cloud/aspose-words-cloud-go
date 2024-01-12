/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="footnote.go">
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

// DTO container with a footnote.

type IFootnote interface {
    IsFootnote() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetNodeId() *string
    SetNodeId(value *string)
    GetPosition() IDocumentPosition
    SetPosition(value IDocumentPosition)
    GetFootnoteType() *string
    SetFootnoteType(value *string)
    GetReferenceMark() *string
    SetReferenceMark(value *string)
    GetText() *string
    SetText(value *string)
    GetContent() IStoryChildNodes
    SetContent(value IStoryChildNodes)
}

type Footnote struct {
    // DTO container with a footnote.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with a footnote.
    NodeId *string `json:"NodeId,omitempty"`

    // DTO container with a footnote.
    Position IDocumentPosition `json:"Position,omitempty"`

    // DTO container with a footnote.
    FootnoteType *string `json:"FootnoteType,omitempty"`

    // DTO container with a footnote.
    ReferenceMark *string `json:"ReferenceMark,omitempty"`

    // DTO container with a footnote.
    Text *string `json:"Text,omitempty"`

    // DTO container with a footnote.
    Content IStoryChildNodes `json:"Content,omitempty"`
}

func (Footnote) IsFootnote() bool {
    return true
}

func (Footnote) IsFootnoteLink() bool {
    return true
}

func (Footnote) IsNodeLink() bool {
    return true
}

func (Footnote) IsLinkElement() bool {
    return true
}

func (obj *Footnote) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.Position != nil) {
        obj.Position.Initialize()
    }

    if (obj.Content != nil) {
        obj.Content.Initialize()
    }


}

func (obj *Footnote) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["NodeId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NodeId = &parsedValue
        }

    } else if jsonValue, exists := json["nodeId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NodeId = &parsedValue
        }

    }

    if jsonValue, exists := json["Position"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IDocumentPosition = new(DocumentPosition)
            modelInstance.Deserialize(parsedValue)
            obj.Position = modelInstance
        }

    } else if jsonValue, exists := json["position"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IDocumentPosition = new(DocumentPosition)
            modelInstance.Deserialize(parsedValue)
            obj.Position = modelInstance
        }

    }

    if jsonValue, exists := json["FootnoteType"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FootnoteType = &parsedValue
        }

    } else if jsonValue, exists := json["footnoteType"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FootnoteType = &parsedValue
        }

    }

    if jsonValue, exists := json["ReferenceMark"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ReferenceMark = &parsedValue
        }

    } else if jsonValue, exists := json["referenceMark"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ReferenceMark = &parsedValue
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

func (obj *Footnote) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *Footnote) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Link != nil {
        if err := obj.Link.Validate(); err != nil {
            return err
        }
    }
    if obj.Position != nil {
        if err := obj.Position.Validate(); err != nil {
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

func (obj *Footnote) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *Footnote) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *Footnote) GetNodeId() *string {
    return obj.NodeId
}

func (obj *Footnote) SetNodeId(value *string) {
    obj.NodeId = value
}

func (obj *Footnote) GetPosition() IDocumentPosition {
    return obj.Position
}

func (obj *Footnote) SetPosition(value IDocumentPosition) {
    obj.Position = value
}

func (obj *Footnote) GetFootnoteType() *string {
    return obj.FootnoteType
}

func (obj *Footnote) SetFootnoteType(value *string) {
    obj.FootnoteType = value
}

func (obj *Footnote) GetReferenceMark() *string {
    return obj.ReferenceMark
}

func (obj *Footnote) SetReferenceMark(value *string) {
    obj.ReferenceMark = value
}

func (obj *Footnote) GetText() *string {
    return obj.Text
}

func (obj *Footnote) SetText(value *string) {
    obj.Text = value
}

func (obj *Footnote) GetContent() IStoryChildNodes {
    return obj.Content
}

func (obj *Footnote) SetContent(value IStoryChildNodes) {
    obj.Content = value
}

