/*
* MIT License

* Copyright (c) 2019 Aspose Pty Ltd

* Permission is hereby granted, free of charge, to any person obtaining a copy
* of this software and associated documentation files (the "Software"), to deal
* in the Software without restriction, including without limitation the rights
* to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
* copies of the Software, and to permit persons to whom the Software is
* furnished to do so, subject to the following conditions:

* The above copyright notice and this permission notice shall be included in all
* copies or substantial portions of the Software.

* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
* IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
* FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
* AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
* LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
* OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
* SOFTWARE.
*/

package api

import (
	"net/url"
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"os"
	"io/ioutil"
	"encoding/json"
    "fmt"
	"github.com/aspose-words-cloud/aspose-words-cloud-go/api/models"
)

// Linger please
var (
	_ context.Context
)

type WordsApiService service

/* WordsApiService Accepts all revisions in document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
 @return RevisionsModificationResponse*/
func (a *WordsApiService) AcceptAllRevisions(ctx context.Context, name models.I, localVarOptionals map[string]interface{}) ( models.RevisionsModificationResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.RevisionsModificationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/revisions/acceptAll"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Appends documents to original document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name Original document name.
 @param documentList DocumentEntryList with a list of documents to append.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return DocumentResponse*/
func (a *WordsApiService) AppendDocument(ctx context.Context, name models.I, documentList models.I, localVarOptionals map[string]interface{}) ( models.DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.DocumentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/appendDocument"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &documentList
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Classifies raw text.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param text Text to classify.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "bestClassesCount" () Number of the best classes to return.
 @return ClassificationResponse*/
func (a *WordsApiService) Classify(ctx context.Context, text models.I, localVarOptionals map[string]interface{}) ( models.ClassificationResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.ClassificationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/classify"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &text
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Classifies document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param documentName The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "bestClassesCount" () Count of the best classes to return.
     @param "taxonomy" () Taxonomy to use for classification return.
 @return ClassificationResponse*/
func (a *WordsApiService) ClassifyDocument(ctx context.Context, documentName models.I, localVarOptionals map[string]interface{}) ( models.ClassificationResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.ClassificationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{documentName}/classify"
	localVarPath = strings.Replace(localVarPath, "{"+"documentName"+"}", fmt.Sprintf("%v", documentName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Compares document with original document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name Original document name.
 @param compareData CompareData with a document to compare.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
 @return DocumentResponse*/
func (a *WordsApiService) CompareDocument(ctx context.Context, name models.I, compareData models.I, localVarOptionals map[string]interface{}) ( models.DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.DocumentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/compareDocument"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &compareData
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Converts document from the request&#39;s content to the specified format .
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param document Converting document
 @param format Format to convert.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" () Original document storage.
     @param "outPath" () Path for saving operation result to the local storage.
     @param "fileNameFieldValue" () This file name will be used when resulting document has dynamic field for document file name {filename}. If it is not set, &quot;sourceFilename&quot; will be used instead. 
     @param "fontsLocation" () Folder in filestorage with custom fonts.
 @return *os.File*/
func (a *WordsApiService) ConvertDocument(ctx context.Context, document , format models.I, localVarOptionals map[string]interface{}) (  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/convert"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	_document := document
	if _document != nil {
		fbs, _ := ioutil.ReadAll(_document)
		_document.Close()
		localFiles[_document.Name()] = fbs
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	return localVarHttpResponse, err
}
/* WordsApiService Copy file
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param destPath Destination file path
 @param srcPath Source file&#39;s path e.g. &#39;/Folder 1/file.ext&#39; or &#39;/Bucket/Folder 1/file.ext&#39;
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "srcStorageName" () Source storage name
     @param "destStorageName" () Destination storage name
     @param "versionId" () File version ID to copy
 @return */
func (a *WordsApiService) CopyFile(ctx context.Context, destPath models.I, srcPath models.I, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/storage/file/copy/{srcPath}"
	localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", fmt.Sprintf("%v", srcPath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Copy folder
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param destPath Destination folder path e.g. &#39;/dst&#39;
 @param srcPath Source folder path e.g. /Folder1
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "srcStorageName" () Source storage name
     @param "destStorageName" () Destination storage name
 @return */
func (a *WordsApiService) CopyFolder(ctx context.Context, destPath models.I, srcPath models.I, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/storage/folder/copy/{srcPath}"
	localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", fmt.Sprintf("%v", srcPath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Creates new document. Document is created with format which is recognized from file extensions. Supported extensions: \&quot;.doc\&quot;, \&quot;.docx\&quot;, \&quot;.docm\&quot;, \&quot;.dot\&quot;, \&quot;.dotm\&quot;, \&quot;.dotx\&quot;, \&quot;.flatopc\&quot;, \&quot;.fopc\&quot;, \&quot;.flatopc_macro\&quot;, \&quot;.fopc_macro\&quot;, \&quot;.flatopc_template\&quot;, \&quot;.fopc_template\&quot;, \&quot;.flatopc_template_macro\&quot;, \&quot;.fopc_template_macro\&quot;, \&quot;.wordml\&quot;, \&quot;.wml\&quot;, \&quot;.rtf\&quot;.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" () Original document storage.
     @param "fileName" () The document name.
     @param "folder" () The document folder.
 @return DocumentResponse*/
func (a *WordsApiService) CreateDocument(ctx context.Context, localVarOptionals map[string]interface{}) ( models.DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.DocumentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/create"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Create the folder
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param path Target folder&#39;s path e.g. Folder1/Folder2/. The folders will be created recursively
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" () Storage name
 @return */
func (a *WordsApiService) CreateFolder(ctx context.Context, path models.I, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/storage/folder/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", path), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Adds new or update existing document property.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param propertyName The property name.
 @param property The property with new value.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return DocumentPropertyResponse*/
func (a *WordsApiService) CreateOrUpdateDocumentProperty(ctx context.Context, name models.I, propertyName models.I, property models.I, localVarOptionals map[string]interface{}) ( models.DocumentPropertyResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.DocumentPropertyResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/documentProperties/{propertyName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", fmt.Sprintf("%v", propertyName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &property
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Resets border properties to default values.             
 &#39;nodePath&#39; should refer to paragraph, cell or row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param nodePath Path to the node with border(node should be paragraph, cell or row).
 @param borderType Border type.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return BorderResponse*/
func (a *WordsApiService) DeleteBorder(ctx context.Context, name models.I, nodePath models.I, borderType models.I, localVarOptionals map[string]interface{}) ( models.BorderResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.BorderResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/borders/{borderType}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"borderType"+"}", fmt.Sprintf("%v", borderType), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Resets borders properties to default values.             
 &#39;nodePath&#39; should refer to paragraph, cell or row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param nodePath Path to the node with borders(node should be paragraph, cell or row).
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return BordersResponse*/
func (a *WordsApiService) DeleteBorders(ctx context.Context, name models.I, nodePath models.I, localVarOptionals map[string]interface{}) ( models.BordersResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.BordersResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/borders"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Removes comment from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param commentIndex The comment index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return */
func (a *WordsApiService) DeleteComment(ctx context.Context, name models.I, commentIndex models.I, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/comments/{commentIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"commentIndex"+"}", fmt.Sprintf("%v", commentIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Deletes document property.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param propertyName The property name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return */
func (a *WordsApiService) DeleteDocumentProperty(ctx context.Context, name models.I, propertyName models.I, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/documentProperties/{propertyName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", fmt.Sprintf("%v", propertyName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Removes drawing object from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param nodePath Path to the node, which contains collection of drawing objects.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return */
func (a *WordsApiService) DeleteDrawingObject(ctx context.Context, name models.I, nodePath models.I, index models.I, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/drawingObjects/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Removes drawing object from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return */
func (a *WordsApiService) DeleteDrawingObjectWithoutNodePath(ctx context.Context, name models.I, index models.I, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/drawingObjects/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Deletes field from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param nodePath Path to the node, which contains collection of fields.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return */
func (a *WordsApiService) DeleteField(ctx context.Context, name models.I, nodePath models.I, index models.I, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/fields/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Deletes field from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return */
func (a *WordsApiService) DeleteFieldWithoutNodePath(ctx context.Context, name models.I, index models.I, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/fields/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Removes fields from section paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param nodePath Path to the node, which contains collection of fields.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return */
func (a *WordsApiService) DeleteFields(ctx context.Context, name models.I, nodePath models.I, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/fields"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Removes fields from section paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return */
func (a *WordsApiService) DeleteFieldsWithoutNodePath(ctx context.Context, name models.I, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/fields"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Delete file
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param path Path of the file including file name and extension e.g. /Folder1/file.ext
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" () Storage name
     @param "versionId" () File version ID to delete
 @return */
func (a *WordsApiService) DeleteFile(ctx context.Context, path models.I, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/storage/file/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", path), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Delete folder
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param path Folder path e.g. /Folder1s
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" () Storage name
     @param "recursive" () Enable to delete folders, subfolders and files
 @return */
func (a *WordsApiService) DeleteFolder(ctx context.Context, path models.I, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/storage/folder/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", path), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Removes footnote from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param nodePath Path to the node, which contains collection of footnotes.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return */
func (a *WordsApiService) DeleteFootnote(ctx context.Context, name models.I, nodePath models.I, index models.I, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/footnotes/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Removes footnote from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return */
func (a *WordsApiService) DeleteFootnoteWithoutNodePath(ctx context.Context, name models.I, index models.I, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/footnotes/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Removes form field from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param nodePath Path to the node that contains collection of formfields.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return */
func (a *WordsApiService) DeleteFormField(ctx context.Context, name models.I, nodePath models.I, index models.I, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/formfields/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Removes form field from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return */
func (a *WordsApiService) DeleteFormFieldWithoutNodePath(ctx context.Context, name models.I, index models.I, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/formfields/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Deletes header/footer from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param sectionPath Path to parent section.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return */
func (a *WordsApiService) DeleteHeaderFooter(ctx context.Context, name models.I, sectionPath models.I, index models.I, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{sectionPath}/headersfooters/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sectionPath"+"}", fmt.Sprintf("%v", sectionPath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Deletes document headers and footers.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param sectionPath Path to parent section.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
     @param "headersFootersTypes" () List of types of headers and footers.
 @return */
func (a *WordsApiService) DeleteHeadersFooters(ctx context.Context, name models.I, sectionPath models.I, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{sectionPath}/headersfooters"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sectionPath"+"}", fmt.Sprintf("%v", sectionPath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Removes macros from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return */
func (a *WordsApiService) DeleteMacros(ctx context.Context, name models.I, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/macros"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Removes OfficeMath object from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param nodePath Path to the node, which contains collection of OfficeMath objects.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return */
func (a *WordsApiService) DeleteOfficeMathObject(ctx context.Context, name models.I, nodePath models.I, index models.I, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/OfficeMathObjects/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Removes OfficeMath object from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return */
func (a *WordsApiService) DeleteOfficeMathObjectWithoutNodePath(ctx context.Context, name models.I, index models.I, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/OfficeMathObjects/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Removes paragraph from section.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The file name.
 @param nodePath Path to the node which contains paragraphs.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return */
func (a *WordsApiService) DeleteParagraph(ctx context.Context, name models.I, nodePath models.I, index models.I, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/paragraphs/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Removes paragraph from section.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The file name.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return */
func (a *WordsApiService) DeleteParagraphWithoutNodePath(ctx context.Context, name models.I, index models.I, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/paragraphs/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Removes run from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param paragraphPath Path to parent paragraph.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return */
func (a *WordsApiService) DeleteRun(ctx context.Context, name models.I, paragraphPath models.I, index models.I, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{paragraphPath}/runs/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"paragraphPath"+"}", fmt.Sprintf("%v", paragraphPath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Removes section from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param sectionIndex Section index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return */
func (a *WordsApiService) DeleteSection(ctx context.Context, name models.I, sectionIndex models.I, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/sections/{sectionIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sectionIndex"+"}", fmt.Sprintf("%v", sectionIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Deletes a table.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param nodePath Path to the node, which contains tables.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return */
func (a *WordsApiService) DeleteTable(ctx context.Context, name models.I, nodePath models.I, index models.I, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/tables/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Deletes a table cell.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param tableRowPath Path to table row.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return */
func (a *WordsApiService) DeleteTableCell(ctx context.Context, name models.I, tableRowPath models.I, index models.I, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{tableRowPath}/cells/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tableRowPath"+"}", fmt.Sprintf("%v", tableRowPath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Deletes a table row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param tablePath Path to table.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return */
func (a *WordsApiService) DeleteTableRow(ctx context.Context, name models.I, tablePath models.I, index models.I, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{tablePath}/rows/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tablePath"+"}", fmt.Sprintf("%v", tablePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Deletes a table.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return */
func (a *WordsApiService) DeleteTableWithoutNodePath(ctx context.Context, name models.I, index models.I, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/tables/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Deletes watermark (for deleting last watermark from the document).
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return DocumentResponse*/
func (a *WordsApiService) DeleteWatermark(ctx context.Context, name models.I, localVarOptionals map[string]interface{}) ( models.DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.DocumentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/watermarks/deleteLast"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Download file
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param path Path of the file including the file name and extension e.g. /folder1/file.ext
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" () Storage name
     @param "versionId" () File version ID to download
 @return *os.File*/
func (a *WordsApiService) DownloadFile(ctx context.Context, path models.I, localVarOptionals map[string]interface{}) (  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/storage/file/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", path), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	return localVarHttpResponse, err
}
/* WordsApiService Executes document mail merge operation.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "data" () Mail merge data
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "withRegions" () With regions flag.
     @param "mailMergeDataFile" () Mail merge data.
     @param "cleanup" () Clean up options.
     @param "useWholeParagraphAsRegion" () Gets or sets a value indicating whether paragraph with TableStart or             TableEnd field should be fully included into mail merge region or particular range between TableStart and TableEnd fields.             The default value is true.
     @param "destFileName" () Result name of the document after the operation. If this parameter is omitted then result of the operation will be saved with autogenerated name.
 @return DocumentResponse*/
func (a *WordsApiService) ExecuteMailMerge(ctx context.Context, name models.I, localVarOptionals map[string]interface{}) ( models.DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.DocumentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/MailMerge"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["data"].(); localVarOk {
		localVarFormParams.Add("Data", parameterToString(localVarTempParam, ""))
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Executes document mail merge online.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param template File with template
 @param data File with mailmerge data
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "withRegions" () With regions flag.
     @param "cleanup" () Clean up options.
     @param "documentFileName" () This file name will be used when resulting document has dynamic field for document file name {filename}. If it is not setted, &quot;template&quot; will be used instead. 
 @return *os.File*/
func (a *WordsApiService) ExecuteMailMergeOnline(ctx context.Context, template , data , localVarOptionals map[string]interface{}) (  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/MailMerge"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	_template := template
	if _template != nil {
		fbs, _ := ioutil.ReadAll(_template)
		_template.Close()
		localFiles[_template.Name()] = fbs
	}
	_data := data
	if _data != nil {
		fbs, _ := ioutil.ReadAll(_data)
		_data.Close()
		localFiles[_data.Name()] = fbs
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	return localVarHttpResponse, err
}
/* WordsApiService Gets the list of fonts, available for document processing.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "fontsLocation" () Folder in filestorage with custom fonts.
 @return AvailableFontsResponse*/
func (a *WordsApiService) GetAvailableFonts(ctx context.Context, localVarOptionals map[string]interface{}) ( models.AvailableFontsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.AvailableFontsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/fonts/available"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads document bookmark data by its name.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param bookmarkName The bookmark name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return BookmarkResponse*/
func (a *WordsApiService) GetBookmarkByName(ctx context.Context, name models.I, bookmarkName models.I, localVarOptionals map[string]interface{}) ( models.BookmarkResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.BookmarkResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/bookmarks/{bookmarkName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bookmarkName"+"}", fmt.Sprintf("%v", bookmarkName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads document bookmarks common info.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return BookmarksResponse*/
func (a *WordsApiService) GetBookmarks(ctx context.Context, name models.I, localVarOptionals map[string]interface{}) ( models.BookmarksResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.BookmarksResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/bookmarks"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a border.
 &#39;nodePath&#39; should refer to paragraph, cell or row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param nodePath Path to the node with border(node should be paragraph, cell or row).
 @param borderType Border type.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return BorderResponse*/
func (a *WordsApiService) GetBorder(ctx context.Context, name models.I, nodePath models.I, borderType models.I, localVarOptionals map[string]interface{}) ( models.BorderResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.BorderResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/borders/{borderType}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"borderType"+"}", fmt.Sprintf("%v", borderType), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a collection of borders.
 &#39;nodePath&#39; should refer to paragraph, cell or row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param nodePath Path to the node with borders (node should be paragraph, cell or row).
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return BordersResponse*/
func (a *WordsApiService) GetBorders(ctx context.Context, name models.I, nodePath models.I, localVarOptionals map[string]interface{}) ( models.BordersResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.BordersResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/borders"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Gets comment from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param commentIndex The comment index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return CommentResponse*/
func (a *WordsApiService) GetComment(ctx context.Context, name models.I, commentIndex models.I, localVarOptionals map[string]interface{}) ( models.CommentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.CommentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/comments/{commentIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"commentIndex"+"}", fmt.Sprintf("%v", commentIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Gets comments from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return CommentsResponse*/
func (a *WordsApiService) GetComments(ctx context.Context, name models.I, localVarOptionals map[string]interface{}) ( models.CommentsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.CommentsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/comments"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads document common info.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param documentName The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return DocumentResponse*/
func (a *WordsApiService) GetDocument(ctx context.Context, documentName models.I, localVarOptionals map[string]interface{}) ( models.DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.DocumentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{documentName}"
	localVarPath = strings.Replace(localVarPath, "{"+"documentName"+"}", fmt.Sprintf("%v", documentName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads document drawing object common info by its index or convert to format specified.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param nodePath Path to the node, which contains collection of drawing objects.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return DrawingObjectResponse*/
func (a *WordsApiService) GetDocumentDrawingObjectByIndex(ctx context.Context, name models.I, nodePath models.I, index models.I, localVarOptionals map[string]interface{}) ( models.DrawingObjectResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.DrawingObjectResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/drawingObjects/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads document drawing object common info by its index or convert to format specified.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return DrawingObjectResponse*/
func (a *WordsApiService) GetDocumentDrawingObjectByIndexWithoutNodePath(ctx context.Context, name models.I, index models.I, localVarOptionals map[string]interface{}) ( models.DrawingObjectResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.DrawingObjectResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/drawingObjects/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads drawing object image data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param nodePath Path to the node, which contains collection of drawing objects.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return *os.File*/
func (a *WordsApiService) GetDocumentDrawingObjectImageData(ctx context.Context, name models.I, nodePath models.I, index models.I, localVarOptionals map[string]interface{}) (  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/drawingObjects/{index}/imageData"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	return localVarHttpResponse, err
}
/* WordsApiService Reads drawing object image data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return *os.File*/
func (a *WordsApiService) GetDocumentDrawingObjectImageDataWithoutNodePath(ctx context.Context, name models.I, index models.I, localVarOptionals map[string]interface{}) (  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/drawingObjects/{index}/imageData"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	return localVarHttpResponse, err
}
/* WordsApiService Gets drawing object OLE data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param nodePath Path to the node, which contains collection of drawing objects.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return *os.File*/
func (a *WordsApiService) GetDocumentDrawingObjectOleData(ctx context.Context, name models.I, nodePath models.I, index models.I, localVarOptionals map[string]interface{}) (  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/drawingObjects/{index}/oleData"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	return localVarHttpResponse, err
}
/* WordsApiService Gets drawing object OLE data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return *os.File*/
func (a *WordsApiService) GetDocumentDrawingObjectOleDataWithoutNodePath(ctx context.Context, name models.I, index models.I, localVarOptionals map[string]interface{}) (  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/drawingObjects/{index}/oleData"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	return localVarHttpResponse, err
}
/* WordsApiService Reads document drawing objects common info.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param nodePath Path to the node, which contains collection of drawing objects.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return DrawingObjectsResponse*/
func (a *WordsApiService) GetDocumentDrawingObjects(ctx context.Context, name models.I, nodePath models.I, localVarOptionals map[string]interface{}) ( models.DrawingObjectsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.DrawingObjectsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/drawingObjects"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads document drawing objects common info.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return DrawingObjectsResponse*/
func (a *WordsApiService) GetDocumentDrawingObjectsWithoutNodePath(ctx context.Context, name models.I, localVarOptionals map[string]interface{}) ( models.DrawingObjectsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.DrawingObjectsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/drawingObjects"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads document field names.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "useNonMergeFields" () If true, result includes &quot;mustache&quot; field names.
 @return FieldNamesResponse*/
func (a *WordsApiService) GetDocumentFieldNames(ctx context.Context, name models.I, localVarOptionals map[string]interface{}) ( models.FieldNamesResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.FieldNamesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/mailMerge/FieldNames"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads document field names.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param template File with template
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "useNonMergeFields" () Use non merge fields or not.
 @return FieldNamesResponse*/
func (a *WordsApiService) GetDocumentFieldNamesOnline(ctx context.Context, template , localVarOptionals map[string]interface{}) ( models.FieldNamesResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.FieldNamesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/mailMerge/FieldNames"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	_template := template
	if _template != nil {
		fbs, _ := ioutil.ReadAll(_template)
		_template.Close()
		localFiles[_template.Name()] = fbs
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads document hyperlink by its index.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param hyperlinkIndex The hyperlink index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return HyperlinkResponse*/
func (a *WordsApiService) GetDocumentHyperlinkByIndex(ctx context.Context, name models.I, hyperlinkIndex models.I, localVarOptionals map[string]interface{}) ( models.HyperlinkResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.HyperlinkResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/hyperlinks/{hyperlinkIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"hyperlinkIndex"+"}", fmt.Sprintf("%v", hyperlinkIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads document hyperlinks common info.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return HyperlinksResponse*/
func (a *WordsApiService) GetDocumentHyperlinks(ctx context.Context, name models.I, localVarOptionals map[string]interface{}) ( models.HyperlinksResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.HyperlinksResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/hyperlinks"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads document properties info.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document&#39;s name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return DocumentPropertiesResponse*/
func (a *WordsApiService) GetDocumentProperties(ctx context.Context, name models.I, localVarOptionals map[string]interface{}) ( models.DocumentPropertiesResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.DocumentPropertiesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/documentProperties"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads document property info by the property name.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param propertyName The property name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return DocumentPropertyResponse*/
func (a *WordsApiService) GetDocumentProperty(ctx context.Context, name models.I, propertyName models.I, localVarOptionals map[string]interface{}) ( models.DocumentPropertyResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.DocumentPropertyResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/documentProperties/{propertyName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", fmt.Sprintf("%v", propertyName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads document protection common info.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return ProtectionDataResponse*/
func (a *WordsApiService) GetDocumentProtection(ctx context.Context, name models.I, localVarOptionals map[string]interface{}) ( models.ProtectionDataResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.ProtectionDataResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/protection"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads document statistics.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "includeComments" () Support including/excluding comments from the WordCount. Default value is &quot;false&quot;.
     @param "includeFootnotes" () Support including/excluding footnotes from the WordCount. Default value is &quot;false&quot;.
     @param "includeTextInShapes" () Support including/excluding shape&#39;s text from the WordCount. Default value is &quot;false&quot;.
 @return StatDataResponse*/
func (a *WordsApiService) GetDocumentStatistics(ctx context.Context, name models.I, localVarOptionals map[string]interface{}) ( models.StatDataResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.StatDataResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/statistics"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Exports the document into the specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param format The destination format.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "outPath" () Path to save the result.
     @param "fontsLocation" () Folder in filestorage with custom fonts.
 @return *os.File*/
func (a *WordsApiService) GetDocumentWithFormat(ctx context.Context, name models.I, format models.I, localVarOptionals map[string]interface{}) (  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	return localVarHttpResponse, err
}
/* WordsApiService Gets field from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param nodePath Path to the node, which contains collection of fields.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return FieldResponse*/
func (a *WordsApiService) GetField(ctx context.Context, name models.I, nodePath models.I, index models.I, localVarOptionals map[string]interface{}) ( models.FieldResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.FieldResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/fields/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Gets field from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return FieldResponse*/
func (a *WordsApiService) GetFieldWithoutNodePath(ctx context.Context, name models.I, index models.I, localVarOptionals map[string]interface{}) ( models.FieldResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.FieldResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/fields/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Get fields from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param nodePath Path to the node, which contains collection of fields.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return FieldsResponse*/
func (a *WordsApiService) GetFields(ctx context.Context, name models.I, nodePath models.I, localVarOptionals map[string]interface{}) ( models.FieldsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.FieldsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/fields"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Get fields from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return FieldsResponse*/
func (a *WordsApiService) GetFieldsWithoutNodePath(ctx context.Context, name models.I, localVarOptionals map[string]interface{}) ( models.FieldsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.FieldsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/fields"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Get all files and folders within a folder
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param path Folder path e.g. /Folder1
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" () Storage name
 @return FilesList*/
func (a *WordsApiService) GetFilesList(ctx context.Context, path models.I, localVarOptionals map[string]interface{}) ( models.FilesList,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.FilesList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/storage/folder/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", path), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads footnote by index.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param nodePath Path to the node, which contains collection of footnotes.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return FootnoteResponse*/
func (a *WordsApiService) GetFootnote(ctx context.Context, name models.I, nodePath models.I, index models.I, localVarOptionals map[string]interface{}) ( models.FootnoteResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.FootnoteResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/footnotes/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads footnote by index.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return FootnoteResponse*/
func (a *WordsApiService) GetFootnoteWithoutNodePath(ctx context.Context, name models.I, index models.I, localVarOptionals map[string]interface{}) ( models.FootnoteResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.FootnoteResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/footnotes/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Gets footnotes from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param nodePath Path to the node, which contains collection of footnotes.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return FootnotesResponse*/
func (a *WordsApiService) GetFootnotes(ctx context.Context, name models.I, nodePath models.I, localVarOptionals map[string]interface{}) ( models.FootnotesResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.FootnotesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/footnotes"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Gets footnotes from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return FootnotesResponse*/
func (a *WordsApiService) GetFootnotesWithoutNodePath(ctx context.Context, name models.I, localVarOptionals map[string]interface{}) ( models.FootnotesResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.FootnotesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/footnotes"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns representation of an one of the form field.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param nodePath Path to the node that contains collection of formfields.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return FormFieldResponse*/
func (a *WordsApiService) GetFormField(ctx context.Context, name models.I, nodePath models.I, index models.I, localVarOptionals map[string]interface{}) ( models.FormFieldResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.FormFieldResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/formfields/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns representation of an one of the form field.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return FormFieldResponse*/
func (a *WordsApiService) GetFormFieldWithoutNodePath(ctx context.Context, name models.I, index models.I, localVarOptionals map[string]interface{}) ( models.FormFieldResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.FormFieldResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/formfields/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Gets form fields from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param nodePath Path to the node containing collection of form fields.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return FormFieldsResponse*/
func (a *WordsApiService) GetFormFields(ctx context.Context, name models.I, nodePath models.I, localVarOptionals map[string]interface{}) ( models.FormFieldsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.FormFieldsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/formfields"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Gets form fields from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return FormFieldsResponse*/
func (a *WordsApiService) GetFormFieldsWithoutNodePath(ctx context.Context, name models.I, localVarOptionals map[string]interface{}) ( models.FormFieldsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.FormFieldsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/formfields"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a header/footer from the document by index.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param headerFooterIndex Header/footer index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "filterByType" () List of types of headers and footers.
 @return HeaderFooterResponse*/
func (a *WordsApiService) GetHeaderFooter(ctx context.Context, name models.I, headerFooterIndex models.I, localVarOptionals map[string]interface{}) ( models.HeaderFooterResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.HeaderFooterResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/headersfooters/{headerFooterIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"headerFooterIndex"+"}", fmt.Sprintf("%v", headerFooterIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a header/footer from the document section.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param headerFooterIndex Header/footer index.
 @param sectionIndex Section index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "filterByType" () List of types of headers and footers.
 @return HeaderFooterResponse*/
func (a *WordsApiService) GetHeaderFooterOfSection(ctx context.Context, name models.I, headerFooterIndex models.I, sectionIndex models.I, localVarOptionals map[string]interface{}) ( models.HeaderFooterResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.HeaderFooterResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/sections/{sectionIndex}/headersfooters/{headerFooterIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"headerFooterIndex"+"}", fmt.Sprintf("%v", headerFooterIndex), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sectionIndex"+"}", fmt.Sprintf("%v", sectionIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a list of header/footers from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param sectionPath Path to parent section.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "filterByType" () List of types of headers and footers.
 @return HeaderFootersResponse*/
func (a *WordsApiService) GetHeaderFooters(ctx context.Context, name models.I, sectionPath models.I, localVarOptionals map[string]interface{}) ( models.HeaderFootersResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.HeaderFootersResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{sectionPath}/headersfooters"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sectionPath"+"}", fmt.Sprintf("%v", sectionPath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads OfficeMath object by index.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param nodePath Path to the node, which contains collection of OfficeMath objects.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return OfficeMathObjectResponse*/
func (a *WordsApiService) GetOfficeMathObject(ctx context.Context, name models.I, nodePath models.I, index models.I, localVarOptionals map[string]interface{}) ( models.OfficeMathObjectResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.OfficeMathObjectResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/OfficeMathObjects/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Reads OfficeMath object by index.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return OfficeMathObjectResponse*/
func (a *WordsApiService) GetOfficeMathObjectWithoutNodePath(ctx context.Context, name models.I, index models.I, localVarOptionals map[string]interface{}) ( models.OfficeMathObjectResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.OfficeMathObjectResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/OfficeMathObjects/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Gets OfficeMath objects from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param nodePath Path to the node, which contains collection of OfficeMath objects.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return OfficeMathObjectsResponse*/
func (a *WordsApiService) GetOfficeMathObjects(ctx context.Context, name models.I, nodePath models.I, localVarOptionals map[string]interface{}) ( models.OfficeMathObjectsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.OfficeMathObjectsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/OfficeMathObjects"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Gets OfficeMath objects from document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return OfficeMathObjectsResponse*/
func (a *WordsApiService) GetOfficeMathObjectsWithoutNodePath(ctx context.Context, name models.I, localVarOptionals map[string]interface{}) ( models.OfficeMathObjectsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.OfficeMathObjectsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/OfficeMathObjects"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService This resource represents one of the paragraphs contained in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param nodePath Path to the node which contains paragraphs.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return ParagraphResponse*/
func (a *WordsApiService) GetParagraph(ctx context.Context, name models.I, nodePath models.I, index models.I, localVarOptionals map[string]interface{}) ( models.ParagraphResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.ParagraphResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/paragraphs/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Represents all the formatting for a paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param nodePath Path to the node which contains paragraphs.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return ParagraphFormatResponse*/
func (a *WordsApiService) GetParagraphFormat(ctx context.Context, name models.I, nodePath models.I, index models.I, localVarOptionals map[string]interface{}) ( models.ParagraphFormatResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.ParagraphFormatResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/paragraphs/{index}/format"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Represents all the formatting for a paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return ParagraphFormatResponse*/
func (a *WordsApiService) GetParagraphFormatWithoutNodePath(ctx context.Context, name models.I, index models.I, localVarOptionals map[string]interface{}) ( models.ParagraphFormatResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.ParagraphFormatResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/paragraphs/{index}/format"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService This resource represents one of the paragraphs contained in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return ParagraphResponse*/
func (a *WordsApiService) GetParagraphWithoutNodePath(ctx context.Context, name models.I, index models.I, localVarOptionals map[string]interface{}) ( models.ParagraphResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.ParagraphResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/paragraphs/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a list of paragraphs that are contained in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param nodePath Path to the node which contains paragraphs.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return ParagraphLinkCollectionResponse*/
func (a *WordsApiService) GetParagraphs(ctx context.Context, name models.I, nodePath models.I, localVarOptionals map[string]interface{}) ( models.ParagraphLinkCollectionResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.ParagraphLinkCollectionResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/paragraphs"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a list of paragraphs that are contained in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return ParagraphLinkCollectionResponse*/
func (a *WordsApiService) GetParagraphsWithoutNodePath(ctx context.Context, name models.I, localVarOptionals map[string]interface{}) ( models.ParagraphLinkCollectionResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.ParagraphLinkCollectionResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/paragraphs"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Gets the text from the range.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document.
 @param rangeStartIdentifier The range start identifier. Identifier is the value of the &quot;nodeId&quot; field, which every document node has, extended with the prefix &quot;id&quot;. It looks like &quot;id0.0.7&quot;. Also values like &quot;image5&quot; and &quot;table3&quot; can be used as an identifier for images and tables, where the number is an index of the image/table.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "rangeEndIdentifier" () The range end identifier.
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return RangeTextResponse*/
func (a *WordsApiService) GetRangeText(ctx context.Context, name models.I, rangeStartIdentifier models.I, localVarOptionals map[string]interface{}) ( models.RangeTextResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.RangeTextResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/range/{rangeStartIdentifier}/{rangeEndIdentifier}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rangeStartIdentifier"+"}", fmt.Sprintf("%v", rangeStartIdentifier), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rangeEndIdentifier"+"}", fmt.Sprintf("%v", localVarOptionals["rangeEndIdentifier"]), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService This resource represents run of text contained in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param paragraphPath Path to parent paragraph.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return RunResponse*/
func (a *WordsApiService) GetRun(ctx context.Context, name models.I, paragraphPath models.I, index models.I, localVarOptionals map[string]interface{}) ( models.RunResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.RunResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{paragraphPath}/runs/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"paragraphPath"+"}", fmt.Sprintf("%v", paragraphPath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService This resource represents font of run.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param paragraphPath Path to parent paragraph.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return FontResponse*/
func (a *WordsApiService) GetRunFont(ctx context.Context, name models.I, paragraphPath models.I, index models.I, localVarOptionals map[string]interface{}) ( models.FontResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.FontResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{paragraphPath}/runs/{index}/font"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"paragraphPath"+"}", fmt.Sprintf("%v", paragraphPath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService This resource represents collection of runs in the paragraph.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param paragraphPath Path to parent paragraph.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return RunsResponse*/
func (a *WordsApiService) GetRuns(ctx context.Context, name models.I, paragraphPath models.I, localVarOptionals map[string]interface{}) ( models.RunsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.RunsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{paragraphPath}/runs"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"paragraphPath"+"}", fmt.Sprintf("%v", paragraphPath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Gets document section by index.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param sectionIndex Section index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return SectionResponse*/
func (a *WordsApiService) GetSection(ctx context.Context, name models.I, sectionIndex models.I, localVarOptionals map[string]interface{}) ( models.SectionResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.SectionResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/sections/{sectionIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sectionIndex"+"}", fmt.Sprintf("%v", sectionIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Gets page setup of section.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param sectionIndex Section index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return SectionPageSetupResponse*/
func (a *WordsApiService) GetSectionPageSetup(ctx context.Context, name models.I, sectionIndex models.I, localVarOptionals map[string]interface{}) ( models.SectionPageSetupResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.SectionPageSetupResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/sections/{sectionIndex}/pageSetup"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sectionIndex"+"}", fmt.Sprintf("%v", sectionIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a list of sections that are contained in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return SectionLinkCollectionResponse*/
func (a *WordsApiService) GetSections(ctx context.Context, name models.I, localVarOptionals map[string]interface{}) ( models.SectionLinkCollectionResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.SectionLinkCollectionResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/sections"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a table.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param nodePath Path to the node, which contains tables.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return TableResponse*/
func (a *WordsApiService) GetTable(ctx context.Context, name models.I, nodePath models.I, index models.I, localVarOptionals map[string]interface{}) ( models.TableResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.TableResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/tables/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a table cell.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param tableRowPath Path to table row.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return TableCellResponse*/
func (a *WordsApiService) GetTableCell(ctx context.Context, name models.I, tableRowPath models.I, index models.I, localVarOptionals map[string]interface{}) ( models.TableCellResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.TableCellResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{tableRowPath}/cells/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tableRowPath"+"}", fmt.Sprintf("%v", tableRowPath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a table cell format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param tableRowPath Path to table row.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return TableCellFormatResponse*/
func (a *WordsApiService) GetTableCellFormat(ctx context.Context, name models.I, tableRowPath models.I, index models.I, localVarOptionals map[string]interface{}) ( models.TableCellFormatResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.TableCellFormatResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{tableRowPath}/cells/{index}/cellformat"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tableRowPath"+"}", fmt.Sprintf("%v", tableRowPath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a table properties.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param nodePath Path to the node, which contains tables.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return TablePropertiesResponse*/
func (a *WordsApiService) GetTableProperties(ctx context.Context, name models.I, nodePath models.I, index models.I, localVarOptionals map[string]interface{}) ( models.TablePropertiesResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.TablePropertiesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/tables/{index}/properties"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a table properties.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return TablePropertiesResponse*/
func (a *WordsApiService) GetTablePropertiesWithoutNodePath(ctx context.Context, name models.I, index models.I, localVarOptionals map[string]interface{}) ( models.TablePropertiesResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.TablePropertiesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/tables/{index}/properties"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a table row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param tablePath Path to table.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return TableRowResponse*/
func (a *WordsApiService) GetTableRow(ctx context.Context, name models.I, tablePath models.I, index models.I, localVarOptionals map[string]interface{}) ( models.TableRowResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.TableRowResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{tablePath}/rows/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tablePath"+"}", fmt.Sprintf("%v", tablePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a table row format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param tablePath Path to table.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return TableRowFormatResponse*/
func (a *WordsApiService) GetTableRowFormat(ctx context.Context, name models.I, tablePath models.I, index models.I, localVarOptionals map[string]interface{}) ( models.TableRowFormatResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.TableRowFormatResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{tablePath}/rows/{index}/rowformat"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tablePath"+"}", fmt.Sprintf("%v", tablePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a table.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return TableResponse*/
func (a *WordsApiService) GetTableWithoutNodePath(ctx context.Context, name models.I, index models.I, localVarOptionals map[string]interface{}) ( models.TableResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.TableResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/tables/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a list of tables that are contained in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param nodePath Path to the node, which contains tables.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return TableLinkCollectionResponse*/
func (a *WordsApiService) GetTables(ctx context.Context, name models.I, nodePath models.I, localVarOptionals map[string]interface{}) ( models.TableLinkCollectionResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.TableLinkCollectionResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/tables"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Returns a list of tables that are contained in the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return TableLinkCollectionResponse*/
func (a *WordsApiService) GetTablesWithoutNodePath(ctx context.Context, name models.I, localVarOptionals map[string]interface{}) ( models.TableLinkCollectionResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.TableLinkCollectionResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/tables"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds comment to document, returns inserted comment data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param comment The comment data.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return CommentResponse*/
func (a *WordsApiService) InsertComment(ctx context.Context, name models.I, comment models.I, localVarOptionals map[string]interface{}) ( models.CommentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.CommentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/comments"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &comment
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds drawing object to document, returns added  drawing object&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param drawingObject Drawing object parameters
 @param imageFile File with image
 @param nodePath Path to the node, which contains collection of drawing objects.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return DrawingObjectResponse*/
func (a *WordsApiService) InsertDrawingObject(ctx context.Context, name models.I, drawingObject models.I, imageFile , nodePath models.I, localVarOptionals map[string]interface{}) ( models.DrawingObjectResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.DrawingObjectResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/drawingObjects"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarFormParams.Add("DrawingObject", parameterToString(drawingObject, ""))
	_imageFile := imageFile
	if _imageFile != nil {
		fbs, _ := ioutil.ReadAll(_imageFile)
		_imageFile.Close()
		localFiles[_imageFile.Name()] = fbs
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds drawing object to document, returns added  drawing object&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param drawingObject Drawing object parameters
 @param imageFile File with image
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return DrawingObjectResponse*/
func (a *WordsApiService) InsertDrawingObjectWithoutNodePath(ctx context.Context, name models.I, drawingObject models.I, imageFile , localVarOptionals map[string]interface{}) ( models.DrawingObjectResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.DrawingObjectResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/drawingObjects"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarFormParams.Add("DrawingObject", parameterToString(drawingObject, ""))
	_imageFile := imageFile
	if _imageFile != nil {
		fbs, _ := ioutil.ReadAll(_imageFile)
		_imageFile.Close()
		localFiles[_imageFile.Name()] = fbs
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds field to document, returns inserted field&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param field Field data.
 @param nodePath Path to the node, which contains collection of fields.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
     @param "insertBeforeNode" () Field will be inserted before node with id&#x3D;&quot;nodeId&quot;.
 @return FieldResponse*/
func (a *WordsApiService) InsertField(ctx context.Context, name models.I, field models.I, nodePath models.I, localVarOptionals map[string]interface{}) ( models.FieldResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.FieldResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/fields"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &field
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds field to document, returns inserted field&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param field Field data.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
     @param "insertBeforeNode" () Field will be inserted before node with id&#x3D;&quot;nodeId&quot;.
 @return FieldResponse*/
func (a *WordsApiService) InsertFieldWithoutNodePath(ctx context.Context, name models.I, field models.I, localVarOptionals map[string]interface{}) ( models.FieldResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.FieldResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/fields"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &field
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds footnote to document, returns added footnote&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param footnoteDto Footnote data.
 @param nodePath Path to the node, which contains collection of footnotes.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return FootnoteResponse*/
func (a *WordsApiService) InsertFootnote(ctx context.Context, name models.I, footnoteDto models.I, nodePath models.I, localVarOptionals map[string]interface{}) ( models.FootnoteResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.FootnoteResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/footnotes"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &footnoteDto
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds footnote to document, returns added footnote&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param footnoteDto Footnote data.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return FootnoteResponse*/
func (a *WordsApiService) InsertFootnoteWithoutNodePath(ctx context.Context, name models.I, footnoteDto models.I, localVarOptionals map[string]interface{}) ( models.FootnoteResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.FootnoteResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/footnotes"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &footnoteDto
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds form field to paragraph, returns added form field&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param formField From field data.
 @param nodePath Path to the node that contains collection of formfields.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
     @param "insertBeforeNode" () Form field will be inserted before node with index.
 @return FormFieldResponse*/
func (a *WordsApiService) InsertFormField(ctx context.Context, name models.I, formField models.I, nodePath models.I, localVarOptionals map[string]interface{}) ( models.FormFieldResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.FormFieldResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/formfields"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &formField
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds form field to paragraph, returns added form field&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param formField From field data.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
     @param "insertBeforeNode" () Form field will be inserted before node with index.
 @return FormFieldResponse*/
func (a *WordsApiService) InsertFormFieldWithoutNodePath(ctx context.Context, name models.I, formField models.I, localVarOptionals map[string]interface{}) ( models.FormFieldResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.FormFieldResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/formfields"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &formField
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Inserts to document header or footer.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param headerFooterType Type of header/footer.
 @param sectionPath Path to parent section.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return HeaderFooterResponse*/
func (a *WordsApiService) InsertHeaderFooter(ctx context.Context, name models.I, headerFooterType models.I, sectionPath models.I, localVarOptionals map[string]interface{}) ( models.HeaderFooterResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.HeaderFooterResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{sectionPath}/headersfooters"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sectionPath"+"}", fmt.Sprintf("%v", sectionPath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &headerFooterType
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Inserts document page numbers.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name A document name.
 @param pageNumber PageNumber with the page numbers settings.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return DocumentResponse*/
func (a *WordsApiService) InsertPageNumbers(ctx context.Context, name models.I, pageNumber models.I, localVarOptionals map[string]interface{}) ( models.DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.DocumentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/PageNumbers"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &pageNumber
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds paragraph to document, returns added paragraph&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param paragraph Paragraph data.
 @param nodePath Path to the node which contains paragraphs.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
     @param "insertBeforeNode" () Paragraph will be inserted before node with index.
 @return ParagraphResponse*/
func (a *WordsApiService) InsertParagraph(ctx context.Context, name models.I, paragraph models.I, nodePath models.I, localVarOptionals map[string]interface{}) ( models.ParagraphResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.ParagraphResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/paragraphs"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &paragraph
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds run to document, returns added paragraph&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param paragraphPath Path to parent paragraph.
 @param run Run data.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
     @param "insertBeforeNode" () Paragraph will be inserted before node with index.
 @return RunResponse*/
func (a *WordsApiService) InsertRun(ctx context.Context, name models.I, paragraphPath models.I, run models.I, localVarOptionals map[string]interface{}) ( models.RunResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.RunResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{paragraphPath}/runs"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"paragraphPath"+"}", fmt.Sprintf("%v", paragraphPath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &run
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds table to document, returns added table&#39;s data.             
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param table Table parameters/.
 @param nodePath Path to the node, which contains tables.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return TableResponse*/
func (a *WordsApiService) InsertTable(ctx context.Context, name models.I, table models.I, nodePath models.I, localVarOptionals map[string]interface{}) ( models.TableResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.TableResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/tables"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &table
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds table cell to table, returns added cell&#39;s data.             
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param cell Table cell parameters/.
 @param tableRowPath Path to table row.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return TableCellResponse*/
func (a *WordsApiService) InsertTableCell(ctx context.Context, name models.I, cell models.I, tableRowPath models.I, localVarOptionals map[string]interface{}) ( models.TableCellResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.TableCellResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{tableRowPath}/cells"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tableRowPath"+"}", fmt.Sprintf("%v", tableRowPath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &cell
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds table row to table, returns added row&#39;s data.             
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param row Table row parameters/.
 @param tablePath Path to table.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return TableRowResponse*/
func (a *WordsApiService) InsertTableRow(ctx context.Context, name models.I, row models.I, tablePath models.I, localVarOptionals map[string]interface{}) ( models.TableRowResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.TableRowResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{tablePath}/rows"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tablePath"+"}", fmt.Sprintf("%v", tablePath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &row
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Adds table to document, returns added table&#39;s data.             
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param table Table parameters/.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return TableResponse*/
func (a *WordsApiService) InsertTableWithoutNodePath(ctx context.Context, name models.I, table models.I, localVarOptionals map[string]interface{}) ( models.TableResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.TableResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/tables"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &table
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Inserts document watermark image.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "imageFile" () File with image
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
     @param "rotationAngle" () The watermark rotation angle.
     @param "image" () The image file server full name. If the name is empty the image is expected in request content.
 @return DocumentResponse*/
func (a *WordsApiService) InsertWatermarkImage(ctx context.Context, name models.I, localVarOptionals map[string]interface{}) ( models.DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.DocumentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/watermarks/images"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var imageFile (*os.File)
	if localVarTempParam, localVarOk := localVarOptionals["imageFile"].(); localVarOk {
		imageFile = localVarTempParam
	}
	_imageFile := imageFile
	if _imageFile != nil {
		fbs, _ := ioutil.ReadAll(_imageFile)
		_imageFile.Close()
		localFiles[_imageFile.Name()] = fbs
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Inserts document watermark text.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param watermarkText WatermarkText with the watermark data.             
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return DocumentResponse*/
func (a *WordsApiService) InsertWatermarkText(ctx context.Context, name models.I, watermarkText models.I, localVarOptionals map[string]interface{}) ( models.DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.DocumentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/watermarks/texts"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &watermarkText
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Loads new document from web into the file with any supported format of data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param data Parameters of loading.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" () Original document storage.
 @return SaveResponse*/
func (a *WordsApiService) LoadWebDocument(ctx context.Context, data models.I, localVarOptionals map[string]interface{}) ( models.SaveResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.SaveResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/loadWebDocument"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &data
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Move file
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param destPath Destination file path e.g. &#39;/dest.ext&#39;
 @param srcPath Source file&#39;s path e.g. &#39;/Folder 1/file.ext&#39; or &#39;/Bucket/Folder 1/file.ext&#39;
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "srcStorageName" () Source storage name
     @param "destStorageName" () Destination storage name
     @param "versionId" () File version ID to move
 @return */
func (a *WordsApiService) MoveFile(ctx context.Context, destPath models.I, srcPath models.I, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/storage/file/move/{srcPath}"
	localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", fmt.Sprintf("%v", srcPath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Move folder
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param destPath Destination folder path to move to e.g &#39;/dst&#39;
 @param srcPath Source folder path e.g. /Folder1
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "srcStorageName" () Source storage name
     @param "destStorageName" () Destination storage name
 @return */
func (a *WordsApiService) MoveFolder(ctx context.Context, destPath models.I, srcPath models.I, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/storage/folder/move/{srcPath}"
	localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", fmt.Sprintf("%v", srcPath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Protects document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param protectionRequest ProtectionRequest with protection settings.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
 @return ProtectionDataResponse*/
func (a *WordsApiService) ProtectDocument(ctx context.Context, name models.I, protectionRequest models.I, localVarOptionals map[string]interface{}) ( models.ProtectionDataResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.ProtectionDataResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/protection"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &protectionRequest
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Rejects all revisions in document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
 @return RevisionsModificationResponse*/
func (a *WordsApiService) RejectAllRevisions(ctx context.Context, name models.I, localVarOptionals map[string]interface{}) ( models.RevisionsModificationResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.RevisionsModificationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/revisions/rejectAll"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Removes the range from the document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document.
 @param rangeStartIdentifier The range start identifier. Identifier is the value of the &quot;nodeId&quot; field, which every document node has, extended with the prefix &quot;id&quot;. It looks like &quot;id0.0.7&quot;. Also values like &quot;image5&quot; and &quot;table3&quot; can be used as an identifier for images and tables, where the number is an index of the image/table.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "rangeEndIdentifier" () The range end identifier.
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
 @return DocumentResponse*/
func (a *WordsApiService) RemoveRange(ctx context.Context, name models.I, rangeStartIdentifier models.I, localVarOptionals map[string]interface{}) ( models.DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.DocumentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/range/{rangeStartIdentifier}/{rangeEndIdentifier}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rangeStartIdentifier"+"}", fmt.Sprintf("%v", rangeStartIdentifier), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rangeEndIdentifier"+"}", fmt.Sprintf("%v", localVarOptionals["rangeEndIdentifier"]), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Renders drawing object to specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param format The destination format.
 @param nodePath Path to the node, which contains drawing objects.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "fontsLocation" () Folder in filestorage with custom fonts.
 @return *os.File*/
func (a *WordsApiService) RenderDrawingObject(ctx context.Context, name models.I, format models.I, nodePath models.I, index models.I, localVarOptionals map[string]interface{}) (  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/drawingObjects/{index}/render"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	return localVarHttpResponse, err
}
/* WordsApiService Renders drawing object to specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param format The destination format.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "fontsLocation" () Folder in filestorage with custom fonts.
 @return *os.File*/
func (a *WordsApiService) RenderDrawingObjectWithoutNodePath(ctx context.Context, name models.I, format models.I, index models.I, localVarOptionals map[string]interface{}) (  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/drawingObjects/{index}/render"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	return localVarHttpResponse, err
}
/* WordsApiService Renders math object to specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param format The destination format.
 @param nodePath Path to the node, which contains office math objects.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "fontsLocation" () Folder in filestorage with custom fonts.
 @return *os.File*/
func (a *WordsApiService) RenderMathObject(ctx context.Context, name models.I, format models.I, nodePath models.I, index models.I, localVarOptionals map[string]interface{}) (  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/OfficeMathObjects/{index}/render"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	return localVarHttpResponse, err
}
/* WordsApiService Renders math object to specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param format The destination format.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "fontsLocation" () Folder in filestorage with custom fonts.
 @return *os.File*/
func (a *WordsApiService) RenderMathObjectWithoutNodePath(ctx context.Context, name models.I, format models.I, index models.I, localVarOptionals map[string]interface{}) (  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/OfficeMathObjects/{index}/render"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	return localVarHttpResponse, err
}
/* WordsApiService Renders page to specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param pageIndex Comment index.
 @param format The destination format.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "fontsLocation" () Folder in filestorage with custom fonts.
 @return *os.File*/
func (a *WordsApiService) RenderPage(ctx context.Context, name models.I, pageIndex models.I, format models.I, localVarOptionals map[string]interface{}) (  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/pages/{pageIndex}/render"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageIndex"+"}", fmt.Sprintf("%v", pageIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	return localVarHttpResponse, err
}
/* WordsApiService Renders paragraph to specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param format The destination format.
 @param nodePath Path to the node, which contains paragraphs.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "fontsLocation" () Folder in filestorage with custom fonts.
 @return *os.File*/
func (a *WordsApiService) RenderParagraph(ctx context.Context, name models.I, format models.I, nodePath models.I, index models.I, localVarOptionals map[string]interface{}) (  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/paragraphs/{index}/render"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	return localVarHttpResponse, err
}
/* WordsApiService Renders paragraph to specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param format The destination format.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "fontsLocation" () Folder in filestorage with custom fonts.
 @return *os.File*/
func (a *WordsApiService) RenderParagraphWithoutNodePath(ctx context.Context, name models.I, format models.I, index models.I, localVarOptionals map[string]interface{}) (  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/paragraphs/{index}/render"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	return localVarHttpResponse, err
}
/* WordsApiService Renders table to specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param format The destination format.
 @param nodePath Path to the node, which contains tables.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "fontsLocation" () Folder in filestorage with custom fonts.
 @return *os.File*/
func (a *WordsApiService) RenderTable(ctx context.Context, name models.I, format models.I, nodePath models.I, index models.I, localVarOptionals map[string]interface{}) (  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/tables/{index}/render"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	return localVarHttpResponse, err
}
/* WordsApiService Renders table to specified format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param format The destination format.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "fontsLocation" () Folder in filestorage with custom fonts.
 @return *os.File*/
func (a *WordsApiService) RenderTableWithoutNodePath(ctx context.Context, name models.I, format models.I, index models.I, localVarOptionals map[string]interface{}) (  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/tables/{index}/render"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	return localVarHttpResponse, err
}
/* WordsApiService Replaces document text.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param replaceText ReplaceTextResponse with the replace operation settings.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return ReplaceTextResponse*/
func (a *WordsApiService) ReplaceText(ctx context.Context, name models.I, replaceText models.I, localVarOptionals map[string]interface{}) ( models.ReplaceTextResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.ReplaceTextResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/replaceText"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &replaceText
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Replaces the content in the range.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document.
 @param rangeStartIdentifier The range start identifier. Identifier is the value of the &quot;nodeId&quot; field, which every document node has, extended with the prefix &quot;id&quot;. It looks like &quot;id0.0.7&quot;. Also values like &quot;image5&quot; and &quot;table3&quot; can be used as an identifier for images and tables, where the number is an index of the image/table.
 @param rangeText Model with text for replacement.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "rangeEndIdentifier" () The range end identifier.
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
 @return DocumentResponse*/
func (a *WordsApiService) ReplaceWithText(ctx context.Context, name models.I, rangeStartIdentifier models.I, rangeText models.I, localVarOptionals map[string]interface{}) ( models.DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.DocumentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/range/{rangeStartIdentifier}/{rangeEndIdentifier}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rangeStartIdentifier"+"}", fmt.Sprintf("%v", rangeStartIdentifier), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rangeEndIdentifier"+"}", fmt.Sprintf("%v", localVarOptionals["rangeEndIdentifier"]), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &rangeText
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Resets font&#39;s cache.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @return */
func (a *WordsApiService) ResetCache(ctx context.Context) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/fonts/cache"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}
/* WordsApiService Converts document to destination format with detailed settings and saves result to storage.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param saveOptionsData Save options.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "fontsLocation" () Folder in filestorage with custom fonts.
 @return SaveResponse*/
func (a *WordsApiService) SaveAs(ctx context.Context, name models.I, saveOptionsData models.I, localVarOptionals map[string]interface{}) ( models.SaveResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.SaveResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/saveAs"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &saveOptionsData
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Saves the selected range as a new document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document.
 @param rangeStartIdentifier The range start identifier. Identifier is the value of the &quot;nodeId&quot; field, which every document node has, extended with the prefix &quot;id&quot;. It looks like &quot;id0.0.7&quot;. Also values like &quot;image5&quot; and &quot;table3&quot; can be used as an identifier for images and tables, where the number is an index of the image/table.
 @param documentParameters Parameters of a new document.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "rangeEndIdentifier" () The range end identifier.
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return DocumentResponse*/
func (a *WordsApiService) SaveAsRange(ctx context.Context, name models.I, rangeStartIdentifier models.I, documentParameters models.I, localVarOptionals map[string]interface{}) ( models.DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.DocumentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/range/{rangeStartIdentifier}/{rangeEndIdentifier}/SaveAs"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rangeStartIdentifier"+"}", fmt.Sprintf("%v", rangeStartIdentifier), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rangeEndIdentifier"+"}", fmt.Sprintf("%v", localVarOptionals["rangeEndIdentifier"]), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &documentParameters
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Converts document to tiff with detailed settings and saves result to storage.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param saveOptions Tiff save options.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "useAntiAliasing" () Use antialiasing flag.
     @param "useHighQualityRendering" () Use high quality flag.
     @param "imageBrightness" () Brightness for the generated images.
     @param "imageColorMode" () Color mode for the generated images.
     @param "imageContrast" () The contrast for the generated images.
     @param "numeralFormat" () The images numeral format.
     @param "pageCount" () Number of pages to render.
     @param "pageIndex" () Page index to start rendering.
     @param "paperColor" () Background image color.
     @param "pixelFormat" () The pixel format of generated images.
     @param "resolution" () The resolution of generated images.
     @param "scale" () Zoom factor for generated images.
     @param "tiffCompression" () The compression tipe.
     @param "dmlRenderingMode" () Optional, default is Fallback.
     @param "dmlEffectsRenderingMode" () Optional, default is Simplified.
     @param "tiffBinarizationMethod" () Optional, Tiff binarization method, possible values are: FloydSteinbergDithering, Threshold.
     @param "zipOutput" () Optional. A value determining zip output or not.
     @param "fontsLocation" () Folder in filestorage with custom fonts.
 @return SaveResponse*/
func (a *WordsApiService) SaveAsTiff(ctx context.Context, name models.I, saveOptions models.I, localVarOptionals map[string]interface{}) ( models.SaveResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.SaveResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/saveAs/tiff"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &saveOptions
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Searches text in document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param pattern The regular expression used to find matches.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
 @return SearchResponse*/
func (a *WordsApiService) Search(ctx context.Context, name models.I, pattern models.I, localVarOptionals map[string]interface{}) ( models.SearchResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.SearchResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/search"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Splits document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name Original document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "format" () Format to split.
     @param "from" () Start page.
     @param "to" () End page.
     @param "zipOutput" () ZipOutput or not.
     @param "fontsLocation" () Folder in filestorage with custom fonts.
 @return SplitDocumentResponse*/
func (a *WordsApiService) SplitDocument(ctx context.Context, name models.I, localVarOptionals map[string]interface{}) ( models.SplitDocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.SplitDocumentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/split"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Unprotects document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param protectionRequest ProtectionRequest with protection settings.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
 @return ProtectionDataResponse*/
func (a *WordsApiService) UnprotectDocument(ctx context.Context, name models.I, protectionRequest models.I, localVarOptionals map[string]interface{}) ( models.ProtectionDataResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.ProtectionDataResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/protection"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &protectionRequest
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates document bookmark.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param bookmarkData BookmarkData with new bookmark data.
 @param bookmarkName The bookmark name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return BookmarkResponse*/
func (a *WordsApiService) UpdateBookmark(ctx context.Context, name models.I, bookmarkData models.I, bookmarkName models.I, localVarOptionals map[string]interface{}) ( models.BookmarkResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.BookmarkResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/bookmarks/{bookmarkName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bookmarkName"+"}", fmt.Sprintf("%v", bookmarkName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &bookmarkData
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates border properties.             
 &#39;nodePath&#39; should refer to paragraph, cell or row.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param borderProperties Border properties.
 @param nodePath Path to the node with border(node should be paragraph, cell or row).
 @param borderType Border type.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return BorderResponse*/
func (a *WordsApiService) UpdateBorder(ctx context.Context, name models.I, borderProperties models.I, nodePath models.I, borderType models.I, localVarOptionals map[string]interface{}) ( models.BorderResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.BorderResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/borders/{borderType}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"borderType"+"}", fmt.Sprintf("%v", borderType), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &borderProperties
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates the comment, returns updated comment data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param commentIndex The comment index.
 @param comment The comment data.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return CommentResponse*/
func (a *WordsApiService) UpdateComment(ctx context.Context, name models.I, commentIndex models.I, comment models.I, localVarOptionals map[string]interface{}) ( models.CommentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.CommentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/comments/{commentIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"commentIndex"+"}", fmt.Sprintf("%v", commentIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &comment
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates drawing object, returns updated  drawing object&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param drawingObject Drawing object parameters
 @param imageFile File with image
 @param nodePath Path to the node, which contains collection of drawing objects.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return DrawingObjectResponse*/
func (a *WordsApiService) UpdateDrawingObject(ctx context.Context, name models.I, drawingObject models.I, imageFile , nodePath models.I, index models.I, localVarOptionals map[string]interface{}) ( models.DrawingObjectResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.DrawingObjectResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/drawingObjects/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarFormParams.Add("DrawingObject", parameterToString(drawingObject, ""))
	_imageFile := imageFile
	if _imageFile != nil {
		fbs, _ := ioutil.ReadAll(_imageFile)
		_imageFile.Close()
		localFiles[_imageFile.Name()] = fbs
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates drawing object, returns updated  drawing object&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param drawingObject Drawing object parameters
 @param imageFile File with image
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return DrawingObjectResponse*/
func (a *WordsApiService) UpdateDrawingObjectWithoutNodePath(ctx context.Context, name models.I, drawingObject models.I, imageFile , index models.I, localVarOptionals map[string]interface{}) ( models.DrawingObjectResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.DrawingObjectResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/drawingObjects/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarFormParams.Add("DrawingObject", parameterToString(drawingObject, ""))
	_imageFile := imageFile
	if _imageFile != nil {
		fbs, _ := ioutil.ReadAll(_imageFile)
		_imageFile.Close()
		localFiles[_imageFile.Name()] = fbs
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates field&#39;s properties, returns updated field&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param field Field data.
 @param nodePath Path to the node, which contains collection of fields.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return FieldResponse*/
func (a *WordsApiService) UpdateField(ctx context.Context, name models.I, field models.I, nodePath models.I, index models.I, localVarOptionals map[string]interface{}) ( models.FieldResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.FieldResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/fields/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &field
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates (reevaluate) fields in document.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
 @return DocumentResponse*/
func (a *WordsApiService) UpdateFields(ctx context.Context, name models.I, localVarOptionals map[string]interface{}) ( models.DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.DocumentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/updateFields"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates footnote&#39;s properties, returns updated run&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param footnoteDto Footnote data.
 @param nodePath Path to the node, which contains collection of footnotes.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return FootnoteResponse*/
func (a *WordsApiService) UpdateFootnote(ctx context.Context, name models.I, footnoteDto models.I, nodePath models.I, index models.I, localVarOptionals map[string]interface{}) ( models.FootnoteResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.FootnoteResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/footnotes/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &footnoteDto
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates footnote&#39;s properties, returns updated run&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param footnoteDto Footnote data.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return FootnoteResponse*/
func (a *WordsApiService) UpdateFootnoteWithoutNodePath(ctx context.Context, name models.I, footnoteDto models.I, index models.I, localVarOptionals map[string]interface{}) ( models.FootnoteResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.FootnoteResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/footnotes/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &footnoteDto
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates properties of form field, returns updated form field.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param formField From field data.
 @param nodePath Path to the node that contains collection of formfields.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return FormFieldResponse*/
func (a *WordsApiService) UpdateFormField(ctx context.Context, name models.I, formField models.I, nodePath models.I, index models.I, localVarOptionals map[string]interface{}) ( models.FormFieldResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.FormFieldResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/formfields/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &formField
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates properties of form field, returns updated form field.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param formField From field data.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return FormFieldResponse*/
func (a *WordsApiService) UpdateFormFieldWithoutNodePath(ctx context.Context, name models.I, formField models.I, index models.I, localVarOptionals map[string]interface{}) ( models.FormFieldResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.FormFieldResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/formfields/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &formField
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates paragraph format properties, returns updated format properties.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param dto Paragraph format object.
 @param nodePath Path to the node which contains paragraphs.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return ParagraphFormatResponse*/
func (a *WordsApiService) UpdateParagraphFormat(ctx context.Context, name models.I, dto models.I, nodePath models.I, index models.I, localVarOptionals map[string]interface{}) ( models.ParagraphFormatResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.ParagraphFormatResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/paragraphs/{index}/format"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &dto
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates run&#39;s properties, returns updated run&#39;s data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param run Run data.
 @param paragraphPath Path to parent paragraph.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return RunResponse*/
func (a *WordsApiService) UpdateRun(ctx context.Context, name models.I, run models.I, paragraphPath models.I, index models.I, localVarOptionals map[string]interface{}) ( models.RunResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.RunResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{paragraphPath}/runs/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"paragraphPath"+"}", fmt.Sprintf("%v", paragraphPath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &run
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates font properties, returns updated font data.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param fontDto Font dto object.
 @param paragraphPath Path to parent paragraph.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return FontResponse*/
func (a *WordsApiService) UpdateRunFont(ctx context.Context, name models.I, fontDto models.I, paragraphPath models.I, index models.I, localVarOptionals map[string]interface{}) ( models.FontResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.FontResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{paragraphPath}/runs/{index}/font"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"paragraphPath"+"}", fmt.Sprintf("%v", paragraphPath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &fontDto
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates page setup of section.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param sectionIndex Section index.
 @param pageSetup Page setup properties dto.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return SectionPageSetupResponse*/
func (a *WordsApiService) UpdateSectionPageSetup(ctx context.Context, name models.I, sectionIndex models.I, pageSetup models.I, localVarOptionals map[string]interface{}) ( models.SectionPageSetupResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.SectionPageSetupResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/sections/{sectionIndex}/pageSetup"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sectionIndex"+"}", fmt.Sprintf("%v", sectionIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &pageSetup
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates a table cell format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param format The properties.
 @param tableRowPath Path to table row.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return TableCellFormatResponse*/
func (a *WordsApiService) UpdateTableCellFormat(ctx context.Context, name models.I, format models.I, tableRowPath models.I, index models.I, localVarOptionals map[string]interface{}) ( models.TableCellFormatResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.TableCellFormatResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{tableRowPath}/cells/{index}/cellformat"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tableRowPath"+"}", fmt.Sprintf("%v", tableRowPath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &format
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates a table properties.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param properties The properties.
 @param nodePath Path to the node, which contains tables.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return TablePropertiesResponse*/
func (a *WordsApiService) UpdateTableProperties(ctx context.Context, name models.I, properties models.I, nodePath models.I, index models.I, localVarOptionals map[string]interface{}) ( models.TablePropertiesResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.TablePropertiesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{nodePath}/tables/{index}/properties"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodePath"+"}", fmt.Sprintf("%v", nodePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &properties
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates a table properties.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param properties The properties.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return TablePropertiesResponse*/
func (a *WordsApiService) UpdateTablePropertiesWithoutNodePath(ctx context.Context, name models.I, properties models.I, index models.I, localVarOptionals map[string]interface{}) ( models.TablePropertiesResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.TablePropertiesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/tables/{index}/properties"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &properties
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Updates a table row format.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param name The document name.
 @param format Table row format.
 @param tablePath Path to table.
 @param index Object index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" () Original document folder.
     @param "storage" () Original document storage.
     @param "loadEncoding" () Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML.
     @param "password" () Password for opening an encrypted document.
     @param "destFileName" () Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document.
     @param "revisionAuthor" () Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions.
     @param "revisionDateTime" () The date and time to use for revisions.
 @return TableRowFormatResponse*/
func (a *WordsApiService) UpdateTableRowFormat(ctx context.Context, name models.I, format models.I, tablePath models.I, index models.I, localVarOptionals map[string]interface{}) ( models.TableRowFormatResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.TableRowFormatResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/{name}/{tablePath}/rows/{index}/rowformat"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tablePath"+"}", fmt.Sprintf("%v", tablePath), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &format
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
/* WordsApiService Upload file
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param fileContent File to upload
 @param path Path where to upload including filename and extension e.g. /file.ext or /Folder 1/file.ext             If the content is multipart and path does not contains the file name it tries to get them from filename parameter             from Content-Disposition header.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" () Storage name
 @return FilesUploadResult*/
func (a *WordsApiService) UploadFile(ctx context.Context, fileContent , path models.I, localVarOptionals map[string]interface{}) ( models.FilesUploadResult,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload models.FilesUploadResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/words/storage/file/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", path), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	_fileContent := fileContent
	if _fileContent != nil {
		fbs, _ := ioutil.ReadAll(_fileContent)
		_fileContent.Close()
		localFiles[_fileContent.Name()] = fbs
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError models.WordsApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}
