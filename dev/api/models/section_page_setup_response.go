/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="section_page_setup_response.go">
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

// The REST response with a page setup of a section.
// This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/sections/{0}/PageSetup" REST API requests.

type ISectionPageSetupResponse interface {
    IsSectionPageSetupResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetPageSetup() IPageSetup
    SetPageSetup(value IPageSetup)
}

type SectionPageSetupResponse struct {
    // The REST response with a page setup of a section.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/sections/{0}/PageSetup" REST API requests.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a page setup of a section.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/sections/{0}/PageSetup" REST API requests.
    PageSetup IPageSetup `json:"PageSetup,omitempty"`
}

func (SectionPageSetupResponse) IsSectionPageSetupResponse() bool {
    return true
}

func (SectionPageSetupResponse) IsWordsResponse() bool {
    return true
}

func (obj *SectionPageSetupResponse) Initialize() {
    if (obj.PageSetup != nil) {
        obj.PageSetup.Initialize()
    }


}

func (obj *SectionPageSetupResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["PageSetup"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IPageSetup = new(PageSetup)
            modelInstance.Deserialize(parsedValue)
            obj.PageSetup = modelInstance
        }

    } else if jsonValue, exists := json["pageSetup"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IPageSetup = new(PageSetup)
            modelInstance.Deserialize(parsedValue)
            obj.PageSetup = modelInstance
        }

    }
}

func (obj *SectionPageSetupResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *SectionPageSetupResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    return nil;
}

func (obj *SectionPageSetupResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *SectionPageSetupResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *SectionPageSetupResponse) GetPageSetup() IPageSetup {
    return obj.PageSetup
}

func (obj *SectionPageSetupResponse) SetPageSetup(value IPageSetup) {
    obj.PageSetup = value
}

