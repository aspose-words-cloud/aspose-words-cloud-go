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



// OfficeMath object.
type OfficeMathObject struct {

	Content *StoryChildNodes `json:"Content,omitempty"`

	// Gets or sets /sets Office Math display format type which represents whether an equation is displayed inline with the text or displayed on its own line.
	DisplayType string `json:"DisplayType,omitempty"`

	// Gets or sets /sets Office Math justification.
	Justification string `json:"Justification,omitempty"`

	// Gets or sets type Aspose.Words.Math.OfficeMath.MathObjectType of this Office Math object.
	MathObjectType string `json:"MathObjectType,omitempty"`
}

type IOfficeMathObject interface {
	IsOfficeMathObject() bool
}
func (OfficeMathObject) IsOfficeMathObject() bool {
	return true;
}
func (OfficeMathObject) IsOfficeMathLink() bool {
	return true;
}
