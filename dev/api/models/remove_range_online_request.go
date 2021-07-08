/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="remove_range_online_request.go">
 *   Copyright (c) 2021 Aspose.Words for Cloud
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
	"io/ioutil"
	"net/url"
	"strings"
    "io"
    "encoding/json"
	"mime/multipart"
)

// RemoveRangeOnlineRequest contains request data for WordsApiService.RemoveRangeOnline method.
type RemoveRangeOnlineRequest struct {
        // The document.
        Document io.ReadCloser
        // The range start identifier.
        RangeStartIdentifier *string
    /* optional (nil or map[string]interface{}) with one or more of key / value pairs:
        key: "rangeEndIdentifier" value: (*string) The range end identifier.
        key: "loadEncoding" value: (*string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
        key: "password" value: (*string) Password for opening an encrypted document.
        key: "destFileName" value: (*string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. */
    Optionals map[string]interface{}
}


func (data *RemoveRangeOnlineRequest) CreateRequestData() (RequestData, error) {

    var result RequestData

    result.Method = strings.ToUpper("put")

    // create path and map variables
    result.Path = "/words/online/delete/range/{rangeStartIdentifier}/{rangeEndIdentifier}"
    result.Path = strings.Replace(result.Path, "{"+"rangeStartIdentifier"+"}", fmt.Sprintf("%v", *data.RangeStartIdentifier), -1)
    result.Path = strings.Replace(result.Path, "{"+"rangeEndIdentifier"+"}", fmt.Sprintf("%v", data.Optionals["rangeEndIdentifier"]), -1)

    result.Path = strings.Replace(result.Path, "/<nil>", "", -1)
    result.Path = strings.Replace(result.Path, "//", "/", -1)

    result.HeaderParams = make(map[string]string)
    result.QueryParams = url.Values{}
    result.FormParams = make([]FormParamContainer, 0)

    if err := typeCheckParameter(data.Optionals["rangeEndIdentifier"], "string", "data.Optionals[rangeEndIdentifier]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["loadEncoding"], "string", "data.Optionals[loadEncoding]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["password"], "string", "data.Optionals[password]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["destFileName"], "string", "data.Optionals[destFileName]"); err != nil {
        return result, err
    }


    if localVarTempParam, localVarOk := data.Optionals["loadEncoding"].(string); localVarOk {
        result.QueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["password"].(string); localVarOk {
        result.QueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["destFileName"].(string); localVarOk {
        result.QueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
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


    _document := data.Document
    if _document != nil {
        fbs, _ := ioutil.ReadAll(_document)
        _document.Close()
        result.FormParams = append(result.FormParams, NewFileFormParamContainer("document", fbs))
    }



    return result, nil
}

func (data *RemoveRangeOnlineRequest) CreateResponse(reader io.Reader, boundary string) (response interface{}, err error) {
    var successPayload RemoveRangeOnlineResponse
    var part *multipart.Part

	mr := multipart.NewReader(reader, boundary)
    part, err = mr.NextPart()

    if err != nil && err != io.EOF {
        return successPayload, err
    }

    if err = json.NewDecoder(part).Decode(&successPayload.Model); err != nil {
        return successPayload, err
    }


    part, err = mr.NextPart()

    if err != nil && err != io.EOF {
        return successPayload, err
    }

    successPayload.Document = part


    return successPayload, err
}