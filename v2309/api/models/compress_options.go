/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="compress_options.go">
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

// Options of document compress.

type ICompressOptions interface {
    IsCompressOptions() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetImagesQuality() *int32
    SetImagesQuality(value *int32)
    GetImagesReduceSizeFactor() *int32
    SetImagesReduceSizeFactor(value *int32)
}

type CompressOptions struct {
    // Options of document compress.
    ImagesQuality *int32 `json:"ImagesQuality,omitempty"`

    // Options of document compress.
    ImagesReduceSizeFactor *int32 `json:"ImagesReduceSizeFactor,omitempty"`
}

func (CompressOptions) IsCompressOptions() bool {
    return true
}


func (obj *CompressOptions) Initialize() {
}

func (obj *CompressOptions) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["ImagesQuality"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ImagesQuality = new(int32)
            *obj.ImagesQuality = int32(parsedValue)
        }

    } else if jsonValue, exists := json["imagesQuality"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ImagesQuality = new(int32)
            *obj.ImagesQuality = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["ImagesReduceSizeFactor"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ImagesReduceSizeFactor = new(int32)
            *obj.ImagesReduceSizeFactor = int32(parsedValue)
        }

    } else if jsonValue, exists := json["imagesReduceSizeFactor"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ImagesReduceSizeFactor = new(int32)
            *obj.ImagesReduceSizeFactor = int32(parsedValue)
        }

    }
}

func (obj *CompressOptions) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *CompressOptions) GetImagesQuality() *int32 {
    return obj.ImagesQuality
}

func (obj *CompressOptions) SetImagesQuality(value *int32) {
    obj.ImagesQuality = value
}

func (obj *CompressOptions) GetImagesReduceSizeFactor() *int32 {
    return obj.ImagesReduceSizeFactor
}

func (obj *CompressOptions) SetImagesReduceSizeFactor(value *int32) {
    obj.ImagesReduceSizeFactor = value
}

