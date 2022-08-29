/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="drawing_object_collection.go">
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

// DTO container with a collection of DrawingObjects links.
type DrawingObjectCollectionResult struct {
    // DTO container with a collection of DrawingObjects links.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // DTO container with a collection of DrawingObjects links.
    List []LinkElementResult `json:"List,omitempty"`
}

type DrawingObjectCollection struct {
    // DTO container with a collection of DrawingObjects links.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with a collection of DrawingObjects links.
    List []LinkElement `json:"List,omitempty"`
}

type IDrawingObjectCollection interface {
    IsDrawingObjectCollection() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileContent) []FileContent
}

func (DrawingObjectCollection) IsDrawingObjectCollection() bool {
    return true
}

func (DrawingObjectCollection) IsLinkElement() bool {
    return true
}

func (obj *DrawingObjectCollection) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.List != nil) {
        for _, element := range obj.List {
            element.Initialize()
        }
    }


}

func (obj *DrawingObjectCollection) CollectFilesContent(resultFilesContent []FileContent) []FileContent {
    if (obj.List != nil) {
        for _, element := range obj.List {
            resultFilesContent = element.CollectFilesContent(resultFilesContent)
        }
    }

    return resultFilesContent
}


