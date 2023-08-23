/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="storage_file.go">
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

// File or folder information.

type IStorageFile interface {
    IsStorageFile() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetIsFolder() *bool
    SetIsFolder(value *bool)
    GetModifiedDate() *Time
    SetModifiedDate(value *Time)
    GetName() *string
    SetName(value *string)
    GetPath() *string
    SetPath(value *string)
    GetSize() *int32
    SetSize(value *int32)
}

type StorageFile struct {
    // File or folder information.
    IsFolder *bool

    // File or folder information.
    ModifiedDate *Time

    // File or folder information.
    Name *string

    // File or folder information.
    Path *string

    // File or folder information.
    Size *int32
}

func (StorageFile) IsStorageFile() bool {
    return true
}


func (obj *StorageFile) Initialize() {
}

func (obj *StorageFile) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["IsFolder"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsFolder = &parsedValue
        }

    } else if jsonValue, exists := json["isFolder"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsFolder = &parsedValue
        }

    }

    if jsonValue, exists := json["ModifiedDate"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ModifiedDate = new(Time)
            obj.ModifiedDate.Parse(parsedValue)
        }

    } else if jsonValue, exists := json["modifiedDate"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ModifiedDate = new(Time)
            obj.ModifiedDate.Parse(parsedValue)
        }

    }

    if jsonValue, exists := json["Name"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Name = &parsedValue
        }

    } else if jsonValue, exists := json["name"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Name = &parsedValue
        }

    }

    if jsonValue, exists := json["Path"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Path = &parsedValue
        }

    } else if jsonValue, exists := json["path"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Path = &parsedValue
        }

    }

    if jsonValue, exists := json["Size"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Size = new(int32)
            *obj.Size = int32(parsedValue)
        }

    } else if jsonValue, exists := json["size"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Size = new(int32)
            *obj.Size = int32(parsedValue)
        }

    }
}

func (obj *StorageFile) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *StorageFile) GetIsFolder() *bool {
    return obj.IsFolder
}

func (obj *StorageFile) SetIsFolder(value *bool) {
    obj.IsFolder = value
}

func (obj *StorageFile) GetModifiedDate() *Time {
    return obj.ModifiedDate
}

func (obj *StorageFile) SetModifiedDate(value *Time) {
    obj.ModifiedDate = value
}

func (obj *StorageFile) GetName() *string {
    return obj.Name
}

func (obj *StorageFile) SetName(value *string) {
    obj.Name = value
}

func (obj *StorageFile) GetPath() *string {
    return obj.Path
}

func (obj *StorageFile) SetPath(value *string) {
    obj.Path = value
}

func (obj *StorageFile) GetSize() *int32 {
    return obj.Size
}

func (obj *StorageFile) SetSize(value *int32) {
    obj.Size = value
}

