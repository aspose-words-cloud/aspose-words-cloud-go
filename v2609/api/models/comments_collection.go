/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="comments_collection.go">
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

// The collection of comments.

type ICommentsCollection interface {
    IsCommentsCollection() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetCommentList() []IComment
    SetCommentList(value []IComment)
}

type CommentsCollection struct {
    // The collection of comments.
    Link IWordsApiLink `json:"Link,omitempty"`

    // The collection of comments.
    CommentList []IComment `json:"CommentList,omitempty"`
}

func (CommentsCollection) IsCommentsCollection() bool {
    return true
}

func (CommentsCollection) IsLinkElement() bool {
    return true
}

func (obj *CommentsCollection) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.CommentList != nil) {
        for _, objElementCommentList := range obj.CommentList {
            objElementCommentList.Initialize()
        }
    }

}

func (obj *CommentsCollection) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["CommentList"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.CommentList = make([]IComment, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IComment = new(Comment)
                    modelElementInstance.Deserialize(elementValue)
                    obj.CommentList = append(obj.CommentList, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["commentList"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.CommentList = make([]IComment, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IComment = new(Comment)
                    modelElementInstance.Deserialize(elementValue)
                    obj.CommentList = append(obj.CommentList, modelElementInstance)
                }

            }
        }

    }
}

func (obj *CommentsCollection) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *CommentsCollection) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Link != nil {
        if err := obj.Link.Validate(); err != nil {
            return err
        }
    }
    if obj.CommentList != nil {
        for _, elementCommentList := range obj.CommentList {
            if elementCommentList != nil {
                if err := elementCommentList.Validate(); err != nil {
                    return err
                }
            }
        }
    }

    return nil;
}

func (obj *CommentsCollection) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *CommentsCollection) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *CommentsCollection) GetCommentList() []IComment {
    return obj.CommentList
}

func (obj *CommentsCollection) SetCommentList(value []IComment) {
    obj.CommentList = value
}

