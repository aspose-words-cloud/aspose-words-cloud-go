/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="save_as_tiff_request.go">
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

import (
    "fmt"
    "io/ioutil"
    "net/url"
    "strings"
    "io"
    "encoding/json"
)

// SaveAsTiffRequest contains request data for WordsApiService.SaveAsTiff method.
type SaveAsTiffRequest struct {
        // The filename of the input document.
        Name *string
        // Tiff save options.
        SaveOptions ITiffSaveOptionsData
    /* optional (nil or map[string]interface{}) with one or more of key / value pairs:
        key: "folder" value: (*string) Original document folder.
        key: "storage" value: (*string) Original document storage.
        key: "loadEncoding" value: (*string) Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
        key: "password" value: (*string) Password of protected Word document. Use the parameter to pass a password via SDK. SDK encrypts it automatically. We don't recommend to use the parameter to pass a plain password for direct call of API.
        key: "encryptedPassword" value: (*string) Password of protected Word document. Use the parameter to pass an encrypted password for direct calls of API. See SDK code for encyption details.
        key: "useAntiAliasing" value: (*bool) The flag indicating whether to use antialiasing.
        key: "useHighQualityRendering" value: (*bool) The flag indicating whether to use high quality.
        key: "imageBrightness" value: (*float64) The level of brightness for the generated images.
        key: "imageColorMode" value: (*string) The color mode for the generated images.
        key: "imageContrast" value: (*float64) The contrast for the generated images.
        key: "numeralFormat" value: (*string) The images numeral format.
        key: "pageCount" value: (*int32) The number of pages to render.
        key: "pageIndex" value: (*int32) The index of the page to start rendering.
        key: "paperColor" value: (*string) The background image color.
        key: "pixelFormat" value: (*string) The pixel format of the generated images.
        key: "resolution" value: (*float64) The resolution of the generated images.
        key: "scale" value: (*float64) The zoom factor for the generated images.
        key: "tiffCompression" value: (*string) The compression tipe.
        key: "dmlRenderingMode" value: (*string) The optional dml rendering mode. The default value is Fallback.
        key: "dmlEffectsRenderingMode" value: (*string) The optional dml effects rendering mode. The default value is Simplified.
        key: "tiffBinarizationMethod" value: (*string) The optional TIFF binarization method. Possible values are: FloydSteinbergDithering, Threshold.
        key: "zipOutput" value: (*bool) The flag indicating whether to ZIP the output.
        key: "fontsLocation" value: (*string) Folder in filestorage with custom fonts. */
    Optionals map[string]interface{}
}


func (data *SaveAsTiffRequest) CreateRequestData() (RequestData, error) {

    var result RequestData
    var filesContentData = make([]FileReference, 0)

    result.Method = strings.ToUpper("put")

    // create path and map variables
    result.Path = "/words/{name}/saveAs/tiff"
    result.Path = strings.Replace(result.Path, "{"+"name"+"}", fmt.Sprintf("%v", *data.Name), -1)

    result.Path = strings.Replace(result.Path, "/<nil>", "", -1)
    result.Path = strings.Replace(result.Path, "//", "/", -1)

    result.HeaderParams = make(map[string]string)
    result.QueryParams = url.Values{}
    result.FormParams = make([]FormParamContainer, 0)

    if (data.SaveOptions != nil) {
        data.SaveOptions.Initialize()
    }

    if err := typeCheckParameter(data.Optionals["folder"], "string", "data.Optionals[folder]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["storage"], "string", "data.Optionals[storage]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["loadEncoding"], "string", "data.Optionals[loadEncoding]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["password"], "string", "data.Optionals[password]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["encryptedPassword"], "string", "data.Optionals[encryptedPassword]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["useAntiAliasing"], "bool", "data.Optionals[useAntiAliasing]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["useHighQualityRendering"], "bool", "data.Optionals[useHighQualityRendering]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["imageBrightness"], "float64", "data.Optionals[imageBrightness]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["imageColorMode"], "string", "data.Optionals[imageColorMode]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["imageContrast"], "float64", "data.Optionals[imageContrast]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["numeralFormat"], "string", "data.Optionals[numeralFormat]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["pageCount"], "int32", "data.Optionals[pageCount]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["pageIndex"], "int32", "data.Optionals[pageIndex]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["paperColor"], "string", "data.Optionals[paperColor]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["pixelFormat"], "string", "data.Optionals[pixelFormat]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["resolution"], "float64", "data.Optionals[resolution]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["scale"], "float64", "data.Optionals[scale]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["tiffCompression"], "string", "data.Optionals[tiffCompression]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["dmlRenderingMode"], "string", "data.Optionals[dmlRenderingMode]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["dmlEffectsRenderingMode"], "string", "data.Optionals[dmlEffectsRenderingMode]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["tiffBinarizationMethod"], "string", "data.Optionals[tiffBinarizationMethod]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["zipOutput"], "bool", "data.Optionals[zipOutput]"); err != nil {
        return result, err
    }
    if err := typeCheckParameter(data.Optionals["fontsLocation"], "string", "data.Optionals[fontsLocation]"); err != nil {
        return result, err
    }


    if localVarTempParam, localVarOk := data.Optionals["folder"].(string); localVarOk {
        result.QueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["storage"].(string); localVarOk {
        result.QueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["loadEncoding"].(string); localVarOk {
        result.QueryParams.Add("LoadEncoding", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["password"].(string); localVarOk {
        result.QueryParams.Add("Password", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["encryptedPassword"].(string); localVarOk {
        result.QueryParams.Add("EncryptedPassword", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["useAntiAliasing"].(bool); localVarOk {
        result.QueryParams.Add("UseAntiAliasing", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["useHighQualityRendering"].(bool); localVarOk {
        result.QueryParams.Add("UseHighQualityRendering", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["imageBrightness"].(float64); localVarOk {
        result.QueryParams.Add("ImageBrightness", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["imageColorMode"].(string); localVarOk {
        result.QueryParams.Add("ImageColorMode", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["imageContrast"].(float64); localVarOk {
        result.QueryParams.Add("ImageContrast", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["numeralFormat"].(string); localVarOk {
        result.QueryParams.Add("NumeralFormat", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["pageCount"].(int32); localVarOk {
        result.QueryParams.Add("PageCount", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["pageIndex"].(int32); localVarOk {
        result.QueryParams.Add("PageIndex", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["paperColor"].(string); localVarOk {
        result.QueryParams.Add("PaperColor", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["pixelFormat"].(string); localVarOk {
        result.QueryParams.Add("PixelFormat", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["resolution"].(float64); localVarOk {
        result.QueryParams.Add("Resolution", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["scale"].(float64); localVarOk {
        result.QueryParams.Add("Scale", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["tiffCompression"].(string); localVarOk {
        result.QueryParams.Add("TiffCompression", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["dmlRenderingMode"].(string); localVarOk {
        result.QueryParams.Add("DmlRenderingMode", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["dmlEffectsRenderingMode"].(string); localVarOk {
        result.QueryParams.Add("DmlEffectsRenderingMode", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["tiffBinarizationMethod"].(string); localVarOk {
        result.QueryParams.Add("TiffBinarizationMethod", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["zipOutput"].(bool); localVarOk {
        result.QueryParams.Add("ZipOutput", parameterToString(localVarTempParam, ""))
    }


    if localVarTempParam, localVarOk := data.Optionals["fontsLocation"].(string); localVarOk {
        result.QueryParams.Add("FontsLocation", parameterToString(localVarTempParam, ""))
    }



    result.FormParams = append(result.FormParams, NewJsonFormParamContainer("SaveOptions", parameterToString(data.SaveOptions, "")))


    for _, fileContentData := range filesContentData {
        fbs, _ := ioutil.ReadAll(fileContentData.Content)
        result.FormParams = append(result.FormParams, NewFileFormParamContainer(fileContentData.Reference, fbs))
    }

    return result, nil
}

func (data *SaveAsTiffRequest) CreateResponse(reader io.Reader, boundary string) (response interface{}, err error) {
            var successPayload SaveResponse
            if err = json.NewDecoder(reader).Decode(&successPayload); err != nil {
                return nil, err
            }

            return successPayload, err
}