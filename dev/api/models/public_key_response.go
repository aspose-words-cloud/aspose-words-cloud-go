/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="public_key_response.go">
 *   Copyright (c) 2023 Aspose.Words for Cloud
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

// REST response for RSA public key info.

type IPublicKeyResponse interface {
    IsPublicKeyResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetRequestId() *string
    SetRequestId(value *string)
    GetExponent() *string
    SetExponent(value *string)
    GetModulus() *string
    SetModulus(value *string)
}

type PublicKeyResponse struct {
    // REST response for RSA public key info.
    RequestId *string `json:"RequestId,omitempty"`

    // REST response for RSA public key info.
    Exponent *string `json:"Exponent,omitempty"`

    // REST response for RSA public key info.
    Modulus *string `json:"Modulus,omitempty"`
}

func (PublicKeyResponse) IsPublicKeyResponse() bool {
    return true
}

func (PublicKeyResponse) IsWordsResponse() bool {
    return true
}

func (obj *PublicKeyResponse) Initialize() {
}

func (obj *PublicKeyResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["Exponent"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Exponent = &parsedValue
        }

    } else if jsonValue, exists := json["exponent"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Exponent = &parsedValue
        }

    }

    if jsonValue, exists := json["Modulus"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Modulus = &parsedValue
        }

    } else if jsonValue, exists := json["modulus"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Modulus = &parsedValue
        }

    }
}

func (obj *PublicKeyResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *PublicKeyResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *PublicKeyResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *PublicKeyResponse) GetExponent() *string {
    return obj.Exponent
}

func (obj *PublicKeyResponse) SetExponent(value *string) {
    obj.Exponent = value
}

func (obj *PublicKeyResponse) GetModulus() *string {
    return obj.Modulus
}

func (obj *PublicKeyResponse) SetModulus(value *string) {
    obj.Modulus = value
}

