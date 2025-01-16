/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="load_web_document_data.go">
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

// Contains data for load web document.

type ILoadWebDocumentData interface {
    IsLoadWebDocumentData() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetSaveOptions() ISaveOptionsData
    SetSaveOptions(value ISaveOptionsData)
    GetLoadingDocumentUrl() *string
    SetLoadingDocumentUrl(value *string)
}

type LoadWebDocumentData struct {
    // Contains data for load web document.
    SaveOptions ISaveOptionsData `json:"SaveOptions,omitempty"`

    // Contains data for load web document.
    LoadingDocumentUrl *string `json:"LoadingDocumentUrl,omitempty"`
}

func (LoadWebDocumentData) IsLoadWebDocumentData() bool {
    return true
}


func (obj *LoadWebDocumentData) Initialize() {
    if (obj.SaveOptions != nil) {
        obj.SaveOptions.Initialize()
    }


}

func (obj *LoadWebDocumentData) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["SaveOptions"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ISaveOptionsData = nil
            if jsonType, found := parsedValue["$type"]; found {
                jsonTypeStr := jsonType.(string)
                if jsonTypeStr == "Azw3SaveOptionsData, _" { modelInstance = new(Azw3SaveOptionsData) }
                if jsonTypeStr == "BmpSaveOptionsData, _" { modelInstance = new(BmpSaveOptionsData) }
                if jsonTypeStr == "DocmSaveOptionsData, _" { modelInstance = new(DocmSaveOptionsData) }
                if jsonTypeStr == "DocSaveOptionsData, _" { modelInstance = new(DocSaveOptionsData) }
                if jsonTypeStr == "DocxSaveOptionsData, _" { modelInstance = new(DocxSaveOptionsData) }
                if jsonTypeStr == "DotmSaveOptionsData, _" { modelInstance = new(DotmSaveOptionsData) }
                if jsonTypeStr == "DotSaveOptionsData, _" { modelInstance = new(DotSaveOptionsData) }
                if jsonTypeStr == "DotxSaveOptionsData, _" { modelInstance = new(DotxSaveOptionsData) }
                if jsonTypeStr == "EmfSaveOptionsData, _" { modelInstance = new(EmfSaveOptionsData) }
                if jsonTypeStr == "EpsSaveOptionsData, _" { modelInstance = new(EpsSaveOptionsData) }
                if jsonTypeStr == "EpubSaveOptionsData, _" { modelInstance = new(EpubSaveOptionsData) }
                if jsonTypeStr == "FixedPageSaveOptionsData, _" {  }
                if jsonTypeStr == "FlatOpcMacroSaveOptionsData, _" { modelInstance = new(FlatOpcMacroSaveOptionsData) }
                if jsonTypeStr == "FlatOpcSaveOptionsData, _" { modelInstance = new(FlatOpcSaveOptionsData) }
                if jsonTypeStr == "FlatOpcTemplateMacroSaveOptionsData, _" { modelInstance = new(FlatOpcTemplateMacroSaveOptionsData) }
                if jsonTypeStr == "FlatOpcTemplateSaveOptionsData, _" { modelInstance = new(FlatOpcTemplateSaveOptionsData) }
                if jsonTypeStr == "GifSaveOptionsData, _" { modelInstance = new(GifSaveOptionsData) }
                if jsonTypeStr == "HtmlFixedSaveOptionsData, _" { modelInstance = new(HtmlFixedSaveOptionsData) }
                if jsonTypeStr == "HtmlSaveOptionsData, _" { modelInstance = new(HtmlSaveOptionsData) }
                if jsonTypeStr == "ImageSaveOptionsData, _" {  }
                if jsonTypeStr == "JpegSaveOptionsData, _" { modelInstance = new(JpegSaveOptionsData) }
                if jsonTypeStr == "MarkdownSaveOptionsData, _" { modelInstance = new(MarkdownSaveOptionsData) }
                if jsonTypeStr == "MhtmlSaveOptionsData, _" { modelInstance = new(MhtmlSaveOptionsData) }
                if jsonTypeStr == "OdtSaveOptionsData, _" { modelInstance = new(OdtSaveOptionsData) }
                if jsonTypeStr == "OoxmlSaveOptionsData, _" {  }
                if jsonTypeStr == "OpenXpsSaveOptionsData, _" { modelInstance = new(OpenXpsSaveOptionsData) }
                if jsonTypeStr == "OttSaveOptionsData, _" { modelInstance = new(OttSaveOptionsData) }
                if jsonTypeStr == "PclSaveOptionsData, _" { modelInstance = new(PclSaveOptionsData) }
                if jsonTypeStr == "PdfSaveOptionsData, _" { modelInstance = new(PdfSaveOptionsData) }
                if jsonTypeStr == "PngSaveOptionsData, _" { modelInstance = new(PngSaveOptionsData) }
                if jsonTypeStr == "PsSaveOptionsData, _" { modelInstance = new(PsSaveOptionsData) }
                if jsonTypeStr == "RtfSaveOptionsData, _" { modelInstance = new(RtfSaveOptionsData) }
                if jsonTypeStr == "SvgSaveOptionsData, _" { modelInstance = new(SvgSaveOptionsData) }
                if jsonTypeStr == "TextSaveOptionsData, _" { modelInstance = new(TextSaveOptionsData) }
                if jsonTypeStr == "TiffSaveOptionsData, _" { modelInstance = new(TiffSaveOptionsData) }
                if jsonTypeStr == "TxtSaveOptionsBaseData, _" {  }
                if jsonTypeStr == "WordMLSaveOptionsData, _" { modelInstance = new(WordMLSaveOptionsData) }
                if jsonTypeStr == "XamlFixedSaveOptionsData, _" { modelInstance = new(XamlFixedSaveOptionsData) }
                if jsonTypeStr == "XamlFlowPackSaveOptionsData, _" { modelInstance = new(XamlFlowPackSaveOptionsData) }
                if jsonTypeStr == "XamlFlowSaveOptionsData, _" { modelInstance = new(XamlFlowSaveOptionsData) }
                if jsonTypeStr == "XpsSaveOptionsData, _" { modelInstance = new(XpsSaveOptionsData) }
            }

            modelInstance.Deserialize(parsedValue)
            obj.SaveOptions = modelInstance
        }

    } else if jsonValue, exists := json["saveOptions"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ISaveOptionsData = nil
            if jsonType, found := parsedValue["$type"]; found {
                jsonTypeStr := jsonType.(string)
                if jsonTypeStr == "Azw3SaveOptionsData, _" { modelInstance = new(Azw3SaveOptionsData) }
                if jsonTypeStr == "BmpSaveOptionsData, _" { modelInstance = new(BmpSaveOptionsData) }
                if jsonTypeStr == "DocmSaveOptionsData, _" { modelInstance = new(DocmSaveOptionsData) }
                if jsonTypeStr == "DocSaveOptionsData, _" { modelInstance = new(DocSaveOptionsData) }
                if jsonTypeStr == "DocxSaveOptionsData, _" { modelInstance = new(DocxSaveOptionsData) }
                if jsonTypeStr == "DotmSaveOptionsData, _" { modelInstance = new(DotmSaveOptionsData) }
                if jsonTypeStr == "DotSaveOptionsData, _" { modelInstance = new(DotSaveOptionsData) }
                if jsonTypeStr == "DotxSaveOptionsData, _" { modelInstance = new(DotxSaveOptionsData) }
                if jsonTypeStr == "EmfSaveOptionsData, _" { modelInstance = new(EmfSaveOptionsData) }
                if jsonTypeStr == "EpsSaveOptionsData, _" { modelInstance = new(EpsSaveOptionsData) }
                if jsonTypeStr == "EpubSaveOptionsData, _" { modelInstance = new(EpubSaveOptionsData) }
                if jsonTypeStr == "FixedPageSaveOptionsData, _" {  }
                if jsonTypeStr == "FlatOpcMacroSaveOptionsData, _" { modelInstance = new(FlatOpcMacroSaveOptionsData) }
                if jsonTypeStr == "FlatOpcSaveOptionsData, _" { modelInstance = new(FlatOpcSaveOptionsData) }
                if jsonTypeStr == "FlatOpcTemplateMacroSaveOptionsData, _" { modelInstance = new(FlatOpcTemplateMacroSaveOptionsData) }
                if jsonTypeStr == "FlatOpcTemplateSaveOptionsData, _" { modelInstance = new(FlatOpcTemplateSaveOptionsData) }
                if jsonTypeStr == "GifSaveOptionsData, _" { modelInstance = new(GifSaveOptionsData) }
                if jsonTypeStr == "HtmlFixedSaveOptionsData, _" { modelInstance = new(HtmlFixedSaveOptionsData) }
                if jsonTypeStr == "HtmlSaveOptionsData, _" { modelInstance = new(HtmlSaveOptionsData) }
                if jsonTypeStr == "ImageSaveOptionsData, _" {  }
                if jsonTypeStr == "JpegSaveOptionsData, _" { modelInstance = new(JpegSaveOptionsData) }
                if jsonTypeStr == "MarkdownSaveOptionsData, _" { modelInstance = new(MarkdownSaveOptionsData) }
                if jsonTypeStr == "MhtmlSaveOptionsData, _" { modelInstance = new(MhtmlSaveOptionsData) }
                if jsonTypeStr == "OdtSaveOptionsData, _" { modelInstance = new(OdtSaveOptionsData) }
                if jsonTypeStr == "OoxmlSaveOptionsData, _" {  }
                if jsonTypeStr == "OpenXpsSaveOptionsData, _" { modelInstance = new(OpenXpsSaveOptionsData) }
                if jsonTypeStr == "OttSaveOptionsData, _" { modelInstance = new(OttSaveOptionsData) }
                if jsonTypeStr == "PclSaveOptionsData, _" { modelInstance = new(PclSaveOptionsData) }
                if jsonTypeStr == "PdfSaveOptionsData, _" { modelInstance = new(PdfSaveOptionsData) }
                if jsonTypeStr == "PngSaveOptionsData, _" { modelInstance = new(PngSaveOptionsData) }
                if jsonTypeStr == "PsSaveOptionsData, _" { modelInstance = new(PsSaveOptionsData) }
                if jsonTypeStr == "RtfSaveOptionsData, _" { modelInstance = new(RtfSaveOptionsData) }
                if jsonTypeStr == "SvgSaveOptionsData, _" { modelInstance = new(SvgSaveOptionsData) }
                if jsonTypeStr == "TextSaveOptionsData, _" { modelInstance = new(TextSaveOptionsData) }
                if jsonTypeStr == "TiffSaveOptionsData, _" { modelInstance = new(TiffSaveOptionsData) }
                if jsonTypeStr == "TxtSaveOptionsBaseData, _" {  }
                if jsonTypeStr == "WordMLSaveOptionsData, _" { modelInstance = new(WordMLSaveOptionsData) }
                if jsonTypeStr == "XamlFixedSaveOptionsData, _" { modelInstance = new(XamlFixedSaveOptionsData) }
                if jsonTypeStr == "XamlFlowPackSaveOptionsData, _" { modelInstance = new(XamlFlowPackSaveOptionsData) }
                if jsonTypeStr == "XamlFlowSaveOptionsData, _" { modelInstance = new(XamlFlowSaveOptionsData) }
                if jsonTypeStr == "XpsSaveOptionsData, _" { modelInstance = new(XpsSaveOptionsData) }
            }

            modelInstance.Deserialize(parsedValue)
            obj.SaveOptions = modelInstance
        }

    }

    if jsonValue, exists := json["LoadingDocumentUrl"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.LoadingDocumentUrl = &parsedValue
        }

    } else if jsonValue, exists := json["loadingDocumentUrl"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.LoadingDocumentUrl = &parsedValue
        }

    }
}

func (obj *LoadWebDocumentData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *LoadWebDocumentData) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.LoadingDocumentUrl == nil {
        return errors.New("Property LoadingDocumentUrl in LoadWebDocumentData is required.")
    }
    if obj.SaveOptions != nil {
        if err := obj.SaveOptions.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *LoadWebDocumentData) GetSaveOptions() ISaveOptionsData {
    return obj.SaveOptions
}

func (obj *LoadWebDocumentData) SetSaveOptions(value ISaveOptionsData) {
    obj.SaveOptions = value
}

func (obj *LoadWebDocumentData) GetLoadingDocumentUrl() *string {
    return obj.LoadingDocumentUrl
}

func (obj *LoadWebDocumentData) SetLoadingDocumentUrl(value *string) {
    obj.LoadingDocumentUrl = value
}

