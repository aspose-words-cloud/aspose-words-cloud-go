/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="drawing_object_collection.go">
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

// DTO container with a collection of DrawingObjects links.

type IDrawingObjectCollection interface {
    IsDrawingObjectCollection() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetList() []ILinkElement
    SetList(value []ILinkElement)
}

type DrawingObjectCollection struct {
    // DTO container with a collection of DrawingObjects links.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with a collection of DrawingObjects links.
    List []ILinkElement `json:"List,omitempty"`
}

func (DrawingObjectCollection) IsDrawingObjectCollection() bool {
    return true
}

func (DrawingObjectCollection) IsLinkElement() bool {
    return true
}

func (obj *DrawingObjectCollection) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.List != nil) {
        for _, objElementList := range obj.List {
            objElementList.Initialize()
        }
    }

}

func (obj *DrawingObjectCollection) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["List"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.List = make([]ILinkElement, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance ILinkElement = nil
                    if jsonElementType, found := elementValue["$type"]; found {
                        jsonTypeStr := jsonElementType.(string)
                        if jsonTypeStr == "Bookmark, _" { modelElementInstance = new(Bookmark) }
                        if jsonTypeStr == "Bookmarks, _" { modelElementInstance = new(Bookmarks) }
                        if jsonTypeStr == "Border, _" { modelElementInstance = new(Border) }
                        if jsonTypeStr == "BordersCollection, _" { modelElementInstance = new(BordersCollection) }
                        if jsonTypeStr == "Comment, _" { modelElementInstance = new(Comment) }
                        if jsonTypeStr == "CommentLink, _" { modelElementInstance = new(CommentLink) }
                        if jsonTypeStr == "CommentsCollection, _" { modelElementInstance = new(CommentsCollection) }
                        if jsonTypeStr == "CustomXmlPart, _" { modelElementInstance = new(CustomXmlPart) }
                        if jsonTypeStr == "CustomXmlPartInsert, _" { modelElementInstance = new(CustomXmlPartInsert) }
                        if jsonTypeStr == "CustomXmlPartLink, _" { modelElementInstance = new(CustomXmlPartLink) }
                        if jsonTypeStr == "CustomXmlPartsCollection, _" { modelElementInstance = new(CustomXmlPartsCollection) }
                        if jsonTypeStr == "CustomXmlPartUpdate, _" { modelElementInstance = new(CustomXmlPartUpdate) }
                        if jsonTypeStr == "DocumentProperties, _" { modelElementInstance = new(DocumentProperties) }
                        if jsonTypeStr == "DocumentProperty, _" { modelElementInstance = new(DocumentProperty) }
                        if jsonTypeStr == "DrawingObject, _" { modelElementInstance = new(DrawingObject) }
                        if jsonTypeStr == "DrawingObjectCollection, _" { modelElementInstance = new(DrawingObjectCollection) }
                        if jsonTypeStr == "DrawingObjectLink, _" { modelElementInstance = new(DrawingObjectLink) }
                        if jsonTypeStr == "Field, _" { modelElementInstance = new(Field) }
                        if jsonTypeStr == "FieldCollection, _" { modelElementInstance = new(FieldCollection) }
                        if jsonTypeStr == "FieldLink, _" { modelElementInstance = new(FieldLink) }
                        if jsonTypeStr == "FieldNames, _" { modelElementInstance = new(FieldNames) }
                        if jsonTypeStr == "Font, _" { modelElementInstance = new(Font) }
                        if jsonTypeStr == "Footnote, _" { modelElementInstance = new(Footnote) }
                        if jsonTypeStr == "FootnoteCollection, _" { modelElementInstance = new(FootnoteCollection) }
                        if jsonTypeStr == "FootnoteLink, _" { modelElementInstance = new(FootnoteLink) }
                        if jsonTypeStr == "FormField, _" {  }
                        if jsonTypeStr == "FormFieldCheckbox, _" { modelElementInstance = new(FormFieldCheckbox) }
                        if jsonTypeStr == "FormFieldCollection, _" { modelElementInstance = new(FormFieldCollection) }
                        if jsonTypeStr == "FormFieldDropDown, _" { modelElementInstance = new(FormFieldDropDown) }
                        if jsonTypeStr == "FormFieldTextInput, _" { modelElementInstance = new(FormFieldTextInput) }
                        if jsonTypeStr == "HeaderFooter, _" { modelElementInstance = new(HeaderFooter) }
                        if jsonTypeStr == "HeaderFooterLink, _" { modelElementInstance = new(HeaderFooterLink) }
                        if jsonTypeStr == "HeaderFooterLinkCollection, _" { modelElementInstance = new(HeaderFooterLinkCollection) }
                        if jsonTypeStr == "Hyperlink, _" { modelElementInstance = new(Hyperlink) }
                        if jsonTypeStr == "Hyperlinks, _" { modelElementInstance = new(Hyperlinks) }
                        if jsonTypeStr == "ListFormat, _" { modelElementInstance = new(ListFormat) }
                        if jsonTypeStr == "ListInfo, _" { modelElementInstance = new(ListInfo) }
                        if jsonTypeStr == "ListLevel, _" { modelElementInstance = new(ListLevel) }
                        if jsonTypeStr == "ListLevels, _" { modelElementInstance = new(ListLevels) }
                        if jsonTypeStr == "Lists, _" { modelElementInstance = new(Lists) }
                        if jsonTypeStr == "NodeLink, _" { modelElementInstance = new(NodeLink) }
                        if jsonTypeStr == "OfficeMathLink, _" { modelElementInstance = new(OfficeMathLink) }
                        if jsonTypeStr == "OfficeMathObject, _" { modelElementInstance = new(OfficeMathObject) }
                        if jsonTypeStr == "OfficeMathObjectsCollection, _" { modelElementInstance = new(OfficeMathObjectsCollection) }
                        if jsonTypeStr == "PageSetup, _" { modelElementInstance = new(PageSetup) }
                        if jsonTypeStr == "Paragraph, _" { modelElementInstance = new(Paragraph) }
                        if jsonTypeStr == "ParagraphFormat, _" { modelElementInstance = new(ParagraphFormat) }
                        if jsonTypeStr == "ParagraphFormatBase, _" {  }
                        if jsonTypeStr == "ParagraphFormatUpdate, _" { modelElementInstance = new(ParagraphFormatUpdate) }
                        if jsonTypeStr == "ParagraphLink, _" { modelElementInstance = new(ParagraphLink) }
                        if jsonTypeStr == "ParagraphLinkCollection, _" { modelElementInstance = new(ParagraphLinkCollection) }
                        if jsonTypeStr == "Run, _" { modelElementInstance = new(Run) }
                        if jsonTypeStr == "RunLink, _" { modelElementInstance = new(RunLink) }
                        if jsonTypeStr == "Runs, _" { modelElementInstance = new(Runs) }
                        if jsonTypeStr == "SearchResultsCollection, _" { modelElementInstance = new(SearchResultsCollection) }
                        if jsonTypeStr == "Section, _" { modelElementInstance = new(Section) }
                        if jsonTypeStr == "SectionLink, _" { modelElementInstance = new(SectionLink) }
                        if jsonTypeStr == "SectionLinkCollection, _" { modelElementInstance = new(SectionLinkCollection) }
                        if jsonTypeStr == "StructuredDocumentTag, _" { modelElementInstance = new(StructuredDocumentTag) }
                        if jsonTypeStr == "StructuredDocumentTagBase, _" {  }
                        if jsonTypeStr == "StructuredDocumentTagCollection, _" { modelElementInstance = new(StructuredDocumentTagCollection) }
                        if jsonTypeStr == "StructuredDocumentTagInsert, _" { modelElementInstance = new(StructuredDocumentTagInsert) }
                        if jsonTypeStr == "StructuredDocumentTagUpdate, _" { modelElementInstance = new(StructuredDocumentTagUpdate) }
                        if jsonTypeStr == "Style, _" { modelElementInstance = new(Style) }
                        if jsonTypeStr == "Table, _" { modelElementInstance = new(Table) }
                        if jsonTypeStr == "TableCell, _" { modelElementInstance = new(TableCell) }
                        if jsonTypeStr == "TableCellFormat, _" { modelElementInstance = new(TableCellFormat) }
                        if jsonTypeStr == "TableLink, _" { modelElementInstance = new(TableLink) }
                        if jsonTypeStr == "TableLinkCollection, _" { modelElementInstance = new(TableLinkCollection) }
                        if jsonTypeStr == "TableProperties, _" { modelElementInstance = new(TableProperties) }
                        if jsonTypeStr == "TableRow, _" { modelElementInstance = new(TableRow) }
                        if jsonTypeStr == "TableRowFormat, _" { modelElementInstance = new(TableRowFormat) }
                    }

                    if modelElementInstance == nil { modelElementInstance = new(LinkElement) }
                    modelElementInstance.Deserialize(elementValue)
                    obj.List = append(obj.List, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["list"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.List = make([]ILinkElement, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance ILinkElement = nil
                    if jsonElementType, found := elementValue["$type"]; found {
                        jsonTypeStr := jsonElementType.(string)
                        if jsonTypeStr == "Bookmark, _" { modelElementInstance = new(Bookmark) }
                        if jsonTypeStr == "Bookmarks, _" { modelElementInstance = new(Bookmarks) }
                        if jsonTypeStr == "Border, _" { modelElementInstance = new(Border) }
                        if jsonTypeStr == "BordersCollection, _" { modelElementInstance = new(BordersCollection) }
                        if jsonTypeStr == "Comment, _" { modelElementInstance = new(Comment) }
                        if jsonTypeStr == "CommentLink, _" { modelElementInstance = new(CommentLink) }
                        if jsonTypeStr == "CommentsCollection, _" { modelElementInstance = new(CommentsCollection) }
                        if jsonTypeStr == "CustomXmlPart, _" { modelElementInstance = new(CustomXmlPart) }
                        if jsonTypeStr == "CustomXmlPartInsert, _" { modelElementInstance = new(CustomXmlPartInsert) }
                        if jsonTypeStr == "CustomXmlPartLink, _" { modelElementInstance = new(CustomXmlPartLink) }
                        if jsonTypeStr == "CustomXmlPartsCollection, _" { modelElementInstance = new(CustomXmlPartsCollection) }
                        if jsonTypeStr == "CustomXmlPartUpdate, _" { modelElementInstance = new(CustomXmlPartUpdate) }
                        if jsonTypeStr == "DocumentProperties, _" { modelElementInstance = new(DocumentProperties) }
                        if jsonTypeStr == "DocumentProperty, _" { modelElementInstance = new(DocumentProperty) }
                        if jsonTypeStr == "DrawingObject, _" { modelElementInstance = new(DrawingObject) }
                        if jsonTypeStr == "DrawingObjectCollection, _" { modelElementInstance = new(DrawingObjectCollection) }
                        if jsonTypeStr == "DrawingObjectLink, _" { modelElementInstance = new(DrawingObjectLink) }
                        if jsonTypeStr == "Field, _" { modelElementInstance = new(Field) }
                        if jsonTypeStr == "FieldCollection, _" { modelElementInstance = new(FieldCollection) }
                        if jsonTypeStr == "FieldLink, _" { modelElementInstance = new(FieldLink) }
                        if jsonTypeStr == "FieldNames, _" { modelElementInstance = new(FieldNames) }
                        if jsonTypeStr == "Font, _" { modelElementInstance = new(Font) }
                        if jsonTypeStr == "Footnote, _" { modelElementInstance = new(Footnote) }
                        if jsonTypeStr == "FootnoteCollection, _" { modelElementInstance = new(FootnoteCollection) }
                        if jsonTypeStr == "FootnoteLink, _" { modelElementInstance = new(FootnoteLink) }
                        if jsonTypeStr == "FormField, _" {  }
                        if jsonTypeStr == "FormFieldCheckbox, _" { modelElementInstance = new(FormFieldCheckbox) }
                        if jsonTypeStr == "FormFieldCollection, _" { modelElementInstance = new(FormFieldCollection) }
                        if jsonTypeStr == "FormFieldDropDown, _" { modelElementInstance = new(FormFieldDropDown) }
                        if jsonTypeStr == "FormFieldTextInput, _" { modelElementInstance = new(FormFieldTextInput) }
                        if jsonTypeStr == "HeaderFooter, _" { modelElementInstance = new(HeaderFooter) }
                        if jsonTypeStr == "HeaderFooterLink, _" { modelElementInstance = new(HeaderFooterLink) }
                        if jsonTypeStr == "HeaderFooterLinkCollection, _" { modelElementInstance = new(HeaderFooterLinkCollection) }
                        if jsonTypeStr == "Hyperlink, _" { modelElementInstance = new(Hyperlink) }
                        if jsonTypeStr == "Hyperlinks, _" { modelElementInstance = new(Hyperlinks) }
                        if jsonTypeStr == "ListFormat, _" { modelElementInstance = new(ListFormat) }
                        if jsonTypeStr == "ListInfo, _" { modelElementInstance = new(ListInfo) }
                        if jsonTypeStr == "ListLevel, _" { modelElementInstance = new(ListLevel) }
                        if jsonTypeStr == "ListLevels, _" { modelElementInstance = new(ListLevels) }
                        if jsonTypeStr == "Lists, _" { modelElementInstance = new(Lists) }
                        if jsonTypeStr == "NodeLink, _" { modelElementInstance = new(NodeLink) }
                        if jsonTypeStr == "OfficeMathLink, _" { modelElementInstance = new(OfficeMathLink) }
                        if jsonTypeStr == "OfficeMathObject, _" { modelElementInstance = new(OfficeMathObject) }
                        if jsonTypeStr == "OfficeMathObjectsCollection, _" { modelElementInstance = new(OfficeMathObjectsCollection) }
                        if jsonTypeStr == "PageSetup, _" { modelElementInstance = new(PageSetup) }
                        if jsonTypeStr == "Paragraph, _" { modelElementInstance = new(Paragraph) }
                        if jsonTypeStr == "ParagraphFormat, _" { modelElementInstance = new(ParagraphFormat) }
                        if jsonTypeStr == "ParagraphFormatBase, _" {  }
                        if jsonTypeStr == "ParagraphFormatUpdate, _" { modelElementInstance = new(ParagraphFormatUpdate) }
                        if jsonTypeStr == "ParagraphLink, _" { modelElementInstance = new(ParagraphLink) }
                        if jsonTypeStr == "ParagraphLinkCollection, _" { modelElementInstance = new(ParagraphLinkCollection) }
                        if jsonTypeStr == "Run, _" { modelElementInstance = new(Run) }
                        if jsonTypeStr == "RunLink, _" { modelElementInstance = new(RunLink) }
                        if jsonTypeStr == "Runs, _" { modelElementInstance = new(Runs) }
                        if jsonTypeStr == "SearchResultsCollection, _" { modelElementInstance = new(SearchResultsCollection) }
                        if jsonTypeStr == "Section, _" { modelElementInstance = new(Section) }
                        if jsonTypeStr == "SectionLink, _" { modelElementInstance = new(SectionLink) }
                        if jsonTypeStr == "SectionLinkCollection, _" { modelElementInstance = new(SectionLinkCollection) }
                        if jsonTypeStr == "StructuredDocumentTag, _" { modelElementInstance = new(StructuredDocumentTag) }
                        if jsonTypeStr == "StructuredDocumentTagBase, _" {  }
                        if jsonTypeStr == "StructuredDocumentTagCollection, _" { modelElementInstance = new(StructuredDocumentTagCollection) }
                        if jsonTypeStr == "StructuredDocumentTagInsert, _" { modelElementInstance = new(StructuredDocumentTagInsert) }
                        if jsonTypeStr == "StructuredDocumentTagUpdate, _" { modelElementInstance = new(StructuredDocumentTagUpdate) }
                        if jsonTypeStr == "Style, _" { modelElementInstance = new(Style) }
                        if jsonTypeStr == "Table, _" { modelElementInstance = new(Table) }
                        if jsonTypeStr == "TableCell, _" { modelElementInstance = new(TableCell) }
                        if jsonTypeStr == "TableCellFormat, _" { modelElementInstance = new(TableCellFormat) }
                        if jsonTypeStr == "TableLink, _" { modelElementInstance = new(TableLink) }
                        if jsonTypeStr == "TableLinkCollection, _" { modelElementInstance = new(TableLinkCollection) }
                        if jsonTypeStr == "TableProperties, _" { modelElementInstance = new(TableProperties) }
                        if jsonTypeStr == "TableRow, _" { modelElementInstance = new(TableRow) }
                        if jsonTypeStr == "TableRowFormat, _" { modelElementInstance = new(TableRowFormat) }
                    }

                    if modelElementInstance == nil { modelElementInstance = new(LinkElement) }
                    modelElementInstance.Deserialize(elementValue)
                    obj.List = append(obj.List, modelElementInstance)
                }

            }
        }

    }
}

func (obj *DrawingObjectCollection) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *DrawingObjectCollection) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Link != nil {
        if err := obj.Link.Validate(); err != nil {
            return err
        }
    }
    if obj.List != nil {
        for _, elementList := range obj.List {
            if elementList != nil {
                if err := elementList.Validate(); err != nil {
                    return err
                }
            }
        }
    }

    return nil;
}

func (obj *DrawingObjectCollection) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *DrawingObjectCollection) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *DrawingObjectCollection) GetList() []ILinkElement {
    return obj.List
}

func (obj *DrawingObjectCollection) SetList(value []ILinkElement) {
    obj.List = value
}

