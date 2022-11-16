/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="custom_xml_parts_collection.go">
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

// The collection of CustomXmlPart.
type CustomXmlPartsCollectionResult struct {
    // The collection of CustomXmlPart.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // The collection of CustomXmlPart.
    CustomXmlPartsList []CustomXmlPartResult `json:"CustomXmlPartsList,omitempty"`
}

type CustomXmlPartsCollection struct {
    // The collection of CustomXmlPart.
    Link IWordsApiLink `json:"Link,omitempty"`

    // The collection of CustomXmlPart.
    CustomXmlPartsList []CustomXmlPart `json:"CustomXmlPartsList,omitempty"`
}

type ICustomXmlPartsCollection interface {
    IsCustomXmlPartsCollection() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
}

func (CustomXmlPartsCollection) IsCustomXmlPartsCollection() bool {
    return true
}

func (CustomXmlPartsCollection) IsLinkElement() bool {
    return true
}

func (obj *CustomXmlPartsCollection) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.CustomXmlPartsList != nil) {
        for _, objElementCustomXmlPartsList := range obj.CustomXmlPartsList {
            objElementCustomXmlPartsList.Initialize()
        }
    }

}

func (obj *CustomXmlPartsCollection) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}


