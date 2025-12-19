/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="hyperlink_response.go">
 *   Copyright (c) 2025 Aspose.Words for Cloud
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

// The REST response with a hyperlink.
// This response should be returned by the service when handling: GET /{name}/hyperlinks/{hyperlinkIndex}.

type IHyperlinkResponse interface {
    IsHyperlinkResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetHyperlink() IHyperlink
    SetHyperlink(value IHyperlink)
}

type HyperlinkResponse struct {
    // The REST response with a hyperlink.
    // This response should be returned by the service when handling: GET /{name}/hyperlinks/{hyperlinkIndex}.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a hyperlink.
    // This response should be returned by the service when handling: GET /{name}/hyperlinks/{hyperlinkIndex}.
    Hyperlink IHyperlink `json:"Hyperlink,omitempty"`
}

func (HyperlinkResponse) IsHyperlinkResponse() bool {
    return true
}

func (HyperlinkResponse) IsWordsResponse() bool {
    return true
}

func (obj *HyperlinkResponse) Initialize() {
    if (obj.Hyperlink != nil) {
        obj.Hyperlink.Initialize()
    }


}

func (obj *HyperlinkResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["Hyperlink"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IHyperlink = new(Hyperlink)
            modelInstance.Deserialize(parsedValue)
            obj.Hyperlink = modelInstance
        }

    } else if jsonValue, exists := json["hyperlink"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IHyperlink = new(Hyperlink)
            modelInstance.Deserialize(parsedValue)
            obj.Hyperlink = modelInstance
        }

    }
}

func (obj *HyperlinkResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *HyperlinkResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Hyperlink != nil {
        if err := obj.Hyperlink.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *HyperlinkResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *HyperlinkResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *HyperlinkResponse) GetHyperlink() IHyperlink {
    return obj.Hyperlink
}

func (obj *HyperlinkResponse) SetHyperlink(value IHyperlink) {
    obj.Hyperlink = value
}

