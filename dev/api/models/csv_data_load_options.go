/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="csv_data_load_options.go">
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

import (
    "errors"
)

// Represents options for parsing CSV data.
// To learn more, visit the LINQ Reporting Engine documentation article.
// An instance of this class can be passed into constructors of CsvDataSource.

type ICsvDataLoadOptions interface {
    IsCsvDataLoadOptions() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetCommentChar() *string
    SetCommentChar(value *string)
    GetDelimiter() *string
    SetDelimiter(value *string)
    GetHasHeaders() *bool
    SetHasHeaders(value *bool)
    GetQuoteChar() *string
    SetQuoteChar(value *string)
}

type CsvDataLoadOptions struct {
    // Represents options for parsing CSV data.
    // To learn more, visit the LINQ Reporting Engine documentation article.
    // An instance of this class can be passed into constructors of CsvDataSource.
    CommentChar *string `json:"CommentChar,omitempty"`

    // Represents options for parsing CSV data.
    // To learn more, visit the LINQ Reporting Engine documentation article.
    // An instance of this class can be passed into constructors of CsvDataSource.
    Delimiter *string `json:"Delimiter,omitempty"`

    // Represents options for parsing CSV data.
    // To learn more, visit the LINQ Reporting Engine documentation article.
    // An instance of this class can be passed into constructors of CsvDataSource.
    HasHeaders *bool `json:"HasHeaders,omitempty"`

    // Represents options for parsing CSV data.
    // To learn more, visit the LINQ Reporting Engine documentation article.
    // An instance of this class can be passed into constructors of CsvDataSource.
    QuoteChar *string `json:"QuoteChar,omitempty"`
}

func (CsvDataLoadOptions) IsCsvDataLoadOptions() bool {
    return true
}


func (obj *CsvDataLoadOptions) Initialize() {
}

func (obj *CsvDataLoadOptions) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["CommentChar"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.CommentChar = &parsedValue
        }

    } else if jsonValue, exists := json["commentChar"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.CommentChar = &parsedValue
        }

    }

    if jsonValue, exists := json["Delimiter"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Delimiter = &parsedValue
        }

    } else if jsonValue, exists := json["delimiter"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Delimiter = &parsedValue
        }

    }

    if jsonValue, exists := json["HasHeaders"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.HasHeaders = &parsedValue
        }

    } else if jsonValue, exists := json["hasHeaders"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.HasHeaders = &parsedValue
        }

    }

    if jsonValue, exists := json["QuoteChar"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.QuoteChar = &parsedValue
        }

    } else if jsonValue, exists := json["quoteChar"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.QuoteChar = &parsedValue
        }

    }
}

func (obj *CsvDataLoadOptions) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *CsvDataLoadOptions) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.CommentChar == nil {
        return errors.New("Property CommentChar in CsvDataLoadOptions is required.")
    }

    if obj.Delimiter == nil {
        return errors.New("Property Delimiter in CsvDataLoadOptions is required.")
    }

    if obj.HasHeaders == nil {
        return errors.New("Property HasHeaders in CsvDataLoadOptions is required.")
    }

    if obj.QuoteChar == nil {
        return errors.New("Property QuoteChar in CsvDataLoadOptions is required.")
    }

    return nil;
}

func (obj *CsvDataLoadOptions) GetCommentChar() *string {
    return obj.CommentChar
}

func (obj *CsvDataLoadOptions) SetCommentChar(value *string) {
    obj.CommentChar = value
}

func (obj *CsvDataLoadOptions) GetDelimiter() *string {
    return obj.Delimiter
}

func (obj *CsvDataLoadOptions) SetDelimiter(value *string) {
    obj.Delimiter = value
}

func (obj *CsvDataLoadOptions) GetHasHeaders() *bool {
    return obj.HasHeaders
}

func (obj *CsvDataLoadOptions) SetHasHeaders(value *bool) {
    obj.HasHeaders = value
}

func (obj *CsvDataLoadOptions) GetQuoteChar() *string {
    return obj.QuoteChar
}

func (obj *CsvDataLoadOptions) SetQuoteChar(value *string) {
    obj.QuoteChar = value
}

