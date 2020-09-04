/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="list_format.go">
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

// Paragraph list format element.
type ListFormat struct {
    // Paragraph list format element.
    Link *WordsApiLink `json:"Link,omitempty"`

    // Paragraph list format element.
    IsListItem bool `json:"IsListItem,omitempty"`

    // Paragraph list format element.
    ListId int32 `json:"ListId,omitempty"`

    // Paragraph list format element.
    ListLevelNumber int32 `json:"ListLevelNumber,omitempty"`
}

type IListFormat interface {
    IsListFormat() bool
}
func (ListFormat) IsListFormat() bool {
    return true
}

func (ListFormat) IsLinkElement() bool {
    return true
}
