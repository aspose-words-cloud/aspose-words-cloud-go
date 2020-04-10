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



// Represents a single document list.
type ListInfo struct {

	// Gets or sets the unique identifier of the list.
	ListId int32 `json:"ListId,omitempty"`

	// Gets or sets a value indicating whether returns true when the list contains 9 levels; false when 1 level.
	IsMultiLevel bool `json:"IsMultiLevel,omitempty"`

	// Gets or sets a value indicating whether specifies whether list should be restarted at each section. Default value is false.
	IsRestartAtEachSection bool `json:"IsRestartAtEachSection,omitempty"`

	// Gets or sets a value indicating whether returns true if this list is a definition of a list style.
	IsListStyleDefinition bool `json:"IsListStyleDefinition,omitempty"`

	// Gets or sets a value indicating whether returns true if this list is a reference to a list style.
	IsListStyleReference bool `json:"IsListStyleReference,omitempty"`

	Style *Style `json:"Style,omitempty"`

	ListLevels *ListLevels `json:"ListLevels,omitempty"`
}

type IListInfo interface {
	IsListInfo() bool
}
func (ListInfo) IsListInfo() bool {
	return true;
}
