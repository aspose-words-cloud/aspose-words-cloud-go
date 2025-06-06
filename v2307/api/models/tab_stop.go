/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="tab_stop.go">
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

// DTO container with paragraph format tab stop.
type TabStopResult struct {
    // DTO container with paragraph format tab stop.
    Alignment string `json:"Alignment,omitempty"`

    // DTO container with paragraph format tab stop.
    Leader string `json:"Leader,omitempty"`

    // DTO container with paragraph format tab stop.
    Position float64 `json:"Position,omitempty"`

    // DTO container with paragraph format tab stop.
    IsClear bool `json:"IsClear,omitempty"`
}

type TabStop struct {
    // DTO container with paragraph format tab stop.
    Alignment *string `json:"Alignment,omitempty"`

    // DTO container with paragraph format tab stop.
    Leader *string `json:"Leader,omitempty"`

    // DTO container with paragraph format tab stop.
    Position *float64 `json:"Position,omitempty"`

    // DTO container with paragraph format tab stop.
    IsClear *bool `json:"IsClear,omitempty"`
}

type ITabStop interface {
    IsTabStop() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
}

func (TabStop) IsTabStop() bool {
    return true
}

func (TabStop) IsTabStopBase() bool {
    return true
}

func (obj *TabStop) Initialize() {
}

func (obj *TabStop) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}


