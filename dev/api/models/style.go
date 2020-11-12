/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="style.go">
 *   Copyright (c) 2020 Aspose.Words for Cloud
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

// Represents a single document style.
type Style struct {
    // Represents a single document style.
    Link *WordsApiLink `json:"Link,omitempty"`

    // Represents a single document style.
    Aliases []string `json:"Aliases,omitempty"`

    // Represents a single document style.
    BaseStyleName *string `json:"BaseStyleName,omitempty"`

    // Represents a single document style.
    BuiltIn *bool `json:"BuiltIn,omitempty"`

    // Represents a single document style.
    Font *Font `json:"Font,omitempty"`

    // Represents a single document style.
    IsHeading *bool `json:"IsHeading,omitempty"`

    // Represents a single document style.
    IsQuickStyle *bool `json:"IsQuickStyle,omitempty"`

    // Represents a single document style.
    LinkedStyleName *string `json:"LinkedStyleName,omitempty"`

    // Represents a single document style.
    Name *string `json:"Name,omitempty"`

    // Represents a single document style.
    NextParagraphStyleName *string `json:"NextParagraphStyleName,omitempty"`

    // Represents a single document style.
    StyleIdentifier string `json:"StyleIdentifier,omitempty"`

    // Represents a single document style.
    Type string `json:"Type,omitempty"`
}

type IStyle interface {
    IsStyle() bool
}
func (Style) IsStyle() bool {
    return true
}

func (Style) IsLinkElement() bool {
    return true
}
