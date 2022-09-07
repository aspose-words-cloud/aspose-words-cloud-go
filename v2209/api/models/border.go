/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="border.go">
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

// Represents a border of an object.
type BorderResult struct {
    // Represents a border of an object.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // Represents a border of an object.
    BorderType string `json:"BorderType,omitempty"`

    // Represents a border of an object.
    Color XmlColorResult `json:"Color,omitempty"`

    // Represents a border of an object.
    DistanceFromText float64 `json:"DistanceFromText,omitempty"`

    // Represents a border of an object.
    LineStyle string `json:"LineStyle,omitempty"`

    // Represents a border of an object.
    LineWidth float64 `json:"LineWidth,omitempty"`

    // Represents a border of an object.
    Shadow bool `json:"Shadow,omitempty"`
}

type Border struct {
    // Represents a border of an object.
    Link IWordsApiLink `json:"Link,omitempty"`

    // Represents a border of an object.
    BorderType *string `json:"BorderType,omitempty"`

    // Represents a border of an object.
    Color IXmlColor `json:"Color,omitempty"`

    // Represents a border of an object.
    DistanceFromText *float64 `json:"DistanceFromText,omitempty"`

    // Represents a border of an object.
    LineStyle *string `json:"LineStyle,omitempty"`

    // Represents a border of an object.
    LineWidth *float64 `json:"LineWidth,omitempty"`

    // Represents a border of an object.
    Shadow *bool `json:"Shadow,omitempty"`
}

type IBorder interface {
    IsBorder() bool
    Initialize()
}

func (Border) IsBorder() bool {
    return true
}

func (Border) IsLinkElement() bool {
    return true
}

func (obj *Border) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }





    if (obj.Color != nil) {
        obj.Color.Initialize()
    }









}


