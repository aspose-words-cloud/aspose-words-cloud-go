/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="build_report_online_request.go">
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

// BuildReportOnlineRequest contains request data for WordsApiService.BuildReportOnline method.
type BuildReportOnlineRequest struct {
        // File with template.
        Template io.ReadCloser
        // A string providing a data to populate the specified template. The string must be of one of the following types: xml, json, csv.
        Data *string
        // An object providing a settings of report engine.
        ReportEngineSettings IReportEngineSettings
    /* optional (nil or map[string]interface{}) with one or more of key / value pairs:
        key: "documentFileName" value: (*string) The filename of the output document, that will be used when the resulting document has a dynamic field {filename}. If it is not set, the "template" will be used instead. */
    Optionals map[string]interface{}
}


func (data *BuildReportOnlineRequest) CreateRequestData() (RequestData, error) {

    var result RequestData
    var filesContentData = make([]FileReference, 0)

    result.Method = strings.ToUpper("put")

    // create path and map variables
    result.Path = "/words/buildReport"

    result.Path = strings.Replace(result.Path, "/<nil>", "", -1)
    result.Path = strings.Replace(result.Path, "//", "/", -1)

    result.HeaderParams = make(map[string]string)
    result.QueryParams = url.Values{}
    result.FormParams = make([]FormParamContainer, 0)

    if (data.ReportEngineSettings != nil) {
        data.ReportEngineSettings.Initialize()
    }

    if err := typeCheckParameter(data.Optionals["documentFileName"], "string", "data.Optionals[documentFileName]"); err != nil {
        return result, err
    }


    if localVarTempParam, localVarOk := data.Optionals["documentFileName"].(string); localVarOk {
        result.QueryParams.Add("DocumentFileName", parameterToString(localVarTempParam, ""))
    }



    _template := data.Template
    if _template != nil {
        fbs, _ := ioutil.ReadAll(_template)
        _template.Close()
        result.FormParams = append(result.FormParams, NewFileFormParamContainer("template", fbs))
    }

    result.FormParams = append(result.FormParams, NewTextFormParamContainer("Data", parameterToString(data.Data, "")))

    result.FormParams = append(result.FormParams, NewJsonFormParamContainer("ReportEngineSettings", parameterToString(data.ReportEngineSettings, "")))


    for _, fileContentData := range filesContentData {
        fbs, _ := ioutil.ReadAll(fileContentData.Content)
        result.FormParams = append(result.FormParams, NewFileFormParamContainer(fileContentData.Id, fbs))
    }

    return result, nil
}

func (data *BuildReportOnlineRequest) CreateResponse(reader io.Reader, boundary string) (response interface{}, err error) {
            return reader, nil
}