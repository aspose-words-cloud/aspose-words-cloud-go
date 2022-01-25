/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="execute_mail_merge_online_request.go">
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
)

// ExecuteMailMergeOnlineRequest contains request data for WordsApiService.ExecuteMailMergeOnline method.
type ExecuteMailMergeOnlineRequest struct {
        // File with template.
        Template io.ReadCloser
        // File with mailmerge data.
        Data io.ReadCloser
    /* optional (nil or map[string]interface{}) with one or more of key / value pairs:
        key: "options" value: (IFieldOptions) Field options.
        key: "withRegions" value: (*bool) The flag indicating whether to execute Mail Merge operation with regions.
        key: "cleanup" value: (*string) The cleanup options.
        key: "documentFileName" value: (*string) The filename of the output document, that will be used when the resulting document has a dynamic field {filename}. If it is not set, the "template" will be used instead. */
    Optionals map[string]interface{}
}


func (data *ExecuteMailMergeOnlineRequest) CreateRequestData() (RequestData, error) {

    var result RequestData

    result.Method = strings.ToUpper("put")

    // create path and map variables
    result.Path = "/words/MailMerge"

    result.Path = strings.Replace(result.Path, "/<nil>", "", -1)
    result.Path = strings.Replace(result.Path, "//", "/", -1)

    result.HeaderParams = make(map[string]string)
    result.QueryParams = url.Values{}
    result.FormParams = make([]FormParamContainer, 0)


    if err := typeCheckParameter(data.Optionals["withRegions"], "bool", "data.Optionals[withRegions]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["cleanup"], "string", "data.Optionals[cleanup]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["documentFileName"], "string", "data.Optionals[documentFileName]"); err != nil {
        return result, err
    }


    if localVarTempParam, localVarOk := data.Optionals["withRegions"].(bool); localVarOk {
        result.QueryParams.Add("WithRegions", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["cleanup"].(string); localVarOk {
        result.QueryParams.Add("Cleanup", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["documentFileName"].(string); localVarOk {
        result.QueryParams.Add("DocumentFileName", parameterToString(localVarTempParam, ""))
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


    _data := data.Data
    if _data != nil {
        fbs, _ := ioutil.ReadAll(_data)
        _data.Close()
        result.FormParams = append(result.FormParams, NewFileFormParamContainer("data", fbs))
    }


    if localVarTempParam, localVarOk := data.Optionals["options"].(FieldOptions); localVarOk {
        result.FormParams = append(result.FormParams, NewTextFormParamContainer("Options", parameterToString(localVarTempParam, "")))
    }



    return result, nil
}

func (data *ExecuteMailMergeOnlineRequest) CreateResponse(reader io.Reader, boundary string) (response interface{}, err error) {
            return reader, nil
}