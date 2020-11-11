/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="list_info.go">
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

// Represents a single document list.
type ListInfo struct {
    // Represents a single document list.
    Link *WordsApiLink `json:"Link,omitempty"`

    // Represents a single document list.
    IsListStyleDefinition *bool `json:"IsListStyleDefinition,omitempty"`

    // Represents a single document list.
    IsListStyleReference *bool `json:"IsListStyleReference,omitempty"`

    // Represents a single document list.
    IsMultiLevel *bool `json:"IsMultiLevel,omitempty"`

    // Represents a single document list.
    IsRestartAtEachSection *bool `json:"IsRestartAtEachSection,omitempty"`

    // Represents a single document list.
    ListId *int32 `json:"ListId,omitempty"`

    // Represents a single document list.
    ListLevels *ListLevels `json:"ListLevels,omitempty"`

    // Represents a single document list.
    Style *Style `json:"Style,omitempty"`
}

type IListInfo interface {
    IsListInfo() bool
}
func (ListInfo) IsListInfo() bool {
    return true
}

func (ListInfo) IsLinkElement() bool {
    return true
}
