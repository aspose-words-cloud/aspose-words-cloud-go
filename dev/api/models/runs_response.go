/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="runs_response.go">
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

// The REST response with a collection of Run elements.
// This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/{paragraphPath}/runs" REST API requests.

type IRunsResponse interface {
    IsRunsResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetRequestId() *string
    SetRequestId(value *string)
    GetRuns() IRuns
    SetRuns(value IRuns)
}

type RunsResponse struct {
    // The REST response with a collection of Run elements.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/{paragraphPath}/runs" REST API requests.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a collection of Run elements.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/{paragraphPath}/runs" REST API requests.
    Runs IRuns `json:"Runs,omitempty"`
}

func (RunsResponse) IsRunsResponse() bool {
    return true
}

func (RunsResponse) IsWordsResponse() bool {
    return true
}

func (obj *RunsResponse) Initialize() {
    if (obj.Runs != nil) {
        obj.Runs.Initialize()
    }


}

func (obj *RunsResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["Runs"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IRuns = new(Runs)
            modelInstance.Deserialize(parsedValue)
            obj.Runs = modelInstance
        }

    } else if jsonValue, exists := json["runs"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IRuns = new(Runs)
            modelInstance.Deserialize(parsedValue)
            obj.Runs = modelInstance
        }

    }
}

func (obj *RunsResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *RunsResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *RunsResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *RunsResponse) GetRuns() IRuns {
    return obj.Runs
}

func (obj *RunsResponse) SetRuns(value IRuns) {
    obj.Runs = value
}

