/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="get_document_statistics_online_request.go">
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
    "io/ioutil"
    "net/url"
    "strings"
    "io"
    "encoding/json"
)

// GetDocumentStatisticsOnlineRequest contains request data for WordsApiService.GetDocumentStatisticsOnline method.
type GetDocumentStatisticsOnlineRequest struct {
        // The document.
        Document io.ReadCloser
    /* optional (nil or map[string]interface{}) with one or more of key / value pairs:
        key: "loadEncoding" value: (*string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
        key: "password" value: (*string) Password of protected Word document. Use the parameter to pass a password via SDK. SDK encrypts it automatically. We don't recommend to use the parameter to pass a plain password for direct call of API.
        key: "encryptedPassword" value: (*string) Password of protected Word document. Use the parameter to pass an encrypted password for direct calls of API. See SDK code for encyption details.
        key: "includeComments" value: (*bool) The flag indicating whether to include comments from the WordCount. The default value is "false".
        key: "includeFootnotes" value: (*bool) The flag indicating whether to include footnotes from the WordCount. The default value is "false".
        key: "includeTextInShapes" value: (*bool) The flag indicating whether to include shape's text from the WordCount. The default value is "false". */
    Optionals map[string]interface{}
}


func (data *GetDocumentStatisticsOnlineRequest) CreateRequestData() (RequestData, error) {
    var result RequestData
    var filesContentData = make([]FileReference, 0)
    if data == nil {
        return result, errors.New("Invalid object.")
    }

    result.Method = strings.ToUpper("put")

    // create path and map variables
    result.Path = "/words/online/get/statistics"

    result.Path = strings.Replace(result.Path, "/<nil>", "", -1)
    result.Path = strings.Replace(result.Path, "//", "/", -1)

    result.HeaderParams = make(map[string]string)
    result.QueryParams = url.Values{}
    result.FormParams = make([]FormParamContainer, 0)

    if (data.Document == nil) {
        return result, errors.New("Parameter Document is required.")
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
    if err := typeCheckParameter(data.Optionals["includeComments"], "bool", "data.Optionals[includeComments]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["includeFootnotes"], "bool", "data.Optionals[includeFootnotes]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["includeTextInShapes"], "bool", "data.Optionals[includeTextInShapes]"); err != nil {
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


    if localVarTempParam, localVarOk := data.Optionals["includeComments"].(bool); localVarOk {
        result.QueryParams.Add("IncludeComments", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["includeFootnotes"].(bool); localVarOk {
        result.QueryParams.Add("IncludeFootnotes", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["includeTextInShapes"].(bool); localVarOk {
        result.QueryParams.Add("IncludeTextInShapes", parameterToString(localVarTempParam, ""))
    }



    _document := data.Document
    if _document != nil {
        fbs, _ := ioutil.ReadAll(_document)
        _document.Close()
        result.FormParams = append(result.FormParams, NewFileFormParamContainer("document", fbs))
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

func (data *GetDocumentStatisticsOnlineRequest) CreateResponse(reader io.Reader, boundary string) (response interface{}, err error) {
            var successPayload IStatDataResponse = new(StatDataResponse)
            var jsonMap map[string]interface{}
            if err = json.NewDecoder(reader).Decode(&jsonMap); err != nil {
                return nil, err
            }

            successPayload.Deserialize(jsonMap)
            return successPayload, err
}