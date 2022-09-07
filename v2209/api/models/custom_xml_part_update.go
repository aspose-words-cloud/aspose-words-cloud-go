/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="custom_xml_part_update.go">
 *   Copyright (c) 2022 Aspose.Words for Cloud
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

// Custom xml part update.
type CustomXmlPartUpdateResult struct {
    // Custom xml part update.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // Custom xml part update.
    Data string `json:"Data,omitempty"`

    // Custom xml part update.
    Id string `json:"Id,omitempty"`
}

type CustomXmlPartUpdate struct {
    // Custom xml part update.
    Link IWordsApiLink `json:"Link,omitempty"`

    // Custom xml part update.
    Data *string `json:"Data,omitempty"`

    // Custom xml part update.
    Id *string `json:"Id,omitempty"`
}

type ICustomXmlPartUpdate interface {
    IsCustomXmlPartUpdate() bool
    Initialize()
}

func (CustomXmlPartUpdate) IsCustomXmlPartUpdate() bool {
    return true
}

func (CustomXmlPartUpdate) IsCustomXmlPart() bool {
    return true
}

func (CustomXmlPartUpdate) IsCustomXmlPartLink() bool {
    return true
}

func (CustomXmlPartUpdate) IsLinkElement() bool {
    return true
}

func (obj *CustomXmlPartUpdate) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }





}


