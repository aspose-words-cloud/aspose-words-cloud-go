/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="batch_test.go">
 *   Copyright (c) 2021 Aspose.Words for Cloud
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

package api_test

import (
	// "os"
	"reflect"
	"testing"

	"github.com/aspose-words-cloud/aspose-words-cloud-go/dev/api/models"
	"github.com/stretchr/testify/assert"
)

// Test for a batch request.
func Test_BatchRequest(t *testing.T) {
	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Paragraphs"
	localFile := "Common/test_multi_pages.docx"
	// reportingFolder := "DocumentActions/Reporting"
	remoteFileName := "TestGetDocumentParagraphByIndex.docx"

	// localDocumentFile := "ReportTemplate.docx"
	// localDataFile := reportingFolder + "/ReportData.json"

	UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder+"/"+remoteFileName)

	options := map[string]interface{}{
		"nodePath": "sections/0",
		"folder":   remoteDataFolder,
	}

	request1 := &models.GetParagraphsRequest{
		Name:      ToStringPointer(remoteFileName),
		Optionals: options,
	}

	request2 := &models.GetParagraphRequest{
		Name:      ToStringPointer(remoteFileName),
		Index:     ToInt32Pointer(0),
		Optionals: options,
	}

	request3 := &models.InsertParagraphRequest{
		Name: ToStringPointer(remoteFileName),
		Paragraph: models.ParagraphInsert{
			Text: ToStringPointer("This is a new paragraph for your document"),
		},
		Optionals: options,
	}

	request4 := &models.DeleteParagraphRequest{
		Name:  ToStringPointer(remoteFileName),
		Index: ToInt32Pointer(0),
		Optionals: map[string]interface{}{
			"nodePath": "",
			"folder":   remoteDataFolder,
		},
	}

/*
	templateFile, err := os.Open(GetLocalFile(reportingFolder + "/" + localDocumentFile))
	if err != nil {
		t.Error(err)
	}

	request5 := &models.BuildReportOnlineRequest{
		Template: templateFile,
		Data:     ToStringPointer(ReadFile(t, localDataFile)),
		ReportEngineSettings: models.ReportEngineSettings{
			DataSourceType: ToStringPointer("Json"),
			DataSourceName: ToStringPointer("persons"),
		},
	}

	actual, _, err := client.WordsApi.Batch(ctx, request1, request2, request3, request4, request5)

	assert.Nil(t, err, err)
	assert.True(t, len(actual) == 5)
	assert.Equal(t, "ParagraphLinkCollectionResponse", reflect.TypeOf(actual[0]).Name(), "Validate GetParagraphs response.")
	assert.Equal(t, "ParagraphResponse", reflect.TypeOf(actual[1]).Name(), "Validate GetParagraph response.")
	assert.Equal(t, "ParagraphResponse", reflect.TypeOf(actual[2]).Name(), "Validate InsertParagraph response.")
	assert.Nil(t, actual[3], "Validate DeleteParagraph response.")
	assert.Equal(t, "Reader", reflect.TypeOf(actual[4]).Elem().Name(), "Validate BuildReportOnline response.")
*/

	actual, _, err := client.WordsApi.Batch(ctx, request1, request2, request3, request4)

	assert.Nil(t, err, err)
	assert.True(t, len(actual) == 4)
	assert.Equal(t, "ParagraphLinkCollectionResponse", reflect.TypeOf(actual[0]).Name(), "Validate GetParagraphs response.")
	assert.Equal(t, "ParagraphResponse", reflect.TypeOf(actual[1]).Name(), "Validate GetParagraph response.")
	assert.Equal(t, "ParagraphResponse", reflect.TypeOf(actual[2]).Name(), "Validate InsertParagraph response.")
	assert.Nil(t, actual[3], "Validate DeleteParagraph response.")
}