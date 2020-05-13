Aspose.Wors Cloud SDK for Go wraps Aspose.Words REST API so you could seamlessly integrate Microsoft Word® document generation, manipulation, conversion & inspection features into your own Go applications.

# Word Document Processing in the Cloud

[Aspose.Words Cloud SDK for Go](https://products.aspose.cloud/words/go) allows to work with document headers, footers, page numbering, tables, sections, document comments, drawing objects, FormFields, fonts, hyperlinks, ranges, paragraphs, math objects, watermarks, track changes and document protection. It also assists in appending documents, splitting documents as well as converting document to other supported file formats. 

Feel free to explore the [Developer's Guide](https://docs.aspose.cloud/display/wordscloud/Developer+Guide) & [API Reference](https://apireference.aspose.cloud/words/) to know all about Aspose.Words Cloud API. 

## Document Processing Features

- Convert between various document-related formats, including Word to PDF & vice versa.
- Mail merge and report generation in the Cloud.
- Split & merge Word documents.
- Access Word document metadata.
- Find and replace text.
- Add & remove watermarks and protection.
- Read & write access to Document Object Model.

## Enhancements in Version 20.5

- Added methods to work with Word document lists
  - getLists
  - getList
  - insertList
  - updateList
  - updateListLevel
- Added methods to work with styles
  - getStyles
  - updateStyle
  - insertStyle
  - copyStyle
  - getStyleFromDocumentElement
  - applyStyleToDocumentElement
- Added methods to work with paragraph list format
  - getParagraphListFormat
  - getParagraphListFormatWithoutNodePath
  - updateParagraphListFormat
  - deleteParagraphListFormat
- Added methods to work with paragraph tab stops
  - g1etParagraphTabStops
  - insertOrUpdateParagraphTabStop
  - deleteAllParagraphTabStops
  - deleteParagraphTabStop
- Added methods to build reports
  - buildReport
  - buildReportOnline
- Added Shading property to ParagraphFormat

## Read & Write Document Formats

**Microsoft Word:** DOC, DOCX, RTF, DOT, DOTX, DOTM, FlatOPC (XML)
**OpenOffice:** ODT, OTT
**WordprocessingML:** XML
**Web:** HTML, MHTML, HtmlFixed
**Text:** TXT
**Fixed Layout:** PDF

## Save Document As

**Fixed Layout:** PDF/A, XPS, OpenXPS, PS
**Images:** JPEG, PNG, BMP, SVG, TIFF, EMF
**Others:** PCL

## Getting Started with Aspose.Words Cloud SDK for Go

Firstly, create an account at [Aspose for Cloud](https://dashboard.aspose.cloud/#/apps) to get your application information and free quota to use the API. 

#### Install Aspose.Words-Cloud

From Visual Stuio Code:

	Add "github.com/aspose-words-cloud/aspose-words-cloud-go/v2005/api" and "github.com/aspose-words-cloud/aspose-words-cloud-go/v2005/api/models" in the import section of your code

From the command line:

	go get -v github.com/aspose-words-cloud/aspose-words-cloud-go/2005/api

The complete source code is available at [GitHub Repository](https://github.com/aspose-words-cloud/aspose-words-cloud-go).

### SDK Dependencies

The libraray doesn't uses any non-Google Golang packages.

## Sample usage via the SDK

The examples below show how your application have to initiate and get a text of the first paragraph using Aspose.Words-Cloud library:

Config.json file:
```csharp
{
	"AppKey": "your app key",
	"AppSid": "your app sid",
	"BaseUrl": "https://api.aspose.cloud"
} 
```
Go code:

```csharp
import (
	"fmt"
	"github.com/aspose-words-cloud/aspose-words-cloud-go/v2005/api"
	"os"
)

// init words cloud api
config, _ := api.NewConfiguration("config.json")
wordsApi, ctx, _ := api.CreateWordsApi(config)

// upload test.docx to a cloud
// remote.docx is a name in the cloud
file, _ := os.Open("test.docx")
wordsApi.UploadFile(ctx, file, "remote.docx", nil)

// get a text for the first paragraph of the first section
result, _, _ := wordsApi.GetParagraphs(ctx, "remote.docx", "sections/0", nil)

fmt.Println(result.Paragraphs.ParagraphLinkList[0].Text)
```

[Product Page](https://products.aspose.cloud/words/go) | [Documentation](https://docs.aspose.cloud/display/wordscloud/Home) | [API Reference](https://apireference.aspose.cloud/words/) | [Code Samples](https://github.com/aspose-words-cloud/aspose-words-cloud-go) | [Blog](https://blog.aspose.cloud/category/words/) | [Free Support](https://forum.aspose.cloud/c/words) | [Free Trial](https://dashboard.aspose.cloud/#/apps)