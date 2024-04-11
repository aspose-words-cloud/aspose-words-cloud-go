/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="header_footers_response.go">
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

// The REST response with a collection of HeaderFooter elements.
// This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/headersfooters" REST API requests.

type IHeaderFootersResponse interface {
    IsHeaderFootersResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetHeaderFooters() IHeaderFooterLinkCollection
    SetHeaderFooters(value IHeaderFooterLinkCollection)
}

type HeaderFootersResponse struct {
    // The REST response with a collection of HeaderFooter elements.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/headersfooters" REST API requests.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a collection of HeaderFooter elements.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/headersfooters" REST API requests.
    HeaderFooters IHeaderFooterLinkCollection `json:"HeaderFooters,omitempty"`
}

func (HeaderFootersResponse) IsHeaderFootersResponse() bool {
    return true
}

func (HeaderFootersResponse) IsWordsResponse() bool {
    return true
}

func (obj *HeaderFootersResponse) Initialize() {
    if (obj.HeaderFooters != nil) {
        obj.HeaderFooters.Initialize()
    }


}

func (obj *HeaderFootersResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["HeaderFooters"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IHeaderFooterLinkCollection = new(HeaderFooterLinkCollection)
            modelInstance.Deserialize(parsedValue)
            obj.HeaderFooters = modelInstance
        }

    } else if jsonValue, exists := json["headerFooters"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IHeaderFooterLinkCollection = new(HeaderFooterLinkCollection)
            modelInstance.Deserialize(parsedValue)
            obj.HeaderFooters = modelInstance
        }

    }
}

func (obj *HeaderFootersResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *HeaderFootersResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.HeaderFooters != nil {
        if err := obj.HeaderFooters.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *HeaderFootersResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *HeaderFootersResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *HeaderFootersResponse) GetHeaderFooters() IHeaderFooterLinkCollection {
    return obj.HeaderFooters
}

func (obj *HeaderFootersResponse) SetHeaderFooters(value IHeaderFooterLinkCollection) {
    obj.HeaderFooters = value
}

