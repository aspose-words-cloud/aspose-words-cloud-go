/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="math_object_test.go">
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

// Example of how to work with MathObjects.
package api_test

import (
    "testing"
)

// Test for getting mathObjects.
func Test_MathObject_GetOfficeMathObjects(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/MathObjects"
    localFile := "DocumentElements/MathObjects/MathObjects.docx"
    remoteFileName := "TestGetOfficeMathObjects.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetOfficeMathObjects(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }
}

// Test for getting mathObjects without node path.
func Test_MathObject_GetOfficeMathObjectsWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/MathObjects"
    localFile := "DocumentElements/MathObjects/MathObjects.docx"
    remoteFileName := "TestGetOfficeMathObjectsWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetOfficeMathObjects(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }
}

// Test for getting mathObject.
func Test_MathObject_GetOfficeMathObject(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/MathObjects"
    localFile := "DocumentElements/MathObjects/MathObjects.docx"
    remoteFileName := "TestGetOfficeMathObject.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetOfficeMathObject(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }
}

// Test for getting mathObject without node path.
func Test_MathObject_GetOfficeMathObjectWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/MathObjects"
    localFile := "DocumentElements/MathObjects/MathObjects.docx"
    remoteFileName := "TestGetOfficeMathObjectWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetOfficeMathObject(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }
}

// Test for rendering mathObject.
func Test_MathObject_RenderMathObject(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/MathObjects"
    localFile := "DocumentElements/MathObjects/MathObjects.docx"
    remoteFileName := "TestRenderMathObject.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }
    _, err := client.WordsApi.RenderMathObject(ctx, remoteFileName, "png", int32(0), options)

    if err != nil {
        t.Error(err)
    }
}

// Test for rendering mathObject without node path.
func Test_MathObject_RenderMathObjectWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/MathObjects"
    localFile := "DocumentElements/MathObjects/MathObjects.docx"
    remoteFileName := "TestRenderMathObjectWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, err := client.WordsApi.RenderMathObject(ctx, remoteFileName, "png", int32(0), options)

    if err != nil {
        t.Error(err)
    }
}

// Test for deleting mathObject.
func Test_MathObject_DeleteOfficeMathObject(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/MathObjects"
    localFile := "DocumentElements/MathObjects/MathObjects.docx"
    remoteFileName := "TestDeleteOfficeMathObject.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }
    _, err := client.WordsApi.DeleteOfficeMathObject(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }
}

// Test for deleting mathObject without node path.
func Test_MathObject_DeleteOfficeMathObjectWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/MathObjects"
    localFile := "DocumentElements/MathObjects/MathObjects.docx"
    remoteFileName := "TestDeleteOfficeMathObjectWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, err := client.WordsApi.DeleteOfficeMathObject(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }
}
