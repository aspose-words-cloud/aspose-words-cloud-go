/*
* MIT License

* Copyright (c) 2019 Aspose Pty Ltd

* Permission is hereby granted, free of charge, to any person obtaining a copy
* of this software and associated documentation files (the "Software"), to deal
* in the Software without restriction, including without limitation the rights
* to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
* copies of the Software, and to permit persons to whom the Software is
* furnished to do so, subject to the following conditions:

* The above copyright notice and this permission notice shall be included in all
* copies or substantial portions of the Software.

* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
* IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
* FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
* AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
* LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
* OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
* SOFTWARE.
*/


package models



// Represents Drawing Object DTO.
type DrawingObject struct {

	Link *WordsApiLink `json:"link,omitempty"`

	// Gets or sets node id.
	NodeId string `json:"NodeId,omitempty"`

	// Gets or sets height of the drawing object in points.
	Height float64 `json:"Height,omitempty"`

	ImageDataLink *WordsApiLink `json:"ImageDataLink,omitempty"`

	// Gets or sets distance in points from the origin to the left side of the image.             
	Left float64 `json:"Left,omitempty"`

	OleDataLink *WordsApiLink `json:"OleDataLink,omitempty"`

	// Gets or sets specifies where the distance to the image is measured from.             
	RelativeHorizontalPosition string `json:"RelativeHorizontalPosition,omitempty"`

	// Gets or sets specifies where the distance to the image measured from.
	RelativeVerticalPosition string `json:"RelativeVerticalPosition,omitempty"`

	// Gets or sets a list of links that originate from this DrawingObjectDto.
	RenderLinks []WordsApiLink `json:"RenderLinks,omitempty"`

	// Gets or sets distance in points from the origin to the top side of the image.
	Top float64 `json:"Top,omitempty"`

	// Gets or sets width of the drawing objects in points.
	Width float64 `json:"Width,omitempty"`

	// Gets or sets specifies how to wrap text around the image.
	WrapType string `json:"WrapType,omitempty"`
}

type IDrawingObject interface {
	IsDrawingObject() bool
}
func (DrawingObject) IsDrawingObject() bool {
	return true;
}
func (DrawingObject) IsDrawingObjectLink() bool {
	return true;
}
