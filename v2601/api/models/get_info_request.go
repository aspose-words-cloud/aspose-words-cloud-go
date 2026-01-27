/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="get_info_request.go">
 *   Copyright (c) 2026 Aspose.Words for Cloud
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
    "io/ioutil"
    "net/url"
    "strings"
    "io"
    "encoding/json"
)

// GetInfoRequest contains request data for WordsApiService.GetInfo method.
type GetInfoRequest struct {
}


func (data *GetInfoRequest) CreateRequestData() (RequestData, error) {
    var result RequestData
    var filesContentData = make([]FileReference, 0)
    if data == nil {
        return result, errors.New("Invalid object.")
    }

    result.Method = strings.ToUpper("get")

    // create path and map variables
    result.Path = "/words/info"

    result.Path = strings.Replace(result.Path, "/<nil>", "", -1)
    result.Path = strings.Replace(result.Path, "//", "/", -1)

    result.HeaderParams = make(map[string]string)
    result.QueryParams = url.Values{}
    result.FormParams = make([]FormParamContainer, 0)







    result.FileReferences = filesContentData
    for _, fileContentData := range filesContentData {
        if fileContentData.Source == "Request" {
            fbs, _ := ioutil.ReadAll(fileContentData.Content)
            result.FormParams = append(result.FormParams, NewFileFormParamContainer(fileContentData.Reference, fbs))
        }
    }

    return result, nil
}

func (data *GetInfoRequest) CreateResponse(reader io.Reader, boundary string) (response interface{}, err error) {
            var successPayload IInfoResponse = new(InfoResponse)
            var jsonMap map[string]interface{}
            if err = json.NewDecoder(reader).Decode(&jsonMap); err != nil {
                return nil, err
            }

            successPayload.Deserialize(jsonMap)
            return successPayload, err
}