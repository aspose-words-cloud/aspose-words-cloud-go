/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="run_response.go">
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

// The REST response with a Run element.

type IRunResponse interface {
    IsRunResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetRequestId() *string
    SetRequestId(value *string)
    GetRun() IRun
    SetRun(value IRun)
}

type RunResponse struct {
    // The REST response with a Run element.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a Run element.
    Run IRun `json:"Run,omitempty"`
}

func (RunResponse) IsRunResponse() bool {
    return true
}

func (RunResponse) IsWordsResponse() bool {
    return true
}

func (obj *RunResponse) Initialize() {
    if (obj.Run != nil) {
        obj.Run.Initialize()
    }


}

func (obj *RunResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["Run"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IRun = new(Run)
            modelInstance.Deserialize(parsedValue)
            obj.Run = modelInstance
        }

    } else if jsonValue, exists := json["run"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IRun = new(Run)
            modelInstance.Deserialize(parsedValue)
            obj.Run = modelInstance
        }

    }
}

func (obj *RunResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *RunResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *RunResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *RunResponse) GetRun() IRun {
    return obj.Run
}

func (obj *RunResponse) SetRun(value IRun) {
    obj.Run = value
}

