/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="range_document_dto.go">
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

// DTO container with a Range element.

type IRangeDocumentDto interface {
    IsRangeDocumentDto() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetDocumentName() *string
    SetDocumentName(value *string)
}

type RangeDocumentDto struct {
    // DTO container with a Range element.
    DocumentName *string `json:"DocumentName,omitempty"`
}

func (RangeDocumentDto) IsRangeDocumentDto() bool {
    return true
}


func (obj *RangeDocumentDto) Initialize() {
}

func (obj *RangeDocumentDto) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["DocumentName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DocumentName = &parsedValue
        }

    } else if jsonValue, exists := json["documentName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DocumentName = &parsedValue
        }

    }
}

func (obj *RangeDocumentDto) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *RangeDocumentDto) GetDocumentName() *string {
    return obj.DocumentName
}

func (obj *RangeDocumentDto) SetDocumentName(value *string) {
    obj.DocumentName = value
}

