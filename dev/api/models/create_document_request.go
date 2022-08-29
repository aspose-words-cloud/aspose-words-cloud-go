/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="create_document_request.go">
 *   Copyright (c) 2022 Aspose.Words for Cloud
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
    "io/ioutil"
    "net/url"
    "strings"
    "io"
    "encoding/json"
)

// CreateDocumentRequest contains request data for WordsApiService.CreateDocument method.
type CreateDocumentRequest struct {
    /* optional (nil or map[string]interface{}) with one or more of key / value pairs:
        key: "fileName" value: (*string) The filename of the document.
        key: "folder" value: (*string) The path to the document folder.
        key: "storage" value: (*string) Original document storage. */
    Optionals map[string]interface{}
}


func (data *CreateDocumentRequest) CreateRequestData() (RequestData, error) {

    var result RequestData
    var filesContentData = make([]FileContent, 0)

    result.Method = strings.ToUpper("put")

    // create path and map variables
    result.Path = "/words/create"

    result.Path = strings.Replace(result.Path, "/<nil>", "", -1)
    result.Path = strings.Replace(result.Path, "//", "/", -1)

    result.HeaderParams = make(map[string]string)
    result.QueryParams = url.Values{}
    result.FormParams = make([]FormParamContainer, 0)


    if err := typeCheckParameter(data.Optionals["fileName"], "string", "data.Optionals[fileName]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["folder"], "string", "data.Optionals[folder]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["storage"], "string", "data.Optionals[storage]"); err != nil {
        return result, err
    }


    if localVarTempParam, localVarOk := data.Optionals["fileName"].(string); localVarOk {
        result.QueryParams.Add("FileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["folder"].(string); localVarOk {
        result.QueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["storage"].(string); localVarOk {
        result.QueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }




    for _, fileContentData := range filesContentData {
        fbs, _ := ioutil.ReadAll(fileContentData.Content)
        result.FormParams = append(result.FormParams, NewFileFormParamContainer(fileContentData.Id, fbs))
    }

    return result, nil
}

func (data *CreateDocumentRequest) CreateResponse(reader io.Reader, boundary string) (response interface{}, err error) {
            var successPayload DocumentResponse
            if err = json.NewDecoder(reader).Decode(&successPayload); err != nil {
                return nil, err
            }

            return successPayload, err
}