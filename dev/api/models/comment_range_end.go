/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="comment_range_end.go">
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

// Comment range end link.

type ICommentRangeEnd interface {
    IsCommentRangeEnd() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetNodeId() *string
    SetNodeId(value *string)
    GetCommentLink() ICommentLink
    SetCommentLink(value ICommentLink)
}

type CommentRangeEnd struct {
    // Comment range end link.
    Link IWordsApiLink `json:"Link,omitempty"`

    // Comment range end link.
    NodeId *string `json:"NodeId,omitempty"`

    // Comment range end link.
    CommentLink ICommentLink `json:"CommentLink,omitempty"`
}

func (CommentRangeEnd) IsCommentRangeEnd() bool {
    return true
}

func (CommentRangeEnd) IsNodeLink() bool {
    return true
}

func (CommentRangeEnd) IsLinkElement() bool {
    return true
}

func (obj *CommentRangeEnd) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.CommentLink != nil) {
        obj.CommentLink.Initialize()
    }


}

func (obj *CommentRangeEnd) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["CommentLink"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ICommentLink = nil
            if jsonType, found := parsedValue["$type"]; found {
                jsonTypeStr := jsonType.(string)
                if jsonTypeStr == "Comment, _" { modelInstance = new(Comment) }
            }

            if modelInstance == nil { modelInstance = new(CommentLink) }
            modelInstance.Deserialize(parsedValue)
            obj.CommentLink = modelInstance
        }

    } else if jsonValue, exists := json["commentLink"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ICommentLink = nil
            if jsonType, found := parsedValue["$type"]; found {
                jsonTypeStr := jsonType.(string)
                if jsonTypeStr == "Comment, _" { modelInstance = new(Comment) }
            }

            if modelInstance == nil { modelInstance = new(CommentLink) }
            modelInstance.Deserialize(parsedValue)
            obj.CommentLink = modelInstance
        }

    }
}

func (obj *CommentRangeEnd) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *CommentRangeEnd) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Link != nil {
        if err := obj.Link.Validate(); err != nil {
            return err
        }
    }
    if obj.CommentLink != nil {
        if err := obj.CommentLink.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *CommentRangeEnd) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *CommentRangeEnd) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *CommentRangeEnd) GetNodeId() *string {
    return obj.NodeId
}

func (obj *CommentRangeEnd) SetNodeId(value *string) {
    obj.NodeId = value
}

func (obj *CommentRangeEnd) GetCommentLink() ICommentLink {
    return obj.CommentLink
}

func (obj *CommentRangeEnd) SetCommentLink(value ICommentLink) {
    obj.CommentLink = value
}

