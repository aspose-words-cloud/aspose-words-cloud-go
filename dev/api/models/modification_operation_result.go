/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="modification_operation_result.go">
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

// result of the operation which modifies the original document and saves the result.
type ModificationOperationResultResult struct {
    // result of the operation which modifies the original document and saves the result.
    Dest FileLinkResult `json:"Dest,omitempty"`

    // result of the operation which modifies the original document and saves the result.
    Source FileLinkResult `json:"Source,omitempty"`
}

type ModificationOperationResult struct {
    // result of the operation which modifies the original document and saves the result.
    Dest IFileLink `json:"Dest,omitempty"`

    // result of the operation which modifies the original document and saves the result.
    Source IFileLink `json:"Source,omitempty"`
}

type IModificationOperationResult interface {
    IsModificationOperationResult() bool
}
func (ModificationOperationResult) IsModificationOperationResult() bool {
    return true
}


