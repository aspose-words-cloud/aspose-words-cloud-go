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

// Represents the page setup properties of a section.
type PageSetupResult struct {
    // Represents the page setup properties of a section.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // Represents the page setup properties of a section.
    Bidi bool `json:"Bidi,omitempty"`

    // Represents the page setup properties of a section.
    BorderAlwaysInFront bool `json:"BorderAlwaysInFront,omitempty"`

    // Represents the page setup properties of a section.
    BorderAppliesTo string `json:"BorderAppliesTo,omitempty"`

    // Represents the page setup properties of a section.
    BorderDistanceFrom string `json:"BorderDistanceFrom,omitempty"`

    // Represents the page setup properties of a section.
    BottomMargin float64 `json:"BottomMargin,omitempty"`

    // Represents the page setup properties of a section.
    DifferentFirstPageHeaderFooter bool `json:"DifferentFirstPageHeaderFooter,omitempty"`

    // Represents the page setup properties of a section.
    FirstPageTray int32 `json:"FirstPageTray,omitempty"`

    // Represents the page setup properties of a section.
    FooterDistance float64 `json:"FooterDistance,omitempty"`

    // Represents the page setup properties of a section.
    Gutter float64 `json:"Gutter,omitempty"`

    // Represents the page setup properties of a section.
    HeaderDistance float64 `json:"HeaderDistance,omitempty"`

    // Represents the page setup properties of a section.
    LeftMargin float64 `json:"LeftMargin,omitempty"`

    // Represents the page setup properties of a section.
    LineNumberCountBy int32 `json:"LineNumberCountBy,omitempty"`

    // Represents the page setup properties of a section.
    LineNumberDistanceFromText float64 `json:"LineNumberDistanceFromText,omitempty"`

    // Represents the page setup properties of a section.
    LineNumberRestartMode string `json:"LineNumberRestartMode,omitempty"`

    // Represents the page setup properties of a section.
    LineStartingNumber int32 `json:"LineStartingNumber,omitempty"`

    // Represents the page setup properties of a section.
    Orientation string `json:"Orientation,omitempty"`

    // Represents the page setup properties of a section.
    OtherPagesTray int32 `json:"OtherPagesTray,omitempty"`

    // Represents the page setup properties of a section.
    PageHeight float64 `json:"PageHeight,omitempty"`

    // Represents the page setup properties of a section.
    PageNumberStyle string `json:"PageNumberStyle,omitempty"`

    // Represents the page setup properties of a section.
    PageStartingNumber int32 `json:"PageStartingNumber,omitempty"`

    // Represents the page setup properties of a section.
    PageWidth float64 `json:"PageWidth,omitempty"`

    // Represents the page setup properties of a section.
    PaperSize string `json:"PaperSize,omitempty"`

    // Represents the page setup properties of a section.
    RestartPageNumbering bool `json:"RestartPageNumbering,omitempty"`

    // Represents the page setup properties of a section.
    RightMargin float64 `json:"RightMargin,omitempty"`

    // Represents the page setup properties of a section.
    RtlGutter bool `json:"RtlGutter,omitempty"`

    // Represents the page setup properties of a section.
    SectionStart string `json:"SectionStart,omitempty"`

    // Represents the page setup properties of a section.
    SuppressEndnotes bool `json:"SuppressEndnotes,omitempty"`

    // Represents the page setup properties of a section.
    TopMargin float64 `json:"TopMargin,omitempty"`

    // Represents the page setup properties of a section.
    VerticalAlignment string `json:"VerticalAlignment,omitempty"`
}

type PageSetup struct {
    // Represents the page setup properties of a section.
    Link IWordsApiLink `json:"Link,omitempty"`

    // Represents the page setup properties of a section.
    Bidi *bool `json:"Bidi,omitempty"`

    // Represents the page setup properties of a section.
    BorderAlwaysInFront *bool `json:"BorderAlwaysInFront,omitempty"`

    // Represents the page setup properties of a section.
    BorderAppliesTo *string `json:"BorderAppliesTo,omitempty"`

    // Represents the page setup properties of a section.
    BorderDistanceFrom *string `json:"BorderDistanceFrom,omitempty"`

    // Represents the page setup properties of a section.
    BottomMargin *float64 `json:"BottomMargin,omitempty"`

    // Represents the page setup properties of a section.
    DifferentFirstPageHeaderFooter *bool `json:"DifferentFirstPageHeaderFooter,omitempty"`

    // Represents the page setup properties of a section.
    FirstPageTray *int32 `json:"FirstPageTray,omitempty"`

    // Represents the page setup properties of a section.
    FooterDistance *float64 `json:"FooterDistance,omitempty"`

    // Represents the page setup properties of a section.
    Gutter *float64 `json:"Gutter,omitempty"`

    // Represents the page setup properties of a section.
    HeaderDistance *float64 `json:"HeaderDistance,omitempty"`

    // Represents the page setup properties of a section.
    LeftMargin *float64 `json:"LeftMargin,omitempty"`

    // Represents the page setup properties of a section.
    LineNumberCountBy *int32 `json:"LineNumberCountBy,omitempty"`

    // Represents the page setup properties of a section.
    LineNumberDistanceFromText *float64 `json:"LineNumberDistanceFromText,omitempty"`

    // Represents the page setup properties of a section.
    LineNumberRestartMode *string `json:"LineNumberRestartMode,omitempty"`

    // Represents the page setup properties of a section.
    LineStartingNumber *int32 `json:"LineStartingNumber,omitempty"`

    // Represents the page setup properties of a section.
    Orientation *string `json:"Orientation,omitempty"`

    // Represents the page setup properties of a section.
    OtherPagesTray *int32 `json:"OtherPagesTray,omitempty"`

    // Represents the page setup properties of a section.
    PageHeight *float64 `json:"PageHeight,omitempty"`

    // Represents the page setup properties of a section.
    PageNumberStyle *string `json:"PageNumberStyle,omitempty"`

    // Represents the page setup properties of a section.
    PageStartingNumber *int32 `json:"PageStartingNumber,omitempty"`

    // Represents the page setup properties of a section.
    PageWidth *float64 `json:"PageWidth,omitempty"`

    // Represents the page setup properties of a section.
    PaperSize *string `json:"PaperSize,omitempty"`

    // Represents the page setup properties of a section.
    RestartPageNumbering *bool `json:"RestartPageNumbering,omitempty"`

    // Represents the page setup properties of a section.
    RightMargin *float64 `json:"RightMargin,omitempty"`

    // Represents the page setup properties of a section.
    RtlGutter *bool `json:"RtlGutter,omitempty"`

    // Represents the page setup properties of a section.
    SectionStart *string `json:"SectionStart,omitempty"`

    // Represents the page setup properties of a section.
    SuppressEndnotes *bool `json:"SuppressEndnotes,omitempty"`

    // Represents the page setup properties of a section.
    TopMargin *float64 `json:"TopMargin,omitempty"`

    // Represents the page setup properties of a section.
    VerticalAlignment *string `json:"VerticalAlignment,omitempty"`
}

type IPageSetup interface {
    IsPageSetup() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
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

func (obj *PageSetup) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}


