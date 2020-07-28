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



// container class for outline options.
type OutlineOptionsData struct {

	// Gets or sets allows to specify individual bookmarks outline level.
	BookmarksOutlineLevels []BookmarksOutlineLevelData `json:"BookmarksOutlineLevels,omitempty"`

	// Gets or sets specifies the default level in the document outline at which to display Word bookmarks.
	DefaultBookmarksOutlineLevel int32 `json:"DefaultBookmarksOutlineLevel,omitempty"`

	// Gets or sets a value determining whether or not to create missing outline levels     when the document is exported.     Default value for this property is false.
	CreateMissingOutlineLevels bool `json:"CreateMissingOutlineLevels,omitempty"`

	// Gets or sets specifies whether or not to create outlines for headings (paragraphs formatted     with the Heading styles) inside tables.
	CreateOutlinesForHeadingsInTables bool `json:"CreateOutlinesForHeadingsInTables,omitempty"`

	// Gets or sets specifies how many levels in the document outline to show expanded when the file is viewed.
	ExpandedOutlineLevels int32 `json:"ExpandedOutlineLevels,omitempty"`

	// Gets or sets specifies how many levels of headings (paragraphs formatted with the Heading styles) to include in the document outline.
	HeadingsOutlineLevels int32 `json:"HeadingsOutlineLevels,omitempty"`
}

type IOutlineOptionsData interface {
	IsOutlineOptionsData() bool
}
func (OutlineOptionsData) IsOutlineOptionsData() bool {
	return true;
}
