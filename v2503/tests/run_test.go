/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="run_test.go">
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

// Example of how to work with runs.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2503/api/models"
)

// Test for updating run.
func Test_Run_UpdateRun(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Runs"
    localFile := "DocumentElements/Runs/Run.doc"
    remoteFileName := "TestUpdateRun.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestRun := models.RunUpdate{
        Text: ToStringPointer("run with text"),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.UpdateRunRequest{
        Name: ToStringPointer(remoteFileName),
        ParagraphPath: ToStringPointer("paragraphs/1"),
        Index: ToInt32Pointer(int32(0)),
        Run: &requestRun,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.UpdateRun(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetRun(), "Validate UpdateRun response.");
    assert.Equal(t, "run with text", DereferenceValue(actual.GetRun().GetText()), "Validate UpdateRun response.");
}

// Test for updating run online.
func Test_Run_UpdateRunOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Runs/Run.doc"

    requestDocument := OpenFile(t, localFile)
    requestRun := models.RunUpdate{
        Text: ToStringPointer("run with text"),
    }

    options := map[string]interface{}{
    }

    request := &models.UpdateRunOnlineRequest{
        Document: requestDocument,
        ParagraphPath: ToStringPointer("paragraphs/1"),
        Run: &requestRun,
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.UpdateRunOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for adding run.
func Test_Run_InsertRun(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Runs"
    localFile := "DocumentElements/Runs/Run.doc"
    remoteFileName := "TestInsertRun.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestRun := models.RunInsert{
        Text: ToStringPointer("run with text"),
    }

    options := map[string]interface{}{
        "paragraphPath": "paragraphs/1",
        "folder": remoteDataFolder,
    }

    request := &models.InsertRunRequest{
        Name: ToStringPointer(remoteFileName),
        Run: &requestRun,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.InsertRun(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetRun(), "Validate InsertRun response.");
    assert.Equal(t, "run with text", DereferenceValue(actual.GetRun().GetText()), "Validate InsertRun response.");
    assert.Equal(t, "0.0.1.3", DereferenceValue(actual.GetRun().GetNodeId()), "Validate InsertRun response.");
}

// Test for adding run online.
func Test_Run_InsertRunOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Runs/Run.doc"

    requestDocument := OpenFile(t, localFile)
    requestRun := models.RunInsert{
        Text: ToStringPointer("run with text"),
    }

    options := map[string]interface{}{
        "paragraphPath": "paragraphs/1",
    }

    request := &models.InsertRunOnlineRequest{
        Document: requestDocument,
        Run: &requestRun,
        Optionals: options,
    }

    _, _, err := client.WordsApi.InsertRunOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for deleting run.
func Test_Run_DeleteRun(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Runs"
    localFile := "DocumentElements/Runs/Run.doc"
    remoteFileName := "TestDeleteRun.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.DeleteRunRequest{
        Name: ToStringPointer(remoteFileName),
        ParagraphPath: ToStringPointer("paragraphs/1"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, err := client.WordsApi.DeleteRun(ctx, request)

    if err != nil {
        t.Error(err)
    }

}

// Test for deleting run online.
func Test_Run_DeleteRunOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Runs/Run.doc"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.DeleteRunOnlineRequest{
        Document: requestDocument,
        ParagraphPath: ToStringPointer("paragraphs/1"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.DeleteRunOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}
