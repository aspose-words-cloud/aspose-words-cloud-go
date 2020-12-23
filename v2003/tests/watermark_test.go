//
// MIT License

// Copyright (c) 2019 Aspose Pty Ltd

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

package api_test

import (
	"path"
	"testing"

	"github.com/aspose-words-cloud/aspose-words-cloud-go/v2003/api/models"
)

func TestDeleteWatermark(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "Watermark")
	remoteName := "TestDeleteWatermark.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.DeleteWatermark(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertWatermarkImage(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "Watermark")
	remoteName := "TestInsertWatermarkImage.docx"
	localImagePath := GetLocalPath("Common", "aspose-cloud.png")
	remoteImageName := "TestInsertWatermarkImage.png"

	options := map[string]interface{}{
		"folder":        remoteFolder,
		"image":         path.Join(remoteFolder, remoteImageName),
		"rotationAngle": 0.0,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))
	UploadNextFileToStorage(t, ctx, client, localImagePath, path.Join(remoteFolder, remoteImageName))

	_, _, err := client.WordsApi.InsertWatermarkImage(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertWatermarkText(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "Watermark")
	remoteName := "TestInsertWatermarkText.docx"
	watermarkText := models.WatermarkText{
		Text:          "This is the text",
		RotationAngle: 90.0,
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertWatermarkText(ctx, remoteName, watermarkText, options)

	if err != nil {
		t.Error(err)
	}
}
