/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="revision_collection.go">
 *   Copyright (c) 2025 Aspose.Words for Cloud
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

// RevisionCollection DTO.

type IRevisionCollection interface {
    IsRevisionCollection() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRevisions() []IRevision
    SetRevisions(value []IRevision)
}

type RevisionCollection struct {
    // RevisionCollection DTO.
    Revisions []IRevision `json:"Revisions,omitempty"`
}

func (RevisionCollection) IsRevisionCollection() bool {
    return true
}


func (obj *RevisionCollection) Initialize() {
    if (obj.Revisions != nil) {
        for _, objElementRevisions := range obj.Revisions {
            objElementRevisions.Initialize()
        }
    }

}

func (obj *RevisionCollection) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Revisions"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.Revisions = make([]IRevision, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IRevision = new(Revision)
                    modelElementInstance.Deserialize(elementValue)
                    obj.Revisions = append(obj.Revisions, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["revisions"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.Revisions = make([]IRevision, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IRevision = new(Revision)
                    modelElementInstance.Deserialize(elementValue)
                    obj.Revisions = append(obj.Revisions, modelElementInstance)
                }

            }
        }

    }
}

func (obj *RevisionCollection) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *RevisionCollection) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Revisions != nil {
        for _, elementRevisions := range obj.Revisions {
            if elementRevisions != nil {
                if err := elementRevisions.Validate(); err != nil {
                    return err
                }
            }
        }
    }

    return nil;
}

func (obj *RevisionCollection) GetRevisions() []IRevision {
    return obj.Revisions
}

func (obj *RevisionCollection) SetRevisions(value []IRevision) {
    obj.Revisions = value
}

