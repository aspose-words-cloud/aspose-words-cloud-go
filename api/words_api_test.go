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
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/aspose-words-cloud/aspose-words-cloud-go/api"
	"github.com/google/uuid"
)

var remoteBaseTestDataFolder string = "Temp/SdkTests/TestData"

var commonTestFile string = GetLocalPath("Common", "test_multi_pages.docx")

func GetLocalPath(folderPath string, filename string) string {
	_, currentFilePath, _, _ := runtime.Caller(0)
	currentFolder := filepath.Dir(currentFilePath)
	return filepath.Join(currentFolder, "testdata", folderPath, filename)
}

func ReadConfiguration(t *testing.T) (cfg *api.Configuration) {
	_, filename, _, _ := runtime.Caller(0)
	configFile := filepath.Join(filepath.Dir(filename), "../config.json")
	configuration, err := api.NewConfiguration(configFile)

	if err != nil {
		t.Error(err)
	}

	return configuration
}

func PrepareTest(t *testing.T, config *api.Configuration) (apiClient *api.APIClient, ctx context.Context) {
	client, err := api.NewAPIClient(config)

	if err != nil {
		t.Error(err)
	}

	context, err := client.NewContextWithToken(nil)

	if err != nil {
		t.Error(err)
	}

	return client, context
}

func UploadFileToStorage(t *testing.T, fileName string, path string) (*api.APIClient, context.Context) {
	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	UploadNextFileToStorage(t, ctx, client, fileName, path)

	return client, ctx
}

func UploadNextFileToStorage(t *testing.T, ctx context.Context, client *api.APIClient, fileName string, path string) {
	file, fileErr := os.Open(fileName)
	if fileErr != nil {
		t.Error(fileErr)
	}

	_, _, err := client.WordsApi.UploadFile(ctx, file, path, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestDebugMode(t *testing.T) {
	config := ReadConfiguration(t)
	config.DebugMode = true
	client, ctx := PrepareTest(t, config)

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	_, _, err := client.WordsApi.CreateDocument(ctx, map[string]interface{}{
		"fileName": "test.docx",
	})

	if err != nil {
		t.Error(err)
	}

	logData := buf.String()

	url := "PUT /v4.0/words/create?FileName=test.docx"
	success := " 200 OK"

	if !strings.Contains(logData, url) {
		t.Errorf("Output doesn't contain the string '%s'", url)
	}

	if !strings.Contains(logData, success) {
		t.Errorf("Output doesn't contain the string '%s'", success)
	}
}

func TestTestCoverage(t *testing.T) {
	_, testFileName, _, _ := runtime.Caller(0)
	apiFileName := filepath.Join(filepath.Dir(testFileName), "words_api.go")
	api, err := ioutil.ReadFile(apiFileName)

	if err != nil {
		t.Error(err)
	}

	declareFuncRegex := regexp.MustCompile(`func \(a \*WordsApiService\) (\w+)`)
	allApiFuncNames := declareFuncRegex.FindAllSubmatch(api, -1)

	var exists = struct{}{}

	notTestedApiFuncNames := make(map[string]struct{})

	for _, name := range allApiFuncNames {
		notTestedApiFuncNames[string(name[1])] = exists
	}

	tests, err := ioutil.ReadFile(testFileName)
	if err != nil {
		t.Error(err)
	}

	callFuncRegex := regexp.MustCompile(`\.WordsApi\.(\w+)`)
	testedlApiFuncNames := callFuncRegex.FindAllSubmatch(tests, -1)

	for _, name := range testedlApiFuncNames {
		delete(notTestedApiFuncNames, string(name[1]))
	}

	if len(notTestedApiFuncNames) > 0 {
		report := "Not covered methods:\n"

		for name := range notTestedApiFuncNames {
			report = report + fmt.Sprintf("\t%s\n", name)
		}

		t.Error(report)
	}
}

func TestDeleteField(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Fields"), "GetField.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Fields")
	remoteName := "TestDeleteField.docx"
	nodePath := "sections/0/paragraphs/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteField(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestAcceptAllRevisions(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "Revisions")
	remoteName := "TestAcceptAllRevisions.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.AcceptAllRevisions(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestAppendDocument(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "AppendDocument")
	remoteName := "TestAppendDocument.docx"
	remoteFilePath := path.Join(remoteFolder, remoteName)
	documentList := api.DocumentEntryList{
		DocumentEntries: []api.DocumentEntry{
			api.DocumentEntry{
				Href:             remoteFilePath,
				ImportFormatMode: "KeepSourceFormatting",
			},
		},
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, remoteFilePath)

	_, _, err := client.WordsApi.AppendDocument(ctx, remoteName, documentList, options)

	if err != nil {
		t.Error(err)
	}
}

func TestClassify(t *testing.T) {

	text := "Try text classification"
	options := map[string]interface{}{
		"bestClassesCount": "3",
	}

	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	_, _, err := client.WordsApi.Classify(ctx, text, options)

	if err != nil {
		t.Error(err)
	}
}

func TestClassifyDocument(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("Common"), "test_multi_pages.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "Common")
	remoteName := "TestClassifyDocument.docx"
	options := map[string]interface{}{
		"folder":           remoteFolder,
		"bestClassesCount": "3",
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.ClassifyDocument(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestCompareDocument(t *testing.T) {

	localFilePath1 := GetLocalPath(filepath.Join("DocumentActions", "CompareDocument"), "compareTestDoc1.doc")
	localFilePath2 := GetLocalPath(filepath.Join("DocumentActions", "CompareDocument"), "compareTestDoc2.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "CompareDocument")
	remoteName1 := "TestCompareDocument1.docx"
	remoteName2 := "TestCompareDocument2.docx"
	remotePath1 := path.Join(remoteFolder, remoteName1)
	remotePath2 := path.Join(remoteFolder, remoteName2)
	compareData := api.CompareData{
		ComparingWithDocument: remotePath2,
		Author:                "author",
		DateTime:              time.Date(2015, time.October, 26, 0, 0, 0, 0, time.UTC),
	}
	options := map[string]interface{}{
		"folder":       remoteFolder,
		"destFileName": path.Join("TestOut", "TestCompareDocumentOut.doc"),
	}

	client, ctx := UploadFileToStorage(t, localFilePath1, remotePath1)
	UploadNextFileToStorage(t, ctx, client, localFilePath2, remotePath2)

	_, _, err := client.WordsApi.CompareDocument(ctx, remoteName1, compareData, options)

	if err != nil {
		t.Error(err)
	}
}

func TestConvertDocument(t *testing.T) {

	format := "pdf"
	localFilePath := GetLocalPath(filepath.Join("DocumentActions", "ConvertDocument"), "test_uploadfile.docx")
	document, fileErr := os.Open(localFilePath)
	if fileErr != nil {
		t.Error(fileErr)
	}

	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	output, err := client.WordsApi.ConvertDocument(ctx, document, format, nil)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}

func TestCopyFile(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "Storage")
	remoteSrcName := "TestCopyFile.docx"
	remoteDstName := fmt.Sprintf("TestCopyFile%s.docx", uuid.New().String())
	remoteSrcPath := path.Join(remoteFolder, remoteSrcName)
	remoteDstPath := path.Join(remoteFolder, remoteDstName)

	client, ctx := UploadFileToStorage(t, localFilePath, remoteSrcPath)

	_, err := client.WordsApi.CopyFile(ctx, remoteDstPath, remoteSrcPath, nil)

	if err != nil {
		t.Error(err)
	}
}

func TestCopyFolder(t *testing.T) {

	guid := uuid.New().String()
	remoteSrcFolder := path.Join(remoteBaseTestDataFolder, "Storage", fmt.Sprintf("TestCopyFolder%sSrc", guid))
	remoteDstFolder := path.Join(remoteBaseTestDataFolder, "Storage", fmt.Sprintf("TestCopyFolder%sDst", guid))

	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	_, folderErr := client.WordsApi.CreateFolder(ctx, remoteSrcFolder, nil)
	if folderErr != nil {
		t.Error(folderErr)
	}

	_, err := client.WordsApi.CopyFolder(ctx, remoteDstFolder, remoteSrcFolder, nil)

	if err != nil {
		t.Error(err)
	}
}

func TestCreateDocument(t *testing.T) {

	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "Document")
	remoteName := "TestCreateDocument.docx"
	options := map[string]interface{}{
		"fileName": remoteName,
		"folder":   remoteFolder,
	}

	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	_, _, err := client.WordsApi.CreateDocument(ctx, options)

	if err != nil {
		t.Error(err)
	}
}

func TestCreateFolder(t *testing.T) {

	remoteFolder := path.Join(remoteBaseTestDataFolder, "Storage", fmt.Sprintf("TestCreateFolder%s", uuid.New().String()))

	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	_, err := client.WordsApi.CreateFolder(ctx, remoteFolder, nil)

	if err != nil {
		t.Error(err)
	}
}

func TestCreateOrUpdateDocumentProperty(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := "TestOut"
	remoteName := "TestCreateOrUpdateDocumentProperty.docx"
	propertyName := "AsposeAuthor"
	property := api.DocumentProperty{
		Name:  "Author",
		Value: "Imran Anwar",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.CreateOrUpdateDocumentProperty(ctx, remoteName, propertyName, property, options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteBorder(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestDeleteBorder.docx"
	nodePath := "tables/1/rows/0/cells/0"
	borderType := "Left"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.DeleteBorder(ctx, remoteName, nodePath, borderType, options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteBorders(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestDeleteBorders.docx"
	nodePath := "tables/1/rows/0/cells/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.DeleteBorders(ctx, remoteName, nodePath, options)

	if err != nil {
		t.Error(err)
	}
}

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

func TestDeleteDocumentProperty(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DocumentProperties")
	remoteName := "TestDeleteDocumentProperty.docx"
	propertyName := "testProp"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteDocumentProperty(ctx, remoteName, propertyName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteDrawingObject(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DrawingObjects")
	remoteName := "TestDeleteDrawingObject.docx"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteDrawingObject(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteDrawingObjectWithoutNodePath(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DrawingObjects")
	remoteName := "TestDeleteDrawingObject.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteDrawingObjectWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteFieldWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Fields"), "GetField.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Fields")
	remoteName := "TestDeleteFieldWithoutNode.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteFieldWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteFields(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Fields")
	remoteName := "TestDeleteFields.docx"
	nodePath := "sections/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteFields(ctx, remoteName, nodePath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteFieldsWithoutNodePath(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Fields")
	remoteName := "TestDeleteFieldsWithoutNodePath.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteFieldsWithoutNodePath(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteFile(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "Storage")
	remoteName := "TestDeleteFile.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteFile(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteFolder(t *testing.T) {

	guid := uuid.New().String()
	remoteFolder := path.Join(remoteBaseTestDataFolder, "Storage", fmt.Sprintf("TestDeleteFolder%s", guid))

	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	_, folderErr := client.WordsApi.CreateFolder(ctx, remoteFolder, nil)
	if folderErr != nil {
		t.Error(folderErr)
	}

	_, err := client.WordsApi.DeleteFolder(ctx, remoteFolder, nil)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteFootnote(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Footnotes"), "Footnote.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Footnotes")
	remoteName := "TestDeleteFootnote.docx"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteFootnote(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteFootnoteWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Footnotes"), "Footnote.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Footnotes")
	remoteName := "TestDeleteFootnoteWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteFootnoteWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteFormField(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "FormFields"), "FormFilled.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "FormFields")
	remoteName := "TestDeleteFormField.docx"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteFormField(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteFormFieldWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "FormFields"), "FormFilled.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "FormFields")
	remoteName := "TestDeleteFormFieldWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteFormFieldWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteHeaderFooter(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "HeaderFooters"), "HeadersFooters.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "HeaderFooters")
	remoteName := "TestDeleteHeaderFooter.docx"
	sectionPath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteHeaderFooter(ctx, remoteName, sectionPath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteHeadersFooters(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "HeaderFooters"), "HeadersFooters.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "HeaderFooters")
	remoteName := "TestDeleteHeaderFooters.docx"
	sectionPath := "sections/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteHeadersFooters(ctx, remoteName, sectionPath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteMacros(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Macros")
	remoteName := "TestDeleteMacros.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteMacros(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteOfficeMathObject(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "MathObjects"), "MathObjects.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "MathObjects")
	remoteName := "TestDeleteOfficeMathObject.docx"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteOfficeMathObject(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteOfficeMathObjectWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "MathObjects"), "MathObjects.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "MathObjects")
	remoteName := "TestDeleteOfficeMathObjectWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteOfficeMathObjectWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteParagraph(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestDeleteParagraph.docx"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteParagraph(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteParagraphWithoutNodePath(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestDeleteParagraphWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteParagraphWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteRun(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Runs"), "Run.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Runs")
	remoteName := "TestDeleteRun.docx"
	paragraphPath := "paragraphs/1"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteRun(ctx, remoteName, paragraphPath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteSection(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Section")
	remoteName := "TestDeleteSection.docx"
	sectionIndex := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteSection(ctx, remoteName, int32(sectionIndex), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteTable(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestDeleteTable.docx"
	nodePath := "sections/0"
	index := 1
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteTable(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteTableCell(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestDeleteTableCell.docx"
	tableRowPath := "sections/0/tables/2/rows/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteTableCell(ctx, remoteName, tableRowPath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteTableRow(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestDeleteTableRow.docx"
	tablePath := "tables/1"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteTableRow(ctx, remoteName, tablePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteTableWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestDeleteTableWithoutNode.docx"
	index := 1
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteTableWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteWatermark(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "Watermark")
	remoteName := "TestDeleteWatermark.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.DeleteWatermark(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestDownloadFile(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "Storage")
	remoteName := "TestDownloadFile.docx"
	fullName := path.Join(remoteFolder, remoteName)
	options := map[string]interface{}{}

	client, ctx := UploadFileToStorage(t, localFilePath, fullName)

	output, err := client.WordsApi.DownloadFile(ctx, fullName, options)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}

func TestExecuteMailMerge(t *testing.T) {

	localFolder := filepath.Join("DocumentActions", "MailMerge")
	localFilePath := GetLocalPath(localFolder, "SampleMailMergeTemplate.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "MailMerge")
	remoteName := "TestExecuteMailMerge.docx"
	data, fileErr := ioutil.ReadFile(GetLocalPath(localFolder, "SampleMailMergeTemplateData.txt"))
	if fileErr != nil {
		t.Error(fileErr)
	}
	options := map[string]interface{}{
		"data":        string(data),
		"folder":      remoteFolder,
		"withRegions": false,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.ExecuteMailMerge(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestExecuteMailMergeOnline(t *testing.T) {

	localFolder := filepath.Join("DocumentActions", "MailMerge")
	localTemplatePath := GetLocalPath(localFolder, "SampleMailMergeTemplate.docx")
	localDataPath := GetLocalPath(localFolder, "SampleMailMergeTemplateData.txt")

	template, err := os.Open(localTemplatePath)
	if err != nil {
		t.Error(err)
	}
	data, err := os.Open(localDataPath)
	if err != nil {
		t.Error(err)
	}

	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	output, err := client.WordsApi.ExecuteMailMergeOnline(ctx, template, data, nil)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}

func TestGetAvailableFonts(t *testing.T) {

	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	_, _, err := client.WordsApi.GetAvailableFonts(ctx, nil)

	if err != nil {
		t.Error(err)
	}
}

func TestGetBookmarkByName(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Bookmarks")
	remoteName := "TestGetBookmarkByName.docx"
	bookmarkName := "aspose"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetBookmarkByName(ctx, remoteName, bookmarkName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetBookmarks(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Bookmarks")
	remoteName := "TestGetBookmarks.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetBookmarks(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetBorder(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestGetBorder.docx"
	nodePath := "tables/1/rows/0/cells/0"
	borderType := "Left"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetBorder(ctx, remoteName, nodePath, borderType, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetBorders(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestGetBorders.docx"
	nodePath := "tables/1/rows/0/cells/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetBorders(ctx, remoteName, nodePath, options)

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

func TestGetDocument(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "Document")
	remoteName := "TestGetDocument.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetDocument(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocumentDrawingObjectByIndex(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DrawingObjects")
	remoteName := "TestGetDocumentDrawingObjectByIndex.docx"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetDocumentDrawingObjectByIndex(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocumentDrawingObjectByIndexWithoutNodePath(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DrawingObjects")
	remoteName := "TestGetDocumentDrawingObjectByIndexWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetDocumentDrawingObjectByIndexWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocumentDrawingObjectImageData(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DrawingObjects")
	remoteName := "TestGetDocumentDrawingObjectImageData.docx"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	output, err := client.WordsApi.GetDocumentDrawingObjectImageData(ctx, remoteName, nodePath, int32(index), options)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocumentDrawingObjectImageDataWithoutNodePath(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DrawingObjects")
	remoteName := "TestGetDocumentDrawingObjectImageDataWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	output, err := client.WordsApi.GetDocumentDrawingObjectImageDataWithoutNodePath(ctx, remoteName, int32(index), options)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocumentDrawingObjectOleData(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "DrawingObjects"), "sample_EmbeddedOLE.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DrawingObjects")
	remoteName := "TestGetDocumentDrawingObjectOleData.docx"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	output, err := client.WordsApi.GetDocumentDrawingObjectOleData(ctx, remoteName, nodePath, int32(index), options)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocumentDrawingObjectOleDataWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "DrawingObjects"), "sample_EmbeddedOLE.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DrawingObjects")
	remoteName := "TestGetDocumentDrawingObjectOleDataWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	output, err := client.WordsApi.GetDocumentDrawingObjectOleDataWithoutNodePath(ctx, remoteName, int32(index), options)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocumentDrawingObjects(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DrawingObjects")
	remoteName := "TestGetDocumentDrawingObjects.docx"
	nodePath := "sections/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetDocumentDrawingObjects(ctx, remoteName, nodePath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocumentDrawingObjectsWithoutNodePath(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DrawingObjects")
	remoteName := "TestGetDocumentDrawingObjectsWithoutNodePath.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetDocumentDrawingObjectsWithoutNodePath(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocumentFieldNames(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "MailMerge")
	remoteName := "TestGetDocumentFieldNames.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetDocumentFieldNames(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocumentFieldNamesOnline(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentActions", "MailMerge"), "SampleExecuteTemplate.docx")
	template, fileErr := os.Open(localFilePath)
	if fileErr != nil {
		t.Error(fileErr)
	}

	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	_, _, err := client.WordsApi.GetDocumentFieldNamesOnline(ctx, template, nil)

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocumentHyperlinkByIndex(t *testing.T) {

	localFilePath := GetLocalPath("Common", "test_doc.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Hyperlink")
	remoteName := "TestGetDocumentHyperlinkByIndex.docx"
	hyperlinkIndex := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetDocumentHyperlinkByIndex(ctx, remoteName, int32(hyperlinkIndex), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocumentHyperlinks(t *testing.T) {

	localFilePath := GetLocalPath("Common", "test_doc.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Hyperlink")
	remoteName := "TestGetDocumentHyperlinks.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetDocumentHyperlinks(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocumentProperties(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DocumentProperties")
	remoteName := "TestGetDocumentProperties.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetDocumentProperties(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocumentProperty(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DocumentProperties")
	remoteName := "TestGetDocumentProperty.docx"
	propertyName := "Author"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetDocumentProperty(ctx, remoteName, propertyName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocumentProtection(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "DocumentProtection")
	remoteName := "TestGetDocumentProtection.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetDocumentProtection(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocumentStatistics(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "Statistics")
	remoteName := "TestGetDocumentStatistics.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetDocumentStatistics(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocumentWithFormat(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "DocumentWithFormat")
	remoteName := "TestGetDocumentWithFormat.docx"
	format := "text"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	output, err := client.WordsApi.GetDocumentWithFormat(ctx, remoteName, format, options)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}

func TestGetField(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Fields"), "GetField.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Fields")
	remoteName := "TestGetField.docx"
	nodePath := "sections/0/paragraphs/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetField(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetFieldWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Fields"), "GetField.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Fields")
	remoteName := "TestGetFieldWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetFieldWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetFields(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Fields"), "GetField.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Fields")
	remoteName := "TestGetFields.docx"
	nodePath := "sections/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetFields(ctx, remoteName, nodePath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetFieldsWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Fields"), "GetField.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Fields")
	remoteName := "TestGetFieldsWithoutNodePath.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetFieldsWithoutNodePath(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetFilesList(t *testing.T) {

	remoteFolder := path.Join(remoteBaseTestDataFolder, "Storage")

	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	_, _, err := client.WordsApi.GetFilesList(ctx, remoteFolder, nil)

	if err != nil {
		t.Error(err)
	}
}

func TestGetFootnote(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Footnotes"), "Footnote.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Footnotes")
	remoteName := "TestGetFootnote.docx"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetFootnote(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetFootnoteWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Footnotes"), "Footnote.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Footnotes")
	remoteName := "TestGetFootnoteWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetFootnoteWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetFootnotes(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Footnotes"), "Footnote.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Footnotes")
	nodePath := "sections/0"
	remoteName := "TestGetFootnotes.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetFootnotes(ctx, remoteName, nodePath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetFootnotesWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Footnotes"), "Footnote.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Footnotes")
	remoteName := "TestGetFootnotesWithoutNodePath.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetFootnotesWithoutNodePath(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetFormField(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "FormFields"), "FormFilled.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "FormFields")
	remoteName := "TestGetFormField.docx"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetFormField(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetFormFieldWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "FormFields"), "FormFilled.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "FormFields")
	remoteName := "TestGetFormFieldWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetFormFieldWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetFormFields(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "FormFields"), "FormFilled.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "FormFields")
	remoteName := "TestGetFormFields.docx"
	nodePath := "sections/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetFormFields(ctx, remoteName, nodePath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetFormFieldsWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "FormFields"), "FormFilled.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "FormFields")
	remoteName := "TestGetFormFieldsWithoutNodePath.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetFormFieldsWithoutNodePath(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetHeaderFooter(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "HeaderFooters"), "HeadersFooters.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "HeaderFooters")
	remoteName := "TestGetHeaderFooter.docx"
	headerFooterIndex := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetHeaderFooter(ctx, remoteName, int32(headerFooterIndex), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetHeaderFooterOfSection(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "HeaderFooters"), "HeadersFooters.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "HeaderFooters")
	remoteName := "TestGetHeaderFooterOfSection.docx"
	sectionIndex := 0
	headerFooterIndex := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetHeaderFooterOfSection(ctx, remoteName, int32(headerFooterIndex), int32(sectionIndex), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetHeaderFooters(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "HeaderFooters"), "HeadersFooters.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "HeaderFooters")
	remoteName := "TestGetHeaderFooters.docx"
	sectionPath := "sections/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetHeaderFooters(ctx, remoteName, sectionPath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetOfficeMathObject(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "MathObjects"), "MathObjects.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "MathObjects")
	remoteName := "TestGetOfficeMathObject.docx"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetOfficeMathObject(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetOfficeMathObjectWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "MathObjects"), "MathObjects.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "MathObjects")
	remoteName := "TestGetOfficeMathObjectWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetOfficeMathObjectWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetOfficeMathObjects(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "MathObjects"), "MathObjects.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "MathObjects")
	remoteName := "TestGetOfficeMathObjects.docx"
	nodePath := "sections/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetOfficeMathObjects(ctx, remoteName, nodePath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetOfficeMathObjectsWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "MathObjects"), "MathObjects.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "MathObjects")
	remoteName := "TestGetOfficeMathObjectsWithoutNodePath.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetOfficeMathObjectsWithoutNodePath(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetParagraph(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestGetParagraph.docx"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetParagraph(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetParagraphFormat(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestGetParagraphFormat.docx"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetParagraphFormat(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetParagraphFormatWithoutNodePath(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestGetParagraphFormatWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetParagraphFormatWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetParagraphWithoutNodePath(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestGetParagraphWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetParagraphWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetParagraphs(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestGetParagraphs.docx"
	nodePath := "sections/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetParagraphs(ctx, remoteName, nodePath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetParagraphsWithoutNodePath(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestGetParagraphsWithoutNodePath.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetParagraphsWithoutNodePath(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetRangeText(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Range"), "RangeGet.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Range")
	remoteName := "TestGetRangeText.docx"
	rangeStartIdentifier := "id0.0.0"
	rangeEndIdentifier := "id0.0.1"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetRangeText(ctx, remoteName, rangeStartIdentifier, rangeEndIdentifier, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetRun(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestGetRun.docx"
	paragraphPath := "paragraphs/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetRun(ctx, remoteName, paragraphPath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetRunFont(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestGetRunFont.docx"
	paragraphPath := "paragraphs/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetRunFont(ctx, remoteName, paragraphPath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetRuns(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestGetRuns.docx"
	paragraphPath := "paragraphs/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetRuns(ctx, remoteName, paragraphPath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetSection(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Section")
	remoteName := "TestGetSection.docx"
	sectionIndex := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetSection(ctx, remoteName, int32(sectionIndex), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetSectionPageSetup(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "PageSetup")
	remoteName := "TestGetSectionPageSetup.docx"
	sectionIndex := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetSectionPageSetup(ctx, remoteName, int32(sectionIndex), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetSections(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Section")
	remoteName := "TestGetSections.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetSections(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetTable(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestGetTable.docx"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetTable(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetTableCell(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	tableRowPath := "sections/0/tables/2/rows/0"
	index := 0
	remoteName := "TestGetTableCell.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetTableCell(ctx, remoteName, tableRowPath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetTableCellFormat(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	tableRowPath := "sections/0/tables/2/rows/0"
	index := 0
	remoteName := "TestGetTableCellFormat.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetTableCellFormat(ctx, remoteName, tableRowPath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetTableProperties(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	nodePath := "sections/0"
	index := 0
	remoteName := "TestGetTableProperties.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetTableProperties(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetTablePropertiesWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestGetTablePropertiesWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetTablePropertiesWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetTableRow(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestGetTableRow.docx"
	tablePath := "tables/1"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetTableRow(ctx, remoteName, tablePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetTableRowFormat(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestGetTableRowFormat.docx"
	tablePath := "tables/1"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetTableRowFormat(ctx, remoteName, tablePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetTableWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestGetTableWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetTableWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetTables(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestGetTables.docx"
	nodePath := "sections/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetTables(ctx, remoteName, nodePath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetTablesWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestGetTablesWithoutNodePath.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetTablesWithoutNodePath(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertComment(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Comments")
	remoteName := "TestInsertComment.docx"
	position := api.DocumentPosition{
		Node: &api.NodeLink{
			NodeId: "0.3.0.3",
		},
		Offset: 0,
	}
	comment := api.Comment{
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

func TestInsertDrawingObject(t *testing.T) {

	localFilePath := commonTestFile
	localImagePath := GetLocalPath("Common", "aspose-cloud.png")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DrawingObjects")
	remoteName := "TestInsertDrawingObject.docx"
	imageFile, fileErr := os.Open(localImagePath)
	if fileErr != nil {
		t.Error(fileErr)
	}
	drawingObject := "{\"Left\": 0}"
	nodePath := "sections/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertDrawingObject(ctx, remoteName, drawingObject, imageFile, nodePath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertDrawingObjectWithoutNodePath(t *testing.T) {

	localFilePath := commonTestFile
	localImagePath := GetLocalPath("Common", "aspose-cloud.png")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DrawingObjects")
	remoteName := "TestInsertDrawingObjectWithoutNodePath.docx"
	imageFile, fileErr := os.Open(localImagePath)
	if fileErr != nil {
		t.Error(fileErr)
	}
	drawingObject := "{\"Left\": 0}"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertDrawingObjectWithoutNodePath(ctx, remoteName, drawingObject, imageFile, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertField(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Text"), "SampleWordDocument.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Fields")
	remoteName := "TestInsertField.docx"
	nodePath := "sections/0/paragraphs/0"
	field := api.Field{
		NodeId:    "0.0.3",
		FieldCode: "{ NUMPAGES }",
		Result:    "3",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertField(ctx, remoteName, field, nodePath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertFieldWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Text"), "SampleWordDocument.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Fields")
	remoteName := "TestInsertFieldWithoutNodePath.docx"
	field := api.Field{
		NodeId:    "0.0.3",
		FieldCode: "{ NUMPAGES }",
		Result:    "3",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertFieldWithoutNodePath(ctx, remoteName, field, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertFootnote(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Footnotes"), "Footnote.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Footnotes")
	remoteName := "TestInsertFootnote.docx"
	footnoteDto := api.Footnote{
		FootnoteType: "Endnote",
		Text:         "test endnote",
	}
	nodePath := "sections/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertFootnote(ctx, remoteName, footnoteDto, nodePath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertFootnoteWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Footnotes"), "Footnote.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Footnotes")
	remoteName := "TestInsertFootnoteWithoutNodePath.docx"
	footnoteDto := api.Footnote{
		FootnoteType: "Endnote",
		Text:         "test endnote",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertFootnoteWithoutNodePath(ctx, remoteName, footnoteDto, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertFormField(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "FormFields")
	remoteName := "TestInsertFormField.docx"
	var formField api.IFormField = api.FormFieldTextInput{
		Name:             "FullName",
		Enabled:          true,
		CalculateOnExit:  true,
		StatusText:       "",
		TextInputType:    "Regular",
		TextInputDefault: "123",
		TextInputFormat:  "UPPERCASE",
	}
	nodePath := "sections/0/paragraphs/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertFormField(ctx, remoteName, formField, nodePath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertFormFieldWithoutNodePath(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "FormFields")
	remoteName := "TestInsertFormFieldWithoutNodePath.docx"
	formField := api.FormFieldTextInput{
		Name:             "FullName",
		Enabled:          true,
		CalculateOnExit:  true,
		StatusText:       "",
		TextInputType:    "Regular",
		TextInputDefault: "123",
		TextInputFormat:  "UPPERCASE",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertFormFieldWithoutNodePath(ctx, remoteName, formField, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertHeaderFooter(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "HeaderFooters"), "HeadersFooters.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "HeaderFooters")
	remoteName := "TestInsertHeaderFooter.docx"
	headerFooterType := "FooterEven"
	sectionPath := "sections/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertHeaderFooter(ctx, remoteName, headerFooterType, sectionPath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertPageNumbers(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Fields")
	remoteName := "TestInsertPageNumbers.docx"
	pageNumber := api.PageNumber{
		Alignment: "center",
		Format:    "{PAGE} of {NUMPAGES}",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertPageNumbers(ctx, remoteName, pageNumber, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertParagraph(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestInsertParagraph.docx"
	nodePath := "sections/0"
	paragraph := api.ParagraphInsert{
		Text: "This is a new paragraph for your document",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertParagraph(ctx, remoteName, paragraph, nodePath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertRun(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Runs"), "Run.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Runs")
	remoteName := "TestInsertRun.docx"
	paragraphPath := "paragraphs/1"
	run := api.Run{
		Text: "run with text",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertRun(ctx, remoteName, paragraphPath, run, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertTable(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestInsertTable.docx"
	nodePath := "sections/0"

	options := map[string]interface{}{
		"folder": remoteFolder,
		"table": api.TableInsert{
			ColumnsCount: 5,
			RowsCount:    4,
		},
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertTable(ctx, remoteName, nodePath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertTableCell(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestInsertTableCell.docx"
	tableRowPath := "sections/0/tables/2/rows/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
		"cell":   api.TableCellInsert{},
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertTableCell(ctx, remoteName, tableRowPath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertTableRow(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	tablePath := "sections/0/tables/2"
	remoteName := "TestInsertTableRow.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
		"row": api.TableRowInsert{
			ColumnsCount: 5,
		},
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertTableRow(ctx, remoteName, tablePath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertTableWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestInsertTableWithoutNodePath.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
		"table": api.TableInsert{
			ColumnsCount: 5,
			RowsCount:    4,
		},
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertTableWithoutNodePath(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertWatermarkImage(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "Watermark")
	remoteName := "TestInsertWatermarkImage.docx"
	localImagePath := GetLocalPath("Common", "aspose-cloud.png")
	remoteImageName := "TestInsertWatermarkImage.png"

	options := map[string]interface{}{
		"folder":        remoteFolder,
		"image":         path.Join(remoteFolder, remoteImageName),
		"rotationAngle": 0.0,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))
	UploadNextFileToStorage(t, ctx, client, localImagePath, path.Join(remoteFolder, remoteImageName))

	_, _, err := client.WordsApi.InsertWatermarkImage(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertWatermarkText(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "Watermark")
	remoteName := "TestInsertWatermarkText.docx"
	watermarkText := api.WatermarkText{
		Text:          "This is the text",
		RotationAngle: 90.0,
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertWatermarkText(ctx, remoteName, watermarkText, options)

	if err != nil {
		t.Error(err)
	}
}

func TestLoadWebDocument(t *testing.T) {

	data := api.LoadWebDocumentData{
		LoadingDocumentUrl: "http://google.com",
		SaveOptions: &api.SaveOptionsData{
			FileName:                "google.doc",
			SaveFormat:              "doc",
			DmlEffectsRenderingMode: "1",
			DmlRenderingMode:        "1",
			UpdateSdtContent:        false,
			ZipOutput:               false,
		},
	}

	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	_, _, err := client.WordsApi.LoadWebDocument(ctx, data, nil)

	if err != nil {
		t.Error(err)
	}
}

func TestMoveFile(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "Storage")
	remoteSrcName := "TestMoveFile.docx"
	remoteDstName := fmt.Sprintf("TesMoveFile%s.docx", uuid.New().String())
	remoteSrcPath := path.Join(remoteFolder, remoteSrcName)
	remoteDstPath := path.Join(remoteFolder, remoteDstName)

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteSrcName))

	_, err := client.WordsApi.MoveFile(ctx, remoteDstPath, remoteSrcPath, nil)

	if err != nil {
		t.Error(err)
	}
}

func TestMoveFolder(t *testing.T) {

	guid := uuid.New().String()
	remoteSrcFolder := path.Join(remoteBaseTestDataFolder, "Storage", fmt.Sprintf("TestMoveFolder%sSrc", guid))
	remoteDstFolder := path.Join(remoteBaseTestDataFolder, "Storage", fmt.Sprintf("TestMoveFolder%sDst", guid))

	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	_, err := client.WordsApi.MoveFolder(ctx, remoteSrcFolder, remoteDstFolder, nil)

	if err != nil {
		t.Error(err)
	}
}

func TestProtectDocument(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "DocumentProtection")
	remoteName := "TestProtectDocument.docx"
	protectionRequest := api.ProtectionRequest{
		NewPassword: "123",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.ProtectDocument(ctx, remoteName, protectionRequest, options)

	if err != nil {
		t.Error(err)
	}
}

func TestRejectAllRevisions(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "Revisions")
	remoteName := "TestRejectAllRevisions.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.RejectAllRevisions(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestRemoveRange(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Range"), "RangeGet.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Range")
	remoteName := "TestRemoveRange.docx"
	rangeStartIdentifier := "id0.0.0"
	rangeEndIdentifier := "id0.0.1"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.RemoveRange(ctx, remoteName, rangeStartIdentifier, rangeEndIdentifier, options)

	if err != nil {
		t.Error(err)
	}
}

func TestRenderDrawingObject(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DrawingObjects")
	remoteName := "TestRenderDrawingObject.docx"
	format := "png"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	output, err := client.WordsApi.RenderDrawingObject(ctx, remoteName, format, nodePath, int32(index), options)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}

func TestRenderDrawingObjectWithoutNodePath(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DrawingObjects")
	remoteName := "TestRenderDrawingObjectWithoutNodePath.docx"
	format := "png"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	output, err := client.WordsApi.RenderDrawingObjectWithoutNodePath(ctx, remoteName, format, int32(index), options)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}

func TestRenderMathObject(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "MathObjects"), "MathObjects.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "MathObjects")
	remoteName := "TestRenderMathObject.docx"
	format := "png"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	output, err := client.WordsApi.RenderMathObject(ctx, remoteName, format, nodePath, int32(index), options)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}

func TestRenderMathObjectWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "MathObjects"), "MathObjects.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "MathObjects")
	remoteName := "TestRenderMathObjectWithoutNodePath.docx"
	format := "png"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	output, err := client.WordsApi.RenderMathObjectWithoutNodePath(ctx, remoteName, format, int32(index), options)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}

func TestRenderPage(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Text"), "SampleWordDocument.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "PageSetup")
	remoteName := "TestRenderPage.docx"
	format := "bmp"
	pageIndex := 1
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	output, err := client.WordsApi.RenderPage(ctx, remoteName, int32(pageIndex), format, options)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}

func TestRenderParagraph(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestRenderParagraph.docx"
	nodePath := "sections/0"
	format := "png"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	output, err := client.WordsApi.RenderParagraph(ctx, remoteName, format, nodePath, int32(index), options)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}

func TestRenderParagraphWithoutNodePath(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestRenderParagraphWithoutNodePath.docx"
	format := "png"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	output, err := client.WordsApi.RenderParagraphWithoutNodePath(ctx, remoteName, format, int32(index), options)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}

func TestRenderTable(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestRenderTable.docx"
	format := "png"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	output, err := client.WordsApi.RenderTable(ctx, remoteName, format, nodePath, int32(index), options)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}

func TestRenderTableWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestRenderTableWithoutNodePath.docx"
	format := "png"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	output, err := client.WordsApi.RenderTableWithoutNodePath(ctx, remoteName, format, int32(index), options)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}

func TestReplaceText(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Text")
	remoteName := "TestReplaceText.docx"
	replaceText := api.ReplaceTextParameters{
		OldValue: "aspose",
		NewValue: "aspose new",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.ReplaceText(ctx, remoteName, replaceText, options)

	if err != nil {
		t.Error(err)
	}
}

func TestReplaceWithText(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Range"), "RangeGet.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Range")
	remoteName := "TestReplaceWithText.docx"
	rangeStartIdentifier := "id0.0.0"
	rangeEndIdentifier := "id0.0.1"
	rangeText := api.ReplaceRange{
		Text: "Replaced header",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.ReplaceWithText(ctx, remoteName, rangeStartIdentifier, rangeText, rangeEndIdentifier, options)

	if err != nil {
		t.Error(err)
	}
}

func TestResetCache(t *testing.T) {

	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	_, err := client.WordsApi.ResetCache(ctx)

	if err != nil {
		t.Error(err)
	}
}

func TestSaveAs(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "ConvertDocument")
	remoteName := "TestSaveAs.docx"
	remoteDstName := "TestSaveAs.pdf"
	saveOptionsData := api.SaveOptionsData{
		SaveFormat: "pdf",
		FileName:   path.Join(remoteFolder, remoteDstName),
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.SaveAs(ctx, remoteName, saveOptionsData, options)

	if err != nil {
		t.Error(err)
	}
}

func TestSaveAsRange(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Range"), "RangeGet.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Range")
	remoteName := "TestSaveAsRange.docx"
	rangeStartIdentifier := "id0.0.0"
	rangeEndIdentifier := "id0.0.1"
	newDocName := "NewDoc.docx"
	documentParameters := api.RangeDocument{
		DocumentName: path.Join(remoteFolder, newDocName),
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.SaveAsRange(ctx, remoteName, rangeStartIdentifier, documentParameters, rangeEndIdentifier, options)

	if err != nil {
		t.Error(err)
	}
}

func TestSaveAsTiff(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "ConvertDocument")
	remoteName := "TestSaveAsTiff.docx"
	destFileName := "TestSaveAsTiff.tiff"
	saveOptions := api.TiffSaveOptionsData{
		FileName:   path.Join(remoteFolder, destFileName),
		SaveFormat: "tiff",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.SaveAsTiff(ctx, remoteName, saveOptions, options)

	if err != nil {
		t.Error(err)
	}
}

func TestSearch(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Text"), "SampleWordDocument.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Text")
	remoteName := "TestSearch.docx"
	pattern := "aspose"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.Search(ctx, remoteName, pattern, options)

	if err != nil {
		t.Error(err)
	}
}

func TestSplitDocument(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "SplitDocument")
	remoteName := "TestSplitDocument.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
		"format": "text",
		"from":   int32(1),
		"to":     int32(2),
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.SplitDocument(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestUnprotectDocument(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentActions", "DocumentProtection"), "SampleProtectedBlankWordDocument.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "DocumentProtection")
	remoteName := "TestUnprotectDocument.docx"
	protectionRequest := api.ProtectionRequest{
		Password: "aspose",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UnprotectDocument(ctx, remoteName, protectionRequest, options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateBookmark(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Bookmarks")
	remoteName := "TestUpdateBookmark.docx"
	bookmarkName := "aspose"
	bookmarkData := api.BookmarkData{
		Name: bookmarkName,
		Text: "This will be the text for Aspose",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateBookmark(ctx, remoteName, bookmarkData, bookmarkName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateBorder(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestUpdateBorder.docx"
	nodePath := "tables/1/rows/0/cells/0"
	borderType := "Left"
	borderProperties := api.Border{
		BorderType: "Left",
		Color: &api.XmlColor{
			Alpha: 2,
		},
		DistanceFromText: 6.0,
		LineStyle:        "DashDotStroker",
		LineWidth:        2.0,
		Shadow:           true,
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateBorder(ctx, remoteName, borderProperties, nodePath, borderType, options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateComment(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Comments")
	remoteName := "TestUpdateComment.docx"
	commentIndex := 0
	documentPosition := api.DocumentPosition{
		Node: &api.NodeLink{
			NodeId: "0.3.0",
		},
		Offset: 0,
	}
	comment := api.Comment{
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

func TestUpdateDrawingObject(t *testing.T) {

	localFilePath := commonTestFile
	localImagePath := GetLocalPath("Common", "aspose-cloud.png")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DrawingObjects")
	remoteName := "TestUpdateDrawingObject.docx"
	imageFile, fileErr := os.Open(localImagePath)
	if fileErr != nil {
		t.Error(fileErr)
	}
	nodePath := "sections/0"
	index := 0
	drawingObject := "{\"Left\": 0}"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateDrawingObject(ctx, remoteName, drawingObject, imageFile, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateDrawingObjectWithoutNodePath(t *testing.T) {

	localFilePath := commonTestFile
	localImagePath := GetLocalPath("Common", "aspose-cloud.png")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DrawingObjects")
	remoteName := "TestUpdateDrawingObjectWithoutNodePath.docx"
	imageFile, fileErr := os.Open(localImagePath)
	if fileErr != nil {
		t.Error(fileErr)
	}
	index := 0
	drawingObject := "{\"Left\": 0}"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateDrawingObjectWithoutNodePath(ctx, remoteName, drawingObject, imageFile, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateField(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Fields"), "GetField.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Fields")
	remoteName := "TestUpdateField.docx"
	nodePath := "sections/0/paragraphs/0"
	index := 0
	field := api.Field{
		Result:    "3",
		FieldCode: "{ NUMPAGES }",
		NodeId:    "0.0.3",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateField(ctx, remoteName, field, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateFields(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Fields")
	remoteName := "TestUpdateFields.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateFields(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateFootnote(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Footnotes"), "Footnote.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Footnotes")
	remoteName := "TestUpdateFootnote.docx"
	nodePath := "sections/0"
	index := 0
	footnoteDto := api.Footnote{
		Text: "new text is here",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateFootnote(ctx, remoteName, footnoteDto, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateFootnoteWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Footnotes"), "Footnote.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Footnotes")
	remoteName := "TestUpdateFootnoteWithoutNodePath.docx"
	index := 0
	footnoteDto := api.Footnote{
		Text: "new text is here",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateFootnoteWithoutNodePath(ctx, remoteName, footnoteDto, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateFormField(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "FormFields"), "FormFilled.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "FormFields")
	remoteName := "TestUpdateFormField.docx"
	nodePath := "sections/0"
	index := 0
	formField := api.FormFieldTextInput{
		Name:             "FullName",
		Enabled:          true,
		CalculateOnExit:  true,
		TextInputType:    "Regular",
		TextInputDefault: "No name",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateFormField(ctx, remoteName, formField, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateFormFieldWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "FormFields"), "FormFilled.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "FormFields")
	remoteName := "TestUpdateFormFieldWithoutNodePath.docx"
	index := 0
	formField := api.FormFieldTextInput{
		Name:             "FullName",
		Enabled:          true,
		CalculateOnExit:  true,
		TextInputType:    "Regular",
		TextInputDefault: "No name",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateFormFieldWithoutNodePath(ctx, remoteName, formField, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateParagraphFormat(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestUpdateParagraphFormat.docx"
	nodePath := "sections/0"
	index := 0
	dto := api.ParagraphFormat{
		Alignment: "Right",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateParagraphFormat(ctx, remoteName, dto, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateRun(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Runs"), "Run.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Runs")
	remoteName := "TestUpdateRun.docx"
	paragraphPath := "paragraphs/1"
	index := 0
	run := api.Run{
		Text: "run with text",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateRun(ctx, remoteName, run, paragraphPath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateRunFont(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestUpdateRunFont.docx"
	paragraphPath := "paragraphs/0"
	index := 0
	fontDto := api.Font{
		Bold: true,
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateRunFont(ctx, remoteName, fontDto, paragraphPath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateSectionPageSetup(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "PageSetup")
	remoteName := "TestUpdateSectionPageSetup.docx"
	sectionIndex := 0
	pageSetup := api.PageSetup{
		RtlGutter:   true,
		LeftMargin:  10.0,
		Orientation: "Landscape",
		PaperSize:   "A5",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateSectionPageSetup(ctx, remoteName, int32(sectionIndex), pageSetup, options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateTableCellFormat(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestUpdateTableCellFormat.docx"
	tableRowPath := "sections/0/tables/2/rows/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
		"format": api.TableCellFormat{
			BottomPadding:   5,
			FitText:         true,
			HorizontalMerge: "First",
			WrapText:        true,
		},
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateTableCellFormat(ctx, remoteName, tableRowPath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateTableProperties(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestUpdateTableProperties.docx"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
		"properties": api.TableProperties{
			Alignment:     "Right",
			AllowAutoFit:  false,
			Bidi:          true,
			BottomPadding: 1,
			CellSpacing:   2,
			LeftIndent:    3,
			LeftPadding:   4,
			RightPadding:  5,
			StyleOptions:  "ColumnBands",
			TopPadding:    6,
		},
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateTableProperties(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateTablePropertiesWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestUpdateTablePropertiesWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
		"properties": api.TableProperties{
			Alignment:     "Right",
			AllowAutoFit:  false,
			Bidi:          true,
			BottomPadding: 1,
			CellSpacing:   2,
			LeftIndent:    3,
			LeftPadding:   4,
			RightPadding:  5,
			StyleOptions:  "ColumnBands",
			TopPadding:    6,
		},
	}
	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateTablePropertiesWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateTableRowFormat(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Tables"), "TablesGet.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Tables")
	remoteName := "TestUpdateTableRowFormat.docx"
	tablePath := "sections/0/tables/2"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
		"format": api.TableRowFormat{
			AllowBreakAcrossPages: true,
			HeadingFormat:         true,
			Height:                10,
			HeightRule:            "Auto",
		},
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateTableRowFormat(ctx, remoteName, tablePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}
