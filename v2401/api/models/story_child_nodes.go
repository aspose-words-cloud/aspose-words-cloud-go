/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="story_child_nodes.go">
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

// Child nodes of Story or InlineStory.

type IStoryChildNodes interface {
    IsStoryChildNodes() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetChildNodes() []INodeLink
    SetChildNodes(value []INodeLink)
}

type StoryChildNodes struct {
    // Child nodes of Story or InlineStory.
    ChildNodes []INodeLink `json:"ChildNodes,omitempty"`
}

func (StoryChildNodes) IsStoryChildNodes() bool {
    return true
}


func (obj *StoryChildNodes) Initialize() {
    if (obj.ChildNodes != nil) {
        for _, objElementChildNodes := range obj.ChildNodes {
            objElementChildNodes.Initialize()
        }
    }

}

func (obj *StoryChildNodes) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["ChildNodes"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.ChildNodes = make([]INodeLink, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance INodeLink = nil
                    if jsonElementType, found := elementValue["$type"]; found {
                        jsonTypeStr := jsonElementType.(string)
                        if jsonTypeStr == "DrawingObject, _" { modelElementInstance = new(DrawingObject) }
                        if jsonTypeStr == "DrawingObjectLink, _" { modelElementInstance = new(DrawingObjectLink) }
                        if jsonTypeStr == "Field, _" { modelElementInstance = new(Field) }
                        if jsonTypeStr == "FieldLink, _" { modelElementInstance = new(FieldLink) }
                        if jsonTypeStr == "Footnote, _" { modelElementInstance = new(Footnote) }
                        if jsonTypeStr == "FootnoteLink, _" { modelElementInstance = new(FootnoteLink) }
                        if jsonTypeStr == "FormField, _" {  }
                        if jsonTypeStr == "FormFieldCheckbox, _" { modelElementInstance = new(FormFieldCheckbox) }
                        if jsonTypeStr == "FormFieldDropDown, _" { modelElementInstance = new(FormFieldDropDown) }
                        if jsonTypeStr == "FormFieldTextInput, _" { modelElementInstance = new(FormFieldTextInput) }
                        if jsonTypeStr == "OfficeMathLink, _" { modelElementInstance = new(OfficeMathLink) }
                        if jsonTypeStr == "OfficeMathObject, _" { modelElementInstance = new(OfficeMathObject) }
                        if jsonTypeStr == "Paragraph, _" { modelElementInstance = new(Paragraph) }
                        if jsonTypeStr == "ParagraphLink, _" { modelElementInstance = new(ParagraphLink) }
                        if jsonTypeStr == "Run, _" { modelElementInstance = new(Run) }
                        if jsonTypeStr == "RunLink, _" { modelElementInstance = new(RunLink) }
                        if jsonTypeStr == "SectionLink, _" { modelElementInstance = new(SectionLink) }
                        if jsonTypeStr == "StructuredDocumentTag, _" { modelElementInstance = new(StructuredDocumentTag) }
                        if jsonTypeStr == "StructuredDocumentTagBase, _" {  }
                        if jsonTypeStr == "StructuredDocumentTagInsert, _" { modelElementInstance = new(StructuredDocumentTagInsert) }
                        if jsonTypeStr == "StructuredDocumentTagUpdate, _" { modelElementInstance = new(StructuredDocumentTagUpdate) }
                        if jsonTypeStr == "Table, _" { modelElementInstance = new(Table) }
                        if jsonTypeStr == "TableCell, _" { modelElementInstance = new(TableCell) }
                        if jsonTypeStr == "TableLink, _" { modelElementInstance = new(TableLink) }
                        if jsonTypeStr == "TableRow, _" { modelElementInstance = new(TableRow) }
                    }

                    if modelElementInstance == nil { modelElementInstance = new(NodeLink) }
                    modelElementInstance.Deserialize(elementValue)
                    obj.ChildNodes = append(obj.ChildNodes, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["childNodes"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.ChildNodes = make([]INodeLink, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance INodeLink = nil
                    if jsonElementType, found := elementValue["$type"]; found {
                        jsonTypeStr := jsonElementType.(string)
                        if jsonTypeStr == "DrawingObject, _" { modelElementInstance = new(DrawingObject) }
                        if jsonTypeStr == "DrawingObjectLink, _" { modelElementInstance = new(DrawingObjectLink) }
                        if jsonTypeStr == "Field, _" { modelElementInstance = new(Field) }
                        if jsonTypeStr == "FieldLink, _" { modelElementInstance = new(FieldLink) }
                        if jsonTypeStr == "Footnote, _" { modelElementInstance = new(Footnote) }
                        if jsonTypeStr == "FootnoteLink, _" { modelElementInstance = new(FootnoteLink) }
                        if jsonTypeStr == "FormField, _" {  }
                        if jsonTypeStr == "FormFieldCheckbox, _" { modelElementInstance = new(FormFieldCheckbox) }
                        if jsonTypeStr == "FormFieldDropDown, _" { modelElementInstance = new(FormFieldDropDown) }
                        if jsonTypeStr == "FormFieldTextInput, _" { modelElementInstance = new(FormFieldTextInput) }
                        if jsonTypeStr == "OfficeMathLink, _" { modelElementInstance = new(OfficeMathLink) }
                        if jsonTypeStr == "OfficeMathObject, _" { modelElementInstance = new(OfficeMathObject) }
                        if jsonTypeStr == "Paragraph, _" { modelElementInstance = new(Paragraph) }
                        if jsonTypeStr == "ParagraphLink, _" { modelElementInstance = new(ParagraphLink) }
                        if jsonTypeStr == "Run, _" { modelElementInstance = new(Run) }
                        if jsonTypeStr == "RunLink, _" { modelElementInstance = new(RunLink) }
                        if jsonTypeStr == "SectionLink, _" { modelElementInstance = new(SectionLink) }
                        if jsonTypeStr == "StructuredDocumentTag, _" { modelElementInstance = new(StructuredDocumentTag) }
                        if jsonTypeStr == "StructuredDocumentTagBase, _" {  }
                        if jsonTypeStr == "StructuredDocumentTagInsert, _" { modelElementInstance = new(StructuredDocumentTagInsert) }
                        if jsonTypeStr == "StructuredDocumentTagUpdate, _" { modelElementInstance = new(StructuredDocumentTagUpdate) }
                        if jsonTypeStr == "Table, _" { modelElementInstance = new(Table) }
                        if jsonTypeStr == "TableCell, _" { modelElementInstance = new(TableCell) }
                        if jsonTypeStr == "TableLink, _" { modelElementInstance = new(TableLink) }
                        if jsonTypeStr == "TableRow, _" { modelElementInstance = new(TableRow) }
                    }

                    if modelElementInstance == nil { modelElementInstance = new(NodeLink) }
                    modelElementInstance.Deserialize(elementValue)
                    obj.ChildNodes = append(obj.ChildNodes, modelElementInstance)
                }

            }
        }

    }
}

func (obj *StoryChildNodes) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *StoryChildNodes) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.ChildNodes != nil {
        for _, elementChildNodes := range obj.ChildNodes {
            if elementChildNodes != nil {
                if err := elementChildNodes.Validate(); err != nil {
                    return err
                }
            }
        }
    }

    return nil;
}

func (obj *StoryChildNodes) GetChildNodes() []INodeLink {
    return obj.ChildNodes
}

func (obj *StoryChildNodes) SetChildNodes(value []INodeLink) {
    obj.ChildNodes = value
}

