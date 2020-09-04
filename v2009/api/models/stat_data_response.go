/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="stat_data_response.go">
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

// Response for the request of the document's statistical data.
type StatDataResponse struct {
    // Response for the request of the document's statistical data.
    RequestId string `json:"RequestId,omitempty"`

    // Response for the request of the document's statistical data.
    DocumentLink *FileLink `json:"DocumentLink,omitempty"`

    // Response for the request of the document's statistical data.
    StatData *DocumentStatData `json:"StatData,omitempty"`
}

type IStatDataResponse interface {
    IsStatDataResponse() bool
}
func (StatDataResponse) IsStatDataResponse() bool {
    return true
}

func (StatDataResponse) IsWordsResponse() bool {
    return true
}
