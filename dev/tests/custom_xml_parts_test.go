/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="custom_xml_parts_test.go">
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

// Example of how to use custom xml parts in documents.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/dev/api/models"
)

// Test for getting custom xml part by specified index.
func Test_CustomXmlParts_GetCustomXmlPart(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/CustomXmlParts"
    localFile := "DocumentElements/CustomXmlParts/MultipleCustomXmlParts.docx"
    remoteFileName := "TestGetCustomXmlPart.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.GetCustomXmlPartRequest{
        Name: ToStringPointer(remoteFileName),
        CustomXmlPartIndex: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetCustomXmlPart(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetCustomXmlPart(), "Validate GetCustomXmlPart response.");
    assert.Equal(t, "aspose", DereferenceValue(actual.GetCustomXmlPart().GetId()), "Validate GetCustomXmlPart response.");
    assert.Equal(t, "<Metadata><Author>author1</Author><Initial>initial</Initial><DateTime>2015-01-22T00:00:00</DateTime><Text>text</Text></Metadata>", DereferenceValue(actual.GetCustomXmlPart().GetData()), "Validate GetCustomXmlPart response.");
}

// Test for getting custom xml part by specified index online.
func Test_CustomXmlParts_GetCustomXmlPartOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/CustomXmlParts/MultipleCustomXmlParts.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.GetCustomXmlPartOnlineRequest{
        Document: requestDocument,
        CustomXmlPartIndex: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetCustomXmlPartOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetCustomXmlPart(), "Validate GetCustomXmlPartOnline response.");
    assert.Equal(t, "aspose", DereferenceValue(actual.GetCustomXmlPart().GetId()), "Validate GetCustomXmlPartOnline response.");
    assert.Equal(t, "<Metadata><Author>author1</Author><Initial>initial</Initial><DateTime>2015-01-22T00:00:00</DateTime><Text>text</Text></Metadata>", DereferenceValue(actual.GetCustomXmlPart().GetData()), "Validate GetCustomXmlPartOnline response.");
}

// Test for getting all custom xml parts from document.
func Test_CustomXmlParts_GetCustomXmlParts(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/CustomXmlParts"
    localFile := "DocumentElements/CustomXmlParts/MultipleCustomXmlParts.docx"
    remoteFileName := "TestGetCustomXmlParts.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.GetCustomXmlPartsRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetCustomXmlParts(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetCustomXmlParts(), "Validate GetCustomXmlParts response.");
    assert.NotNil(t, actual.GetCustomXmlParts().GetCustomXmlPartsList(), "Validate GetCustomXmlParts response.");
    assert.Equal(t, 2, len(actual.GetCustomXmlParts().GetCustomXmlPartsList()), "Validate GetCustomXmlParts response.");
    assert.Equal(t, "aspose", DereferenceValue(actual.GetCustomXmlParts().GetCustomXmlPartsList()[0].GetId()), "Validate GetCustomXmlParts response.");
    assert.Equal(t, "<Metadata><Author>author1</Author><Initial>initial</Initial><DateTime>2015-01-22T00:00:00</DateTime><Text>text</Text></Metadata>", DereferenceValue(actual.GetCustomXmlParts().GetCustomXmlPartsList()[0].GetData()), "Validate GetCustomXmlParts response.");
}

// Test for getting all custom xml parts from document online.
func Test_CustomXmlParts_GetCustomXmlPartsOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/CustomXmlParts/MultipleCustomXmlParts.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.GetCustomXmlPartsOnlineRequest{
        Document: requestDocument,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetCustomXmlPartsOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetCustomXmlParts(), "Validate GetCustomXmlPartsOnline response.");
    assert.NotNil(t, actual.GetCustomXmlParts().GetCustomXmlPartsList(), "Validate GetCustomXmlPartsOnline response.");
    assert.Equal(t, 2, len(actual.GetCustomXmlParts().GetCustomXmlPartsList()), "Validate GetCustomXmlPartsOnline response.");
    assert.Equal(t, "aspose", DereferenceValue(actual.GetCustomXmlParts().GetCustomXmlPartsList()[0].GetId()), "Validate GetCustomXmlPartsOnline response.");
    assert.Equal(t, "<Metadata><Author>author1</Author><Initial>initial</Initial><DateTime>2015-01-22T00:00:00</DateTime><Text>text</Text></Metadata>", DereferenceValue(actual.GetCustomXmlParts().GetCustomXmlPartsList()[0].GetData()), "Validate GetCustomXmlPartsOnline response.");
}

// Test for adding custom xml part.
func Test_CustomXmlParts_InsertCustomXmlPart(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/CustomXmlParts"
    localFile := "DocumentElements/CustomXmlParts/MultipleCustomXmlParts.docx"
    remoteFileName := "TestInsertCustomXmlPart.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestCustomXmlPart := models.CustomXmlPartInsert{
        Id: ToStringPointer("hello"),
        Data: ToStringPointer("<data>Hello world</data>"),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.InsertCustomXmlPartRequest{
        Name: ToStringPointer(remoteFileName),
        CustomXmlPart: &requestCustomXmlPart,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.InsertCustomXmlPart(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetCustomXmlPart(), "Validate InsertCustomXmlPart response.");
    assert.Equal(t, "hello", DereferenceValue(actual.GetCustomXmlPart().GetId()), "Validate InsertCustomXmlPart response.");
    assert.Equal(t, "<data>Hello world</data>", DereferenceValue(actual.GetCustomXmlPart().GetData()), "Validate InsertCustomXmlPart response.");
}

// Test for adding custom xml part online.
func Test_CustomXmlParts_InsertCustomXmlPartOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/CustomXmlParts/MultipleCustomXmlParts.docx"

    requestDocument := OpenFile(t, localFile)
    requestCustomXmlPart := models.CustomXmlPartInsert{
        Id: ToStringPointer("hello"),
        Data: ToStringPointer("<data>Hello world</data>"),
    }

    options := map[string]interface{}{
    }

    request := &models.InsertCustomXmlPartOnlineRequest{
        Document: requestDocument,
        CustomXmlPart: &requestCustomXmlPart,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.InsertCustomXmlPartOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetModel().GetCustomXmlPart(), "Validate InsertCustomXmlPartOnline response.");
    assert.Equal(t, "hello", DereferenceValue(actual.GetModel().GetCustomXmlPart().GetId()), "Validate InsertCustomXmlPartOnline response.");
    assert.Equal(t, "<data>Hello world</data>", DereferenceValue(actual.GetModel().GetCustomXmlPart().GetData()), "Validate InsertCustomXmlPartOnline response.");
}

// Test for updating custom xml part.
func Test_CustomXmlParts_UpdateCustomXmlPart(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/CustomXmlParts"
    localFile := "DocumentElements/CustomXmlParts/MultipleCustomXmlParts.docx"
    remoteFileName := "TestUpdateCustomXmlPart.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestCustomXmlPart := models.CustomXmlPartUpdate{
        Data: ToStringPointer("<data>Hello world</data>"),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.UpdateCustomXmlPartRequest{
        Name: ToStringPointer(remoteFileName),
        CustomXmlPartIndex: ToInt32Pointer(int32(0)),
        CustomXmlPart: &requestCustomXmlPart,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.UpdateCustomXmlPart(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetCustomXmlPart(), "Validate UpdateCustomXmlPart response.");
    assert.Equal(t, "aspose", DereferenceValue(actual.GetCustomXmlPart().GetId()), "Validate UpdateCustomXmlPart response.");
    assert.Equal(t, "<data>Hello world</data>", DereferenceValue(actual.GetCustomXmlPart().GetData()), "Validate UpdateCustomXmlPart response.");
}

// Test for updating custom xml part online.
func Test_CustomXmlParts_UpdateCustomXmlPartOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/CustomXmlParts/MultipleCustomXmlParts.docx"

    requestDocument := OpenFile(t, localFile)
    requestCustomXmlPart := models.CustomXmlPartUpdate{
        Data: ToStringPointer("<data>Hello world</data>"),
    }

    options := map[string]interface{}{
    }

    request := &models.UpdateCustomXmlPartOnlineRequest{
        Document: requestDocument,
        CustomXmlPartIndex: ToInt32Pointer(int32(0)),
        CustomXmlPart: &requestCustomXmlPart,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.UpdateCustomXmlPartOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetModel().GetCustomXmlPart(), "Validate UpdateCustomXmlPartOnline response.");
    assert.Equal(t, "aspose", DereferenceValue(actual.GetModel().GetCustomXmlPart().GetId()), "Validate UpdateCustomXmlPartOnline response.");
    assert.Equal(t, "<data>Hello world</data>", DereferenceValue(actual.GetModel().GetCustomXmlPart().GetData()), "Validate UpdateCustomXmlPartOnline response.");
}

// A test for DeleteCustomXmlPart.
func Test_CustomXmlParts_DeleteCustomXmlPart(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/CustomXmlParts"
    localFile := "DocumentElements/CustomXmlParts/MultipleCustomXmlParts.docx"
    remoteFileName := "TestDeleteCustomXmlPart.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }

    request := &models.DeleteCustomXmlPartRequest{
        Name: ToStringPointer(remoteFileName),
        CustomXmlPartIndex: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, err := client.WordsApi.DeleteCustomXmlPart(ctx, request)

    if err != nil {
        t.Error(err)
    }

}

// A test for DeleteCustomXmlPart online.
func Test_CustomXmlParts_DeleteCustomXmlPartOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/CustomXmlParts/MultipleCustomXmlParts.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.DeleteCustomXmlPartOnlineRequest{
        Document: requestDocument,
        CustomXmlPartIndex: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.DeleteCustomXmlPartOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// A test for DeleteCustomXmlParts.
func Test_CustomXmlParts_DeleteCustomXmlParts(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/CustomXmlParts"
    localFile := "DocumentElements/CustomXmlParts/MultipleCustomXmlParts.docx"
    remoteFileName := "TestDeleteCustomXmlPart.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }

    request := &models.DeleteCustomXmlPartsRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: options,
    }

    _, err := client.WordsApi.DeleteCustomXmlParts(ctx, request)

    if err != nil {
        t.Error(err)
    }

}

// A test for DeleteCustomXmlParts online.
func Test_CustomXmlParts_DeleteCustomXmlPartsOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/CustomXmlParts/MultipleCustomXmlParts.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.DeleteCustomXmlPartsOnlineRequest{
        Document: requestDocument,
        Optionals: options,
    }

    _, _, err := client.WordsApi.DeleteCustomXmlPartsOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}
