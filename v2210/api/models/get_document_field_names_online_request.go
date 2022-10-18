/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="get_document_field_names_online_request.go">
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

// GetDocumentFieldNamesOnlineRequest contains request data for WordsApiService.GetDocumentFieldNamesOnline method.
type GetDocumentFieldNamesOnlineRequest struct {
        // The template document.
        Template io.ReadCloser
    /* optional (nil or map[string]interface{}) with one or more of key / value pairs:
        key: "loadEncoding" value: (*string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
        key: "password" value: (*string) Password of protected Word document. Use the parameter to pass a password via SDK. SDK encrypts it automatically. We don't recommend to use the parameter to pass a plain password for direct call of API.
        key: "encryptedPassword" value: (*string) Password of protected Word document. Use the parameter to pass an encrypted password for direct calls of API. See SDK code for encyption details.
        key: "useNonMergeFields" value: (*bool) The flag indicating whether to use non merge fields. If true, result includes "mustache" field names. */
    Optionals map[string]interface{}
}


func (data *GetDocumentFieldNamesOnlineRequest) CreateRequestData() (RequestData, error) {

    var result RequestData
    var filesContentData = make([]FileReference, 0)

    result.Method = strings.ToUpper("put")

    // create path and map variables
    result.Path = "/words/online/get/mailMerge/FieldNames"

    result.Path = strings.Replace(result.Path, "/<nil>", "", -1)
    result.Path = strings.Replace(result.Path, "//", "/", -1)

    result.HeaderParams = make(map[string]string)
    result.QueryParams = url.Values{}
    result.FormParams = make([]FormParamContainer, 0)


    if err := typeCheckParameter(data.Optionals["loadEncoding"], "string", "data.Optionals[loadEncoding]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["password"], "string", "data.Optionals[password]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["encryptedPassword"], "string", "data.Optionals[encryptedPassword]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["useNonMergeFields"], "bool", "data.Optionals[useNonMergeFields]"); err != nil {
        return result, err
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


    if localVarTempParam, localVarOk := data.Optionals["useNonMergeFields"].(bool); localVarOk {
        result.QueryParams.Add("UseNonMergeFields", parameterToString(localVarTempParam, ""))
    }



    _template := data.Template
    if _template != nil {
        fbs, _ := ioutil.ReadAll(_template)
        _template.Close()
        result.FormParams = append(result.FormParams, NewFileFormParamContainer("template", fbs))
    }


    for _, fileContentData := range filesContentData {
        fbs, _ := ioutil.ReadAll(fileContentData.Content)
        result.FormParams = append(result.FormParams, NewFileFormParamContainer(fileContentData.Reference, fbs))
    }

    return result, nil
}

func (data *GetDocumentFieldNamesOnlineRequest) CreateResponse(reader io.Reader, boundary string) (response interface{}, err error) {
            var successPayload FieldNamesResponse
            if err = json.NewDecoder(reader).Decode(&successPayload); err != nil {
                return nil, err
            }

            return successPayload, err
}