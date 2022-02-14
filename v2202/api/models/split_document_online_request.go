/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="split_document_online_request.go">
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
	"mime/multipart"
)

// SplitDocumentOnlineRequest contains request data for WordsApiService.SplitDocumentOnline method.
type SplitDocumentOnlineRequest struct {
        // The document.
        Document io.ReadCloser
        // The format to split.
        Format *string
    /* optional (nil or map[string]interface{}) with one or more of key / value pairs:
        key: "loadEncoding" value: (*string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
        key: "password" value: (*string) Password of protected Word document. Use the parameter to pass a password via SDK. SDK encrypts it automatically. We don't recommend to use the parameter to pass a plain password for direct call of API.
        key: "encryptedPassword" value: (*string) Password of protected Word document. Use the parameter to pass an encrypted password for direct calls of API. See SDK code for encyption details.
        key: "destFileName" value: (*string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
        key: "from" value: (*int32) The start page.
        key: "to" value: (*int32) The end page.
        key: "zipOutput" value: (*bool) The flag indicating whether to ZIP the output.
        key: "fontsLocation" value: (*string) Folder in filestorage with custom fonts. */
    Optionals map[string]interface{}
}


func (data *SplitDocumentOnlineRequest) CreateRequestData() (RequestData, error) {

    var result RequestData

    result.Method = strings.ToUpper("put")

    // create path and map variables
    result.Path = "/words/online/put/split"

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
    if err := typeCheckParameter(data.Optionals["destFileName"], "string", "data.Optionals[destFileName]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["from"], "int32", "data.Optionals[from]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["to"], "int32", "data.Optionals[to]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["zipOutput"], "bool", "data.Optionals[zipOutput]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["fontsLocation"], "string", "data.Optionals[fontsLocation]"); err != nil {
        return result, err
    }


    result.QueryParams.Add("Format", parameterToString(*data.Format, ""))


    if localVarTempParam, localVarOk := data.Optionals["loadEncoding"].(string); localVarOk {
        result.QueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["password"].(string); localVarOk {
        result.QueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["encryptedPassword"].(string); localVarOk {
        result.QueryParams.Add("EncryptedPassword", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["destFileName"].(string); localVarOk {
        result.QueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["from"].(int32); localVarOk {
        result.QueryParams.Add("From", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["to"].(int32); localVarOk {
        result.QueryParams.Add("To", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["zipOutput"].(bool); localVarOk {
        result.QueryParams.Add("ZipOutput", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["fontsLocation"].(string); localVarOk {
        result.QueryParams.Add("FontsLocation", parameterToString(localVarTempParam, ""))
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

func (data *SplitDocumentOnlineRequest) CreateResponse(reader io.Reader, boundary string) (response interface{}, err error) {
    var successPayload SplitDocumentOnlineResponse
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