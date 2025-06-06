/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="report_engine_settings.go">
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

// Report engine settings.
type ReportEngineSettingsResult struct {
    // Report engine settings.
    CsvDataLoadOptions CsvDataLoadOptionsResult `json:"CsvDataLoadOptions,omitempty"`

    // Report engine settings.
    DataSourceName string `json:"DataSourceName,omitempty"`

    // Report engine settings.
    DataSourceType string `json:"DataSourceType,omitempty"`

    // Report engine settings.
    JsonDataLoadOptions JsonDataLoadOptionsResult `json:"JsonDataLoadOptions,omitempty"`

    // Report engine settings.
    ReportBuildOptions []string `json:"ReportBuildOptions,omitempty"`

    // Report engine settings.
    XmlDataLoadOptions XmlDataLoadOptionsResult `json:"XmlDataLoadOptions,omitempty"`
}

type ReportEngineSettings struct {
    // Report engine settings.
    CsvDataLoadOptions ICsvDataLoadOptions `json:"CsvDataLoadOptions,omitempty"`

    // Report engine settings.
    DataSourceName *string `json:"DataSourceName,omitempty"`

    // Report engine settings.
    DataSourceType *string `json:"DataSourceType,omitempty"`

    // Report engine settings.
    JsonDataLoadOptions IJsonDataLoadOptions `json:"JsonDataLoadOptions,omitempty"`

    // Report engine settings.
    ReportBuildOptions []string `json:"ReportBuildOptions,omitempty"`

    // Report engine settings.
    XmlDataLoadOptions IXmlDataLoadOptions `json:"XmlDataLoadOptions,omitempty"`
}

type IReportEngineSettings interface {
    IsReportEngineSettings() bool
}
func (ReportEngineSettings) IsReportEngineSettings() bool {
    return true
}


