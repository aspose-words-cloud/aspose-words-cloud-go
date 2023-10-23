/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="font_response.go">
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

// The REST response with a font.
// This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/paragraphs/{0}/runs/{1}/font" REST API requests.

type IFontResponse interface {
    IsFontResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetFont() IFont
    SetFont(value IFont)
}

type FontResponse struct {
    // The REST response with a font.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/paragraphs/{0}/runs/{1}/font" REST API requests.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a font.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/paragraphs/{0}/runs/{1}/font" REST API requests.
    Font IFont `json:"Font,omitempty"`
}

func (FontResponse) IsFontResponse() bool {
    return true
}

func (FontResponse) IsWordsResponse() bool {
    return true
}

func (obj *FontResponse) Initialize() {
    if (obj.Font != nil) {
        obj.Font.Initialize()
    }


}

func (obj *FontResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["Font"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFont = new(Font)
            modelInstance.Deserialize(parsedValue)
            obj.Font = modelInstance
        }

    } else if jsonValue, exists := json["font"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFont = new(Font)
            modelInstance.Deserialize(parsedValue)
            obj.Font = modelInstance
        }

    }
}

func (obj *FontResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *FontResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.RequestId == nil {
        return errors.New("Property RequestId in FontResponse is required.")
    }

    if obj.Font == nil {
        return errors.New("Property Font in FontResponse is required.")
    }

    if obj.Font != nil {
        if err := obj.Font.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *FontResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *FontResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *FontResponse) GetFont() IFont {
    return obj.Font
}

func (obj *FontResponse) SetFont(value IFont) {
    obj.Font = value
}

