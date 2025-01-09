/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="report_engine_settings.go">
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

// Report engine settings.

type IReportEngineSettings interface {
    IsReportEngineSettings() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetCsvDataLoadOptions() ICsvDataLoadOptions
    SetCsvDataLoadOptions(value ICsvDataLoadOptions)
    GetDataSourceName() *string
    SetDataSourceName(value *string)
    GetDataSourceType() *string
    SetDataSourceType(value *string)
    GetJsonDataLoadOptions() IJsonDataLoadOptions
    SetJsonDataLoadOptions(value IJsonDataLoadOptions)
    GetReportBuildOptions() []string
    SetReportBuildOptions(value []string)
    GetXmlDataLoadOptions() IXmlDataLoadOptions
    SetXmlDataLoadOptions(value IXmlDataLoadOptions)
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

func (ReportEngineSettings) IsReportEngineSettings() bool {
    return true
}


func (obj *ReportEngineSettings) Initialize() {
    if (obj.CsvDataLoadOptions != nil) {
        obj.CsvDataLoadOptions.Initialize()
    }

    if (obj.JsonDataLoadOptions != nil) {
        obj.JsonDataLoadOptions.Initialize()
    }

    if (obj.XmlDataLoadOptions != nil) {
        obj.XmlDataLoadOptions.Initialize()
    }


}

func (obj *ReportEngineSettings) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["CsvDataLoadOptions"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ICsvDataLoadOptions = new(CsvDataLoadOptions)
            modelInstance.Deserialize(parsedValue)
            obj.CsvDataLoadOptions = modelInstance
        }

    } else if jsonValue, exists := json["csvDataLoadOptions"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ICsvDataLoadOptions = new(CsvDataLoadOptions)
            modelInstance.Deserialize(parsedValue)
            obj.CsvDataLoadOptions = modelInstance
        }

    }

    if jsonValue, exists := json["DataSourceName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DataSourceName = &parsedValue
        }

    } else if jsonValue, exists := json["dataSourceName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DataSourceName = &parsedValue
        }

    }

    if jsonValue, exists := json["DataSourceType"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DataSourceType = &parsedValue
        }

    } else if jsonValue, exists := json["dataSourceType"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DataSourceType = &parsedValue
        }

    }

    if jsonValue, exists := json["JsonDataLoadOptions"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IJsonDataLoadOptions = new(JsonDataLoadOptions)
            modelInstance.Deserialize(parsedValue)
            obj.JsonDataLoadOptions = modelInstance
        }

    } else if jsonValue, exists := json["jsonDataLoadOptions"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IJsonDataLoadOptions = new(JsonDataLoadOptions)
            modelInstance.Deserialize(parsedValue)
            obj.JsonDataLoadOptions = modelInstance
        }

    }

    if jsonValue, exists := json["ReportBuildOptions"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.ReportBuildOptions = make([]string, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(string); valid {
                    obj.ReportBuildOptions = append(obj.ReportBuildOptions, elementValue)
                }

            }
        }

    } else if jsonValue, exists := json["reportBuildOptions"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.ReportBuildOptions = make([]string, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(string); valid {
                    obj.ReportBuildOptions = append(obj.ReportBuildOptions, elementValue)
                }

            }
        }

    }

    if jsonValue, exists := json["XmlDataLoadOptions"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IXmlDataLoadOptions = new(XmlDataLoadOptions)
            modelInstance.Deserialize(parsedValue)
            obj.XmlDataLoadOptions = modelInstance
        }

    } else if jsonValue, exists := json["xmlDataLoadOptions"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IXmlDataLoadOptions = new(XmlDataLoadOptions)
            modelInstance.Deserialize(parsedValue)
            obj.XmlDataLoadOptions = modelInstance
        }

    }
}

func (obj *ReportEngineSettings) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *ReportEngineSettings) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.DataSourceType == nil {
        return errors.New("Property DataSourceType in ReportEngineSettings is required.")
    }
    if obj.CsvDataLoadOptions != nil {
        if err := obj.CsvDataLoadOptions.Validate(); err != nil {
            return err
        }
    }
    if obj.JsonDataLoadOptions != nil {
        if err := obj.JsonDataLoadOptions.Validate(); err != nil {
            return err
        }
    }
    if obj.XmlDataLoadOptions != nil {
        if err := obj.XmlDataLoadOptions.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *ReportEngineSettings) GetCsvDataLoadOptions() ICsvDataLoadOptions {
    return obj.CsvDataLoadOptions
}

func (obj *ReportEngineSettings) SetCsvDataLoadOptions(value ICsvDataLoadOptions) {
    obj.CsvDataLoadOptions = value
}

func (obj *ReportEngineSettings) GetDataSourceName() *string {
    return obj.DataSourceName
}

func (obj *ReportEngineSettings) SetDataSourceName(value *string) {
    obj.DataSourceName = value
}

func (obj *ReportEngineSettings) GetDataSourceType() *string {
    return obj.DataSourceType
}

func (obj *ReportEngineSettings) SetDataSourceType(value *string) {
    obj.DataSourceType = value
}

func (obj *ReportEngineSettings) GetJsonDataLoadOptions() IJsonDataLoadOptions {
    return obj.JsonDataLoadOptions
}

func (obj *ReportEngineSettings) SetJsonDataLoadOptions(value IJsonDataLoadOptions) {
    obj.JsonDataLoadOptions = value
}

func (obj *ReportEngineSettings) GetReportBuildOptions() []string {
    return obj.ReportBuildOptions
}

func (obj *ReportEngineSettings) SetReportBuildOptions(value []string) {
    obj.ReportBuildOptions = value
}

func (obj *ReportEngineSettings) GetXmlDataLoadOptions() IXmlDataLoadOptions {
    return obj.XmlDataLoadOptions
}

func (obj *ReportEngineSettings) SetXmlDataLoadOptions(value IXmlDataLoadOptions) {
    obj.XmlDataLoadOptions = value
}

