/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="signature.go">
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

// The REST response with a document signature collection.
// This response is returned by the Service when handling any "https://api.aspose.cloud/v4.0/words/Test.doc/signatures" REST API requests.

type ISignature interface {
    IsSignature() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetComments() *string
    SetComments(value *string)
    GetIssuerName() *string
    SetIssuerName(value *string)
    GetIsValid() *bool
    SetIsValid(value *bool)
    GetSignatureType() *string
    SetSignatureType(value *string)
    GetSignatureValue() *string
    SetSignatureValue(value *string)
    GetSignTime() *Time
    SetSignTime(value *Time)
    GetSubjectName() *string
    SetSubjectName(value *string)
}

type Signature struct {
    // The REST response with a document signature collection.
    // This response is returned by the Service when handling any "https://api.aspose.cloud/v4.0/words/Test.doc/signatures" REST API requests.
    Comments *string `json:"Comments,omitempty"`

    // The REST response with a document signature collection.
    // This response is returned by the Service when handling any "https://api.aspose.cloud/v4.0/words/Test.doc/signatures" REST API requests.
    IssuerName *string `json:"IssuerName,omitempty"`

    // The REST response with a document signature collection.
    // This response is returned by the Service when handling any "https://api.aspose.cloud/v4.0/words/Test.doc/signatures" REST API requests.
    IsValid *bool `json:"IsValid,omitempty"`

    // The REST response with a document signature collection.
    // This response is returned by the Service when handling any "https://api.aspose.cloud/v4.0/words/Test.doc/signatures" REST API requests.
    SignatureType *string `json:"SignatureType,omitempty"`

    // The REST response with a document signature collection.
    // This response is returned by the Service when handling any "https://api.aspose.cloud/v4.0/words/Test.doc/signatures" REST API requests.
    SignatureValue *string `json:"SignatureValue,omitempty"`

    // The REST response with a document signature collection.
    // This response is returned by the Service when handling any "https://api.aspose.cloud/v4.0/words/Test.doc/signatures" REST API requests.
    SignTime *Time `json:"SignTime,omitempty"`

    // The REST response with a document signature collection.
    // This response is returned by the Service when handling any "https://api.aspose.cloud/v4.0/words/Test.doc/signatures" REST API requests.
    SubjectName *string `json:"SubjectName,omitempty"`
}

func (Signature) IsSignature() bool {
    return true
}


func (obj *Signature) Initialize() {
}

func (obj *Signature) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Comments"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Comments = &parsedValue
        }

    } else if jsonValue, exists := json["comments"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Comments = &parsedValue
        }

    }

    if jsonValue, exists := json["IssuerName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.IssuerName = &parsedValue
        }

    } else if jsonValue, exists := json["issuerName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.IssuerName = &parsedValue
        }

    }

    if jsonValue, exists := json["IsValid"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsValid = &parsedValue
        }

    } else if jsonValue, exists := json["isValid"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsValid = &parsedValue
        }

    }

    if jsonValue, exists := json["SignatureType"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.SignatureType = &parsedValue
        }

    } else if jsonValue, exists := json["signatureType"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.SignatureType = &parsedValue
        }

    }

    if jsonValue, exists := json["SignatureValue"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.SignatureValue = &parsedValue
        }

    } else if jsonValue, exists := json["signatureValue"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.SignatureValue = &parsedValue
        }

    }

    if jsonValue, exists := json["SignTime"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.SignTime = new(Time)
            obj.SignTime.Parse(parsedValue)
        }

    } else if jsonValue, exists := json["signTime"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.SignTime = new(Time)
            obj.SignTime.Parse(parsedValue)
        }

    }

    if jsonValue, exists := json["SubjectName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.SubjectName = &parsedValue
        }

    } else if jsonValue, exists := json["subjectName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.SubjectName = &parsedValue
        }

    }
}

func (obj *Signature) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *Signature) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.IsValid == nil {
        return errors.New("Property IsValid in Signature is required.")
    }
    if obj.SignTime == nil {
        return errors.New("Property SignTime in Signature is required.")
    }
    return nil;
}

func (obj *Signature) GetComments() *string {
    return obj.Comments
}

func (obj *Signature) SetComments(value *string) {
    obj.Comments = value
}

func (obj *Signature) GetIssuerName() *string {
    return obj.IssuerName
}

func (obj *Signature) SetIssuerName(value *string) {
    obj.IssuerName = value
}

func (obj *Signature) GetIsValid() *bool {
    return obj.IsValid
}

func (obj *Signature) SetIsValid(value *bool) {
    obj.IsValid = value
}

func (obj *Signature) GetSignatureType() *string {
    return obj.SignatureType
}

func (obj *Signature) SetSignatureType(value *string) {
    obj.SignatureType = value
}

func (obj *Signature) GetSignatureValue() *string {
    return obj.SignatureValue
}

func (obj *Signature) SetSignatureValue(value *string) {
    obj.SignatureValue = value
}

func (obj *Signature) GetSignTime() *Time {
    return obj.SignTime
}

func (obj *Signature) SetSignTime(value *Time) {
    obj.SignTime = value
}

func (obj *Signature) GetSubjectName() *string {
    return obj.SubjectName
}

func (obj *Signature) SetSubjectName(value *string) {
    obj.SubjectName = value
}

