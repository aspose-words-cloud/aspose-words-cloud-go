/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="office_math_link.go">
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

// OfficeMath object link element.
type OfficeMathLinkResult struct {
    // OfficeMath object link element.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // OfficeMath object link element.
    NodeId string `json:"NodeId,omitempty"`
}

type OfficeMathLink struct {
    // OfficeMath object link element.
    Link IWordsApiLink `json:"Link,omitempty"`

    // OfficeMath object link element.
    NodeId *string `json:"NodeId,omitempty"`
}

type IOfficeMathLink interface {
    IsOfficeMathLink() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
}

func (OfficeMathLink) IsOfficeMathLink() bool {
    return true
}

func (OfficeMathLink) IsNodeLink() bool {
    return true
}

func (OfficeMathLink) IsLinkElement() bool {
    return true
}

func (obj *OfficeMathLink) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }


}

func (obj *OfficeMathLink) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}


