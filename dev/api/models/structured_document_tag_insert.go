/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="structured_document_tag_insert.go">
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

// DTO container with a StructuredDocumentTag.
type StructuredDocumentTagInsertResult struct {
    // DTO container with a StructuredDocumentTag.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // DTO container with a StructuredDocumentTag.
    NodeId string `json:"NodeId,omitempty"`

    // DTO container with a StructuredDocumentTag.
    ListItems []StructuredDocumentTagListItemResult `json:"ListItems,omitempty"`

    // DTO container with a StructuredDocumentTag.
    Checked bool `json:"Checked,omitempty"`

    // DTO container with a StructuredDocumentTag.
    Appearance string `json:"Appearance,omitempty"`

    // DTO container with a StructuredDocumentTag.
    DateDisplayLocale int32 `json:"DateDisplayLocale,omitempty"`

    // DTO container with a StructuredDocumentTag.
    DateDisplayFormat string `json:"DateDisplayFormat,omitempty"`

    // DTO container with a StructuredDocumentTag.
    FullDate Time `json:"FullDate,omitempty"`

    // DTO container with a StructuredDocumentTag.
    Title string `json:"Title,omitempty"`

    // DTO container with a StructuredDocumentTag.
    DateStorageFormat string `json:"DateStorageFormat,omitempty"`

    // DTO container with a StructuredDocumentTag.
    BuildingBlockGallery string `json:"BuildingBlockGallery,omitempty"`

    // DTO container with a StructuredDocumentTag.
    BuildingBlockCategory string `json:"BuildingBlockCategory,omitempty"`

    // DTO container with a StructuredDocumentTag.
    Multiline bool `json:"Multiline,omitempty"`

    // DTO container with a StructuredDocumentTag.
    Color string `json:"Color,omitempty"`

    // DTO container with a StructuredDocumentTag.
    StyleName string `json:"StyleName,omitempty"`

    // DTO container with a StructuredDocumentTag.
    CalendarType string `json:"CalendarType,omitempty"`

    // DTO container with a StructuredDocumentTag.
    IsTemporary bool `json:"IsTemporary,omitempty"`

    // DTO container with a StructuredDocumentTag.
    Level string `json:"Level,omitempty"`

    // DTO container with a StructuredDocumentTag.
    SdtType string `json:"SdtType,omitempty"`

    // DTO container with a StructuredDocumentTag.
    PlaceholderName string `json:"PlaceholderName,omitempty"`

    // DTO container with a StructuredDocumentTag.
    LockContentControl bool `json:"LockContentControl,omitempty"`

    // DTO container with a StructuredDocumentTag.
    LockContents bool `json:"LockContents,omitempty"`

    // DTO container with a StructuredDocumentTag.
    IsShowingPlaceholderText bool `json:"IsShowingPlaceholderText,omitempty"`

    // DTO container with a StructuredDocumentTag.
    Tag string `json:"Tag,omitempty"`

    // DTO container with a StructuredDocumentTag.
    Id int32 `json:"Id,omitempty"`

    // DTO container with a StructuredDocumentTag.
    WordOpenXML string `json:"WordOpenXML,omitempty"`
}

type StructuredDocumentTagInsert struct {
    // DTO container with a StructuredDocumentTag.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with a StructuredDocumentTag.
    NodeId *string `json:"NodeId,omitempty"`

    // DTO container with a StructuredDocumentTag.
    ListItems []StructuredDocumentTagListItem `json:"ListItems,omitempty"`

    // DTO container with a StructuredDocumentTag.
    Checked *bool `json:"Checked,omitempty"`

    // DTO container with a StructuredDocumentTag.
    Appearance *string `json:"Appearance,omitempty"`

    // DTO container with a StructuredDocumentTag.
    DateDisplayLocale *int32 `json:"DateDisplayLocale,omitempty"`

    // DTO container with a StructuredDocumentTag.
    DateDisplayFormat *string `json:"DateDisplayFormat,omitempty"`

    // DTO container with a StructuredDocumentTag.
    FullDate *Time `json:"FullDate,omitempty"`

    // DTO container with a StructuredDocumentTag.
    Title *string `json:"Title,omitempty"`

    // DTO container with a StructuredDocumentTag.
    DateStorageFormat *string `json:"DateStorageFormat,omitempty"`

    // DTO container with a StructuredDocumentTag.
    BuildingBlockGallery *string `json:"BuildingBlockGallery,omitempty"`

    // DTO container with a StructuredDocumentTag.
    BuildingBlockCategory *string `json:"BuildingBlockCategory,omitempty"`

    // DTO container with a StructuredDocumentTag.
    Multiline *bool `json:"Multiline,omitempty"`

    // DTO container with a StructuredDocumentTag.
    Color *string `json:"Color,omitempty"`

    // DTO container with a StructuredDocumentTag.
    StyleName *string `json:"StyleName,omitempty"`

    // DTO container with a StructuredDocumentTag.
    CalendarType *string `json:"CalendarType,omitempty"`

    // DTO container with a StructuredDocumentTag.
    IsTemporary *bool `json:"IsTemporary,omitempty"`

    // DTO container with a StructuredDocumentTag.
    Level *string `json:"Level,omitempty"`

    // DTO container with a StructuredDocumentTag.
    SdtType *string `json:"SdtType,omitempty"`

    // DTO container with a StructuredDocumentTag.
    PlaceholderName *string `json:"PlaceholderName,omitempty"`

    // DTO container with a StructuredDocumentTag.
    LockContentControl *bool `json:"LockContentControl,omitempty"`

    // DTO container with a StructuredDocumentTag.
    LockContents *bool `json:"LockContents,omitempty"`

    // DTO container with a StructuredDocumentTag.
    IsShowingPlaceholderText *bool `json:"IsShowingPlaceholderText,omitempty"`

    // DTO container with a StructuredDocumentTag.
    Tag *string `json:"Tag,omitempty"`

    // DTO container with a StructuredDocumentTag.
    Id *int32 `json:"Id,omitempty"`

    // DTO container with a StructuredDocumentTag.
    WordOpenXML *string `json:"WordOpenXML,omitempty"`
}

type IStructuredDocumentTagInsert interface {
    IsStructuredDocumentTagInsert() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
}

func (StructuredDocumentTagInsert) IsStructuredDocumentTagInsert() bool {
    return true
}

func (StructuredDocumentTagInsert) IsStructuredDocumentTag() bool {
    return true
}

func (StructuredDocumentTagInsert) IsNodeLink() bool {
    return true
}

func (StructuredDocumentTagInsert) IsLinkElement() bool {
    return true
}

func (obj *StructuredDocumentTagInsert) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.ListItems != nil) {
        for _, objElementListItems := range obj.ListItems {
            objElementListItems.Initialize()
        }
    }

}

func (obj *StructuredDocumentTagInsert) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}


