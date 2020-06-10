/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="header_footer.go">
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

// Section element.
type HeaderFooter struct {
    // Section element.
    Link *WordsApiLink `json:"Link,omitempty"`

    // Section element.
    Type string `json:"Type,omitempty"`

    // Section element.
    ChildNodes []NodeLink `json:"ChildNodes,omitempty"`

    // Section element.
    DrawingObjects *LinkElement `json:"DrawingObjects,omitempty"`

    // Section element.
    Paragraphs *LinkElement `json:"Paragraphs,omitempty"`
}

type IHeaderFooter interface {
    IsHeaderFooter() bool
}
func (HeaderFooter) IsHeaderFooter() bool {
    return true
}

func (HeaderFooter) IsHeaderFooterLink() bool {
    return true
}
