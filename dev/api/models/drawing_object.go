/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="drawing_object.go">
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

// DTO container with a DrawingObject.
type DrawingObjectResult struct {
    // DTO container with a DrawingObject.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // DTO container with a DrawingObject.
    NodeId string `json:"NodeId,omitempty"`

    // DTO container with a DrawingObject.
    Height float64 `json:"Height,omitempty"`

    // DTO container with a DrawingObject.
    ImageDataLink WordsApiLinkResult `json:"ImageDataLink,omitempty"`

    // DTO container with a DrawingObject.
    Left float64 `json:"Left,omitempty"`

    // DTO container with a DrawingObject.
    OleDataLink WordsApiLinkResult `json:"OleDataLink,omitempty"`

    // DTO container with a DrawingObject.
    RelativeHorizontalPosition string `json:"RelativeHorizontalPosition,omitempty"`

    // DTO container with a DrawingObject.
    RelativeVerticalPosition string `json:"RelativeVerticalPosition,omitempty"`

    // DTO container with a DrawingObject.
    RenderLinks []WordsApiLinkResult `json:"RenderLinks,omitempty"`

    // DTO container with a DrawingObject.
    Top float64 `json:"Top,omitempty"`

    // DTO container with a DrawingObject.
    Width float64 `json:"Width,omitempty"`

    // DTO container with a DrawingObject.
    WrapType string `json:"WrapType,omitempty"`
}

type DrawingObject struct {
    // DTO container with a DrawingObject.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with a DrawingObject.
    NodeId *string `json:"NodeId,omitempty"`

    // DTO container with a DrawingObject.
    Height *float64 `json:"Height,omitempty"`

    // DTO container with a DrawingObject.
    ImageDataLink IWordsApiLink `json:"ImageDataLink,omitempty"`

    // DTO container with a DrawingObject.
    Left *float64 `json:"Left,omitempty"`

    // DTO container with a DrawingObject.
    OleDataLink IWordsApiLink `json:"OleDataLink,omitempty"`

    // DTO container with a DrawingObject.
    RelativeHorizontalPosition *string `json:"RelativeHorizontalPosition,omitempty"`

    // DTO container with a DrawingObject.
    RelativeVerticalPosition *string `json:"RelativeVerticalPosition,omitempty"`

    // DTO container with a DrawingObject.
    RenderLinks []WordsApiLink `json:"RenderLinks,omitempty"`

    // DTO container with a DrawingObject.
    Top *float64 `json:"Top,omitempty"`

    // DTO container with a DrawingObject.
    Width *float64 `json:"Width,omitempty"`

    // DTO container with a DrawingObject.
    WrapType *string `json:"WrapType,omitempty"`
}

type IDrawingObject interface {
    IsDrawingObject() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileContent) []FileContent
}

func (DrawingObject) IsDrawingObject() bool {
    return true
}

func (DrawingObject) IsDrawingObjectLink() bool {
    return true
}

func (DrawingObject) IsNodeLink() bool {
    return true
}

func (DrawingObject) IsLinkElement() bool {
    return true
}

func (obj *DrawingObject) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.ImageDataLink != nil) {
        obj.ImageDataLink.Initialize()
    }

    if (obj.OleDataLink != nil) {
        obj.OleDataLink.Initialize()
    }

    if (obj.RenderLinks != nil) {
        for _, element := range obj.RenderLinks {
            element.Initialize()
        }
    }


}

func (obj *DrawingObject) CollectFilesContent(resultFilesContent []FileContent) []FileContent {
    if (obj.ImageDataLink != nil) {
        resultFilesContent = obj.ImageDataLink.CollectFilesContent(resultFilesContent)
    }


    if (obj.OleDataLink != nil) {
        resultFilesContent = obj.OleDataLink.CollectFilesContent(resultFilesContent)
    }



    if (obj.RenderLinks != nil) {
        for _, element := range obj.RenderLinks {
            resultFilesContent = element.CollectFilesContent(resultFilesContent)
        }
    }




    return resultFilesContent
}


