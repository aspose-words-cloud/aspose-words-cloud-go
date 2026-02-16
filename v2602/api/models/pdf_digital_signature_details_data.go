/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="pdf_digital_signature_details_data.go">
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

package models

import (
    "errors"
)

// Container class for details of digital signature.

type IPdfDigitalSignatureDetailsData interface {
    IsPdfDigitalSignatureDetailsData() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetCertificateFilename() *string
    SetCertificateFilename(value *string)
    GetHashAlgorithm() *string
    SetHashAlgorithm(value *string)
    GetLocation() *string
    SetLocation(value *string)
    GetReason() *string
    SetReason(value *string)
    GetSignatureDate() *Time
    SetSignatureDate(value *Time)
}

type PdfDigitalSignatureDetailsData struct {
    // Container class for details of digital signature.
    CertificateFilename *string `json:"CertificateFilename,omitempty"`

    // Container class for details of digital signature.
    HashAlgorithm *string `json:"HashAlgorithm,omitempty"`

    // Container class for details of digital signature.
    Location *string `json:"Location,omitempty"`

    // Container class for details of digital signature.
    Reason *string `json:"Reason,omitempty"`

    // Container class for details of digital signature.
    SignatureDate *Time `json:"SignatureDate,omitempty"`
}

func (PdfDigitalSignatureDetailsData) IsPdfDigitalSignatureDetailsData() bool {
    return true
}


func (obj *PdfDigitalSignatureDetailsData) Initialize() {
}

func (obj *PdfDigitalSignatureDetailsData) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["CertificateFilename"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.CertificateFilename = &parsedValue
        }

    } else if jsonValue, exists := json["certificateFilename"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.CertificateFilename = &parsedValue
        }

    }

    if jsonValue, exists := json["HashAlgorithm"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.HashAlgorithm = &parsedValue
        }

    } else if jsonValue, exists := json["hashAlgorithm"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.HashAlgorithm = &parsedValue
        }

    }

    if jsonValue, exists := json["Location"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Location = &parsedValue
        }

    } else if jsonValue, exists := json["location"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Location = &parsedValue
        }

    }

    if jsonValue, exists := json["Reason"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Reason = &parsedValue
        }

    } else if jsonValue, exists := json["reason"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Reason = &parsedValue
        }

    }

    if jsonValue, exists := json["SignatureDate"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.SignatureDate = new(Time)
            obj.SignatureDate.Parse(parsedValue)
        }

    } else if jsonValue, exists := json["signatureDate"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.SignatureDate = new(Time)
            obj.SignatureDate.Parse(parsedValue)
        }

    }
}

func (obj *PdfDigitalSignatureDetailsData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *PdfDigitalSignatureDetailsData) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    return nil;
}

func (obj *PdfDigitalSignatureDetailsData) GetCertificateFilename() *string {
    return obj.CertificateFilename
}

func (obj *PdfDigitalSignatureDetailsData) SetCertificateFilename(value *string) {
    obj.CertificateFilename = value
}

func (obj *PdfDigitalSignatureDetailsData) GetHashAlgorithm() *string {
    return obj.HashAlgorithm
}

func (obj *PdfDigitalSignatureDetailsData) SetHashAlgorithm(value *string) {
    obj.HashAlgorithm = value
}

func (obj *PdfDigitalSignatureDetailsData) GetLocation() *string {
    return obj.Location
}

func (obj *PdfDigitalSignatureDetailsData) SetLocation(value *string) {
    obj.Location = value
}

func (obj *PdfDigitalSignatureDetailsData) GetReason() *string {
    return obj.Reason
}

func (obj *PdfDigitalSignatureDetailsData) SetReason(value *string) {
    obj.Reason = value
}

func (obj *PdfDigitalSignatureDetailsData) GetSignatureDate() *Time {
    return obj.SignatureDate
}

func (obj *PdfDigitalSignatureDetailsData) SetSignatureDate(value *Time) {
    obj.SignatureDate = value
}

