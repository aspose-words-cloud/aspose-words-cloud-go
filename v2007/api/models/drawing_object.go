/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="drawing_object.go">
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

// Represents Drawing Object DTO.
type DrawingObject struct {
    // Represents Drawing Object DTO.
    Link *WordsApiLink `json:"Link,omitempty"`

    // Represents Drawing Object DTO.
    NodeId string `json:"NodeId,omitempty"`

    // Represents Drawing Object DTO.
    Height float64 `json:"Height,omitempty"`

    // Represents Drawing Object DTO.
    ImageDataLink *WordsApiLink `json:"ImageDataLink,omitempty"`

    // Represents Drawing Object DTO.
    Left float64 `json:"Left,omitempty"`

    // Represents Drawing Object DTO.
    OleDataLink *WordsApiLink `json:"OleDataLink,omitempty"`

    // Represents Drawing Object DTO.
    RelativeHorizontalPosition string `json:"RelativeHorizontalPosition,omitempty"`

    // Represents Drawing Object DTO.
    RelativeVerticalPosition string `json:"RelativeVerticalPosition,omitempty"`

    // Represents Drawing Object DTO.
    RenderLinks []WordsApiLink `json:"RenderLinks,omitempty"`

    // Represents Drawing Object DTO.
    Top float64 `json:"Top,omitempty"`

    // Represents Drawing Object DTO.
    Width float64 `json:"Width,omitempty"`

    // Represents Drawing Object DTO.
    WrapType string `json:"WrapType,omitempty"`
}

type IDrawingObject interface {
    IsDrawingObject() bool
}
func (DrawingObject) IsDrawingObject() bool {
    return true
}

func (DrawingObject) IsDrawingObjectLink() bool {
    return true
}
