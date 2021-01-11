/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="copy_folder_request.go">
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
	"net/url"
	"strings"
    "io"
)

// CopyFolderRequest contains request data for WordsApiService.CopyFolder method.
type CopyFolderRequest struct {
        // Destination folder path e.g. '/dst'.
        DestPath *string
        // Source folder path e.g. /Folder1.
        SrcPath *string
    /* optional (nil or map[string]interface{}) with one or more of key / value pairs:
        key: "srcStorageName" value: (string) Source storage name.
        key: "destStorageName" value: (string) Destination storage name. */
    Optionals map[string]interface{}
}


func (data *CopyFolderRequest) CreateRequestData() (RequestData, error) {

    var result RequestData

    result.Method = strings.ToUpper("put")

    // create path and map variables
    result.Path = "/words/storage/folder/copy/{srcPath}"
    result.Path = strings.Replace(result.Path, "{"+"srcPath"+"}", fmt.Sprintf("%v", *data.SrcPath), -1)

    result.Path = strings.Replace(result.Path, "/<nil>", "", -1)
    result.Path = strings.Replace(result.Path, "//", "/", -1)

    result.HeaderParams = make(map[string]string)
    result.QueryParams = url.Values{}
    result.FormParams = make([]FormParamContainer, 0)

    if err := typeCheckParameter(data.Optionals["srcStorageName"], "string", "data.Optionals[srcStorageName]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["destStorageName"], "string", "data.Optionals[destStorageName]"); err != nil {
        return result, err
    }


    result.QueryParams.Add("DestPath", parameterToString(*data.DestPath, ""))


    if localVarTempParam, localVarOk := data.Optionals["srcStorageName"].(string); localVarOk {
        result.QueryParams.Add("SrcStorageName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["destStorageName"].(string); localVarOk {
        result.QueryParams.Add("DestStorageName", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

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




    return result, nil
}

func (data *CopyFolderRequest) CreateResponse(reader io.Reader) (response interface{}, err error) {
        return nil, nil
}