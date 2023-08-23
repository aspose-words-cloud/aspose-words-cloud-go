/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="available_fonts_response.go">
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

// The REST response with data on system, additional and custom fonts, available for document processing.

type IAvailableFontsResponse interface {
    IsAvailableFontsResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetRequestId() *string
    SetRequestId(value *string)
    GetAdditionalFonts() []IFontInfo
    SetAdditionalFonts(value []IFontInfo)
    GetCustomFonts() []IFontInfo
    SetCustomFonts(value []IFontInfo)
    GetSystemFonts() []IFontInfo
    SetSystemFonts(value []IFontInfo)
}

type AvailableFontsResponse struct {
    // The REST response with data on system, additional and custom fonts, available for document processing.
    RequestId *string

    // The REST response with data on system, additional and custom fonts, available for document processing.
    AdditionalFonts []IFontInfo

    // The REST response with data on system, additional and custom fonts, available for document processing.
    CustomFonts []IFontInfo

    // The REST response with data on system, additional and custom fonts, available for document processing.
    SystemFonts []IFontInfo
}

func (AvailableFontsResponse) IsAvailableFontsResponse() bool {
    return true
}

func (AvailableFontsResponse) IsWordsResponse() bool {
    return true
}

func (obj *AvailableFontsResponse) Initialize() {
    if (obj.AdditionalFonts != nil) {
        for _, objElementAdditionalFonts := range obj.AdditionalFonts {
            objElementAdditionalFonts.Initialize()
        }
    }
    if (obj.CustomFonts != nil) {
        for _, objElementCustomFonts := range obj.CustomFonts {
            objElementCustomFonts.Initialize()
        }
    }
    if (obj.SystemFonts != nil) {
        for _, objElementSystemFonts := range obj.SystemFonts {
            objElementSystemFonts.Initialize()
        }
    }

}

func (obj *AvailableFontsResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["AdditionalFonts"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IFontInfo = new(FontInfo)
                    modelElementInstance.Deserialize(elementValue)
                    obj.AdditionalFonts = append(obj.AdditionalFonts, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["additionalFonts"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IFontInfo = new(FontInfo)
                    modelElementInstance.Deserialize(elementValue)
                    obj.AdditionalFonts = append(obj.AdditionalFonts, modelElementInstance)
                }

            }
        }

    }

    if jsonValue, exists := json["CustomFonts"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IFontInfo = new(FontInfo)
                    modelElementInstance.Deserialize(elementValue)
                    obj.CustomFonts = append(obj.CustomFonts, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["customFonts"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IFontInfo = new(FontInfo)
                    modelElementInstance.Deserialize(elementValue)
                    obj.CustomFonts = append(obj.CustomFonts, modelElementInstance)
                }

            }
        }

    }

    if jsonValue, exists := json["SystemFonts"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IFontInfo = new(FontInfo)
                    modelElementInstance.Deserialize(elementValue)
                    obj.SystemFonts = append(obj.SystemFonts, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["systemFonts"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IFontInfo = new(FontInfo)
                    modelElementInstance.Deserialize(elementValue)
                    obj.SystemFonts = append(obj.SystemFonts, modelElementInstance)
                }

            }
        }

    }
}

func (obj *AvailableFontsResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *AvailableFontsResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *AvailableFontsResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *AvailableFontsResponse) GetAdditionalFonts() []IFontInfo {
    return obj.AdditionalFonts
}

func (obj *AvailableFontsResponse) SetAdditionalFonts(value []IFontInfo) {
    obj.AdditionalFonts = value
}

func (obj *AvailableFontsResponse) GetCustomFonts() []IFontInfo {
    return obj.CustomFonts
}

func (obj *AvailableFontsResponse) SetCustomFonts(value []IFontInfo) {
    obj.CustomFonts = value
}

func (obj *AvailableFontsResponse) GetSystemFonts() []IFontInfo {
    return obj.SystemFonts
}

func (obj *AvailableFontsResponse) SetSystemFonts(value []IFontInfo) {
    obj.SystemFonts = value
}

