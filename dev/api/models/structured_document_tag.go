/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="structured_document_tag.go">
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

// DTO container with a StructuredDocumentTag.

type IStructuredDocumentTag interface {
    IsStructuredDocumentTag() bool
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
    GetColor() *string
    SetColor(value *string)
    GetStyleName() *string
    SetStyleName(value *string)
    GetCalendarType() *string
    SetCalendarType(value *string)
    GetIsTemporary() *bool
    SetIsTemporary(value *bool)
    GetLevel() *string
    SetLevel(value *string)
    GetSdtType() *string
    SetSdtType(value *string)
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

type StructuredDocumentTag struct {
    // DTO container with a StructuredDocumentTag.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with a StructuredDocumentTag.
    NodeId *string `json:"NodeId,omitempty"`

    // DTO container with a StructuredDocumentTag.
    ListItems []IStructuredDocumentTagListItem `json:"ListItems,omitempty"`

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

func (StructuredDocumentTag) IsStructuredDocumentTag() bool {
    return true
}

func (StructuredDocumentTag) IsNodeLink() bool {
    return true
}

func (StructuredDocumentTag) IsLinkElement() bool {
    return true
}

func (obj *StructuredDocumentTag) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.ListItems != nil) {
        for _, objElementListItems := range obj.ListItems {
            objElementListItems.Initialize()
        }
    }

}

func (obj *StructuredDocumentTag) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["ListItems"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.ListItems = make([]IStructuredDocumentTagListItem, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IStructuredDocumentTagListItem = new(StructuredDocumentTagListItem)
                    modelElementInstance.Deserialize(elementValue)
                    obj.ListItems = append(obj.ListItems, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["listItems"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.ListItems = make([]IStructuredDocumentTagListItem, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IStructuredDocumentTagListItem = new(StructuredDocumentTagListItem)
                    modelElementInstance.Deserialize(elementValue)
                    obj.ListItems = append(obj.ListItems, modelElementInstance)
                }

            }
        }

    }

    if jsonValue, exists := json["Checked"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Checked = &parsedValue
        }

    } else if jsonValue, exists := json["checked"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Checked = &parsedValue
        }

    }

    if jsonValue, exists := json["Appearance"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Appearance = &parsedValue
        }

    } else if jsonValue, exists := json["appearance"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Appearance = &parsedValue
        }

    }

    if jsonValue, exists := json["DateDisplayLocale"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.DateDisplayLocale = new(int32)
            *obj.DateDisplayLocale = int32(parsedValue)
        }

    } else if jsonValue, exists := json["dateDisplayLocale"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.DateDisplayLocale = new(int32)
            *obj.DateDisplayLocale = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["DateDisplayFormat"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DateDisplayFormat = &parsedValue
        }

    } else if jsonValue, exists := json["dateDisplayFormat"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DateDisplayFormat = &parsedValue
        }

    }

    if jsonValue, exists := json["FullDate"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FullDate = new(Time)
            obj.FullDate.Parse(parsedValue)
        }

    } else if jsonValue, exists := json["fullDate"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FullDate = new(Time)
            obj.FullDate.Parse(parsedValue)
        }

    }

    if jsonValue, exists := json["Title"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Title = &parsedValue
        }

    } else if jsonValue, exists := json["title"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Title = &parsedValue
        }

    }

    if jsonValue, exists := json["DateStorageFormat"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DateStorageFormat = &parsedValue
        }

    } else if jsonValue, exists := json["dateStorageFormat"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DateStorageFormat = &parsedValue
        }

    }

    if jsonValue, exists := json["BuildingBlockGallery"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.BuildingBlockGallery = &parsedValue
        }

    } else if jsonValue, exists := json["buildingBlockGallery"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.BuildingBlockGallery = &parsedValue
        }

    }

    if jsonValue, exists := json["BuildingBlockCategory"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.BuildingBlockCategory = &parsedValue
        }

    } else if jsonValue, exists := json["buildingBlockCategory"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.BuildingBlockCategory = &parsedValue
        }

    }

    if jsonValue, exists := json["Multiline"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Multiline = &parsedValue
        }

    } else if jsonValue, exists := json["multiline"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Multiline = &parsedValue
        }

    }

    if jsonValue, exists := json["Color"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Color = &parsedValue
        }

    } else if jsonValue, exists := json["color"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Color = &parsedValue
        }

    }

    if jsonValue, exists := json["StyleName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.StyleName = &parsedValue
        }

    } else if jsonValue, exists := json["styleName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.StyleName = &parsedValue
        }

    }

    if jsonValue, exists := json["CalendarType"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.CalendarType = &parsedValue
        }

    } else if jsonValue, exists := json["calendarType"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.CalendarType = &parsedValue
        }

    }

    if jsonValue, exists := json["IsTemporary"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsTemporary = &parsedValue
        }

    } else if jsonValue, exists := json["isTemporary"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsTemporary = &parsedValue
        }

    }

    if jsonValue, exists := json["Level"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Level = &parsedValue
        }

    } else if jsonValue, exists := json["level"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Level = &parsedValue
        }

    }

    if jsonValue, exists := json["SdtType"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.SdtType = &parsedValue
        }

    } else if jsonValue, exists := json["sdtType"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.SdtType = &parsedValue
        }

    }

    if jsonValue, exists := json["PlaceholderName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.PlaceholderName = &parsedValue
        }

    } else if jsonValue, exists := json["placeholderName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.PlaceholderName = &parsedValue
        }

    }

    if jsonValue, exists := json["LockContentControl"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.LockContentControl = &parsedValue
        }

    } else if jsonValue, exists := json["lockContentControl"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.LockContentControl = &parsedValue
        }

    }

    if jsonValue, exists := json["LockContents"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.LockContents = &parsedValue
        }

    } else if jsonValue, exists := json["lockContents"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.LockContents = &parsedValue
        }

    }

    if jsonValue, exists := json["IsShowingPlaceholderText"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsShowingPlaceholderText = &parsedValue
        }

    } else if jsonValue, exists := json["isShowingPlaceholderText"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsShowingPlaceholderText = &parsedValue
        }

    }

    if jsonValue, exists := json["Tag"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Tag = &parsedValue
        }

    } else if jsonValue, exists := json["tag"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Tag = &parsedValue
        }

    }

    if jsonValue, exists := json["Id"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Id = new(int32)
            *obj.Id = int32(parsedValue)
        }

    } else if jsonValue, exists := json["id"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Id = new(int32)
            *obj.Id = int32(parsedValue)
        }

    }
}

func (obj *StructuredDocumentTag) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *StructuredDocumentTag) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Link == nil {
        return errors.New("Property Link in StructuredDocumentTag is required.")
    }

    if obj.Link != nil {
        if err := obj.Link.Validate(); err != nil {
            return err
        }
    }

    if obj.NodeId == nil {
        return errors.New("Property NodeId in StructuredDocumentTag is required.")
    }

    if obj.ListItems == nil {
        return errors.New("Property ListItems in StructuredDocumentTag is required.")
    }

    if obj.ListItems != nil {
        for _, elementListItems := range obj.ListItems {
            if elementListItems != nil {
                if err := elementListItems.Validate(); err != nil {
                    return err
                }
            }
        }
    }

    if obj.Checked == nil {
        return errors.New("Property Checked in StructuredDocumentTag is required.")
    }

    if obj.Appearance == nil {
        return errors.New("Property Appearance in StructuredDocumentTag is required.")
    }

    if obj.DateDisplayLocale == nil {
        return errors.New("Property DateDisplayLocale in StructuredDocumentTag is required.")
    }

    if obj.DateDisplayFormat == nil {
        return errors.New("Property DateDisplayFormat in StructuredDocumentTag is required.")
    }

    if obj.FullDate == nil {
        return errors.New("Property FullDate in StructuredDocumentTag is required.")
    }

    if obj.Title == nil {
        return errors.New("Property Title in StructuredDocumentTag is required.")
    }

    if obj.DateStorageFormat == nil {
        return errors.New("Property DateStorageFormat in StructuredDocumentTag is required.")
    }

    if obj.BuildingBlockGallery == nil {
        return errors.New("Property BuildingBlockGallery in StructuredDocumentTag is required.")
    }

    if obj.BuildingBlockCategory == nil {
        return errors.New("Property BuildingBlockCategory in StructuredDocumentTag is required.")
    }

    if obj.Multiline == nil {
        return errors.New("Property Multiline in StructuredDocumentTag is required.")
    }

    if obj.Color == nil {
        return errors.New("Property Color in StructuredDocumentTag is required.")
    }

    if obj.StyleName == nil {
        return errors.New("Property StyleName in StructuredDocumentTag is required.")
    }

    if obj.CalendarType == nil {
        return errors.New("Property CalendarType in StructuredDocumentTag is required.")
    }

    if obj.IsTemporary == nil {
        return errors.New("Property IsTemporary in StructuredDocumentTag is required.")
    }

    if obj.Level == nil {
        return errors.New("Property Level in StructuredDocumentTag is required.")
    }

    if obj.SdtType == nil {
        return errors.New("Property SdtType in StructuredDocumentTag is required.")
    }

    if obj.PlaceholderName == nil {
        return errors.New("Property PlaceholderName in StructuredDocumentTag is required.")
    }

    if obj.LockContentControl == nil {
        return errors.New("Property LockContentControl in StructuredDocumentTag is required.")
    }

    if obj.LockContents == nil {
        return errors.New("Property LockContents in StructuredDocumentTag is required.")
    }

    if obj.IsShowingPlaceholderText == nil {
        return errors.New("Property IsShowingPlaceholderText in StructuredDocumentTag is required.")
    }

    if obj.Tag == nil {
        return errors.New("Property Tag in StructuredDocumentTag is required.")
    }

    if obj.Id == nil {
        return errors.New("Property Id in StructuredDocumentTag is required.")
    }

    if obj.WordOpenXML == nil {
        return errors.New("Property WordOpenXML in StructuredDocumentTag is required.")
    }

    return nil;
}

func (obj *StructuredDocumentTag) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *StructuredDocumentTag) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *StructuredDocumentTag) GetNodeId() *string {
    return obj.NodeId
}

func (obj *StructuredDocumentTag) SetNodeId(value *string) {
    obj.NodeId = value
}

func (obj *StructuredDocumentTag) GetListItems() []IStructuredDocumentTagListItem {
    return obj.ListItems
}

func (obj *StructuredDocumentTag) SetListItems(value []IStructuredDocumentTagListItem) {
    obj.ListItems = value
}

func (obj *StructuredDocumentTag) GetChecked() *bool {
    return obj.Checked
}

func (obj *StructuredDocumentTag) SetChecked(value *bool) {
    obj.Checked = value
}

func (obj *StructuredDocumentTag) GetAppearance() *string {
    return obj.Appearance
}

func (obj *StructuredDocumentTag) SetAppearance(value *string) {
    obj.Appearance = value
}

func (obj *StructuredDocumentTag) GetDateDisplayLocale() *int32 {
    return obj.DateDisplayLocale
}

func (obj *StructuredDocumentTag) SetDateDisplayLocale(value *int32) {
    obj.DateDisplayLocale = value
}

func (obj *StructuredDocumentTag) GetDateDisplayFormat() *string {
    return obj.DateDisplayFormat
}

func (obj *StructuredDocumentTag) SetDateDisplayFormat(value *string) {
    obj.DateDisplayFormat = value
}

func (obj *StructuredDocumentTag) GetFullDate() *Time {
    return obj.FullDate
}

func (obj *StructuredDocumentTag) SetFullDate(value *Time) {
    obj.FullDate = value
}

func (obj *StructuredDocumentTag) GetTitle() *string {
    return obj.Title
}

func (obj *StructuredDocumentTag) SetTitle(value *string) {
    obj.Title = value
}

func (obj *StructuredDocumentTag) GetDateStorageFormat() *string {
    return obj.DateStorageFormat
}

func (obj *StructuredDocumentTag) SetDateStorageFormat(value *string) {
    obj.DateStorageFormat = value
}

func (obj *StructuredDocumentTag) GetBuildingBlockGallery() *string {
    return obj.BuildingBlockGallery
}

func (obj *StructuredDocumentTag) SetBuildingBlockGallery(value *string) {
    obj.BuildingBlockGallery = value
}

func (obj *StructuredDocumentTag) GetBuildingBlockCategory() *string {
    return obj.BuildingBlockCategory
}

func (obj *StructuredDocumentTag) SetBuildingBlockCategory(value *string) {
    obj.BuildingBlockCategory = value
}

func (obj *StructuredDocumentTag) GetMultiline() *bool {
    return obj.Multiline
}

func (obj *StructuredDocumentTag) SetMultiline(value *bool) {
    obj.Multiline = value
}

func (obj *StructuredDocumentTag) GetColor() *string {
    return obj.Color
}

func (obj *StructuredDocumentTag) SetColor(value *string) {
    obj.Color = value
}

func (obj *StructuredDocumentTag) GetStyleName() *string {
    return obj.StyleName
}

func (obj *StructuredDocumentTag) SetStyleName(value *string) {
    obj.StyleName = value
}

func (obj *StructuredDocumentTag) GetCalendarType() *string {
    return obj.CalendarType
}

func (obj *StructuredDocumentTag) SetCalendarType(value *string) {
    obj.CalendarType = value
}

func (obj *StructuredDocumentTag) GetIsTemporary() *bool {
    return obj.IsTemporary
}

func (obj *StructuredDocumentTag) SetIsTemporary(value *bool) {
    obj.IsTemporary = value
}

func (obj *StructuredDocumentTag) GetLevel() *string {
    return obj.Level
}

func (obj *StructuredDocumentTag) SetLevel(value *string) {
    obj.Level = value
}

func (obj *StructuredDocumentTag) GetSdtType() *string {
    return obj.SdtType
}

func (obj *StructuredDocumentTag) SetSdtType(value *string) {
    obj.SdtType = value
}

func (obj *StructuredDocumentTag) GetPlaceholderName() *string {
    return obj.PlaceholderName
}

func (obj *StructuredDocumentTag) SetPlaceholderName(value *string) {
    obj.PlaceholderName = value
}

func (obj *StructuredDocumentTag) GetLockContentControl() *bool {
    return obj.LockContentControl
}

func (obj *StructuredDocumentTag) SetLockContentControl(value *bool) {
    obj.LockContentControl = value
}

func (obj *StructuredDocumentTag) GetLockContents() *bool {
    return obj.LockContents
}

func (obj *StructuredDocumentTag) SetLockContents(value *bool) {
    obj.LockContents = value
}

func (obj *StructuredDocumentTag) GetIsShowingPlaceholderText() *bool {
    return obj.IsShowingPlaceholderText
}

func (obj *StructuredDocumentTag) SetIsShowingPlaceholderText(value *bool) {
    obj.IsShowingPlaceholderText = value
}

func (obj *StructuredDocumentTag) GetTag() *string {
    return obj.Tag
}

func (obj *StructuredDocumentTag) SetTag(value *string) {
    obj.Tag = value
}

func (obj *StructuredDocumentTag) GetId() *int32 {
    return obj.Id
}

func (obj *StructuredDocumentTag) SetId(value *int32) {
    obj.Id = value
}

func (obj *StructuredDocumentTag) GetWordOpenXML() *string {
    return obj.WordOpenXML
}

func (obj *StructuredDocumentTag) SetWordOpenXML(value *string) {
    obj.WordOpenXML = value
}

