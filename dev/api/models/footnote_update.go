/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="footnote_update.go">
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

// Footnote for update.

type IFootnoteUpdate interface {
    IsFootnoteUpdate() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetFootnoteType() *string
    SetFootnoteType(value *string)
    GetPosition() INewDocumentPosition
    SetPosition(value INewDocumentPosition)
    GetReferenceMark() *string
    SetReferenceMark(value *string)
    GetText() *string
    SetText(value *string)
}

type FootnoteUpdate struct {
    // Footnote for update.
    FootnoteType *string `json:"FootnoteType,omitempty"`

    // Footnote for update.
    Position INewDocumentPosition `json:"Position,omitempty"`

    // Footnote for update.
    ReferenceMark *string `json:"ReferenceMark,omitempty"`

    // Footnote for update.
    Text *string `json:"Text,omitempty"`
}

func (FootnoteUpdate) IsFootnoteUpdate() bool {
    return true
}

func (FootnoteUpdate) IsFootnoteBase() bool {
    return true
}

func (obj *FootnoteUpdate) Initialize() {
    if (obj.Position != nil) {
        obj.Position.Initialize()
    }


}

func (obj *FootnoteUpdate) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["FootnoteType"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FootnoteType = &parsedValue
        }

    } else if jsonValue, exists := json["footnoteType"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FootnoteType = &parsedValue
        }

    }

    if jsonValue, exists := json["Position"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance INewDocumentPosition = new(NewDocumentPosition)
            modelInstance.Deserialize(parsedValue)
            obj.Position = modelInstance
        }

    } else if jsonValue, exists := json["position"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance INewDocumentPosition = new(NewDocumentPosition)
            modelInstance.Deserialize(parsedValue)
            obj.Position = modelInstance
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

func (obj *FootnoteUpdate) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *FootnoteUpdate) GetFootnoteType() *string {
    return obj.FootnoteType
}

func (obj *FootnoteUpdate) SetFootnoteType(value *string) {
    obj.FootnoteType = value
}

func (obj *FootnoteUpdate) GetPosition() INewDocumentPosition {
    return obj.Position
}

func (obj *FootnoteUpdate) SetPosition(value INewDocumentPosition) {
    obj.Position = value
}

func (obj *FootnoteUpdate) GetReferenceMark() *string {
    return obj.ReferenceMark
}

func (obj *FootnoteUpdate) SetReferenceMark(value *string) {
    obj.ReferenceMark = value
}

func (obj *FootnoteUpdate) GetText() *string {
    return obj.Text
}

func (obj *FootnoteUpdate) SetText(value *string) {
    obj.Text = value
}

