/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="upload_file_request.go">
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
    "fmt"
    "io/ioutil"
    "net/url"
    "strings"
    "io"
    "encoding/json"
)

// UploadFileRequest contains request data for WordsApiService.UploadFile method.
type UploadFileRequest struct {
        // File to upload.
        FileContent io.ReadCloser
        // Path where to upload including filename and extension e.g. /file.ext or /Folder 1/file.ext If the content is multipart and path does not contains the file name it tries to get them from filename parameter from Content-Disposition header.
        Path *string
    /* optional (nil or map[string]interface{}) with one or more of key / value pairs:
        key: "storageName" value: (*string) Storage name. */
    Optionals map[string]interface{}
}


func (data *UploadFileRequest) CreateRequestData() (RequestData, error) {

    var result RequestData
    var filesContentData = make([]FileReference, 0)

    result.Method = strings.ToUpper("put")

    // create path and map variables
    result.Path = "/words/storage/file/{path}"
    result.Path = strings.Replace(result.Path, "{"+"path"+"}", fmt.Sprintf("%v", *data.Path), -1)

    result.Path = strings.Replace(result.Path, "/<nil>", "", -1)
    result.Path = strings.Replace(result.Path, "//", "/", -1)

    result.HeaderParams = make(map[string]string)
    result.QueryParams = url.Values{}
    result.FormParams = make([]FormParamContainer, 0)


    if err := typeCheckParameter(data.Optionals["storageName"], "string", "data.Optionals[storageName]"); err != nil {
        return result, err
    }


    if localVarTempParam, localVarOk := data.Optionals["storageName"].(string); localVarOk {
        result.QueryParams.Add("StorageName", parameterToString(localVarTempParam, ""))
    }



    _fileContent := data.FileContent
    if _fileContent != nil {
        fbs, _ := ioutil.ReadAll(_fileContent)
        _fileContent.Close()
        result.FormParams = append(result.FormParams, NewFileFormParamContainer("fileContent", fbs))
    }


    for _, fileContentData := range filesContentData {
        fbs, _ := ioutil.ReadAll(fileContentData.Content)
        result.FormParams = append(result.FormParams, NewFileFormParamContainer(fileContentData.Id, fbs))
    }

    return result, nil
}

func (data *UploadFileRequest) CreateResponse(reader io.Reader, boundary string) (response interface{}, err error) {
            var successPayload FilesUploadResult
            if err = json.NewDecoder(reader).Decode(&successPayload); err != nil {
                return nil, err
            }

            return successPayload, err
}