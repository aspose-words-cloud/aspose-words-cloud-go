/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="image_entry_list.go">
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

// Represents a list of images which will be appended to the original resource document or image.
type ImageEntryListResult struct {
    // Represents a list of images which will be appended to the original resource document or image.
    AppendEachImageOnNewPage bool `json:"AppendEachImageOnNewPage,omitempty"`

    // Represents a list of images which will be appended to the original resource document or image.
    ImageEntries []ImageEntryResult `json:"ImageEntries,omitempty"`
}

type ImageEntryList struct {
    // Represents a list of images which will be appended to the original resource document or image.
    AppendEachImageOnNewPage *bool `json:"AppendEachImageOnNewPage,omitempty"`

    // Represents a list of images which will be appended to the original resource document or image.
    ImageEntries []ImageEntry `json:"ImageEntries,omitempty"`
}

type IImageEntryList interface {
    IsImageEntryList() bool
    Initialize()
}

func (ImageEntryList) IsImageEntryList() bool {
    return true
}

func (ImageEntryList) IsBaseEntryList() bool {
    return true
}

func (obj *ImageEntryList) Initialize() {
}


