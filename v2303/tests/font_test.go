/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="font_test.go">
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

// Example of how to work with font.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2303/api/models"
)

// Test for reseting cache.
func Test_Font_ResetCache(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)

    request := &models.ResetCacheRequest{
    }

    _, err := client.WordsApi.ResetCache(ctx, request)

    if err != nil {
        t.Error(err)
    }

}

// Test for GetAvailableFonts resource.
func Test_Font_GetAvailableFonts(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)

    options := map[string]interface{}{
    }

    request := &models.GetAvailableFontsRequest{
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetAvailableFonts(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.SystemFonts, "Validate GetAvailableFonts response.");
}
