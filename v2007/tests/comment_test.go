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

	"github.com/aspose-words-cloud/aspose-words-cloud-go/v2006/api/models"
)

func TestDeleteComment(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Comments")
	remoteName := "TestDeleteComment.docx"
	commentIndex := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteComment(ctx, remoteName, int32(commentIndex), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetComment(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Comments")
	remoteName := "TestGetComment.docx"
	commentIndex := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetComment(ctx, remoteName, int32(commentIndex), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetComments(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Comments")
	remoteName := "TestGetComments.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetComments(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertComment(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Comments")
	remoteName := "TestInsertComment.docx"
	position := models.DocumentPosition{
		Node: &models.NodeLink{
			NodeId: "0.3.0.3",
		},
		Offset: 0,
	}
	comment := models.CommentInsert{
		Author:     "Imran Anwar",
		Initial:    "IA",
		RangeStart: &position,
		RangeEnd:   &position,
		Text:       "A new comment",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertComment(ctx, remoteName, comment, options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateComment(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Comments")
	remoteName := "TestUpdateComment.docx"
	commentIndex := 0
	documentPosition := models.DocumentPosition{
		Node: &models.NodeLink{
			NodeId: "0.3.0",
		},
		Offset: 0,
	}
	comment := models.CommentUpdate{
		RangeStart: &documentPosition,
		RangeEnd:   &documentPosition,
		Initial:    "IA",
		Author:     "Imran Anwar",
		Text:       "A new Comment",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateComment(ctx, remoteName, int32(commentIndex), comment, options)

	if err != nil {
		t.Error(err)
	}
}
