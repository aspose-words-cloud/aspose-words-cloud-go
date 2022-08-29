/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="style.go">
 *   Copyright (c) 2022 Aspose.Words for Cloud
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

// DTO container with a single document style.
type StyleResult struct {
    // DTO container with a single document style.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // DTO container with a single document style.
    Aliases []string `json:"Aliases,omitempty"`

    // DTO container with a single document style.
    BaseStyleName string `json:"BaseStyleName,omitempty"`

    // DTO container with a single document style.
    BuiltIn bool `json:"BuiltIn,omitempty"`

    // DTO container with a single document style.
    Font FontResult `json:"Font,omitempty"`

    // DTO container with a single document style.
    IsHeading bool `json:"IsHeading,omitempty"`

    // DTO container with a single document style.
    IsQuickStyle bool `json:"IsQuickStyle,omitempty"`

    // DTO container with a single document style.
    LinkedStyleName string `json:"LinkedStyleName,omitempty"`

    // DTO container with a single document style.
    Name string `json:"Name,omitempty"`

    // DTO container with a single document style.
    NextParagraphStyleName string `json:"NextParagraphStyleName,omitempty"`

    // DTO container with a single document style.
    StyleIdentifier string `json:"StyleIdentifier,omitempty"`

    // DTO container with a single document style.
    Type string `json:"Type,omitempty"`
}

type Style struct {
    // DTO container with a single document style.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with a single document style.
    Aliases []string `json:"Aliases,omitempty"`

    // DTO container with a single document style.
    BaseStyleName *string `json:"BaseStyleName,omitempty"`

    // DTO container with a single document style.
    BuiltIn *bool `json:"BuiltIn,omitempty"`

    // DTO container with a single document style.
    Font IFont `json:"Font,omitempty"`

    // DTO container with a single document style.
    IsHeading *bool `json:"IsHeading,omitempty"`

    // DTO container with a single document style.
    IsQuickStyle *bool `json:"IsQuickStyle,omitempty"`

    // DTO container with a single document style.
    LinkedStyleName *string `json:"LinkedStyleName,omitempty"`

    // DTO container with a single document style.
    Name *string `json:"Name,omitempty"`

    // DTO container with a single document style.
    NextParagraphStyleName *string `json:"NextParagraphStyleName,omitempty"`

    // DTO container with a single document style.
    StyleIdentifier *string `json:"StyleIdentifier,omitempty"`

    // DTO container with a single document style.
    Type *string `json:"Type,omitempty"`
}

type IStyle interface {
    IsStyle() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileContent) []FileContent
}

func (Style) IsStyle() bool {
    return true
}

func (Style) IsLinkElement() bool {
    return true
}

func (obj *Style) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.Font != nil) {
        obj.Font.Initialize()
    }


}

func (obj *Style) CollectFilesContent(resultFilesContent []FileContent) []FileContent {
    if (obj.Font != nil) {
        resultFilesContent = obj.Font.CollectFilesContent(resultFilesContent)
    }








    return resultFilesContent
}


