/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="load_web_document_test.go">
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

// Example of how to load web document.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2506/api/models"
)

// Test for loading web document.
func Test_LoadWebDocument_LoadWebDocument(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    requestDataSaveOptions := models.DocSaveOptionsData{
        FileName: ToStringPointer("google.doc"),
        DmlEffectsRenderingMode: ToStringPointer("None"),
        DmlRenderingMode: ToStringPointer("DrawingML"),
        ZipOutput: ToBoolPointer(false),
    }
    requestData := models.LoadWebDocumentData{
        LoadingDocumentUrl: ToStringPointer("http://google.com"),
        SaveOptions: &requestDataSaveOptions,
    }

    options := map[string]interface{}{
    }

    request := &models.LoadWebDocumentRequest{
        Data: &requestData,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.LoadWebDocument(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetSaveResult(), "Validate LoadWebDocument response.");
    assert.NotNil(t, actual.GetSaveResult().GetDestDocument(), "Validate LoadWebDocument response.");
    assert.Equal(t, "google.doc", DereferenceValue(actual.GetSaveResult().GetDestDocument().GetHref()), "Validate LoadWebDocument response.");
}
