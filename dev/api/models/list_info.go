/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="list_info.go">
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

// DTO container with a single document list.
type ListInfoResult struct {
    // DTO container with a single document list.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // DTO container with a single document list.
    IsListStyleDefinition bool `json:"IsListStyleDefinition,omitempty"`

    // DTO container with a single document list.
    IsListStyleReference bool `json:"IsListStyleReference,omitempty"`

    // DTO container with a single document list.
    IsMultiLevel bool `json:"IsMultiLevel,omitempty"`

    // DTO container with a single document list.
    IsRestartAtEachSection bool `json:"IsRestartAtEachSection,omitempty"`

    // DTO container with a single document list.
    ListId int32 `json:"ListId,omitempty"`

    // DTO container with a single document list.
    ListLevels ListLevelsResult `json:"ListLevels,omitempty"`

    // DTO container with a single document list.
    Style StyleResult `json:"Style,omitempty"`
}

type ListInfo struct {
    // DTO container with a single document list.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with a single document list.
    IsListStyleDefinition *bool `json:"IsListStyleDefinition,omitempty"`

    // DTO container with a single document list.
    IsListStyleReference *bool `json:"IsListStyleReference,omitempty"`

    // DTO container with a single document list.
    IsMultiLevel *bool `json:"IsMultiLevel,omitempty"`

    // DTO container with a single document list.
    IsRestartAtEachSection *bool `json:"IsRestartAtEachSection,omitempty"`

    // DTO container with a single document list.
    ListId *int32 `json:"ListId,omitempty"`

    // DTO container with a single document list.
    ListLevels IListLevels `json:"ListLevels,omitempty"`

    // DTO container with a single document list.
    Style IStyle `json:"Style,omitempty"`
}

type IListInfo interface {
    IsListInfo() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileContent) []FileContent
}

func (ListInfo) IsListInfo() bool {
    return true
}

func (ListInfo) IsLinkElement() bool {
    return true
}

func (obj *ListInfo) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.ListLevels != nil) {
        obj.ListLevels.Initialize()
    }

    if (obj.Style != nil) {
        obj.Style.Initialize()
    }


}

func (obj *ListInfo) CollectFilesContent(resultFilesContent []FileContent) []FileContent {
    if (obj.Link != nil) {
        resultFilesContent = obj.Link.CollectFilesContent(resultFilesContent)
    }






    if (obj.ListLevels != nil) {
        resultFilesContent = obj.ListLevels.CollectFilesContent(resultFilesContent)
    }

    if (obj.Style != nil) {
        resultFilesContent = obj.Style.CollectFilesContent(resultFilesContent)
    }

    return resultFilesContent
}


