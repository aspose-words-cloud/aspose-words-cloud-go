/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="image_entry_list.go">
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

// Represents a list of images which will be appended to the original resource document or image.

type IImageEntryList interface {
    IsImageEntryList() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetAppendEachImageOnNewPage() *bool
    SetAppendEachImageOnNewPage(value *bool)
    GetImageEntries() []IImageEntry
    SetImageEntries(value []IImageEntry)
}

type ImageEntryList struct {
    // Represents a list of images which will be appended to the original resource document or image.
    AppendEachImageOnNewPage *bool `json:"AppendEachImageOnNewPage,omitempty"`

    // Represents a list of images which will be appended to the original resource document or image.
    ImageEntries []IImageEntry `json:"ImageEntries,omitempty"`
}

func (ImageEntryList) IsImageEntryList() bool {
    return true
}

func (ImageEntryList) IsBaseEntryList() bool {
    return true
}

func (obj *ImageEntryList) Initialize() {
    if (obj.ImageEntries != nil) {
        for _, objElementImageEntries := range obj.ImageEntries {
            objElementImageEntries.Initialize()
        }
    }

}

func (obj *ImageEntryList) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["AppendEachImageOnNewPage"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.AppendEachImageOnNewPage = &parsedValue
        }

    } else if jsonValue, exists := json["appendEachImageOnNewPage"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.AppendEachImageOnNewPage = &parsedValue
        }

    }

    if jsonValue, exists := json["ImageEntries"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.ImageEntries = make([]IImageEntry, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IImageEntry = new(ImageEntry)
                    modelElementInstance.Deserialize(elementValue)
                    obj.ImageEntries = append(obj.ImageEntries, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["imageEntries"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.ImageEntries = make([]IImageEntry, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IImageEntry = new(ImageEntry)
                    modelElementInstance.Deserialize(elementValue)
                    obj.ImageEntries = append(obj.ImageEntries, modelElementInstance)
                }

            }
        }

    }
}

func (obj *ImageEntryList) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    if (obj.ImageEntries != nil) {
        for _, element := range obj.ImageEntries {
            resultFilesContent = element.CollectFilesContent(resultFilesContent)
        }
    }

    return resultFilesContent
}

func (obj *ImageEntryList) GetAppendEachImageOnNewPage() *bool {
    return obj.AppendEachImageOnNewPage
}

func (obj *ImageEntryList) SetAppendEachImageOnNewPage(value *bool) {
    obj.AppendEachImageOnNewPage = value
}

func (obj *ImageEntryList) GetImageEntries() []IImageEntry {
    return obj.ImageEntries
}

func (obj *ImageEntryList) SetImageEntries(value []IImageEntry) {
    obj.ImageEntries = value
}

