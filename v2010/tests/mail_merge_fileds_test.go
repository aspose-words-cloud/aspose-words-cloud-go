/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="mail_merge_fileds_test.go">
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

// Example of how to work with merge fields.
package api_test

import (
    "testing"
)

// Test for putting new fields.
func Test_MailMergeFileds_GetDocumentFieldNamesOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    mailMergeFolder := "DocumentActions/MailMerge"
    localDocumentFile := "SampleExecuteTemplate.docx"


    options := map[string]interface{}{
        "useNonMergeFields": true,
    }
    _, _, err := client.WordsApi.GetDocumentFieldNamesOnline(ctx, OpenFile(t, mailMergeFolder + "/" + localDocumentFile), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for getting mailmerge fields.
func Test_MailMergeFileds_GetDocumentFieldNames(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentActions/MailMerge"
    remoteFileName := "TestGetDocumentFieldNames.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile("Common/test_multi_pages.docx"), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetDocumentFieldNames(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }

}
