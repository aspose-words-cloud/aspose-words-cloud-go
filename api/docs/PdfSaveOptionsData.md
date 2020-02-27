# PdfSaveOptionsData

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
**Compliance** | **string** | Gets or sets specifies the PDF standards compliance level for output documents. | [optional] [default to null]
**CreateNoteHyperlinks** | **bool** | Gets or sets specifies whether to convert footnote/endnote references in main text story into active hyperlinks. When clicked the hyperlink will lead to the corresponding footnote/endnote. Default is false. | [optional] [default to null]
**CustomPropertiesExport** | **string** | Gets or sets a value determining the way  are exported to PDF file. Default value is . | [optional] [default to null]
**DigitalSignatureDetails** | [***PdfDigitalSignatureDetailsData**](PdfDigitalSignatureDetailsData.md) | Gets or sets specifies the details for signing the output PDF document. | [optional] [default to null]
**DisplayDocTitle** | **bool** | Gets or sets a flag specifying whether the window�s title bar should display the document title taken from the Title entry of the document information dictionary. | [optional] [default to null]
**DownsampleOptions** | [***DownsampleOptionsData**](DownsampleOptionsData.md) | Gets or sets allows to specify downsample options. | [optional] [default to null]
**EmbedFullFonts** | **bool** | Gets or sets controls how fonts are embedded into the resulting PDF documents. | [optional] [default to null]
**EncryptionDetails** | [***PdfEncryptionDetailsData**](PdfEncryptionDetailsData.md) | Gets or sets specifies the details for encrypting the output PDF document. | [optional] [default to null]
**EscapeUri** | **bool** | Gets or sets a flag specifying whether URI should be escaped before writing.              | [optional] [default to null]
**ExportDocumentStructure** | **bool** | Gets or sets determines whether or not to export document structure. | [optional] [default to null]
**FontEmbeddingMode** | **string** | Gets or sets specifies the font embedding mode. | [optional] [default to null]
**HeaderFooterBookmarksExportMode** | **string** | Gets or sets determines how bookmarks in headers/footers are exported. The default value is Aspose.Words.Saving.HeaderFooterBookmarksExportMode.All. | [optional] [default to null]
**ImageColorSpaceExportMode** | **string** | Gets or sets specifies how the color space will be selected for the images in PDF document. | [optional] [default to null]
**ImageCompression** | **string** | Gets or sets specifies compression type to be used for all images in the document. | [optional] [default to null]
**OpenHyperlinksInNewWindow** | **bool** | Gets or sets determines whether hyperlinks in the output Pdf document are forced to be opened in a new window (or tab) of a browser. | [optional] [default to null]
**OutlineOptions** | [***OutlineOptionsData**](OutlineOptionsData.md) | Gets or sets allows to specify outline options. | [optional] [default to null]
**PageMode** | **string** | Gets or sets specifies how the PDF document should be displayed when opened in the PDF reader. | [optional] [default to null]
**PreblendImages** | **bool** | Gets or sets a value determining whether or not to preblend transparent images with black background color. | [optional] [default to null]
**PreserveFormFields** | **bool** | Gets or sets specifies whether to preserve Microsoft Word form fields as form fields in PDF or convert them to text. | [optional] [default to null]
**TextCompression** | **string** | Gets or sets specifies compression type to be used for all textual content in the document. | [optional] [default to null]
**UseBookFoldPrintingSettings** | **bool** | Gets or sets determines whether the document should be saved using a booklet printing layout. | [optional] [default to null]
**UseCoreFonts** | **bool** | Gets or sets determines whether or not to substitute TrueType fonts Arial, Times New Roman, Courier New and Symbol with core PDF Type 1 fonts. | [optional] [default to null]
**ZoomBehavior** | **string** | Gets or sets determines what type of zoom should be applied when a document is opened with a PDF viewer. | [optional] [default to null]
**ZoomFactor** | **int32** | Gets or sets determines zoom factor (in percentages) for a document. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


