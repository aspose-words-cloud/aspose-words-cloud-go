/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="drawing_object_insert.go">
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

// Drawing object element for insert.
type DrawingObjectInsertResult struct {
    // Drawing object element for insert.
    Height float64 `json:"Height,omitempty"`

    // Drawing object element for insert.
    Left float64 `json:"Left,omitempty"`

    // Drawing object element for insert.
    Position DocumentPositionResult `json:"Position,omitempty"`

    // Drawing object element for insert.
    RelativeHorizontalPosition string `json:"RelativeHorizontalPosition,omitempty"`

    // Drawing object element for insert.
    RelativeVerticalPosition string `json:"RelativeVerticalPosition,omitempty"`

    // Drawing object element for insert.
    Top float64 `json:"Top,omitempty"`

    // Drawing object element for insert.
    Width float64 `json:"Width,omitempty"`

    // Drawing object element for insert.
    WrapType string `json:"WrapType,omitempty"`
}

type DrawingObjectInsert struct {
    // Drawing object element for insert.
    Height *float64 `json:"Height,omitempty"`

    // Drawing object element for insert.
    Left *float64 `json:"Left,omitempty"`

    // Drawing object element for insert.
    Position IDocumentPosition `json:"Position,omitempty"`

    // Drawing object element for insert.
    RelativeHorizontalPosition *string `json:"RelativeHorizontalPosition,omitempty"`

    // Drawing object element for insert.
    RelativeVerticalPosition *string `json:"RelativeVerticalPosition,omitempty"`

    // Drawing object element for insert.
    Top *float64 `json:"Top,omitempty"`

    // Drawing object element for insert.
    Width *float64 `json:"Width,omitempty"`

    // Drawing object element for insert.
    WrapType *string `json:"WrapType,omitempty"`
}

type IDrawingObjectInsert interface {
    IsDrawingObjectInsert() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileContent) []FileContent
}

func (DrawingObjectInsert) IsDrawingObjectInsert() bool {
    return true
}


func (obj *DrawingObjectInsert) Initialize() {
    if (obj.Position != nil) {
        obj.Position.Initialize()
    }


}

func (obj *DrawingObjectInsert) CollectFilesContent(resultFilesContent []FileContent) []FileContent {
    if (obj.Position != nil) {
        resultFilesContent = obj.Position.CollectFilesContent(resultFilesContent)
    }






    return resultFilesContent
}


