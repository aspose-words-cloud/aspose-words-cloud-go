/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="page_setup.go">
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

// Represents the page setup properties of a section.
// PageSetup object contains all the page setup attributes of a section (left margin, bottom margin, paper size, and so on) as properties.

type IPageSetup interface {
    IsPageSetup() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetBidi() *bool
    SetBidi(value *bool)
    GetBorderAlwaysInFront() *bool
    SetBorderAlwaysInFront(value *bool)
    GetBorderAppliesTo() *string
    SetBorderAppliesTo(value *string)
    GetBorderDistanceFrom() *string
    SetBorderDistanceFrom(value *string)
    GetBottomMargin() *float64
    SetBottomMargin(value *float64)
    GetDifferentFirstPageHeaderFooter() *bool
    SetDifferentFirstPageHeaderFooter(value *bool)
    GetFirstPageTray() *int32
    SetFirstPageTray(value *int32)
    GetFooterDistance() *float64
    SetFooterDistance(value *float64)
    GetGutter() *float64
    SetGutter(value *float64)
    GetHeaderDistance() *float64
    SetHeaderDistance(value *float64)
    GetLeftMargin() *float64
    SetLeftMargin(value *float64)
    GetLineNumberCountBy() *int32
    SetLineNumberCountBy(value *int32)
    GetLineNumberDistanceFromText() *float64
    SetLineNumberDistanceFromText(value *float64)
    GetLineNumberRestartMode() *string
    SetLineNumberRestartMode(value *string)
    GetLineStartingNumber() *int32
    SetLineStartingNumber(value *int32)
    GetOrientation() *string
    SetOrientation(value *string)
    GetOtherPagesTray() *int32
    SetOtherPagesTray(value *int32)
    GetPageHeight() *float64
    SetPageHeight(value *float64)
    GetPageNumberStyle() *string
    SetPageNumberStyle(value *string)
    GetPageStartingNumber() *int32
    SetPageStartingNumber(value *int32)
    GetPageWidth() *float64
    SetPageWidth(value *float64)
    GetPaperSize() *string
    SetPaperSize(value *string)
    GetRestartPageNumbering() *bool
    SetRestartPageNumbering(value *bool)
    GetRightMargin() *float64
    SetRightMargin(value *float64)
    GetRtlGutter() *bool
    SetRtlGutter(value *bool)
    GetSectionStart() *string
    SetSectionStart(value *string)
    GetSuppressEndnotes() *bool
    SetSuppressEndnotes(value *bool)
    GetTopMargin() *float64
    SetTopMargin(value *float64)
    GetVerticalAlignment() *string
    SetVerticalAlignment(value *string)
}

type PageSetup struct {
    // Represents the page setup properties of a section.
    // PageSetup object contains all the page setup attributes of a section (left margin, bottom margin, paper size, and so on) as properties.
    Link IWordsApiLink `json:"Link,omitempty"`

    // Represents the page setup properties of a section.
    // PageSetup object contains all the page setup attributes of a section (left margin, bottom margin, paper size, and so on) as properties.
    Bidi *bool `json:"Bidi,omitempty"`

    // Represents the page setup properties of a section.
    // PageSetup object contains all the page setup attributes of a section (left margin, bottom margin, paper size, and so on) as properties.
    BorderAlwaysInFront *bool `json:"BorderAlwaysInFront,omitempty"`

    // Represents the page setup properties of a section.
    // PageSetup object contains all the page setup attributes of a section (left margin, bottom margin, paper size, and so on) as properties.
    BorderAppliesTo *string `json:"BorderAppliesTo,omitempty"`

    // Represents the page setup properties of a section.
    // PageSetup object contains all the page setup attributes of a section (left margin, bottom margin, paper size, and so on) as properties.
    BorderDistanceFrom *string `json:"BorderDistanceFrom,omitempty"`

    // Represents the page setup properties of a section.
    // PageSetup object contains all the page setup attributes of a section (left margin, bottom margin, paper size, and so on) as properties.
    BottomMargin *float64 `json:"BottomMargin,omitempty"`

    // Represents the page setup properties of a section.
    // PageSetup object contains all the page setup attributes of a section (left margin, bottom margin, paper size, and so on) as properties.
    DifferentFirstPageHeaderFooter *bool `json:"DifferentFirstPageHeaderFooter,omitempty"`

    // Represents the page setup properties of a section.
    // PageSetup object contains all the page setup attributes of a section (left margin, bottom margin, paper size, and so on) as properties.
    FirstPageTray *int32 `json:"FirstPageTray,omitempty"`

    // Represents the page setup properties of a section.
    // PageSetup object contains all the page setup attributes of a section (left margin, bottom margin, paper size, and so on) as properties.
    FooterDistance *float64 `json:"FooterDistance,omitempty"`

    // Represents the page setup properties of a section.
    // PageSetup object contains all the page setup attributes of a section (left margin, bottom margin, paper size, and so on) as properties.
    Gutter *float64 `json:"Gutter,omitempty"`

    // Represents the page setup properties of a section.
    // PageSetup object contains all the page setup attributes of a section (left margin, bottom margin, paper size, and so on) as properties.
    HeaderDistance *float64 `json:"HeaderDistance,omitempty"`

    // Represents the page setup properties of a section.
    // PageSetup object contains all the page setup attributes of a section (left margin, bottom margin, paper size, and so on) as properties.
    LeftMargin *float64 `json:"LeftMargin,omitempty"`

    // Represents the page setup properties of a section.
    // PageSetup object contains all the page setup attributes of a section (left margin, bottom margin, paper size, and so on) as properties.
    LineNumberCountBy *int32 `json:"LineNumberCountBy,omitempty"`

    // Represents the page setup properties of a section.
    // PageSetup object contains all the page setup attributes of a section (left margin, bottom margin, paper size, and so on) as properties.
    LineNumberDistanceFromText *float64 `json:"LineNumberDistanceFromText,omitempty"`

    // Represents the page setup properties of a section.
    // PageSetup object contains all the page setup attributes of a section (left margin, bottom margin, paper size, and so on) as properties.
    LineNumberRestartMode *string `json:"LineNumberRestartMode,omitempty"`

    // Represents the page setup properties of a section.
    // PageSetup object contains all the page setup attributes of a section (left margin, bottom margin, paper size, and so on) as properties.
    LineStartingNumber *int32 `json:"LineStartingNumber,omitempty"`

    // Represents the page setup properties of a section.
    // PageSetup object contains all the page setup attributes of a section (left margin, bottom margin, paper size, and so on) as properties.
    Orientation *string `json:"Orientation,omitempty"`

    // Represents the page setup properties of a section.
    // PageSetup object contains all the page setup attributes of a section (left margin, bottom margin, paper size, and so on) as properties.
    OtherPagesTray *int32 `json:"OtherPagesTray,omitempty"`

    // Represents the page setup properties of a section.
    // PageSetup object contains all the page setup attributes of a section (left margin, bottom margin, paper size, and so on) as properties.
    PageHeight *float64 `json:"PageHeight,omitempty"`

    // Represents the page setup properties of a section.
    // PageSetup object contains all the page setup attributes of a section (left margin, bottom margin, paper size, and so on) as properties.
    PageNumberStyle *string `json:"PageNumberStyle,omitempty"`

    // Represents the page setup properties of a section.
    // PageSetup object contains all the page setup attributes of a section (left margin, bottom margin, paper size, and so on) as properties.
    PageStartingNumber *int32 `json:"PageStartingNumber,omitempty"`

    // Represents the page setup properties of a section.
    // PageSetup object contains all the page setup attributes of a section (left margin, bottom margin, paper size, and so on) as properties.
    PageWidth *float64 `json:"PageWidth,omitempty"`

    // Represents the page setup properties of a section.
    // PageSetup object contains all the page setup attributes of a section (left margin, bottom margin, paper size, and so on) as properties.
    PaperSize *string `json:"PaperSize,omitempty"`

    // Represents the page setup properties of a section.
    // PageSetup object contains all the page setup attributes of a section (left margin, bottom margin, paper size, and so on) as properties.
    RestartPageNumbering *bool `json:"RestartPageNumbering,omitempty"`

    // Represents the page setup properties of a section.
    // PageSetup object contains all the page setup attributes of a section (left margin, bottom margin, paper size, and so on) as properties.
    RightMargin *float64 `json:"RightMargin,omitempty"`

    // Represents the page setup properties of a section.
    // PageSetup object contains all the page setup attributes of a section (left margin, bottom margin, paper size, and so on) as properties.
    RtlGutter *bool `json:"RtlGutter,omitempty"`

    // Represents the page setup properties of a section.
    // PageSetup object contains all the page setup attributes of a section (left margin, bottom margin, paper size, and so on) as properties.
    SectionStart *string `json:"SectionStart,omitempty"`

    // Represents the page setup properties of a section.
    // PageSetup object contains all the page setup attributes of a section (left margin, bottom margin, paper size, and so on) as properties.
    SuppressEndnotes *bool `json:"SuppressEndnotes,omitempty"`

    // Represents the page setup properties of a section.
    // PageSetup object contains all the page setup attributes of a section (left margin, bottom margin, paper size, and so on) as properties.
    TopMargin *float64 `json:"TopMargin,omitempty"`

    // Represents the page setup properties of a section.
    // PageSetup object contains all the page setup attributes of a section (left margin, bottom margin, paper size, and so on) as properties.
    VerticalAlignment *string `json:"VerticalAlignment,omitempty"`
}

func (PageSetup) IsPageSetup() bool {
    return true
}

func (PageSetup) IsLinkElement() bool {
    return true
}

func (obj *PageSetup) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }


}

func (obj *PageSetup) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["Bidi"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Bidi = &parsedValue
        }

    } else if jsonValue, exists := json["bidi"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Bidi = &parsedValue
        }

    }

    if jsonValue, exists := json["BorderAlwaysInFront"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.BorderAlwaysInFront = &parsedValue
        }

    } else if jsonValue, exists := json["borderAlwaysInFront"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.BorderAlwaysInFront = &parsedValue
        }

    }

    if jsonValue, exists := json["BorderAppliesTo"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.BorderAppliesTo = &parsedValue
        }

    } else if jsonValue, exists := json["borderAppliesTo"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.BorderAppliesTo = &parsedValue
        }

    }

    if jsonValue, exists := json["BorderDistanceFrom"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.BorderDistanceFrom = &parsedValue
        }

    } else if jsonValue, exists := json["borderDistanceFrom"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.BorderDistanceFrom = &parsedValue
        }

    }

    if jsonValue, exists := json["BottomMargin"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.BottomMargin = &parsedValue
        }

    } else if jsonValue, exists := json["bottomMargin"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.BottomMargin = &parsedValue
        }

    }

    if jsonValue, exists := json["DifferentFirstPageHeaderFooter"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.DifferentFirstPageHeaderFooter = &parsedValue
        }

    } else if jsonValue, exists := json["differentFirstPageHeaderFooter"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.DifferentFirstPageHeaderFooter = &parsedValue
        }

    }

    if jsonValue, exists := json["FirstPageTray"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.FirstPageTray = new(int32)
            *obj.FirstPageTray = int32(parsedValue)
        }

    } else if jsonValue, exists := json["firstPageTray"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.FirstPageTray = new(int32)
            *obj.FirstPageTray = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["FooterDistance"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.FooterDistance = &parsedValue
        }

    } else if jsonValue, exists := json["footerDistance"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.FooterDistance = &parsedValue
        }

    }

    if jsonValue, exists := json["Gutter"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Gutter = &parsedValue
        }

    } else if jsonValue, exists := json["gutter"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Gutter = &parsedValue
        }

    }

    if jsonValue, exists := json["HeaderDistance"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.HeaderDistance = &parsedValue
        }

    } else if jsonValue, exists := json["headerDistance"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.HeaderDistance = &parsedValue
        }

    }

    if jsonValue, exists := json["LeftMargin"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.LeftMargin = &parsedValue
        }

    } else if jsonValue, exists := json["leftMargin"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.LeftMargin = &parsedValue
        }

    }

    if jsonValue, exists := json["LineNumberCountBy"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.LineNumberCountBy = new(int32)
            *obj.LineNumberCountBy = int32(parsedValue)
        }

    } else if jsonValue, exists := json["lineNumberCountBy"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.LineNumberCountBy = new(int32)
            *obj.LineNumberCountBy = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["LineNumberDistanceFromText"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.LineNumberDistanceFromText = &parsedValue
        }

    } else if jsonValue, exists := json["lineNumberDistanceFromText"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.LineNumberDistanceFromText = &parsedValue
        }

    }

    if jsonValue, exists := json["LineNumberRestartMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.LineNumberRestartMode = &parsedValue
        }

    } else if jsonValue, exists := json["lineNumberRestartMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.LineNumberRestartMode = &parsedValue
        }

    }

    if jsonValue, exists := json["LineStartingNumber"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.LineStartingNumber = new(int32)
            *obj.LineStartingNumber = int32(parsedValue)
        }

    } else if jsonValue, exists := json["lineStartingNumber"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.LineStartingNumber = new(int32)
            *obj.LineStartingNumber = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["Orientation"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Orientation = &parsedValue
        }

    } else if jsonValue, exists := json["orientation"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Orientation = &parsedValue
        }

    }

    if jsonValue, exists := json["OtherPagesTray"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.OtherPagesTray = new(int32)
            *obj.OtherPagesTray = int32(parsedValue)
        }

    } else if jsonValue, exists := json["otherPagesTray"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.OtherPagesTray = new(int32)
            *obj.OtherPagesTray = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["PageHeight"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.PageHeight = &parsedValue
        }

    } else if jsonValue, exists := json["pageHeight"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.PageHeight = &parsedValue
        }

    }

    if jsonValue, exists := json["PageNumberStyle"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.PageNumberStyle = &parsedValue
        }

    } else if jsonValue, exists := json["pageNumberStyle"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.PageNumberStyle = &parsedValue
        }

    }

    if jsonValue, exists := json["PageStartingNumber"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.PageStartingNumber = new(int32)
            *obj.PageStartingNumber = int32(parsedValue)
        }

    } else if jsonValue, exists := json["pageStartingNumber"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.PageStartingNumber = new(int32)
            *obj.PageStartingNumber = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["PageWidth"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.PageWidth = &parsedValue
        }

    } else if jsonValue, exists := json["pageWidth"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.PageWidth = &parsedValue
        }

    }

    if jsonValue, exists := json["PaperSize"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.PaperSize = &parsedValue
        }

    } else if jsonValue, exists := json["paperSize"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.PaperSize = &parsedValue
        }

    }

    if jsonValue, exists := json["RestartPageNumbering"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.RestartPageNumbering = &parsedValue
        }

    } else if jsonValue, exists := json["restartPageNumbering"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.RestartPageNumbering = &parsedValue
        }

    }

    if jsonValue, exists := json["RightMargin"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.RightMargin = &parsedValue
        }

    } else if jsonValue, exists := json["rightMargin"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.RightMargin = &parsedValue
        }

    }

    if jsonValue, exists := json["RtlGutter"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.RtlGutter = &parsedValue
        }

    } else if jsonValue, exists := json["rtlGutter"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.RtlGutter = &parsedValue
        }

    }

    if jsonValue, exists := json["SectionStart"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.SectionStart = &parsedValue
        }

    } else if jsonValue, exists := json["sectionStart"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.SectionStart = &parsedValue
        }

    }

    if jsonValue, exists := json["SuppressEndnotes"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.SuppressEndnotes = &parsedValue
        }

    } else if jsonValue, exists := json["suppressEndnotes"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.SuppressEndnotes = &parsedValue
        }

    }

    if jsonValue, exists := json["TopMargin"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.TopMargin = &parsedValue
        }

    } else if jsonValue, exists := json["topMargin"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.TopMargin = &parsedValue
        }

    }

    if jsonValue, exists := json["VerticalAlignment"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.VerticalAlignment = &parsedValue
        }

    } else if jsonValue, exists := json["verticalAlignment"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.VerticalAlignment = &parsedValue
        }

    }
}

func (obj *PageSetup) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *PageSetup) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Link != nil {
        if err := obj.Link.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *PageSetup) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *PageSetup) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *PageSetup) GetBidi() *bool {
    return obj.Bidi
}

func (obj *PageSetup) SetBidi(value *bool) {
    obj.Bidi = value
}

func (obj *PageSetup) GetBorderAlwaysInFront() *bool {
    return obj.BorderAlwaysInFront
}

func (obj *PageSetup) SetBorderAlwaysInFront(value *bool) {
    obj.BorderAlwaysInFront = value
}

func (obj *PageSetup) GetBorderAppliesTo() *string {
    return obj.BorderAppliesTo
}

func (obj *PageSetup) SetBorderAppliesTo(value *string) {
    obj.BorderAppliesTo = value
}

func (obj *PageSetup) GetBorderDistanceFrom() *string {
    return obj.BorderDistanceFrom
}

func (obj *PageSetup) SetBorderDistanceFrom(value *string) {
    obj.BorderDistanceFrom = value
}

func (obj *PageSetup) GetBottomMargin() *float64 {
    return obj.BottomMargin
}

func (obj *PageSetup) SetBottomMargin(value *float64) {
    obj.BottomMargin = value
}

func (obj *PageSetup) GetDifferentFirstPageHeaderFooter() *bool {
    return obj.DifferentFirstPageHeaderFooter
}

func (obj *PageSetup) SetDifferentFirstPageHeaderFooter(value *bool) {
    obj.DifferentFirstPageHeaderFooter = value
}

func (obj *PageSetup) GetFirstPageTray() *int32 {
    return obj.FirstPageTray
}

func (obj *PageSetup) SetFirstPageTray(value *int32) {
    obj.FirstPageTray = value
}

func (obj *PageSetup) GetFooterDistance() *float64 {
    return obj.FooterDistance
}

func (obj *PageSetup) SetFooterDistance(value *float64) {
    obj.FooterDistance = value
}

func (obj *PageSetup) GetGutter() *float64 {
    return obj.Gutter
}

func (obj *PageSetup) SetGutter(value *float64) {
    obj.Gutter = value
}

func (obj *PageSetup) GetHeaderDistance() *float64 {
    return obj.HeaderDistance
}

func (obj *PageSetup) SetHeaderDistance(value *float64) {
    obj.HeaderDistance = value
}

func (obj *PageSetup) GetLeftMargin() *float64 {
    return obj.LeftMargin
}

func (obj *PageSetup) SetLeftMargin(value *float64) {
    obj.LeftMargin = value
}

func (obj *PageSetup) GetLineNumberCountBy() *int32 {
    return obj.LineNumberCountBy
}

func (obj *PageSetup) SetLineNumberCountBy(value *int32) {
    obj.LineNumberCountBy = value
}

func (obj *PageSetup) GetLineNumberDistanceFromText() *float64 {
    return obj.LineNumberDistanceFromText
}

func (obj *PageSetup) SetLineNumberDistanceFromText(value *float64) {
    obj.LineNumberDistanceFromText = value
}

func (obj *PageSetup) GetLineNumberRestartMode() *string {
    return obj.LineNumberRestartMode
}

func (obj *PageSetup) SetLineNumberRestartMode(value *string) {
    obj.LineNumberRestartMode = value
}

func (obj *PageSetup) GetLineStartingNumber() *int32 {
    return obj.LineStartingNumber
}

func (obj *PageSetup) SetLineStartingNumber(value *int32) {
    obj.LineStartingNumber = value
}

func (obj *PageSetup) GetOrientation() *string {
    return obj.Orientation
}

func (obj *PageSetup) SetOrientation(value *string) {
    obj.Orientation = value
}

func (obj *PageSetup) GetOtherPagesTray() *int32 {
    return obj.OtherPagesTray
}

func (obj *PageSetup) SetOtherPagesTray(value *int32) {
    obj.OtherPagesTray = value
}

func (obj *PageSetup) GetPageHeight() *float64 {
    return obj.PageHeight
}

func (obj *PageSetup) SetPageHeight(value *float64) {
    obj.PageHeight = value
}

func (obj *PageSetup) GetPageNumberStyle() *string {
    return obj.PageNumberStyle
}

func (obj *PageSetup) SetPageNumberStyle(value *string) {
    obj.PageNumberStyle = value
}

func (obj *PageSetup) GetPageStartingNumber() *int32 {
    return obj.PageStartingNumber
}

func (obj *PageSetup) SetPageStartingNumber(value *int32) {
    obj.PageStartingNumber = value
}

func (obj *PageSetup) GetPageWidth() *float64 {
    return obj.PageWidth
}

func (obj *PageSetup) SetPageWidth(value *float64) {
    obj.PageWidth = value
}

func (obj *PageSetup) GetPaperSize() *string {
    return obj.PaperSize
}

func (obj *PageSetup) SetPaperSize(value *string) {
    obj.PaperSize = value
}

func (obj *PageSetup) GetRestartPageNumbering() *bool {
    return obj.RestartPageNumbering
}

func (obj *PageSetup) SetRestartPageNumbering(value *bool) {
    obj.RestartPageNumbering = value
}

func (obj *PageSetup) GetRightMargin() *float64 {
    return obj.RightMargin
}

func (obj *PageSetup) SetRightMargin(value *float64) {
    obj.RightMargin = value
}

func (obj *PageSetup) GetRtlGutter() *bool {
    return obj.RtlGutter
}

func (obj *PageSetup) SetRtlGutter(value *bool) {
    obj.RtlGutter = value
}

func (obj *PageSetup) GetSectionStart() *string {
    return obj.SectionStart
}

func (obj *PageSetup) SetSectionStart(value *string) {
    obj.SectionStart = value
}

func (obj *PageSetup) GetSuppressEndnotes() *bool {
    return obj.SuppressEndnotes
}

func (obj *PageSetup) SetSuppressEndnotes(value *bool) {
    obj.SuppressEndnotes = value
}

func (obj *PageSetup) GetTopMargin() *float64 {
    return obj.TopMargin
}

func (obj *PageSetup) SetTopMargin(value *float64) {
    obj.TopMargin = value
}

func (obj *PageSetup) GetVerticalAlignment() *string {
    return obj.VerticalAlignment
}

func (obj *PageSetup) SetVerticalAlignment(value *string) {
    obj.VerticalAlignment = value
}

