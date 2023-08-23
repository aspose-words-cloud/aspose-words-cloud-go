/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="split_document_result.go">
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

// Result of splitting document.

type ISplitDocumentResult interface {
    IsSplitDocumentResult() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetSourceDocument() IFileLink
    SetSourceDocument(value IFileLink)
    GetZippedPages() IFileLink
    SetZippedPages(value IFileLink)
    GetPages() []IFileLink
    SetPages(value []IFileLink)
}

type SplitDocumentResult struct {
    // Result of splitting document.
    SourceDocument IFileLink

    // Result of splitting document.
    ZippedPages IFileLink

    // Result of splitting document.
    Pages []IFileLink
}

func (SplitDocumentResult) IsSplitDocumentResult() bool {
    return true
}


func (obj *SplitDocumentResult) Initialize() {
    if (obj.SourceDocument != nil) {
        obj.SourceDocument.Initialize()
    }

    if (obj.ZippedPages != nil) {
        obj.ZippedPages.Initialize()
    }

    if (obj.Pages != nil) {
        for _, objElementPages := range obj.Pages {
            objElementPages.Initialize()
        }
    }

}

func (obj *SplitDocumentResult) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["SourceDocument"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFileLink = new(FileLink)
            modelInstance.Deserialize(parsedValue)
            obj.SourceDocument = modelInstance
        }

    } else if jsonValue, exists := json["sourceDocument"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFileLink = new(FileLink)
            modelInstance.Deserialize(parsedValue)
            obj.SourceDocument = modelInstance
        }

    }

    if jsonValue, exists := json["ZippedPages"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFileLink = new(FileLink)
            modelInstance.Deserialize(parsedValue)
            obj.ZippedPages = modelInstance
        }

    } else if jsonValue, exists := json["zippedPages"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFileLink = new(FileLink)
            modelInstance.Deserialize(parsedValue)
            obj.ZippedPages = modelInstance
        }

    }

    if jsonValue, exists := json["Pages"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IFileLink = new(FileLink)
                    modelElementInstance.Deserialize(elementValue)
                    obj.Pages = append(obj.Pages, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["pages"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IFileLink = new(FileLink)
                    modelElementInstance.Deserialize(elementValue)
                    obj.Pages = append(obj.Pages, modelElementInstance)
                }

            }
        }

    }
}

func (obj *SplitDocumentResult) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *SplitDocumentResult) GetSourceDocument() IFileLink {
    return obj.SourceDocument
}

func (obj *SplitDocumentResult) SetSourceDocument(value IFileLink) {
    obj.SourceDocument = value
}

func (obj *SplitDocumentResult) GetZippedPages() IFileLink {
    return obj.ZippedPages
}

func (obj *SplitDocumentResult) SetZippedPages(value IFileLink) {
    obj.ZippedPages = value
}

func (obj *SplitDocumentResult) GetPages() []IFileLink {
    return obj.Pages
}

func (obj *SplitDocumentResult) SetPages(value []IFileLink) {
    obj.Pages = value
}

