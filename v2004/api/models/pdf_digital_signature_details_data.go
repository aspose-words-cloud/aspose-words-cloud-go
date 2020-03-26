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

import (
	"time"
)


// container class for details of digital signature.
type PdfDigitalSignatureDetailsData struct {

	// Gets or sets certificate's filename using for signing.
	CertificateFilename string `json:"CertificateFilename,omitempty"`

	// Gets or sets hash algorithm.
	HashAlgorithm string `json:"HashAlgorithm,omitempty"`

	// Gets or sets location of the signing.
	Location string `json:"Location,omitempty"`

	// Gets or sets reason for the signing.
	Reason string `json:"Reason,omitempty"`

	// Gets or sets date of the signing.
	SignatureDate time.Time `json:"SignatureDate,omitempty"`
}

type IPdfDigitalSignatureDetailsData interface {
	IsPdfDigitalSignatureDetailsData() bool
}
func (PdfDigitalSignatureDetailsData) IsPdfDigitalSignatureDetailsData() bool {
	return true;
}
