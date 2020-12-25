/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="get_document_field_names_online_request.go">
 *   Copyright (c) 2020 Aspose.Words for Cloud
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
    "os"
	"net/url"
	"strings"
    "io"
    "encoding/json"
)

// GetDocumentFieldNamesOnlineRequest contains request data for WordsApiService.GetDocumentFieldNamesOnline method.
type GetDocumentFieldNamesOnlineRequest struct {
        // File with template.
        Template *os.File
    /* optional (nil or map[string]interface{}) with one or more of key / value pairs:
        key: "useNonMergeFields" value: (bool) The flag indicating whether to use non merge fields. If true, result includes "mustache" field names. */
    Optionals map[string]interface{}
}


func (data *GetDocumentFieldNamesOnlineRequest) CreateRequestData() (RequestData, error) {

    var result RequestData

    result.Method = strings.ToUpper("put")

    // create path and map variables
    result.Path = "/words/mailMerge/FieldNames"

    result.Path = strings.Replace(result.Path, "/<nil>", "", -1)
    result.Path = strings.Replace(result.Path, "//", "/", -1)

    result.HeaderParams = make(map[string]string)
    result.QueryParams = url.Values{}
    result.FormParams = make([]FormParamContainer, 0)

    if err := typeCheckParameter(data.Optionals["useNonMergeFields"], "bool", "data.Optionals[useNonMergeFields]"); err != nil {
        return result, err
    }


    if localVarTempParam, localVarOk := data.Optionals["useNonMergeFields"].(bool); localVarOk {
        result.QueryParams.Add("UseNonMergeFields", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "multipart/form-data", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        result.HeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        result.HeaderParams["Accept"] = localVarHttpHeaderAccept
    }


    _template := data.Template
    if _template != nil {
        fbs, _ := ioutil.ReadAll(_template)
        _template.Close()
        result.FormParams = append(result.FormParams, NewFileFormParamContainer("template", fbs))
    }



    return result, nil
}

func (data *GetDocumentFieldNamesOnlineRequest) CreateResponse(reader io.Reader) (response interface{}, err error) {
            var successPayload FieldNamesResponse
            if err = json.NewDecoder(reader).Decode(&successPayload); err != nil {
                return nil, err
            }

            return successPayload, err
}