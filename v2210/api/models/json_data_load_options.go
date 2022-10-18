/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="json_data_load_options.go">
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

// Represents options for parsing JSON data.
type JsonDataLoadOptionsResult struct {
    // Represents options for parsing JSON data.
    AlwaysGenerateRootObject bool `json:"AlwaysGenerateRootObject,omitempty"`

    // Represents options for parsing JSON data.
    ExactDateTimeParseFormats []string `json:"ExactDateTimeParseFormats,omitempty"`

    // Represents options for parsing JSON data.
    SimpleValueParseMode string `json:"SimpleValueParseMode,omitempty"`
}

type JsonDataLoadOptions struct {
    // Represents options for parsing JSON data.
    AlwaysGenerateRootObject *bool `json:"AlwaysGenerateRootObject,omitempty"`

    // Represents options for parsing JSON data.
    ExactDateTimeParseFormats []string `json:"ExactDateTimeParseFormats,omitempty"`

    // Represents options for parsing JSON data.
    SimpleValueParseMode *string `json:"SimpleValueParseMode,omitempty"`
}

type IJsonDataLoadOptions interface {
    IsJsonDataLoadOptions() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
}

func (JsonDataLoadOptions) IsJsonDataLoadOptions() bool {
    return true
}


func (obj *JsonDataLoadOptions) Initialize() {
}

func (obj *JsonDataLoadOptions) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}


