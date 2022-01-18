/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="comment_test.go">
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

// Example of how to get comments from document.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/dev/api/models"
)

// Test for getting comment by specified comment's index.
func Test_Comment_GetComment(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/Comments"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetComment.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.GetCommentRequest{
        Name: ToStringPointer(remoteFileName),
        CommentIndex: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetComment(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Comment, "Validate GetComment response.");
    assert.Equal(t, "Comment 1" + "\r\n\r\n", actual.Comment.Text, "Validate GetComment response.");
}

// Test for getting comment by specified comment's index online.
func Test_Comment_GetCommentOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.GetCommentOnlineRequest{
        Document: requestDocument,
        CommentIndex: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetCommentOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for getting all comments from document.
func Test_Comment_GetComments(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/Comments"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetComments.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.GetCommentsRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetComments(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Comments, "Validate GetComments response.");
    assert.NotNil(t, actual.Comments.CommentList, "Validate GetComments response.");
    assert.Equal(t, 1, len(actual.Comments.CommentList), "Validate GetComments response.");
    assert.Equal(t, "Comment 1" + "\r\n\r\n", actual.Comments.CommentList[0].Text, "Validate GetComments response.");
}

// Test for getting all comments from document online.
func Test_Comment_GetCommentsOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.GetCommentsOnlineRequest{
        Document: requestDocument,
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetCommentsOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for adding comment.
func Test_Comment_InsertComment(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/Comments"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestInsertComment.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestCommentRangeStartNode := models.NodeLink{
        NodeId: ToStringPointer("0.3.0.3"),
    }
    requestCommentRangeStart := models.DocumentPosition{
        Node: requestCommentRangeStartNode,
        Offset: ToInt32Pointer(int32(0)),
    }
    requestCommentRangeEndNode := models.NodeLink{
        NodeId: ToStringPointer("0.3.0.3"),
    }
    requestCommentRangeEnd := models.DocumentPosition{
        Node: requestCommentRangeEndNode,
        Offset: ToInt32Pointer(int32(0)),
    }
    requestComment := models.CommentInsert{
        RangeStart: requestCommentRangeStart,
        RangeEnd: requestCommentRangeEnd,
        Initial: ToStringPointer("IA"),
        Author: ToStringPointer("Imran Anwar"),
        Text: ToStringPointer("A new Comment"),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.InsertCommentRequest{
        Name: ToStringPointer(remoteFileName),
        Comment: requestComment,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.InsertComment(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Comment, "Validate InsertComment response.");
    assert.Equal(t, "A new Comment" + "\r\n", actual.Comment.Text, "Validate InsertComment response.");
    assert.NotNil(t, actual.Comment.RangeStart, "Validate InsertComment response.");
    assert.NotNil(t, actual.Comment.RangeStart.Node, "Validate InsertComment response.");
    assert.Equal(t, "0.3.0.4", actual.Comment.RangeStart.Node.NodeId, "Validate InsertComment response.");
}

// Test for adding comment online.
func Test_Comment_InsertCommentOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)
    requestCommentRangeStartNode := models.NodeLink{
        NodeId: ToStringPointer("0.3.0.3"),
    }
    requestCommentRangeStart := models.DocumentPosition{
        Node: requestCommentRangeStartNode,
        Offset: ToInt32Pointer(int32(0)),
    }
    requestCommentRangeEndNode := models.NodeLink{
        NodeId: ToStringPointer("0.3.0.3"),
    }
    requestCommentRangeEnd := models.DocumentPosition{
        Node: requestCommentRangeEndNode,
        Offset: ToInt32Pointer(int32(0)),
    }
    requestComment := models.CommentInsert{
        RangeStart: requestCommentRangeStart,
        RangeEnd: requestCommentRangeEnd,
        Initial: ToStringPointer("IA"),
        Author: ToStringPointer("Imran Anwar"),
        Text: ToStringPointer("A new Comment"),
    }

    options := map[string]interface{}{
    }

    request := &models.InsertCommentOnlineRequest{
        Document: requestDocument,
        Comment: requestComment,
        Optionals: options,
    }

    _, _, err := client.WordsApi.InsertCommentOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for updating comment.
func Test_Comment_UpdateComment(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/Comments"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestUpdateComment.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestCommentRangeStartNode := models.NodeLink{
        NodeId: ToStringPointer("0.3.0"),
    }
    requestCommentRangeStart := models.DocumentPosition{
        Node: requestCommentRangeStartNode,
        Offset: ToInt32Pointer(int32(0)),
    }
    requestCommentRangeEndNode := models.NodeLink{
        NodeId: ToStringPointer("0.3.0"),
    }
    requestCommentRangeEnd := models.DocumentPosition{
        Node: requestCommentRangeEndNode,
        Offset: ToInt32Pointer(int32(0)),
    }
    requestComment := models.CommentUpdate{
        RangeStart: requestCommentRangeStart,
        RangeEnd: requestCommentRangeEnd,
        Initial: ToStringPointer("IA"),
        Author: ToStringPointer("Imran Anwar"),
        Text: ToStringPointer("A new Comment"),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.UpdateCommentRequest{
        Name: ToStringPointer(remoteFileName),
        CommentIndex: ToInt32Pointer(int32(0)),
        Comment: requestComment,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.UpdateComment(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Comment, "Validate UpdateComment response.");
    assert.Equal(t, "A new Comment" + "\r\n", actual.Comment.Text, "Validate UpdateComment response.");
    assert.NotNil(t, actual.Comment.RangeStart, "Validate UpdateComment response.");
    assert.NotNil(t, actual.Comment.RangeStart.Node, "Validate UpdateComment response.");
    assert.Equal(t, "0.3.0.1", actual.Comment.RangeStart.Node.NodeId, "Validate UpdateComment response.");
}

// Test for updating comment online.
func Test_Comment_UpdateCommentOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)
    requestCommentRangeStartNode := models.NodeLink{
        NodeId: ToStringPointer("0.3.0"),
    }
    requestCommentRangeStart := models.DocumentPosition{
        Node: requestCommentRangeStartNode,
        Offset: ToInt32Pointer(int32(0)),
    }
    requestCommentRangeEndNode := models.NodeLink{
        NodeId: ToStringPointer("0.3.0"),
    }
    requestCommentRangeEnd := models.DocumentPosition{
        Node: requestCommentRangeEndNode,
        Offset: ToInt32Pointer(int32(0)),
    }
    requestComment := models.CommentUpdate{
        RangeStart: requestCommentRangeStart,
        RangeEnd: requestCommentRangeEnd,
        Initial: ToStringPointer("IA"),
        Author: ToStringPointer("Imran Anwar"),
        Text: ToStringPointer("A new Comment"),
    }

    options := map[string]interface{}{
    }

    request := &models.UpdateCommentOnlineRequest{
        Document: requestDocument,
        CommentIndex: ToInt32Pointer(int32(0)),
        Comment: requestComment,
        Optionals: options,
    }

    _, _, err := client.WordsApi.UpdateCommentOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// A test for DeleteComment.
func Test_Comment_DeleteComment(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/Comments"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestDeleteComment.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }

    request := &models.DeleteCommentRequest{
        Name: ToStringPointer(remoteFileName),
        CommentIndex: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, err := client.WordsApi.DeleteComment(ctx, request)

    if err != nil {
        t.Error(err)
    }

}

// A test for DeleteComment online.
func Test_Comment_DeleteCommentOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.DeleteCommentOnlineRequest{
        Document: requestDocument,
        CommentIndex: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, err := client.WordsApi.DeleteCommentOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// A test for DeleteComments.
func Test_Comment_DeleteComments(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/Comments"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestDeleteComment.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }

    request := &models.DeleteCommentsRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: options,
    }

    _, err := client.WordsApi.DeleteComments(ctx, request)

    if err != nil {
        t.Error(err)
    }

}

// A test for DeleteComments online.
func Test_Comment_DeleteCommentsOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.DeleteCommentsOnlineRequest{
        Document: requestDocument,
        Optionals: options,
    }

    _, err := client.WordsApi.DeleteCommentsOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}
