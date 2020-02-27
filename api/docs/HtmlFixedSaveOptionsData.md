# HtmlFixedSaveOptionsData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SaveFormat** | **string** | Gets or sets format of save. | [optional] [default to null]
**FileName** | **string** | Gets or sets name of destination file. | [optional] [default to null]
**DmlRenderingMode** | **string** | Gets or sets a value determining how DrawingML shapes are rendered. { Fallback | DrawingML }. | [optional] [default to null]
**DmlEffectsRenderingMode** | **string** | Gets or sets a value determining how DrawingML effects are rendered. { Simplified | None | Fine }. | [optional] [default to null]
**ZipOutput** | **bool** | Gets or sets controls zip output or not. Default value is false. | [optional] [default to null]
**UpdateLastSavedTimeProperty** | **bool** | Gets or sets a value determining whether the Aspose.Words.Properties.BuiltInDocumentProperties.LastSavedTime property is updated before saving. | [optional] [default to null]
**UpdateSdtContent** | **bool** | Gets or sets value determining whether content of  is updated before saving. | [optional] [default to null]
**UpdateFields** | **bool** | Gets or sets a value determining if fields should be updated before saving the document to a fixed page format. Default value for this property is. true | [optional] [default to null]
**ColorMode** | **string** | Gets or sets a value determining how colors are rendered. { Normal | Grayscale}. | [optional] [default to null]
**JpegQuality** | **int32** | Gets or sets determines the quality of the JPEG images inside PDF document. | [optional] [default to null]
**MetafileRenderingOptions** | [***MetafileRenderingOptionsData**](MetafileRenderingOptionsData.md) | Gets or sets allows to specify metafile rendering options. | [optional] [default to null]
**NumeralFormat** | **string** | Gets or sets indicates the symbol set that is used to represent numbers while rendering to fixed page formats. | [optional] [default to null]
**OptimizeOutput** | **bool** | Gets or sets flag indicates whether it is required to optimize output of XPS. If this flag is set redundant nested canvases and empty canvases are removed, also neighbor glyphs with the same formatting are concatenated. Note: The accuracy of the content display may be affected if this property is set to true.  Default is false. | [optional] [default to null]
**PageCount** | **int32** | Gets or sets determines number of pages to render. | [optional] [default to null]
**PageIndex** | **int32** | Gets or sets determines 0-based index of the first page to render. | [optional] [default to null]
**CssClassNamesPrefix** | **string** | Gets or sets specifies prefix which is added to all class names in style.css file. Default value is \&quot;aw\&quot;. | [optional] [default to null]
**Encoding** | **string** | Gets or sets encoding. | [optional] [default to null]
**ExportEmbeddedCss** | **bool** | Gets or sets specifies whether the CSS (Cascading Style Sheet) should be embedded into Html document. | [optional] [default to null]
**ExportEmbeddedFonts** | **bool** | Gets or sets specifies whether fonts should be embedded into Html document in Base64 format. | [optional] [default to null]
**ExportEmbeddedImages** | **bool** | Gets or sets specifies whether images should be embedded into Html document in Base64 format. | [optional] [default to null]
**ExportFormFields** | **bool** | Gets or sets indication of whether form fields are exported as interactive items (as &#39;input&#39; tag) rather than converted to text or graphics. | [optional] [default to null]
**FontFormat** | **string** | Gets or sets specifies export format of fonts. | [optional] [default to null]
**PageHorizontalAlignment** | **string** | Gets or sets specifies the horizontal alignment of pages in an HTML document. Default value is HtmlFixedHorizontalPageAlignment.Center. | [optional] [default to null]
**PageMargins** | **float64** | Gets or sets specifies the margins around pages in an HTML document. The margins value is measured in points and should be equal to or greater than 0. Default value is 10 points. | [optional] [default to null]
**ResourcesFolder** | **string** | Gets or sets specifies the physical folder where resources are saved when exporting a document. | [optional] [default to null]
**ResourcesFolderAlias** | **string** | Gets or sets specifies the name of the folder used to construct resource URIs. | [optional] [default to null]
**SaveFontFaceCssSeparately** | **bool** | Gets or sets flag indicates whether \&quot;@font-face\&quot; CSS rules should be placed into a separate file \&quot;fontFaces.css\&quot; when a document is being saved with external stylesheet (that is, when Aspose.Words.Saving.HtmlFixedSaveOptions.ExportEmbeddedCss is false). Default value is false, all CSS rules are written into single file \&quot;styles.css\&quot;. | [optional] [default to null]
**ShowPageBorder** | **bool** | Gets or sets specifies whether border around pages should be shown. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


