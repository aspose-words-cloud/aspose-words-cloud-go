/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="document_entry_list.go">
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

import (
    "errors"
)

// Represents a list of documents which will be appended to the original resource document.

type IDocumentEntryList interface {
    IsDocumentEntryList() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetApplyBaseDocumentHeadersAndFootersToAppendingDocuments() *bool
    SetApplyBaseDocumentHeadersAndFootersToAppendingDocuments(value *bool)
    GetDocumentEntries() []IDocumentEntry
    SetDocumentEntries(value []IDocumentEntry)
}

type DocumentEntryList struct {
    // Represents a list of documents which will be appended to the original resource document.
    ApplyBaseDocumentHeadersAndFootersToAppendingDocuments *bool `json:"ApplyBaseDocumentHeadersAndFootersToAppendingDocuments,omitempty"`

    // Represents a list of documents which will be appended to the original resource document.
    DocumentEntries []IDocumentEntry `json:"DocumentEntries,omitempty"`
}

func (DocumentEntryList) IsDocumentEntryList() bool {
    return true
}

func (DocumentEntryList) IsBaseEntryList() bool {
    return true
}

func (obj *DocumentEntryList) Initialize() {
    if (obj.DocumentEntries != nil) {
        for _, objElementDocumentEntries := range obj.DocumentEntries {
            objElementDocumentEntries.Initialize()
        }
    }

}

func (obj *DocumentEntryList) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["ApplyBaseDocumentHeadersAndFootersToAppendingDocuments"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ApplyBaseDocumentHeadersAndFootersToAppendingDocuments = &parsedValue
        }

    } else if jsonValue, exists := json["applyBaseDocumentHeadersAndFootersToAppendingDocuments"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ApplyBaseDocumentHeadersAndFootersToAppendingDocuments = &parsedValue
        }

    }

    if jsonValue, exists := json["DocumentEntries"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.DocumentEntries = make([]IDocumentEntry, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IDocumentEntry = new(DocumentEntry)
                    modelElementInstance.Deserialize(elementValue)
                    obj.DocumentEntries = append(obj.DocumentEntries, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["documentEntries"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.DocumentEntries = make([]IDocumentEntry, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IDocumentEntry = new(DocumentEntry)
                    modelElementInstance.Deserialize(elementValue)
                    obj.DocumentEntries = append(obj.DocumentEntries, modelElementInstance)
                }

            }
        }

    }
}

func (obj *DocumentEntryList) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    if (obj.DocumentEntries != nil) {
        for _, element := range obj.DocumentEntries {
            resultFilesContent = element.CollectFilesContent(resultFilesContent)
        }
    }

    return resultFilesContent
}

func (obj *DocumentEntryList) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.DocumentEntries == nil {
        return errors.New("Property DocumentEntries in DocumentEntryList is required.")
    }
    if obj.DocumentEntries != nil {
        for _, elementDocumentEntries := range obj.DocumentEntries {
            if elementDocumentEntries != nil {
                if err := elementDocumentEntries.Validate(); err != nil {
                    return err
                }
            }
        }
    }

    return nil;
}

func (obj *DocumentEntryList) GetApplyBaseDocumentHeadersAndFootersToAppendingDocuments() *bool {
    return obj.ApplyBaseDocumentHeadersAndFootersToAppendingDocuments
}

func (obj *DocumentEntryList) SetApplyBaseDocumentHeadersAndFootersToAppendingDocuments(value *bool) {
    obj.ApplyBaseDocumentHeadersAndFootersToAppendingDocuments = value
}

func (obj *DocumentEntryList) GetDocumentEntries() []IDocumentEntry {
    return obj.DocumentEntries
}

func (obj *DocumentEntryList) SetDocumentEntries(value []IDocumentEntry) {
    obj.DocumentEntries = value
}

