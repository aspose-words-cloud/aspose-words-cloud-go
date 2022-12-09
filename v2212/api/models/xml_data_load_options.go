/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="xml_data_load_options.go">
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

// Represents options for XML data loading.
// To learn more, visit the LINQ Reporting Engine documentation article.
type XmlDataLoadOptionsResult struct {
    // Represents options for XML data loading.
    // To learn more, visit the LINQ Reporting Engine documentation article.
    AlwaysGenerateRootObject bool `json:"AlwaysGenerateRootObject,omitempty"`
}

type XmlDataLoadOptions struct {
    // Represents options for XML data loading.
    // To learn more, visit the LINQ Reporting Engine documentation article.
    AlwaysGenerateRootObject *bool `json:"AlwaysGenerateRootObject,omitempty"`
}

type IXmlDataLoadOptions interface {
    IsXmlDataLoadOptions() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
}

func (XmlDataLoadOptions) IsXmlDataLoadOptions() bool {
    return true
}


func (obj *XmlDataLoadOptions) Initialize() {
}

func (obj *XmlDataLoadOptions) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}


