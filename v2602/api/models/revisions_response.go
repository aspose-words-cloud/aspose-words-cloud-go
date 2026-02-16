/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="revisions_response.go">
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

// The revision response.

type IRevisionsResponse interface {
    IsRevisionsResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetRevisions() IRevisionCollection
    SetRevisions(value IRevisionCollection)
}

type RevisionsResponse struct {
    // The revision response.
    RequestId *string `json:"RequestId,omitempty"`

    // The revision response.
    Revisions IRevisionCollection `json:"Revisions,omitempty"`
}

func (RevisionsResponse) IsRevisionsResponse() bool {
    return true
}

func (RevisionsResponse) IsWordsResponse() bool {
    return true
}

func (obj *RevisionsResponse) Initialize() {
    if (obj.Revisions != nil) {
        obj.Revisions.Initialize()
    }


}

func (obj *RevisionsResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["Revisions"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IRevisionCollection = new(RevisionCollection)
            modelInstance.Deserialize(parsedValue)
            obj.Revisions = modelInstance
        }

    } else if jsonValue, exists := json["revisions"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IRevisionCollection = new(RevisionCollection)
            modelInstance.Deserialize(parsedValue)
            obj.Revisions = modelInstance
        }

    }
}

func (obj *RevisionsResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *RevisionsResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Revisions != nil {
        if err := obj.Revisions.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *RevisionsResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *RevisionsResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *RevisionsResponse) GetRevisions() IRevisionCollection {
    return obj.Revisions
}

func (obj *RevisionsResponse) SetRevisions(value IRevisionCollection) {
    obj.Revisions = value
}

