/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="readme_test.go">
 *   Copyright (c) 2026 Aspose.Words for Cloud
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
    "bufio"
    "bytes"
    "fmt"
    "io"
    "os"
    "path"
    "path/filepath"
    "regexp"
    "runtime"
    "testing"

    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2602/api"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2602/api/models"
)

func TestReadmeCode(t *testing.T) {

    localFilePath := commonTestFile
    remoteFolder := path.Join(remoteBaseTestDataFolder, "ReadmeCode")
    remoteName := "ReadmeCode.docx"
    remotePath := path.Join(remoteFolder, remoteName)
    configFilePath := GetConfigFilePath()

    // Start README example

    // init words cloud api
    config, _ := models.NewConfiguration(configFilePath)
    wordsApi, ctx, _ := api.CreateWordsApi(config)

    // upload test.docx to a cloud
    // remote.docx is a name in the cloud
        file, _ := os.Open(localFilePath)
    uploadRequest := &models.UploadFileRequest{
        FileContent: file,
        Path: &remotePath,
    }
    wordsApi.UploadFile(ctx, uploadRequest)

    // get a text for the first paragraph of the first section
    options := map[string]interface{}{
        "folder": remoteFolder,
    }
    request := &models.GetParagraphsRequest{
        Name: &remoteName,
        Optionals: options,
    }

    result, _, _ := wordsApi.GetParagraphs(ctx, request)

    fmt.Println(result.Paragraphs.GetParagraphLinkList()[0].GetText())

    // End README example

    writeToReadme(t)
}

func writeToReadme(t *testing.T) {
    // set regexes
    startPatern := regexp.MustCompile("^\\s*// Start README example\\s*$")
    endPattern := regexp.MustCompile("^\\s*// End README example\\s*$")

    // set paths
    _, filename, _, _ := runtime.Caller(0)
    testFolder := filepath.Dir(filename)
    sourcePath := filepath.Join(testFolder, "readme_test.go")
    readmePath := filepath.Join(testFolder, "../../README.md")

    // read the file
    codeLines, err := readLines(sourcePath)
    if err != nil {
        t.Error(err)
    }

    // extract example code
    var readmeCode []string
    skipMode := true
    for i := 0; i < len(codeLines); i++ {
        if skipMode {
            skipMode = !startPatern.Match([]byte(codeLines[i]))
        }

        if !skipMode {
            readmeCode = append(readmeCode, codeLines[i])
            skipMode = endPattern.Match([]byte(codeLines[i]))
        }
    }

    if len(readmeCode) < 2 {
        t.Error("Code has not found in readme_test.go")
    }

    // read readme.md
    readmeLines, err := readLines(readmePath)
    if err != nil {
        t.Error(err)
    }

    // replace example code
    var newReadmeLines []string
    codeMode := false
    for i := 0; i < len(readmeLines); i++ {
        if !codeMode {
            codeMode = startPatern.Match([]byte(readmeLines[i]))

            if codeMode {
                newReadmeLines = append(newReadmeLines, readmeCode...)
            }
        }

        if codeMode {
            codeMode = !endPattern.Match([]byte(readmeLines[i]))
            continue
        }

        if !codeMode {
            newReadmeLines = append(newReadmeLines, readmeLines[i])
        }
    }

    // write new readme
    err = writeLines(newReadmeLines, readmePath)
    if err != nil {
        t.Error(err)
    }
}

// Read a whole file into the memory and store it as array of lines
func readLines(path string) (lines []string, err error) {
    var (
        file   *os.File
        part   []byte
        prefix bool
    )
    if file, err = os.Open(path); err != nil {
        return
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    buffer := bytes.NewBuffer(make([]byte, 0))
    for {
        if part, prefix, err = reader.ReadLine(); err != nil {
            break
        }
        buffer.Write(part)
        if !prefix {
            lines = append(lines, buffer.String())
            buffer.Reset()
        }
    }
    if err == io.EOF {
        err = nil
    }
    return
}

func writeLines(lines []string, path string) (err error) {
    var (
        file *os.File
    )

    if file, err = os.Create(path); err != nil {
        return
    }
    defer file.Close()

    for _, item := range lines {
        _, err := file.WriteString(item + "\n")

        if err != nil {
            break
        }
    }

    return
}
