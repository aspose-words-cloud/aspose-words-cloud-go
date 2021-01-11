/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="update_list_online_request.go">
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
    "os"
	"net/url"
	"strings"
    "io"
    "encoding/json"
)

// UpdateListOnlineRequest contains request data for WordsApiService.UpdateListOnline method.
type UpdateListOnlineRequest struct {
        // The document.
        Document *os.File
        // The list Id.
        ListId *int32
        // List object.
        ListUpdate IListUpdate
    /* optional (nil or map[string]interface{}) with one or more of key / value pairs:
        key: "loadEncoding" value: (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
        key: "password" value: (string) Password for opening an encrypted document.
        key: "destFileName" value: (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
        key: "revisionAuthor" value: (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
        key: "revisionDateTime" value: (string) The date and time to use for revisions. */
    Optionals map[string]interface{}
}


func (data *UpdateListOnlineRequest) CreateRequestData() (RequestData, error) {

    var result RequestData

    result.Method = strings.ToUpper("put")

    // create path and map variables
    result.Path = "/words/online/put/lists/{listId}"
    result.Path = strings.Replace(result.Path, "{"+"listId"+"}", fmt.Sprintf("%v", *data.ListId), -1)

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
    if err := typeCheckParameter(data.Optionals["destFileName"], "string", "data.Optionals[destFileName]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["revisionAuthor"], "string", "data.Optionals[revisionAuthor]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["revisionDateTime"], "string", "data.Optionals[revisionDateTime]"); err != nil {
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


    if localVarTempParam, localVarOk := data.Optionals["revisionAuthor"].(string); localVarOk {
        result.QueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["revisionDateTime"].(string); localVarOk {
        result.QueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
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


    result.FormParams = append(result.FormParams, NewTextFormParamContainer("ListUpdate", parameterToString(data.ListUpdate, "")))



    return result, nil
}

func (data *UpdateListOnlineRequest) CreateResponse(reader io.Reader) (response interface{}, err error) {
    var successPayload UpdateListOnlineResponse
            if err = json.NewDecoder(reader).Decode(&successPayload); err != nil {
                return nil, err
            }

            return successPayload, err
}