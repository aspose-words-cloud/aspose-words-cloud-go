/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="signature_collection_response.go">
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

package models

import (
    "errors"
)

// The REST response with a document signature collection.
// This response is returned by the Service when handling any "https://api.aspose.cloud/v4.0/words/Test.doc/signatures" REST API requests.

type ISignatureCollectionResponse interface {
    IsSignatureCollectionResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetIsValid() *bool
    SetIsValid(value *bool)
    GetSignatures() []ISignature
    SetSignatures(value []ISignature)
}

type SignatureCollectionResponse struct {
    // The REST response with a document signature collection.
    // This response is returned by the Service when handling any "https://api.aspose.cloud/v4.0/words/Test.doc/signatures" REST API requests.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a document signature collection.
    // This response is returned by the Service when handling any "https://api.aspose.cloud/v4.0/words/Test.doc/signatures" REST API requests.
    IsValid *bool `json:"IsValid,omitempty"`

    // The REST response with a document signature collection.
    // This response is returned by the Service when handling any "https://api.aspose.cloud/v4.0/words/Test.doc/signatures" REST API requests.
    Signatures []ISignature `json:"Signatures,omitempty"`
}

func (SignatureCollectionResponse) IsSignatureCollectionResponse() bool {
    return true
}

func (SignatureCollectionResponse) IsWordsResponse() bool {
    return true
}

func (obj *SignatureCollectionResponse) Initialize() {
    if (obj.Signatures != nil) {
        for _, objElementSignatures := range obj.Signatures {
            objElementSignatures.Initialize()
        }
    }

}

func (obj *SignatureCollectionResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
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

    if jsonValue, exists := json["Signatures"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.Signatures = make([]ISignature, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance ISignature = new(Signature)
                    modelElementInstance.Deserialize(elementValue)
                    obj.Signatures = append(obj.Signatures, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["signatures"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.Signatures = make([]ISignature, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance ISignature = new(Signature)
                    modelElementInstance.Deserialize(elementValue)
                    obj.Signatures = append(obj.Signatures, modelElementInstance)
                }

            }
        }

    }
}

func (obj *SignatureCollectionResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *SignatureCollectionResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.IsValid == nil {
        return errors.New("Property IsValid in SignatureCollectionResponse is required.")
    }
    if obj.Signatures != nil {
        for _, elementSignatures := range obj.Signatures {
            if elementSignatures != nil {
                if err := elementSignatures.Validate(); err != nil {
                    return err
                }
            }
        }
    }

    return nil;
}

func (obj *SignatureCollectionResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *SignatureCollectionResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *SignatureCollectionResponse) GetIsValid() *bool {
    return obj.IsValid
}

func (obj *SignatureCollectionResponse) SetIsValid(value *bool) {
    obj.IsValid = value
}

func (obj *SignatureCollectionResponse) GetSignatures() []ISignature {
    return obj.Signatures
}

func (obj *SignatureCollectionResponse) SetSignatures(value []ISignature) {
    obj.Signatures = value
}

