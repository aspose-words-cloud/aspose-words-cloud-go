/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="replace_range_dto.go">
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

// DTO container with a range element.

type IReplaceRangeDto interface {
    IsReplaceRangeDto() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetText() *string
    SetText(value *string)
    GetTextType() *string
    SetTextType(value *string)
}

type ReplaceRangeDto struct {
    // DTO container with a range element.
    Text *string `json:"Text,omitempty"`

    // DTO container with a range element.
    TextType *string `json:"TextType,omitempty"`
}

func (ReplaceRangeDto) IsReplaceRangeDto() bool {
    return true
}


func (obj *ReplaceRangeDto) Initialize() {
}

func (obj *ReplaceRangeDto) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Text"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Text = &parsedValue
        }

    } else if jsonValue, exists := json["text"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Text = &parsedValue
        }

    }

    if jsonValue, exists := json["TextType"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.TextType = &parsedValue
        }

    } else if jsonValue, exists := json["textType"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.TextType = &parsedValue
        }

    }
}

func (obj *ReplaceRangeDto) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *ReplaceRangeDto) GetText() *string {
    return obj.Text
}

func (obj *ReplaceRangeDto) SetText(value *string) {
    obj.Text = value
}

func (obj *ReplaceRangeDto) GetTextType() *string {
    return obj.TextType
}

func (obj *ReplaceRangeDto) SetTextType(value *string) {
    obj.TextType = value
}

