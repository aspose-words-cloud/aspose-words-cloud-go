/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="words_api.go">
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

package api

import (
    "net/url"
    "net/http"
    "strings"
    "golang.org/x/net/context"
    "os"
    "io/ioutil"
    "encoding/json"
    "fmt"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/dev/api/models"
    "errors"
)

// Linger please
var (
    _ context.Context
)

type WordsApiService service

/* WordsApiService Accepts all revisions in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
@return models.RevisionsModificationResponse*/
func (a *WordsApiService) AcceptAllRevisions(ctx context.Context, name string, localVarOptionals map[string]interface{}) (models.RevisionsModificationResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.RevisionsModificationResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/revisions/acceptAll"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Appends documents to the original document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param documentList The collection of documents to append.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.DocumentResponse*/
func (a *WordsApiService) AppendDocument(ctx context.Context, name string, documentList models.IDocumentEntryList, localVarOptionals map[string]interface{}) (models.DocumentResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.DocumentResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/appendDocument"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &documentList
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Applies a style to the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param styleApply The style to apply.
 @param styledNodePath The path to the node in the document tree, that supports styles: ParagraphFormat, List, ListLevel, Table.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.WordsResponse*/
func (a *WordsApiService) ApplyStyleToDocumentElement(ctx context.Context, name string, styleApply models.IStyleApply, styledNodePath string, localVarOptionals map[string]interface{}) (models.WordsResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.WordsResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{styledNodePath}/style"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"styledNodePath"+"}", fmt.Sprintf("%v", styledNodePath), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &styleApply
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Executes the report generation process using the specified document template and the external data source in XML, JSON or CSV format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param data A string providing a data to populate the specified template. The string must be of one of the following types: xml, json, csv.
 @param reportEngineSettings An object providing a settings of report engine.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) The filename of the output document. If this parameter is omitted, the result will be saved with autogenerated name.
@return models.DocumentResponse*/
func (a *WordsApiService) BuildReport(ctx context.Context, name string, data string, reportEngineSettings models.IReportEngineSettings, localVarOptionals map[string]interface{}) (models.DocumentResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.DocumentResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/buildReport"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "multipart/form-data", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }


    localVarFormParams = append(localVarFormParams, NewTextFormParamContainer("Data", parameterToString(data, "")))


    localVarFormParams = append(localVarFormParams, NewTextFormParamContainer("ReportEngineSettings", parameterToString(reportEngineSettings, "")))


    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Executes the report generation process online using the specified document template and the external data source in XML, JSON or CSV format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param template File with template.
 @param data A string providing a data to populate the specified template. The string must be of one of the following types: xml, json, csv.
 @param reportEngineSettings An object providing a settings of report engine.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "documentFileName" (string) The filename of the output document, that will be used when the resulting document has a dynamic field {filename}. If it is not set, the "template" will be used instead.
@return *os.File*/
func (a *WordsApiService) BuildReportOnline(ctx context.Context, template *os.File, data string, reportEngineSettings models.IReportEngineSettings, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/buildReport"

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["documentFileName"], "string", "documentFileName"); err != nil {
        return nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["documentFileName"].(string); localVarOk {
        localVarQueryParams.Add("DocumentFileName", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "multipart/form-data", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }


    _template := template
    if _template != nil {
        fbs, _ := ioutil.ReadAll(_template)
        _template.Close()
        localVarFormParams = append(localVarFormParams, NewFileFormParamContainer(_template.Name(), fbs))
    }


    localVarFormParams = append(localVarFormParams, NewTextFormParamContainer("Data", parameterToString(data, "")))


    localVarFormParams = append(localVarFormParams, NewTextFormParamContainer("ReportEngineSettings", parameterToString(reportEngineSettings, "")))


    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    return localVarHttpResponse, err
}
/* WordsApiService Runs a multi-class text classification for the specified raw text.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param text The text to classify.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "bestClassesCount" (string) The number of the best classes to return.
@return models.ClassificationResponse*/
func (a *WordsApiService) Classify(ctx context.Context, text string, localVarOptionals map[string]interface{}) (models.ClassificationResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.ClassificationResponse
    )

    // create path and map variables
    localVarPath := "/words/classify"

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["bestClassesCount"], "string", "bestClassesCount"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["bestClassesCount"].(string); localVarOk {
        localVarQueryParams.Add("BestClassesCount", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &text
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Runs a multi-class text classification for the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param documentName The filename of the input document.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "bestClassesCount" (string) The number of the best classes to return.
     @param "taxonomy" (string) The taxonomy to use.
@return models.ClassificationResponse*/
func (a *WordsApiService) ClassifyDocument(ctx context.Context, documentName string, localVarOptionals map[string]interface{}) (models.ClassificationResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.ClassificationResponse
    )

    // create path and map variables
    localVarPath := "/words/{documentName}/classify"
    localVarPath = strings.Replace(localVarPath, "{"+"documentName"+"}", fmt.Sprintf("%v", documentName), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["bestClassesCount"], "string", "bestClassesCount"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["taxonomy"], "string", "taxonomy"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["bestClassesCount"].(string); localVarOk {
        localVarQueryParams.Add("BestClassesCount", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["taxonomy"].(string); localVarOk {
        localVarQueryParams.Add("Taxonomy", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Compares two documents.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param compareData The properties of the document to compare with.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
@return models.DocumentResponse*/
func (a *WordsApiService) CompareDocument(ctx context.Context, name string, compareData models.ICompareData, localVarOptionals map[string]interface{}) (models.DocumentResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.DocumentResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/compareDocument"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &compareData
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Converts a document on a local drive to the specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param document Converting document.
 @param format The format to convert.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) Original document storage.
     @param "outPath" (string) The path to the output document on a local storage.
     @param "fileNameFieldValue" (string) The filename of the output document, that will be used when the resulting document has a dynamic field {filename}. If it is not set, the "sourceFilename" will be used instead.
     @param "fontsLocation" (string) Folder in filestorage with custom fonts.
@return *os.File*/
func (a *WordsApiService) ConvertDocument(ctx context.Context, document *os.File, format string, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/convert"

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["outPath"], "string", "outPath"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["fileNameFieldValue"], "string", "fileNameFieldValue"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["fontsLocation"], "string", "fontsLocation"); err != nil {
        return nil, err
    }


    localVarQueryParams.Add("Format", parameterToString(format, ""))


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["outPath"].(string); localVarOk {
        localVarQueryParams.Add("OutPath", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["fileNameFieldValue"].(string); localVarOk {
        localVarQueryParams.Add("FileNameFieldValue", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["fontsLocation"].(string); localVarOk {
        localVarQueryParams.Add("FontsLocation", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "multipart/form-data", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }


    _document := document
    if _document != nil {
        fbs, _ := ioutil.ReadAll(_document)
        _document.Close()
        localVarFormParams = append(localVarFormParams, NewFileFormParamContainer(_document.Name(), fbs))
    }


    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    return localVarHttpResponse, err
}
/* WordsApiService Copy file.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param destPath Destination file path.
 @param srcPath Source file's path e.g. '/Folder 1/file.ext' or '/Bucket/Folder 1/file.ext'.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "srcStorageName" (string) Source storage name.
     @param "destStorageName" (string) Destination storage name.
     @param "versionId" (string) File version ID to copy.
@return */
func (a *WordsApiService) CopyFile(ctx context.Context, destPath string, srcPath string, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/storage/file/copy/{srcPath}"
    localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", fmt.Sprintf("%v", srcPath), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["srcStorageName"], "string", "srcStorageName"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destStorageName"], "string", "destStorageName"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["versionId"], "string", "versionId"); err != nil {
        return nil, err
    }


    localVarQueryParams.Add("DestPath", parameterToString(destPath, ""))


    if localVarTempParam, localVarOk := localVarOptionals["srcStorageName"].(string); localVarOk {
        localVarQueryParams.Add("SrcStorageName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destStorageName"].(string); localVarOk {
        localVarQueryParams.Add("DestStorageName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["versionId"].(string); localVarOk {
        localVarQueryParams.Add("VersionId", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    return localVarHttpResponse, err
}
/* WordsApiService Copy folder.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param destPath Destination folder path e.g. '/dst'.
 @param srcPath Source folder path e.g. /Folder1.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "srcStorageName" (string) Source storage name.
     @param "destStorageName" (string) Destination storage name.
@return */
func (a *WordsApiService) CopyFolder(ctx context.Context, destPath string, srcPath string, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/storage/folder/copy/{srcPath}"
    localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", fmt.Sprintf("%v", srcPath), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["srcStorageName"], "string", "srcStorageName"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destStorageName"], "string", "destStorageName"); err != nil {
        return nil, err
    }


    localVarQueryParams.Add("DestPath", parameterToString(destPath, ""))


    if localVarTempParam, localVarOk := localVarOptionals["srcStorageName"].(string); localVarOk {
        localVarQueryParams.Add("SrcStorageName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destStorageName"].(string); localVarOk {
        localVarQueryParams.Add("DestStorageName", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    return localVarHttpResponse, err
}
/* WordsApiService Makes a copy of the style in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param styleCopy The properties of the style.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.StyleResponse*/
func (a *WordsApiService) CopyStyle(ctx context.Context, name string, styleCopy models.IStyleCopy, localVarOptionals map[string]interface{}) (models.StyleResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("post")
        localVarPostBody interface{}
        successPayload models.StyleResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/styles/copy"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &styleCopy
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Supported extensions: ".doc", ".docx", ".docm", ".dot", ".dotm", ".dotx", ".flatopc", ".fopc", ".flatopc_macro", ".fopc_macro", ".flatopc_template", ".fopc_template", ".flatopc_template_macro", ".fopc_template_macro", ".wordml", ".wml", ".rtf".
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) Original document storage.
     @param "fileName" (string) The filename of the document.
     @param "folder" (string) The path to the document folder.
@return models.DocumentResponse*/
func (a *WordsApiService) CreateDocument(ctx context.Context, localVarOptionals map[string]interface{}) (models.DocumentResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.DocumentResponse
    )

    // create path and map variables
    localVarPath := "/words/create"

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["fileName"], "string", "fileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["fileName"].(string); localVarOk {
        localVarQueryParams.Add("FileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Create the folder.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param path Target folder's path e.g. Folder1/Folder2/. The folders will be created recursively.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" (string) Storage name.
@return */
func (a *WordsApiService) CreateFolder(ctx context.Context, path string, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/storage/folder/{path}"
    localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", path), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["storageName"], "string", "storageName"); err != nil {
        return nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["storageName"].(string); localVarOk {
        localVarQueryParams.Add("StorageName", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    return localVarHttpResponse, err
}
/* WordsApiService Adds a new or updates an existing document property.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param propertyName The name of the property.
 @param property The property with a new value.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.DocumentPropertyResponse*/
func (a *WordsApiService) CreateOrUpdateDocumentProperty(ctx context.Context, name string, propertyName string, property models.IDocumentPropertyCreateOrUpdate, localVarOptionals map[string]interface{}) (models.DocumentPropertyResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.DocumentPropertyResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/documentProperties/{propertyName}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", fmt.Sprintf("%v", propertyName), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &property
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Removes paragraph tab stops from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
@return models.TabStopsResponse*/
func (a *WordsApiService) DeleteAllParagraphTabStops(ctx context.Context, name string, index int32, localVarOptionals map[string]interface{}) (models.TabStopsResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("delete")
        localVarPostBody interface{}
        successPayload models.TabStopsResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/paragraphs/{index}/tabstops"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService The 'nodePath' parameter should refer to a paragraph, a cell or a row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param borderType Border type.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.BorderResponse*/
func (a *WordsApiService) DeleteBorder(ctx context.Context, name string, borderType string, localVarOptionals map[string]interface{}) (models.BorderResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("delete")
        localVarPostBody interface{}
        successPayload models.BorderResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/borders/{borderType}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"borderType"+"}", fmt.Sprintf("%v", borderType), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService The 'nodePath' parameter should refer to a paragraph, a cell or a row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.BordersResponse*/
func (a *WordsApiService) DeleteBorders(ctx context.Context, name string, localVarOptionals map[string]interface{}) (models.BordersResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("delete")
        localVarPostBody interface{}
        successPayload models.BordersResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/borders"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Removes a comment from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param commentIndex The index of the comment.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return */
func (a *WordsApiService) DeleteComment(ctx context.Context, name string, commentIndex int32, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("delete")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/{name}/comments/{commentIndex}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"commentIndex"+"}", fmt.Sprintf("%v", commentIndex), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    return localVarHttpResponse, err
}
/* WordsApiService Removes a document property.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param propertyName The name of the property.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return */
func (a *WordsApiService) DeleteDocumentProperty(ctx context.Context, name string, propertyName string, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("delete")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/{name}/documentProperties/{propertyName}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", fmt.Sprintf("%v", propertyName), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    return localVarHttpResponse, err
}
/* WordsApiService Removes a DrawingObject from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return */
func (a *WordsApiService) DeleteDrawingObject(ctx context.Context, name string, index int32, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("delete")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/drawingObjects/{index}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    return localVarHttpResponse, err
}
/* WordsApiService Removes a field from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return */
func (a *WordsApiService) DeleteField(ctx context.Context, name string, index int32, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("delete")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/fields/{index}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    return localVarHttpResponse, err
}
/* WordsApiService Removes fields from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return */
func (a *WordsApiService) DeleteFields(ctx context.Context, name string, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("delete")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/fields"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    return localVarHttpResponse, err
}
/* WordsApiService Delete file.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param path Path of the file including the file name and extension e.g. /folder1/file.ext.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" (string) Storage name.
     @param "versionId" (string) File version ID to delete.
@return */
func (a *WordsApiService) DeleteFile(ctx context.Context, path string, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("delete")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/storage/file/{path}"
    localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", path), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["storageName"], "string", "storageName"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["versionId"], "string", "versionId"); err != nil {
        return nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["storageName"].(string); localVarOk {
        localVarQueryParams.Add("StorageName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["versionId"].(string); localVarOk {
        localVarQueryParams.Add("VersionId", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    return localVarHttpResponse, err
}
/* WordsApiService Delete folder.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param path Folder path e.g. '/folder'.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" (string) Storage name.
     @param "recursive" (bool) Enable to delete folders, subfolders and files.
@return */
func (a *WordsApiService) DeleteFolder(ctx context.Context, path string, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("delete")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/storage/folder/{path}"
    localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", path), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["storageName"], "string", "storageName"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["recursive"], "bool", "recursive"); err != nil {
        return nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["storageName"].(string); localVarOk {
        localVarQueryParams.Add("StorageName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["recursive"].(bool); localVarOk {
        localVarQueryParams.Add("Recursive", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    return localVarHttpResponse, err
}
/* WordsApiService Removes a footnote from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return */
func (a *WordsApiService) DeleteFootnote(ctx context.Context, name string, index int32, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("delete")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/footnotes/{index}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    return localVarHttpResponse, err
}
/* WordsApiService Removes a form field from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return */
func (a *WordsApiService) DeleteFormField(ctx context.Context, name string, index int32, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("delete")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/formfields/{index}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    return localVarHttpResponse, err
}
/* WordsApiService Removes a HeaderFooter object from the document section.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param sectionPath The path to the section in the document tree.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return */
func (a *WordsApiService) DeleteHeaderFooter(ctx context.Context, name string, sectionPath string, index int32, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("delete")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/{name}/{sectionPath}/headersfooters/{index}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sectionPath"+"}", fmt.Sprintf("%v", sectionPath), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    return localVarHttpResponse, err
}
/* WordsApiService Removes HeaderFooter objects from the document section.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param sectionPath The path to the section in the document tree.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
     @param "headersFootersTypes" (string) The list of HeaderFooter types.
@return */
func (a *WordsApiService) DeleteHeadersFooters(ctx context.Context, name string, sectionPath string, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("delete")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/{name}/{sectionPath}/headersfooters"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sectionPath"+"}", fmt.Sprintf("%v", sectionPath), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["headersFootersTypes"], "string", "headersFootersTypes"); err != nil {
        return nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["headersFootersTypes"].(string); localVarOk {
        localVarQueryParams.Add("HeadersFootersTypes", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    return localVarHttpResponse, err
}
/* WordsApiService Removes macros from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return */
func (a *WordsApiService) DeleteMacros(ctx context.Context, name string, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("delete")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/{name}/macros"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    return localVarHttpResponse, err
}
/* WordsApiService Removes an OfficeMath object from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return */
func (a *WordsApiService) DeleteOfficeMathObject(ctx context.Context, name string, index int32, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("delete")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/OfficeMathObjects/{index}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    return localVarHttpResponse, err
}
/* WordsApiService Removes a paragraph from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return */
func (a *WordsApiService) DeleteParagraph(ctx context.Context, name string, index int32, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("delete")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/paragraphs/{index}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    return localVarHttpResponse, err
}
/* WordsApiService Removes the formatting properties of a paragraph list from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.ParagraphListFormatResponse*/
func (a *WordsApiService) DeleteParagraphListFormat(ctx context.Context, name string, index int32, localVarOptionals map[string]interface{}) (models.ParagraphListFormatResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("delete")
        localVarPostBody interface{}
        successPayload models.ParagraphListFormatResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/paragraphs/{index}/listFormat"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Removes a paragraph tab stop from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param position The position of a tab stop to remove.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
@return models.TabStopsResponse*/
func (a *WordsApiService) DeleteParagraphTabStop(ctx context.Context, name string, position float64, index int32, localVarOptionals map[string]interface{}) (models.TabStopsResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("delete")
        localVarPostBody interface{}
        successPayload models.TabStopsResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/paragraphs/{index}/tabstop"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }


    localVarQueryParams.Add("Position", parameterToString(position, ""))


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Removes a Run object from the paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param paragraphPath The path to the paragraph in the document tree.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return */
func (a *WordsApiService) DeleteRun(ctx context.Context, name string, paragraphPath string, index int32, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("delete")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/{name}/{paragraphPath}/runs/{index}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"paragraphPath"+"}", fmt.Sprintf("%v", paragraphPath), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    return localVarHttpResponse, err
}
/* WordsApiService Removes a section from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param sectionIndex The index of the section.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return */
func (a *WordsApiService) DeleteSection(ctx context.Context, name string, sectionIndex int32, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("delete")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/{name}/sections/{sectionIndex}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sectionIndex"+"}", fmt.Sprintf("%v", sectionIndex), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    return localVarHttpResponse, err
}
/* WordsApiService Removes a table from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return */
func (a *WordsApiService) DeleteTable(ctx context.Context, name string, index int32, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("delete")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/tables/{index}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    return localVarHttpResponse, err
}
/* WordsApiService Removes a cell from the table row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param tableRowPath The path to the table row in the document tree.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return */
func (a *WordsApiService) DeleteTableCell(ctx context.Context, name string, tableRowPath string, index int32, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("delete")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/{name}/{tableRowPath}/cells/{index}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"tableRowPath"+"}", fmt.Sprintf("%v", tableRowPath), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    return localVarHttpResponse, err
}
/* WordsApiService Removes a row from the table.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param tablePath The path to the table in the document tree.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return */
func (a *WordsApiService) DeleteTableRow(ctx context.Context, name string, tablePath string, index int32, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("delete")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/{name}/{tablePath}/rows/{index}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"tablePath"+"}", fmt.Sprintf("%v", tablePath), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    return localVarHttpResponse, err
}
/* WordsApiService Removes a watermark from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.DocumentResponse*/
func (a *WordsApiService) DeleteWatermark(ctx context.Context, name string, localVarOptionals map[string]interface{}) (models.DocumentResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("post")
        localVarPostBody interface{}
        successPayload models.DocumentResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/watermarks/deleteLast"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Download file.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param path Path of the file including the file name and extension e.g. /folder1/file.ext.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" (string) Storage name.
     @param "versionId" (string) File version ID to download.
@return *os.File*/
func (a *WordsApiService) DownloadFile(ctx context.Context, path string, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/storage/file/{path}"
    localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", path), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["storageName"], "string", "storageName"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["versionId"], "string", "versionId"); err != nil {
        return nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["storageName"].(string); localVarOk {
        localVarQueryParams.Add("StorageName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["versionId"].(string); localVarOk {
        localVarQueryParams.Add("VersionId", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    return localVarHttpResponse, err
}
/* WordsApiService Executes a Mail Merge operation.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "data" (string) Mail merge data.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "withRegions" (bool) The flag indicating whether to execute Mail Merge operation with regions.
     @param "mailMergeDataFile" (string) The data file.
     @param "cleanup" (string) The cleanup options.
     @param "useWholeParagraphAsRegion" (bool) The flag indicating whether paragraph with TableStart or TableEnd field should be fully included into mail merge region or particular range between TableStart and TableEnd fields. The default value is true.
     @param "destFileName" (string) The filename of the output document. If this parameter is omitted, the result will be saved with autogenerated name.
@return models.DocumentResponse*/
func (a *WordsApiService) ExecuteMailMerge(ctx context.Context, name string, localVarOptionals map[string]interface{}) (models.DocumentResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.DocumentResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/MailMerge"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["data"], "string", "data"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["withRegions"], "bool", "withRegions"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["mailMergeDataFile"], "string", "mailMergeDataFile"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["cleanup"], "string", "cleanup"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["useWholeParagraphAsRegion"], "bool", "useWholeParagraphAsRegion"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["withRegions"].(bool); localVarOk {
        localVarQueryParams.Add("WithRegions", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["mailMergeDataFile"].(string); localVarOk {
        localVarQueryParams.Add("MailMergeDataFile", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["cleanup"].(string); localVarOk {
        localVarQueryParams.Add("Cleanup", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["useWholeParagraphAsRegion"].(bool); localVarOk {
        localVarQueryParams.Add("UseWholeParagraphAsRegion", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "multipart/form-data", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }


    if localVarTempParam, localVarOk := localVarOptionals["data"].(string); localVarOk {
        localVarFormParams = append(localVarFormParams, NewTextFormParamContainer("Data", parameterToString(localVarTempParam, "")))
    }


    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Executes a Mail Merge operation online.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param template File with template.
 @param data File with mailmerge data.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "withRegions" (bool) The flag indicating whether to execute Mail Merge operation with regions.
     @param "cleanup" (string) The cleanup options.
     @param "documentFileName" (string) The filename of the output document, that will be used when the resulting document has a dynamic field {filename}. If it is not set, the "template" will be used instead.
@return *os.File*/
func (a *WordsApiService) ExecuteMailMergeOnline(ctx context.Context, template *os.File, data *os.File, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/MailMerge"

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["withRegions"], "bool", "withRegions"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["cleanup"], "string", "cleanup"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["documentFileName"], "string", "documentFileName"); err != nil {
        return nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["withRegions"].(bool); localVarOk {
        localVarQueryParams.Add("WithRegions", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["cleanup"].(string); localVarOk {
        localVarQueryParams.Add("Cleanup", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["documentFileName"].(string); localVarOk {
        localVarQueryParams.Add("DocumentFileName", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "multipart/form-data", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }


    _template := template
    if _template != nil {
        fbs, _ := ioutil.ReadAll(_template)
        _template.Close()
        localVarFormParams = append(localVarFormParams, NewFileFormParamContainer(_template.Name(), fbs))
    }


    _data := data
    if _data != nil {
        fbs, _ := ioutil.ReadAll(_data)
        _data.Close()
        localVarFormParams = append(localVarFormParams, NewFileFormParamContainer(_data.Name(), fbs))
    }


    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    return localVarHttpResponse, err
}
/* WordsApiService Reads available fonts from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "fontsLocation" (string) The folder in cloud storage with custom fonts.
@return models.AvailableFontsResponse*/
func (a *WordsApiService) GetAvailableFonts(ctx context.Context, localVarOptionals map[string]interface{}) (models.AvailableFontsResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.AvailableFontsResponse
    )

    // create path and map variables
    localVarPath := "/words/fonts/available"

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["fontsLocation"], "string", "fontsLocation"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["fontsLocation"].(string); localVarOk {
        localVarQueryParams.Add("FontsLocation", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads a bookmark, specified by name, from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param bookmarkName The name of the bookmark.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.BookmarkResponse*/
func (a *WordsApiService) GetBookmarkByName(ctx context.Context, name string, bookmarkName string, localVarOptionals map[string]interface{}) (models.BookmarkResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.BookmarkResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/bookmarks/{bookmarkName}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"bookmarkName"+"}", fmt.Sprintf("%v", bookmarkName), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads bookmarks from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.BookmarksResponse*/
func (a *WordsApiService) GetBookmarks(ctx context.Context, name string, localVarOptionals map[string]interface{}) (models.BookmarksResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.BookmarksResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/bookmarks"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService The 'nodePath' parameter should refer to a paragraph, a cell or a row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param borderType Border type.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.BorderResponse*/
func (a *WordsApiService) GetBorder(ctx context.Context, name string, borderType string, localVarOptionals map[string]interface{}) (models.BorderResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.BorderResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/borders/{borderType}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"borderType"+"}", fmt.Sprintf("%v", borderType), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads borders from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.BordersResponse*/
func (a *WordsApiService) GetBorders(ctx context.Context, name string, localVarOptionals map[string]interface{}) (models.BordersResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.BordersResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/borders"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads a comment from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param commentIndex The index of the comment.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.CommentResponse*/
func (a *WordsApiService) GetComment(ctx context.Context, name string, commentIndex int32, localVarOptionals map[string]interface{}) (models.CommentResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.CommentResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/comments/{commentIndex}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"commentIndex"+"}", fmt.Sprintf("%v", commentIndex), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads comments from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.CommentsResponse*/
func (a *WordsApiService) GetComments(ctx context.Context, name string, localVarOptionals map[string]interface{}) (models.CommentsResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.CommentsResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/comments"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads common information from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param documentName The filename of the input document.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.DocumentResponse*/
func (a *WordsApiService) GetDocument(ctx context.Context, documentName string, localVarOptionals map[string]interface{}) (models.DocumentResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.DocumentResponse
    )

    // create path and map variables
    localVarPath := "/words/{documentName}"
    localVarPath = strings.Replace(localVarPath, "{"+"documentName"+"}", fmt.Sprintf("%v", documentName), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads a DrawingObject from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.DrawingObjectResponse*/
func (a *WordsApiService) GetDocumentDrawingObjectByIndex(ctx context.Context, name string, index int32, localVarOptionals map[string]interface{}) (models.DrawingObjectResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.DrawingObjectResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/drawingObjects/{index}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads image data of a DrawingObject from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return *os.File*/
func (a *WordsApiService) GetDocumentDrawingObjectImageData(ctx context.Context, name string, index int32, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/drawingObjects/{index}/imageData"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    return localVarHttpResponse, err
}
/* WordsApiService Reads OLE data of a DrawingObject from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return *os.File*/
func (a *WordsApiService) GetDocumentDrawingObjectOleData(ctx context.Context, name string, index int32, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/drawingObjects/{index}/oleData"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    return localVarHttpResponse, err
}
/* WordsApiService Reads DrawingObjects from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.DrawingObjectsResponse*/
func (a *WordsApiService) GetDocumentDrawingObjects(ctx context.Context, name string, localVarOptionals map[string]interface{}) (models.DrawingObjectsResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.DrawingObjectsResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/drawingObjects"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads merge field names from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "useNonMergeFields" (bool) The flag indicating whether to use non merge fields. If true, result includes "mustache" field names.
@return models.FieldNamesResponse*/
func (a *WordsApiService) GetDocumentFieldNames(ctx context.Context, name string, localVarOptionals map[string]interface{}) (models.FieldNamesResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.FieldNamesResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/mailMerge/FieldNames"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["useNonMergeFields"], "bool", "useNonMergeFields"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["useNonMergeFields"].(bool); localVarOk {
        localVarQueryParams.Add("UseNonMergeFields", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads merge field names from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param template File with template.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "useNonMergeFields" (bool) The flag indicating whether to use non merge fields. If true, result includes "mustache" field names.
@return models.FieldNamesResponse*/
func (a *WordsApiService) GetDocumentFieldNamesOnline(ctx context.Context, template *os.File, localVarOptionals map[string]interface{}) (models.FieldNamesResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.FieldNamesResponse
    )

    // create path and map variables
    localVarPath := "/words/mailMerge/FieldNames"

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["useNonMergeFields"], "bool", "useNonMergeFields"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["useNonMergeFields"].(bool); localVarOk {
        localVarQueryParams.Add("UseNonMergeFields", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "multipart/form-data", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }


    _template := template
    if _template != nil {
        fbs, _ := ioutil.ReadAll(_template)
        _template.Close()
        localVarFormParams = append(localVarFormParams, NewFileFormParamContainer(_template.Name(), fbs))
    }


    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads a hyperlink from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param hyperlinkIndex The index of the hyperlink.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.HyperlinkResponse*/
func (a *WordsApiService) GetDocumentHyperlinkByIndex(ctx context.Context, name string, hyperlinkIndex int32, localVarOptionals map[string]interface{}) (models.HyperlinkResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.HyperlinkResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/hyperlinks/{hyperlinkIndex}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"hyperlinkIndex"+"}", fmt.Sprintf("%v", hyperlinkIndex), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads hyperlinks from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.HyperlinksResponse*/
func (a *WordsApiService) GetDocumentHyperlinks(ctx context.Context, name string, localVarOptionals map[string]interface{}) (models.HyperlinksResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.HyperlinksResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/hyperlinks"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads document properties.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.DocumentPropertiesResponse*/
func (a *WordsApiService) GetDocumentProperties(ctx context.Context, name string, localVarOptionals map[string]interface{}) (models.DocumentPropertiesResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.DocumentPropertiesResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/documentProperties"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads a document property.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param propertyName The name of the property.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.DocumentPropertyResponse*/
func (a *WordsApiService) GetDocumentProperty(ctx context.Context, name string, propertyName string, localVarOptionals map[string]interface{}) (models.DocumentPropertyResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.DocumentPropertyResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/documentProperties/{propertyName}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", fmt.Sprintf("%v", propertyName), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads protection properties from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.ProtectionDataResponse*/
func (a *WordsApiService) GetDocumentProtection(ctx context.Context, name string, localVarOptionals map[string]interface{}) (models.ProtectionDataResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.ProtectionDataResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/protection"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads document statistics.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "includeComments" (bool) The flag indicating whether to include comments from the WordCount. The default value is "false".
     @param "includeFootnotes" (bool) The flag indicating whether to include footnotes from the WordCount. The default value is "false".
     @param "includeTextInShapes" (bool) The flag indicating whether to include shape's text from the WordCount. The default value is "false".
@return models.StatDataResponse*/
func (a *WordsApiService) GetDocumentStatistics(ctx context.Context, name string, localVarOptionals map[string]interface{}) (models.StatDataResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.StatDataResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/statistics"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["includeComments"], "bool", "includeComments"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["includeFootnotes"], "bool", "includeFootnotes"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["includeTextInShapes"], "bool", "includeTextInShapes"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["includeComments"].(bool); localVarOk {
        localVarQueryParams.Add("IncludeComments", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["includeFootnotes"].(bool); localVarOk {
        localVarQueryParams.Add("IncludeFootnotes", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["includeTextInShapes"].(bool); localVarOk {
        localVarQueryParams.Add("IncludeTextInShapes", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Converts a document in cloud storage to the specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param format The destination format.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "outPath" (string) The path to the output document.
     @param "fontsLocation" (string) Folder in filestorage with custom fonts.
@return *os.File*/
func (a *WordsApiService) GetDocumentWithFormat(ctx context.Context, name string, format string, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/{name}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["outPath"], "string", "outPath"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["fontsLocation"], "string", "fontsLocation"); err != nil {
        return nil, err
    }


    localVarQueryParams.Add("Format", parameterToString(format, ""))


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["outPath"].(string); localVarOk {
        localVarQueryParams.Add("OutPath", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["fontsLocation"].(string); localVarOk {
        localVarQueryParams.Add("FontsLocation", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    return localVarHttpResponse, err
}
/* WordsApiService Reads a field from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.FieldResponse*/
func (a *WordsApiService) GetField(ctx context.Context, name string, index int32, localVarOptionals map[string]interface{}) (models.FieldResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.FieldResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/fields/{index}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads fields from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.FieldsResponse*/
func (a *WordsApiService) GetFields(ctx context.Context, name string, localVarOptionals map[string]interface{}) (models.FieldsResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.FieldsResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/fields"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Get all files and folders within a folder.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param path Folder path e.g. '/folder'.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" (string) Storage name.
@return models.FilesList*/
func (a *WordsApiService) GetFilesList(ctx context.Context, path string, localVarOptionals map[string]interface{}) (models.FilesList, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.FilesList
    )

    // create path and map variables
    localVarPath := "/words/storage/folder/{path}"
    localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", path), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["storageName"], "string", "storageName"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["storageName"].(string); localVarOk {
        localVarQueryParams.Add("StorageName", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads a footnote from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.FootnoteResponse*/
func (a *WordsApiService) GetFootnote(ctx context.Context, name string, index int32, localVarOptionals map[string]interface{}) (models.FootnoteResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.FootnoteResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/footnotes/{index}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads footnotes from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.FootnotesResponse*/
func (a *WordsApiService) GetFootnotes(ctx context.Context, name string, localVarOptionals map[string]interface{}) (models.FootnotesResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.FootnotesResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/footnotes"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads a form field from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.FormFieldResponse*/
func (a *WordsApiService) GetFormField(ctx context.Context, name string, index int32, localVarOptionals map[string]interface{}) (models.FormFieldResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.FormFieldResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/formfields/{index}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads form fields from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.FormFieldsResponse*/
func (a *WordsApiService) GetFormFields(ctx context.Context, name string, localVarOptionals map[string]interface{}) (models.FormFieldsResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.FormFieldsResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/formfields"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads a HeaderFooter object from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param headerFooterIndex The index of the HeaderFooter object.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "filterByType" (string) The list of HeaderFooter types.
@return models.HeaderFooterResponse*/
func (a *WordsApiService) GetHeaderFooter(ctx context.Context, name string, headerFooterIndex int32, localVarOptionals map[string]interface{}) (models.HeaderFooterResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.HeaderFooterResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/headersfooters/{headerFooterIndex}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"headerFooterIndex"+"}", fmt.Sprintf("%v", headerFooterIndex), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["filterByType"], "string", "filterByType"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["filterByType"].(string); localVarOk {
        localVarQueryParams.Add("FilterByType", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads a HeaderFooter object from the document section.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param headerFooterIndex The index of the HeaderFooter object.
 @param sectionIndex The index of the section.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "filterByType" (string) The list of HeaderFooter types.
@return models.HeaderFooterResponse*/
func (a *WordsApiService) GetHeaderFooterOfSection(ctx context.Context, name string, headerFooterIndex int32, sectionIndex int32, localVarOptionals map[string]interface{}) (models.HeaderFooterResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.HeaderFooterResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/sections/{sectionIndex}/headersfooters/{headerFooterIndex}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"headerFooterIndex"+"}", fmt.Sprintf("%v", headerFooterIndex), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sectionIndex"+"}", fmt.Sprintf("%v", sectionIndex), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["filterByType"], "string", "filterByType"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["filterByType"].(string); localVarOk {
        localVarQueryParams.Add("FilterByType", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads HeaderFooter objects from the document section.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param sectionPath The path to the section in the document tree.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "filterByType" (string) The list of HeaderFooter types.
@return models.HeaderFootersResponse*/
func (a *WordsApiService) GetHeaderFooters(ctx context.Context, name string, sectionPath string, localVarOptionals map[string]interface{}) (models.HeaderFootersResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.HeaderFootersResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{sectionPath}/headersfooters"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sectionPath"+"}", fmt.Sprintf("%v", sectionPath), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["filterByType"], "string", "filterByType"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["filterByType"].(string); localVarOk {
        localVarQueryParams.Add("FilterByType", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads a list from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param listId The list Id.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.ListResponse*/
func (a *WordsApiService) GetList(ctx context.Context, name string, listId int32, localVarOptionals map[string]interface{}) (models.ListResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.ListResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/lists/{listId}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"listId"+"}", fmt.Sprintf("%v", listId), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads lists from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.ListsResponse*/
func (a *WordsApiService) GetLists(ctx context.Context, name string, localVarOptionals map[string]interface{}) (models.ListsResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.ListsResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/lists"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads an OfficeMath object from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.OfficeMathObjectResponse*/
func (a *WordsApiService) GetOfficeMathObject(ctx context.Context, name string, index int32, localVarOptionals map[string]interface{}) (models.OfficeMathObjectResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.OfficeMathObjectResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/OfficeMathObjects/{index}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads OfficeMath objects from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.OfficeMathObjectsResponse*/
func (a *WordsApiService) GetOfficeMathObjects(ctx context.Context, name string, localVarOptionals map[string]interface{}) (models.OfficeMathObjectsResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.OfficeMathObjectsResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/OfficeMathObjects"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads a paragraph from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.ParagraphResponse*/
func (a *WordsApiService) GetParagraph(ctx context.Context, name string, index int32, localVarOptionals map[string]interface{}) (models.ParagraphResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.ParagraphResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/paragraphs/{index}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads the formatting properties of a paragraph from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.ParagraphFormatResponse*/
func (a *WordsApiService) GetParagraphFormat(ctx context.Context, name string, index int32, localVarOptionals map[string]interface{}) (models.ParagraphFormatResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.ParagraphFormatResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/paragraphs/{index}/format"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads the formatting properties of a paragraph list from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.ParagraphListFormatResponse*/
func (a *WordsApiService) GetParagraphListFormat(ctx context.Context, name string, index int32, localVarOptionals map[string]interface{}) (models.ParagraphListFormatResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.ParagraphListFormatResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/paragraphs/{index}/listFormat"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads paragraphs from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.ParagraphLinkCollectionResponse*/
func (a *WordsApiService) GetParagraphs(ctx context.Context, name string, localVarOptionals map[string]interface{}) (models.ParagraphLinkCollectionResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.ParagraphLinkCollectionResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/paragraphs"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads paragraph tab stops from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.TabStopsResponse*/
func (a *WordsApiService) GetParagraphTabStops(ctx context.Context, name string, index int32, localVarOptionals map[string]interface{}) (models.TabStopsResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.TabStopsResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/paragraphs/{index}/tabstops"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads range text from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param rangeStartIdentifier The range start identifier.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "rangeEndIdentifier" (string) The range end identifier.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.RangeTextResponse*/
func (a *WordsApiService) GetRangeText(ctx context.Context, name string, rangeStartIdentifier string, localVarOptionals map[string]interface{}) (models.RangeTextResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.RangeTextResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/range/{rangeStartIdentifier}/{rangeEndIdentifier}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"rangeStartIdentifier"+"}", fmt.Sprintf("%v", rangeStartIdentifier), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"rangeEndIdentifier"+"}", fmt.Sprintf("%v", localVarOptionals["rangeEndIdentifier"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["rangeEndIdentifier"], "string", "rangeEndIdentifier"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads a Run object from the paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param paragraphPath The path to the paragraph in the document tree.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.RunResponse*/
func (a *WordsApiService) GetRun(ctx context.Context, name string, paragraphPath string, index int32, localVarOptionals map[string]interface{}) (models.RunResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.RunResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{paragraphPath}/runs/{index}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"paragraphPath"+"}", fmt.Sprintf("%v", paragraphPath), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads the font properties of a Run object from the paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param paragraphPath The path to the paragraph in the document tree.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.FontResponse*/
func (a *WordsApiService) GetRunFont(ctx context.Context, name string, paragraphPath string, index int32, localVarOptionals map[string]interface{}) (models.FontResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.FontResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{paragraphPath}/runs/{index}/font"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"paragraphPath"+"}", fmt.Sprintf("%v", paragraphPath), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads Run objects from the paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param paragraphPath The path to the paragraph in the document tree.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.RunsResponse*/
func (a *WordsApiService) GetRuns(ctx context.Context, name string, paragraphPath string, localVarOptionals map[string]interface{}) (models.RunsResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.RunsResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{paragraphPath}/runs"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"paragraphPath"+"}", fmt.Sprintf("%v", paragraphPath), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads a section from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param sectionIndex The index of the section.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.SectionResponse*/
func (a *WordsApiService) GetSection(ctx context.Context, name string, sectionIndex int32, localVarOptionals map[string]interface{}) (models.SectionResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.SectionResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/sections/{sectionIndex}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sectionIndex"+"}", fmt.Sprintf("%v", sectionIndex), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads the page setup of a section from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param sectionIndex The index of the section.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.SectionPageSetupResponse*/
func (a *WordsApiService) GetSectionPageSetup(ctx context.Context, name string, sectionIndex int32, localVarOptionals map[string]interface{}) (models.SectionPageSetupResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.SectionPageSetupResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/sections/{sectionIndex}/pageSetup"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sectionIndex"+"}", fmt.Sprintf("%v", sectionIndex), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads sections from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.SectionLinkCollectionResponse*/
func (a *WordsApiService) GetSections(ctx context.Context, name string, localVarOptionals map[string]interface{}) (models.SectionLinkCollectionResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.SectionLinkCollectionResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/sections"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads a style from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param styleName The name of the style.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.StyleResponse*/
func (a *WordsApiService) GetStyle(ctx context.Context, name string, styleName string, localVarOptionals map[string]interface{}) (models.StyleResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.StyleResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/styles/{styleName}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"styleName"+"}", fmt.Sprintf("%v", styleName), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads a style from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param styledNodePath The path to the node in the document tree, that supports styles: ParagraphFormat, List, ListLevel, Table.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.StyleResponse*/
func (a *WordsApiService) GetStyleFromDocumentElement(ctx context.Context, name string, styledNodePath string, localVarOptionals map[string]interface{}) (models.StyleResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.StyleResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{styledNodePath}/style"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"styledNodePath"+"}", fmt.Sprintf("%v", styledNodePath), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads styles from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.StylesResponse*/
func (a *WordsApiService) GetStyles(ctx context.Context, name string, localVarOptionals map[string]interface{}) (models.StylesResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.StylesResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/styles"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads a table from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.TableResponse*/
func (a *WordsApiService) GetTable(ctx context.Context, name string, index int32, localVarOptionals map[string]interface{}) (models.TableResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.TableResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/tables/{index}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads a cell from the table row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param tableRowPath The path to the table row in the document tree.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.TableCellResponse*/
func (a *WordsApiService) GetTableCell(ctx context.Context, name string, tableRowPath string, index int32, localVarOptionals map[string]interface{}) (models.TableCellResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.TableCellResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{tableRowPath}/cells/{index}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"tableRowPath"+"}", fmt.Sprintf("%v", tableRowPath), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads the formatting properties of a table cell.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param tableRowPath The path to the table row in the document tree.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.TableCellFormatResponse*/
func (a *WordsApiService) GetTableCellFormat(ctx context.Context, name string, tableRowPath string, index int32, localVarOptionals map[string]interface{}) (models.TableCellFormatResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.TableCellFormatResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{tableRowPath}/cells/{index}/cellformat"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"tableRowPath"+"}", fmt.Sprintf("%v", tableRowPath), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads properties of a table from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.TablePropertiesResponse*/
func (a *WordsApiService) GetTableProperties(ctx context.Context, name string, index int32, localVarOptionals map[string]interface{}) (models.TablePropertiesResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.TablePropertiesResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/tables/{index}/properties"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads a row from the table.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param tablePath The path to the table in the document tree.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.TableRowResponse*/
func (a *WordsApiService) GetTableRow(ctx context.Context, name string, tablePath string, index int32, localVarOptionals map[string]interface{}) (models.TableRowResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.TableRowResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{tablePath}/rows/{index}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"tablePath"+"}", fmt.Sprintf("%v", tablePath), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads the formatting properties of a table row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param tablePath The path to the table in the document tree.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.TableRowFormatResponse*/
func (a *WordsApiService) GetTableRowFormat(ctx context.Context, name string, tablePath string, index int32, localVarOptionals map[string]interface{}) (models.TableRowFormatResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.TableRowFormatResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{tablePath}/rows/{index}/rowformat"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"tablePath"+"}", fmt.Sprintf("%v", tablePath), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads tables from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.TableLinkCollectionResponse*/
func (a *WordsApiService) GetTables(ctx context.Context, name string, localVarOptionals map[string]interface{}) (models.TableLinkCollectionResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.TableLinkCollectionResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/tables"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Inserts a new comment to the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param comment The properties of the comment.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.CommentResponse*/
func (a *WordsApiService) InsertComment(ctx context.Context, name string, comment models.ICommentInsert, localVarOptionals map[string]interface{}) (models.CommentResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("post")
        localVarPostBody interface{}
        successPayload models.CommentResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/comments"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &comment
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Inserts a new DrawingObject to the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param drawingObject Drawing object parameters.
 @param imageFile File with image.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.DrawingObjectResponse*/
func (a *WordsApiService) InsertDrawingObject(ctx context.Context, name string, drawingObject models.IDrawingObjectInsert, imageFile *os.File, localVarOptionals map[string]interface{}) (models.DrawingObjectResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("post")
        localVarPostBody interface{}
        successPayload models.DrawingObjectResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/drawingObjects"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "multipart/form-data", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }


    localVarFormParams = append(localVarFormParams, NewTextFormParamContainer("DrawingObject", parameterToString(drawingObject, "")))


    _imageFile := imageFile
    if _imageFile != nil {
        fbs, _ := ioutil.ReadAll(_imageFile)
        _imageFile.Close()
        localVarFormParams = append(localVarFormParams, NewFileFormParamContainer(_imageFile.Name(), fbs))
    }


    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Inserts a new field to the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param field The properties of the field.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
     @param "insertBeforeNode" (string) The index of the node. A new field will be inserted before the node with the specified node Id.
@return models.FieldResponse*/
func (a *WordsApiService) InsertField(ctx context.Context, name string, field models.IFieldInsert, localVarOptionals map[string]interface{}) (models.FieldResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("post")
        localVarPostBody interface{}
        successPayload models.FieldResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/fields"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["insertBeforeNode"], "string", "insertBeforeNode"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["insertBeforeNode"].(string); localVarOk {
        localVarQueryParams.Add("InsertBeforeNode", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &field
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Inserts a new footnote to the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param footnoteDto The properties of the footnote.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.FootnoteResponse*/
func (a *WordsApiService) InsertFootnote(ctx context.Context, name string, footnoteDto models.IFootnoteInsert, localVarOptionals map[string]interface{}) (models.FootnoteResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("post")
        localVarPostBody interface{}
        successPayload models.FootnoteResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/footnotes"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &footnoteDto
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Inserts a new form field to the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param formField The properties of the form field.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
     @param "insertBeforeNode" (string) The index of the node. A new form field will be inserted before the node with the specified node Id.
@return models.FormFieldResponse*/
func (a *WordsApiService) InsertFormField(ctx context.Context, name string, formField models.IFormField, localVarOptionals map[string]interface{}) (models.FormFieldResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("post")
        localVarPostBody interface{}
        successPayload models.FormFieldResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/formfields"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["insertBeforeNode"], "string", "insertBeforeNode"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["insertBeforeNode"].(string); localVarOk {
        localVarQueryParams.Add("InsertBeforeNode", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &formField
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Inserts a new HeaderFooter object to the document section.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param headerFooterType The type of a HeaderFooter object.
 @param sectionPath The path to the section in the document tree.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.HeaderFooterResponse*/
func (a *WordsApiService) InsertHeaderFooter(ctx context.Context, name string, headerFooterType string, sectionPath string, localVarOptionals map[string]interface{}) (models.HeaderFooterResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.HeaderFooterResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{sectionPath}/headersfooters"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sectionPath"+"}", fmt.Sprintf("%v", sectionPath), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &headerFooterType
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Inserts a new list to the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param listInsert The properties of the list.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.ListResponse*/
func (a *WordsApiService) InsertList(ctx context.Context, name string, listInsert models.IListInsert, localVarOptionals map[string]interface{}) (models.ListResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("post")
        localVarPostBody interface{}
        successPayload models.ListResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/lists"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &listInsert
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Inserts a new or updates an existing paragraph tab stop in the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param dto The properties of the paragraph tab stop.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
@return models.TabStopsResponse*/
func (a *WordsApiService) InsertOrUpdateParagraphTabStop(ctx context.Context, name string, dto models.ITabStopInsert, index int32, localVarOptionals map[string]interface{}) (models.TabStopsResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("post")
        localVarPostBody interface{}
        successPayload models.TabStopsResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/paragraphs/{index}/tabstops"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &dto
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Inserts page numbers to the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param pageNumber The page numbers settings.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.DocumentResponse*/
func (a *WordsApiService) InsertPageNumbers(ctx context.Context, name string, pageNumber models.IPageNumber, localVarOptionals map[string]interface{}) (models.DocumentResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.DocumentResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/PageNumbers"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &pageNumber
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Inserts a new paragraph to the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param paragraph The properties of the paragraph.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
     @param "insertBeforeNode" (string) The index of the node. A new paragraph will be inserted before the node with the specified index.
@return models.ParagraphResponse*/
func (a *WordsApiService) InsertParagraph(ctx context.Context, name string, paragraph models.IParagraphInsert, localVarOptionals map[string]interface{}) (models.ParagraphResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("post")
        localVarPostBody interface{}
        successPayload models.ParagraphResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/paragraphs"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["insertBeforeNode"], "string", "insertBeforeNode"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["insertBeforeNode"].(string); localVarOk {
        localVarQueryParams.Add("InsertBeforeNode", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &paragraph
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Inserts a new Run object to the paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param paragraphPath The path to the paragraph in the document tree.
 @param run The properties of the Run object.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
     @param "insertBeforeNode" (string) The index of the node. A new Run object will be inserted before the node with the specified node Id.
@return models.RunResponse*/
func (a *WordsApiService) InsertRun(ctx context.Context, name string, paragraphPath string, run models.IRunInsert, localVarOptionals map[string]interface{}) (models.RunResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("post")
        localVarPostBody interface{}
        successPayload models.RunResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{paragraphPath}/runs"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"paragraphPath"+"}", fmt.Sprintf("%v", paragraphPath), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["insertBeforeNode"], "string", "insertBeforeNode"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["insertBeforeNode"].(string); localVarOk {
        localVarQueryParams.Add("InsertBeforeNode", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &run
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Inserts a new style to the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param styleInsert The properties of the style.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.StyleResponse*/
func (a *WordsApiService) InsertStyle(ctx context.Context, name string, styleInsert models.IStyleInsert, localVarOptionals map[string]interface{}) (models.StyleResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("post")
        localVarPostBody interface{}
        successPayload models.StyleResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/styles/insert"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &styleInsert
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Inserts a new table to the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param table The properties of the table.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.TableResponse*/
func (a *WordsApiService) InsertTable(ctx context.Context, name string, table models.ITableInsert, localVarOptionals map[string]interface{}) (models.TableResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("post")
        localVarPostBody interface{}
        successPayload models.TableResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/tables"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &table
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Inserts a new cell to the table row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param cell The properties of the cell.
 @param tableRowPath The path to the table row in the document tree.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.TableCellResponse*/
func (a *WordsApiService) InsertTableCell(ctx context.Context, name string, cell models.ITableCellInsert, tableRowPath string, localVarOptionals map[string]interface{}) (models.TableCellResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("post")
        localVarPostBody interface{}
        successPayload models.TableCellResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{tableRowPath}/cells"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"tableRowPath"+"}", fmt.Sprintf("%v", tableRowPath), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &cell
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Inserts a new row to the table.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param row The properties of the row.
 @param tablePath The path to the table in the document tree.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.TableRowResponse*/
func (a *WordsApiService) InsertTableRow(ctx context.Context, name string, row models.ITableRowInsert, tablePath string, localVarOptionals map[string]interface{}) (models.TableRowResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("post")
        localVarPostBody interface{}
        successPayload models.TableRowResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{tablePath}/rows"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"tablePath"+"}", fmt.Sprintf("%v", tablePath), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &row
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Inserts a new watermark image to the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "imageFile" (*os.File) File with image.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
     @param "rotationAngle" (float64) The rotation angle of the watermark.
     @param "image" (string) The filename of the image. If the parameter value is missing — the image data is expected in the request content.
@return models.DocumentResponse*/
func (a *WordsApiService) InsertWatermarkImage(ctx context.Context, name string, localVarOptionals map[string]interface{}) (models.DocumentResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("post")
        localVarPostBody interface{}
        successPayload models.DocumentResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/watermarks/images"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["rotationAngle"], "float64", "rotationAngle"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["image"], "string", "image"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["rotationAngle"].(float64); localVarOk {
        localVarQueryParams.Add("RotationAngle", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["image"].(string); localVarOk {
        localVarQueryParams.Add("Image", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "multipart/form-data", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }


    var imageFile (*os.File)
    if localVarTempParam, localVarOk := localVarOptionals["imageFile"].(*os.File); localVarOk {
        imageFile = localVarTempParam
    }
    _imageFile := imageFile
    if _imageFile != nil {
        fbs, _ := ioutil.ReadAll(_imageFile)
        _imageFile.Close()
        localVarFormParams = append(localVarFormParams, NewFileFormParamContainer(_imageFile.Name(), fbs))
    }


    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Inserts a new watermark text to the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param watermarkText The watermark text to insert.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.DocumentResponse*/
func (a *WordsApiService) InsertWatermarkText(ctx context.Context, name string, watermarkText models.IWatermarkText, localVarOptionals map[string]interface{}) (models.DocumentResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("post")
        localVarPostBody interface{}
        successPayload models.DocumentResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/watermarks/texts"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &watermarkText
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Downloads a document from the Web using URL and saves it to cloud storage in the specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param data The properties of data downloading.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) Original document storage.
@return models.SaveResponse*/
func (a *WordsApiService) LoadWebDocument(ctx context.Context, data models.ILoadWebDocumentData, localVarOptionals map[string]interface{}) (models.SaveResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.SaveResponse
    )

    // create path and map variables
    localVarPath := "/words/loadWebDocument"

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &data
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Move file.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param destPath Destination file path e.g. '/dest.ext'.
 @param srcPath Source file's path e.g. '/Folder 1/file.ext' or '/Bucket/Folder 1/file.ext'.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "srcStorageName" (string) Source storage name.
     @param "destStorageName" (string) Destination storage name.
     @param "versionId" (string) File version ID to move.
@return */
func (a *WordsApiService) MoveFile(ctx context.Context, destPath string, srcPath string, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/storage/file/move/{srcPath}"
    localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", fmt.Sprintf("%v", srcPath), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["srcStorageName"], "string", "srcStorageName"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destStorageName"], "string", "destStorageName"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["versionId"], "string", "versionId"); err != nil {
        return nil, err
    }


    localVarQueryParams.Add("DestPath", parameterToString(destPath, ""))


    if localVarTempParam, localVarOk := localVarOptionals["srcStorageName"].(string); localVarOk {
        localVarQueryParams.Add("SrcStorageName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destStorageName"].(string); localVarOk {
        localVarQueryParams.Add("DestStorageName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["versionId"].(string); localVarOk {
        localVarQueryParams.Add("VersionId", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    return localVarHttpResponse, err
}
/* WordsApiService Move folder.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param destPath Destination folder path to move to e.g '/dst'.
 @param srcPath Source folder path e.g. /Folder1.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "srcStorageName" (string) Source storage name.
     @param "destStorageName" (string) Destination storage name.
@return */
func (a *WordsApiService) MoveFolder(ctx context.Context, destPath string, srcPath string, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/storage/folder/move/{srcPath}"
    localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", fmt.Sprintf("%v", srcPath), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["srcStorageName"], "string", "srcStorageName"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destStorageName"], "string", "destStorageName"); err != nil {
        return nil, err
    }


    localVarQueryParams.Add("DestPath", parameterToString(destPath, ""))


    if localVarTempParam, localVarOk := localVarOptionals["srcStorageName"].(string); localVarOk {
        localVarQueryParams.Add("SrcStorageName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destStorageName"].(string); localVarOk {
        localVarQueryParams.Add("DestStorageName", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    return localVarHttpResponse, err
}
/* WordsApiService Applies document content optimization options, specific to a particular versions of Microsoft Word.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param options The document optimization options.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return */
func (a *WordsApiService) OptimizeDocument(ctx context.Context, name string, options models.IOptimizationOptions, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/{name}/compatibility/optimize"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &options
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    return localVarHttpResponse, err
}
/* WordsApiService Adds protection to the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param protectionRequest The protection settings.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
@return models.ProtectionDataResponse*/
func (a *WordsApiService) ProtectDocument(ctx context.Context, name string, protectionRequest models.IProtectionRequest, localVarOptionals map[string]interface{}) (models.ProtectionDataResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.ProtectionDataResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/protection"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &protectionRequest
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Rejects all revisions in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
@return models.RevisionsModificationResponse*/
func (a *WordsApiService) RejectAllRevisions(ctx context.Context, name string, localVarOptionals map[string]interface{}) (models.RevisionsModificationResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.RevisionsModificationResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/revisions/rejectAll"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Removes a range from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param rangeStartIdentifier The range start identifier.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "rangeEndIdentifier" (string) The range end identifier.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
@return models.DocumentResponse*/
func (a *WordsApiService) RemoveRange(ctx context.Context, name string, rangeStartIdentifier string, localVarOptionals map[string]interface{}) (models.DocumentResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("delete")
        localVarPostBody interface{}
        successPayload models.DocumentResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/range/{rangeStartIdentifier}/{rangeEndIdentifier}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"rangeStartIdentifier"+"}", fmt.Sprintf("%v", rangeStartIdentifier), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"rangeEndIdentifier"+"}", fmt.Sprintf("%v", localVarOptionals["rangeEndIdentifier"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["rangeEndIdentifier"], "string", "rangeEndIdentifier"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Renders a DrawingObject to the specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param format The destination format.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "fontsLocation" (string) Folder in filestorage with custom fonts.
@return *os.File*/
func (a *WordsApiService) RenderDrawingObject(ctx context.Context, name string, format string, index int32, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/drawingObjects/{index}/render"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["fontsLocation"], "string", "fontsLocation"); err != nil {
        return nil, err
    }


    localVarQueryParams.Add("Format", parameterToString(format, ""))


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["fontsLocation"].(string); localVarOk {
        localVarQueryParams.Add("FontsLocation", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    return localVarHttpResponse, err
}
/* WordsApiService Renders an OfficeMath object to the specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param format The destination format.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "fontsLocation" (string) Folder in filestorage with custom fonts.
@return *os.File*/
func (a *WordsApiService) RenderMathObject(ctx context.Context, name string, format string, index int32, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/OfficeMathObjects/{index}/render"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["fontsLocation"], "string", "fontsLocation"); err != nil {
        return nil, err
    }


    localVarQueryParams.Add("Format", parameterToString(format, ""))


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["fontsLocation"].(string); localVarOk {
        localVarQueryParams.Add("FontsLocation", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    return localVarHttpResponse, err
}
/* WordsApiService Renders a page to the specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param pageIndex The index of the page.
 @param format The destination format.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "fontsLocation" (string) Folder in filestorage with custom fonts.
@return *os.File*/
func (a *WordsApiService) RenderPage(ctx context.Context, name string, pageIndex int32, format string, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/{name}/pages/{pageIndex}/render"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"pageIndex"+"}", fmt.Sprintf("%v", pageIndex), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["fontsLocation"], "string", "fontsLocation"); err != nil {
        return nil, err
    }


    localVarQueryParams.Add("Format", parameterToString(format, ""))


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["fontsLocation"].(string); localVarOk {
        localVarQueryParams.Add("FontsLocation", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    return localVarHttpResponse, err
}
/* WordsApiService Renders a paragraph to the specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param format The destination format.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "fontsLocation" (string) Folder in filestorage with custom fonts.
@return *os.File*/
func (a *WordsApiService) RenderParagraph(ctx context.Context, name string, format string, index int32, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/paragraphs/{index}/render"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["fontsLocation"], "string", "fontsLocation"); err != nil {
        return nil, err
    }


    localVarQueryParams.Add("Format", parameterToString(format, ""))


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["fontsLocation"].(string); localVarOk {
        localVarQueryParams.Add("FontsLocation", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    return localVarHttpResponse, err
}
/* WordsApiService Renders a table to the specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param format The destination format.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "fontsLocation" (string) Folder in filestorage with custom fonts.
@return *os.File*/
func (a *WordsApiService) RenderTable(ctx context.Context, name string, format string, index int32, localVarOptionals map[string]interface{}) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/tables/{index}/render"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return nil, err
    }
    if err := typeCheckParameter(localVarOptionals["fontsLocation"], "string", "fontsLocation"); err != nil {
        return nil, err
    }


    localVarQueryParams.Add("Format", parameterToString(format, ""))


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["fontsLocation"].(string); localVarOk {
        localVarQueryParams.Add("FontsLocation", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    return localVarHttpResponse, err
}
/* WordsApiService Replaces text in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param replaceText The text replacement parameters.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.ReplaceTextResponse*/
func (a *WordsApiService) ReplaceText(ctx context.Context, name string, replaceText models.IReplaceTextParameters, localVarOptionals map[string]interface{}) (models.ReplaceTextResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.ReplaceTextResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/replaceText"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &replaceText
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Replaces a range with text in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param rangeStartIdentifier The range start identifier.
 @param rangeText The text replacement properties.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "rangeEndIdentifier" (string) The range end identifier.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
@return models.DocumentResponse*/
func (a *WordsApiService) ReplaceWithText(ctx context.Context, name string, rangeStartIdentifier string, rangeText models.IReplaceRange, localVarOptionals map[string]interface{}) (models.DocumentResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("post")
        localVarPostBody interface{}
        successPayload models.DocumentResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/range/{rangeStartIdentifier}/{rangeEndIdentifier}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"rangeStartIdentifier"+"}", fmt.Sprintf("%v", rangeStartIdentifier), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"rangeEndIdentifier"+"}", fmt.Sprintf("%v", localVarOptionals["rangeEndIdentifier"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["rangeEndIdentifier"], "string", "rangeEndIdentifier"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &rangeText
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Clears the font cache.
 * @param ctx context.Context for authentication, logging, tracing, etc.
@return */
func (a *WordsApiService) ResetCache(ctx context.Context) (*http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("delete")
        localVarPostBody interface{}
    )

    // create path and map variables
    localVarPath := "/words/fonts/cache"

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)



    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    return localVarHttpResponse, err
}
/* WordsApiService Converts a document in cloud storage to the specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param saveOptionsData The save options.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "fontsLocation" (string) Folder in filestorage with custom fonts.
@return models.SaveResponse*/
func (a *WordsApiService) SaveAs(ctx context.Context, name string, saveOptionsData models.ISaveOptionsData, localVarOptionals map[string]interface{}) (models.SaveResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.SaveResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/saveAs"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["fontsLocation"], "string", "fontsLocation"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["fontsLocation"].(string); localVarOk {
        localVarQueryParams.Add("FontsLocation", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &saveOptionsData
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Saves a range as a new document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param rangeStartIdentifier The range start identifier.
 @param documentParameters The parameters of a new document.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "rangeEndIdentifier" (string) The range end identifier.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.DocumentResponse*/
func (a *WordsApiService) SaveAsRange(ctx context.Context, name string, rangeStartIdentifier string, documentParameters models.IRangeDocument, localVarOptionals map[string]interface{}) (models.DocumentResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("post")
        localVarPostBody interface{}
        successPayload models.DocumentResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/range/{rangeStartIdentifier}/{rangeEndIdentifier}/SaveAs"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"rangeStartIdentifier"+"}", fmt.Sprintf("%v", rangeStartIdentifier), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"rangeEndIdentifier"+"}", fmt.Sprintf("%v", localVarOptionals["rangeEndIdentifier"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["rangeEndIdentifier"], "string", "rangeEndIdentifier"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &documentParameters
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Converts a document in cloud storage to TIFF format using detailed conversion settings.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param saveOptions The save options to TIFF format.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "useAntiAliasing" (bool) The flag indicating whether to use antialiasing.
     @param "useHighQualityRendering" (bool) The flag indicating whether to use high quality.
     @param "imageBrightness" (float64) The level of brightness for the generated images.
     @param "imageColorMode" (string) The color mode for the generated images.
     @param "imageContrast" (float64) The contrast for the generated images.
     @param "numeralFormat" (string) The images numeral format.
     @param "pageCount" (int32) The number of pages to render.
     @param "pageIndex" (int32) The index of the page to start rendering.
     @param "paperColor" (string) The background image color.
     @param "pixelFormat" (string) The pixel format of the generated images.
     @param "resolution" (float64) The resolution of the generated images.
     @param "scale" (float64) The zoom factor for the generated images.
     @param "tiffCompression" (string) The compression tipe.
     @param "dmlRenderingMode" (string) The optional dml rendering mode. The default value is Fallback.
     @param "dmlEffectsRenderingMode" (string) The optional dml effects rendering mode. The default value is Simplified.
     @param "tiffBinarizationMethod" (string) The optional TIFF binarization method. Possible values are: FloydSteinbergDithering, Threshold.
     @param "zipOutput" (bool) The flag indicating whether to ZIP the output.
     @param "fontsLocation" (string) Folder in filestorage with custom fonts.
@return models.SaveResponse*/
func (a *WordsApiService) SaveAsTiff(ctx context.Context, name string, saveOptions models.ITiffSaveOptionsData, localVarOptionals map[string]interface{}) (models.SaveResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.SaveResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/saveAs/tiff"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["useAntiAliasing"], "bool", "useAntiAliasing"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["useHighQualityRendering"], "bool", "useHighQualityRendering"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["imageBrightness"], "float64", "imageBrightness"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["imageColorMode"], "string", "imageColorMode"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["imageContrast"], "float64", "imageContrast"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["numeralFormat"], "string", "numeralFormat"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["pageCount"], "int32", "pageCount"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["pageIndex"], "int32", "pageIndex"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["paperColor"], "string", "paperColor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["pixelFormat"], "string", "pixelFormat"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["resolution"], "float64", "resolution"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["scale"], "float64", "scale"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["tiffCompression"], "string", "tiffCompression"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["dmlRenderingMode"], "string", "dmlRenderingMode"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["dmlEffectsRenderingMode"], "string", "dmlEffectsRenderingMode"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["tiffBinarizationMethod"], "string", "tiffBinarizationMethod"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["zipOutput"], "bool", "zipOutput"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["fontsLocation"], "string", "fontsLocation"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["useAntiAliasing"].(bool); localVarOk {
        localVarQueryParams.Add("UseAntiAliasing", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["useHighQualityRendering"].(bool); localVarOk {
        localVarQueryParams.Add("UseHighQualityRendering", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["imageBrightness"].(float64); localVarOk {
        localVarQueryParams.Add("ImageBrightness", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["imageColorMode"].(string); localVarOk {
        localVarQueryParams.Add("ImageColorMode", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["imageContrast"].(float64); localVarOk {
        localVarQueryParams.Add("ImageContrast", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["numeralFormat"].(string); localVarOk {
        localVarQueryParams.Add("NumeralFormat", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["pageCount"].(int32); localVarOk {
        localVarQueryParams.Add("PageCount", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["pageIndex"].(int32); localVarOk {
        localVarQueryParams.Add("PageIndex", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["paperColor"].(string); localVarOk {
        localVarQueryParams.Add("PaperColor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["pixelFormat"].(string); localVarOk {
        localVarQueryParams.Add("PixelFormat", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["resolution"].(float64); localVarOk {
        localVarQueryParams.Add("Resolution", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["scale"].(float64); localVarOk {
        localVarQueryParams.Add("Scale", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["tiffCompression"].(string); localVarOk {
        localVarQueryParams.Add("TiffCompression", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["dmlRenderingMode"].(string); localVarOk {
        localVarQueryParams.Add("DmlRenderingMode", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["dmlEffectsRenderingMode"].(string); localVarOk {
        localVarQueryParams.Add("DmlEffectsRenderingMode", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["tiffBinarizationMethod"].(string); localVarOk {
        localVarQueryParams.Add("TiffBinarizationMethod", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["zipOutput"].(bool); localVarOk {
        localVarQueryParams.Add("ZipOutput", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["fontsLocation"].(string); localVarOk {
        localVarQueryParams.Add("FontsLocation", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &saveOptions
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Searches text, specified by the regular expression, in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param pattern The regular expression used to find matches.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
@return models.SearchResponse*/
func (a *WordsApiService) Search(ctx context.Context, name string, pattern string, localVarOptionals map[string]interface{}) (models.SearchResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("get")
        localVarPostBody interface{}
        successPayload models.SearchResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/search"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }


    localVarQueryParams.Add("Pattern", parameterToString(pattern, ""))


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Splits a document into parts and saves them in the specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param format The format to split.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "from" (int32) The start page.
     @param "to" (int32) The end page.
     @param "zipOutput" (bool) The flag indicating whether to ZIP the output.
     @param "fontsLocation" (string) Folder in filestorage with custom fonts.
@return models.SplitDocumentResponse*/
func (a *WordsApiService) SplitDocument(ctx context.Context, name string, format string, localVarOptionals map[string]interface{}) (models.SplitDocumentResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.SplitDocumentResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/split"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["from"], "int32", "from"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["to"], "int32", "to"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["zipOutput"], "bool", "zipOutput"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["fontsLocation"], "string", "fontsLocation"); err != nil {
        return successPayload, nil, err
    }


    localVarQueryParams.Add("Format", parameterToString(format, ""))


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["from"].(int32); localVarOk {
        localVarQueryParams.Add("From", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["to"].(int32); localVarOk {
        localVarQueryParams.Add("To", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["zipOutput"].(bool); localVarOk {
        localVarQueryParams.Add("ZipOutput", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["fontsLocation"].(string); localVarOk {
        localVarQueryParams.Add("FontsLocation", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Removes protection from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param protectionRequest The protection settings.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
@return models.ProtectionDataResponse*/
func (a *WordsApiService) UnprotectDocument(ctx context.Context, name string, protectionRequest models.IProtectionRequest, localVarOptionals map[string]interface{}) (models.ProtectionDataResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("delete")
        localVarPostBody interface{}
        successPayload models.ProtectionDataResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/protection"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &protectionRequest
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates a bookmark in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param bookmarkData The properties of the bookmark.
 @param bookmarkName The name of the bookmark.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.BookmarkResponse*/
func (a *WordsApiService) UpdateBookmark(ctx context.Context, name string, bookmarkData models.IBookmarkData, bookmarkName string, localVarOptionals map[string]interface{}) (models.BookmarkResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.BookmarkResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/bookmarks/{bookmarkName}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"bookmarkName"+"}", fmt.Sprintf("%v", bookmarkName), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &bookmarkData
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService The 'nodePath' parameter should refer to a paragraph, a cell or a row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param borderProperties The new border properties to update.
 @param borderType Border type.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.BorderResponse*/
func (a *WordsApiService) UpdateBorder(ctx context.Context, name string, borderProperties models.IBorder, borderType string, localVarOptionals map[string]interface{}) (models.BorderResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.BorderResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/borders/{borderType}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"borderType"+"}", fmt.Sprintf("%v", borderType), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &borderProperties
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates a comment in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param commentIndex The index of the comment.
 @param comment The properties of the comment.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.CommentResponse*/
func (a *WordsApiService) UpdateComment(ctx context.Context, name string, commentIndex int32, comment models.ICommentUpdate, localVarOptionals map[string]interface{}) (models.CommentResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.CommentResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/comments/{commentIndex}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"commentIndex"+"}", fmt.Sprintf("%v", commentIndex), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &comment
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates a DrawingObject in the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param drawingObject Drawing object parameters.
 @param imageFile File with image.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.DrawingObjectResponse*/
func (a *WordsApiService) UpdateDrawingObject(ctx context.Context, name string, drawingObject models.IDrawingObjectUpdate, imageFile *os.File, index int32, localVarOptionals map[string]interface{}) (models.DrawingObjectResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.DrawingObjectResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/drawingObjects/{index}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "multipart/form-data", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }


    localVarFormParams = append(localVarFormParams, NewTextFormParamContainer("DrawingObject", parameterToString(drawingObject, "")))


    _imageFile := imageFile
    if _imageFile != nil {
        fbs, _ := ioutil.ReadAll(_imageFile)
        _imageFile.Close()
        localVarFormParams = append(localVarFormParams, NewFileFormParamContainer(_imageFile.Name(), fbs))
    }


    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates a field in the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param field The properties of the field.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.FieldResponse*/
func (a *WordsApiService) UpdateField(ctx context.Context, name string, field models.IFieldUpdate, index int32, localVarOptionals map[string]interface{}) (models.FieldResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.FieldResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/fields/{index}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &field
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reevaluates field values in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
@return models.DocumentResponse*/
func (a *WordsApiService) UpdateFields(ctx context.Context, name string, localVarOptionals map[string]interface{}) (models.DocumentResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.DocumentResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/updateFields"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates a footnote in the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param footnoteDto The properties of the footnote.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.FootnoteResponse*/
func (a *WordsApiService) UpdateFootnote(ctx context.Context, name string, footnoteDto models.IFootnoteUpdate, index int32, localVarOptionals map[string]interface{}) (models.FootnoteResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.FootnoteResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/footnotes/{index}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &footnoteDto
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates a form field in the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param formField The new form field properties.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.FormFieldResponse*/
func (a *WordsApiService) UpdateFormField(ctx context.Context, name string, formField models.IFormField, index int32, localVarOptionals map[string]interface{}) (models.FormFieldResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.FormFieldResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/formfields/{index}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &formField
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates a list in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param listUpdate The properties of the list.
 @param listId The list Id.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.ListResponse*/
func (a *WordsApiService) UpdateList(ctx context.Context, name string, listUpdate models.IListUpdate, listId int32, localVarOptionals map[string]interface{}) (models.ListResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.ListResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/lists/{listId}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"listId"+"}", fmt.Sprintf("%v", listId), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &listUpdate
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates the level of a List element in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param listUpdate The properties of the List element.
 @param listId The list Id.
 @param listLevel The list level.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.ListResponse*/
func (a *WordsApiService) UpdateListLevel(ctx context.Context, name string, listUpdate models.IListLevelUpdate, listId int32, listLevel int32, localVarOptionals map[string]interface{}) (models.ListResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.ListResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/lists/{listId}/listLevels/{listLevel}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"listId"+"}", fmt.Sprintf("%v", listId), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"listLevel"+"}", fmt.Sprintf("%v", listLevel), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &listUpdate
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates the formatting properties of a paragraph in the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param dto The formatting properties of a paragraph.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.ParagraphFormatResponse*/
func (a *WordsApiService) UpdateParagraphFormat(ctx context.Context, name string, dto models.IParagraphFormatUpdate, index int32, localVarOptionals map[string]interface{}) (models.ParagraphFormatResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.ParagraphFormatResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/paragraphs/{index}/format"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &dto
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates the formatting properties of a paragraph list in the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param dto The formatting properties of a paragraph list.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.ParagraphListFormatResponse*/
func (a *WordsApiService) UpdateParagraphListFormat(ctx context.Context, name string, dto models.IListFormatUpdate, index int32, localVarOptionals map[string]interface{}) (models.ParagraphListFormatResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.ParagraphListFormatResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/paragraphs/{index}/listFormat"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &dto
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates a Run object in the paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param run The properties of the Run object.
 @param paragraphPath The path to the paragraph in the document tree.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.RunResponse*/
func (a *WordsApiService) UpdateRun(ctx context.Context, name string, run models.IRunUpdate, paragraphPath string, index int32, localVarOptionals map[string]interface{}) (models.RunResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.RunResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{paragraphPath}/runs/{index}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"paragraphPath"+"}", fmt.Sprintf("%v", paragraphPath), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &run
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates the font properties of a Run object in the paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param fontDto The font properties of a Run object.
 @param paragraphPath The path to the paragraph in the document tree.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.FontResponse*/
func (a *WordsApiService) UpdateRunFont(ctx context.Context, name string, fontDto models.IFont, paragraphPath string, index int32, localVarOptionals map[string]interface{}) (models.FontResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.FontResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{paragraphPath}/runs/{index}/font"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"paragraphPath"+"}", fmt.Sprintf("%v", paragraphPath), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &fontDto
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates the page setup of a section in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param sectionIndex The index of the section.
 @param pageSetup The properties of the page setup.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.SectionPageSetupResponse*/
func (a *WordsApiService) UpdateSectionPageSetup(ctx context.Context, name string, sectionIndex int32, pageSetup models.IPageSetup, localVarOptionals map[string]interface{}) (models.SectionPageSetupResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.SectionPageSetupResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/sections/{sectionIndex}/pageSetup"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sectionIndex"+"}", fmt.Sprintf("%v", sectionIndex), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &pageSetup
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates a style in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param styleUpdate The properties of the style.
 @param styleName The name of the style.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.StyleResponse*/
func (a *WordsApiService) UpdateStyle(ctx context.Context, name string, styleUpdate models.IStyleUpdate, styleName string, localVarOptionals map[string]interface{}) (models.StyleResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.StyleResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/styles/{styleName}/update"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"styleName"+"}", fmt.Sprintf("%v", styleName), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &styleUpdate
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates the formatting properties of a cell in the table row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param format The cell format.
 @param tableRowPath The path to the table row in the document tree.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.TableCellFormatResponse*/
func (a *WordsApiService) UpdateTableCellFormat(ctx context.Context, name string, format models.ITableCellFormat, tableRowPath string, index int32, localVarOptionals map[string]interface{}) (models.TableCellFormatResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.TableCellFormatResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{tableRowPath}/cells/{index}/cellformat"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"tableRowPath"+"}", fmt.Sprintf("%v", tableRowPath), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &format
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates properties of a table in the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param properties The properties of the table.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "nodePath" (string) The path to the node in the document tree.
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.TablePropertiesResponse*/
func (a *WordsApiService) UpdateTableProperties(ctx context.Context, name string, properties models.ITableProperties, index int32, localVarOptionals map[string]interface{}) (models.TablePropertiesResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.TablePropertiesResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{nodePath}/tables/{index}/properties"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", localVarOptionals["nodePath"]), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["nodePath"], "string", "nodePath"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &properties
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates the formatting properties of a table row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The filename of the input document.
 @param format The row format.
 @param tablePath The path to the table in the document tree.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Original document folder.
     @param "storage" (string) Original document storage.
     @param "loadEncoding" (string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" (string) Password for opening an encrypted document.
     @param "destFileName" (string) Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" (string) Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" (string) The date and time to use for revisions.
@return models.TableRowFormatResponse*/
func (a *WordsApiService) UpdateTableRowFormat(ctx context.Context, name string, format models.ITableRowFormat, tablePath string, index int32, localVarOptionals map[string]interface{}) (models.TableRowFormatResponse, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.TableRowFormatResponse
    )

    // create path and map variables
    localVarPath := "/words/{name}/{tablePath}/rows/{index}/rowformat"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"tablePath"+"}", fmt.Sprintf("%v", tablePath), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["loadEncoding"], "string", "loadEncoding"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["password"], "string", "password"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["destFileName"], "string", "destFileName"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionAuthor"], "string", "revisionAuthor"); err != nil {
        return successPayload, nil, err
    }
    if err := typeCheckParameter(localVarOptionals["revisionDateTime"], "string", "revisionDateTime"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
        localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
        localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["loadEncoding"].(string); localVarOk {
        localVarQueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["password"].(string); localVarOk {
        localVarQueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["destFileName"].(string); localVarOk {
        localVarQueryParams.Add("DestFileName", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionAuthor"].(string); localVarOk {
        localVarQueryParams.Add("RevisionAuthor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := localVarOptionals["revisionDateTime"].(string); localVarOk {
        localVarQueryParams.Add("RevisionDateTime", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "application/xml", "application/json", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }



    localVarPostBody = &format
    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}
/* WordsApiService Upload file.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param fileContent File to upload.
 @param path Path where to upload including filename and extension e.g. /file.ext or /Folder 1/file.ext If the content is multipart and path does not contains the file name it tries to get them from filename parameter from Content-Disposition header.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" (string) Storage name.
@return models.FilesUploadResult*/
func (a *WordsApiService) UploadFile(ctx context.Context, fileContent *os.File, path string, localVarOptionals map[string]interface{}) (models.FilesUploadResult, *http.Response, error) {
    var (
        localVarHttpMethod = strings.ToUpper("put")
        localVarPostBody interface{}
        successPayload models.FilesUploadResult
    )

    // create path and map variables
    localVarPath := "/words/storage/file/{path}"
    localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", path), -1)

    localVarPath = strings.Replace(localVarPath, "/<nil>", "", -1)
    localVarPath = a.client.cfg.BaseUrl + strings.Replace(localVarPath, "//", "/", -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := make([]FormParamContainer, 0)

    if err := typeCheckParameter(localVarOptionals["storageName"], "string", "storageName"); err != nil {
        return successPayload, nil, err
    }


    if localVarTempParam, localVarOk := localVarOptionals["storageName"].(string); localVarOk {
        localVarQueryParams.Add("StorageName", parameterToString(localVarTempParam, ""))
    }


    // to determine the Content-Type header
    localVarHttpContentTypes := []string{ "multipart/form-data", }

    // set Content-Type header
    localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHttpContentType
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string{
        "application/xml",
        "application/json",
    }

    // set Accept header
    localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
    }


    _fileContent := fileContent
    if _fileContent != nil {
        fbs, _ := ioutil.ReadAll(_fileContent)
        _fileContent.Close()
        localVarFormParams = append(localVarFormParams, NewFileFormParamContainer(_fileContent.Name(), fbs))
    }


    r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
    if err != nil {
        return successPayload, nil, err
    }

    localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return successPayload, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return successPayload, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return successPayload, localVarHttpResponse, err
        }

        return successPayload, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
        return successPayload, localVarHttpResponse, err
    }

    return successPayload, localVarHttpResponse, err
}