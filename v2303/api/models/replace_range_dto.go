/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="replace_range_dto.go">
 *   Copyright (c) 2023 Aspose.Words for Cloud
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

// DTO container with a range element.
type ReplaceRangeDtoResult struct {
    // DTO container with a range element.
    Text string `json:"Text,omitempty"`

    // DTO container with a range element.
    TextType string `json:"TextType,omitempty"`
}

type ReplaceRangeDto struct {
    // DTO container with a range element.
    Text *string `json:"Text,omitempty"`

    // DTO container with a range element.
    TextType *string `json:"TextType,omitempty"`
}

type IReplaceRangeDto interface {
    IsReplaceRangeDto() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
}

func (ReplaceRangeDto) IsReplaceRangeDto() bool {
    return true
}


func (obj *ReplaceRangeDto) Initialize() {
}

func (obj *ReplaceRangeDto) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}


