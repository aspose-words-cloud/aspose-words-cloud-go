/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="style_insert.go">
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

// Represents a single document style to insert.
type StyleInsertResult struct {
    // Represents a single document style to insert.
    StyleName string `json:"StyleName,omitempty"`

    // Represents a single document style to insert.
    StyleType string `json:"StyleType,omitempty"`
}

type StyleInsert struct {
    // Represents a single document style to insert.
    StyleName *string `json:"StyleName,omitempty"`

    // Represents a single document style to insert.
    StyleType *string `json:"StyleType,omitempty"`
}

type IStyleInsert interface {
    IsStyleInsert() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
}

func (StyleInsert) IsStyleInsert() bool {
    return true
}


func (obj *StyleInsert) Initialize() {
}

func (obj *StyleInsert) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}


