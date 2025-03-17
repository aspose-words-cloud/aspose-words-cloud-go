/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="move_file_request.go">
 *   Copyright (c) 2025 Aspose.Words for Cloud
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
)

// MoveFileRequest contains request data for WordsApiService.MoveFile method.
type MoveFileRequest struct {
        // Destination file path e.g. '/dest.ext'.
        DestPath *string
        // Source file's path e.g. '/Folder 1/file.ext' or '/Bucket/Folder 1/file.ext'.
        SrcPath *string
    /* optional (nil or map[string]interface{}) with one or more of key / value pairs:
        key: "srcStorageName" value: (*string) Source storage name.
        key: "destStorageName" value: (*string) Destination storage name.
        key: "versionId" value: (*string) File version ID to move. */
    Optionals map[string]interface{}
}


func (data *MoveFileRequest) CreateRequestData() (RequestData, error) {
    var result RequestData
    var filesContentData = make([]FileReference, 0)
    if data == nil {
        return result, errors.New("Invalid object.")
    }

    result.Method = strings.ToUpper("put")

    // create path and map variables
    result.Path = "/words/storage/file/move/{srcPath}"
    result.Path = strings.Replace(result.Path, "{"+"srcPath"+"}", fmt.Sprintf("%v", *data.SrcPath), -1)

    result.Path = strings.Replace(result.Path, "/<nil>", "", -1)
    result.Path = strings.Replace(result.Path, "//", "/", -1)

    result.HeaderParams = make(map[string]string)
    result.QueryParams = url.Values{}
    result.FormParams = make([]FormParamContainer, 0)

    if (data.DestPath == nil) {
        return result, errors.New("Parameter DestPath is required.")
    }

    if (data.SrcPath == nil) {
        return result, errors.New("Parameter SrcPath is required.")
    }


    if err := typeCheckParameter(data.Optionals["srcStorageName"], "string", "data.Optionals[srcStorageName]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["destStorageName"], "string", "data.Optionals[destStorageName]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["versionId"], "string", "data.Optionals[versionId]"); err != nil {
        return result, err
    }



    result.QueryParams.Add("DestPath", parameterToString(*data.DestPath, ""))


    if localVarTempParam, localVarOk := data.Optionals["srcStorageName"].(string); localVarOk {
        result.QueryParams.Add("SrcStorageName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["destStorageName"].(string); localVarOk {
        result.QueryParams.Add("DestStorageName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["versionId"].(string); localVarOk {
        result.QueryParams.Add("VersionId", parameterToString(localVarTempParam, ""))
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

func (data *MoveFileRequest) CreateResponse(reader io.Reader, boundary string) (response interface{}, err error) {
        return nil, nil
}