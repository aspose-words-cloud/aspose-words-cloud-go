/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="list_level.go">
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

// DTO container with a document list level.
type ListLevelResult struct {
    // DTO container with a document list level.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // DTO container with a document list level.
    Alignment string `json:"Alignment,omitempty"`

    // DTO container with a document list level.
    Font FontResult `json:"Font,omitempty"`

    // DTO container with a document list level.
    IsLegal bool `json:"IsLegal,omitempty"`

    // DTO container with a document list level.
    LinkedStyle StyleResult `json:"LinkedStyle,omitempty"`

    // DTO container with a document list level.
    NumberFormat string `json:"NumberFormat,omitempty"`

    // DTO container with a document list level.
    NumberPosition float64 `json:"NumberPosition,omitempty"`

    // DTO container with a document list level.
    NumberStyle string `json:"NumberStyle,omitempty"`

    // DTO container with a document list level.
    RestartAfterLevel int32 `json:"RestartAfterLevel,omitempty"`

    // DTO container with a document list level.
    StartAt int32 `json:"StartAt,omitempty"`

    // DTO container with a document list level.
    TabPosition float64 `json:"TabPosition,omitempty"`

    // DTO container with a document list level.
    TextPosition float64 `json:"TextPosition,omitempty"`

    // DTO container with a document list level.
    TrailingCharacter string `json:"TrailingCharacter,omitempty"`
}

type ListLevel struct {
    // DTO container with a document list level.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with a document list level.
    Alignment *string `json:"Alignment,omitempty"`

    // DTO container with a document list level.
    Font IFont `json:"Font,omitempty"`

    // DTO container with a document list level.
    IsLegal *bool `json:"IsLegal,omitempty"`

    // DTO container with a document list level.
    LinkedStyle IStyle `json:"LinkedStyle,omitempty"`

    // DTO container with a document list level.
    NumberFormat *string `json:"NumberFormat,omitempty"`

    // DTO container with a document list level.
    NumberPosition *float64 `json:"NumberPosition,omitempty"`

    // DTO container with a document list level.
    NumberStyle *string `json:"NumberStyle,omitempty"`

    // DTO container with a document list level.
    RestartAfterLevel *int32 `json:"RestartAfterLevel,omitempty"`

    // DTO container with a document list level.
    StartAt *int32 `json:"StartAt,omitempty"`

    // DTO container with a document list level.
    TabPosition *float64 `json:"TabPosition,omitempty"`

    // DTO container with a document list level.
    TextPosition *float64 `json:"TextPosition,omitempty"`

    // DTO container with a document list level.
    TrailingCharacter *string `json:"TrailingCharacter,omitempty"`
}

type IListLevel interface {
    IsListLevel() bool
    Initialize()
}

func (ListLevel) IsListLevel() bool {
    return true
}

func (ListLevel) IsLinkElement() bool {
    return true
}

func (obj *ListLevel) Initialize() {
}


