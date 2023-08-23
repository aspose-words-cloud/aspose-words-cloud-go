/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="stat_data_response.go">
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

// The REST response with document's statistical data.

type IStatDataResponse interface {
    IsStatDataResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetRequestId() *string
    SetRequestId(value *string)
    GetDocumentLink() IFileLink
    SetDocumentLink(value IFileLink)
    GetStatData() IDocumentStatData
    SetStatData(value IDocumentStatData)
}

type StatDataResponse struct {
    // The REST response with document's statistical data.
    RequestId *string

    // The REST response with document's statistical data.
    DocumentLink IFileLink

    // The REST response with document's statistical data.
    StatData IDocumentStatData
}

func (StatDataResponse) IsStatDataResponse() bool {
    return true
}

func (StatDataResponse) IsWordsResponse() bool {
    return true
}

func (obj *StatDataResponse) Initialize() {
    if (obj.DocumentLink != nil) {
        obj.DocumentLink.Initialize()
    }

    if (obj.StatData != nil) {
        obj.StatData.Initialize()
    }


}

func (obj *StatDataResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["DocumentLink"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFileLink = new(FileLink)
            modelInstance.Deserialize(parsedValue)
            obj.DocumentLink = modelInstance
        }

    } else if jsonValue, exists := json["documentLink"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFileLink = new(FileLink)
            modelInstance.Deserialize(parsedValue)
            obj.DocumentLink = modelInstance
        }

    }

    if jsonValue, exists := json["StatData"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IDocumentStatData = new(DocumentStatData)
            modelInstance.Deserialize(parsedValue)
            obj.StatData = modelInstance
        }

    } else if jsonValue, exists := json["statData"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IDocumentStatData = new(DocumentStatData)
            modelInstance.Deserialize(parsedValue)
            obj.StatData = modelInstance
        }

    }
}

func (obj *StatDataResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *StatDataResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *StatDataResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *StatDataResponse) GetDocumentLink() IFileLink {
    return obj.DocumentLink
}

func (obj *StatDataResponse) SetDocumentLink(value IFileLink) {
    obj.DocumentLink = value
}

func (obj *StatDataResponse) GetStatData() IDocumentStatData {
    return obj.StatData
}

func (obj *StatDataResponse) SetStatData(value IDocumentStatData) {
    obj.StatData = value
}

