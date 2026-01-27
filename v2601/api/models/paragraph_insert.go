/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="paragraph_insert.go">
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

// DTO container with a paragraph's text.

type IParagraphInsert interface {
    IsParagraphInsert() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetText() *string
    SetText(value *string)
    GetPosition() IPosition
    SetPosition(value IPosition)
}

type ParagraphInsert struct {
    // DTO container with a paragraph's text.
    Text *string `json:"Text,omitempty"`

    // DTO container with a paragraph's text.
    Position IPosition `json:"Position,omitempty"`
}

func (ParagraphInsert) IsParagraphInsert() bool {
    return true
}


func (obj *ParagraphInsert) Initialize() {
    if (obj.Position != nil) {
        obj.Position.Initialize()
    }


}

func (obj *ParagraphInsert) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Text"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Text = &parsedValue
        }

    } else if jsonValue, exists := json["text"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Text = &parsedValue
        }

    }

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
}

func (obj *ParagraphInsert) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *ParagraphInsert) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Text == nil {
        return errors.New("Property Text in ParagraphInsert is required.")
    }
    if obj.Position != nil {
        if err := obj.Position.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *ParagraphInsert) GetText() *string {
    return obj.Text
}

func (obj *ParagraphInsert) SetText(value *string) {
    obj.Text = value
}

func (obj *ParagraphInsert) GetPosition() IPosition {
    return obj.Position
}

func (obj *ParagraphInsert) SetPosition(value IPosition) {
    obj.Position = value
}

