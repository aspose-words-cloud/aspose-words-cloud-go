/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="outline_options_data.go">
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

// container class for outline options.
type OutlineOptionsData struct {
    // container class for outline options.
    BookmarksOutlineLevels []BookmarksOutlineLevelData `json:"BookmarksOutlineLevels,omitempty"`

    // container class for outline options.
    CreateMissingOutlineLevels *bool `json:"CreateMissingOutlineLevels,omitempty"`

    // container class for outline options.
    CreateOutlinesForHeadingsInTables *bool `json:"CreateOutlinesForHeadingsInTables,omitempty"`

    // container class for outline options.
    DefaultBookmarksOutlineLevel *int32 `json:"DefaultBookmarksOutlineLevel,omitempty"`

    // container class for outline options.
    ExpandedOutlineLevels *int32 `json:"ExpandedOutlineLevels,omitempty"`

    // container class for outline options.
    HeadingsOutlineLevels *int32 `json:"HeadingsOutlineLevels,omitempty"`
}

type IOutlineOptionsData interface {
    IsOutlineOptionsData() bool
}
func (OutlineOptionsData) IsOutlineOptionsData() bool {
    return true
}

