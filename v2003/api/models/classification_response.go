/*
* MIT License

* Copyright (c) 2019 Aspose Pty Ltd

* Permission is hereby granted, free of charge, to any person obtaining a copy
* of this software and associated documentation files (the "Software"), to deal
* in the Software without restriction, including without limitation the rights
* to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
* copies of the Software, and to permit persons to whom the Software is
* furnished to do so, subject to the following conditions:

* The above copyright notice and this permission notice shall be included in all
* copies or substantial portions of the Software.

* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
* IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
* FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
* AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
* LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
* OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
* SOFTWARE.
*/


package models



// This response should be returned by the service when handling: PUT http://api.aspose.com/v4.0/words/classify.
type ClassificationResponse struct {

	// Gets or sets request Id.
	RequestId string `json:"RequestId,omitempty"`

	// Gets or sets best class name.
	BestClassName string `json:"BestClassName,omitempty"`

	// Gets or sets best class probability.
	BestClassProbability float64 `json:"BestClassProbability,omitempty"`

	// Gets or sets array of best classes results.
	BestResults []ClassificationResult `json:"BestResults,omitempty"`
}

type IClassificationResponse interface {
	IsClassificationResponse() bool
}
func (ClassificationResponse) IsClassificationResponse() bool {
	return true;
}
func (ClassificationResponse) IsWordsResponse() bool {
	return true;
}
