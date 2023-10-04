/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="table_properties_response.go">
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

// The REST response with a table.
// This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/tables/{0}/properties" REST API requests.

type ITablePropertiesResponse interface {
    IsTablePropertiesResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetRequestId() *string
    SetRequestId(value *string)
    GetProperties() ITableProperties
    SetProperties(value ITableProperties)
}

type TablePropertiesResponse struct {
    // The REST response with a table.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/tables/{0}/properties" REST API requests.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a table.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/tables/{0}/properties" REST API requests.
    Properties ITableProperties `json:"Properties,omitempty"`
}

func (TablePropertiesResponse) IsTablePropertiesResponse() bool {
    return true
}

func (TablePropertiesResponse) IsWordsResponse() bool {
    return true
}

func (obj *TablePropertiesResponse) Initialize() {
    if (obj.Properties != nil) {
        obj.Properties.Initialize()
    }


}

func (obj *TablePropertiesResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["Properties"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ITableProperties = new(TableProperties)
            modelInstance.Deserialize(parsedValue)
            obj.Properties = modelInstance
        }

    } else if jsonValue, exists := json["properties"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ITableProperties = new(TableProperties)
            modelInstance.Deserialize(parsedValue)
            obj.Properties = modelInstance
        }

    }
}

func (obj *TablePropertiesResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *TablePropertiesResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *TablePropertiesResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *TablePropertiesResponse) GetProperties() ITableProperties {
    return obj.Properties
}

func (obj *TablePropertiesResponse) SetProperties(value ITableProperties) {
    obj.Properties = value
}

