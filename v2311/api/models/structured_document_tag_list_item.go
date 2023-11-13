/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="structured_document_tag_list_item.go">
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

// DTO container with a SdtListItem for StructuredDocumentTag.

type IStructuredDocumentTagListItem interface {
    IsStructuredDocumentTagListItem() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetDisplayText() *string
    SetDisplayText(value *string)
    GetValue() *string
    SetValue(value *string)
}

type StructuredDocumentTagListItem struct {
    // DTO container with a SdtListItem for StructuredDocumentTag.
    DisplayText *string `json:"DisplayText,omitempty"`

    // DTO container with a SdtListItem for StructuredDocumentTag.
    Value *string `json:"Value,omitempty"`
}

func (StructuredDocumentTagListItem) IsStructuredDocumentTagListItem() bool {
    return true
}


func (obj *StructuredDocumentTagListItem) Initialize() {
}

func (obj *StructuredDocumentTagListItem) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["DisplayText"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DisplayText = &parsedValue
        }

    } else if jsonValue, exists := json["displayText"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DisplayText = &parsedValue
        }

    }

    if jsonValue, exists := json["Value"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Value = &parsedValue
        }

    } else if jsonValue, exists := json["value"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Value = &parsedValue
        }

    }
}

func (obj *StructuredDocumentTagListItem) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *StructuredDocumentTagListItem) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    return nil;
}

func (obj *StructuredDocumentTagListItem) GetDisplayText() *string {
    return obj.DisplayText
}

func (obj *StructuredDocumentTagListItem) SetDisplayText(value *string) {
    obj.DisplayText = value
}

func (obj *StructuredDocumentTagListItem) GetValue() *string {
    return obj.Value
}

func (obj *StructuredDocumentTagListItem) SetValue(value *string) {
    obj.Value = value
}

