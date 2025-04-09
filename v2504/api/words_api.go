/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="words_api.go">
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

package api

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strings"

	"github.com/aspose-words-cloud/aspose-words-cloud-go/v2504/api/models"
)

// Linger please
var (
    _ context.Context
)

type WordsApiService service

/* WordsApiService Accepts all revisions in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.RevisionsModificationResponse*/
func (a *WordsApiService) AcceptAllRevisions(ctx context.Context, data *models.AcceptAllRevisionsRequest) (models.RevisionsModificationResponse, *http.Response, error) {
    var (
        successPayload models.RevisionsModificationResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Accepts all revisions in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return AcceptAllRevisionsOnlineResponse*/
func (a *WordsApiService) AcceptAllRevisionsOnline(ctx context.Context, data *models.AcceptAllRevisionsOnlineRequest) (models.AcceptAllRevisionsOnlineResponse, *http.Response, error) {
    var (
        successPayload models.AcceptAllRevisionsOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.AcceptAllRevisionsOnlineResponse), response, err
}

/* WordsApiService Appends documents to the original document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.DocumentResponse*/
func (a *WordsApiService) AppendDocument(ctx context.Context, data *models.AppendDocumentRequest) (models.DocumentResponse, *http.Response, error) {
    var (
        successPayload models.DocumentResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Appends documents to the original document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return AppendDocumentOnlineResponse*/
func (a *WordsApiService) AppendDocumentOnline(ctx context.Context, data *models.AppendDocumentOnlineRequest) (models.AppendDocumentOnlineResponse, *http.Response, error) {
    var (
        successPayload models.AppendDocumentOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.AppendDocumentOnlineResponse), response, err
}

/* WordsApiService Applies a style to the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.WordsResponse*/
func (a *WordsApiService) ApplyStyleToDocumentElement(ctx context.Context, data *models.ApplyStyleToDocumentElementRequest) (models.WordsResponse, *http.Response, error) {
    var (
        successPayload models.WordsResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Applies a style to the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return ApplyStyleToDocumentElementOnlineResponse*/
func (a *WordsApiService) ApplyStyleToDocumentElementOnline(ctx context.Context, data *models.ApplyStyleToDocumentElementOnlineRequest) (models.ApplyStyleToDocumentElementOnlineResponse, *http.Response, error) {
    var (
        successPayload models.ApplyStyleToDocumentElementOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.ApplyStyleToDocumentElementOnlineResponse), response, err
}

/* WordsApiService Executes the report generation process using the specified document template and the external data source in XML, JSON or CSV format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.DocumentResponse*/
func (a *WordsApiService) BuildReport(ctx context.Context, data *models.BuildReportRequest) (models.DocumentResponse, *http.Response, error) {
    var (
        successPayload models.DocumentResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Executes the report generation process online using the specified document template and the external data source in XML, JSON or CSV format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return http.Response.Body contains binary result data*/
func (a *WordsApiService) BuildReportOnline(ctx context.Context, data *models.BuildReportOnlineRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Runs a multi-class text classification for the specified raw text.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.ClassificationResponse*/
func (a *WordsApiService) Classify(ctx context.Context, data *models.ClassifyRequest) (models.ClassificationResponse, *http.Response, error) {
    var (
        successPayload models.ClassificationResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Runs a multi-class text classification for the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.ClassificationResponse*/
func (a *WordsApiService) ClassifyDocument(ctx context.Context, data *models.ClassifyDocumentRequest) (models.ClassificationResponse, *http.Response, error) {
    var (
        successPayload models.ClassificationResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Runs a multi-class text classification for the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.ClassificationResponse*/
func (a *WordsApiService) ClassifyDocumentOnline(ctx context.Context, data *models.ClassifyDocumentOnlineRequest) (models.ClassificationResponse, *http.Response, error) {
    var (
        successPayload models.ClassificationResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Compares two documents.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.DocumentResponse*/
func (a *WordsApiService) CompareDocument(ctx context.Context, data *models.CompareDocumentRequest) (models.DocumentResponse, *http.Response, error) {
    var (
        successPayload models.DocumentResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Compares two documents.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return CompareDocumentOnlineResponse*/
func (a *WordsApiService) CompareDocumentOnline(ctx context.Context, data *models.CompareDocumentOnlineRequest) (models.CompareDocumentOnlineResponse, *http.Response, error) {
    var (
        successPayload models.CompareDocumentOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.CompareDocumentOnlineResponse), response, err
}

/* WordsApiService Compress and resize images inside the document. The default settings allows to reduce the size of the document without any visible degradation of images quality.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.CompressResponse*/
func (a *WordsApiService) CompressDocument(ctx context.Context, data *models.CompressDocumentRequest) (models.CompressResponse, *http.Response, error) {
    var (
        successPayload models.CompressResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Compress and resize images inside the document. The default settings allows to reduce the size of the document without any visible degradation of images quality.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return CompressDocumentOnlineResponse*/
func (a *WordsApiService) CompressDocumentOnline(ctx context.Context, data *models.CompressDocumentOnlineRequest) (models.CompressDocumentOnlineResponse, *http.Response, error) {
    var (
        successPayload models.CompressDocumentOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.CompressDocumentOnlineResponse), response, err
}

/* WordsApiService Converts a document on a local drive to the specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return http.Response.Body contains binary result data*/
func (a *WordsApiService) ConvertDocument(ctx context.Context, data *models.ConvertDocumentRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Copy file.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) CopyFile(ctx context.Context, data *models.CopyFileRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Copy folder.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) CopyFolder(ctx context.Context, data *models.CopyFolderRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Makes a copy of the style in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.StyleResponse*/
func (a *WordsApiService) CopyStyle(ctx context.Context, data *models.CopyStyleRequest) (models.StyleResponse, *http.Response, error) {
    var (
        successPayload models.StyleResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Makes a copy of the style in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return CopyStyleOnlineResponse*/
func (a *WordsApiService) CopyStyleOnline(ctx context.Context, data *models.CopyStyleOnlineRequest) (models.CopyStyleOnlineResponse, *http.Response, error) {
    var (
        successPayload models.CopyStyleOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.CopyStyleOnlineResponse), response, err
}

/* WordsApiService Copies styles from the origin document to the target document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.WordsResponse*/
func (a *WordsApiService) CopyStylesFromTemplate(ctx context.Context, data *models.CopyStylesFromTemplateRequest) (models.WordsResponse, *http.Response, error) {
    var (
        successPayload models.WordsResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Creates a new document in cloud storage in the format, determined by the file extension. Supported all save format extensions.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.DocumentResponse*/
func (a *WordsApiService) CreateDocument(ctx context.Context, data *models.CreateDocumentRequest) (models.DocumentResponse, *http.Response, error) {
    var (
        successPayload models.DocumentResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Create the folder.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) CreateFolder(ctx context.Context, data *models.CreateFolderRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Adds a new or updates an existing document property.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.DocumentPropertyResponse*/
func (a *WordsApiService) CreateOrUpdateDocumentProperty(ctx context.Context, data *models.CreateOrUpdateDocumentPropertyRequest) (models.DocumentPropertyResponse, *http.Response, error) {
    var (
        successPayload models.DocumentPropertyResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Adds a new or updates an existing document property.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return CreateOrUpdateDocumentPropertyOnlineResponse*/
func (a *WordsApiService) CreateOrUpdateDocumentPropertyOnline(ctx context.Context, data *models.CreateOrUpdateDocumentPropertyOnlineRequest) (models.CreateOrUpdateDocumentPropertyOnlineResponse, *http.Response, error) {
    var (
        successPayload models.CreateOrUpdateDocumentPropertyOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.CreateOrUpdateDocumentPropertyOnlineResponse), response, err
}

/* WordsApiService Removes paragraph tab stops from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.TabStopsResponse*/
func (a *WordsApiService) DeleteAllParagraphTabStops(ctx context.Context, data *models.DeleteAllParagraphTabStopsRequest) (models.TabStopsResponse, *http.Response, error) {
    var (
        successPayload models.TabStopsResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Removes paragraph tab stops from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return DeleteAllParagraphTabStopsOnlineResponse*/
func (a *WordsApiService) DeleteAllParagraphTabStopsOnline(ctx context.Context, data *models.DeleteAllParagraphTabStopsOnlineRequest) (models.DeleteAllParagraphTabStopsOnlineResponse, *http.Response, error) {
    var (
        successPayload models.DeleteAllParagraphTabStopsOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.DeleteAllParagraphTabStopsOnlineResponse), response, err
}

/* WordsApiService Removes a bookmark from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) DeleteBookmark(ctx context.Context, data *models.DeleteBookmarkRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Removes a bookmark from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return map[string]io.Reader*/
func (a *WordsApiService) DeleteBookmarkOnline(ctx context.Context, data *models.DeleteBookmarkOnlineRequest) (map[string]io.Reader, *http.Response, error) {
    var (
        successPayload map[string]io.Reader
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    if err != nil {
        return successPayload, response, err
    }

    successPayload, err = models.ParseFilesCollection(response)
    return successPayload, response, err
}

/* WordsApiService Removes all bookmarks from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) DeleteBookmarks(ctx context.Context, data *models.DeleteBookmarksRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Removes all bookmarks from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return map[string]io.Reader*/
func (a *WordsApiService) DeleteBookmarksOnline(ctx context.Context, data *models.DeleteBookmarksOnlineRequest) (map[string]io.Reader, *http.Response, error) {
    var (
        successPayload map[string]io.Reader
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    if err != nil {
        return successPayload, response, err
    }

    successPayload, err = models.ParseFilesCollection(response)
    return successPayload, response, err
}

/* WordsApiService Removes a border from the document node. The 'nodePath' parameter should refer to a paragraph, a cell or a row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.BorderResponse*/
func (a *WordsApiService) DeleteBorder(ctx context.Context, data *models.DeleteBorderRequest) (models.BorderResponse, *http.Response, error) {
    var (
        successPayload models.BorderResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Removes a border from the document node. The 'nodePath' parameter should refer to a paragraph, a cell or a row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return DeleteBorderOnlineResponse*/
func (a *WordsApiService) DeleteBorderOnline(ctx context.Context, data *models.DeleteBorderOnlineRequest) (models.DeleteBorderOnlineResponse, *http.Response, error) {
    var (
        successPayload models.DeleteBorderOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.DeleteBorderOnlineResponse), response, err
}

/* WordsApiService Removes borders from the document node. The 'nodePath' parameter should refer to a paragraph, a cell or a row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.BordersResponse*/
func (a *WordsApiService) DeleteBorders(ctx context.Context, data *models.DeleteBordersRequest) (models.BordersResponse, *http.Response, error) {
    var (
        successPayload models.BordersResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Removes borders from the document node. The 'nodePath' parameter should refer to a paragraph, a cell or a row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return DeleteBordersOnlineResponse*/
func (a *WordsApiService) DeleteBordersOnline(ctx context.Context, data *models.DeleteBordersOnlineRequest) (models.DeleteBordersOnlineResponse, *http.Response, error) {
    var (
        successPayload models.DeleteBordersOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.DeleteBordersOnlineResponse), response, err
}

/* WordsApiService Removes a comment from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) DeleteComment(ctx context.Context, data *models.DeleteCommentRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Removes a comment from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return map[string]io.Reader*/
func (a *WordsApiService) DeleteCommentOnline(ctx context.Context, data *models.DeleteCommentOnlineRequest) (map[string]io.Reader, *http.Response, error) {
    var (
        successPayload map[string]io.Reader
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    if err != nil {
        return successPayload, response, err
    }

    successPayload, err = models.ParseFilesCollection(response)
    return successPayload, response, err
}

/* WordsApiService Removes all comments from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) DeleteComments(ctx context.Context, data *models.DeleteCommentsRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Removes all comments from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return map[string]io.Reader*/
func (a *WordsApiService) DeleteCommentsOnline(ctx context.Context, data *models.DeleteCommentsOnlineRequest) (map[string]io.Reader, *http.Response, error) {
    var (
        successPayload map[string]io.Reader
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    if err != nil {
        return successPayload, response, err
    }

    successPayload, err = models.ParseFilesCollection(response)
    return successPayload, response, err
}

/* WordsApiService Removes the custom xml part from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) DeleteCustomXmlPart(ctx context.Context, data *models.DeleteCustomXmlPartRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Removes the custom xml part from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return map[string]io.Reader*/
func (a *WordsApiService) DeleteCustomXmlPartOnline(ctx context.Context, data *models.DeleteCustomXmlPartOnlineRequest) (map[string]io.Reader, *http.Response, error) {
    var (
        successPayload map[string]io.Reader
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    if err != nil {
        return successPayload, response, err
    }

    successPayload, err = models.ParseFilesCollection(response)
    return successPayload, response, err
}

/* WordsApiService Removes all custom xml parts from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) DeleteCustomXmlParts(ctx context.Context, data *models.DeleteCustomXmlPartsRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Removes all custom xml parts from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return map[string]io.Reader*/
func (a *WordsApiService) DeleteCustomXmlPartsOnline(ctx context.Context, data *models.DeleteCustomXmlPartsOnlineRequest) (map[string]io.Reader, *http.Response, error) {
    var (
        successPayload map[string]io.Reader
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    if err != nil {
        return successPayload, response, err
    }

    successPayload, err = models.ParseFilesCollection(response)
    return successPayload, response, err
}

/* WordsApiService Removes a document property.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) DeleteDocumentProperty(ctx context.Context, data *models.DeleteDocumentPropertyRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Removes a document property.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return map[string]io.Reader*/
func (a *WordsApiService) DeleteDocumentPropertyOnline(ctx context.Context, data *models.DeleteDocumentPropertyOnlineRequest) (map[string]io.Reader, *http.Response, error) {
    var (
        successPayload map[string]io.Reader
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    if err != nil {
        return successPayload, response, err
    }

    successPayload, err = models.ParseFilesCollection(response)
    return successPayload, response, err
}

/* WordsApiService Removes a DrawingObject from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) DeleteDrawingObject(ctx context.Context, data *models.DeleteDrawingObjectRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Removes a DrawingObject from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return map[string]io.Reader*/
func (a *WordsApiService) DeleteDrawingObjectOnline(ctx context.Context, data *models.DeleteDrawingObjectOnlineRequest) (map[string]io.Reader, *http.Response, error) {
    var (
        successPayload map[string]io.Reader
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    if err != nil {
        return successPayload, response, err
    }

    successPayload, err = models.ParseFilesCollection(response)
    return successPayload, response, err
}

/* WordsApiService Removes a field from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) DeleteField(ctx context.Context, data *models.DeleteFieldRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Removes a field from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return map[string]io.Reader*/
func (a *WordsApiService) DeleteFieldOnline(ctx context.Context, data *models.DeleteFieldOnlineRequest) (map[string]io.Reader, *http.Response, error) {
    var (
        successPayload map[string]io.Reader
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    if err != nil {
        return successPayload, response, err
    }

    successPayload, err = models.ParseFilesCollection(response)
    return successPayload, response, err
}

/* WordsApiService Removes fields from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) DeleteFields(ctx context.Context, data *models.DeleteFieldsRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Removes fields from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return map[string]io.Reader*/
func (a *WordsApiService) DeleteFieldsOnline(ctx context.Context, data *models.DeleteFieldsOnlineRequest) (map[string]io.Reader, *http.Response, error) {
    var (
        successPayload map[string]io.Reader
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    if err != nil {
        return successPayload, response, err
    }

    successPayload, err = models.ParseFilesCollection(response)
    return successPayload, response, err
}

/* WordsApiService Delete file.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) DeleteFile(ctx context.Context, data *models.DeleteFileRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Delete folder.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) DeleteFolder(ctx context.Context, data *models.DeleteFolderRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Removes a footnote from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) DeleteFootnote(ctx context.Context, data *models.DeleteFootnoteRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Removes a footnote from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return map[string]io.Reader*/
func (a *WordsApiService) DeleteFootnoteOnline(ctx context.Context, data *models.DeleteFootnoteOnlineRequest) (map[string]io.Reader, *http.Response, error) {
    var (
        successPayload map[string]io.Reader
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    if err != nil {
        return successPayload, response, err
    }

    successPayload, err = models.ParseFilesCollection(response)
    return successPayload, response, err
}

/* WordsApiService Removes a form field from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) DeleteFormField(ctx context.Context, data *models.DeleteFormFieldRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Removes a form field from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return map[string]io.Reader*/
func (a *WordsApiService) DeleteFormFieldOnline(ctx context.Context, data *models.DeleteFormFieldOnlineRequest) (map[string]io.Reader, *http.Response, error) {
    var (
        successPayload map[string]io.Reader
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    if err != nil {
        return successPayload, response, err
    }

    successPayload, err = models.ParseFilesCollection(response)
    return successPayload, response, err
}

/* WordsApiService Removes a HeaderFooter object from the document section.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) DeleteHeaderFooter(ctx context.Context, data *models.DeleteHeaderFooterRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Removes a HeaderFooter object from the document section.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return map[string]io.Reader*/
func (a *WordsApiService) DeleteHeaderFooterOnline(ctx context.Context, data *models.DeleteHeaderFooterOnlineRequest) (map[string]io.Reader, *http.Response, error) {
    var (
        successPayload map[string]io.Reader
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    if err != nil {
        return successPayload, response, err
    }

    successPayload, err = models.ParseFilesCollection(response)
    return successPayload, response, err
}

/* WordsApiService Removes HeaderFooter objects from the document section.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) DeleteHeadersFooters(ctx context.Context, data *models.DeleteHeadersFootersRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Removes HeaderFooter objects from the document section.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return map[string]io.Reader*/
func (a *WordsApiService) DeleteHeadersFootersOnline(ctx context.Context, data *models.DeleteHeadersFootersOnlineRequest) (map[string]io.Reader, *http.Response, error) {
    var (
        successPayload map[string]io.Reader
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    if err != nil {
        return successPayload, response, err
    }

    successPayload, err = models.ParseFilesCollection(response)
    return successPayload, response, err
}

/* WordsApiService Removes macros from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) DeleteMacros(ctx context.Context, data *models.DeleteMacrosRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Removes macros from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return map[string]io.Reader*/
func (a *WordsApiService) DeleteMacrosOnline(ctx context.Context, data *models.DeleteMacrosOnlineRequest) (map[string]io.Reader, *http.Response, error) {
    var (
        successPayload map[string]io.Reader
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    if err != nil {
        return successPayload, response, err
    }

    successPayload, err = models.ParseFilesCollection(response)
    return successPayload, response, err
}

/* WordsApiService Removes an OfficeMath object from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) DeleteOfficeMathObject(ctx context.Context, data *models.DeleteOfficeMathObjectRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Removes an OfficeMath object from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return map[string]io.Reader*/
func (a *WordsApiService) DeleteOfficeMathObjectOnline(ctx context.Context, data *models.DeleteOfficeMathObjectOnlineRequest) (map[string]io.Reader, *http.Response, error) {
    var (
        successPayload map[string]io.Reader
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    if err != nil {
        return successPayload, response, err
    }

    successPayload, err = models.ParseFilesCollection(response)
    return successPayload, response, err
}

/* WordsApiService Removes all office math objects from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) DeleteOfficeMathObjects(ctx context.Context, data *models.DeleteOfficeMathObjectsRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Removes all office math objects from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return map[string]io.Reader*/
func (a *WordsApiService) DeleteOfficeMathObjectsOnline(ctx context.Context, data *models.DeleteOfficeMathObjectsOnlineRequest) (map[string]io.Reader, *http.Response, error) {
    var (
        successPayload map[string]io.Reader
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    if err != nil {
        return successPayload, response, err
    }

    successPayload, err = models.ParseFilesCollection(response)
    return successPayload, response, err
}

/* WordsApiService Removes a paragraph from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) DeleteParagraph(ctx context.Context, data *models.DeleteParagraphRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Removes the formatting properties of a paragraph list from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.ParagraphListFormatResponse*/
func (a *WordsApiService) DeleteParagraphListFormat(ctx context.Context, data *models.DeleteParagraphListFormatRequest) (models.ParagraphListFormatResponse, *http.Response, error) {
    var (
        successPayload models.ParagraphListFormatResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Removes the formatting properties of a paragraph list from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return DeleteParagraphListFormatOnlineResponse*/
func (a *WordsApiService) DeleteParagraphListFormatOnline(ctx context.Context, data *models.DeleteParagraphListFormatOnlineRequest) (models.DeleteParagraphListFormatOnlineResponse, *http.Response, error) {
    var (
        successPayload models.DeleteParagraphListFormatOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.DeleteParagraphListFormatOnlineResponse), response, err
}

/* WordsApiService Removes a paragraph from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return map[string]io.Reader*/
func (a *WordsApiService) DeleteParagraphOnline(ctx context.Context, data *models.DeleteParagraphOnlineRequest) (map[string]io.Reader, *http.Response, error) {
    var (
        successPayload map[string]io.Reader
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    if err != nil {
        return successPayload, response, err
    }

    successPayload, err = models.ParseFilesCollection(response)
    return successPayload, response, err
}

/* WordsApiService Removes a paragraph tab stop from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.TabStopsResponse*/
func (a *WordsApiService) DeleteParagraphTabStop(ctx context.Context, data *models.DeleteParagraphTabStopRequest) (models.TabStopsResponse, *http.Response, error) {
    var (
        successPayload models.TabStopsResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Removes a paragraph tab stop from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return DeleteParagraphTabStopOnlineResponse*/
func (a *WordsApiService) DeleteParagraphTabStopOnline(ctx context.Context, data *models.DeleteParagraphTabStopOnlineRequest) (models.DeleteParagraphTabStopOnlineResponse, *http.Response, error) {
    var (
        successPayload models.DeleteParagraphTabStopOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.DeleteParagraphTabStopOnlineResponse), response, err
}

/* WordsApiService Removes a Run object from the paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) DeleteRun(ctx context.Context, data *models.DeleteRunRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Removes a Run object from the paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return map[string]io.Reader*/
func (a *WordsApiService) DeleteRunOnline(ctx context.Context, data *models.DeleteRunOnlineRequest) (map[string]io.Reader, *http.Response, error) {
    var (
        successPayload map[string]io.Reader
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    if err != nil {
        return successPayload, response, err
    }

    successPayload, err = models.ParseFilesCollection(response)
    return successPayload, response, err
}

/* WordsApiService Removes a section from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) DeleteSection(ctx context.Context, data *models.DeleteSectionRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Removes a section from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return map[string]io.Reader*/
func (a *WordsApiService) DeleteSectionOnline(ctx context.Context, data *models.DeleteSectionOnlineRequest) (map[string]io.Reader, *http.Response, error) {
    var (
        successPayload map[string]io.Reader
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    if err != nil {
        return successPayload, response, err
    }

    successPayload, err = models.ParseFilesCollection(response)
    return successPayload, response, err
}

/* WordsApiService Removes a StructuredDocumentTag (SDT) from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) DeleteStructuredDocumentTag(ctx context.Context, data *models.DeleteStructuredDocumentTagRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Removes a StructuredDocumentTag (SDT) from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return map[string]io.Reader*/
func (a *WordsApiService) DeleteStructuredDocumentTagOnline(ctx context.Context, data *models.DeleteStructuredDocumentTagOnlineRequest) (map[string]io.Reader, *http.Response, error) {
    var (
        successPayload map[string]io.Reader
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    if err != nil {
        return successPayload, response, err
    }

    successPayload, err = models.ParseFilesCollection(response)
    return successPayload, response, err
}

/* WordsApiService Removes a table from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) DeleteTable(ctx context.Context, data *models.DeleteTableRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Removes a cell from the table row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) DeleteTableCell(ctx context.Context, data *models.DeleteTableCellRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Removes a cell from the table row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return map[string]io.Reader*/
func (a *WordsApiService) DeleteTableCellOnline(ctx context.Context, data *models.DeleteTableCellOnlineRequest) (map[string]io.Reader, *http.Response, error) {
    var (
        successPayload map[string]io.Reader
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    if err != nil {
        return successPayload, response, err
    }

    successPayload, err = models.ParseFilesCollection(response)
    return successPayload, response, err
}

/* WordsApiService Removes a table from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return map[string]io.Reader*/
func (a *WordsApiService) DeleteTableOnline(ctx context.Context, data *models.DeleteTableOnlineRequest) (map[string]io.Reader, *http.Response, error) {
    var (
        successPayload map[string]io.Reader
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    if err != nil {
        return successPayload, response, err
    }

    successPayload, err = models.ParseFilesCollection(response)
    return successPayload, response, err
}

/* WordsApiService Removes a row from the table.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) DeleteTableRow(ctx context.Context, data *models.DeleteTableRowRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Removes a row from the table.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return map[string]io.Reader*/
func (a *WordsApiService) DeleteTableRowOnline(ctx context.Context, data *models.DeleteTableRowOnlineRequest) (map[string]io.Reader, *http.Response, error) {
    var (
        successPayload map[string]io.Reader
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    if err != nil {
        return successPayload, response, err
    }

    successPayload, err = models.ParseFilesCollection(response)
    return successPayload, response, err
}

/* WordsApiService Removes a watermark from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.DocumentResponse*/
func (a *WordsApiService) DeleteWatermark(ctx context.Context, data *models.DeleteWatermarkRequest) (models.DocumentResponse, *http.Response, error) {
    var (
        successPayload models.DocumentResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Removes a watermark from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return DeleteWatermarkOnlineResponse*/
func (a *WordsApiService) DeleteWatermarkOnline(ctx context.Context, data *models.DeleteWatermarkOnlineRequest) (models.DeleteWatermarkOnlineResponse, *http.Response, error) {
    var (
        successPayload models.DeleteWatermarkOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.DeleteWatermarkOnlineResponse), response, err
}

/* WordsApiService Download file.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return http.Response.Body contains binary result data*/
func (a *WordsApiService) DownloadFile(ctx context.Context, data *models.DownloadFileRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Executes a Mail Merge operation.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.DocumentResponse*/
func (a *WordsApiService) ExecuteMailMerge(ctx context.Context, data *models.ExecuteMailMergeRequest) (models.DocumentResponse, *http.Response, error) {
    var (
        successPayload models.DocumentResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Executes a Mail Merge operation online.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return http.Response.Body contains binary result data*/
func (a *WordsApiService) ExecuteMailMergeOnline(ctx context.Context, data *models.ExecuteMailMergeOnlineRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Get all information about revisions.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.RevisionsResponse*/
func (a *WordsApiService) GetAllRevisions(ctx context.Context, data *models.GetAllRevisionsRequest) (models.RevisionsResponse, *http.Response, error) {
    var (
        successPayload models.RevisionsResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Get all information about revisions.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.RevisionsResponse*/
func (a *WordsApiService) GetAllRevisionsOnline(ctx context.Context, data *models.GetAllRevisionsOnlineRequest) (models.RevisionsResponse, *http.Response, error) {
    var (
        successPayload models.RevisionsResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads available fonts from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.AvailableFontsResponse*/
func (a *WordsApiService) GetAvailableFonts(ctx context.Context, data *models.GetAvailableFontsRequest) (models.AvailableFontsResponse, *http.Response, error) {
    var (
        successPayload models.AvailableFontsResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a bookmark, specified by name, from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.BookmarkResponse*/
func (a *WordsApiService) GetBookmarkByName(ctx context.Context, data *models.GetBookmarkByNameRequest) (models.BookmarkResponse, *http.Response, error) {
    var (
        successPayload models.BookmarkResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a bookmark, specified by name, from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.BookmarkResponse*/
func (a *WordsApiService) GetBookmarkByNameOnline(ctx context.Context, data *models.GetBookmarkByNameOnlineRequest) (models.BookmarkResponse, *http.Response, error) {
    var (
        successPayload models.BookmarkResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads bookmarks from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.BookmarksResponse*/
func (a *WordsApiService) GetBookmarks(ctx context.Context, data *models.GetBookmarksRequest) (models.BookmarksResponse, *http.Response, error) {
    var (
        successPayload models.BookmarksResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads bookmarks from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.BookmarksResponse*/
func (a *WordsApiService) GetBookmarksOnline(ctx context.Context, data *models.GetBookmarksOnlineRequest) (models.BookmarksResponse, *http.Response, error) {
    var (
        successPayload models.BookmarksResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a border from the document node. The 'nodePath' parameter should refer to a paragraph, a cell or a row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.BorderResponse*/
func (a *WordsApiService) GetBorder(ctx context.Context, data *models.GetBorderRequest) (models.BorderResponse, *http.Response, error) {
    var (
        successPayload models.BorderResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a border from the document node. The 'nodePath' parameter should refer to a paragraph, a cell or a row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.BorderResponse*/
func (a *WordsApiService) GetBorderOnline(ctx context.Context, data *models.GetBorderOnlineRequest) (models.BorderResponse, *http.Response, error) {
    var (
        successPayload models.BorderResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads borders from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.BordersResponse*/
func (a *WordsApiService) GetBorders(ctx context.Context, data *models.GetBordersRequest) (models.BordersResponse, *http.Response, error) {
    var (
        successPayload models.BordersResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads borders from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.BordersResponse*/
func (a *WordsApiService) GetBordersOnline(ctx context.Context, data *models.GetBordersOnlineRequest) (models.BordersResponse, *http.Response, error) {
    var (
        successPayload models.BordersResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a comment from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.CommentResponse*/
func (a *WordsApiService) GetComment(ctx context.Context, data *models.GetCommentRequest) (models.CommentResponse, *http.Response, error) {
    var (
        successPayload models.CommentResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a comment from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.CommentResponse*/
func (a *WordsApiService) GetCommentOnline(ctx context.Context, data *models.GetCommentOnlineRequest) (models.CommentResponse, *http.Response, error) {
    var (
        successPayload models.CommentResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads comments from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.CommentsResponse*/
func (a *WordsApiService) GetComments(ctx context.Context, data *models.GetCommentsRequest) (models.CommentsResponse, *http.Response, error) {
    var (
        successPayload models.CommentsResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads comments from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.CommentsResponse*/
func (a *WordsApiService) GetCommentsOnline(ctx context.Context, data *models.GetCommentsOnlineRequest) (models.CommentsResponse, *http.Response, error) {
    var (
        successPayload models.CommentsResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads the custom xml part from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.CustomXmlPartResponse*/
func (a *WordsApiService) GetCustomXmlPart(ctx context.Context, data *models.GetCustomXmlPartRequest) (models.CustomXmlPartResponse, *http.Response, error) {
    var (
        successPayload models.CustomXmlPartResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads the custom xml part from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.CustomXmlPartResponse*/
func (a *WordsApiService) GetCustomXmlPartOnline(ctx context.Context, data *models.GetCustomXmlPartOnlineRequest) (models.CustomXmlPartResponse, *http.Response, error) {
    var (
        successPayload models.CustomXmlPartResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads custom xml parts from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.CustomXmlPartsResponse*/
func (a *WordsApiService) GetCustomXmlParts(ctx context.Context, data *models.GetCustomXmlPartsRequest) (models.CustomXmlPartsResponse, *http.Response, error) {
    var (
        successPayload models.CustomXmlPartsResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads custom xml parts from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.CustomXmlPartsResponse*/
func (a *WordsApiService) GetCustomXmlPartsOnline(ctx context.Context, data *models.GetCustomXmlPartsOnlineRequest) (models.CustomXmlPartsResponse, *http.Response, error) {
    var (
        successPayload models.CustomXmlPartsResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads common information from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.DocumentResponse*/
func (a *WordsApiService) GetDocument(ctx context.Context, data *models.GetDocumentRequest) (models.DocumentResponse, *http.Response, error) {
    var (
        successPayload models.DocumentResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a DrawingObject from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.DrawingObjectResponse*/
func (a *WordsApiService) GetDocumentDrawingObjectByIndex(ctx context.Context, data *models.GetDocumentDrawingObjectByIndexRequest) (models.DrawingObjectResponse, *http.Response, error) {
    var (
        successPayload models.DrawingObjectResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a DrawingObject from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.DrawingObjectResponse*/
func (a *WordsApiService) GetDocumentDrawingObjectByIndexOnline(ctx context.Context, data *models.GetDocumentDrawingObjectByIndexOnlineRequest) (models.DrawingObjectResponse, *http.Response, error) {
    var (
        successPayload models.DrawingObjectResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads image data of a DrawingObject from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return http.Response.Body contains binary result data*/
func (a *WordsApiService) GetDocumentDrawingObjectImageData(ctx context.Context, data *models.GetDocumentDrawingObjectImageDataRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Reads image data of a DrawingObject from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return http.Response.Body contains binary result data*/
func (a *WordsApiService) GetDocumentDrawingObjectImageDataOnline(ctx context.Context, data *models.GetDocumentDrawingObjectImageDataOnlineRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Reads OLE data of a DrawingObject from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return http.Response.Body contains binary result data*/
func (a *WordsApiService) GetDocumentDrawingObjectOleData(ctx context.Context, data *models.GetDocumentDrawingObjectOleDataRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Reads OLE data of a DrawingObject from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return http.Response.Body contains binary result data*/
func (a *WordsApiService) GetDocumentDrawingObjectOleDataOnline(ctx context.Context, data *models.GetDocumentDrawingObjectOleDataOnlineRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Reads DrawingObjects from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.DrawingObjectsResponse*/
func (a *WordsApiService) GetDocumentDrawingObjects(ctx context.Context, data *models.GetDocumentDrawingObjectsRequest) (models.DrawingObjectsResponse, *http.Response, error) {
    var (
        successPayload models.DrawingObjectsResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads DrawingObjects from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.DrawingObjectsResponse*/
func (a *WordsApiService) GetDocumentDrawingObjectsOnline(ctx context.Context, data *models.GetDocumentDrawingObjectsOnlineRequest) (models.DrawingObjectsResponse, *http.Response, error) {
    var (
        successPayload models.DrawingObjectsResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads merge field names from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.FieldNamesResponse*/
func (a *WordsApiService) GetDocumentFieldNames(ctx context.Context, data *models.GetDocumentFieldNamesRequest) (models.FieldNamesResponse, *http.Response, error) {
    var (
        successPayload models.FieldNamesResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads merge field names from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.FieldNamesResponse*/
func (a *WordsApiService) GetDocumentFieldNamesOnline(ctx context.Context, data *models.GetDocumentFieldNamesOnlineRequest) (models.FieldNamesResponse, *http.Response, error) {
    var (
        successPayload models.FieldNamesResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a hyperlink from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.HyperlinkResponse*/
func (a *WordsApiService) GetDocumentHyperlinkByIndex(ctx context.Context, data *models.GetDocumentHyperlinkByIndexRequest) (models.HyperlinkResponse, *http.Response, error) {
    var (
        successPayload models.HyperlinkResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a hyperlink from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.HyperlinkResponse*/
func (a *WordsApiService) GetDocumentHyperlinkByIndexOnline(ctx context.Context, data *models.GetDocumentHyperlinkByIndexOnlineRequest) (models.HyperlinkResponse, *http.Response, error) {
    var (
        successPayload models.HyperlinkResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads hyperlinks from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.HyperlinksResponse*/
func (a *WordsApiService) GetDocumentHyperlinks(ctx context.Context, data *models.GetDocumentHyperlinksRequest) (models.HyperlinksResponse, *http.Response, error) {
    var (
        successPayload models.HyperlinksResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads hyperlinks from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.HyperlinksResponse*/
func (a *WordsApiService) GetDocumentHyperlinksOnline(ctx context.Context, data *models.GetDocumentHyperlinksOnlineRequest) (models.HyperlinksResponse, *http.Response, error) {
    var (
        successPayload models.HyperlinksResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads document properties.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.DocumentPropertiesResponse*/
func (a *WordsApiService) GetDocumentProperties(ctx context.Context, data *models.GetDocumentPropertiesRequest) (models.DocumentPropertiesResponse, *http.Response, error) {
    var (
        successPayload models.DocumentPropertiesResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads document properties.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.DocumentPropertiesResponse*/
func (a *WordsApiService) GetDocumentPropertiesOnline(ctx context.Context, data *models.GetDocumentPropertiesOnlineRequest) (models.DocumentPropertiesResponse, *http.Response, error) {
    var (
        successPayload models.DocumentPropertiesResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a document property.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.DocumentPropertyResponse*/
func (a *WordsApiService) GetDocumentProperty(ctx context.Context, data *models.GetDocumentPropertyRequest) (models.DocumentPropertyResponse, *http.Response, error) {
    var (
        successPayload models.DocumentPropertyResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a document property.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.DocumentPropertyResponse*/
func (a *WordsApiService) GetDocumentPropertyOnline(ctx context.Context, data *models.GetDocumentPropertyOnlineRequest) (models.DocumentPropertyResponse, *http.Response, error) {
    var (
        successPayload models.DocumentPropertyResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads protection properties from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.ProtectionDataResponse*/
func (a *WordsApiService) GetDocumentProtection(ctx context.Context, data *models.GetDocumentProtectionRequest) (models.ProtectionDataResponse, *http.Response, error) {
    var (
        successPayload models.ProtectionDataResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads protection properties from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.ProtectionDataResponse*/
func (a *WordsApiService) GetDocumentProtectionOnline(ctx context.Context, data *models.GetDocumentProtectionOnlineRequest) (models.ProtectionDataResponse, *http.Response, error) {
    var (
        successPayload models.ProtectionDataResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads document statistics.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.StatDataResponse*/
func (a *WordsApiService) GetDocumentStatistics(ctx context.Context, data *models.GetDocumentStatisticsRequest) (models.StatDataResponse, *http.Response, error) {
    var (
        successPayload models.StatDataResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads document statistics.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.StatDataResponse*/
func (a *WordsApiService) GetDocumentStatisticsOnline(ctx context.Context, data *models.GetDocumentStatisticsOnlineRequest) (models.StatDataResponse, *http.Response, error) {
    var (
        successPayload models.StatDataResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Converts a document in cloud storage to the specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return http.Response.Body contains binary result data*/
func (a *WordsApiService) GetDocumentWithFormat(ctx context.Context, data *models.GetDocumentWithFormatRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Reads a field from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.FieldResponse*/
func (a *WordsApiService) GetField(ctx context.Context, data *models.GetFieldRequest) (models.FieldResponse, *http.Response, error) {
    var (
        successPayload models.FieldResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a field from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.FieldResponse*/
func (a *WordsApiService) GetFieldOnline(ctx context.Context, data *models.GetFieldOnlineRequest) (models.FieldResponse, *http.Response, error) {
    var (
        successPayload models.FieldResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads fields from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.FieldsResponse*/
func (a *WordsApiService) GetFields(ctx context.Context, data *models.GetFieldsRequest) (models.FieldsResponse, *http.Response, error) {
    var (
        successPayload models.FieldsResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads fields from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.FieldsResponse*/
func (a *WordsApiService) GetFieldsOnline(ctx context.Context, data *models.GetFieldsOnlineRequest) (models.FieldsResponse, *http.Response, error) {
    var (
        successPayload models.FieldsResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Get all files and folders within a folder.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.FilesList*/
func (a *WordsApiService) GetFilesList(ctx context.Context, data *models.GetFilesListRequest) (models.FilesList, *http.Response, error) {
    var (
        successPayload models.FilesList
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a footnote from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.FootnoteResponse*/
func (a *WordsApiService) GetFootnote(ctx context.Context, data *models.GetFootnoteRequest) (models.FootnoteResponse, *http.Response, error) {
    var (
        successPayload models.FootnoteResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a footnote from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.FootnoteResponse*/
func (a *WordsApiService) GetFootnoteOnline(ctx context.Context, data *models.GetFootnoteOnlineRequest) (models.FootnoteResponse, *http.Response, error) {
    var (
        successPayload models.FootnoteResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads footnotes from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.FootnotesResponse*/
func (a *WordsApiService) GetFootnotes(ctx context.Context, data *models.GetFootnotesRequest) (models.FootnotesResponse, *http.Response, error) {
    var (
        successPayload models.FootnotesResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads footnotes from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.FootnotesResponse*/
func (a *WordsApiService) GetFootnotesOnline(ctx context.Context, data *models.GetFootnotesOnlineRequest) (models.FootnotesResponse, *http.Response, error) {
    var (
        successPayload models.FootnotesResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a form field from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.FormFieldResponse*/
func (a *WordsApiService) GetFormField(ctx context.Context, data *models.GetFormFieldRequest) (models.FormFieldResponse, *http.Response, error) {
    var (
        successPayload models.FormFieldResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a form field from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.FormFieldResponse*/
func (a *WordsApiService) GetFormFieldOnline(ctx context.Context, data *models.GetFormFieldOnlineRequest) (models.FormFieldResponse, *http.Response, error) {
    var (
        successPayload models.FormFieldResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads form fields from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.FormFieldsResponse*/
func (a *WordsApiService) GetFormFields(ctx context.Context, data *models.GetFormFieldsRequest) (models.FormFieldsResponse, *http.Response, error) {
    var (
        successPayload models.FormFieldsResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads form fields from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.FormFieldsResponse*/
func (a *WordsApiService) GetFormFieldsOnline(ctx context.Context, data *models.GetFormFieldsOnlineRequest) (models.FormFieldsResponse, *http.Response, error) {
    var (
        successPayload models.FormFieldsResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a HeaderFooter object from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.HeaderFooterResponse*/
func (a *WordsApiService) GetHeaderFooter(ctx context.Context, data *models.GetHeaderFooterRequest) (models.HeaderFooterResponse, *http.Response, error) {
    var (
        successPayload models.HeaderFooterResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a HeaderFooter object from the document section.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.HeaderFooterResponse*/
func (a *WordsApiService) GetHeaderFooterOfSection(ctx context.Context, data *models.GetHeaderFooterOfSectionRequest) (models.HeaderFooterResponse, *http.Response, error) {
    var (
        successPayload models.HeaderFooterResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a HeaderFooter object from the document section.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.HeaderFooterResponse*/
func (a *WordsApiService) GetHeaderFooterOfSectionOnline(ctx context.Context, data *models.GetHeaderFooterOfSectionOnlineRequest) (models.HeaderFooterResponse, *http.Response, error) {
    var (
        successPayload models.HeaderFooterResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a HeaderFooter object from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.HeaderFooterResponse*/
func (a *WordsApiService) GetHeaderFooterOnline(ctx context.Context, data *models.GetHeaderFooterOnlineRequest) (models.HeaderFooterResponse, *http.Response, error) {
    var (
        successPayload models.HeaderFooterResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads HeaderFooter objects from the document section.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.HeaderFootersResponse*/
func (a *WordsApiService) GetHeaderFooters(ctx context.Context, data *models.GetHeaderFootersRequest) (models.HeaderFootersResponse, *http.Response, error) {
    var (
        successPayload models.HeaderFootersResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads HeaderFooter objects from the document section.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.HeaderFootersResponse*/
func (a *WordsApiService) GetHeaderFootersOnline(ctx context.Context, data *models.GetHeaderFootersOnlineRequest) (models.HeaderFootersResponse, *http.Response, error) {
    var (
        successPayload models.HeaderFootersResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Returns application info.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.InfoResponse*/
func (a *WordsApiService) GetInfo(ctx context.Context, data *models.GetInfoRequest) (models.InfoResponse, *http.Response, error) {
    var (
        successPayload models.InfoResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a list from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.ListResponse*/
func (a *WordsApiService) GetList(ctx context.Context, data *models.GetListRequest) (models.ListResponse, *http.Response, error) {
    var (
        successPayload models.ListResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a list from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.ListResponse*/
func (a *WordsApiService) GetListOnline(ctx context.Context, data *models.GetListOnlineRequest) (models.ListResponse, *http.Response, error) {
    var (
        successPayload models.ListResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads lists from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.ListsResponse*/
func (a *WordsApiService) GetLists(ctx context.Context, data *models.GetListsRequest) (models.ListsResponse, *http.Response, error) {
    var (
        successPayload models.ListsResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads lists from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.ListsResponse*/
func (a *WordsApiService) GetListsOnline(ctx context.Context, data *models.GetListsOnlineRequest) (models.ListsResponse, *http.Response, error) {
    var (
        successPayload models.ListsResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads an OfficeMath object from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.OfficeMathObjectResponse*/
func (a *WordsApiService) GetOfficeMathObject(ctx context.Context, data *models.GetOfficeMathObjectRequest) (models.OfficeMathObjectResponse, *http.Response, error) {
    var (
        successPayload models.OfficeMathObjectResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads an OfficeMath object from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.OfficeMathObjectResponse*/
func (a *WordsApiService) GetOfficeMathObjectOnline(ctx context.Context, data *models.GetOfficeMathObjectOnlineRequest) (models.OfficeMathObjectResponse, *http.Response, error) {
    var (
        successPayload models.OfficeMathObjectResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads OfficeMath objects from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.OfficeMathObjectsResponse*/
func (a *WordsApiService) GetOfficeMathObjects(ctx context.Context, data *models.GetOfficeMathObjectsRequest) (models.OfficeMathObjectsResponse, *http.Response, error) {
    var (
        successPayload models.OfficeMathObjectsResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads OfficeMath objects from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.OfficeMathObjectsResponse*/
func (a *WordsApiService) GetOfficeMathObjectsOnline(ctx context.Context, data *models.GetOfficeMathObjectsOnlineRequest) (models.OfficeMathObjectsResponse, *http.Response, error) {
    var (
        successPayload models.OfficeMathObjectsResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a paragraph from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.ParagraphResponse*/
func (a *WordsApiService) GetParagraph(ctx context.Context, data *models.GetParagraphRequest) (models.ParagraphResponse, *http.Response, error) {
    var (
        successPayload models.ParagraphResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads the formatting properties of a paragraph from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.ParagraphFormatResponse*/
func (a *WordsApiService) GetParagraphFormat(ctx context.Context, data *models.GetParagraphFormatRequest) (models.ParagraphFormatResponse, *http.Response, error) {
    var (
        successPayload models.ParagraphFormatResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads the formatting properties of a paragraph from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.ParagraphFormatResponse*/
func (a *WordsApiService) GetParagraphFormatOnline(ctx context.Context, data *models.GetParagraphFormatOnlineRequest) (models.ParagraphFormatResponse, *http.Response, error) {
    var (
        successPayload models.ParagraphFormatResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads the formatting properties of a paragraph list from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.ParagraphListFormatResponse*/
func (a *WordsApiService) GetParagraphListFormat(ctx context.Context, data *models.GetParagraphListFormatRequest) (models.ParagraphListFormatResponse, *http.Response, error) {
    var (
        successPayload models.ParagraphListFormatResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads the formatting properties of a paragraph list from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.ParagraphListFormatResponse*/
func (a *WordsApiService) GetParagraphListFormatOnline(ctx context.Context, data *models.GetParagraphListFormatOnlineRequest) (models.ParagraphListFormatResponse, *http.Response, error) {
    var (
        successPayload models.ParagraphListFormatResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a paragraph from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.ParagraphResponse*/
func (a *WordsApiService) GetParagraphOnline(ctx context.Context, data *models.GetParagraphOnlineRequest) (models.ParagraphResponse, *http.Response, error) {
    var (
        successPayload models.ParagraphResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads paragraphs from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.ParagraphLinkCollectionResponse*/
func (a *WordsApiService) GetParagraphs(ctx context.Context, data *models.GetParagraphsRequest) (models.ParagraphLinkCollectionResponse, *http.Response, error) {
    var (
        successPayload models.ParagraphLinkCollectionResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads paragraphs from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.ParagraphLinkCollectionResponse*/
func (a *WordsApiService) GetParagraphsOnline(ctx context.Context, data *models.GetParagraphsOnlineRequest) (models.ParagraphLinkCollectionResponse, *http.Response, error) {
    var (
        successPayload models.ParagraphLinkCollectionResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads paragraph tab stops from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.TabStopsResponse*/
func (a *WordsApiService) GetParagraphTabStops(ctx context.Context, data *models.GetParagraphTabStopsRequest) (models.TabStopsResponse, *http.Response, error) {
    var (
        successPayload models.TabStopsResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads paragraph tab stops from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.TabStopsResponse*/
func (a *WordsApiService) GetParagraphTabStopsOnline(ctx context.Context, data *models.GetParagraphTabStopsOnlineRequest) (models.TabStopsResponse, *http.Response, error) {
    var (
        successPayload models.TabStopsResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Get assymetric public key.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.PublicKeyResponse*/
func (a *WordsApiService) GetPublicKey(ctx context.Context, data *models.GetPublicKeyRequest) (models.PublicKeyResponse, *http.Response, error) {
    var (
        successPayload models.PublicKeyResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads range text from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.RangeTextResponse*/
func (a *WordsApiService) GetRangeText(ctx context.Context, data *models.GetRangeTextRequest) (models.RangeTextResponse, *http.Response, error) {
    var (
        successPayload models.RangeTextResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads range text from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.RangeTextResponse*/
func (a *WordsApiService) GetRangeTextOnline(ctx context.Context, data *models.GetRangeTextOnlineRequest) (models.RangeTextResponse, *http.Response, error) {
    var (
        successPayload models.RangeTextResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a Run object from the paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.RunResponse*/
func (a *WordsApiService) GetRun(ctx context.Context, data *models.GetRunRequest) (models.RunResponse, *http.Response, error) {
    var (
        successPayload models.RunResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads the font properties of a Run object from the paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.FontResponse*/
func (a *WordsApiService) GetRunFont(ctx context.Context, data *models.GetRunFontRequest) (models.FontResponse, *http.Response, error) {
    var (
        successPayload models.FontResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads the font properties of a Run object from the paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.FontResponse*/
func (a *WordsApiService) GetRunFontOnline(ctx context.Context, data *models.GetRunFontOnlineRequest) (models.FontResponse, *http.Response, error) {
    var (
        successPayload models.FontResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a Run object from the paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.RunResponse*/
func (a *WordsApiService) GetRunOnline(ctx context.Context, data *models.GetRunOnlineRequest) (models.RunResponse, *http.Response, error) {
    var (
        successPayload models.RunResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads Run objects from the paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.RunsResponse*/
func (a *WordsApiService) GetRuns(ctx context.Context, data *models.GetRunsRequest) (models.RunsResponse, *http.Response, error) {
    var (
        successPayload models.RunsResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads Run objects from the paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.RunsResponse*/
func (a *WordsApiService) GetRunsOnline(ctx context.Context, data *models.GetRunsOnlineRequest) (models.RunsResponse, *http.Response, error) {
    var (
        successPayload models.RunsResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a section from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.SectionResponse*/
func (a *WordsApiService) GetSection(ctx context.Context, data *models.GetSectionRequest) (models.SectionResponse, *http.Response, error) {
    var (
        successPayload models.SectionResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a section from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.SectionResponse*/
func (a *WordsApiService) GetSectionOnline(ctx context.Context, data *models.GetSectionOnlineRequest) (models.SectionResponse, *http.Response, error) {
    var (
        successPayload models.SectionResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads the page setup of a section from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.SectionPageSetupResponse*/
func (a *WordsApiService) GetSectionPageSetup(ctx context.Context, data *models.GetSectionPageSetupRequest) (models.SectionPageSetupResponse, *http.Response, error) {
    var (
        successPayload models.SectionPageSetupResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads the page setup of a section from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.SectionPageSetupResponse*/
func (a *WordsApiService) GetSectionPageSetupOnline(ctx context.Context, data *models.GetSectionPageSetupOnlineRequest) (models.SectionPageSetupResponse, *http.Response, error) {
    var (
        successPayload models.SectionPageSetupResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads sections from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.SectionLinkCollectionResponse*/
func (a *WordsApiService) GetSections(ctx context.Context, data *models.GetSectionsRequest) (models.SectionLinkCollectionResponse, *http.Response, error) {
    var (
        successPayload models.SectionLinkCollectionResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads sections from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.SectionLinkCollectionResponse*/
func (a *WordsApiService) GetSectionsOnline(ctx context.Context, data *models.GetSectionsOnlineRequest) (models.SectionLinkCollectionResponse, *http.Response, error) {
    var (
        successPayload models.SectionLinkCollectionResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Gets signatures from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.SignatureCollectionResponse*/
func (a *WordsApiService) GetSignatures(ctx context.Context, data *models.GetSignaturesRequest) (models.SignatureCollectionResponse, *http.Response, error) {
    var (
        successPayload models.SignatureCollectionResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Gets signatures from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.SignatureCollectionResponse*/
func (a *WordsApiService) GetSignaturesOnline(ctx context.Context, data *models.GetSignaturesOnlineRequest) (models.SignatureCollectionResponse, *http.Response, error) {
    var (
        successPayload models.SignatureCollectionResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a StructuredDocumentTag (SDT) from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.StructuredDocumentTagResponse*/
func (a *WordsApiService) GetStructuredDocumentTag(ctx context.Context, data *models.GetStructuredDocumentTagRequest) (models.StructuredDocumentTagResponse, *http.Response, error) {
    var (
        successPayload models.StructuredDocumentTagResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a StructuredDocumentTag (SDT) from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.StructuredDocumentTagResponse*/
func (a *WordsApiService) GetStructuredDocumentTagOnline(ctx context.Context, data *models.GetStructuredDocumentTagOnlineRequest) (models.StructuredDocumentTagResponse, *http.Response, error) {
    var (
        successPayload models.StructuredDocumentTagResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads StructuredDocumentTags (SDT) from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.StructuredDocumentTagsResponse*/
func (a *WordsApiService) GetStructuredDocumentTags(ctx context.Context, data *models.GetStructuredDocumentTagsRequest) (models.StructuredDocumentTagsResponse, *http.Response, error) {
    var (
        successPayload models.StructuredDocumentTagsResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads StructuredDocumentTags (SDT) from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.StructuredDocumentTagsResponse*/
func (a *WordsApiService) GetStructuredDocumentTagsOnline(ctx context.Context, data *models.GetStructuredDocumentTagsOnlineRequest) (models.StructuredDocumentTagsResponse, *http.Response, error) {
    var (
        successPayload models.StructuredDocumentTagsResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a style from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.StyleResponse*/
func (a *WordsApiService) GetStyle(ctx context.Context, data *models.GetStyleRequest) (models.StyleResponse, *http.Response, error) {
    var (
        successPayload models.StyleResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a style from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.StyleResponse*/
func (a *WordsApiService) GetStyleFromDocumentElement(ctx context.Context, data *models.GetStyleFromDocumentElementRequest) (models.StyleResponse, *http.Response, error) {
    var (
        successPayload models.StyleResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a style from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.StyleResponse*/
func (a *WordsApiService) GetStyleFromDocumentElementOnline(ctx context.Context, data *models.GetStyleFromDocumentElementOnlineRequest) (models.StyleResponse, *http.Response, error) {
    var (
        successPayload models.StyleResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a style from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.StyleResponse*/
func (a *WordsApiService) GetStyleOnline(ctx context.Context, data *models.GetStyleOnlineRequest) (models.StyleResponse, *http.Response, error) {
    var (
        successPayload models.StyleResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads styles from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.StylesResponse*/
func (a *WordsApiService) GetStyles(ctx context.Context, data *models.GetStylesRequest) (models.StylesResponse, *http.Response, error) {
    var (
        successPayload models.StylesResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads styles from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.StylesResponse*/
func (a *WordsApiService) GetStylesOnline(ctx context.Context, data *models.GetStylesOnlineRequest) (models.StylesResponse, *http.Response, error) {
    var (
        successPayload models.StylesResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a table from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.TableResponse*/
func (a *WordsApiService) GetTable(ctx context.Context, data *models.GetTableRequest) (models.TableResponse, *http.Response, error) {
    var (
        successPayload models.TableResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a cell from the table row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.TableCellResponse*/
func (a *WordsApiService) GetTableCell(ctx context.Context, data *models.GetTableCellRequest) (models.TableCellResponse, *http.Response, error) {
    var (
        successPayload models.TableCellResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads the formatting properties of a table cell.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.TableCellFormatResponse*/
func (a *WordsApiService) GetTableCellFormat(ctx context.Context, data *models.GetTableCellFormatRequest) (models.TableCellFormatResponse, *http.Response, error) {
    var (
        successPayload models.TableCellFormatResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads the formatting properties of a table cell.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.TableCellFormatResponse*/
func (a *WordsApiService) GetTableCellFormatOnline(ctx context.Context, data *models.GetTableCellFormatOnlineRequest) (models.TableCellFormatResponse, *http.Response, error) {
    var (
        successPayload models.TableCellFormatResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a cell from the table row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.TableCellResponse*/
func (a *WordsApiService) GetTableCellOnline(ctx context.Context, data *models.GetTableCellOnlineRequest) (models.TableCellResponse, *http.Response, error) {
    var (
        successPayload models.TableCellResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a table from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.TableResponse*/
func (a *WordsApiService) GetTableOnline(ctx context.Context, data *models.GetTableOnlineRequest) (models.TableResponse, *http.Response, error) {
    var (
        successPayload models.TableResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads properties of a table from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.TablePropertiesResponse*/
func (a *WordsApiService) GetTableProperties(ctx context.Context, data *models.GetTablePropertiesRequest) (models.TablePropertiesResponse, *http.Response, error) {
    var (
        successPayload models.TablePropertiesResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads properties of a table from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.TablePropertiesResponse*/
func (a *WordsApiService) GetTablePropertiesOnline(ctx context.Context, data *models.GetTablePropertiesOnlineRequest) (models.TablePropertiesResponse, *http.Response, error) {
    var (
        successPayload models.TablePropertiesResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a row from the table.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.TableRowResponse*/
func (a *WordsApiService) GetTableRow(ctx context.Context, data *models.GetTableRowRequest) (models.TableRowResponse, *http.Response, error) {
    var (
        successPayload models.TableRowResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads the formatting properties of a table row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.TableRowFormatResponse*/
func (a *WordsApiService) GetTableRowFormat(ctx context.Context, data *models.GetTableRowFormatRequest) (models.TableRowFormatResponse, *http.Response, error) {
    var (
        successPayload models.TableRowFormatResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads the formatting properties of a table row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.TableRowFormatResponse*/
func (a *WordsApiService) GetTableRowFormatOnline(ctx context.Context, data *models.GetTableRowFormatOnlineRequest) (models.TableRowFormatResponse, *http.Response, error) {
    var (
        successPayload models.TableRowFormatResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads a row from the table.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.TableRowResponse*/
func (a *WordsApiService) GetTableRowOnline(ctx context.Context, data *models.GetTableRowOnlineRequest) (models.TableRowResponse, *http.Response, error) {
    var (
        successPayload models.TableRowResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads tables from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.TableLinkCollectionResponse*/
func (a *WordsApiService) GetTables(ctx context.Context, data *models.GetTablesRequest) (models.TableLinkCollectionResponse, *http.Response, error) {
    var (
        successPayload models.TableLinkCollectionResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reads tables from the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.TableLinkCollectionResponse*/
func (a *WordsApiService) GetTablesOnline(ctx context.Context, data *models.GetTablesOnlineRequest) (models.TableLinkCollectionResponse, *http.Response, error) {
    var (
        successPayload models.TableLinkCollectionResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Inserts a new bookmark to the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.BookmarkResponse*/
func (a *WordsApiService) InsertBookmark(ctx context.Context, data *models.InsertBookmarkRequest) (models.BookmarkResponse, *http.Response, error) {
    var (
        successPayload models.BookmarkResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Inserts a new bookmark to the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return InsertBookmarkOnlineResponse*/
func (a *WordsApiService) InsertBookmarkOnline(ctx context.Context, data *models.InsertBookmarkOnlineRequest) (models.InsertBookmarkOnlineResponse, *http.Response, error) {
    var (
        successPayload models.InsertBookmarkOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.InsertBookmarkOnlineResponse), response, err
}

/* WordsApiService Inserts a new comment to the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.CommentResponse*/
func (a *WordsApiService) InsertComment(ctx context.Context, data *models.InsertCommentRequest) (models.CommentResponse, *http.Response, error) {
    var (
        successPayload models.CommentResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Inserts a new comment to the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return InsertCommentOnlineResponse*/
func (a *WordsApiService) InsertCommentOnline(ctx context.Context, data *models.InsertCommentOnlineRequest) (models.InsertCommentOnlineResponse, *http.Response, error) {
    var (
        successPayload models.InsertCommentOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.InsertCommentOnlineResponse), response, err
}

/* WordsApiService Inserts a new custom xml part to the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.CustomXmlPartResponse*/
func (a *WordsApiService) InsertCustomXmlPart(ctx context.Context, data *models.InsertCustomXmlPartRequest) (models.CustomXmlPartResponse, *http.Response, error) {
    var (
        successPayload models.CustomXmlPartResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Inserts a new custom xml part to the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return InsertCustomXmlPartOnlineResponse*/
func (a *WordsApiService) InsertCustomXmlPartOnline(ctx context.Context, data *models.InsertCustomXmlPartOnlineRequest) (models.InsertCustomXmlPartOnlineResponse, *http.Response, error) {
    var (
        successPayload models.InsertCustomXmlPartOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.InsertCustomXmlPartOnlineResponse), response, err
}

/* WordsApiService Inserts a new DrawingObject to the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.DrawingObjectResponse*/
func (a *WordsApiService) InsertDrawingObject(ctx context.Context, data *models.InsertDrawingObjectRequest) (models.DrawingObjectResponse, *http.Response, error) {
    var (
        successPayload models.DrawingObjectResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Inserts a new DrawingObject to the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return InsertDrawingObjectOnlineResponse*/
func (a *WordsApiService) InsertDrawingObjectOnline(ctx context.Context, data *models.InsertDrawingObjectOnlineRequest) (models.InsertDrawingObjectOnlineResponse, *http.Response, error) {
    var (
        successPayload models.InsertDrawingObjectOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.InsertDrawingObjectOnlineResponse), response, err
}

/* WordsApiService Inserts a new field to the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.FieldResponse*/
func (a *WordsApiService) InsertField(ctx context.Context, data *models.InsertFieldRequest) (models.FieldResponse, *http.Response, error) {
    var (
        successPayload models.FieldResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Inserts a new field to the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return InsertFieldOnlineResponse*/
func (a *WordsApiService) InsertFieldOnline(ctx context.Context, data *models.InsertFieldOnlineRequest) (models.InsertFieldOnlineResponse, *http.Response, error) {
    var (
        successPayload models.InsertFieldOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.InsertFieldOnlineResponse), response, err
}

/* WordsApiService Inserts a new footnote to the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.FootnoteResponse*/
func (a *WordsApiService) InsertFootnote(ctx context.Context, data *models.InsertFootnoteRequest) (models.FootnoteResponse, *http.Response, error) {
    var (
        successPayload models.FootnoteResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Inserts a new footnote to the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return InsertFootnoteOnlineResponse*/
func (a *WordsApiService) InsertFootnoteOnline(ctx context.Context, data *models.InsertFootnoteOnlineRequest) (models.InsertFootnoteOnlineResponse, *http.Response, error) {
    var (
        successPayload models.InsertFootnoteOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.InsertFootnoteOnlineResponse), response, err
}

/* WordsApiService Inserts a new form field to the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.FormFieldResponse*/
func (a *WordsApiService) InsertFormField(ctx context.Context, data *models.InsertFormFieldRequest) (models.FormFieldResponse, *http.Response, error) {
    var (
        successPayload models.FormFieldResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Inserts a new form field to the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return InsertFormFieldOnlineResponse*/
func (a *WordsApiService) InsertFormFieldOnline(ctx context.Context, data *models.InsertFormFieldOnlineRequest) (models.InsertFormFieldOnlineResponse, *http.Response, error) {
    var (
        successPayload models.InsertFormFieldOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.InsertFormFieldOnlineResponse), response, err
}

/* WordsApiService Inserts a new HeaderFooter object to the document section.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.HeaderFooterResponse*/
func (a *WordsApiService) InsertHeaderFooter(ctx context.Context, data *models.InsertHeaderFooterRequest) (models.HeaderFooterResponse, *http.Response, error) {
    var (
        successPayload models.HeaderFooterResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Inserts a new HeaderFooter object to the document section.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return InsertHeaderFooterOnlineResponse*/
func (a *WordsApiService) InsertHeaderFooterOnline(ctx context.Context, data *models.InsertHeaderFooterOnlineRequest) (models.InsertHeaderFooterOnlineResponse, *http.Response, error) {
    var (
        successPayload models.InsertHeaderFooterOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.InsertHeaderFooterOnlineResponse), response, err
}

/* WordsApiService Inserts a new list to the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.ListResponse*/
func (a *WordsApiService) InsertList(ctx context.Context, data *models.InsertListRequest) (models.ListResponse, *http.Response, error) {
    var (
        successPayload models.ListResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Inserts a new list to the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return InsertListOnlineResponse*/
func (a *WordsApiService) InsertListOnline(ctx context.Context, data *models.InsertListOnlineRequest) (models.InsertListOnlineResponse, *http.Response, error) {
    var (
        successPayload models.InsertListOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.InsertListOnlineResponse), response, err
}

/* WordsApiService Inserts a new or updates an existing paragraph tab stop in the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.TabStopsResponse*/
func (a *WordsApiService) InsertOrUpdateParagraphTabStop(ctx context.Context, data *models.InsertOrUpdateParagraphTabStopRequest) (models.TabStopsResponse, *http.Response, error) {
    var (
        successPayload models.TabStopsResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Inserts a new or updates an existing paragraph tab stop in the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return InsertOrUpdateParagraphTabStopOnlineResponse*/
func (a *WordsApiService) InsertOrUpdateParagraphTabStopOnline(ctx context.Context, data *models.InsertOrUpdateParagraphTabStopOnlineRequest) (models.InsertOrUpdateParagraphTabStopOnlineResponse, *http.Response, error) {
    var (
        successPayload models.InsertOrUpdateParagraphTabStopOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.InsertOrUpdateParagraphTabStopOnlineResponse), response, err
}

/* WordsApiService Inserts page numbers to the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.DocumentResponse*/
func (a *WordsApiService) InsertPageNumbers(ctx context.Context, data *models.InsertPageNumbersRequest) (models.DocumentResponse, *http.Response, error) {
    var (
        successPayload models.DocumentResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Inserts page numbers to the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return InsertPageNumbersOnlineResponse*/
func (a *WordsApiService) InsertPageNumbersOnline(ctx context.Context, data *models.InsertPageNumbersOnlineRequest) (models.InsertPageNumbersOnlineResponse, *http.Response, error) {
    var (
        successPayload models.InsertPageNumbersOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.InsertPageNumbersOnlineResponse), response, err
}

/* WordsApiService Inserts a new paragraph to the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.ParagraphResponse*/
func (a *WordsApiService) InsertParagraph(ctx context.Context, data *models.InsertParagraphRequest) (models.ParagraphResponse, *http.Response, error) {
    var (
        successPayload models.ParagraphResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Inserts a new paragraph to the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return InsertParagraphOnlineResponse*/
func (a *WordsApiService) InsertParagraphOnline(ctx context.Context, data *models.InsertParagraphOnlineRequest) (models.InsertParagraphOnlineResponse, *http.Response, error) {
    var (
        successPayload models.InsertParagraphOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.InsertParagraphOnlineResponse), response, err
}

/* WordsApiService Inserts a new Run object to the paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.RunResponse*/
func (a *WordsApiService) InsertRun(ctx context.Context, data *models.InsertRunRequest) (models.RunResponse, *http.Response, error) {
    var (
        successPayload models.RunResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Inserts a new Run object to the paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return InsertRunOnlineResponse*/
func (a *WordsApiService) InsertRunOnline(ctx context.Context, data *models.InsertRunOnlineRequest) (models.InsertRunOnlineResponse, *http.Response, error) {
    var (
        successPayload models.InsertRunOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.InsertRunOnlineResponse), response, err
}

/* WordsApiService Inserts a section to the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) InsertSection(ctx context.Context, data *models.InsertSectionRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Inserts a section to the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return map[string]io.Reader*/
func (a *WordsApiService) InsertSectionOnline(ctx context.Context, data *models.InsertSectionOnlineRequest) (map[string]io.Reader, *http.Response, error) {
    var (
        successPayload map[string]io.Reader
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    if err != nil {
        return successPayload, response, err
    }

    successPayload, err = models.ParseFilesCollection(response)
    return successPayload, response, err
}

/* WordsApiService Inserts a new StructuredDocumentTag (SDT) to the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.StructuredDocumentTagResponse*/
func (a *WordsApiService) InsertStructuredDocumentTag(ctx context.Context, data *models.InsertStructuredDocumentTagRequest) (models.StructuredDocumentTagResponse, *http.Response, error) {
    var (
        successPayload models.StructuredDocumentTagResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Inserts a new StructuredDocumentTag (SDT) to the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return InsertStructuredDocumentTagOnlineResponse*/
func (a *WordsApiService) InsertStructuredDocumentTagOnline(ctx context.Context, data *models.InsertStructuredDocumentTagOnlineRequest) (models.InsertStructuredDocumentTagOnlineResponse, *http.Response, error) {
    var (
        successPayload models.InsertStructuredDocumentTagOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.InsertStructuredDocumentTagOnlineResponse), response, err
}

/* WordsApiService Inserts a new style to the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.StyleResponse*/
func (a *WordsApiService) InsertStyle(ctx context.Context, data *models.InsertStyleRequest) (models.StyleResponse, *http.Response, error) {
    var (
        successPayload models.StyleResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Inserts a new style to the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return InsertStyleOnlineResponse*/
func (a *WordsApiService) InsertStyleOnline(ctx context.Context, data *models.InsertStyleOnlineRequest) (models.InsertStyleOnlineResponse, *http.Response, error) {
    var (
        successPayload models.InsertStyleOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.InsertStyleOnlineResponse), response, err
}

/* WordsApiService Inserts a new table to the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.TableResponse*/
func (a *WordsApiService) InsertTable(ctx context.Context, data *models.InsertTableRequest) (models.TableResponse, *http.Response, error) {
    var (
        successPayload models.TableResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Inserts a new cell to the table row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.TableCellResponse*/
func (a *WordsApiService) InsertTableCell(ctx context.Context, data *models.InsertTableCellRequest) (models.TableCellResponse, *http.Response, error) {
    var (
        successPayload models.TableCellResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Inserts a new cell to the table row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return InsertTableCellOnlineResponse*/
func (a *WordsApiService) InsertTableCellOnline(ctx context.Context, data *models.InsertTableCellOnlineRequest) (models.InsertTableCellOnlineResponse, *http.Response, error) {
    var (
        successPayload models.InsertTableCellOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.InsertTableCellOnlineResponse), response, err
}

/* WordsApiService Inserts a new table to the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return InsertTableOnlineResponse*/
func (a *WordsApiService) InsertTableOnline(ctx context.Context, data *models.InsertTableOnlineRequest) (models.InsertTableOnlineResponse, *http.Response, error) {
    var (
        successPayload models.InsertTableOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.InsertTableOnlineResponse), response, err
}

/* WordsApiService Inserts a new row to the table.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.TableRowResponse*/
func (a *WordsApiService) InsertTableRow(ctx context.Context, data *models.InsertTableRowRequest) (models.TableRowResponse, *http.Response, error) {
    var (
        successPayload models.TableRowResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Inserts a new row to the table.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return InsertTableRowOnlineResponse*/
func (a *WordsApiService) InsertTableRowOnline(ctx context.Context, data *models.InsertTableRowOnlineRequest) (models.InsertTableRowOnlineResponse, *http.Response, error) {
    var (
        successPayload models.InsertTableRowOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.InsertTableRowOnlineResponse), response, err
}

/* WordsApiService Insert a watermark to the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.DocumentResponse*/
func (a *WordsApiService) InsertWatermark(ctx context.Context, data *models.InsertWatermarkRequest) (models.DocumentResponse, *http.Response, error) {
    var (
        successPayload models.DocumentResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Inserts a new watermark image to the document.
 * Deprecated: This operation is deprecated and is used for backward compatibility only. Please use InsertWatermark instead.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.DocumentResponse*/
func (a *WordsApiService) InsertWatermarkImage(ctx context.Context, data *models.InsertWatermarkImageRequest) (models.DocumentResponse, *http.Response, error) {
    var (
        successPayload models.DocumentResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Inserts a new watermark image to the document.
 * Deprecated: This operation is deprecated and is used for backward compatibility only. Please use InsertWatermark instead.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return InsertWatermarkImageOnlineResponse*/
func (a *WordsApiService) InsertWatermarkImageOnline(ctx context.Context, data *models.InsertWatermarkImageOnlineRequest) (models.InsertWatermarkImageOnlineResponse, *http.Response, error) {
    var (
        successPayload models.InsertWatermarkImageOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.InsertWatermarkImageOnlineResponse), response, err
}

/* WordsApiService Insert a watermark to the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return InsertWatermarkOnlineResponse*/
func (a *WordsApiService) InsertWatermarkOnline(ctx context.Context, data *models.InsertWatermarkOnlineRequest) (models.InsertWatermarkOnlineResponse, *http.Response, error) {
    var (
        successPayload models.InsertWatermarkOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.InsertWatermarkOnlineResponse), response, err
}

/* WordsApiService Inserts a new watermark text to the document.
 * Deprecated: This operation is deprecated and is used for backward compatibility only. Please use InsertWatermark instead.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.DocumentResponse*/
func (a *WordsApiService) InsertWatermarkText(ctx context.Context, data *models.InsertWatermarkTextRequest) (models.DocumentResponse, *http.Response, error) {
    var (
        successPayload models.DocumentResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Inserts a new watermark text to the document.
 * Deprecated: This operation is deprecated and is used for backward compatibility only. Please use InsertWatermark instead.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return InsertWatermarkTextOnlineResponse*/
func (a *WordsApiService) InsertWatermarkTextOnline(ctx context.Context, data *models.InsertWatermarkTextOnlineRequest) (models.InsertWatermarkTextOnlineResponse, *http.Response, error) {
    var (
        successPayload models.InsertWatermarkTextOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.InsertWatermarkTextOnlineResponse), response, err
}

/* WordsApiService Links headers / footers of the section to the previous one.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) LinkHeaderFootersToPrevious(ctx context.Context, data *models.LinkHeaderFootersToPreviousRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Downloads a document from the Web using URL and saves it to cloud storage in the specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.SaveResponse*/
func (a *WordsApiService) LoadWebDocument(ctx context.Context, data *models.LoadWebDocumentRequest) (models.SaveResponse, *http.Response, error) {
    var (
        successPayload models.SaveResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Merge the section with the next one.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) MergeWithNext(ctx context.Context, data *models.MergeWithNextRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Merge the section with the next one.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return map[string]io.Reader*/
func (a *WordsApiService) MergeWithNextOnline(ctx context.Context, data *models.MergeWithNextOnlineRequest) (map[string]io.Reader, *http.Response, error) {
    var (
        successPayload map[string]io.Reader
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    if err != nil {
        return successPayload, response, err
    }

    successPayload, err = models.ParseFilesCollection(response)
    return successPayload, response, err
}

/* WordsApiService Move file.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) MoveFile(ctx context.Context, data *models.MoveFileRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Move folder.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) MoveFolder(ctx context.Context, data *models.MoveFolderRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Applies document content optimization options, specific to a particular versions of Microsoft Word.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) OptimizeDocument(ctx context.Context, data *models.OptimizeDocumentRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Applies document content optimization options, specific to a particular versions of Microsoft Word.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return map[string]io.Reader*/
func (a *WordsApiService) OptimizeDocumentOnline(ctx context.Context, data *models.OptimizeDocumentOnlineRequest) (map[string]io.Reader, *http.Response, error) {
    var (
        successPayload map[string]io.Reader
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    if err != nil {
        return successPayload, response, err
    }

    successPayload, err = models.ParseFilesCollection(response)
    return successPayload, response, err
}

/* WordsApiService Changes the document protection. The previous protection will be overwritten if it exist.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.ProtectionDataResponse*/
func (a *WordsApiService) ProtectDocument(ctx context.Context, data *models.ProtectDocumentRequest) (models.ProtectionDataResponse, *http.Response, error) {
    var (
        successPayload models.ProtectionDataResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Changes the document protection. The previous protection will be overwritten if it exist.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return ProtectDocumentOnlineResponse*/
func (a *WordsApiService) ProtectDocumentOnline(ctx context.Context, data *models.ProtectDocumentOnlineRequest) (models.ProtectDocumentOnlineResponse, *http.Response, error) {
    var (
        successPayload models.ProtectDocumentOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.ProtectDocumentOnlineResponse), response, err
}

/* WordsApiService Rejects all revisions in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.RevisionsModificationResponse*/
func (a *WordsApiService) RejectAllRevisions(ctx context.Context, data *models.RejectAllRevisionsRequest) (models.RevisionsModificationResponse, *http.Response, error) {
    var (
        successPayload models.RevisionsModificationResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Rejects all revisions in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return RejectAllRevisionsOnlineResponse*/
func (a *WordsApiService) RejectAllRevisionsOnline(ctx context.Context, data *models.RejectAllRevisionsOnlineRequest) (models.RejectAllRevisionsOnlineResponse, *http.Response, error) {
    var (
        successPayload models.RejectAllRevisionsOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.RejectAllRevisionsOnlineResponse), response, err
}

/* WordsApiService Removes all signatures of the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.SignatureCollectionResponse*/
func (a *WordsApiService) RemoveAllSignatures(ctx context.Context, data *models.RemoveAllSignaturesRequest) (models.SignatureCollectionResponse, *http.Response, error) {
    var (
        successPayload models.SignatureCollectionResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Removes all signatures of the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return RemoveAllSignaturesOnlineResponse*/
func (a *WordsApiService) RemoveAllSignaturesOnline(ctx context.Context, data *models.RemoveAllSignaturesOnlineRequest) (models.RemoveAllSignaturesOnlineResponse, *http.Response, error) {
    var (
        successPayload models.RemoveAllSignaturesOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.RemoveAllSignaturesOnlineResponse), response, err
}

/* WordsApiService Removes a range from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.DocumentResponse*/
func (a *WordsApiService) RemoveRange(ctx context.Context, data *models.RemoveRangeRequest) (models.DocumentResponse, *http.Response, error) {
    var (
        successPayload models.DocumentResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Removes a range from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return RemoveRangeOnlineResponse*/
func (a *WordsApiService) RemoveRangeOnline(ctx context.Context, data *models.RemoveRangeOnlineRequest) (models.RemoveRangeOnlineResponse, *http.Response, error) {
    var (
        successPayload models.RemoveRangeOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.RemoveRangeOnlineResponse), response, err
}

/* WordsApiService Renders a DrawingObject to the specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return http.Response.Body contains binary result data*/
func (a *WordsApiService) RenderDrawingObject(ctx context.Context, data *models.RenderDrawingObjectRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Renders a DrawingObject to the specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return http.Response.Body contains binary result data*/
func (a *WordsApiService) RenderDrawingObjectOnline(ctx context.Context, data *models.RenderDrawingObjectOnlineRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Renders an OfficeMath object to the specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return http.Response.Body contains binary result data*/
func (a *WordsApiService) RenderMathObject(ctx context.Context, data *models.RenderMathObjectRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Renders an OfficeMath object to the specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return http.Response.Body contains binary result data*/
func (a *WordsApiService) RenderMathObjectOnline(ctx context.Context, data *models.RenderMathObjectOnlineRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Renders a page to the specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return http.Response.Body contains binary result data*/
func (a *WordsApiService) RenderPage(ctx context.Context, data *models.RenderPageRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Renders a page to the specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return http.Response.Body contains binary result data*/
func (a *WordsApiService) RenderPageOnline(ctx context.Context, data *models.RenderPageOnlineRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Renders a paragraph to the specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return http.Response.Body contains binary result data*/
func (a *WordsApiService) RenderParagraph(ctx context.Context, data *models.RenderParagraphRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Renders a paragraph to the specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return http.Response.Body contains binary result data*/
func (a *WordsApiService) RenderParagraphOnline(ctx context.Context, data *models.RenderParagraphOnlineRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Renders a table to the specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return http.Response.Body contains binary result data*/
func (a *WordsApiService) RenderTable(ctx context.Context, data *models.RenderTableRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Renders a table to the specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return http.Response.Body contains binary result data*/
func (a *WordsApiService) RenderTableOnline(ctx context.Context, data *models.RenderTableOnlineRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Replaces text in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.ReplaceTextResponse*/
func (a *WordsApiService) ReplaceText(ctx context.Context, data *models.ReplaceTextRequest) (models.ReplaceTextResponse, *http.Response, error) {
    var (
        successPayload models.ReplaceTextResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Replaces text in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return ReplaceTextOnlineResponse*/
func (a *WordsApiService) ReplaceTextOnline(ctx context.Context, data *models.ReplaceTextOnlineRequest) (models.ReplaceTextOnlineResponse, *http.Response, error) {
    var (
        successPayload models.ReplaceTextOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.ReplaceTextOnlineResponse), response, err
}

/* WordsApiService Replaces a range with text in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.DocumentResponse*/
func (a *WordsApiService) ReplaceWithText(ctx context.Context, data *models.ReplaceWithTextRequest) (models.DocumentResponse, *http.Response, error) {
    var (
        successPayload models.DocumentResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Replaces a range with text in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return ReplaceWithTextOnlineResponse*/
func (a *WordsApiService) ReplaceWithTextOnline(ctx context.Context, data *models.ReplaceWithTextOnlineRequest) (models.ReplaceWithTextOnlineResponse, *http.Response, error) {
    var (
        successPayload models.ReplaceWithTextOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.ReplaceWithTextOnlineResponse), response, err
}

/* WordsApiService Clears the font cache.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return */
func (a *WordsApiService) ResetCache(ctx context.Context, data *models.ResetCacheRequest) (*http.Response, error) {
    var (
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return nil, err
    }


    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, err
        }

        apiError.Deserialize(jsonMap)
        return response, &apiError
    }
    return response, err
}

/* WordsApiService Converts a document in cloud storage to the specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.SaveResponse*/
func (a *WordsApiService) SaveAs(ctx context.Context, data *models.SaveAsRequest) (models.SaveResponse, *http.Response, error) {
    var (
        successPayload models.SaveResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Converts a document in cloud storage to the specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return SaveAsOnlineResponse*/
func (a *WordsApiService) SaveAsOnline(ctx context.Context, data *models.SaveAsOnlineRequest) (models.SaveAsOnlineResponse, *http.Response, error) {
    var (
        successPayload models.SaveAsOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.SaveAsOnlineResponse), response, err
}

/* WordsApiService Saves a range as a new document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.DocumentResponse*/
func (a *WordsApiService) SaveAsRange(ctx context.Context, data *models.SaveAsRangeRequest) (models.DocumentResponse, *http.Response, error) {
    var (
        successPayload models.DocumentResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Saves a range as a new document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return SaveAsRangeOnlineResponse*/
func (a *WordsApiService) SaveAsRangeOnline(ctx context.Context, data *models.SaveAsRangeOnlineRequest) (models.SaveAsRangeOnlineResponse, *http.Response, error) {
    var (
        successPayload models.SaveAsRangeOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.SaveAsRangeOnlineResponse), response, err
}

/* WordsApiService Converts a document in cloud storage to TIFF format using detailed conversion settings.
 * Deprecated: This operation will be removed in the future.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.SaveResponse*/
func (a *WordsApiService) SaveAsTiff(ctx context.Context, data *models.SaveAsTiffRequest) (models.SaveResponse, *http.Response, error) {
    var (
        successPayload models.SaveResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Converts a document in cloud storage to TIFF format using detailed conversion settings.
 * Deprecated: This operation will be removed in the future.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return SaveAsTiffOnlineResponse*/
func (a *WordsApiService) SaveAsTiffOnline(ctx context.Context, data *models.SaveAsTiffOnlineRequest) (models.SaveAsTiffOnlineResponse, *http.Response, error) {
    var (
        successPayload models.SaveAsTiffOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.SaveAsTiffOnlineResponse), response, err
}

/* WordsApiService Searches text, specified by the regular expression, in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.SearchResponse*/
func (a *WordsApiService) Search(ctx context.Context, data *models.SearchRequest) (models.SearchResponse, *http.Response, error) {
    var (
        successPayload models.SearchResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Searches text, specified by the regular expression, in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.SearchResponse*/
func (a *WordsApiService) SearchOnline(ctx context.Context, data *models.SearchOnlineRequest) (models.SearchResponse, *http.Response, error) {
    var (
        successPayload models.SearchResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Signs the document with given certificate.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.SignatureCollectionResponse*/
func (a *WordsApiService) SignDocument(ctx context.Context, data *models.SignDocumentRequest) (models.SignatureCollectionResponse, *http.Response, error) {
    var (
        successPayload models.SignatureCollectionResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Signs the document with given certificate.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return SignDocumentOnlineResponse*/
func (a *WordsApiService) SignDocumentOnline(ctx context.Context, data *models.SignDocumentOnlineRequest) (models.SignDocumentOnlineResponse, *http.Response, error) {
    var (
        successPayload models.SignDocumentOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.SignDocumentOnlineResponse), response, err
}

/* WordsApiService Splits a document into parts and saves them in the specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.SplitDocumentResponse*/
func (a *WordsApiService) SplitDocument(ctx context.Context, data *models.SplitDocumentRequest) (models.SplitDocumentResponse, *http.Response, error) {
    var (
        successPayload models.SplitDocumentResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Splits a document into parts and saves them in the specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return SplitDocumentOnlineResponse*/
func (a *WordsApiService) SplitDocumentOnline(ctx context.Context, data *models.SplitDocumentOnlineRequest) (models.SplitDocumentOnlineResponse, *http.Response, error) {
    var (
        successPayload models.SplitDocumentOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.SplitDocumentOnlineResponse), response, err
}

/* WordsApiService Translate a node id to a node path.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.TranslateNodeIdResponse*/
func (a *WordsApiService) TranslateNodeId(ctx context.Context, data *models.TranslateNodeIdRequest) (models.TranslateNodeIdResponse, *http.Response, error) {
    var (
        successPayload models.TranslateNodeIdResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Translate a node id to a node path.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.TranslateNodeIdResponse*/
func (a *WordsApiService) TranslateNodeIdOnline(ctx context.Context, data *models.TranslateNodeIdOnlineRequest) (models.TranslateNodeIdResponse, *http.Response, error) {
    var (
        successPayload models.TranslateNodeIdResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Removes protection from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.ProtectionDataResponse*/
func (a *WordsApiService) UnprotectDocument(ctx context.Context, data *models.UnprotectDocumentRequest) (models.ProtectionDataResponse, *http.Response, error) {
    var (
        successPayload models.ProtectionDataResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Removes protection from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return UnprotectDocumentOnlineResponse*/
func (a *WordsApiService) UnprotectDocumentOnline(ctx context.Context, data *models.UnprotectDocumentOnlineRequest) (models.UnprotectDocumentOnlineResponse, *http.Response, error) {
    var (
        successPayload models.UnprotectDocumentOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.UnprotectDocumentOnlineResponse), response, err
}

/* WordsApiService Updates a bookmark in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.BookmarkResponse*/
func (a *WordsApiService) UpdateBookmark(ctx context.Context, data *models.UpdateBookmarkRequest) (models.BookmarkResponse, *http.Response, error) {
    var (
        successPayload models.BookmarkResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Updates a bookmark in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return UpdateBookmarkOnlineResponse*/
func (a *WordsApiService) UpdateBookmarkOnline(ctx context.Context, data *models.UpdateBookmarkOnlineRequest) (models.UpdateBookmarkOnlineResponse, *http.Response, error) {
    var (
        successPayload models.UpdateBookmarkOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.UpdateBookmarkOnlineResponse), response, err
}

/* WordsApiService Updates a border in the document node. The 'nodePath' parameter should refer to a paragraph, a cell or a row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.BorderResponse*/
func (a *WordsApiService) UpdateBorder(ctx context.Context, data *models.UpdateBorderRequest) (models.BorderResponse, *http.Response, error) {
    var (
        successPayload models.BorderResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Updates a border in the document node. The 'nodePath' parameter should refer to a paragraph, a cell or a row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return UpdateBorderOnlineResponse*/
func (a *WordsApiService) UpdateBorderOnline(ctx context.Context, data *models.UpdateBorderOnlineRequest) (models.UpdateBorderOnlineResponse, *http.Response, error) {
    var (
        successPayload models.UpdateBorderOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.UpdateBorderOnlineResponse), response, err
}

/* WordsApiService Updates a comment in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.CommentResponse*/
func (a *WordsApiService) UpdateComment(ctx context.Context, data *models.UpdateCommentRequest) (models.CommentResponse, *http.Response, error) {
    var (
        successPayload models.CommentResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Updates a comment in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return UpdateCommentOnlineResponse*/
func (a *WordsApiService) UpdateCommentOnline(ctx context.Context, data *models.UpdateCommentOnlineRequest) (models.UpdateCommentOnlineResponse, *http.Response, error) {
    var (
        successPayload models.UpdateCommentOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.UpdateCommentOnlineResponse), response, err
}

/* WordsApiService Updates the custom xml part in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.CustomXmlPartResponse*/
func (a *WordsApiService) UpdateCustomXmlPart(ctx context.Context, data *models.UpdateCustomXmlPartRequest) (models.CustomXmlPartResponse, *http.Response, error) {
    var (
        successPayload models.CustomXmlPartResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Updates the custom xml part in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return UpdateCustomXmlPartOnlineResponse*/
func (a *WordsApiService) UpdateCustomXmlPartOnline(ctx context.Context, data *models.UpdateCustomXmlPartOnlineRequest) (models.UpdateCustomXmlPartOnlineResponse, *http.Response, error) {
    var (
        successPayload models.UpdateCustomXmlPartOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.UpdateCustomXmlPartOnlineResponse), response, err
}

/* WordsApiService Updates a DrawingObject in the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.DrawingObjectResponse*/
func (a *WordsApiService) UpdateDrawingObject(ctx context.Context, data *models.UpdateDrawingObjectRequest) (models.DrawingObjectResponse, *http.Response, error) {
    var (
        successPayload models.DrawingObjectResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Updates a DrawingObject in the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return UpdateDrawingObjectOnlineResponse*/
func (a *WordsApiService) UpdateDrawingObjectOnline(ctx context.Context, data *models.UpdateDrawingObjectOnlineRequest) (models.UpdateDrawingObjectOnlineResponse, *http.Response, error) {
    var (
        successPayload models.UpdateDrawingObjectOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.UpdateDrawingObjectOnlineResponse), response, err
}

/* WordsApiService Updates a field in the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.FieldResponse*/
func (a *WordsApiService) UpdateField(ctx context.Context, data *models.UpdateFieldRequest) (models.FieldResponse, *http.Response, error) {
    var (
        successPayload models.FieldResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Updates a field in the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return UpdateFieldOnlineResponse*/
func (a *WordsApiService) UpdateFieldOnline(ctx context.Context, data *models.UpdateFieldOnlineRequest) (models.UpdateFieldOnlineResponse, *http.Response, error) {
    var (
        successPayload models.UpdateFieldOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.UpdateFieldOnlineResponse), response, err
}

/* WordsApiService Reevaluates field values in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.DocumentResponse*/
func (a *WordsApiService) UpdateFields(ctx context.Context, data *models.UpdateFieldsRequest) (models.DocumentResponse, *http.Response, error) {
    var (
        successPayload models.DocumentResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Reevaluates field values in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return UpdateFieldsOnlineResponse*/
func (a *WordsApiService) UpdateFieldsOnline(ctx context.Context, data *models.UpdateFieldsOnlineRequest) (models.UpdateFieldsOnlineResponse, *http.Response, error) {
    var (
        successPayload models.UpdateFieldsOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.UpdateFieldsOnlineResponse), response, err
}

/* WordsApiService Updates a footnote in the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.FootnoteResponse*/
func (a *WordsApiService) UpdateFootnote(ctx context.Context, data *models.UpdateFootnoteRequest) (models.FootnoteResponse, *http.Response, error) {
    var (
        successPayload models.FootnoteResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Updates a footnote in the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return UpdateFootnoteOnlineResponse*/
func (a *WordsApiService) UpdateFootnoteOnline(ctx context.Context, data *models.UpdateFootnoteOnlineRequest) (models.UpdateFootnoteOnlineResponse, *http.Response, error) {
    var (
        successPayload models.UpdateFootnoteOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.UpdateFootnoteOnlineResponse), response, err
}

/* WordsApiService Updates a form field in the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.FormFieldResponse*/
func (a *WordsApiService) UpdateFormField(ctx context.Context, data *models.UpdateFormFieldRequest) (models.FormFieldResponse, *http.Response, error) {
    var (
        successPayload models.FormFieldResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Updates a form field in the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return UpdateFormFieldOnlineResponse*/
func (a *WordsApiService) UpdateFormFieldOnline(ctx context.Context, data *models.UpdateFormFieldOnlineRequest) (models.UpdateFormFieldOnlineResponse, *http.Response, error) {
    var (
        successPayload models.UpdateFormFieldOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.UpdateFormFieldOnlineResponse), response, err
}

/* WordsApiService Updates a list in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.ListResponse*/
func (a *WordsApiService) UpdateList(ctx context.Context, data *models.UpdateListRequest) (models.ListResponse, *http.Response, error) {
    var (
        successPayload models.ListResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Updates the level of a List element in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.ListResponse*/
func (a *WordsApiService) UpdateListLevel(ctx context.Context, data *models.UpdateListLevelRequest) (models.ListResponse, *http.Response, error) {
    var (
        successPayload models.ListResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Updates the level of a List element in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return UpdateListLevelOnlineResponse*/
func (a *WordsApiService) UpdateListLevelOnline(ctx context.Context, data *models.UpdateListLevelOnlineRequest) (models.UpdateListLevelOnlineResponse, *http.Response, error) {
    var (
        successPayload models.UpdateListLevelOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.UpdateListLevelOnlineResponse), response, err
}

/* WordsApiService Updates a list in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return UpdateListOnlineResponse*/
func (a *WordsApiService) UpdateListOnline(ctx context.Context, data *models.UpdateListOnlineRequest) (models.UpdateListOnlineResponse, *http.Response, error) {
    var (
        successPayload models.UpdateListOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.UpdateListOnlineResponse), response, err
}

/* WordsApiService Updates the formatting properties of a paragraph in the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.ParagraphFormatResponse*/
func (a *WordsApiService) UpdateParagraphFormat(ctx context.Context, data *models.UpdateParagraphFormatRequest) (models.ParagraphFormatResponse, *http.Response, error) {
    var (
        successPayload models.ParagraphFormatResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Updates the formatting properties of a paragraph in the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return UpdateParagraphFormatOnlineResponse*/
func (a *WordsApiService) UpdateParagraphFormatOnline(ctx context.Context, data *models.UpdateParagraphFormatOnlineRequest) (models.UpdateParagraphFormatOnlineResponse, *http.Response, error) {
    var (
        successPayload models.UpdateParagraphFormatOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.UpdateParagraphFormatOnlineResponse), response, err
}

/* WordsApiService Updates the formatting properties of a paragraph list in the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.ParagraphListFormatResponse*/
func (a *WordsApiService) UpdateParagraphListFormat(ctx context.Context, data *models.UpdateParagraphListFormatRequest) (models.ParagraphListFormatResponse, *http.Response, error) {
    var (
        successPayload models.ParagraphListFormatResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Updates the formatting properties of a paragraph list in the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return UpdateParagraphListFormatOnlineResponse*/
func (a *WordsApiService) UpdateParagraphListFormatOnline(ctx context.Context, data *models.UpdateParagraphListFormatOnlineRequest) (models.UpdateParagraphListFormatOnlineResponse, *http.Response, error) {
    var (
        successPayload models.UpdateParagraphListFormatOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.UpdateParagraphListFormatOnlineResponse), response, err
}

/* WordsApiService Updates a Run object in the paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.RunResponse*/
func (a *WordsApiService) UpdateRun(ctx context.Context, data *models.UpdateRunRequest) (models.RunResponse, *http.Response, error) {
    var (
        successPayload models.RunResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Updates the font properties of a Run object in the paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.FontResponse*/
func (a *WordsApiService) UpdateRunFont(ctx context.Context, data *models.UpdateRunFontRequest) (models.FontResponse, *http.Response, error) {
    var (
        successPayload models.FontResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Updates the font properties of a Run object in the paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return UpdateRunFontOnlineResponse*/
func (a *WordsApiService) UpdateRunFontOnline(ctx context.Context, data *models.UpdateRunFontOnlineRequest) (models.UpdateRunFontOnlineResponse, *http.Response, error) {
    var (
        successPayload models.UpdateRunFontOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.UpdateRunFontOnlineResponse), response, err
}

/* WordsApiService Updates a Run object in the paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return UpdateRunOnlineResponse*/
func (a *WordsApiService) UpdateRunOnline(ctx context.Context, data *models.UpdateRunOnlineRequest) (models.UpdateRunOnlineResponse, *http.Response, error) {
    var (
        successPayload models.UpdateRunOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.UpdateRunOnlineResponse), response, err
}

/* WordsApiService Updates the page setup of a section in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.SectionPageSetupResponse*/
func (a *WordsApiService) UpdateSectionPageSetup(ctx context.Context, data *models.UpdateSectionPageSetupRequest) (models.SectionPageSetupResponse, *http.Response, error) {
    var (
        successPayload models.SectionPageSetupResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Updates the page setup of a section in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return UpdateSectionPageSetupOnlineResponse*/
func (a *WordsApiService) UpdateSectionPageSetupOnline(ctx context.Context, data *models.UpdateSectionPageSetupOnlineRequest) (models.UpdateSectionPageSetupOnlineResponse, *http.Response, error) {
    var (
        successPayload models.UpdateSectionPageSetupOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.UpdateSectionPageSetupOnlineResponse), response, err
}

/* WordsApiService Updates a StructuredDocumentTag (SDT) in the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.StructuredDocumentTagResponse*/
func (a *WordsApiService) UpdateStructuredDocumentTag(ctx context.Context, data *models.UpdateStructuredDocumentTagRequest) (models.StructuredDocumentTagResponse, *http.Response, error) {
    var (
        successPayload models.StructuredDocumentTagResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Updates a StructuredDocumentTag (SDT) in the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return UpdateStructuredDocumentTagOnlineResponse*/
func (a *WordsApiService) UpdateStructuredDocumentTagOnline(ctx context.Context, data *models.UpdateStructuredDocumentTagOnlineRequest) (models.UpdateStructuredDocumentTagOnlineResponse, *http.Response, error) {
    var (
        successPayload models.UpdateStructuredDocumentTagOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.UpdateStructuredDocumentTagOnlineResponse), response, err
}

/* WordsApiService Updates a style in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.StyleResponse*/
func (a *WordsApiService) UpdateStyle(ctx context.Context, data *models.UpdateStyleRequest) (models.StyleResponse, *http.Response, error) {
    var (
        successPayload models.StyleResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Updates a style in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return UpdateStyleOnlineResponse*/
func (a *WordsApiService) UpdateStyleOnline(ctx context.Context, data *models.UpdateStyleOnlineRequest) (models.UpdateStyleOnlineResponse, *http.Response, error) {
    var (
        successPayload models.UpdateStyleOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.UpdateStyleOnlineResponse), response, err
}

/* WordsApiService Updates the formatting properties of a cell in the table row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.TableCellFormatResponse*/
func (a *WordsApiService) UpdateTableCellFormat(ctx context.Context, data *models.UpdateTableCellFormatRequest) (models.TableCellFormatResponse, *http.Response, error) {
    var (
        successPayload models.TableCellFormatResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Updates the formatting properties of a cell in the table row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return UpdateTableCellFormatOnlineResponse*/
func (a *WordsApiService) UpdateTableCellFormatOnline(ctx context.Context, data *models.UpdateTableCellFormatOnlineRequest) (models.UpdateTableCellFormatOnlineResponse, *http.Response, error) {
    var (
        successPayload models.UpdateTableCellFormatOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.UpdateTableCellFormatOnlineResponse), response, err
}

/* WordsApiService Updates properties of a table in the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.TablePropertiesResponse*/
func (a *WordsApiService) UpdateTableProperties(ctx context.Context, data *models.UpdateTablePropertiesRequest) (models.TablePropertiesResponse, *http.Response, error) {
    var (
        successPayload models.TablePropertiesResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Updates properties of a table in the document node.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return UpdateTablePropertiesOnlineResponse*/
func (a *WordsApiService) UpdateTablePropertiesOnline(ctx context.Context, data *models.UpdateTablePropertiesOnlineRequest) (models.UpdateTablePropertiesOnlineResponse, *http.Response, error) {
    var (
        successPayload models.UpdateTablePropertiesOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.UpdateTablePropertiesOnlineResponse), response, err
}

/* WordsApiService Updates the formatting properties of a table row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.TableRowFormatResponse*/
func (a *WordsApiService) UpdateTableRowFormat(ctx context.Context, data *models.UpdateTableRowFormatRequest) (models.TableRowFormatResponse, *http.Response, error) {
    var (
        successPayload models.TableRowFormatResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}

/* WordsApiService Updates the formatting properties of a table row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return UpdateTableRowFormatOnlineResponse*/
func (a *WordsApiService) UpdateTableRowFormatOnline(ctx context.Context, data *models.UpdateTableRowFormatOnlineRequest) (models.UpdateTableRowFormatOnlineResponse, *http.Response, error) {
    var (
        successPayload models.UpdateTableRowFormatOnlineResponse
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }


    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    boundary := GetBoundary(response)
    result, err := data.CreateResponse(response.Body, boundary)

    if err != nil {
        return successPayload, response, err
    }

    return result.(models.UpdateTableRowFormatOnlineResponse), response, err
}

/* WordsApiService Upload file.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data operation request data.
@return models.FilesUploadResult*/
func (a *WordsApiService) UploadFile(ctx context.Context, data *models.UploadFileRequest) (models.FilesUploadResult, *http.Response, error) {
    var (
        successPayload models.FilesUploadResult
    )

    requestData, err := data.CreateRequestData();
    if err != nil {
        return successPayload, nil, err
    }

    requestData.Path = a.client.cfg.BaseUrl + requestData.Path;

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return successPayload, nil, err
    }

    response, err := a.client.callAPI(r)

    if err != nil || response == nil {
        return successPayload, nil, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, nil, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }
    var jsonMap map[string]interface{}
    if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
        return successPayload, response, err
    }

    successPayload.Deserialize(jsonMap)
    return successPayload, response, err
}


/* WordsApiService Batch request.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @requests to be called as one call.
@return array of results */
/* WordsApiService Batch request.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @requests to be called as one call.
@return array of results */
func (a *WordsApiService) Batch(ctx context.Context, requests ...models.BatchPartRequest) ([]interface{}, *http.Response, error) {
    return a.batch(ctx, false, requests)
}

/* WordsApiService Batch request without intermediate results.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @requests to be called as one call.
@return array with one element (a result of last executed request). */
func (a *WordsApiService) BatchWithoutIntermidiateResults(ctx context.Context, requests ...models.BatchPartRequest) ([]interface{}, *http.Response, error) {
    return a.batch(ctx, true, requests)
}

func (a *WordsApiService) batch(ctx context.Context, withoutIntermediateResults bool, requests []models.BatchPartRequest) ([]interface{}, *http.Response, error) {

    if requests == nil || len(requests) == 0 {
        return nil, nil, errors.New("The Batch method requires one or more requests.")
    }

    // create map of requests
    rs := make(map[string]models.BatchPartRequest)

    // generate HTTP requests
    httpRequests := []http.Request{}

    for _, r := range requests {
        rs[*r.RequestId] = r

        data, err := r.CreateRequestData()
        if err != nil {
            return nil, nil, err
        }

        // fix path
        data.Path = data.Path[7:]

        req, err := a.client.prepareRequest(ctx, data)
        if err != nil {
            return nil, nil, err
        }

        // add the request and its parent IDs
        req.Header.Add("RequestId", *r.RequestId)
        if r.ParentRequestId != nil {
            req.Header.Add("DependsOn", *r.ParentRequestId)
        }

        httpRequests = append(httpRequests, *req)
    }

    // create body
    reader, writer := io.Pipe()
    defer reader.Close()

    bodyWriter := multipart.NewWriter(writer)
    boundary := bodyWriter.Boundary()

    go func() {

        for _, req := range httpRequests {
            headers := make(textproto.MIMEHeader)
            headers.Set("Content-Type", "application/http; msgtype=request")
            partWriter, _ := bodyWriter.CreatePart(headers)

            io.WriteString(partWriter, fmt.Sprintf("%s %s \r\n", req.Method, req.URL))

            if (req.Body != nil) && (req.Header.Get("Content-Type") != "") {
                io.WriteString(partWriter, "Content-Type: "+req.Header.Get("Content-Type"))
                io.WriteString(partWriter, "\r\n")
            }

            io.WriteString(partWriter, "RequestId: "+req.Header.Get("RequestId")+"\r\n")
            if req.Header.Get("DependsOn") != "" {
                io.WriteString(partWriter, "DependsOn: "+req.Header.Get("DependsOn")+"\r\n")
            }

            io.WriteString(partWriter, "\r\n")

            if req.Body != nil {
                io.Copy(partWriter, req.Body)
            }
        }

        bodyWriter.Close()
        writer.Close()
    }()

    url := a.client.cfg.BaseUrl+"/words/batch"
    if withoutIntermediateResults {
        url = url + "?displayIntermediateResults=false"
    }

    request, err := http.NewRequest("PUT", url, reader)
    if err != nil {
        return nil, nil, err
    }

    request.Header.Add("Content-Type", "multipart/batching;boundary="+boundary)

    if ctx != nil {
        // add context to the request
        request = request.WithContext(ctx)

        // AccessToken Authentication
        if auth, ok := ctx.Value(models.ContextAccessToken).(string); ok {
            request.Header.Add("Authorization", "Bearer "+auth)
        }
    }

    for header, value := range a.client.cfg.DefaultHeader {
        request.Header.Add(header, value)
    }

    var successPayload []interface{}
    response, err := a.client.callAPI(request)

    if err != nil || response == nil {
        return successPayload, response, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return successPayload, response, err
        }

        apiError.Deserialize(jsonMap)
        return successPayload, response, &apiError
    }

    _, params, err := mime.ParseMediaType(response.Header.Get("Content-Type"))
    if err != nil {
        return successPayload, response, err
    }

    mr := multipart.NewReader(response.Body, params["boundary"])
    for part, err := mr.NextPart(); err == nil; part, err = mr.NextPart() {

        r := bufio.NewReader(part)
        status, _ := r.ReadString('\n')
        statusParts := strings.Split(status, " ")
        statusCode, _ := atoi(statusParts[0])

        // parse headers
        hs := make(map[string]string)
        line := status
        for line != "" {
            line, _ = r.ReadString('\n')
            line = strings.ReplaceAll(line, "\r\n", "")
            header_parts  := strings.Split(line, ":")
            if len(header_parts) == 2 {
                hs[header_parts[0]] = strings.Trim(header_parts[1], " ")
            }
        }

        id := hs["RequestId"]

        if statusCode >= 300 {
            var apiError models.WordsApiErrorResponse
            var jsonMap map[string]interface{}
            if err = json.NewDecoder(r).Decode(&jsonMap); err != nil {
                return successPayload, response, err
            }

            apiError.Deserialize(jsonMap)
            return successPayload, response, &apiError
        }

        boundary := GetPartBoundary(part)

        body, _ := ioutil.ReadAll(r)

        rq := rs[id]
        currentResult, err := (&rq).CreateResponse(bytes.NewReader(body), boundary)
        if err != nil {
            return successPayload, response, err
        }

        successPayload = append(successPayload, currentResult)
    }

    if err != io.EOF {
        return successPayload, response, err
    }

    return successPayload, response, err
}


/* Encrypts string on APi key.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @data data to encrypt.
@return encrypted data. */
func (a *WordsApiService) Encrypt(ctx context.Context, data string) (string, error) {
    return a.client.encrypt(ctx, data)
}
