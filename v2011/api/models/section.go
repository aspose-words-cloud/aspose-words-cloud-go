/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="section.go">
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

// DTO container with a section element.
type Section struct {
    // DTO container with a section element.
    Link *WordsApiLink `json:"Link,omitempty"`

    // DTO container with a section element.
    ChildNodes []NodeLink `json:"ChildNodes,omitempty"`

    // DTO container with a section element.
    HeaderFooters *LinkElement `json:"HeaderFooters,omitempty"`

    // DTO container with a section element.
    PageSetup *LinkElement `json:"PageSetup,omitempty"`

    // DTO container with a section element.
    Paragraphs *LinkElement `json:"Paragraphs,omitempty"`

    // DTO container with a section element.
    Tables *LinkElement `json:"Tables,omitempty"`
}

type ISection interface {
    IsSection() bool
}
func (Section) IsSection() bool {
    return true
}

func (Section) IsLinkElement() bool {
    return true
}
