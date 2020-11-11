/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="list_level_update.go">
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

// Represents a document list levels.
type ListLevelUpdate struct {
    // Represents a document list levels.
    Alignment string `json:"Alignment,omitempty"`

    // Represents a document list levels.
    IsLegal bool `json:"IsLegal,omitempty"`

    // Represents a document list levels.
    NumberFormat string `json:"NumberFormat,omitempty"`

    // Represents a document list levels.
    NumberPosition float64 `json:"NumberPosition,omitempty"`

    // Represents a document list levels.
    NumberStyle string `json:"NumberStyle,omitempty"`

    // Represents a document list levels.
    RestartAfterLevel int32 `json:"RestartAfterLevel,omitempty"`

    // Represents a document list levels.
    StartAt int32 `json:"StartAt,omitempty"`

    // Represents a document list levels.
    TabPosition float64 `json:"TabPosition,omitempty"`

    // Represents a document list levels.
    TextPosition float64 `json:"TextPosition,omitempty"`

    // Represents a document list levels.
    TrailingCharacter string `json:"TrailingCharacter,omitempty"`
}

type IListLevelUpdate interface {
    IsListLevelUpdate() bool
}
func (ListLevelUpdate) IsListLevelUpdate() bool {
    return true
}

