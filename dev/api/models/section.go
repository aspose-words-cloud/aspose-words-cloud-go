/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="section.go">
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

// DTO container with a section element.

type ISection interface {
    IsSection() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetChildNodes() []INodeLink
    SetChildNodes(value []INodeLink)
    GetParagraphs() ILinkElement
    SetParagraphs(value ILinkElement)
    GetPageSetup() ILinkElement
    SetPageSetup(value ILinkElement)
    GetHeaderFooters() ILinkElement
    SetHeaderFooters(value ILinkElement)
    GetTables() ILinkElement
    SetTables(value ILinkElement)
}

type Section struct {
    // DTO container with a section element.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with a section element.
    ChildNodes []INodeLink `json:"ChildNodes,omitempty"`

    // DTO container with a section element.
    Paragraphs ILinkElement `json:"Paragraphs,omitempty"`

    // DTO container with a section element.
    PageSetup ILinkElement `json:"PageSetup,omitempty"`

    // DTO container with a section element.
    HeaderFooters ILinkElement `json:"HeaderFooters,omitempty"`

    // DTO container with a section element.
    Tables ILinkElement `json:"Tables,omitempty"`
}

func (Section) IsSection() bool {
    return true
}

func (Section) IsLinkElement() bool {
    return true
}

func (obj *Section) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.ChildNodes != nil) {
        for _, objElementChildNodes := range obj.ChildNodes {
            objElementChildNodes.Initialize()
        }
    }
    if (obj.Paragraphs != nil) {
        obj.Paragraphs.Initialize()
    }

    if (obj.PageSetup != nil) {
        obj.PageSetup.Initialize()
    }

    if (obj.HeaderFooters != nil) {
        obj.HeaderFooters.Initialize()
    }

    if (obj.Tables != nil) {
        obj.Tables.Initialize()
    }


}

func (obj *Section) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["Paragraphs"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ILinkElement = nil
            if jsonType, found := parsedValue["$type"]; found {
                jsonTypeStr := jsonType.(string)
                if jsonTypeStr == "Bookmark, _" { modelInstance = new(Bookmark) }
                if jsonTypeStr == "Bookmarks, _" { modelInstance = new(Bookmarks) }
                if jsonTypeStr == "Border, _" { modelInstance = new(Border) }
                if jsonTypeStr == "BordersCollection, _" { modelInstance = new(BordersCollection) }
                if jsonTypeStr == "Comment, _" { modelInstance = new(Comment) }
                if jsonTypeStr == "CommentLink, _" { modelInstance = new(CommentLink) }
                if jsonTypeStr == "CommentsCollection, _" { modelInstance = new(CommentsCollection) }
                if jsonTypeStr == "CustomXmlPart, _" { modelInstance = new(CustomXmlPart) }
                if jsonTypeStr == "CustomXmlPartInsert, _" { modelInstance = new(CustomXmlPartInsert) }
                if jsonTypeStr == "CustomXmlPartLink, _" { modelInstance = new(CustomXmlPartLink) }
                if jsonTypeStr == "CustomXmlPartsCollection, _" { modelInstance = new(CustomXmlPartsCollection) }
                if jsonTypeStr == "CustomXmlPartUpdate, _" { modelInstance = new(CustomXmlPartUpdate) }
                if jsonTypeStr == "DocumentProperties, _" { modelInstance = new(DocumentProperties) }
                if jsonTypeStr == "DocumentProperty, _" { modelInstance = new(DocumentProperty) }
                if jsonTypeStr == "DrawingObject, _" { modelInstance = new(DrawingObject) }
                if jsonTypeStr == "DrawingObjectCollection, _" { modelInstance = new(DrawingObjectCollection) }
                if jsonTypeStr == "DrawingObjectLink, _" { modelInstance = new(DrawingObjectLink) }
                if jsonTypeStr == "Field, _" { modelInstance = new(Field) }
                if jsonTypeStr == "FieldCollection, _" { modelInstance = new(FieldCollection) }
                if jsonTypeStr == "FieldLink, _" { modelInstance = new(FieldLink) }
                if jsonTypeStr == "FieldNames, _" { modelInstance = new(FieldNames) }
                if jsonTypeStr == "Font, _" { modelInstance = new(Font) }
                if jsonTypeStr == "Footnote, _" { modelInstance = new(Footnote) }
                if jsonTypeStr == "FootnoteCollection, _" { modelInstance = new(FootnoteCollection) }
                if jsonTypeStr == "FootnoteLink, _" { modelInstance = new(FootnoteLink) }
                if jsonTypeStr == "FormField, _" {  }
                if jsonTypeStr == "FormFieldCheckbox, _" { modelInstance = new(FormFieldCheckbox) }
                if jsonTypeStr == "FormFieldCollection, _" { modelInstance = new(FormFieldCollection) }
                if jsonTypeStr == "FormFieldDropDown, _" { modelInstance = new(FormFieldDropDown) }
                if jsonTypeStr == "FormFieldTextInput, _" { modelInstance = new(FormFieldTextInput) }
                if jsonTypeStr == "HeaderFooter, _" { modelInstance = new(HeaderFooter) }
                if jsonTypeStr == "HeaderFooterLink, _" { modelInstance = new(HeaderFooterLink) }
                if jsonTypeStr == "HeaderFooterLinkCollection, _" { modelInstance = new(HeaderFooterLinkCollection) }
                if jsonTypeStr == "Hyperlink, _" { modelInstance = new(Hyperlink) }
                if jsonTypeStr == "Hyperlinks, _" { modelInstance = new(Hyperlinks) }
                if jsonTypeStr == "ListFormat, _" { modelInstance = new(ListFormat) }
                if jsonTypeStr == "ListInfo, _" { modelInstance = new(ListInfo) }
                if jsonTypeStr == "ListLevel, _" { modelInstance = new(ListLevel) }
                if jsonTypeStr == "ListLevels, _" { modelInstance = new(ListLevels) }
                if jsonTypeStr == "Lists, _" { modelInstance = new(Lists) }
                if jsonTypeStr == "NodeLink, _" { modelInstance = new(NodeLink) }
                if jsonTypeStr == "OfficeMathLink, _" { modelInstance = new(OfficeMathLink) }
                if jsonTypeStr == "OfficeMathObject, _" { modelInstance = new(OfficeMathObject) }
                if jsonTypeStr == "OfficeMathObjectsCollection, _" { modelInstance = new(OfficeMathObjectsCollection) }
                if jsonTypeStr == "PageSetup, _" { modelInstance = new(PageSetup) }
                if jsonTypeStr == "Paragraph, _" { modelInstance = new(Paragraph) }
                if jsonTypeStr == "ParagraphFormat, _" { modelInstance = new(ParagraphFormat) }
                if jsonTypeStr == "ParagraphFormatBase, _" {  }
                if jsonTypeStr == "ParagraphFormatUpdate, _" { modelInstance = new(ParagraphFormatUpdate) }
                if jsonTypeStr == "ParagraphLink, _" { modelInstance = new(ParagraphLink) }
                if jsonTypeStr == "ParagraphLinkCollection, _" { modelInstance = new(ParagraphLinkCollection) }
                if jsonTypeStr == "Run, _" { modelInstance = new(Run) }
                if jsonTypeStr == "RunLink, _" { modelInstance = new(RunLink) }
                if jsonTypeStr == "Runs, _" { modelInstance = new(Runs) }
                if jsonTypeStr == "SearchResultsCollection, _" { modelInstance = new(SearchResultsCollection) }
                if jsonTypeStr == "Section, _" { modelInstance = new(Section) }
                if jsonTypeStr == "SectionLink, _" { modelInstance = new(SectionLink) }
                if jsonTypeStr == "SectionLinkCollection, _" { modelInstance = new(SectionLinkCollection) }
                if jsonTypeStr == "StructuredDocumentTag, _" { modelInstance = new(StructuredDocumentTag) }
                if jsonTypeStr == "StructuredDocumentTagCollection, _" { modelInstance = new(StructuredDocumentTagCollection) }
                if jsonTypeStr == "StructuredDocumentTagInsert, _" { modelInstance = new(StructuredDocumentTagInsert) }
                if jsonTypeStr == "StructuredDocumentTagUpdate, _" { modelInstance = new(StructuredDocumentTagUpdate) }
                if jsonTypeStr == "Style, _" { modelInstance = new(Style) }
                if jsonTypeStr == "Table, _" { modelInstance = new(Table) }
                if jsonTypeStr == "TableCell, _" { modelInstance = new(TableCell) }
                if jsonTypeStr == "TableCellFormat, _" { modelInstance = new(TableCellFormat) }
                if jsonTypeStr == "TableLink, _" { modelInstance = new(TableLink) }
                if jsonTypeStr == "TableLinkCollection, _" { modelInstance = new(TableLinkCollection) }
                if jsonTypeStr == "TableProperties, _" { modelInstance = new(TableProperties) }
                if jsonTypeStr == "TableRow, _" { modelInstance = new(TableRow) }
                if jsonTypeStr == "TableRowFormat, _" { modelInstance = new(TableRowFormat) }
            }

            if modelInstance == nil { modelInstance = new(LinkElement) }
            modelInstance.Deserialize(parsedValue)
            obj.Paragraphs = modelInstance
        }

    } else if jsonValue, exists := json["paragraphs"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ILinkElement = nil
            if jsonType, found := parsedValue["$type"]; found {
                jsonTypeStr := jsonType.(string)
                if jsonTypeStr == "Bookmark, _" { modelInstance = new(Bookmark) }
                if jsonTypeStr == "Bookmarks, _" { modelInstance = new(Bookmarks) }
                if jsonTypeStr == "Border, _" { modelInstance = new(Border) }
                if jsonTypeStr == "BordersCollection, _" { modelInstance = new(BordersCollection) }
                if jsonTypeStr == "Comment, _" { modelInstance = new(Comment) }
                if jsonTypeStr == "CommentLink, _" { modelInstance = new(CommentLink) }
                if jsonTypeStr == "CommentsCollection, _" { modelInstance = new(CommentsCollection) }
                if jsonTypeStr == "CustomXmlPart, _" { modelInstance = new(CustomXmlPart) }
                if jsonTypeStr == "CustomXmlPartInsert, _" { modelInstance = new(CustomXmlPartInsert) }
                if jsonTypeStr == "CustomXmlPartLink, _" { modelInstance = new(CustomXmlPartLink) }
                if jsonTypeStr == "CustomXmlPartsCollection, _" { modelInstance = new(CustomXmlPartsCollection) }
                if jsonTypeStr == "CustomXmlPartUpdate, _" { modelInstance = new(CustomXmlPartUpdate) }
                if jsonTypeStr == "DocumentProperties, _" { modelInstance = new(DocumentProperties) }
                if jsonTypeStr == "DocumentProperty, _" { modelInstance = new(DocumentProperty) }
                if jsonTypeStr == "DrawingObject, _" { modelInstance = new(DrawingObject) }
                if jsonTypeStr == "DrawingObjectCollection, _" { modelInstance = new(DrawingObjectCollection) }
                if jsonTypeStr == "DrawingObjectLink, _" { modelInstance = new(DrawingObjectLink) }
                if jsonTypeStr == "Field, _" { modelInstance = new(Field) }
                if jsonTypeStr == "FieldCollection, _" { modelInstance = new(FieldCollection) }
                if jsonTypeStr == "FieldLink, _" { modelInstance = new(FieldLink) }
                if jsonTypeStr == "FieldNames, _" { modelInstance = new(FieldNames) }
                if jsonTypeStr == "Font, _" { modelInstance = new(Font) }
                if jsonTypeStr == "Footnote, _" { modelInstance = new(Footnote) }
                if jsonTypeStr == "FootnoteCollection, _" { modelInstance = new(FootnoteCollection) }
                if jsonTypeStr == "FootnoteLink, _" { modelInstance = new(FootnoteLink) }
                if jsonTypeStr == "FormField, _" {  }
                if jsonTypeStr == "FormFieldCheckbox, _" { modelInstance = new(FormFieldCheckbox) }
                if jsonTypeStr == "FormFieldCollection, _" { modelInstance = new(FormFieldCollection) }
                if jsonTypeStr == "FormFieldDropDown, _" { modelInstance = new(FormFieldDropDown) }
                if jsonTypeStr == "FormFieldTextInput, _" { modelInstance = new(FormFieldTextInput) }
                if jsonTypeStr == "HeaderFooter, _" { modelInstance = new(HeaderFooter) }
                if jsonTypeStr == "HeaderFooterLink, _" { modelInstance = new(HeaderFooterLink) }
                if jsonTypeStr == "HeaderFooterLinkCollection, _" { modelInstance = new(HeaderFooterLinkCollection) }
                if jsonTypeStr == "Hyperlink, _" { modelInstance = new(Hyperlink) }
                if jsonTypeStr == "Hyperlinks, _" { modelInstance = new(Hyperlinks) }
                if jsonTypeStr == "ListFormat, _" { modelInstance = new(ListFormat) }
                if jsonTypeStr == "ListInfo, _" { modelInstance = new(ListInfo) }
                if jsonTypeStr == "ListLevel, _" { modelInstance = new(ListLevel) }
                if jsonTypeStr == "ListLevels, _" { modelInstance = new(ListLevels) }
                if jsonTypeStr == "Lists, _" { modelInstance = new(Lists) }
                if jsonTypeStr == "NodeLink, _" { modelInstance = new(NodeLink) }
                if jsonTypeStr == "OfficeMathLink, _" { modelInstance = new(OfficeMathLink) }
                if jsonTypeStr == "OfficeMathObject, _" { modelInstance = new(OfficeMathObject) }
                if jsonTypeStr == "OfficeMathObjectsCollection, _" { modelInstance = new(OfficeMathObjectsCollection) }
                if jsonTypeStr == "PageSetup, _" { modelInstance = new(PageSetup) }
                if jsonTypeStr == "Paragraph, _" { modelInstance = new(Paragraph) }
                if jsonTypeStr == "ParagraphFormat, _" { modelInstance = new(ParagraphFormat) }
                if jsonTypeStr == "ParagraphFormatBase, _" {  }
                if jsonTypeStr == "ParagraphFormatUpdate, _" { modelInstance = new(ParagraphFormatUpdate) }
                if jsonTypeStr == "ParagraphLink, _" { modelInstance = new(ParagraphLink) }
                if jsonTypeStr == "ParagraphLinkCollection, _" { modelInstance = new(ParagraphLinkCollection) }
                if jsonTypeStr == "Run, _" { modelInstance = new(Run) }
                if jsonTypeStr == "RunLink, _" { modelInstance = new(RunLink) }
                if jsonTypeStr == "Runs, _" { modelInstance = new(Runs) }
                if jsonTypeStr == "SearchResultsCollection, _" { modelInstance = new(SearchResultsCollection) }
                if jsonTypeStr == "Section, _" { modelInstance = new(Section) }
                if jsonTypeStr == "SectionLink, _" { modelInstance = new(SectionLink) }
                if jsonTypeStr == "SectionLinkCollection, _" { modelInstance = new(SectionLinkCollection) }
                if jsonTypeStr == "StructuredDocumentTag, _" { modelInstance = new(StructuredDocumentTag) }
                if jsonTypeStr == "StructuredDocumentTagCollection, _" { modelInstance = new(StructuredDocumentTagCollection) }
                if jsonTypeStr == "StructuredDocumentTagInsert, _" { modelInstance = new(StructuredDocumentTagInsert) }
                if jsonTypeStr == "StructuredDocumentTagUpdate, _" { modelInstance = new(StructuredDocumentTagUpdate) }
                if jsonTypeStr == "Style, _" { modelInstance = new(Style) }
                if jsonTypeStr == "Table, _" { modelInstance = new(Table) }
                if jsonTypeStr == "TableCell, _" { modelInstance = new(TableCell) }
                if jsonTypeStr == "TableCellFormat, _" { modelInstance = new(TableCellFormat) }
                if jsonTypeStr == "TableLink, _" { modelInstance = new(TableLink) }
                if jsonTypeStr == "TableLinkCollection, _" { modelInstance = new(TableLinkCollection) }
                if jsonTypeStr == "TableProperties, _" { modelInstance = new(TableProperties) }
                if jsonTypeStr == "TableRow, _" { modelInstance = new(TableRow) }
                if jsonTypeStr == "TableRowFormat, _" { modelInstance = new(TableRowFormat) }
            }

            if modelInstance == nil { modelInstance = new(LinkElement) }
            modelInstance.Deserialize(parsedValue)
            obj.Paragraphs = modelInstance
        }

    }

    if jsonValue, exists := json["PageSetup"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ILinkElement = nil
            if jsonType, found := parsedValue["$type"]; found {
                jsonTypeStr := jsonType.(string)
                if jsonTypeStr == "Bookmark, _" { modelInstance = new(Bookmark) }
                if jsonTypeStr == "Bookmarks, _" { modelInstance = new(Bookmarks) }
                if jsonTypeStr == "Border, _" { modelInstance = new(Border) }
                if jsonTypeStr == "BordersCollection, _" { modelInstance = new(BordersCollection) }
                if jsonTypeStr == "Comment, _" { modelInstance = new(Comment) }
                if jsonTypeStr == "CommentLink, _" { modelInstance = new(CommentLink) }
                if jsonTypeStr == "CommentsCollection, _" { modelInstance = new(CommentsCollection) }
                if jsonTypeStr == "CustomXmlPart, _" { modelInstance = new(CustomXmlPart) }
                if jsonTypeStr == "CustomXmlPartInsert, _" { modelInstance = new(CustomXmlPartInsert) }
                if jsonTypeStr == "CustomXmlPartLink, _" { modelInstance = new(CustomXmlPartLink) }
                if jsonTypeStr == "CustomXmlPartsCollection, _" { modelInstance = new(CustomXmlPartsCollection) }
                if jsonTypeStr == "CustomXmlPartUpdate, _" { modelInstance = new(CustomXmlPartUpdate) }
                if jsonTypeStr == "DocumentProperties, _" { modelInstance = new(DocumentProperties) }
                if jsonTypeStr == "DocumentProperty, _" { modelInstance = new(DocumentProperty) }
                if jsonTypeStr == "DrawingObject, _" { modelInstance = new(DrawingObject) }
                if jsonTypeStr == "DrawingObjectCollection, _" { modelInstance = new(DrawingObjectCollection) }
                if jsonTypeStr == "DrawingObjectLink, _" { modelInstance = new(DrawingObjectLink) }
                if jsonTypeStr == "Field, _" { modelInstance = new(Field) }
                if jsonTypeStr == "FieldCollection, _" { modelInstance = new(FieldCollection) }
                if jsonTypeStr == "FieldLink, _" { modelInstance = new(FieldLink) }
                if jsonTypeStr == "FieldNames, _" { modelInstance = new(FieldNames) }
                if jsonTypeStr == "Font, _" { modelInstance = new(Font) }
                if jsonTypeStr == "Footnote, _" { modelInstance = new(Footnote) }
                if jsonTypeStr == "FootnoteCollection, _" { modelInstance = new(FootnoteCollection) }
                if jsonTypeStr == "FootnoteLink, _" { modelInstance = new(FootnoteLink) }
                if jsonTypeStr == "FormField, _" {  }
                if jsonTypeStr == "FormFieldCheckbox, _" { modelInstance = new(FormFieldCheckbox) }
                if jsonTypeStr == "FormFieldCollection, _" { modelInstance = new(FormFieldCollection) }
                if jsonTypeStr == "FormFieldDropDown, _" { modelInstance = new(FormFieldDropDown) }
                if jsonTypeStr == "FormFieldTextInput, _" { modelInstance = new(FormFieldTextInput) }
                if jsonTypeStr == "HeaderFooter, _" { modelInstance = new(HeaderFooter) }
                if jsonTypeStr == "HeaderFooterLink, _" { modelInstance = new(HeaderFooterLink) }
                if jsonTypeStr == "HeaderFooterLinkCollection, _" { modelInstance = new(HeaderFooterLinkCollection) }
                if jsonTypeStr == "Hyperlink, _" { modelInstance = new(Hyperlink) }
                if jsonTypeStr == "Hyperlinks, _" { modelInstance = new(Hyperlinks) }
                if jsonTypeStr == "ListFormat, _" { modelInstance = new(ListFormat) }
                if jsonTypeStr == "ListInfo, _" { modelInstance = new(ListInfo) }
                if jsonTypeStr == "ListLevel, _" { modelInstance = new(ListLevel) }
                if jsonTypeStr == "ListLevels, _" { modelInstance = new(ListLevels) }
                if jsonTypeStr == "Lists, _" { modelInstance = new(Lists) }
                if jsonTypeStr == "NodeLink, _" { modelInstance = new(NodeLink) }
                if jsonTypeStr == "OfficeMathLink, _" { modelInstance = new(OfficeMathLink) }
                if jsonTypeStr == "OfficeMathObject, _" { modelInstance = new(OfficeMathObject) }
                if jsonTypeStr == "OfficeMathObjectsCollection, _" { modelInstance = new(OfficeMathObjectsCollection) }
                if jsonTypeStr == "PageSetup, _" { modelInstance = new(PageSetup) }
                if jsonTypeStr == "Paragraph, _" { modelInstance = new(Paragraph) }
                if jsonTypeStr == "ParagraphFormat, _" { modelInstance = new(ParagraphFormat) }
                if jsonTypeStr == "ParagraphFormatBase, _" {  }
                if jsonTypeStr == "ParagraphFormatUpdate, _" { modelInstance = new(ParagraphFormatUpdate) }
                if jsonTypeStr == "ParagraphLink, _" { modelInstance = new(ParagraphLink) }
                if jsonTypeStr == "ParagraphLinkCollection, _" { modelInstance = new(ParagraphLinkCollection) }
                if jsonTypeStr == "Run, _" { modelInstance = new(Run) }
                if jsonTypeStr == "RunLink, _" { modelInstance = new(RunLink) }
                if jsonTypeStr == "Runs, _" { modelInstance = new(Runs) }
                if jsonTypeStr == "SearchResultsCollection, _" { modelInstance = new(SearchResultsCollection) }
                if jsonTypeStr == "Section, _" { modelInstance = new(Section) }
                if jsonTypeStr == "SectionLink, _" { modelInstance = new(SectionLink) }
                if jsonTypeStr == "SectionLinkCollection, _" { modelInstance = new(SectionLinkCollection) }
                if jsonTypeStr == "StructuredDocumentTag, _" { modelInstance = new(StructuredDocumentTag) }
                if jsonTypeStr == "StructuredDocumentTagCollection, _" { modelInstance = new(StructuredDocumentTagCollection) }
                if jsonTypeStr == "StructuredDocumentTagInsert, _" { modelInstance = new(StructuredDocumentTagInsert) }
                if jsonTypeStr == "StructuredDocumentTagUpdate, _" { modelInstance = new(StructuredDocumentTagUpdate) }
                if jsonTypeStr == "Style, _" { modelInstance = new(Style) }
                if jsonTypeStr == "Table, _" { modelInstance = new(Table) }
                if jsonTypeStr == "TableCell, _" { modelInstance = new(TableCell) }
                if jsonTypeStr == "TableCellFormat, _" { modelInstance = new(TableCellFormat) }
                if jsonTypeStr == "TableLink, _" { modelInstance = new(TableLink) }
                if jsonTypeStr == "TableLinkCollection, _" { modelInstance = new(TableLinkCollection) }
                if jsonTypeStr == "TableProperties, _" { modelInstance = new(TableProperties) }
                if jsonTypeStr == "TableRow, _" { modelInstance = new(TableRow) }
                if jsonTypeStr == "TableRowFormat, _" { modelInstance = new(TableRowFormat) }
            }

            if modelInstance == nil { modelInstance = new(LinkElement) }
            modelInstance.Deserialize(parsedValue)
            obj.PageSetup = modelInstance
        }

    } else if jsonValue, exists := json["pageSetup"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ILinkElement = nil
            if jsonType, found := parsedValue["$type"]; found {
                jsonTypeStr := jsonType.(string)
                if jsonTypeStr == "Bookmark, _" { modelInstance = new(Bookmark) }
                if jsonTypeStr == "Bookmarks, _" { modelInstance = new(Bookmarks) }
                if jsonTypeStr == "Border, _" { modelInstance = new(Border) }
                if jsonTypeStr == "BordersCollection, _" { modelInstance = new(BordersCollection) }
                if jsonTypeStr == "Comment, _" { modelInstance = new(Comment) }
                if jsonTypeStr == "CommentLink, _" { modelInstance = new(CommentLink) }
                if jsonTypeStr == "CommentsCollection, _" { modelInstance = new(CommentsCollection) }
                if jsonTypeStr == "CustomXmlPart, _" { modelInstance = new(CustomXmlPart) }
                if jsonTypeStr == "CustomXmlPartInsert, _" { modelInstance = new(CustomXmlPartInsert) }
                if jsonTypeStr == "CustomXmlPartLink, _" { modelInstance = new(CustomXmlPartLink) }
                if jsonTypeStr == "CustomXmlPartsCollection, _" { modelInstance = new(CustomXmlPartsCollection) }
                if jsonTypeStr == "CustomXmlPartUpdate, _" { modelInstance = new(CustomXmlPartUpdate) }
                if jsonTypeStr == "DocumentProperties, _" { modelInstance = new(DocumentProperties) }
                if jsonTypeStr == "DocumentProperty, _" { modelInstance = new(DocumentProperty) }
                if jsonTypeStr == "DrawingObject, _" { modelInstance = new(DrawingObject) }
                if jsonTypeStr == "DrawingObjectCollection, _" { modelInstance = new(DrawingObjectCollection) }
                if jsonTypeStr == "DrawingObjectLink, _" { modelInstance = new(DrawingObjectLink) }
                if jsonTypeStr == "Field, _" { modelInstance = new(Field) }
                if jsonTypeStr == "FieldCollection, _" { modelInstance = new(FieldCollection) }
                if jsonTypeStr == "FieldLink, _" { modelInstance = new(FieldLink) }
                if jsonTypeStr == "FieldNames, _" { modelInstance = new(FieldNames) }
                if jsonTypeStr == "Font, _" { modelInstance = new(Font) }
                if jsonTypeStr == "Footnote, _" { modelInstance = new(Footnote) }
                if jsonTypeStr == "FootnoteCollection, _" { modelInstance = new(FootnoteCollection) }
                if jsonTypeStr == "FootnoteLink, _" { modelInstance = new(FootnoteLink) }
                if jsonTypeStr == "FormField, _" {  }
                if jsonTypeStr == "FormFieldCheckbox, _" { modelInstance = new(FormFieldCheckbox) }
                if jsonTypeStr == "FormFieldCollection, _" { modelInstance = new(FormFieldCollection) }
                if jsonTypeStr == "FormFieldDropDown, _" { modelInstance = new(FormFieldDropDown) }
                if jsonTypeStr == "FormFieldTextInput, _" { modelInstance = new(FormFieldTextInput) }
                if jsonTypeStr == "HeaderFooter, _" { modelInstance = new(HeaderFooter) }
                if jsonTypeStr == "HeaderFooterLink, _" { modelInstance = new(HeaderFooterLink) }
                if jsonTypeStr == "HeaderFooterLinkCollection, _" { modelInstance = new(HeaderFooterLinkCollection) }
                if jsonTypeStr == "Hyperlink, _" { modelInstance = new(Hyperlink) }
                if jsonTypeStr == "Hyperlinks, _" { modelInstance = new(Hyperlinks) }
                if jsonTypeStr == "ListFormat, _" { modelInstance = new(ListFormat) }
                if jsonTypeStr == "ListInfo, _" { modelInstance = new(ListInfo) }
                if jsonTypeStr == "ListLevel, _" { modelInstance = new(ListLevel) }
                if jsonTypeStr == "ListLevels, _" { modelInstance = new(ListLevels) }
                if jsonTypeStr == "Lists, _" { modelInstance = new(Lists) }
                if jsonTypeStr == "NodeLink, _" { modelInstance = new(NodeLink) }
                if jsonTypeStr == "OfficeMathLink, _" { modelInstance = new(OfficeMathLink) }
                if jsonTypeStr == "OfficeMathObject, _" { modelInstance = new(OfficeMathObject) }
                if jsonTypeStr == "OfficeMathObjectsCollection, _" { modelInstance = new(OfficeMathObjectsCollection) }
                if jsonTypeStr == "PageSetup, _" { modelInstance = new(PageSetup) }
                if jsonTypeStr == "Paragraph, _" { modelInstance = new(Paragraph) }
                if jsonTypeStr == "ParagraphFormat, _" { modelInstance = new(ParagraphFormat) }
                if jsonTypeStr == "ParagraphFormatBase, _" {  }
                if jsonTypeStr == "ParagraphFormatUpdate, _" { modelInstance = new(ParagraphFormatUpdate) }
                if jsonTypeStr == "ParagraphLink, _" { modelInstance = new(ParagraphLink) }
                if jsonTypeStr == "ParagraphLinkCollection, _" { modelInstance = new(ParagraphLinkCollection) }
                if jsonTypeStr == "Run, _" { modelInstance = new(Run) }
                if jsonTypeStr == "RunLink, _" { modelInstance = new(RunLink) }
                if jsonTypeStr == "Runs, _" { modelInstance = new(Runs) }
                if jsonTypeStr == "SearchResultsCollection, _" { modelInstance = new(SearchResultsCollection) }
                if jsonTypeStr == "Section, _" { modelInstance = new(Section) }
                if jsonTypeStr == "SectionLink, _" { modelInstance = new(SectionLink) }
                if jsonTypeStr == "SectionLinkCollection, _" { modelInstance = new(SectionLinkCollection) }
                if jsonTypeStr == "StructuredDocumentTag, _" { modelInstance = new(StructuredDocumentTag) }
                if jsonTypeStr == "StructuredDocumentTagCollection, _" { modelInstance = new(StructuredDocumentTagCollection) }
                if jsonTypeStr == "StructuredDocumentTagInsert, _" { modelInstance = new(StructuredDocumentTagInsert) }
                if jsonTypeStr == "StructuredDocumentTagUpdate, _" { modelInstance = new(StructuredDocumentTagUpdate) }
                if jsonTypeStr == "Style, _" { modelInstance = new(Style) }
                if jsonTypeStr == "Table, _" { modelInstance = new(Table) }
                if jsonTypeStr == "TableCell, _" { modelInstance = new(TableCell) }
                if jsonTypeStr == "TableCellFormat, _" { modelInstance = new(TableCellFormat) }
                if jsonTypeStr == "TableLink, _" { modelInstance = new(TableLink) }
                if jsonTypeStr == "TableLinkCollection, _" { modelInstance = new(TableLinkCollection) }
                if jsonTypeStr == "TableProperties, _" { modelInstance = new(TableProperties) }
                if jsonTypeStr == "TableRow, _" { modelInstance = new(TableRow) }
                if jsonTypeStr == "TableRowFormat, _" { modelInstance = new(TableRowFormat) }
            }

            if modelInstance == nil { modelInstance = new(LinkElement) }
            modelInstance.Deserialize(parsedValue)
            obj.PageSetup = modelInstance
        }

    }

    if jsonValue, exists := json["HeaderFooters"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ILinkElement = nil
            if jsonType, found := parsedValue["$type"]; found {
                jsonTypeStr := jsonType.(string)
                if jsonTypeStr == "Bookmark, _" { modelInstance = new(Bookmark) }
                if jsonTypeStr == "Bookmarks, _" { modelInstance = new(Bookmarks) }
                if jsonTypeStr == "Border, _" { modelInstance = new(Border) }
                if jsonTypeStr == "BordersCollection, _" { modelInstance = new(BordersCollection) }
                if jsonTypeStr == "Comment, _" { modelInstance = new(Comment) }
                if jsonTypeStr == "CommentLink, _" { modelInstance = new(CommentLink) }
                if jsonTypeStr == "CommentsCollection, _" { modelInstance = new(CommentsCollection) }
                if jsonTypeStr == "CustomXmlPart, _" { modelInstance = new(CustomXmlPart) }
                if jsonTypeStr == "CustomXmlPartInsert, _" { modelInstance = new(CustomXmlPartInsert) }
                if jsonTypeStr == "CustomXmlPartLink, _" { modelInstance = new(CustomXmlPartLink) }
                if jsonTypeStr == "CustomXmlPartsCollection, _" { modelInstance = new(CustomXmlPartsCollection) }
                if jsonTypeStr == "CustomXmlPartUpdate, _" { modelInstance = new(CustomXmlPartUpdate) }
                if jsonTypeStr == "DocumentProperties, _" { modelInstance = new(DocumentProperties) }
                if jsonTypeStr == "DocumentProperty, _" { modelInstance = new(DocumentProperty) }
                if jsonTypeStr == "DrawingObject, _" { modelInstance = new(DrawingObject) }
                if jsonTypeStr == "DrawingObjectCollection, _" { modelInstance = new(DrawingObjectCollection) }
                if jsonTypeStr == "DrawingObjectLink, _" { modelInstance = new(DrawingObjectLink) }
                if jsonTypeStr == "Field, _" { modelInstance = new(Field) }
                if jsonTypeStr == "FieldCollection, _" { modelInstance = new(FieldCollection) }
                if jsonTypeStr == "FieldLink, _" { modelInstance = new(FieldLink) }
                if jsonTypeStr == "FieldNames, _" { modelInstance = new(FieldNames) }
                if jsonTypeStr == "Font, _" { modelInstance = new(Font) }
                if jsonTypeStr == "Footnote, _" { modelInstance = new(Footnote) }
                if jsonTypeStr == "FootnoteCollection, _" { modelInstance = new(FootnoteCollection) }
                if jsonTypeStr == "FootnoteLink, _" { modelInstance = new(FootnoteLink) }
                if jsonTypeStr == "FormField, _" {  }
                if jsonTypeStr == "FormFieldCheckbox, _" { modelInstance = new(FormFieldCheckbox) }
                if jsonTypeStr == "FormFieldCollection, _" { modelInstance = new(FormFieldCollection) }
                if jsonTypeStr == "FormFieldDropDown, _" { modelInstance = new(FormFieldDropDown) }
                if jsonTypeStr == "FormFieldTextInput, _" { modelInstance = new(FormFieldTextInput) }
                if jsonTypeStr == "HeaderFooter, _" { modelInstance = new(HeaderFooter) }
                if jsonTypeStr == "HeaderFooterLink, _" { modelInstance = new(HeaderFooterLink) }
                if jsonTypeStr == "HeaderFooterLinkCollection, _" { modelInstance = new(HeaderFooterLinkCollection) }
                if jsonTypeStr == "Hyperlink, _" { modelInstance = new(Hyperlink) }
                if jsonTypeStr == "Hyperlinks, _" { modelInstance = new(Hyperlinks) }
                if jsonTypeStr == "ListFormat, _" { modelInstance = new(ListFormat) }
                if jsonTypeStr == "ListInfo, _" { modelInstance = new(ListInfo) }
                if jsonTypeStr == "ListLevel, _" { modelInstance = new(ListLevel) }
                if jsonTypeStr == "ListLevels, _" { modelInstance = new(ListLevels) }
                if jsonTypeStr == "Lists, _" { modelInstance = new(Lists) }
                if jsonTypeStr == "NodeLink, _" { modelInstance = new(NodeLink) }
                if jsonTypeStr == "OfficeMathLink, _" { modelInstance = new(OfficeMathLink) }
                if jsonTypeStr == "OfficeMathObject, _" { modelInstance = new(OfficeMathObject) }
                if jsonTypeStr == "OfficeMathObjectsCollection, _" { modelInstance = new(OfficeMathObjectsCollection) }
                if jsonTypeStr == "PageSetup, _" { modelInstance = new(PageSetup) }
                if jsonTypeStr == "Paragraph, _" { modelInstance = new(Paragraph) }
                if jsonTypeStr == "ParagraphFormat, _" { modelInstance = new(ParagraphFormat) }
                if jsonTypeStr == "ParagraphFormatBase, _" {  }
                if jsonTypeStr == "ParagraphFormatUpdate, _" { modelInstance = new(ParagraphFormatUpdate) }
                if jsonTypeStr == "ParagraphLink, _" { modelInstance = new(ParagraphLink) }
                if jsonTypeStr == "ParagraphLinkCollection, _" { modelInstance = new(ParagraphLinkCollection) }
                if jsonTypeStr == "Run, _" { modelInstance = new(Run) }
                if jsonTypeStr == "RunLink, _" { modelInstance = new(RunLink) }
                if jsonTypeStr == "Runs, _" { modelInstance = new(Runs) }
                if jsonTypeStr == "SearchResultsCollection, _" { modelInstance = new(SearchResultsCollection) }
                if jsonTypeStr == "Section, _" { modelInstance = new(Section) }
                if jsonTypeStr == "SectionLink, _" { modelInstance = new(SectionLink) }
                if jsonTypeStr == "SectionLinkCollection, _" { modelInstance = new(SectionLinkCollection) }
                if jsonTypeStr == "StructuredDocumentTag, _" { modelInstance = new(StructuredDocumentTag) }
                if jsonTypeStr == "StructuredDocumentTagCollection, _" { modelInstance = new(StructuredDocumentTagCollection) }
                if jsonTypeStr == "StructuredDocumentTagInsert, _" { modelInstance = new(StructuredDocumentTagInsert) }
                if jsonTypeStr == "StructuredDocumentTagUpdate, _" { modelInstance = new(StructuredDocumentTagUpdate) }
                if jsonTypeStr == "Style, _" { modelInstance = new(Style) }
                if jsonTypeStr == "Table, _" { modelInstance = new(Table) }
                if jsonTypeStr == "TableCell, _" { modelInstance = new(TableCell) }
                if jsonTypeStr == "TableCellFormat, _" { modelInstance = new(TableCellFormat) }
                if jsonTypeStr == "TableLink, _" { modelInstance = new(TableLink) }
                if jsonTypeStr == "TableLinkCollection, _" { modelInstance = new(TableLinkCollection) }
                if jsonTypeStr == "TableProperties, _" { modelInstance = new(TableProperties) }
                if jsonTypeStr == "TableRow, _" { modelInstance = new(TableRow) }
                if jsonTypeStr == "TableRowFormat, _" { modelInstance = new(TableRowFormat) }
            }

            if modelInstance == nil { modelInstance = new(LinkElement) }
            modelInstance.Deserialize(parsedValue)
            obj.HeaderFooters = modelInstance
        }

    } else if jsonValue, exists := json["headerFooters"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ILinkElement = nil
            if jsonType, found := parsedValue["$type"]; found {
                jsonTypeStr := jsonType.(string)
                if jsonTypeStr == "Bookmark, _" { modelInstance = new(Bookmark) }
                if jsonTypeStr == "Bookmarks, _" { modelInstance = new(Bookmarks) }
                if jsonTypeStr == "Border, _" { modelInstance = new(Border) }
                if jsonTypeStr == "BordersCollection, _" { modelInstance = new(BordersCollection) }
                if jsonTypeStr == "Comment, _" { modelInstance = new(Comment) }
                if jsonTypeStr == "CommentLink, _" { modelInstance = new(CommentLink) }
                if jsonTypeStr == "CommentsCollection, _" { modelInstance = new(CommentsCollection) }
                if jsonTypeStr == "CustomXmlPart, _" { modelInstance = new(CustomXmlPart) }
                if jsonTypeStr == "CustomXmlPartInsert, _" { modelInstance = new(CustomXmlPartInsert) }
                if jsonTypeStr == "CustomXmlPartLink, _" { modelInstance = new(CustomXmlPartLink) }
                if jsonTypeStr == "CustomXmlPartsCollection, _" { modelInstance = new(CustomXmlPartsCollection) }
                if jsonTypeStr == "CustomXmlPartUpdate, _" { modelInstance = new(CustomXmlPartUpdate) }
                if jsonTypeStr == "DocumentProperties, _" { modelInstance = new(DocumentProperties) }
                if jsonTypeStr == "DocumentProperty, _" { modelInstance = new(DocumentProperty) }
                if jsonTypeStr == "DrawingObject, _" { modelInstance = new(DrawingObject) }
                if jsonTypeStr == "DrawingObjectCollection, _" { modelInstance = new(DrawingObjectCollection) }
                if jsonTypeStr == "DrawingObjectLink, _" { modelInstance = new(DrawingObjectLink) }
                if jsonTypeStr == "Field, _" { modelInstance = new(Field) }
                if jsonTypeStr == "FieldCollection, _" { modelInstance = new(FieldCollection) }
                if jsonTypeStr == "FieldLink, _" { modelInstance = new(FieldLink) }
                if jsonTypeStr == "FieldNames, _" { modelInstance = new(FieldNames) }
                if jsonTypeStr == "Font, _" { modelInstance = new(Font) }
                if jsonTypeStr == "Footnote, _" { modelInstance = new(Footnote) }
                if jsonTypeStr == "FootnoteCollection, _" { modelInstance = new(FootnoteCollection) }
                if jsonTypeStr == "FootnoteLink, _" { modelInstance = new(FootnoteLink) }
                if jsonTypeStr == "FormField, _" {  }
                if jsonTypeStr == "FormFieldCheckbox, _" { modelInstance = new(FormFieldCheckbox) }
                if jsonTypeStr == "FormFieldCollection, _" { modelInstance = new(FormFieldCollection) }
                if jsonTypeStr == "FormFieldDropDown, _" { modelInstance = new(FormFieldDropDown) }
                if jsonTypeStr == "FormFieldTextInput, _" { modelInstance = new(FormFieldTextInput) }
                if jsonTypeStr == "HeaderFooter, _" { modelInstance = new(HeaderFooter) }
                if jsonTypeStr == "HeaderFooterLink, _" { modelInstance = new(HeaderFooterLink) }
                if jsonTypeStr == "HeaderFooterLinkCollection, _" { modelInstance = new(HeaderFooterLinkCollection) }
                if jsonTypeStr == "Hyperlink, _" { modelInstance = new(Hyperlink) }
                if jsonTypeStr == "Hyperlinks, _" { modelInstance = new(Hyperlinks) }
                if jsonTypeStr == "ListFormat, _" { modelInstance = new(ListFormat) }
                if jsonTypeStr == "ListInfo, _" { modelInstance = new(ListInfo) }
                if jsonTypeStr == "ListLevel, _" { modelInstance = new(ListLevel) }
                if jsonTypeStr == "ListLevels, _" { modelInstance = new(ListLevels) }
                if jsonTypeStr == "Lists, _" { modelInstance = new(Lists) }
                if jsonTypeStr == "NodeLink, _" { modelInstance = new(NodeLink) }
                if jsonTypeStr == "OfficeMathLink, _" { modelInstance = new(OfficeMathLink) }
                if jsonTypeStr == "OfficeMathObject, _" { modelInstance = new(OfficeMathObject) }
                if jsonTypeStr == "OfficeMathObjectsCollection, _" { modelInstance = new(OfficeMathObjectsCollection) }
                if jsonTypeStr == "PageSetup, _" { modelInstance = new(PageSetup) }
                if jsonTypeStr == "Paragraph, _" { modelInstance = new(Paragraph) }
                if jsonTypeStr == "ParagraphFormat, _" { modelInstance = new(ParagraphFormat) }
                if jsonTypeStr == "ParagraphFormatBase, _" {  }
                if jsonTypeStr == "ParagraphFormatUpdate, _" { modelInstance = new(ParagraphFormatUpdate) }
                if jsonTypeStr == "ParagraphLink, _" { modelInstance = new(ParagraphLink) }
                if jsonTypeStr == "ParagraphLinkCollection, _" { modelInstance = new(ParagraphLinkCollection) }
                if jsonTypeStr == "Run, _" { modelInstance = new(Run) }
                if jsonTypeStr == "RunLink, _" { modelInstance = new(RunLink) }
                if jsonTypeStr == "Runs, _" { modelInstance = new(Runs) }
                if jsonTypeStr == "SearchResultsCollection, _" { modelInstance = new(SearchResultsCollection) }
                if jsonTypeStr == "Section, _" { modelInstance = new(Section) }
                if jsonTypeStr == "SectionLink, _" { modelInstance = new(SectionLink) }
                if jsonTypeStr == "SectionLinkCollection, _" { modelInstance = new(SectionLinkCollection) }
                if jsonTypeStr == "StructuredDocumentTag, _" { modelInstance = new(StructuredDocumentTag) }
                if jsonTypeStr == "StructuredDocumentTagCollection, _" { modelInstance = new(StructuredDocumentTagCollection) }
                if jsonTypeStr == "StructuredDocumentTagInsert, _" { modelInstance = new(StructuredDocumentTagInsert) }
                if jsonTypeStr == "StructuredDocumentTagUpdate, _" { modelInstance = new(StructuredDocumentTagUpdate) }
                if jsonTypeStr == "Style, _" { modelInstance = new(Style) }
                if jsonTypeStr == "Table, _" { modelInstance = new(Table) }
                if jsonTypeStr == "TableCell, _" { modelInstance = new(TableCell) }
                if jsonTypeStr == "TableCellFormat, _" { modelInstance = new(TableCellFormat) }
                if jsonTypeStr == "TableLink, _" { modelInstance = new(TableLink) }
                if jsonTypeStr == "TableLinkCollection, _" { modelInstance = new(TableLinkCollection) }
                if jsonTypeStr == "TableProperties, _" { modelInstance = new(TableProperties) }
                if jsonTypeStr == "TableRow, _" { modelInstance = new(TableRow) }
                if jsonTypeStr == "TableRowFormat, _" { modelInstance = new(TableRowFormat) }
            }

            if modelInstance == nil { modelInstance = new(LinkElement) }
            modelInstance.Deserialize(parsedValue)
            obj.HeaderFooters = modelInstance
        }

    }

    if jsonValue, exists := json["Tables"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ILinkElement = nil
            if jsonType, found := parsedValue["$type"]; found {
                jsonTypeStr := jsonType.(string)
                if jsonTypeStr == "Bookmark, _" { modelInstance = new(Bookmark) }
                if jsonTypeStr == "Bookmarks, _" { modelInstance = new(Bookmarks) }
                if jsonTypeStr == "Border, _" { modelInstance = new(Border) }
                if jsonTypeStr == "BordersCollection, _" { modelInstance = new(BordersCollection) }
                if jsonTypeStr == "Comment, _" { modelInstance = new(Comment) }
                if jsonTypeStr == "CommentLink, _" { modelInstance = new(CommentLink) }
                if jsonTypeStr == "CommentsCollection, _" { modelInstance = new(CommentsCollection) }
                if jsonTypeStr == "CustomXmlPart, _" { modelInstance = new(CustomXmlPart) }
                if jsonTypeStr == "CustomXmlPartInsert, _" { modelInstance = new(CustomXmlPartInsert) }
                if jsonTypeStr == "CustomXmlPartLink, _" { modelInstance = new(CustomXmlPartLink) }
                if jsonTypeStr == "CustomXmlPartsCollection, _" { modelInstance = new(CustomXmlPartsCollection) }
                if jsonTypeStr == "CustomXmlPartUpdate, _" { modelInstance = new(CustomXmlPartUpdate) }
                if jsonTypeStr == "DocumentProperties, _" { modelInstance = new(DocumentProperties) }
                if jsonTypeStr == "DocumentProperty, _" { modelInstance = new(DocumentProperty) }
                if jsonTypeStr == "DrawingObject, _" { modelInstance = new(DrawingObject) }
                if jsonTypeStr == "DrawingObjectCollection, _" { modelInstance = new(DrawingObjectCollection) }
                if jsonTypeStr == "DrawingObjectLink, _" { modelInstance = new(DrawingObjectLink) }
                if jsonTypeStr == "Field, _" { modelInstance = new(Field) }
                if jsonTypeStr == "FieldCollection, _" { modelInstance = new(FieldCollection) }
                if jsonTypeStr == "FieldLink, _" { modelInstance = new(FieldLink) }
                if jsonTypeStr == "FieldNames, _" { modelInstance = new(FieldNames) }
                if jsonTypeStr == "Font, _" { modelInstance = new(Font) }
                if jsonTypeStr == "Footnote, _" { modelInstance = new(Footnote) }
                if jsonTypeStr == "FootnoteCollection, _" { modelInstance = new(FootnoteCollection) }
                if jsonTypeStr == "FootnoteLink, _" { modelInstance = new(FootnoteLink) }
                if jsonTypeStr == "FormField, _" {  }
                if jsonTypeStr == "FormFieldCheckbox, _" { modelInstance = new(FormFieldCheckbox) }
                if jsonTypeStr == "FormFieldCollection, _" { modelInstance = new(FormFieldCollection) }
                if jsonTypeStr == "FormFieldDropDown, _" { modelInstance = new(FormFieldDropDown) }
                if jsonTypeStr == "FormFieldTextInput, _" { modelInstance = new(FormFieldTextInput) }
                if jsonTypeStr == "HeaderFooter, _" { modelInstance = new(HeaderFooter) }
                if jsonTypeStr == "HeaderFooterLink, _" { modelInstance = new(HeaderFooterLink) }
                if jsonTypeStr == "HeaderFooterLinkCollection, _" { modelInstance = new(HeaderFooterLinkCollection) }
                if jsonTypeStr == "Hyperlink, _" { modelInstance = new(Hyperlink) }
                if jsonTypeStr == "Hyperlinks, _" { modelInstance = new(Hyperlinks) }
                if jsonTypeStr == "ListFormat, _" { modelInstance = new(ListFormat) }
                if jsonTypeStr == "ListInfo, _" { modelInstance = new(ListInfo) }
                if jsonTypeStr == "ListLevel, _" { modelInstance = new(ListLevel) }
                if jsonTypeStr == "ListLevels, _" { modelInstance = new(ListLevels) }
                if jsonTypeStr == "Lists, _" { modelInstance = new(Lists) }
                if jsonTypeStr == "NodeLink, _" { modelInstance = new(NodeLink) }
                if jsonTypeStr == "OfficeMathLink, _" { modelInstance = new(OfficeMathLink) }
                if jsonTypeStr == "OfficeMathObject, _" { modelInstance = new(OfficeMathObject) }
                if jsonTypeStr == "OfficeMathObjectsCollection, _" { modelInstance = new(OfficeMathObjectsCollection) }
                if jsonTypeStr == "PageSetup, _" { modelInstance = new(PageSetup) }
                if jsonTypeStr == "Paragraph, _" { modelInstance = new(Paragraph) }
                if jsonTypeStr == "ParagraphFormat, _" { modelInstance = new(ParagraphFormat) }
                if jsonTypeStr == "ParagraphFormatBase, _" {  }
                if jsonTypeStr == "ParagraphFormatUpdate, _" { modelInstance = new(ParagraphFormatUpdate) }
                if jsonTypeStr == "ParagraphLink, _" { modelInstance = new(ParagraphLink) }
                if jsonTypeStr == "ParagraphLinkCollection, _" { modelInstance = new(ParagraphLinkCollection) }
                if jsonTypeStr == "Run, _" { modelInstance = new(Run) }
                if jsonTypeStr == "RunLink, _" { modelInstance = new(RunLink) }
                if jsonTypeStr == "Runs, _" { modelInstance = new(Runs) }
                if jsonTypeStr == "SearchResultsCollection, _" { modelInstance = new(SearchResultsCollection) }
                if jsonTypeStr == "Section, _" { modelInstance = new(Section) }
                if jsonTypeStr == "SectionLink, _" { modelInstance = new(SectionLink) }
                if jsonTypeStr == "SectionLinkCollection, _" { modelInstance = new(SectionLinkCollection) }
                if jsonTypeStr == "StructuredDocumentTag, _" { modelInstance = new(StructuredDocumentTag) }
                if jsonTypeStr == "StructuredDocumentTagCollection, _" { modelInstance = new(StructuredDocumentTagCollection) }
                if jsonTypeStr == "StructuredDocumentTagInsert, _" { modelInstance = new(StructuredDocumentTagInsert) }
                if jsonTypeStr == "StructuredDocumentTagUpdate, _" { modelInstance = new(StructuredDocumentTagUpdate) }
                if jsonTypeStr == "Style, _" { modelInstance = new(Style) }
                if jsonTypeStr == "Table, _" { modelInstance = new(Table) }
                if jsonTypeStr == "TableCell, _" { modelInstance = new(TableCell) }
                if jsonTypeStr == "TableCellFormat, _" { modelInstance = new(TableCellFormat) }
                if jsonTypeStr == "TableLink, _" { modelInstance = new(TableLink) }
                if jsonTypeStr == "TableLinkCollection, _" { modelInstance = new(TableLinkCollection) }
                if jsonTypeStr == "TableProperties, _" { modelInstance = new(TableProperties) }
                if jsonTypeStr == "TableRow, _" { modelInstance = new(TableRow) }
                if jsonTypeStr == "TableRowFormat, _" { modelInstance = new(TableRowFormat) }
            }

            if modelInstance == nil { modelInstance = new(LinkElement) }
            modelInstance.Deserialize(parsedValue)
            obj.Tables = modelInstance
        }

    } else if jsonValue, exists := json["tables"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ILinkElement = nil
            if jsonType, found := parsedValue["$type"]; found {
                jsonTypeStr := jsonType.(string)
                if jsonTypeStr == "Bookmark, _" { modelInstance = new(Bookmark) }
                if jsonTypeStr == "Bookmarks, _" { modelInstance = new(Bookmarks) }
                if jsonTypeStr == "Border, _" { modelInstance = new(Border) }
                if jsonTypeStr == "BordersCollection, _" { modelInstance = new(BordersCollection) }
                if jsonTypeStr == "Comment, _" { modelInstance = new(Comment) }
                if jsonTypeStr == "CommentLink, _" { modelInstance = new(CommentLink) }
                if jsonTypeStr == "CommentsCollection, _" { modelInstance = new(CommentsCollection) }
                if jsonTypeStr == "CustomXmlPart, _" { modelInstance = new(CustomXmlPart) }
                if jsonTypeStr == "CustomXmlPartInsert, _" { modelInstance = new(CustomXmlPartInsert) }
                if jsonTypeStr == "CustomXmlPartLink, _" { modelInstance = new(CustomXmlPartLink) }
                if jsonTypeStr == "CustomXmlPartsCollection, _" { modelInstance = new(CustomXmlPartsCollection) }
                if jsonTypeStr == "CustomXmlPartUpdate, _" { modelInstance = new(CustomXmlPartUpdate) }
                if jsonTypeStr == "DocumentProperties, _" { modelInstance = new(DocumentProperties) }
                if jsonTypeStr == "DocumentProperty, _" { modelInstance = new(DocumentProperty) }
                if jsonTypeStr == "DrawingObject, _" { modelInstance = new(DrawingObject) }
                if jsonTypeStr == "DrawingObjectCollection, _" { modelInstance = new(DrawingObjectCollection) }
                if jsonTypeStr == "DrawingObjectLink, _" { modelInstance = new(DrawingObjectLink) }
                if jsonTypeStr == "Field, _" { modelInstance = new(Field) }
                if jsonTypeStr == "FieldCollection, _" { modelInstance = new(FieldCollection) }
                if jsonTypeStr == "FieldLink, _" { modelInstance = new(FieldLink) }
                if jsonTypeStr == "FieldNames, _" { modelInstance = new(FieldNames) }
                if jsonTypeStr == "Font, _" { modelInstance = new(Font) }
                if jsonTypeStr == "Footnote, _" { modelInstance = new(Footnote) }
                if jsonTypeStr == "FootnoteCollection, _" { modelInstance = new(FootnoteCollection) }
                if jsonTypeStr == "FootnoteLink, _" { modelInstance = new(FootnoteLink) }
                if jsonTypeStr == "FormField, _" {  }
                if jsonTypeStr == "FormFieldCheckbox, _" { modelInstance = new(FormFieldCheckbox) }
                if jsonTypeStr == "FormFieldCollection, _" { modelInstance = new(FormFieldCollection) }
                if jsonTypeStr == "FormFieldDropDown, _" { modelInstance = new(FormFieldDropDown) }
                if jsonTypeStr == "FormFieldTextInput, _" { modelInstance = new(FormFieldTextInput) }
                if jsonTypeStr == "HeaderFooter, _" { modelInstance = new(HeaderFooter) }
                if jsonTypeStr == "HeaderFooterLink, _" { modelInstance = new(HeaderFooterLink) }
                if jsonTypeStr == "HeaderFooterLinkCollection, _" { modelInstance = new(HeaderFooterLinkCollection) }
                if jsonTypeStr == "Hyperlink, _" { modelInstance = new(Hyperlink) }
                if jsonTypeStr == "Hyperlinks, _" { modelInstance = new(Hyperlinks) }
                if jsonTypeStr == "ListFormat, _" { modelInstance = new(ListFormat) }
                if jsonTypeStr == "ListInfo, _" { modelInstance = new(ListInfo) }
                if jsonTypeStr == "ListLevel, _" { modelInstance = new(ListLevel) }
                if jsonTypeStr == "ListLevels, _" { modelInstance = new(ListLevels) }
                if jsonTypeStr == "Lists, _" { modelInstance = new(Lists) }
                if jsonTypeStr == "NodeLink, _" { modelInstance = new(NodeLink) }
                if jsonTypeStr == "OfficeMathLink, _" { modelInstance = new(OfficeMathLink) }
                if jsonTypeStr == "OfficeMathObject, _" { modelInstance = new(OfficeMathObject) }
                if jsonTypeStr == "OfficeMathObjectsCollection, _" { modelInstance = new(OfficeMathObjectsCollection) }
                if jsonTypeStr == "PageSetup, _" { modelInstance = new(PageSetup) }
                if jsonTypeStr == "Paragraph, _" { modelInstance = new(Paragraph) }
                if jsonTypeStr == "ParagraphFormat, _" { modelInstance = new(ParagraphFormat) }
                if jsonTypeStr == "ParagraphFormatBase, _" {  }
                if jsonTypeStr == "ParagraphFormatUpdate, _" { modelInstance = new(ParagraphFormatUpdate) }
                if jsonTypeStr == "ParagraphLink, _" { modelInstance = new(ParagraphLink) }
                if jsonTypeStr == "ParagraphLinkCollection, _" { modelInstance = new(ParagraphLinkCollection) }
                if jsonTypeStr == "Run, _" { modelInstance = new(Run) }
                if jsonTypeStr == "RunLink, _" { modelInstance = new(RunLink) }
                if jsonTypeStr == "Runs, _" { modelInstance = new(Runs) }
                if jsonTypeStr == "SearchResultsCollection, _" { modelInstance = new(SearchResultsCollection) }
                if jsonTypeStr == "Section, _" { modelInstance = new(Section) }
                if jsonTypeStr == "SectionLink, _" { modelInstance = new(SectionLink) }
                if jsonTypeStr == "SectionLinkCollection, _" { modelInstance = new(SectionLinkCollection) }
                if jsonTypeStr == "StructuredDocumentTag, _" { modelInstance = new(StructuredDocumentTag) }
                if jsonTypeStr == "StructuredDocumentTagCollection, _" { modelInstance = new(StructuredDocumentTagCollection) }
                if jsonTypeStr == "StructuredDocumentTagInsert, _" { modelInstance = new(StructuredDocumentTagInsert) }
                if jsonTypeStr == "StructuredDocumentTagUpdate, _" { modelInstance = new(StructuredDocumentTagUpdate) }
                if jsonTypeStr == "Style, _" { modelInstance = new(Style) }
                if jsonTypeStr == "Table, _" { modelInstance = new(Table) }
                if jsonTypeStr == "TableCell, _" { modelInstance = new(TableCell) }
                if jsonTypeStr == "TableCellFormat, _" { modelInstance = new(TableCellFormat) }
                if jsonTypeStr == "TableLink, _" { modelInstance = new(TableLink) }
                if jsonTypeStr == "TableLinkCollection, _" { modelInstance = new(TableLinkCollection) }
                if jsonTypeStr == "TableProperties, _" { modelInstance = new(TableProperties) }
                if jsonTypeStr == "TableRow, _" { modelInstance = new(TableRow) }
                if jsonTypeStr == "TableRowFormat, _" { modelInstance = new(TableRowFormat) }
            }

            if modelInstance == nil { modelInstance = new(LinkElement) }
            modelInstance.Deserialize(parsedValue)
            obj.Tables = modelInstance
        }

    }
}

func (obj *Section) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *Section) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Link == nil {
        return errors.New("Property Link in Section is required.")
    }

    if obj.Link != nil {
        if err := obj.Link.Validate(); err != nil {
            return err
        }
    }

    if obj.ChildNodes == nil {
        return errors.New("Property ChildNodes in Section is required.")
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

    if obj.Paragraphs == nil {
        return errors.New("Property Paragraphs in Section is required.")
    }

    if obj.Paragraphs != nil {
        if err := obj.Paragraphs.Validate(); err != nil {
            return err
        }
    }

    if obj.PageSetup == nil {
        return errors.New("Property PageSetup in Section is required.")
    }

    if obj.PageSetup != nil {
        if err := obj.PageSetup.Validate(); err != nil {
            return err
        }
    }

    if obj.HeaderFooters == nil {
        return errors.New("Property HeaderFooters in Section is required.")
    }

    if obj.HeaderFooters != nil {
        if err := obj.HeaderFooters.Validate(); err != nil {
            return err
        }
    }

    if obj.Tables == nil {
        return errors.New("Property Tables in Section is required.")
    }

    if obj.Tables != nil {
        if err := obj.Tables.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *Section) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *Section) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *Section) GetChildNodes() []INodeLink {
    return obj.ChildNodes
}

func (obj *Section) SetChildNodes(value []INodeLink) {
    obj.ChildNodes = value
}

func (obj *Section) GetParagraphs() ILinkElement {
    return obj.Paragraphs
}

func (obj *Section) SetParagraphs(value ILinkElement) {
    obj.Paragraphs = value
}

func (obj *Section) GetPageSetup() ILinkElement {
    return obj.PageSetup
}

func (obj *Section) SetPageSetup(value ILinkElement) {
    obj.PageSetup = value
}

func (obj *Section) GetHeaderFooters() ILinkElement {
    return obj.HeaderFooters
}

func (obj *Section) SetHeaderFooters(value ILinkElement) {
    obj.HeaderFooters = value
}

func (obj *Section) GetTables() ILinkElement {
    return obj.Tables
}

func (obj *Section) SetTables(value ILinkElement) {
    obj.Tables = value
}

