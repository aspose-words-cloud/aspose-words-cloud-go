/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="batch_part_request.go">
 *   Copyright (c) 2024 Aspose.Words for Cloud
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
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

// The REST response with data on system, additional and custom fonts, available for document processing.
type BatchPartRequest struct {
	// Request Id.
	RequestId *string

	// Parent Request Id.
	ParentRequestId *string

	// inner request
	Request RequestInterface
}

// Create new instance of struct.
func NewBatchPartRequest(request RequestInterface) *BatchPartRequest {
	return &BatchPartRequest{RequestId: createRandomGuid(), ParentRequestId: nil, Request: request}
}

// set parent request.
// @param parentRequest parent request.
func (c *BatchPartRequest) DependsOn(parentRequest BatchPartRequest) {
	c.ParentRequestId = parentRequest.RequestId
}

// use request's response as a source
func (c *BatchPartRequest) AsSource() (io.ReadCloser) {
	return ioutil.NopCloser(strings.NewReader("resultOf(" + *(c.RequestId) + ")"))
}

// create request options
func (c *BatchPartRequest) CreateRequestData() (RequestData, error) {
	return c.Request.CreateRequestData()
}

// parse response
func (c *BatchPartRequest) CreateResponse(reader io.Reader, boundary string) (response interface{}, err error) {
	return c.Request.CreateResponse(reader, boundary)
}

func createRandomGuid() *string {
	b := make([]byte, 16)
	rand.Read(b)
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return &uuid
}