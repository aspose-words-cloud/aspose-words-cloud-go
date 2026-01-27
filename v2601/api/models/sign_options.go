/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="sign_options.go">
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

// Container class for digital signature options.

type ISignOptions interface {
    IsSignOptions() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetComments() *string
    SetComments(value *string)
    GetDecryptionPassword() *string
    SetDecryptionPassword(value *string)
    GetProviderId() *string
    SetProviderId(value *string)
    GetSignatureLineId() *string
    SetSignatureLineId(value *string)
    GetSignatureLineImageFilename() *string
    SetSignatureLineImageFilename(value *string)
    GetSignTime() *Time
    SetSignTime(value *Time)
}

type SignOptions struct {
    // Container class for digital signature options.
    Comments *string `json:"Comments,omitempty"`

    // Container class for digital signature options.
    DecryptionPassword *string `json:"DecryptionPassword,omitempty"`

    // Container class for digital signature options.
    ProviderId *string `json:"ProviderId,omitempty"`

    // Container class for digital signature options.
    SignatureLineId *string `json:"SignatureLineId,omitempty"`

    // Container class for digital signature options.
    SignatureLineImageFilename *string `json:"SignatureLineImageFilename,omitempty"`

    // Container class for digital signature options.
    SignTime *Time `json:"SignTime,omitempty"`
}

func (SignOptions) IsSignOptions() bool {
    return true
}


func (obj *SignOptions) Initialize() {
}

func (obj *SignOptions) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Comments"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Comments = &parsedValue
        }

    } else if jsonValue, exists := json["comments"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Comments = &parsedValue
        }

    }

    if jsonValue, exists := json["DecryptionPassword"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DecryptionPassword = &parsedValue
        }

    } else if jsonValue, exists := json["decryptionPassword"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DecryptionPassword = &parsedValue
        }

    }

    if jsonValue, exists := json["ProviderId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ProviderId = &parsedValue
        }

    } else if jsonValue, exists := json["providerId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ProviderId = &parsedValue
        }

    }

    if jsonValue, exists := json["SignatureLineId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.SignatureLineId = &parsedValue
        }

    } else if jsonValue, exists := json["signatureLineId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.SignatureLineId = &parsedValue
        }

    }

    if jsonValue, exists := json["SignatureLineImageFilename"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.SignatureLineImageFilename = &parsedValue
        }

    } else if jsonValue, exists := json["signatureLineImageFilename"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.SignatureLineImageFilename = &parsedValue
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
}

func (obj *SignOptions) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *SignOptions) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    return nil;
}

func (obj *SignOptions) GetComments() *string {
    return obj.Comments
}

func (obj *SignOptions) SetComments(value *string) {
    obj.Comments = value
}

func (obj *SignOptions) GetDecryptionPassword() *string {
    return obj.DecryptionPassword
}

func (obj *SignOptions) SetDecryptionPassword(value *string) {
    obj.DecryptionPassword = value
}

func (obj *SignOptions) GetProviderId() *string {
    return obj.ProviderId
}

func (obj *SignOptions) SetProviderId(value *string) {
    obj.ProviderId = value
}

func (obj *SignOptions) GetSignatureLineId() *string {
    return obj.SignatureLineId
}

func (obj *SignOptions) SetSignatureLineId(value *string) {
    obj.SignatureLineId = value
}

func (obj *SignOptions) GetSignatureLineImageFilename() *string {
    return obj.SignatureLineImageFilename
}

func (obj *SignOptions) SetSignatureLineImageFilename(value *string) {
    obj.SignatureLineImageFilename = value
}

func (obj *SignOptions) GetSignTime() *Time {
    return obj.SignTime
}

func (obj *SignOptions) SetSignTime(value *Time) {
    obj.SignTime = value
}

