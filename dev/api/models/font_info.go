/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="font_info.go">
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

// DTO container with font info.

type IFontInfo interface {
    IsFontInfo() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetFilePath() *string
    SetFilePath(value *string)
    GetFontFamilyName() *string
    SetFontFamilyName(value *string)
    GetFullFontName() *string
    SetFullFontName(value *string)
    GetVersion() *string
    SetVersion(value *string)
}

type FontInfo struct {
    // DTO container with font info.
    FilePath *string

    // DTO container with font info.
    FontFamilyName *string

    // DTO container with font info.
    FullFontName *string

    // DTO container with font info.
    Version *string
}

func (FontInfo) IsFontInfo() bool {
    return true
}


func (obj *FontInfo) Initialize() {
}

func (obj *FontInfo) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["FilePath"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FilePath = &parsedValue
        }

    } else if jsonValue, exists := json["filePath"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FilePath = &parsedValue
        }

    }

    if jsonValue, exists := json["FontFamilyName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FontFamilyName = &parsedValue
        }

    } else if jsonValue, exists := json["fontFamilyName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FontFamilyName = &parsedValue
        }

    }

    if jsonValue, exists := json["FullFontName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FullFontName = &parsedValue
        }

    } else if jsonValue, exists := json["fullFontName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FullFontName = &parsedValue
        }

    }

    if jsonValue, exists := json["Version"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Version = &parsedValue
        }

    } else if jsonValue, exists := json["version"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Version = &parsedValue
        }

    }
}

func (obj *FontInfo) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *FontInfo) GetFilePath() *string {
    return obj.FilePath
}

func (obj *FontInfo) SetFilePath(value *string) {
    obj.FilePath = value
}

func (obj *FontInfo) GetFontFamilyName() *string {
    return obj.FontFamilyName
}

func (obj *FontInfo) SetFontFamilyName(value *string) {
    obj.FontFamilyName = value
}

func (obj *FontInfo) GetFullFontName() *string {
    return obj.FullFontName
}

func (obj *FontInfo) SetFullFontName(value *string) {
    obj.FullFontName = value
}

func (obj *FontInfo) GetVersion() *string {
    return obj.Version
}

func (obj *FontInfo) SetVersion(value *string) {
    obj.Version = value
}

