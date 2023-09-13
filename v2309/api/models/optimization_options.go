/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="optimization_options.go">
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

// Container class for the document optimization options.

type IOptimizationOptions interface {
    IsOptimizationOptions() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetMsWordVersion() *string
    SetMsWordVersion(value *string)
}

type OptimizationOptions struct {
    // Container class for the document optimization options.
    MsWordVersion *string `json:"MsWordVersion,omitempty"`
}

func (OptimizationOptions) IsOptimizationOptions() bool {
    return true
}


func (obj *OptimizationOptions) Initialize() {
}

func (obj *OptimizationOptions) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["MsWordVersion"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.MsWordVersion = &parsedValue
        }

    } else if jsonValue, exists := json["msWordVersion"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.MsWordVersion = &parsedValue
        }

    }
}

func (obj *OptimizationOptions) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *OptimizationOptions) GetMsWordVersion() *string {
    return obj.MsWordVersion
}

func (obj *OptimizationOptions) SetMsWordVersion(value *string) {
    obj.MsWordVersion = value
}

