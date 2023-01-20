/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="lists_test.go">
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

// Example of how to work with lists.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2301/api/models"
)

// Test for getting lists from document.
func Test_Lists_GetLists(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Lists"
    localFile := "DocumentElements/Lists/ListsGet.doc"
    remoteFileName := "TestGetLists.doc"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.GetListsRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetLists(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Lists, "Validate GetLists response.");
    assert.NotNil(t, actual.Lists.ListInfo, "Validate GetLists response.");
    assert.Equal(t, 2, len(actual.Lists.ListInfo), "Validate GetLists response.");
    assert.Equal(t, int32(1), actual.Lists.ListInfo[0].ListId, "Validate GetLists response.");
}

// Test for getting lists from document online.
func Test_Lists_GetListsOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Lists/ListsGet.doc"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.GetListsOnlineRequest{
        Document: requestDocument,
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetListsOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for getting list from document.
func Test_Lists_GetList(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Lists"
    localFile := "DocumentElements/Lists/ListsGet.doc"
    remoteFileName := "TestGetList.doc"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.GetListRequest{
        Name: ToStringPointer(remoteFileName),
        ListId: ToInt32Pointer(int32(1)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetList(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.List, "Validate GetList response.");
    assert.Equal(t, int32(1), actual.List.ListId, "Validate GetList response.");
}

// Test for getting list from document online.
func Test_Lists_GetListOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Lists/ListsGet.doc"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.GetListOnlineRequest{
        Document: requestDocument,
        ListId: ToInt32Pointer(int32(1)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetListOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for updating list from document.
func Test_Lists_UpdateList(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Lists"
    localFile := "DocumentElements/Lists/ListsGet.doc"
    remoteFileName := "TestUpdateList.doc"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestListUpdate := models.ListUpdate{
        IsRestartAtEachSection: ToBoolPointer(true),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.UpdateListRequest{
        Name: ToStringPointer(remoteFileName),
        ListId: ToInt32Pointer(int32(1)),
        ListUpdate: &requestListUpdate,
        Optionals: options,
    }

    _, _, err := client.WordsApi.UpdateList(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for updating list from document online.
func Test_Lists_UpdateListOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Lists/ListsGet.doc"

    requestDocument := OpenFile(t, localFile)
    requestListUpdate := models.ListUpdate{
        IsRestartAtEachSection: ToBoolPointer(true),
    }

    options := map[string]interface{}{
    }

    request := &models.UpdateListOnlineRequest{
        Document: requestDocument,
        ListId: ToInt32Pointer(int32(1)),
        ListUpdate: &requestListUpdate,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.UpdateListOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Model.List, "Validate UpdateListOnline response.");
    assert.Equal(t, int32(1), actual.Model.List.ListId, "Validate UpdateListOnline response.");
    assert.True(t, actual.Model.List.IsRestartAtEachSection, "Validate UpdateListOnline response.");
}

// Test for updating list level from document.
func Test_Lists_UpdateListLevel(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Lists"
    localFile := "DocumentElements/Lists/ListsGet.doc"
    remoteFileName := "TestUpdateListLevel.doc"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestListUpdate := models.ListLevelUpdate{
        Alignment: ToStringPointer("Right"),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.UpdateListLevelRequest{
        Name: ToStringPointer(remoteFileName),
        ListId: ToInt32Pointer(int32(1)),
        ListLevel: ToInt32Pointer(int32(1)),
        ListUpdate: &requestListUpdate,
        Optionals: options,
    }

    _, _, err := client.WordsApi.UpdateListLevel(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for updating list level from document online.
func Test_Lists_UpdateListLevelOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Lists/ListsGet.doc"

    requestDocument := OpenFile(t, localFile)
    requestListUpdate := models.ListLevelUpdate{
        Alignment: ToStringPointer("Right"),
    }

    options := map[string]interface{}{
    }

    request := &models.UpdateListLevelOnlineRequest{
        Document: requestDocument,
        ListId: ToInt32Pointer(int32(1)),
        ListUpdate: &requestListUpdate,
        ListLevel: ToInt32Pointer(int32(1)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.UpdateListLevelOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.Model.List, "Validate UpdateListLevelOnline response.");
    assert.NotNil(t, actual.Model.List.ListLevels, "Validate UpdateListLevelOnline response.");
    assert.NotNil(t, actual.Model.List.ListLevels.ListLevel, "Validate UpdateListLevelOnline response.");
    assert.Equal(t, 9, len(actual.Model.List.ListLevels.ListLevel), "Validate UpdateListLevelOnline response.");

}

// Test for inserting list from document.
func Test_Lists_InsertList(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Lists"
    localFile := "DocumentElements/Lists/ListsGet.doc"
    remoteFileName := "TestInsertList.doc"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestListInsert := models.ListInsert{
        Template: ToStringPointer("OutlineLegal"),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.InsertListRequest{
        Name: ToStringPointer(remoteFileName),
        ListInsert: &requestListInsert,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.InsertList(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.List, "Validate InsertList response.");
    assert.Equal(t, int32(3), actual.List.ListId, "Validate InsertList response.");
}

// Test for inserting list from document online.
func Test_Lists_InsertListOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Lists/ListsGet.doc"

    requestDocument := OpenFile(t, localFile)
    requestListInsert := models.ListInsert{
        Template: ToStringPointer("OutlineLegal"),
    }

    options := map[string]interface{}{
    }

    request := &models.InsertListOnlineRequest{
        Document: requestDocument,
        ListInsert: &requestListInsert,
        Optionals: options,
    }

    _, _, err := client.WordsApi.InsertListOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}
