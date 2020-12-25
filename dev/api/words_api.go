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
	"bufio"
	"bytes"
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

	"github.com/aspose-words-cloud/aspose-words-cloud-go/dev/api/models"
	"golang.org/x/net/context"
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
}

/* WordsApiService Supported extensions: ".doc", ".docx", ".docm", ".dot", ".dotm", ".dotx", ".flatopc", ".fopc", ".flatopc_macro", ".fopc_macro", ".flatopc_template", ".fopc_template", ".flatopc_template_macro", ".fopc_template_macro", ".wordml", ".wml", ".rtf".
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
}

/* WordsApiService The 'nodePath' parameter should refer to a paragraph, a cell or a row.
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
}

/* WordsApiService The 'nodePath' parameter should refer to a paragraph, a cell or a row.
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

        return response, &apiError
    }
    return response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

        return response, &apiError
    }
    return response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

        return response, &apiError
    }
    return response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

        return response, &apiError
    }
    return response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

        return response, &apiError
    }
    return response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

        return response, &apiError
    }
    return response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

        return response, &apiError
    }
    return response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

        return response, &apiError
    }
    return response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

        return response, &apiError
    }
    return response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

        return response, &apiError
    }
    return response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

        return response, &apiError
    }
    return response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

        return response, &apiError
    }
    return response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

        return response, &apiError
    }
    return response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

        return response, &apiError
    }
    return response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

        return response, &apiError
    }
    return response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

        return response, &apiError
    }
    return response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
}

/* WordsApiService The 'nodePath' parameter should refer to a paragraph, a cell or a row.
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
}

/* WordsApiService Inserts a new watermark image to the document.
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
}

/* WordsApiService Inserts a new watermark text to the document.
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

        return response, &apiError
    }
    return response, err
}

/* WordsApiService Adds protection to the document.
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return response, err
    }
    if response.StatusCode == 401 {
        return nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return response, err
        }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
}

/* WordsApiService Converts a document in cloud storage to TIFF format using detailed conversion settings.
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
}

/* WordsApiService The 'nodePath' parameter should refer to a paragraph, a cell or a row.
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
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
    defer response.Body.Close()

    if err != nil || response == nil {
        return successPayload, response, err
    }
    if response.StatusCode == 401 {
        return successPayload, nil, errors.New("Access is denied")
    }
    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse;

        if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
            return successPayload, response, err
        }

        return successPayload, response, &apiError
    }
    if err = json.NewDecoder(response.Body).Decode(&successPayload); err != nil {
        return successPayload, response, err
    }

    return successPayload, response, err
}


/* WordsApiService Batch request.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 * @requests to be called as one call.
@return array of results */
func (a *WordsApiService) Batch(ctx context.Context, requests ...models.RequestInterface) ([]interface{}, *http.Response, error) {

	if requests == nil || len(requests) == 0 {
		return nil, nil, errors.New("The Batch method requires one or more requests.")
	}

	// generate HTTP requests
	httpRequests := []http.Request{}

	for _, r := range requests {
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

			io.WriteString(partWriter, "\r\n")

			if req.Body != nil {
				io.Copy(partWriter, req.Body)
			}
		}

		bodyWriter.Close()
		writer.Close()
	}()

	request, err := http.NewRequest("PUT", a.client.cfg.BaseUrl+"/words/batch", reader)
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
	defer response.Body.Close()

	if err != nil || response == nil {
		return successPayload, response, err
	}
	if response.StatusCode == 401 {
		return successPayload, nil, errors.New("Access is denied")
	}
	if response.StatusCode >= 300 {
		var apiError models.WordsApiErrorResponse

		if err = json.NewDecoder(response.Body).Decode(&apiError); err != nil {
			return successPayload, response, err
		}

		return successPayload, response, &apiError
	}

	index := 0
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

		// skip headers
		line := status
		for line != "" {
			line, _ = r.ReadString('\n')
			line = strings.ReplaceAll(line, "\r\n", "")
		}

		if statusCode >= 300 {
			var apiError models.WordsApiErrorResponse

			if err = json.NewDecoder(r).Decode(&apiError); err != nil {
				return successPayload, response, err
			}

			return successPayload, response, &apiError
		}

		body, _ := ioutil.ReadAll(r)

		currentResult, err := requests[index].CreateResponse(bytes.NewReader(body))
		if err != nil {
			return successPayload, response, err
		}

		successPayload = append(successPayload, currentResult)
		index++
	}

	if err != io.EOF {
		return successPayload, response, err
	}

	return successPayload, response, err
}