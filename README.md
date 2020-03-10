# Aspose.Words Cloud SDK for Go
This repository contains Aspose.Words Cloud SDK for Go source code. This SDK allows you to work with Aspose.Words Cloud REST APIs in your Go applications quickly and easily, with zero initial cost.

[Aspose.Words Cloud](https://products.aspose.cloud/words/family "Aspose.Words Cloud")  
[API Reference](https://apireference.aspose.cloud/words/)  

# Key Features
* Conversion between various document-related formats (20+ formats supported), including PDF<->Word conversion
* Mail merge and reports generation 
* Splitting Word documents
* Accessing Word document metadata and statistics
* Find and replace
* Watermarks and protection
* Full read & write access to Document Object Model, including sections, paragraphs, text, images, tables, headers/footers and many others

## How to use the SDK?

The complete source code is available in this repository folder. You can either directly use it in your project via source code or get [NuGet distribution](https://www.nuget.org/packages/Aspose.Words-Cloud/) (recommended). For more details, please visit our [documentation website](https://docs.aspose.cloud/display/wordscloud/Available+SDKs#AvailableSDKs-Go).

### Prerequisites

To use Aspose Words for Cloud .NET SDK you need :
- have at least Go 1.13 installed.
- have Go module support in your app.

- to register an account with [Aspose Cloud](https://www.aspose.cloud/) and lookup/create App Key and SID at [Cloud Dashboard](https://dashboard.aspose.cloud/#/apps). There is free quota available. For more details, see [Aspose Cloud Pricing](https://purchase.aspose.cloud/pricing).

### Installation

#### Install Aspose.Words-Cloud

From Visual Stuio Code:

	Add "github.com/aspose-words-cloud/aspose-words-cloud-go/api" in the import section of your code

From the command line:

	go get -v github.com/aspose-words-cloud/aspose-words-cloud-go/api

### Sample usage

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
	"github.com/aspose-words-cloud/aspose-words-cloud-go/api"
	"os"
)

// init words cloud api
config, _ := api.NewConfiguration("config.json")
client, _ := api.NewAPIClient(config)
ctx, _ := client.NewContextWithToken(nil)

// upload test.docx to a cloud
// remote.docx is a name in the cloud
file, _ := os.Open("test.docx")
client.WordsApi.UploadFile(ctx, file, "remote.docx", nil)

// get a text for the first paragraph of the first section
result, _, _ := client.WordsApi.GetParagraphs(ctx, "remote.docx", "sections/0", nil)

fmt.Println(result.Paragraphs.ParagraphLinkList[0].Text)
```

[Tests](Aspose.Words.Cloud.Sdk.Tests) contain various examples of using the SDK.  For other examples, check the product [Developer Guide](https://docs.aspose.cloud/display/wordscloud/Developer+Guide).

## Dependencies
The libraray doesn't uses any non-Google Golang packages.

## Licensing
 
All Aspose.Words Cloud SDKs, helper scripts and templates are licensed under [MIT License](https://github.com/aspose-words-cloud/aspose-words-cloud-go/blob/master/License/LICENSE). 

## Contact Us
Your feedback is very important to us. Please feel free to contact us using our [Support Forums](https://forum.aspose.cloud/c/words).

## Resources
 
[Website](https://www.aspose.cloud/)  
[Product Home](https://products.aspose.cloud/words/family)  
[API Reference](https://apireference.aspose.cloud/words/)  
[Documentation](https://docs.aspose.cloud/display/wordscloud/Home)  
[Blog](https://blog.aspose.cloud/category/words/)  
 
## Other languages
We generate our SDKs in different languages so you may check if yours is available in our [list](https://github.com/aspose-words-cloud).
 
If you don't find your language in the list, feel free to request it from us, or use raw REST API requests as you can find it [here](https://products.aspose.cloud/words/curl).
