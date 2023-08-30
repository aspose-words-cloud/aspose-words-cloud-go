/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="document.go">
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

// Represents Words document DTO.

type IDocument interface {
    IsDocument() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetLinks() []ILink
    SetLinks(value []ILink)
    GetDocumentProperties() IDocumentProperties
    SetDocumentProperties(value IDocumentProperties)
    GetFileName() *string
    SetFileName(value *string)
    GetIsEncrypted() *bool
    SetIsEncrypted(value *bool)
    GetIsSigned() *bool
    SetIsSigned(value *bool)
    GetSourceFormat() *string
    SetSourceFormat(value *string)
}

type Document struct {
    // Represents Words document DTO.
    Links []ILink `json:"Links,omitempty"`

    // Represents Words document DTO.
    DocumentProperties IDocumentProperties `json:"DocumentProperties,omitempty"`

    // Represents Words document DTO.
    FileName *string `json:"FileName,omitempty"`

    // Represents Words document DTO.
    IsEncrypted *bool `json:"IsEncrypted,omitempty"`

    // Represents Words document DTO.
    IsSigned *bool `json:"IsSigned,omitempty"`

    // Represents Words document DTO.
    SourceFormat *string `json:"SourceFormat,omitempty"`
}

func (Document) IsDocument() bool {
    return true
}


func (obj *Document) Initialize() {
    if (obj.Links != nil) {
        for _, objElementLinks := range obj.Links {
            objElementLinks.Initialize()
        }
    }
    if (obj.DocumentProperties != nil) {
        obj.DocumentProperties.Initialize()
    }


}

func (obj *Document) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Links"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.Links = make([]ILink, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance ILink = nil
                    if jsonElementType, found := elementValue["$type"]; found {
                        jsonTypeStr := jsonElementType.(string)
                        if jsonTypeStr == "FileLink, _" { modelElementInstance = new(FileLink) }
                        if jsonTypeStr == "WordsApiLink, _" { modelElementInstance = new(WordsApiLink) }
                    }

                    if modelElementInstance == nil { modelElementInstance = new(Link) }
                    modelElementInstance.Deserialize(elementValue)
                    obj.Links = append(obj.Links, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["links"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.Links = make([]ILink, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance ILink = nil
                    if jsonElementType, found := elementValue["$type"]; found {
                        jsonTypeStr := jsonElementType.(string)
                        if jsonTypeStr == "FileLink, _" { modelElementInstance = new(FileLink) }
                        if jsonTypeStr == "WordsApiLink, _" { modelElementInstance = new(WordsApiLink) }
                    }

                    if modelElementInstance == nil { modelElementInstance = new(Link) }
                    modelElementInstance.Deserialize(elementValue)
                    obj.Links = append(obj.Links, modelElementInstance)
                }

            }
        }

    }

    if jsonValue, exists := json["DocumentProperties"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IDocumentProperties = new(DocumentProperties)
            modelInstance.Deserialize(parsedValue)
            obj.DocumentProperties = modelInstance
        }

    } else if jsonValue, exists := json["documentProperties"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IDocumentProperties = new(DocumentProperties)
            modelInstance.Deserialize(parsedValue)
            obj.DocumentProperties = modelInstance
        }

    }

    if jsonValue, exists := json["FileName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FileName = &parsedValue
        }

    } else if jsonValue, exists := json["fileName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FileName = &parsedValue
        }

    }

    if jsonValue, exists := json["IsEncrypted"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsEncrypted = &parsedValue
        }

    } else if jsonValue, exists := json["isEncrypted"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsEncrypted = &parsedValue
        }

    }

    if jsonValue, exists := json["IsSigned"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsSigned = &parsedValue
        }

    } else if jsonValue, exists := json["isSigned"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsSigned = &parsedValue
        }

    }

    if jsonValue, exists := json["SourceFormat"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.SourceFormat = &parsedValue
        }

    } else if jsonValue, exists := json["sourceFormat"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.SourceFormat = &parsedValue
        }

    }
}

func (obj *Document) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *Document) GetLinks() []ILink {
    return obj.Links
}

func (obj *Document) SetLinks(value []ILink) {
    obj.Links = value
}

func (obj *Document) GetDocumentProperties() IDocumentProperties {
    return obj.DocumentProperties
}

func (obj *Document) SetDocumentProperties(value IDocumentProperties) {
    obj.DocumentProperties = value
}

func (obj *Document) GetFileName() *string {
    return obj.FileName
}

func (obj *Document) SetFileName(value *string) {
    obj.FileName = value
}

func (obj *Document) GetIsEncrypted() *bool {
    return obj.IsEncrypted
}

func (obj *Document) SetIsEncrypted(value *bool) {
    obj.IsEncrypted = value
}

func (obj *Document) GetIsSigned() *bool {
    return obj.IsSigned
}

func (obj *Document) SetIsSigned(value *bool) {
    obj.IsSigned = value
}

func (obj *Document) GetSourceFormat() *string {
    return obj.SourceFormat
}

func (obj *Document) SetSourceFormat(value *string) {
    obj.SourceFormat = value
}

