/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="translate_node_id_request.go">
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
    "fmt"
    "errors"
    "io/ioutil"
    "net/url"
    "strings"
    "io"
    "encoding/json"
)

// TranslateNodeIdRequest contains request data for WordsApiService.TranslateNodeId method.
type TranslateNodeIdRequest struct {
        // The filename of the input document.
        Name *string
        // The node identifier. Identifier examples: id0.0.0.
        NodeId *string
    /* optional (nil or map[string]interface{}) with one or more of key / value pairs:
        key: "folder" value: (*string) Original document folder.
        key: "storage" value: (*string) Original document storage.
        key: "loadEncoding" value: (*string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
        key: "password" value: (*string) Password of protected Word document. Use the parameter to pass a password via SDK. SDK encrypts it automatically. We don't recommend to use the parameter to pass a plain password for direct call of API.
        key: "encryptedPassword" value: (*string) Password of protected Word document. Use the parameter to pass an encrypted password for direct calls of API. See SDK code for encyption details.
        key: "openTypeSupport" value: (*bool) The value indicates whether OpenType support is on. */
    Optionals map[string]interface{}
}


func (data *TranslateNodeIdRequest) CreateRequestData() (RequestData, error) {
    var result RequestData
    var filesContentData = make([]FileReference, 0)
    if data == nil {
        return result, errors.New("Invalid object.")
    }

    result.Method = strings.ToUpper("get")

    // create path and map variables
    result.Path = "/words/{name}/translate/{nodeId}"
    result.Path = strings.Replace(result.Path, "{"+"name"+"}", fmt.Sprintf("%v", *data.Name), -1)
    result.Path = strings.Replace(result.Path, "{"+"nodeId"+"}", fmt.Sprintf("%v", *data.NodeId), -1)

    result.Path = strings.Replace(result.Path, "/<nil>", "", -1)
    result.Path = strings.Replace(result.Path, "//", "/", -1)

    result.HeaderParams = make(map[string]string)
    result.QueryParams = url.Values{}
    result.FormParams = make([]FormParamContainer, 0)

    if (data.Name == nil) {
        return result, errors.New("Parameter Name is required.")
    }

    if (data.NodeId == nil) {
        return result, errors.New("Parameter NodeId is required.")
    }


    if err := typeCheckParameter(data.Optionals["folder"], "string", "data.Optionals[folder]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["storage"], "string", "data.Optionals[storage]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["loadEncoding"], "string", "data.Optionals[loadEncoding]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["password"], "string", "data.Optionals[password]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["encryptedPassword"], "string", "data.Optionals[encryptedPassword]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["openTypeSupport"], "bool", "data.Optionals[openTypeSupport]"); err != nil {
        return result, err
    }



    if localVarTempParam, localVarOk := data.Optionals["folder"].(string); localVarOk {
        result.QueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["storage"].(string); localVarOk {
        result.QueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["loadEncoding"].(string); localVarOk {
        result.QueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["password"].(string); localVarOk {
        result.QueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["encryptedPassword"].(string); localVarOk {
        result.QueryParams.Add("EncryptedPassword", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["openTypeSupport"].(bool); localVarOk {
        result.QueryParams.Add("OpenTypeSupport", parameterToString(localVarTempParam, ""))
    }




    result.FileReferences = filesContentData
    for _, fileContentData := range filesContentData {
        if fileContentData.Source == "Request" {
            fbs, _ := ioutil.ReadAll(fileContentData.Content)
            result.FormParams = append(result.FormParams, NewFileFormParamContainer(fileContentData.Reference, fbs))
        }
    }

    return result, nil
}

func (data *TranslateNodeIdRequest) CreateResponse(reader io.Reader, boundary string) (response interface{}, err error) {
            var successPayload ITranslateNodeIdResponse = new(TranslateNodeIdResponse)
            var jsonMap map[string]interface{}
            if err = json.NewDecoder(reader).Decode(&jsonMap); err != nil {
                return nil, err
            }

            successPayload.Deserialize(jsonMap)
            return successPayload, err
}