/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="drawing_objects_test.go">
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

// Example of how to get drawing objects.
package api_test

import (
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2509/api/models"
)

// Test for getting drawing objects from document.
func Test_DrawingObjects_GetDocumentDrawingObjects(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/DrawingObjectss"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetDocumentDrawingObjects.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "sections/0",
        "folder": remoteDataFolder,
    }

    request := &models.GetDocumentDrawingObjectsRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetDocumentDrawingObjects(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for getting drawing objects from document online.
func Test_DrawingObjects_GetDocumentDrawingObjectsOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
        "nodePath": "sections/0",
    }

    request := &models.GetDocumentDrawingObjectsOnlineRequest{
        Document: requestDocument,
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetDocumentDrawingObjectsOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for getting drawing objects from document without node path.
func Test_DrawingObjects_GetDocumentDrawingObjectsWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/DrawingObjectss"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetDocumentDrawingObjectsWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.GetDocumentDrawingObjectsRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetDocumentDrawingObjects(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for getting drawing object by specified index.
func Test_DrawingObjects_GetDocumentDrawingObjectByIndex(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/DrawingObjectss"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetDocumentDrawingObjectByIndex.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "sections/0",
        "folder": remoteDataFolder,
    }

    request := &models.GetDocumentDrawingObjectByIndexRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetDocumentDrawingObjectByIndex(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for getting drawing object by specified index online.
func Test_DrawingObjects_GetDocumentDrawingObjectByIndexOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
        "nodePath": "sections/0",
    }

    request := &models.GetDocumentDrawingObjectByIndexOnlineRequest{
        Document: requestDocument,
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetDocumentDrawingObjectByIndexOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for getting drawing object by specified index without node path.
func Test_DrawingObjects_GetDocumentDrawingObjectByIndexWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/DrawingObjectss"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetDocumentDrawingObjectByIndexWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.GetDocumentDrawingObjectByIndexRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetDocumentDrawingObjectByIndex(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for getting drawing object by specified index and format.
func Test_DrawingObjects_RenderDrawingObject(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/DrawingObjectss"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetDocumentDrawingObjectByIndexWithFormat.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "sections/0",
        "folder": remoteDataFolder,
    }

    request := &models.RenderDrawingObjectRequest{
        Name: ToStringPointer(remoteFileName),
        Format: ToStringPointer("png"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, err := client.WordsApi.RenderDrawingObject(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for getting drawing object by specified index and format online.
func Test_DrawingObjects_RenderDrawingObjectOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
        "nodePath": "sections/0",
    }

    request := &models.RenderDrawingObjectOnlineRequest{
        Document: requestDocument,
        Format: ToStringPointer("png"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, err := client.WordsApi.RenderDrawingObjectOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for getting drawing object by specified index and format without node path.
func Test_DrawingObjects_RenderDrawingObjectWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/DrawingObjectss"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetDocumentDrawingObjectByIndexWithFormatWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.RenderDrawingObjectRequest{
        Name: ToStringPointer(remoteFileName),
        Format: ToStringPointer("png"),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, err := client.WordsApi.RenderDrawingObject(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for reading drawing object's image data.
func Test_DrawingObjects_GetDocumentDrawingObjectImageData(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/DrawingObjectss"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetDocumentDrawingObjectImageData.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "sections/0",
        "folder": remoteDataFolder,
    }

    request := &models.GetDocumentDrawingObjectImageDataRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, err := client.WordsApi.GetDocumentDrawingObjectImageData(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for reading drawing object's image data online.
func Test_DrawingObjects_GetDocumentDrawingObjectImageDataOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
        "nodePath": "sections/0",
    }

    request := &models.GetDocumentDrawingObjectImageDataOnlineRequest{
        Document: requestDocument,
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, err := client.WordsApi.GetDocumentDrawingObjectImageDataOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for reading drawing object's image data without node path.
func Test_DrawingObjects_GetDocumentDrawingObjectImageDataWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/DrawingObjectss"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetDocumentDrawingObjectImageDataWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.GetDocumentDrawingObjectImageDataRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, err := client.WordsApi.GetDocumentDrawingObjectImageData(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for getting drawing object OLE data.
func Test_DrawingObjects_GetDocumentDrawingObjectOleData(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/DrawingObjectss"
    localDrawingFile := "DocumentElements/DrawingObjects/sample_EmbeddedOLE.docx"
    remoteFileName := "TestGetDocumentDrawingObjectOleData.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localDrawingFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "sections/0",
        "folder": remoteDataFolder,
    }

    request := &models.GetDocumentDrawingObjectOleDataRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, err := client.WordsApi.GetDocumentDrawingObjectOleData(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for getting drawing object OLE data online.
func Test_DrawingObjects_GetDocumentDrawingObjectOleDataOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localDrawingFile := "DocumentElements/DrawingObjects/sample_EmbeddedOLE.docx"

    requestDocument := OpenFile(t, localDrawingFile)

    options := map[string]interface{}{
        "nodePath": "sections/0",
    }

    request := &models.GetDocumentDrawingObjectOleDataOnlineRequest{
        Document: requestDocument,
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, err := client.WordsApi.GetDocumentDrawingObjectOleDataOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for getting drawing object OLE data without node path.
func Test_DrawingObjects_GetDocumentDrawingObjectOleDataWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/DrawingObjectss"
    localDrawingFile := "DocumentElements/DrawingObjects/sample_EmbeddedOLE.docx"
    remoteFileName := "TestGetDocumentDrawingObjectOleDataWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localDrawingFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.GetDocumentDrawingObjectOleDataRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, err := client.WordsApi.GetDocumentDrawingObjectOleData(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for adding drawing object.
func Test_DrawingObjects_InsertDrawingObject(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/DrawingObjectss"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestInsetDrawingObject.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestDrawingObject := models.DrawingObjectInsert{
        Height: ToFloat64Pointer(0),
        Left: ToFloat64Pointer(0),
        Top: ToFloat64Pointer(0),
        Width: ToFloat64Pointer(0),
        RelativeHorizontalPosition: ToStringPointer("Margin"),
        RelativeVerticalPosition: ToStringPointer("Margin"),
        WrapType: ToStringPointer("Inline"),
    }
    requestImageFile := OpenFile(t, "Common/aspose-cloud.png")

    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }

    request := &models.InsertDrawingObjectRequest{
        Name: ToStringPointer(remoteFileName),
        DrawingObject: &requestDrawingObject,
        ImageFile: requestImageFile,
        Optionals: options,
    }

    _, _, err := client.WordsApi.InsertDrawingObject(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for adding drawing object online.
func Test_DrawingObjects_InsertDrawingObjectOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)
    requestDrawingObject := models.DrawingObjectInsert{
        Height: ToFloat64Pointer(0),
        Left: ToFloat64Pointer(0),
        Top: ToFloat64Pointer(0),
        Width: ToFloat64Pointer(0),
        RelativeHorizontalPosition: ToStringPointer("Margin"),
        RelativeVerticalPosition: ToStringPointer("Margin"),
        WrapType: ToStringPointer("Inline"),
    }
    requestImageFile := OpenFile(t, "Common/aspose-cloud.png")

    options := map[string]interface{}{
        "nodePath": "",
    }

    request := &models.InsertDrawingObjectOnlineRequest{
        Document: requestDocument,
        DrawingObject: &requestDrawingObject,
        ImageFile: requestImageFile,
        Optionals: options,
    }

    _, _, err := client.WordsApi.InsertDrawingObjectOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for adding drawing object without node path.
func Test_DrawingObjects_InsertDrawingObjectWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/DrawingObjectss"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestInsetDrawingObjectWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestDrawingObject := models.DrawingObjectInsert{
        Height: ToFloat64Pointer(0),
        Left: ToFloat64Pointer(0),
        Top: ToFloat64Pointer(0),
        Width: ToFloat64Pointer(0),
        RelativeHorizontalPosition: ToStringPointer("Margin"),
        RelativeVerticalPosition: ToStringPointer("Margin"),
        WrapType: ToStringPointer("Inline"),
    }
    requestImageFile := OpenFile(t, "Common/aspose-cloud.png")

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.InsertDrawingObjectRequest{
        Name: ToStringPointer(remoteFileName),
        DrawingObject: &requestDrawingObject,
        ImageFile: requestImageFile,
        Optionals: options,
    }

    _, _, err := client.WordsApi.InsertDrawingObject(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for deleting drawing object.
func Test_DrawingObjects_DeleteDrawingObject(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/DrawingObjectss"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestDeleteDrawingObject.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }

    request := &models.DeleteDrawingObjectRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, err := client.WordsApi.DeleteDrawingObject(ctx, request)

    if err != nil {
        t.Error(err)
    }

}

// Test for deleting drawing object online.
func Test_DrawingObjects_DeleteDrawingObjectOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
        "nodePath": "",
    }

    request := &models.DeleteDrawingObjectOnlineRequest{
        Document: requestDocument,
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.DeleteDrawingObjectOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for deleting drawing object without node path.
func Test_DrawingObjects_DeleteDrawingObjectWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/DrawingObjectss"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestDeleteDrawingObjectWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.DeleteDrawingObjectRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, err := client.WordsApi.DeleteDrawingObject(ctx, request)

    if err != nil {
        t.Error(err)
    }

}

// Test for updating drawing object.
func Test_DrawingObjects_UpdateDrawingObject(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/DrawingObjectss"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestUpdateDrawingObject.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestDrawingObject := models.DrawingObjectUpdate{
        Left: ToFloat64Pointer(0),
    }
    requestImageFile := OpenFile(t, "Common/aspose-cloud.png")

    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }

    request := &models.UpdateDrawingObjectRequest{
        Name: ToStringPointer(remoteFileName),
        DrawingObject: &requestDrawingObject,
        ImageFile: requestImageFile,
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.UpdateDrawingObject(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for updating drawing object online.
func Test_DrawingObjects_UpdateDrawingObjectOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)
    requestDrawingObject := models.DrawingObjectUpdate{
        Left: ToFloat64Pointer(0),
    }
    requestImageFile := OpenFile(t, "Common/aspose-cloud.png")

    options := map[string]interface{}{
        "nodePath": "",
    }

    request := &models.UpdateDrawingObjectOnlineRequest{
        Document: requestDocument,
        DrawingObject: &requestDrawingObject,
        ImageFile: requestImageFile,
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.UpdateDrawingObjectOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for updating drawing object without node path.
func Test_DrawingObjects_UpdateDrawingObjectWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/DrawingObjectss"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestUpdateDrawingObjectWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestDrawingObject := models.DrawingObjectUpdate{
        Left: ToFloat64Pointer(0),
    }
    requestImageFile := OpenFile(t, "Common/aspose-cloud.png")

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.UpdateDrawingObjectRequest{
        Name: ToStringPointer(remoteFileName),
        DrawingObject: &requestDrawingObject,
        ImageFile: requestImageFile,
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.UpdateDrawingObject(ctx, request)
    if err != nil {
        t.Error(err)
    }

}
