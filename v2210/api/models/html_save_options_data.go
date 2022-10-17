/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="html_save_options_data.go">
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

// Container class for html save options.
type HtmlSaveOptionsDataResult struct {
    // Container class for html save options.
    AllowEmbeddingPostScriptFonts bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for html save options.
    CustomTimeZoneInfoData TimeZoneInfoDataResult `json:"CustomTimeZoneInfoData,omitempty"`

    // Container class for html save options.
    Dml3DEffectsRenderingMode string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for html save options.
    DmlEffectsRenderingMode string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for html save options.
    DmlRenderingMode string `json:"DmlRenderingMode,omitempty"`

    // Container class for html save options.
    FileName string `json:"FileName,omitempty"`

    // Container class for html save options.
    ImlRenderingMode string `json:"ImlRenderingMode,omitempty"`

    // Container class for html save options.
    UpdateCreatedTimeProperty bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for html save options.
    UpdateFields bool `json:"UpdateFields,omitempty"`

    // Container class for html save options.
    UpdateLastPrintedProperty bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for html save options.
    UpdateLastSavedTimeProperty bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for html save options.
    UpdateSdtContent bool `json:"UpdateSdtContent,omitempty"`

    // Container class for html save options.
    ZipOutput bool `json:"ZipOutput,omitempty"`

    // Container class for html save options.
    AllowNegativeIndent bool `json:"AllowNegativeIndent,omitempty"`

    // Container class for html save options.
    CssClassNamePrefix string `json:"CssClassNamePrefix,omitempty"`

    // Container class for html save options.
    CssStyleSheetFileName string `json:"CssStyleSheetFileName,omitempty"`

    // Container class for html save options.
    CssStyleSheetType string `json:"CssStyleSheetType,omitempty"`

    // Container class for html save options.
    DocumentSplitCriteria string `json:"DocumentSplitCriteria,omitempty"`

    // Container class for html save options.
    DocumentSplitHeadingLevel int32 `json:"DocumentSplitHeadingLevel,omitempty"`

    // Container class for html save options.
    Encoding string `json:"Encoding,omitempty"`

    // Container class for html save options.
    ExportDocumentProperties bool `json:"ExportDocumentProperties,omitempty"`

    // Container class for html save options.
    ExportDropDownFormFieldAsText bool `json:"ExportDropDownFormFieldAsText,omitempty"`

    // Container class for html save options.
    ExportFontResources bool `json:"ExportFontResources,omitempty"`

    // Container class for html save options.
    ExportFontsAsBase64 bool `json:"ExportFontsAsBase64,omitempty"`

    // Container class for html save options.
    ExportHeadersFootersMode string `json:"ExportHeadersFootersMode,omitempty"`

    // Container class for html save options.
    ExportImagesAsBase64 bool `json:"ExportImagesAsBase64,omitempty"`

    // Container class for html save options.
    ExportLanguageInformation bool `json:"ExportLanguageInformation,omitempty"`

    // Container class for html save options.
    ExportListLabels string `json:"ExportListLabels,omitempty"`

    // Container class for html save options.
    ExportOriginalUrlForLinkedImages bool `json:"ExportOriginalUrlForLinkedImages,omitempty"`

    // Container class for html save options.
    ExportPageMargins bool `json:"ExportPageMargins,omitempty"`

    // Container class for html save options.
    ExportPageSetup bool `json:"ExportPageSetup,omitempty"`

    // Container class for html save options.
    ExportRelativeFontSize bool `json:"ExportRelativeFontSize,omitempty"`

    // Container class for html save options.
    ExportRoundtripInformation bool `json:"ExportRoundtripInformation,omitempty"`

    // Container class for html save options.
    ExportTextInputFormFieldAsText bool `json:"ExportTextInputFormFieldAsText,omitempty"`

    // Container class for html save options.
    ExportTocPageNumbers bool `json:"ExportTocPageNumbers,omitempty"`

    // Container class for html save options.
    ExportXhtmlTransitional bool `json:"ExportXhtmlTransitional,omitempty"`

    // Container class for html save options.
    FontResourcesSubsettingSizeThreshold int32 `json:"FontResourcesSubsettingSizeThreshold,omitempty"`

    // Container class for html save options.
    FontsFolder string `json:"FontsFolder,omitempty"`

    // Container class for html save options.
    FontsFolderAlias string `json:"FontsFolderAlias,omitempty"`

    // Container class for html save options.
    HtmlVersion string `json:"HtmlVersion,omitempty"`

    // Container class for html save options.
    ImageResolution int32 `json:"ImageResolution,omitempty"`

    // Container class for html save options.
    ImagesFolder string `json:"ImagesFolder,omitempty"`

    // Container class for html save options.
    ImagesFolderAlias string `json:"ImagesFolderAlias,omitempty"`

    // Container class for html save options.
    MetafileFormat string `json:"MetafileFormat,omitempty"`

    // Container class for html save options.
    OfficeMathOutputMode string `json:"OfficeMathOutputMode,omitempty"`

    // Container class for html save options.
    PrettyFormat bool `json:"PrettyFormat,omitempty"`

    // Container class for html save options.
    ResolveFontNames bool `json:"ResolveFontNames,omitempty"`

    // Container class for html save options.
    ResourceFolder string `json:"ResourceFolder,omitempty"`

    // Container class for html save options.
    ResourceFolderAlias string `json:"ResourceFolderAlias,omitempty"`

    // Container class for html save options.
    SaveFormat string `json:"SaveFormat,omitempty"`

    // Container class for html save options.
    ScaleImageToShapeSize bool `json:"ScaleImageToShapeSize,omitempty"`

    // Container class for html save options.
    TableWidthOutputMode string `json:"TableWidthOutputMode,omitempty"`
}

type HtmlSaveOptionsData struct {
    // Container class for html save options.
    AllowEmbeddingPostScriptFonts *bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for html save options.
    CustomTimeZoneInfoData ITimeZoneInfoData `json:"CustomTimeZoneInfoData,omitempty"`

    // Container class for html save options.
    Dml3DEffectsRenderingMode *string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for html save options.
    DmlEffectsRenderingMode *string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for html save options.
    DmlRenderingMode *string `json:"DmlRenderingMode,omitempty"`

    // Container class for html save options.
    FileName *string `json:"FileName,omitempty"`

    // Container class for html save options.
    ImlRenderingMode *string `json:"ImlRenderingMode,omitempty"`

    // Container class for html save options.
    UpdateCreatedTimeProperty *bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for html save options.
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // Container class for html save options.
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for html save options.
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for html save options.
    UpdateSdtContent *bool `json:"UpdateSdtContent,omitempty"`

    // Container class for html save options.
    ZipOutput *bool `json:"ZipOutput,omitempty"`

    // Container class for html save options.
    AllowNegativeIndent *bool `json:"AllowNegativeIndent,omitempty"`

    // Container class for html save options.
    CssClassNamePrefix *string `json:"CssClassNamePrefix,omitempty"`

    // Container class for html save options.
    CssStyleSheetFileName *string `json:"CssStyleSheetFileName,omitempty"`

    // Container class for html save options.
    CssStyleSheetType *string `json:"CssStyleSheetType,omitempty"`

    // Container class for html save options.
    DocumentSplitCriteria *string `json:"DocumentSplitCriteria,omitempty"`

    // Container class for html save options.
    DocumentSplitHeadingLevel *int32 `json:"DocumentSplitHeadingLevel,omitempty"`

    // Container class for html save options.
    Encoding *string `json:"Encoding,omitempty"`

    // Container class for html save options.
    ExportDocumentProperties *bool `json:"ExportDocumentProperties,omitempty"`

    // Container class for html save options.
    ExportDropDownFormFieldAsText *bool `json:"ExportDropDownFormFieldAsText,omitempty"`

    // Container class for html save options.
    ExportFontResources *bool `json:"ExportFontResources,omitempty"`

    // Container class for html save options.
    ExportFontsAsBase64 *bool `json:"ExportFontsAsBase64,omitempty"`

    // Container class for html save options.
    ExportHeadersFootersMode *string `json:"ExportHeadersFootersMode,omitempty"`

    // Container class for html save options.
    ExportImagesAsBase64 *bool `json:"ExportImagesAsBase64,omitempty"`

    // Container class for html save options.
    ExportLanguageInformation *bool `json:"ExportLanguageInformation,omitempty"`

    // Container class for html save options.
    ExportListLabels *string `json:"ExportListLabels,omitempty"`

    // Container class for html save options.
    ExportOriginalUrlForLinkedImages *bool `json:"ExportOriginalUrlForLinkedImages,omitempty"`

    // Container class for html save options.
    ExportPageMargins *bool `json:"ExportPageMargins,omitempty"`

    // Container class for html save options.
    ExportPageSetup *bool `json:"ExportPageSetup,omitempty"`

    // Container class for html save options.
    ExportRelativeFontSize *bool `json:"ExportRelativeFontSize,omitempty"`

    // Container class for html save options.
    ExportRoundtripInformation *bool `json:"ExportRoundtripInformation,omitempty"`

    // Container class for html save options.
    ExportTextInputFormFieldAsText *bool `json:"ExportTextInputFormFieldAsText,omitempty"`

    // Container class for html save options.
    ExportTocPageNumbers *bool `json:"ExportTocPageNumbers,omitempty"`

    // Container class for html save options.
    ExportXhtmlTransitional *bool `json:"ExportXhtmlTransitional,omitempty"`

    // Container class for html save options.
    FontResourcesSubsettingSizeThreshold *int32 `json:"FontResourcesSubsettingSizeThreshold,omitempty"`

    // Container class for html save options.
    FontsFolder *string `json:"FontsFolder,omitempty"`

    // Container class for html save options.
    FontsFolderAlias *string `json:"FontsFolderAlias,omitempty"`

    // Container class for html save options.
    HtmlVersion *string `json:"HtmlVersion,omitempty"`

    // Container class for html save options.
    ImageResolution *int32 `json:"ImageResolution,omitempty"`

    // Container class for html save options.
    ImagesFolder *string `json:"ImagesFolder,omitempty"`

    // Container class for html save options.
    ImagesFolderAlias *string `json:"ImagesFolderAlias,omitempty"`

    // Container class for html save options.
    MetafileFormat *string `json:"MetafileFormat,omitempty"`

    // Container class for html save options.
    OfficeMathOutputMode *string `json:"OfficeMathOutputMode,omitempty"`

    // Container class for html save options.
    PrettyFormat *bool `json:"PrettyFormat,omitempty"`

    // Container class for html save options.
    ResolveFontNames *bool `json:"ResolveFontNames,omitempty"`

    // Container class for html save options.
    ResourceFolder *string `json:"ResourceFolder,omitempty"`

    // Container class for html save options.
    ResourceFolderAlias *string `json:"ResourceFolderAlias,omitempty"`

    // Container class for html save options.
    SaveFormat *string `json:"SaveFormat,omitempty"`

    // Container class for html save options.
    ScaleImageToShapeSize *bool `json:"ScaleImageToShapeSize,omitempty"`

    // Container class for html save options.
    TableWidthOutputMode *string `json:"TableWidthOutputMode,omitempty"`
}

type IHtmlSaveOptionsData interface {
    IsHtmlSaveOptionsData() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
}

func (HtmlSaveOptionsData) IsHtmlSaveOptionsData() bool {
    return true
}

func (HtmlSaveOptionsData) IsSaveOptionsData() bool {
    return true
}

func (obj *HtmlSaveOptionsData) Initialize() {
    var _SaveFormat = "html"
    obj.SaveFormat = &_SaveFormat


}

func (obj *HtmlSaveOptionsData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}


