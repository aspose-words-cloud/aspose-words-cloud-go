# TiffSaveOptionsData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColorMode** | **string** | Gets or sets a value determining how colors are rendered. { Normal | Grayscale}. | [optional] [default to null]
**SaveFormat** | **string** | Gets or sets format of save. | [optional] [default to null]
**FileName** | **string** | Gets or sets name of destination file. | [optional] [default to null]
**DmlRenderingMode** | **string** | Gets or sets a value determining how DrawingML shapes are rendered. { Fallback | DrawingML }. | [optional] [default to null]
**DmlEffectsRenderingMode** | **string** | Gets or sets a value determining how DrawingML effects are rendered. { Simplified | None | Fine }. | [optional] [default to null]
**ZipOutput** | **bool** | Gets or sets controls zip output or not. Default value is false. | [optional] [default to null]
**UpdateLastSavedTimeProperty** | **bool** | Gets or sets a value determining whether the Aspose.Words.Properties.BuiltInDocumentProperties.LastSavedTime property is updated before saving. | [optional] [default to null]
**UpdateSdtContent** | **bool** | Gets or sets value determining whether content of  is updated before saving. | [optional] [default to null]
**UpdateFields** | **bool** | Gets or sets a value determining if fields should be updated before saving the document to a fixed page format. Default value for this property is. true | [optional] [default to null]
**JpegQuality** | **int32** | Gets or sets determines the quality of the JPEG images inside PDF document. | [optional] [default to null]
**MetafileRenderingOptions** | [***MetafileRenderingOptionsData**](MetafileRenderingOptionsData.md) | Gets or sets allows to specify metafile rendering options. | [optional] [default to null]
**NumeralFormat** | **string** | Gets or sets indicates the symbol set that is used to represent numbers while rendering to fixed page formats. | [optional] [default to null]
**OptimizeOutput** | **bool** | Gets or sets flag indicates whether it is required to optimize output of XPS. If this flag is set redundant nested canvases and empty canvases are removed, also neighbor glyphs with the same formatting are concatenated. Note: The accuracy of the content display may be affected if this property is set to true.  Default is false. | [optional] [default to null]
**PageCount** | **int32** | Gets or sets determines number of pages to render. | [optional] [default to null]
**PageIndex** | **int32** | Gets or sets determines 0-based index of the first page to render. | [optional] [default to null]
**GraphicsQualityOptions** | [***GraphicsQualityOptionsData**](GraphicsQualityOptionsData.md) | Gets or sets allows to specify additional System.Drawing.Graphics quality options. | [optional] [default to null]
**HorizontalResolution** | **float64** | Gets or sets the horizontal resolution for the generated images, in dots per inch.  This property has effect only when saving to raster image formats. The default value is 96. | [optional] [default to null]
**ImageBrightness** | **float64** | Gets or sets brightness of image. | [optional] [default to null]
**ImageColorMode** | **string** | Gets or sets color mode of image. | [optional] [default to null]
**ImageContrast** | **float64** | Gets or sets contrast of image. | [optional] [default to null]
**PaperColor** | **string** | Gets or sets background (paper) color of image. | [optional] [default to null]
**PixelFormat** | **string** | Gets or sets pixel format of image. | [optional] [default to null]
**Resolution** | **float64** | Gets or sets both horizontal and vertical resolution for the generated images, in dots per inch.  This property has effect only when saving to raster image formats. The default value is 96. | [optional] [default to null]
**Scale** | **float64** | Gets or sets zoom factor of image. | [optional] [default to null]
**UseAntiAliasing** | **bool** | Gets or sets determine whether or not to use anti-aliasing for rendering. | [optional] [default to null]
**UseGdiEmfRenderer** | **bool** | Gets or sets a value determining whether to use GDI+ or Aspose.Words metafile renderer when saving to EMF. | [optional] [default to null]
**UseHighQualityRendering** | **bool** | Gets or sets determine whether or not to use high quality (i.e. slow) rendering algorithms. | [optional] [default to null]
**VerticalResolution** | **float64** | Gets or sets the vertical resolution for the generated images, in dots per inch.  This property has effect only when saving to raster image formats. The default value is 96. | [optional] [default to null]
**ThresholdForFloydSteinbergDithering** | **int32** | Gets or sets the threshold that determines the value of the binarization error in the Floyd-Steinberg method. when ImageBinarizationMethod is ImageBinarizationMethod.FloydSteinbergDithering. Default value is 128. | [optional] [default to null]
**TiffBinarizationMethod** | **string** | Gets or sets specifies method used while converting images to 1 bpp format. | [optional] [default to null]
**TiffCompression** | **string** | Gets or sets type of compression. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


