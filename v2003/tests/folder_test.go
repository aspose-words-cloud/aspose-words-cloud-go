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
	"fmt"
	"path"
	"testing"

	"github.com/google/uuid"
)

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

func TestCreateFolder(t *testing.T) {

	remoteFolder := path.Join(remoteBaseTestDataFolder, "Storage", fmt.Sprintf("TestCreateFolder%s", uuid.New().String()))

	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	_, err := client.WordsApi.CreateFolder(ctx, remoteFolder, nil)

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

func TestGetFilesList(t *testing.T) {

	remoteFolder := path.Join(remoteBaseTestDataFolder, "Storage")

	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	_, _, err := client.WordsApi.GetFilesList(ctx, remoteFolder, nil)

	if err != nil {
		t.Error(err)
	}
}
