/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="digital_signature_details.go">
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
    "errors"
)

// Container class for details of digital signature.

type IDigitalSignatureDetails interface {
    IsDigitalSignatureDetails() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetCertificateFilename() *string
    SetCertificateFilename(value *string)
    GetSignOptions() ISignOptions
    SetSignOptions(value ISignOptions)
}

type DigitalSignatureDetails struct {
    // Container class for details of digital signature.
    CertificateFilename *string `json:"CertificateFilename,omitempty"`

    // Container class for details of digital signature.
    SignOptions ISignOptions `json:"SignOptions,omitempty"`
}

func (DigitalSignatureDetails) IsDigitalSignatureDetails() bool {
    return true
}


func (obj *DigitalSignatureDetails) Initialize() {
    if (obj.SignOptions != nil) {
        obj.SignOptions.Initialize()
    }


}

func (obj *DigitalSignatureDetails) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["CertificateFilename"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.CertificateFilename = &parsedValue
        }

    } else if jsonValue, exists := json["certificateFilename"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.CertificateFilename = &parsedValue
        }

    }

    if jsonValue, exists := json["SignOptions"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ISignOptions = new(SignOptions)
            modelInstance.Deserialize(parsedValue)
            obj.SignOptions = modelInstance
        }

    } else if jsonValue, exists := json["signOptions"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ISignOptions = new(SignOptions)
            modelInstance.Deserialize(parsedValue)
            obj.SignOptions = modelInstance
        }

    }
}

func (obj *DigitalSignatureDetails) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *DigitalSignatureDetails) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.SignOptions != nil {
        if err := obj.SignOptions.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *DigitalSignatureDetails) GetCertificateFilename() *string {
    return obj.CertificateFilename
}

func (obj *DigitalSignatureDetails) SetCertificateFilename(value *string) {
    obj.CertificateFilename = value
}

func (obj *DigitalSignatureDetails) GetSignOptions() ISignOptions {
    return obj.SignOptions
}

func (obj *DigitalSignatureDetails) SetSignOptions(value ISignOptions) {
    obj.SignOptions = value
}

