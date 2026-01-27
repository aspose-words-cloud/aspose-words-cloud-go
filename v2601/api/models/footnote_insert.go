/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="footnote_insert.go">
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

// Footnote for insert.

type IFootnoteInsert interface {
    IsFootnoteInsert() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetPosition() IPosition
    SetPosition(value IPosition)
    GetFootnoteType() *string
    SetFootnoteType(value *string)
    GetReferenceMark() *string
    SetReferenceMark(value *string)
    GetText() *string
    SetText(value *string)
}

type FootnoteInsert struct {
    // Footnote for insert.
    Position IPosition `json:"Position,omitempty"`

    // Footnote for insert.
    FootnoteType *string `json:"FootnoteType,omitempty"`

    // Footnote for insert.
    ReferenceMark *string `json:"ReferenceMark,omitempty"`

    // Footnote for insert.
    Text *string `json:"Text,omitempty"`
}

func (FootnoteInsert) IsFootnoteInsert() bool {
    return true
}

func (FootnoteInsert) IsFootnoteBase() bool {
    return true
}

func (obj *FootnoteInsert) Initialize() {
    if (obj.Position != nil) {
        obj.Position.Initialize()
    }


}

func (obj *FootnoteInsert) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Position"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IPosition = nil
            if jsonType, found := parsedValue["$type"]; found {
                jsonTypeStr := jsonType.(string)
                if jsonTypeStr == "PositionAfterNode, _" { modelInstance = new(PositionAfterNode) }
                if jsonTypeStr == "PositionBeforeNode, _" { modelInstance = new(PositionBeforeNode) }
                if jsonTypeStr == "PositionInsideNode, _" { modelInstance = new(PositionInsideNode) }
            }

            modelInstance.Deserialize(parsedValue)
            obj.Position = modelInstance
        }

    } else if jsonValue, exists := json["position"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IPosition = nil
            if jsonType, found := parsedValue["$type"]; found {
                jsonTypeStr := jsonType.(string)
                if jsonTypeStr == "PositionAfterNode, _" { modelInstance = new(PositionAfterNode) }
                if jsonTypeStr == "PositionBeforeNode, _" { modelInstance = new(PositionBeforeNode) }
                if jsonTypeStr == "PositionInsideNode, _" { modelInstance = new(PositionInsideNode) }
            }

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
}

func (obj *FootnoteInsert) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *FootnoteInsert) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Position != nil {
        if err := obj.Position.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *FootnoteInsert) GetPosition() IPosition {
    return obj.Position
}

func (obj *FootnoteInsert) SetPosition(value IPosition) {
    obj.Position = value
}

func (obj *FootnoteInsert) GetFootnoteType() *string {
    return obj.FootnoteType
}

func (obj *FootnoteInsert) SetFootnoteType(value *string) {
    obj.FootnoteType = value
}

func (obj *FootnoteInsert) GetReferenceMark() *string {
    return obj.ReferenceMark
}

func (obj *FootnoteInsert) SetReferenceMark(value *string) {
    obj.ReferenceMark = value
}

func (obj *FootnoteInsert) GetText() *string {
    return obj.Text
}

func (obj *FootnoteInsert) SetText(value *string) {
    obj.Text = value
}

