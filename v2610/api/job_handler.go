/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="job_handler.go">
 *   Copyright (c) 2026 Aspose.Words for Cloud
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
    "net/url"
    "reflect"
    "strings"
    "time"

    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2610/api/models"
)

type JobHandler struct {
    service *WordsApiService
    ctx context.Context
    request models.RequestInterface
    info *models.JobInfo
    result interface{}
}

func NewJobHandler(service *WordsApiService, ctx context.Context, request models.RequestInterface, info *models.JobInfo) *JobHandler {
    return &JobHandler{
        service: service,
        ctx: ctx,
        request: request,
        info: info,
    }
}

func (c *JobHandler) GetStatus() *string {
    if c.info == nil {
        return nil
    }

    return c.info.Status
}

func (c *JobHandler) GetMessage() *string {
    if c.info == nil {
        return nil
    }

    return c.info.Message
}

func (c *JobHandler) GetResult() interface{} {
    return c.result
}

func (c *JobHandler) Update() (interface{}, error) {
    if c.info == nil || c.info.JobId == nil || *c.info.JobId == "" {
        return nil, errors.New("Invalid job id.")
    }

    info, result, _, err := c.service.callJobResult(c.ctx, *c.info.JobId, c.request)
    if err != nil {
        return nil, err
    }

    c.info = info
    if result != nil {
        c.result = result
    }

    return c.result, nil
}

func (c *JobHandler) WaitResult() (interface{}, error) {
    return c.WaitResultWithInterval(time.Second)
}

func (c *JobHandler) WaitResultWithInterval(updateInterval time.Duration) (interface{}, error) {
    for isQueuedStatus(c.GetStatus()) || isProcessingStatus(c.GetStatus()) {
        time.Sleep(updateInterval)
        if _, err := c.Update(); err != nil {
            return nil, err
        }
    }

    if isSucceededStatus(c.GetStatus()) && c.result == nil {
        if _, err := c.Update(); err != nil {
            return nil, err
        }
    }

    if !isSucceededStatus(c.GetStatus()) {
        return nil, fmt.Errorf("Job failed with status %q - %q", dereferenceString(c.GetStatus()), dereferenceString(c.GetMessage()))
    }

    return c.result, nil
}

func (a *WordsApiService) callJobResult(ctx context.Context, jobID string, request models.RequestInterface) (*models.JobInfo, interface{}, *http.Response, error) {
    requestData := models.RequestData{
        Path: a.client.cfg.BaseUrl + "/words/job",
        Method: strings.ToUpper("get"),
        HeaderParams: make(map[string]string),
        QueryParams: url.Values{},
        FormParams: make([]models.FormParamContainer, 0),
        FileReferences: make([]models.FileReference, 0),
    }
    requestData.QueryParams.Add("id", jobID)

    r, err := a.client.prepareRequest(ctx, requestData)
    if err != nil {
        return nil, nil, nil, err
    }

    response, err := a.client.callAPI(r)
    if err != nil || response == nil {
        return nil, nil, response, err
    }

    defer response.Body.Close()

    if response.StatusCode == 401 {
        return nil, nil, nil, errors.New("Access is denied")
    }

    if response.StatusCode >= 300 {
        var apiError models.WordsApiErrorResponse
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(response.Body).Decode(&jsonMap); err != nil {
            return nil, nil, response, err
        }

        apiError.Deserialize(jsonMap)
        return nil, nil, response, &apiError
    }

    info, result, err := deserializeJobMultipartResponse(response, request)
    return info, result, response, err
}

func deserializeJobMultipartResponse(response *http.Response, request models.RequestInterface) (*models.JobInfo, interface{}, error) {
    _, params, err := mime.ParseMediaType(response.Header.Get("Content-Type"))
    if err != nil {
        return nil, nil, err
    }

    mr := multipart.NewReader(response.Body, params["boundary"])

    infoPart, err := mr.NextPart()
    if err != nil {
        return nil, nil, err
    }

    var jsonMap map[string]interface{}
    if err = json.NewDecoder(infoPart).Decode(&jsonMap); err != nil {
        return nil, nil, err
    }

    info := new(models.JobInfo)
    info.Deserialize(jsonMap)

    if !isSucceededStatus(info.Status) {
        return info, nil, nil
    }

    responsePart, err := mr.NextPart()
    if err == io.EOF {
        return info, nil, nil
    }
    if err != nil {
        return nil, nil, err
    }

    result, err := deserializeJobResponsePart(request, responsePart)
    return info, result, err
}

func deserializeJobResponsePart(request models.RequestInterface, part *multipart.Part) (interface{}, error) {
    reader := bufio.NewReader(part)
    statusLine, err := reader.ReadString('\n')
    if err != nil {
        return nil, err
    }

    statusLine = strings.TrimSpace(statusLine)
    statusParts := strings.Split(statusLine, " ")
    if len(statusParts) < 3 || !strings.HasPrefix(statusParts[0], "HTTP/") {
        return nil, errors.New("Failed to parse HTTP response part.")
    }

    statusCode, err := atoi(statusParts[1])
    if err != nil {
        return nil, err
    }

    headers := make(map[string]string)
    for {
        line, err := reader.ReadString('\n')
        if err != nil && err != io.EOF {
            return nil, err
        }

        line = strings.TrimRight(line, "\r\n")
        if line == "" {
            break
        }

        headerParts := strings.SplitN(line, ":", 2)
        if len(headerParts) == 2 {
            headers[strings.TrimSpace(headerParts[0])] = strings.TrimSpace(headerParts[1])
        }

        if err == io.EOF {
            break
        }
    }

    body, err := ioutil.ReadAll(reader)
    if err != nil {
        return nil, err
    }

    if statusCode >= 300 {
        var jsonMap map[string]interface{}
        if err = json.NewDecoder(bytes.NewReader(body)).Decode(&jsonMap); err == nil {
            var apiError models.WordsApiErrorResponse
            apiError.Deserialize(jsonMap)
            return nil, &apiError
        }

        return nil, fmt.Errorf("Job result request failed with status %d", statusCode)
    }

    boundary := getBoundary(headers["Content-Type"])
    result, err := request.CreateResponse(bytes.NewReader(body), boundary)
    if err != nil {
        return nil, err
    }

    return normalizeJobResult(result), nil
}

func normalizeJobResult(result interface{}) interface{} {
    if result == nil {
        return nil
    }

    value := reflect.ValueOf(result)
    if value.Kind() != reflect.Ptr || value.IsNil() {
        return result
    }

    element := value.Elem()
    if !element.IsValid() || element.Kind() != reflect.Struct {
        return result
    }

    return element.Interface()
}

func dereferenceString(value *string) string {
    if value == nil {
        return ""
    }

    return *value
}

func isQueuedStatus(value *string) bool {
    return strings.EqualFold(dereferenceString(value), "queued")
}

func isProcessingStatus(value *string) bool {
    return strings.EqualFold(dereferenceString(value), "processing")
}

func isSucceededStatus(value *string) bool {
    status := dereferenceString(value)
    return strings.EqualFold(status, "succeded") || strings.EqualFold(status, "succeeded")
}
