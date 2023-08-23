/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="paragraph.go">
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

// DTO container with a paragraph element.

type IParagraph interface {
    IsParagraph() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetNodeId() *string
    SetNodeId(value *string)
    GetChildNodes() []INodeLink
    SetChildNodes(value []INodeLink)
}

type Paragraph struct {
    // DTO container with a paragraph element.
    Link IWordsApiLink

    // DTO container with a paragraph element.
    NodeId *string

    // DTO container with a paragraph element.
    ChildNodes []INodeLink
}

func (Paragraph) IsParagraph() bool {
    return true
}

func (Paragraph) IsNodeLink() bool {
    return true
}

func (Paragraph) IsLinkElement() bool {
    return true
}

func (obj *Paragraph) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.ChildNodes != nil) {
        for _, objElementChildNodes := range obj.ChildNodes {
            objElementChildNodes.Initialize()
        }
    }

}

func (obj *Paragraph) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["ChildNodes"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
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

func (obj *Paragraph) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *Paragraph) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *Paragraph) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *Paragraph) GetNodeId() *string {
    return obj.NodeId
}

func (obj *Paragraph) SetNodeId(value *string) {
    obj.NodeId = value
}

func (obj *Paragraph) GetChildNodes() []INodeLink {
    return obj.ChildNodes
}

func (obj *Paragraph) SetChildNodes(value []INodeLink) {
    obj.ChildNodes = value
}

