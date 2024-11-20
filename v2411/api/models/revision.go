/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="revision.go">
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

// Revision Dto.

type IRevision interface {
    IsRevision() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRevisionAuthor() *string
    SetRevisionAuthor(value *string)
    GetRevisionDateTime() *Time
    SetRevisionDateTime(value *Time)
    GetRevisionText() *string
    SetRevisionText(value *string)
    GetRevisionType() *string
    SetRevisionType(value *string)
}

type Revision struct {
    // Revision Dto.
    RevisionAuthor *string `json:"RevisionAuthor,omitempty"`

    // Revision Dto.
    RevisionDateTime *Time `json:"RevisionDateTime,omitempty"`

    // Revision Dto.
    RevisionText *string `json:"RevisionText,omitempty"`

    // Revision Dto.
    RevisionType *string `json:"RevisionType,omitempty"`
}

func (Revision) IsRevision() bool {
    return true
}


func (obj *Revision) Initialize() {
}

func (obj *Revision) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RevisionAuthor"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RevisionAuthor = &parsedValue
        }

    } else if jsonValue, exists := json["revisionAuthor"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RevisionAuthor = &parsedValue
        }

    }

    if jsonValue, exists := json["RevisionDateTime"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RevisionDateTime = new(Time)
            obj.RevisionDateTime.Parse(parsedValue)
        }

    } else if jsonValue, exists := json["revisionDateTime"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RevisionDateTime = new(Time)
            obj.RevisionDateTime.Parse(parsedValue)
        }

    }

    if jsonValue, exists := json["RevisionText"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RevisionText = &parsedValue
        }

    } else if jsonValue, exists := json["revisionText"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RevisionText = &parsedValue
        }

    }

    if jsonValue, exists := json["RevisionType"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RevisionType = &parsedValue
        }

    } else if jsonValue, exists := json["revisionType"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RevisionType = &parsedValue
        }

    }
}

func (obj *Revision) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *Revision) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.RevisionDateTime == nil {
        return errors.New("Property RevisionDateTime in Revision is required.")
    }
    return nil;
}

func (obj *Revision) GetRevisionAuthor() *string {
    return obj.RevisionAuthor
}

func (obj *Revision) SetRevisionAuthor(value *string) {
    obj.RevisionAuthor = value
}

func (obj *Revision) GetRevisionDateTime() *Time {
    return obj.RevisionDateTime
}

func (obj *Revision) SetRevisionDateTime(value *Time) {
    obj.RevisionDateTime = value
}

func (obj *Revision) GetRevisionText() *string {
    return obj.RevisionText
}

func (obj *Revision) SetRevisionText(value *string) {
    obj.RevisionText = value
}

func (obj *Revision) GetRevisionType() *string {
    return obj.RevisionType
}

func (obj *Revision) SetRevisionType(value *string) {
    obj.RevisionType = value
}

