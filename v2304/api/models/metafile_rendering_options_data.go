/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="metafile_rendering_options_data.go">
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

// Container class for options of metafile rendering.
type MetafileRenderingOptionsDataResult struct {
    // Container class for options of metafile rendering.
    EmfPlusDualRenderingMode string `json:"EmfPlusDualRenderingMode,omitempty"`

    // Container class for options of metafile rendering.
    EmulateRasterOperations bool `json:"EmulateRasterOperations,omitempty"`

    // Container class for options of metafile rendering.
    RenderingMode string `json:"RenderingMode,omitempty"`

    // Container class for options of metafile rendering.
    ScaleWmfFontsToMetafileSize bool `json:"ScaleWmfFontsToMetafileSize,omitempty"`

    // Container class for options of metafile rendering.
    UseEmfEmbeddedToWmf bool `json:"UseEmfEmbeddedToWmf,omitempty"`
}

type MetafileRenderingOptionsData struct {
    // Container class for options of metafile rendering.
    EmfPlusDualRenderingMode *string `json:"EmfPlusDualRenderingMode,omitempty"`

    // Container class for options of metafile rendering.
    EmulateRasterOperations *bool `json:"EmulateRasterOperations,omitempty"`

    // Container class for options of metafile rendering.
    RenderingMode *string `json:"RenderingMode,omitempty"`

    // Container class for options of metafile rendering.
    ScaleWmfFontsToMetafileSize *bool `json:"ScaleWmfFontsToMetafileSize,omitempty"`

    // Container class for options of metafile rendering.
    UseEmfEmbeddedToWmf *bool `json:"UseEmfEmbeddedToWmf,omitempty"`
}

type IMetafileRenderingOptionsData interface {
    IsMetafileRenderingOptionsData() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
}

func (MetafileRenderingOptionsData) IsMetafileRenderingOptionsData() bool {
    return true
}


func (obj *MetafileRenderingOptionsData) Initialize() {
}

func (obj *MetafileRenderingOptionsData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}


