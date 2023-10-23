/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="document_position.go">
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

import (
    "errors"
)

// DTO container with a position in the document tree.

type IDocumentPosition interface {
    IsDocumentPosition() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetNode() INodeLink
    SetNode(value INodeLink)
    GetOffset() *int32
    SetOffset(value *int32)
}

type DocumentPosition struct {
    // DTO container with a position in the document tree.
    Node INodeLink `json:"Node,omitempty"`

    // DTO container with a position in the document tree.
    Offset *int32 `json:"Offset,omitempty"`
}

func (DocumentPosition) IsDocumentPosition() bool {
    return true
}


func (obj *DocumentPosition) Initialize() {
    if (obj.Node != nil) {
        obj.Node.Initialize()
    }


}

func (obj *DocumentPosition) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Node"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance INodeLink = nil
            if jsonType, found := parsedValue["$type"]; found {
                jsonTypeStr := jsonType.(string)
                if jsonTypeStr == "DrawingObject, _" { modelInstance = new(DrawingObject) }
                if jsonTypeStr == "DrawingObjectLink, _" { modelInstance = new(DrawingObjectLink) }
                if jsonTypeStr == "Field, _" { modelInstance = new(Field) }
                if jsonTypeStr == "FieldLink, _" { modelInstance = new(FieldLink) }
                if jsonTypeStr == "Footnote, _" { modelInstance = new(Footnote) }
                if jsonTypeStr == "FootnoteLink, _" { modelInstance = new(FootnoteLink) }
                if jsonTypeStr == "FormField, _" {  }
                if jsonTypeStr == "FormFieldCheckbox, _" { modelInstance = new(FormFieldCheckbox) }
                if jsonTypeStr == "FormFieldDropDown, _" { modelInstance = new(FormFieldDropDown) }
                if jsonTypeStr == "FormFieldTextInput, _" { modelInstance = new(FormFieldTextInput) }
                if jsonTypeStr == "OfficeMathLink, _" { modelInstance = new(OfficeMathLink) }
                if jsonTypeStr == "OfficeMathObject, _" { modelInstance = new(OfficeMathObject) }
                if jsonTypeStr == "Paragraph, _" { modelInstance = new(Paragraph) }
                if jsonTypeStr == "ParagraphLink, _" { modelInstance = new(ParagraphLink) }
                if jsonTypeStr == "Run, _" { modelInstance = new(Run) }
                if jsonTypeStr == "RunLink, _" { modelInstance = new(RunLink) }
                if jsonTypeStr == "SectionLink, _" { modelInstance = new(SectionLink) }
                if jsonTypeStr == "StructuredDocumentTag, _" { modelInstance = new(StructuredDocumentTag) }
                if jsonTypeStr == "StructuredDocumentTagInsert, _" { modelInstance = new(StructuredDocumentTagInsert) }
                if jsonTypeStr == "StructuredDocumentTagUpdate, _" { modelInstance = new(StructuredDocumentTagUpdate) }
                if jsonTypeStr == "Table, _" { modelInstance = new(Table) }
                if jsonTypeStr == "TableCell, _" { modelInstance = new(TableCell) }
                if jsonTypeStr == "TableLink, _" { modelInstance = new(TableLink) }
                if jsonTypeStr == "TableRow, _" { modelInstance = new(TableRow) }
            }

            if modelInstance == nil { modelInstance = new(NodeLink) }
            modelInstance.Deserialize(parsedValue)
            obj.Node = modelInstance
        }

    } else if jsonValue, exists := json["node"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance INodeLink = nil
            if jsonType, found := parsedValue["$type"]; found {
                jsonTypeStr := jsonType.(string)
                if jsonTypeStr == "DrawingObject, _" { modelInstance = new(DrawingObject) }
                if jsonTypeStr == "DrawingObjectLink, _" { modelInstance = new(DrawingObjectLink) }
                if jsonTypeStr == "Field, _" { modelInstance = new(Field) }
                if jsonTypeStr == "FieldLink, _" { modelInstance = new(FieldLink) }
                if jsonTypeStr == "Footnote, _" { modelInstance = new(Footnote) }
                if jsonTypeStr == "FootnoteLink, _" { modelInstance = new(FootnoteLink) }
                if jsonTypeStr == "FormField, _" {  }
                if jsonTypeStr == "FormFieldCheckbox, _" { modelInstance = new(FormFieldCheckbox) }
                if jsonTypeStr == "FormFieldDropDown, _" { modelInstance = new(FormFieldDropDown) }
                if jsonTypeStr == "FormFieldTextInput, _" { modelInstance = new(FormFieldTextInput) }
                if jsonTypeStr == "OfficeMathLink, _" { modelInstance = new(OfficeMathLink) }
                if jsonTypeStr == "OfficeMathObject, _" { modelInstance = new(OfficeMathObject) }
                if jsonTypeStr == "Paragraph, _" { modelInstance = new(Paragraph) }
                if jsonTypeStr == "ParagraphLink, _" { modelInstance = new(ParagraphLink) }
                if jsonTypeStr == "Run, _" { modelInstance = new(Run) }
                if jsonTypeStr == "RunLink, _" { modelInstance = new(RunLink) }
                if jsonTypeStr == "SectionLink, _" { modelInstance = new(SectionLink) }
                if jsonTypeStr == "StructuredDocumentTag, _" { modelInstance = new(StructuredDocumentTag) }
                if jsonTypeStr == "StructuredDocumentTagInsert, _" { modelInstance = new(StructuredDocumentTagInsert) }
                if jsonTypeStr == "StructuredDocumentTagUpdate, _" { modelInstance = new(StructuredDocumentTagUpdate) }
                if jsonTypeStr == "Table, _" { modelInstance = new(Table) }
                if jsonTypeStr == "TableCell, _" { modelInstance = new(TableCell) }
                if jsonTypeStr == "TableLink, _" { modelInstance = new(TableLink) }
                if jsonTypeStr == "TableRow, _" { modelInstance = new(TableRow) }
            }

            if modelInstance == nil { modelInstance = new(NodeLink) }
            modelInstance.Deserialize(parsedValue)
            obj.Node = modelInstance
        }

    }

    if jsonValue, exists := json["Offset"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Offset = new(int32)
            *obj.Offset = int32(parsedValue)
        }

    } else if jsonValue, exists := json["offset"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Offset = new(int32)
            *obj.Offset = int32(parsedValue)
        }

    }
}

func (obj *DocumentPosition) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *DocumentPosition) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Node == nil {
        return errors.New("Property Node in DocumentPosition is required.")
    }

    if obj.Node != nil {
        if err := obj.Node.Validate(); err != nil {
            return err
        }
    }

    if obj.Offset == nil {
        return errors.New("Property Offset in DocumentPosition is required.")
    }

    return nil;
}

func (obj *DocumentPosition) GetNode() INodeLink {
    return obj.Node
}

func (obj *DocumentPosition) SetNode(value INodeLink) {
    obj.Node = value
}

func (obj *DocumentPosition) GetOffset() *int32 {
    return obj.Offset
}

func (obj *DocumentPosition) SetOffset(value *int32) {
    obj.Offset = value
}

