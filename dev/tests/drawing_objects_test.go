/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="drawing_objects_test.go">
 *   Copyright (c) 2020 Aspose.Words for Cloud
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
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/dev/api/models"
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
    actual, _, err := client.WordsApi.GetDocumentDrawingObjects(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.DrawingObjects, "Validate GetDocumentDrawingObjects response.");
    assert.NotNil(t, actual.DrawingObjects.List, "Validate GetDocumentDrawingObjects response.");
    assert.Equal(t, 1, len(actual.DrawingObjects.List), "Validate GetDocumentDrawingObjects response.");
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
    actual, _, err := client.WordsApi.GetDocumentDrawingObjects(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.DrawingObjects, "Validate GetDocumentDrawingObjectsWithoutNodePath response.");
    assert.NotNil(t, actual.DrawingObjects.List, "Validate GetDocumentDrawingObjectsWithoutNodePath response.");
    assert.Equal(t, 1, len(actual.DrawingObjects.List), "Validate GetDocumentDrawingObjectsWithoutNodePath response.");
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
    actual, _, err := client.WordsApi.GetDocumentDrawingObjectByIndex(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.DrawingObject, "Validate GetDocumentDrawingObjectByIndex response.");
    assert.Equal(t, 300.0, actual.DrawingObject.Height, "Validate GetDocumentDrawingObjectByIndex response.");
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
    actual, _, err := client.WordsApi.GetDocumentDrawingObjectByIndex(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.DrawingObject, "Validate GetDocumentDrawingObjectByIndexWithoutNodePath response.");
    assert.Equal(t, 300.0, actual.DrawingObject.Height, "Validate GetDocumentDrawingObjectByIndexWithoutNodePath response.");
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
    _, err := client.WordsApi.RenderDrawingObject(ctx, remoteFileName, "png", int32(0), options)

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
    _, err := client.WordsApi.RenderDrawingObject(ctx, remoteFileName, "png", int32(0), options)

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
    _, err := client.WordsApi.GetDocumentDrawingObjectImageData(ctx, remoteFileName, int32(0), options)

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
    _, err := client.WordsApi.GetDocumentDrawingObjectImageData(ctx, remoteFileName, int32(0), options)

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
    _, err := client.WordsApi.GetDocumentDrawingObjectOleData(ctx, remoteFileName, int32(0), options)

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
    _, err := client.WordsApi.GetDocumentDrawingObjectOleData(ctx, remoteFileName, int32(0), options)

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
        Height: 0.0,
        Left: 0.0,
        Top: 0.0,
        Width: 0.0,
        RelativeHorizontalPosition: "Margin",
        RelativeVerticalPosition: "Margin",
        WrapType: "Inline",
    }

    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.InsertDrawingObject(ctx, remoteFileName, requestDrawingObject, OpenFile(t, "Common/aspose-cloud.png"), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.DrawingObject, "Validate InsertDrawingObject response.");
    assert.Equal(t, "0.3.7.1", actual.DrawingObject.NodeId, "Validate InsertDrawingObject response.");
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
        Height: 0.0,
        Left: 0.0,
        Top: 0.0,
        Width: 0.0,
        RelativeHorizontalPosition: "Margin",
        RelativeVerticalPosition: "Margin",
        WrapType: "Inline",
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.InsertDrawingObject(ctx, remoteFileName, requestDrawingObject, OpenFile(t, "Common/aspose-cloud.png"), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.DrawingObject, "Validate InsertDrawingObjectWithoutNodePath response.");
    assert.Equal(t, "0.3.7.1", actual.DrawingObject.NodeId, "Validate InsertDrawingObjectWithoutNodePath response.");
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
    _, err := client.WordsApi.DeleteDrawingObject(ctx, remoteFileName, int32(0), options)

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
    _, err := client.WordsApi.DeleteDrawingObject(ctx, remoteFileName, int32(0), options)

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
        Left: 1.0,
    }

    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.UpdateDrawingObject(ctx, remoteFileName, requestDrawingObject, OpenFile(t, "Common/aspose-cloud.png"), int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.DrawingObject, "Validate UpdateDrawingObject response.");
    assert.Equal(t, 1.0, actual.DrawingObject.Left, "Validate UpdateDrawingObject response.");
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
        Left: 1.0,
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    actual, _, err := client.WordsApi.UpdateDrawingObject(ctx, remoteFileName, requestDrawingObject, OpenFile(t, "Common/aspose-cloud.png"), int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.DrawingObject, "Validate UpdateDrawingObjectWithoutNodePath response.");
    assert.Equal(t, 1.0, actual.DrawingObject.Left, "Validate UpdateDrawingObjectWithoutNodePath response.");
}
