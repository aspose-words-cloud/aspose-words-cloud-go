/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="section_link_collection_response.go">
 *   Copyright (c) 2024 Aspose.Words for Cloud
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

// The REST response with a collection of sections.
// This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/sections" REST API requests.

type ISectionLinkCollectionResponse interface {
    IsSectionLinkCollectionResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetSections() ISectionLinkCollection
    SetSections(value ISectionLinkCollection)
}

type SectionLinkCollectionResponse struct {
    // The REST response with a collection of sections.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/sections" REST API requests.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a collection of sections.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/sections" REST API requests.
    Sections ISectionLinkCollection `json:"Sections,omitempty"`
}

func (SectionLinkCollectionResponse) IsSectionLinkCollectionResponse() bool {
    return true
}

func (SectionLinkCollectionResponse) IsWordsResponse() bool {
    return true
}

func (obj *SectionLinkCollectionResponse) Initialize() {
    if (obj.Sections != nil) {
        obj.Sections.Initialize()
    }


}

func (obj *SectionLinkCollectionResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["Sections"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ISectionLinkCollection = new(SectionLinkCollection)
            modelInstance.Deserialize(parsedValue)
            obj.Sections = modelInstance
        }

    } else if jsonValue, exists := json["sections"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ISectionLinkCollection = new(SectionLinkCollection)
            modelInstance.Deserialize(parsedValue)
            obj.Sections = modelInstance
        }

    }
}

func (obj *SectionLinkCollectionResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *SectionLinkCollectionResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Sections != nil {
        if err := obj.Sections.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *SectionLinkCollectionResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *SectionLinkCollectionResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *SectionLinkCollectionResponse) GetSections() ISectionLinkCollection {
    return obj.Sections
}

func (obj *SectionLinkCollectionResponse) SetSections(value ISectionLinkCollection) {
    obj.Sections = value
}

