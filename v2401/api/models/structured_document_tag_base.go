/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="structured_document_tag_base.go">
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

// DTO container with a StructuredDocumentTagBaseDto.

type IStructuredDocumentTagBase interface {
    IsStructuredDocumentTagBase() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetNodeId() *string
    SetNodeId(value *string)
    GetListItems() []IStructuredDocumentTagListItem
    SetListItems(value []IStructuredDocumentTagListItem)
    GetChecked() *bool
    SetChecked(value *bool)
    GetAppearance() *string
    SetAppearance(value *string)
    GetDateDisplayLocale() *int32
    SetDateDisplayLocale(value *int32)
    GetDateDisplayFormat() *string
    SetDateDisplayFormat(value *string)
    GetFullDate() *Time
    SetFullDate(value *Time)
    GetTitle() *string
    SetTitle(value *string)
    GetDateStorageFormat() *string
    SetDateStorageFormat(value *string)
    GetBuildingBlockGallery() *string
    SetBuildingBlockGallery(value *string)
    GetBuildingBlockCategory() *string
    SetBuildingBlockCategory(value *string)
    GetMultiline() *bool
    SetMultiline(value *bool)
    GetColor() IXmlColor
    SetColor(value IXmlColor)
    GetStyleName() *string
    SetStyleName(value *string)
    GetCalendarType() *string
    SetCalendarType(value *string)
    GetIsTemporary() *bool
    SetIsTemporary(value *bool)
    GetPlaceholderName() *string
    SetPlaceholderName(value *string)
    GetLockContentControl() *bool
    SetLockContentControl(value *bool)
    GetLockContents() *bool
    SetLockContents(value *bool)
    GetIsShowingPlaceholderText() *bool
    SetIsShowingPlaceholderText(value *bool)
    GetTag() *string
    SetTag(value *string)
    GetId() *int32
    SetId(value *int32)
    GetWordOpenXML() *string
    SetWordOpenXML(value *string)
}

