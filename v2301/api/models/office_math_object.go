/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="office_math_object.go">
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

// DTO container with an OfficeMath object.
type OfficeMathObjectResult struct {
    // DTO container with an OfficeMath object.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // DTO container with an OfficeMath object.
    NodeId string `json:"NodeId,omitempty"`

    // DTO container with an OfficeMath object.
    Content StoryChildNodesResult `json:"Content,omitempty"`

    // DTO container with an OfficeMath object.
    DisplayType string `json:"DisplayType,omitempty"`

    // DTO container with an OfficeMath object.
    Justification string `json:"Justification,omitempty"`

    // DTO container with an OfficeMath object.
    MathObjectType string `json:"MathObjectType,omitempty"`
}

type OfficeMathObject struct {
    // DTO container with an OfficeMath object.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with an OfficeMath object.
    NodeId *string `json:"NodeId,omitempty"`

    // DTO container with an OfficeMath object.
    Content IStoryChildNodes `json:"Content,omitempty"`

    // DTO container with an OfficeMath object.
    DisplayType *string `json:"DisplayType,omitempty"`

    // DTO container with an OfficeMath object.
    Justification *string `json:"Justification,omitempty"`

    // DTO container with an OfficeMath object.
    MathObjectType *string `json:"MathObjectType,omitempty"`
}

type IOfficeMathObject interface {
    IsOfficeMathObject() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
}

func (OfficeMathObject) IsOfficeMathObject() bool {
    return true
}

func (OfficeMathObject) IsOfficeMathLink() bool {
    return true
}

func (OfficeMathObject) IsNodeLink() bool {
    return true
}

func (OfficeMathObject) IsLinkElement() bool {
    return true
}

func (obj *OfficeMathObject) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.Content != nil) {
        obj.Content.Initialize()
    }


}

func (obj *OfficeMathObject) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}


