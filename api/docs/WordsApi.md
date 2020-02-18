# \WordsApi

All URIs are relative to *https://localhost/v4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptAllRevisions**](WordsApi.md#AcceptAllRevisions) | **Put** /words/{name}/revisions/acceptAll | Accepts all revisions in document.
[**AppendDocument**](WordsApi.md#AppendDocument) | **Put** /words/{name}/appendDocument | Appends documents to original document.
[**Classify**](WordsApi.md#Classify) | **Put** /words/classify | Classifies raw text.
[**ClassifyDocument**](WordsApi.md#ClassifyDocument) | **Get** /words/{documentName}/classify | Classifies document.
[**CompareDocument**](WordsApi.md#CompareDocument) | **Put** /words/{name}/compareDocument | Compares document with original document.
[**ConvertDocument**](WordsApi.md#ConvertDocument) | **Put** /words/convert | Converts document from the request&#39;s content to the specified format .
[**CopyFile**](WordsApi.md#CopyFile) | **Put** /words/storage/file/copy/{srcPath} | Copy file
[**CopyFolder**](WordsApi.md#CopyFolder) | **Put** /words/storage/folder/copy/{srcPath} | Copy folder
[**CreateDocument**](WordsApi.md#CreateDocument) | **Put** /words/create | Creates new document. Document is created with format which is recognized from file extensions. Supported extensions: \&quot;.doc\&quot;, \&quot;.docx\&quot;, \&quot;.docm\&quot;, \&quot;.dot\&quot;, \&quot;.dotm\&quot;, \&quot;.dotx\&quot;, \&quot;.flatopc\&quot;, \&quot;.fopc\&quot;, \&quot;.flatopc_macro\&quot;, \&quot;.fopc_macro\&quot;, \&quot;.flatopc_template\&quot;, \&quot;.fopc_template\&quot;, \&quot;.flatopc_template_macro\&quot;, \&quot;.fopc_template_macro\&quot;, \&quot;.wordml\&quot;, \&quot;.wml\&quot;, \&quot;.rtf\&quot;.
[**CreateFolder**](WordsApi.md#CreateFolder) | **Put** /words/storage/folder/{path} | Create the folder
[**CreateOrUpdateDocumentProperty**](WordsApi.md#CreateOrUpdateDocumentProperty) | **Put** /words/{name}/documentProperties/{propertyName} | Adds new or update existing document property.
[**DeleteBorder**](WordsApi.md#DeleteBorder) | **Delete** /words/{name}/{nodePath}/borders/{borderType} | Resets border properties to default values.             
[**DeleteBorders**](WordsApi.md#DeleteBorders) | **Delete** /words/{name}/{nodePath}/borders | Resets borders properties to default values.             
[**DeleteComment**](WordsApi.md#DeleteComment) | **Delete** /words/{name}/comments/{commentIndex} | Removes comment from document.
[**DeleteDocumentProperty**](WordsApi.md#DeleteDocumentProperty) | **Delete** /words/{name}/documentProperties/{propertyName} | Deletes document property.
[**DeleteDrawingObject**](WordsApi.md#DeleteDrawingObject) | **Delete** /words/{name}/{nodePath}/drawingObjects/{index} | Removes drawing object from document.
[**DeleteDrawingObjectWithoutNodePath**](WordsApi.md#DeleteDrawingObjectWithoutNodePath) | **Delete** /words/{name}/drawingObjects/{index} | Removes drawing object from document.
[**DeleteField**](WordsApi.md#DeleteField) | **Delete** /words/{name}/{nodePath}/fields/{index} | Deletes field from document.
[**DeleteFieldWithoutNodePath**](WordsApi.md#DeleteFieldWithoutNodePath) | **Delete** /words/{name}/fields/{index} | Deletes field from document.
[**DeleteFields**](WordsApi.md#DeleteFields) | **Delete** /words/{name}/{nodePath}/fields | Removes fields from section paragraph.
[**DeleteFieldsWithoutNodePath**](WordsApi.md#DeleteFieldsWithoutNodePath) | **Delete** /words/{name}/fields | Removes fields from section paragraph.
[**DeleteFile**](WordsApi.md#DeleteFile) | **Delete** /words/storage/file/{path} | Delete file
[**DeleteFolder**](WordsApi.md#DeleteFolder) | **Delete** /words/storage/folder/{path} | Delete folder
[**DeleteFootnote**](WordsApi.md#DeleteFootnote) | **Delete** /words/{name}/{nodePath}/footnotes/{index} | Removes footnote from document.
[**DeleteFootnoteWithoutNodePath**](WordsApi.md#DeleteFootnoteWithoutNodePath) | **Delete** /words/{name}/footnotes/{index} | Removes footnote from document.
[**DeleteFormField**](WordsApi.md#DeleteFormField) | **Delete** /words/{name}/{nodePath}/formfields/{index} | Removes form field from document.
[**DeleteFormFieldWithoutNodePath**](WordsApi.md#DeleteFormFieldWithoutNodePath) | **Delete** /words/{name}/formfields/{index} | Removes form field from document.
[**DeleteHeaderFooter**](WordsApi.md#DeleteHeaderFooter) | **Delete** /words/{name}/{sectionPath}/headersfooters/{index} | Deletes header/footer from document.
[**DeleteHeadersFooters**](WordsApi.md#DeleteHeadersFooters) | **Delete** /words/{name}/{sectionPath}/headersfooters | Deletes document headers and footers.
[**DeleteMacros**](WordsApi.md#DeleteMacros) | **Delete** /words/{name}/macros | Removes macros from document.
[**DeleteOfficeMathObject**](WordsApi.md#DeleteOfficeMathObject) | **Delete** /words/{name}/{nodePath}/OfficeMathObjects/{index} | Removes OfficeMath object from document.
[**DeleteOfficeMathObjectWithoutNodePath**](WordsApi.md#DeleteOfficeMathObjectWithoutNodePath) | **Delete** /words/{name}/OfficeMathObjects/{index} | Removes OfficeMath object from document.
[**DeleteParagraph**](WordsApi.md#DeleteParagraph) | **Delete** /words/{name}/{nodePath}/paragraphs/{index} | Removes paragraph from section.
[**DeleteParagraphWithoutNodePath**](WordsApi.md#DeleteParagraphWithoutNodePath) | **Delete** /words/{name}/paragraphs/{index} | Removes paragraph from section.
[**DeleteRun**](WordsApi.md#DeleteRun) | **Delete** /words/{name}/{paragraphPath}/runs/{index} | Removes run from document.
[**DeleteSection**](WordsApi.md#DeleteSection) | **Delete** /words/{name}/sections/{sectionIndex} | Removes section from document.
[**DeleteTable**](WordsApi.md#DeleteTable) | **Delete** /words/{name}/{nodePath}/tables/{index} | Deletes a table.
[**DeleteTableCell**](WordsApi.md#DeleteTableCell) | **Delete** /words/{name}/{tableRowPath}/cells/{index} | Deletes a table cell.
[**DeleteTableRow**](WordsApi.md#DeleteTableRow) | **Delete** /words/{name}/{tablePath}/rows/{index} | Deletes a table row.
[**DeleteTableWithoutNodePath**](WordsApi.md#DeleteTableWithoutNodePath) | **Delete** /words/{name}/tables/{index} | Deletes a table.
[**DeleteWatermark**](WordsApi.md#DeleteWatermark) | **Post** /words/{name}/watermarks/deleteLast | Deletes watermark (for deleting last watermark from the document).
[**DownloadFile**](WordsApi.md#DownloadFile) | **Get** /words/storage/file/{path} | Download file
[**ExecuteMailMerge**](WordsApi.md#ExecuteMailMerge) | **Put** /words/{name}/MailMerge | Executes document mail merge operation.
[**ExecuteMailMergeOnline**](WordsApi.md#ExecuteMailMergeOnline) | **Put** /words/MailMerge | Executes document mail merge online.
[**GetAvailableFonts**](WordsApi.md#GetAvailableFonts) | **Get** /words/fonts/available | Gets the list of fonts, available for document processing.
[**GetBookmarkByName**](WordsApi.md#GetBookmarkByName) | **Get** /words/{name}/bookmarks/{bookmarkName} | Reads document bookmark data by its name.
[**GetBookmarks**](WordsApi.md#GetBookmarks) | **Get** /words/{name}/bookmarks | Reads document bookmarks common info.
[**GetBorder**](WordsApi.md#GetBorder) | **Get** /words/{name}/{nodePath}/borders/{borderType} | Returns a border.
[**GetBorders**](WordsApi.md#GetBorders) | **Get** /words/{name}/{nodePath}/borders | Returns a collection of borders.
[**GetComment**](WordsApi.md#GetComment) | **Get** /words/{name}/comments/{commentIndex} | Gets comment from document.
[**GetComments**](WordsApi.md#GetComments) | **Get** /words/{name}/comments | Gets comments from document.
[**GetDocument**](WordsApi.md#GetDocument) | **Get** /words/{documentName} | Reads document common info.
[**GetDocumentDrawingObjectByIndex**](WordsApi.md#GetDocumentDrawingObjectByIndex) | **Get** /words/{name}/{nodePath}/drawingObjects/{index} | Reads document drawing object common info by its index or convert to format specified.
[**GetDocumentDrawingObjectByIndexWithoutNodePath**](WordsApi.md#GetDocumentDrawingObjectByIndexWithoutNodePath) | **Get** /words/{name}/drawingObjects/{index} | Reads document drawing object common info by its index or convert to format specified.
[**GetDocumentDrawingObjectImageData**](WordsApi.md#GetDocumentDrawingObjectImageData) | **Get** /words/{name}/{nodePath}/drawingObjects/{index}/imageData | Reads drawing object image data.
[**GetDocumentDrawingObjectImageDataWithoutNodePath**](WordsApi.md#GetDocumentDrawingObjectImageDataWithoutNodePath) | **Get** /words/{name}/drawingObjects/{index}/imageData | Reads drawing object image data.
[**GetDocumentDrawingObjectOleData**](WordsApi.md#GetDocumentDrawingObjectOleData) | **Get** /words/{name}/{nodePath}/drawingObjects/{index}/oleData | Gets drawing object OLE data.
[**GetDocumentDrawingObjectOleDataWithoutNodePath**](WordsApi.md#GetDocumentDrawingObjectOleDataWithoutNodePath) | **Get** /words/{name}/drawingObjects/{index}/oleData | Gets drawing object OLE data.
[**GetDocumentDrawingObjects**](WordsApi.md#GetDocumentDrawingObjects) | **Get** /words/{name}/{nodePath}/drawingObjects | Reads document drawing objects common info.
[**GetDocumentDrawingObjectsWithoutNodePath**](WordsApi.md#GetDocumentDrawingObjectsWithoutNodePath) | **Get** /words/{name}/drawingObjects | Reads document drawing objects common info.
[**GetDocumentFieldNames**](WordsApi.md#GetDocumentFieldNames) | **Get** /words/{name}/mailMerge/FieldNames | Reads document field names.
[**GetDocumentFieldNamesOnline**](WordsApi.md#GetDocumentFieldNamesOnline) | **Put** /words/mailMerge/FieldNames | Reads document field names.
[**GetDocumentHyperlinkByIndex**](WordsApi.md#GetDocumentHyperlinkByIndex) | **Get** /words/{name}/hyperlinks/{hyperlinkIndex} | Reads document hyperlink by its index.
[**GetDocumentHyperlinks**](WordsApi.md#GetDocumentHyperlinks) | **Get** /words/{name}/hyperlinks | Reads document hyperlinks common info.
[**GetDocumentProperties**](WordsApi.md#GetDocumentProperties) | **Get** /words/{name}/documentProperties | Reads document properties info.
[**GetDocumentProperty**](WordsApi.md#GetDocumentProperty) | **Get** /words/{name}/documentProperties/{propertyName} | Reads document property info by the property name.
[**GetDocumentProtection**](WordsApi.md#GetDocumentProtection) | **Get** /words/{name}/protection | Reads document protection common info.
[**GetDocumentStatistics**](WordsApi.md#GetDocumentStatistics) | **Get** /words/{name}/statistics | Reads document statistics.
[**GetDocumentWithFormat**](WordsApi.md#GetDocumentWithFormat) | **Get** /words/{name} | Exports the document into the specified format.
[**GetField**](WordsApi.md#GetField) | **Get** /words/{name}/{nodePath}/fields/{index} | Gets field from document.
[**GetFieldWithoutNodePath**](WordsApi.md#GetFieldWithoutNodePath) | **Get** /words/{name}/fields/{index} | Gets field from document.
[**GetFields**](WordsApi.md#GetFields) | **Get** /words/{name}/{nodePath}/fields | Get fields from document.
[**GetFieldsWithoutNodePath**](WordsApi.md#GetFieldsWithoutNodePath) | **Get** /words/{name}/fields | Get fields from document.
[**GetFilesList**](WordsApi.md#GetFilesList) | **Get** /words/storage/folder/{path} | Get all files and folders within a folder
[**GetFootnote**](WordsApi.md#GetFootnote) | **Get** /words/{name}/{nodePath}/footnotes/{index} | Reads footnote by index.
[**GetFootnoteWithoutNodePath**](WordsApi.md#GetFootnoteWithoutNodePath) | **Get** /words/{name}/footnotes/{index} | Reads footnote by index.
[**GetFootnotes**](WordsApi.md#GetFootnotes) | **Get** /words/{name}/{nodePath}/footnotes | Gets footnotes from document.
[**GetFootnotesWithoutNodePath**](WordsApi.md#GetFootnotesWithoutNodePath) | **Get** /words/{name}/footnotes | Gets footnotes from document.
[**GetFormField**](WordsApi.md#GetFormField) | **Get** /words/{name}/{nodePath}/formfields/{index} | Returns representation of an one of the form field.
[**GetFormFieldWithoutNodePath**](WordsApi.md#GetFormFieldWithoutNodePath) | **Get** /words/{name}/formfields/{index} | Returns representation of an one of the form field.
[**GetFormFields**](WordsApi.md#GetFormFields) | **Get** /words/{name}/{nodePath}/formfields | Gets form fields from document.
[**GetFormFieldsWithoutNodePath**](WordsApi.md#GetFormFieldsWithoutNodePath) | **Get** /words/{name}/formfields | Gets form fields from document.
[**GetHeaderFooter**](WordsApi.md#GetHeaderFooter) | **Get** /words/{name}/headersfooters/{headerFooterIndex} | Returns a header/footer from the document by index.
[**GetHeaderFooterOfSection**](WordsApi.md#GetHeaderFooterOfSection) | **Get** /words/{name}/sections/{sectionIndex}/headersfooters/{headerFooterIndex} | Returns a header/footer from the document section.
[**GetHeaderFooters**](WordsApi.md#GetHeaderFooters) | **Get** /words/{name}/{sectionPath}/headersfooters | Returns a list of header/footers from the document.
[**GetOfficeMathObject**](WordsApi.md#GetOfficeMathObject) | **Get** /words/{name}/{nodePath}/OfficeMathObjects/{index} | Reads OfficeMath object by index.
[**GetOfficeMathObjectWithoutNodePath**](WordsApi.md#GetOfficeMathObjectWithoutNodePath) | **Get** /words/{name}/OfficeMathObjects/{index} | Reads OfficeMath object by index.
[**GetOfficeMathObjects**](WordsApi.md#GetOfficeMathObjects) | **Get** /words/{name}/{nodePath}/OfficeMathObjects | Gets OfficeMath objects from document.
[**GetOfficeMathObjectsWithoutNodePath**](WordsApi.md#GetOfficeMathObjectsWithoutNodePath) | **Get** /words/{name}/OfficeMathObjects | Gets OfficeMath objects from document.
[**GetParagraph**](WordsApi.md#GetParagraph) | **Get** /words/{name}/{nodePath}/paragraphs/{index} | This resource represents one of the paragraphs contained in the document.
[**GetParagraphFormat**](WordsApi.md#GetParagraphFormat) | **Get** /words/{name}/{nodePath}/paragraphs/{index}/format | Represents all the formatting for a paragraph.
[**GetParagraphFormatWithoutNodePath**](WordsApi.md#GetParagraphFormatWithoutNodePath) | **Get** /words/{name}/paragraphs/{index}/format | Represents all the formatting for a paragraph.
[**GetParagraphWithoutNodePath**](WordsApi.md#GetParagraphWithoutNodePath) | **Get** /words/{name}/paragraphs/{index} | This resource represents one of the paragraphs contained in the document.
[**GetParagraphs**](WordsApi.md#GetParagraphs) | **Get** /words/{name}/{nodePath}/paragraphs | Returns a list of paragraphs that are contained in the document.
[**GetParagraphsWithoutNodePath**](WordsApi.md#GetParagraphsWithoutNodePath) | **Get** /words/{name}/paragraphs | Returns a list of paragraphs that are contained in the document.
[**GetRangeText**](WordsApi.md#GetRangeText) | **Get** /words/{name}/range/{rangeStartIdentifier}/{rangeEndIdentifier} | Gets the text from the range.
[**GetRun**](WordsApi.md#GetRun) | **Get** /words/{name}/{paragraphPath}/runs/{index} | This resource represents run of text contained in the document.
[**GetRunFont**](WordsApi.md#GetRunFont) | **Get** /words/{name}/{paragraphPath}/runs/{index}/font | This resource represents font of run.
[**GetRuns**](WordsApi.md#GetRuns) | **Get** /words/{name}/{paragraphPath}/runs | This resource represents collection of runs in the paragraph.
[**GetSection**](WordsApi.md#GetSection) | **Get** /words/{name}/sections/{sectionIndex} | Gets document section by index.
[**GetSectionPageSetup**](WordsApi.md#GetSectionPageSetup) | **Get** /words/{name}/sections/{sectionIndex}/pageSetup | Gets page setup of section.
[**GetSections**](WordsApi.md#GetSections) | **Get** /words/{name}/sections | Returns a list of sections that are contained in the document.
[**GetTable**](WordsApi.md#GetTable) | **Get** /words/{name}/{nodePath}/tables/{index} | Returns a table.
[**GetTableCell**](WordsApi.md#GetTableCell) | **Get** /words/{name}/{tableRowPath}/cells/{index} | Returns a table cell.
[**GetTableCellFormat**](WordsApi.md#GetTableCellFormat) | **Get** /words/{name}/{tableRowPath}/cells/{index}/cellformat | Returns a table cell format.
[**GetTableProperties**](WordsApi.md#GetTableProperties) | **Get** /words/{name}/{nodePath}/tables/{index}/properties | Returns a table properties.
[**GetTablePropertiesWithoutNodePath**](WordsApi.md#GetTablePropertiesWithoutNodePath) | **Get** /words/{name}/tables/{index}/properties | Returns a table properties.
[**GetTableRow**](WordsApi.md#GetTableRow) | **Get** /words/{name}/{tablePath}/rows/{index} | Returns a table row.
[**GetTableRowFormat**](WordsApi.md#GetTableRowFormat) | **Get** /words/{name}/{tablePath}/rows/{index}/rowformat | Returns a table row format.
[**GetTableWithoutNodePath**](WordsApi.md#GetTableWithoutNodePath) | **Get** /words/{name}/tables/{index} | Returns a table.
[**GetTables**](WordsApi.md#GetTables) | **Get** /words/{name}/{nodePath}/tables | Returns a list of tables that are contained in the document.
[**GetTablesWithoutNodePath**](WordsApi.md#GetTablesWithoutNodePath) | **Get** /words/{name}/tables | Returns a list of tables that are contained in the document.
[**InsertComment**](WordsApi.md#InsertComment) | **Post** /words/{name}/comments | Adds comment to document, returns inserted comment data.
[**InsertDrawingObject**](WordsApi.md#InsertDrawingObject) | **Post** /words/{name}/{nodePath}/drawingObjects | Adds drawing object to document, returns added  drawing object&#39;s data.
[**InsertDrawingObjectWithoutNodePath**](WordsApi.md#InsertDrawingObjectWithoutNodePath) | **Post** /words/{name}/drawingObjects | Adds drawing object to document, returns added  drawing object&#39;s data.
[**InsertField**](WordsApi.md#InsertField) | **Post** /words/{name}/{nodePath}/fields | Adds field to document, returns inserted field&#39;s data.
[**InsertFieldWithoutNodePath**](WordsApi.md#InsertFieldWithoutNodePath) | **Post** /words/{name}/fields | Adds field to document, returns inserted field&#39;s data.
[**InsertFootnote**](WordsApi.md#InsertFootnote) | **Post** /words/{name}/{nodePath}/footnotes | Adds footnote to document, returns added footnote&#39;s data.
[**InsertFootnoteWithoutNodePath**](WordsApi.md#InsertFootnoteWithoutNodePath) | **Post** /words/{name}/footnotes | Adds footnote to document, returns added footnote&#39;s data.
[**InsertFormField**](WordsApi.md#InsertFormField) | **Post** /words/{name}/{nodePath}/formfields | Adds form field to paragraph, returns added form field&#39;s data.
[**InsertFormFieldWithoutNodePath**](WordsApi.md#InsertFormFieldWithoutNodePath) | **Post** /words/{name}/formfields | Adds form field to paragraph, returns added form field&#39;s data.
[**InsertHeaderFooter**](WordsApi.md#InsertHeaderFooter) | **Put** /words/{name}/{sectionPath}/headersfooters | Inserts to document header or footer.
[**InsertPageNumbers**](WordsApi.md#InsertPageNumbers) | **Put** /words/{name}/PageNumbers | Inserts document page numbers.
[**InsertParagraph**](WordsApi.md#InsertParagraph) | **Post** /words/{name}/{nodePath}/paragraphs | Adds paragraph to document, returns added paragraph&#39;s data.
[**InsertRun**](WordsApi.md#InsertRun) | **Post** /words/{name}/{paragraphPath}/runs | Adds run to document, returns added paragraph&#39;s data.
[**InsertTable**](WordsApi.md#InsertTable) | **Post** /words/{name}/{nodePath}/tables | Adds table to document, returns added table&#39;s data.             
[**InsertTableCell**](WordsApi.md#InsertTableCell) | **Post** /words/{name}/{tableRowPath}/cells | Adds table cell to table, returns added cell&#39;s data.             
[**InsertTableRow**](WordsApi.md#InsertTableRow) | **Post** /words/{name}/{tablePath}/rows | Adds table row to table, returns added row&#39;s data.             
[**InsertTableWithoutNodePath**](WordsApi.md#InsertTableWithoutNodePath) | **Post** /words/{name}/tables | Adds table to document, returns added table&#39;s data.             
[**InsertWatermarkImage**](WordsApi.md#InsertWatermarkImage) | **Post** /words/{name}/watermarks/images | Inserts document watermark image.
[**InsertWatermarkText**](WordsApi.md#InsertWatermarkText) | **Post** /words/{name}/watermarks/texts | Inserts document watermark text.
[**LoadWebDocument**](WordsApi.md#LoadWebDocument) | **Put** /words/loadWebDocument | Loads new document from web into the file with any supported format of data.
[**MoveFile**](WordsApi.md#MoveFile) | **Put** /words/storage/file/move/{srcPath} | Move file
[**MoveFolder**](WordsApi.md#MoveFolder) | **Put** /words/storage/folder/move/{srcPath} | Move folder
[**ProtectDocument**](WordsApi.md#ProtectDocument) | **Put** /words/{name}/protection | Protects document.
[**RejectAllRevisions**](WordsApi.md#RejectAllRevisions) | **Put** /words/{name}/revisions/rejectAll | Rejects all revisions in document.
[**RemoveRange**](WordsApi.md#RemoveRange) | **Delete** /words/{name}/range/{rangeStartIdentifier}/{rangeEndIdentifier} | Removes the range from the document.
[**RenderDrawingObject**](WordsApi.md#RenderDrawingObject) | **Get** /words/{name}/{nodePath}/drawingObjects/{index}/render | Renders drawing object to specified format.
[**RenderDrawingObjectWithoutNodePath**](WordsApi.md#RenderDrawingObjectWithoutNodePath) | **Get** /words/{name}/drawingObjects/{index}/render | Renders drawing object to specified format.
[**RenderMathObject**](WordsApi.md#RenderMathObject) | **Get** /words/{name}/{nodePath}/OfficeMathObjects/{index}/render | Renders math object to specified format.
[**RenderMathObjectWithoutNodePath**](WordsApi.md#RenderMathObjectWithoutNodePath) | **Get** /words/{name}/OfficeMathObjects/{index}/render | Renders math object to specified format.
[**RenderPage**](WordsApi.md#RenderPage) | **Get** /words/{name}/pages/{pageIndex}/render | Renders page to specified format.
[**RenderParagraph**](WordsApi.md#RenderParagraph) | **Get** /words/{name}/{nodePath}/paragraphs/{index}/render | Renders paragraph to specified format.
[**RenderParagraphWithoutNodePath**](WordsApi.md#RenderParagraphWithoutNodePath) | **Get** /words/{name}/paragraphs/{index}/render | Renders paragraph to specified format.
[**RenderTable**](WordsApi.md#RenderTable) | **Get** /words/{name}/{nodePath}/tables/{index}/render | Renders table to specified format.
[**RenderTableWithoutNodePath**](WordsApi.md#RenderTableWithoutNodePath) | **Get** /words/{name}/tables/{index}/render | Renders table to specified format.
[**ReplaceText**](WordsApi.md#ReplaceText) | **Put** /words/{name}/replaceText | Replaces document text.
[**ReplaceWithText**](WordsApi.md#ReplaceWithText) | **Post** /words/{name}/range/{rangeStartIdentifier}/{rangeEndIdentifier} | Replaces the content in the range.
[**ResetCache**](WordsApi.md#ResetCache) | **Delete** /words/fonts/cache | Resets font&#39;s cache.
[**SaveAs**](WordsApi.md#SaveAs) | **Put** /words/{name}/saveAs | Converts document to destination format with detailed settings and saves result to storage.
[**SaveAsRange**](WordsApi.md#SaveAsRange) | **Post** /words/{name}/range/{rangeStartIdentifier}/{rangeEndIdentifier}/SaveAs | Saves the selected range as a new document.
[**SaveAsTiff**](WordsApi.md#SaveAsTiff) | **Put** /words/{name}/saveAs/tiff | Converts document to tiff with detailed settings and saves result to storage.
[**Search**](WordsApi.md#Search) | **Get** /words/{name}/search | Searches text in document.
[**SplitDocument**](WordsApi.md#SplitDocument) | **Put** /words/{name}/split | Splits document.
[**UnprotectDocument**](WordsApi.md#UnprotectDocument) | **Delete** /words/{name}/protection | Unprotects document.
[**UpdateBookmark**](WordsApi.md#UpdateBookmark) | **Put** /words/{name}/bookmarks/{bookmarkName} | Updates document bookmark.
[**UpdateBorder**](WordsApi.md#UpdateBorder) | **Put** /words/{name}/{nodePath}/borders/{borderType} | Updates border properties.             
[**UpdateComment**](WordsApi.md#UpdateComment) | **Put** /words/{name}/comments/{commentIndex} | Updates the comment, returns updated comment data.
[**UpdateDrawingObject**](WordsApi.md#UpdateDrawingObject) | **Put** /words/{name}/{nodePath}/drawingObjects/{index} | Updates drawing object, returns updated  drawing object&#39;s data.
[**UpdateDrawingObjectWithoutNodePath**](WordsApi.md#UpdateDrawingObjectWithoutNodePath) | **Put** /words/{name}/drawingObjects/{index} | Updates drawing object, returns updated  drawing object&#39;s data.
[**UpdateField**](WordsApi.md#UpdateField) | **Put** /words/{name}/{nodePath}/fields/{index} | Updates field&#39;s properties, returns updated field&#39;s data.
[**UpdateFields**](WordsApi.md#UpdateFields) | **Put** /words/{name}/updateFields | Updates (reevaluate) fields in document.
[**UpdateFootnote**](WordsApi.md#UpdateFootnote) | **Put** /words/{name}/{nodePath}/footnotes/{index} | Updates footnote&#39;s properties, returns updated run&#39;s data.
[**UpdateFootnoteWithoutNodePath**](WordsApi.md#UpdateFootnoteWithoutNodePath) | **Put** /words/{name}/footnotes/{index} | Updates footnote&#39;s properties, returns updated run&#39;s data.
[**UpdateFormField**](WordsApi.md#UpdateFormField) | **Put** /words/{name}/{nodePath}/formfields/{index} | Updates properties of form field, returns updated form field.
[**UpdateFormFieldWithoutNodePath**](WordsApi.md#UpdateFormFieldWithoutNodePath) | **Put** /words/{name}/formfields/{index} | Updates properties of form field, returns updated form field.
[**UpdateParagraphFormat**](WordsApi.md#UpdateParagraphFormat) | **Put** /words/{name}/{nodePath}/paragraphs/{index}/format | Updates paragraph format properties, returns updated format properties.
[**UpdateRun**](WordsApi.md#UpdateRun) | **Put** /words/{name}/{paragraphPath}/runs/{index} | Updates run&#39;s properties, returns updated run&#39;s data.
[**UpdateRunFont**](WordsApi.md#UpdateRunFont) | **Put** /words/{name}/{paragraphPath}/runs/{index}/font | Updates font properties, returns updated font data.
[**UpdateSectionPageSetup**](WordsApi.md#UpdateSectionPageSetup) | **Put** /words/{name}/sections/{sectionIndex}/pageSetup | Updates page setup of section.
[**UpdateTableCellFormat**](WordsApi.md#UpdateTableCellFormat) | **Put** /words/{name}/{tableRowPath}/cells/{index}/cellformat | Updates a table cell format.
[**UpdateTableProperties**](WordsApi.md#UpdateTableProperties) | **Put** /words/{name}/{nodePath}/tables/{index}/properties | Updates a table properties.
[**UpdateTablePropertiesWithoutNodePath**](WordsApi.md#UpdateTablePropertiesWithoutNodePath) | **Put** /words/{name}/tables/{index}/properties | Updates a table properties.
[**UpdateTableRowFormat**](WordsApi.md#UpdateTableRowFormat) | **Put** /words/{name}/{tablePath}/rows/{index}/rowformat | Updates a table row format.
[**UploadFile**](WordsApi.md#UploadFile) | **Put** /words/storage/file/{path} | Upload file


# **AcceptAllRevisions**
> RevisionsModificationResponse AcceptAllRevisions(ctx, name, optional)
Accepts all revisions in document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 

### Return type

[**RevisionsModificationResponse**](RevisionsModificationResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppendDocument**
> DocumentResponse AppendDocument(ctx, name, documentList, optional)
Appends documents to original document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| Original document name. | 
  **documentList** | [**DocumentEntryList**](DocumentEntryList.md)| with a list of documents to append.             | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Original document name. | 
 **documentList** | [**DocumentEntryList**](DocumentEntryList.md)| with a list of documents to append.             | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

[**DocumentResponse**](DocumentResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Classify**
> ClassificationResponse Classify(ctx, text, optional)
Classifies raw text.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **text** | **string**| Text to classify. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | **string**| Text to classify. | 
 **bestClassesCount** | **string**| Number of the best classes to return. | [default to 1]

### Return type

[**ClassificationResponse**](ClassificationResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClassifyDocument**
> ClassificationResponse ClassifyDocument(ctx, documentName, optional)
Classifies document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **documentName** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **documentName** | **string**| The document name. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **bestClassesCount** | **string**| Count of the best classes to return. | [default to 1]
 **taxonomy** | **string**| Taxonomy to use for classification return. | [default to default]

### Return type

[**ClassificationResponse**](ClassificationResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CompareDocument**
> DocumentResponse CompareDocument(ctx, name, compareData, optional)
Compares document with original document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| Original document name. | 
  **compareData** | [**CompareData**](CompareData.md)| with a document to compare.             | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Original document name. | 
 **compareData** | [**CompareData**](CompareData.md)| with a document to compare.             | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 

### Return type

[**DocumentResponse**](DocumentResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConvertDocument**
> *os.File ConvertDocument(ctx, document, format, optional)
Converts document from the request's content to the specified format .

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **document** | ***os.File**| Converting document | 
  **format** | **string**| Format to convert. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **document** | ***os.File**| Converting document | 
 **format** | **string**| Format to convert. | 
 **storage** | **string**| Original document storage. | 
 **outPath** | **string**| Path for saving operation result to the local storage. | 
 **fileNameFieldValue** | **string**| This file name will be used when resulting document has dynamic field for document file name {filename}. If it is not set, \&quot;sourceFilename\&quot; will be used instead.  | [default to sourceFilename]
 **fontsLocation** | **string**| Folder in filestorage with custom fonts. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CopyFile**
> CopyFile(ctx, destPath, srcPath, optional)
Copy file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **destPath** | **string**| Destination file path | 
  **srcPath** | **string**| Source file&#39;s path e.g. &#39;/Folder 1/file.ext&#39; or &#39;/Bucket/Folder 1/file.ext&#39; | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destPath** | **string**| Destination file path | 
 **srcPath** | **string**| Source file&#39;s path e.g. &#39;/Folder 1/file.ext&#39; or &#39;/Bucket/Folder 1/file.ext&#39; | 
 **srcStorageName** | **string**| Source storage name | 
 **destStorageName** | **string**| Destination storage name | 
 **versionId** | **string**| File version ID to copy | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CopyFolder**
> CopyFolder(ctx, destPath, srcPath, optional)
Copy folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **destPath** | **string**| Destination folder path e.g. &#39;/dst&#39; | 
  **srcPath** | **string**| Source folder path e.g. /Folder1 | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destPath** | **string**| Destination folder path e.g. &#39;/dst&#39; | 
 **srcPath** | **string**| Source folder path e.g. /Folder1 | 
 **srcStorageName** | **string**| Source storage name | 
 **destStorageName** | **string**| Destination storage name | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDocument**
> DocumentResponse CreateDocument(ctx, optional)
Creates new document. Document is created with format which is recognized from file extensions. Supported extensions: \".doc\", \".docx\", \".docm\", \".dot\", \".dotm\", \".dotx\", \".flatopc\", \".fopc\", \".flatopc_macro\", \".fopc_macro\", \".flatopc_template\", \".fopc_template\", \".flatopc_template_macro\", \".fopc_template_macro\", \".wordml\", \".wml\", \".rtf\".

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storage** | **string**| Original document storage. | 
 **fileName** | **string**| The document name. | 
 **folder** | **string**| The document folder. | 

### Return type

[**DocumentResponse**](DocumentResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFolder**
> CreateFolder(ctx, path, optional)
Create the folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **path** | **string**| Target folder&#39;s path e.g. Folder1/Folder2/. The folders will be created recursively | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| Target folder&#39;s path e.g. Folder1/Folder2/. The folders will be created recursively | 
 **storageName** | **string**| Storage name | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateDocumentProperty**
> DocumentPropertyResponse CreateOrUpdateDocumentProperty(ctx, name, propertyName, property, optional)
Adds new or update existing document property.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **propertyName** | **string**| The property name. | 
  **property** | [**DocumentProperty**](DocumentProperty.md)| The property with new value. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **propertyName** | **string**| The property name. | 
 **property** | [**DocumentProperty**](DocumentProperty.md)| The property with new value. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

[**DocumentPropertyResponse**](DocumentPropertyResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBorder**
> BorderResponse DeleteBorder(ctx, name, nodePath, borderType, optional)
Resets border properties to default values.             

'nodePath' should refer to paragraph, cell or row.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **nodePath** | **string**| Path to the node with border(node should be paragraph, cell or row). | 
  **borderType** | **string**| Border type. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **nodePath** | **string**| Path to the node with border(node should be paragraph, cell or row). | 
 **borderType** | **string**| Border type. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

[**BorderResponse**](BorderResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBorders**
> BordersResponse DeleteBorders(ctx, name, nodePath, optional)
Resets borders properties to default values.             

'nodePath' should refer to paragraph, cell or row.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **nodePath** | **string**| Path to the node with borders(node should be paragraph, cell or row). | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **nodePath** | **string**| Path to the node with borders(node should be paragraph, cell or row). | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

[**BordersResponse**](BordersResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteComment**
> DeleteComment(ctx, name, commentIndex, optional)
Removes comment from document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **commentIndex** | **int32**| The comment index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **commentIndex** | **int32**| The comment index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDocumentProperty**
> DeleteDocumentProperty(ctx, name, propertyName, optional)
Deletes document property.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **propertyName** | **string**| The property name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **propertyName** | **string**| The property name. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDrawingObject**
> DeleteDrawingObject(ctx, name, nodePath, index, optional)
Removes drawing object from document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **nodePath** | **string**| Path to the node, which contains collection of drawing objects. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **nodePath** | **string**| Path to the node, which contains collection of drawing objects. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDrawingObjectWithoutNodePath**
> DeleteDrawingObjectWithoutNodePath(ctx, name, index, optional)
Removes drawing object from document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteField**
> DeleteField(ctx, name, nodePath, index, optional)
Deletes field from document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **nodePath** | **string**| Path to the node, which contains collection of fields. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **nodePath** | **string**| Path to the node, which contains collection of fields. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFieldWithoutNodePath**
> DeleteFieldWithoutNodePath(ctx, name, index, optional)
Deletes field from document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFields**
> DeleteFields(ctx, name, nodePath, optional)
Removes fields from section paragraph.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **nodePath** | **string**| Path to the node, which contains collection of fields. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **nodePath** | **string**| Path to the node, which contains collection of fields. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFieldsWithoutNodePath**
> DeleteFieldsWithoutNodePath(ctx, name, optional)
Removes fields from section paragraph.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFile**
> DeleteFile(ctx, path, optional)
Delete file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **path** | **string**| Path of the file including file name and extension e.g. /Folder1/file.ext | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| Path of the file including file name and extension e.g. /Folder1/file.ext | 
 **storageName** | **string**| Storage name | 
 **versionId** | **string**| File version ID to delete | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFolder**
> DeleteFolder(ctx, path, optional)
Delete folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **path** | **string**| Folder path e.g. /Folder1s | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| Folder path e.g. /Folder1s | 
 **storageName** | **string**| Storage name | 
 **recursive** | **bool**| Enable to delete folders, subfolders and files | [default to false]

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFootnote**
> DeleteFootnote(ctx, name, nodePath, index, optional)
Removes footnote from document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **nodePath** | **string**| Path to the node, which contains collection of footnotes. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **nodePath** | **string**| Path to the node, which contains collection of footnotes. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFootnoteWithoutNodePath**
> DeleteFootnoteWithoutNodePath(ctx, name, index, optional)
Removes footnote from document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFormField**
> DeleteFormField(ctx, name, nodePath, index, optional)
Removes form field from document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **nodePath** | **string**| Path to the node that contains collection of formfields. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **nodePath** | **string**| Path to the node that contains collection of formfields. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFormFieldWithoutNodePath**
> DeleteFormFieldWithoutNodePath(ctx, name, index, optional)
Removes form field from document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHeaderFooter**
> DeleteHeaderFooter(ctx, name, sectionPath, index, optional)
Deletes header/footer from document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **sectionPath** | **string**| Path to parent section. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **sectionPath** | **string**| Path to parent section. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHeadersFooters**
> DeleteHeadersFooters(ctx, name, sectionPath, optional)
Deletes document headers and footers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **sectionPath** | **string**| Path to parent section. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **sectionPath** | **string**| Path to parent section. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 
 **headersFootersTypes** | **string**| List of types of headers and footers. | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMacros**
> DeleteMacros(ctx, name, optional)
Removes macros from document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOfficeMathObject**
> DeleteOfficeMathObject(ctx, name, nodePath, index, optional)
Removes OfficeMath object from document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **nodePath** | **string**| Path to the node, which contains collection of OfficeMath objects. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **nodePath** | **string**| Path to the node, which contains collection of OfficeMath objects. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOfficeMathObjectWithoutNodePath**
> DeleteOfficeMathObjectWithoutNodePath(ctx, name, index, optional)
Removes OfficeMath object from document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteParagraph**
> DeleteParagraph(ctx, name, nodePath, index, optional)
Removes paragraph from section.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The file name. | 
  **nodePath** | **string**| Path to the node which contains paragraphs. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The file name. | 
 **nodePath** | **string**| Path to the node which contains paragraphs. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteParagraphWithoutNodePath**
> DeleteParagraphWithoutNodePath(ctx, name, index, optional)
Removes paragraph from section.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The file name. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The file name. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRun**
> DeleteRun(ctx, name, paragraphPath, index, optional)
Removes run from document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **paragraphPath** | **string**| Path to parent paragraph. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **paragraphPath** | **string**| Path to parent paragraph. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSection**
> DeleteSection(ctx, name, sectionIndex, optional)
Removes section from document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **sectionIndex** | **int32**| Section index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **sectionIndex** | **int32**| Section index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTable**
> DeleteTable(ctx, name, nodePath, index, optional)
Deletes a table.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **nodePath** | **string**| Path to the node, which contains tables. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **nodePath** | **string**| Path to the node, which contains tables. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTableCell**
> DeleteTableCell(ctx, name, tableRowPath, index, optional)
Deletes a table cell.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **tableRowPath** | **string**| Path to table row. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **tableRowPath** | **string**| Path to table row. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTableRow**
> DeleteTableRow(ctx, name, tablePath, index, optional)
Deletes a table row.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **tablePath** | **string**| Path to table. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **tablePath** | **string**| Path to table. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTableWithoutNodePath**
> DeleteTableWithoutNodePath(ctx, name, index, optional)
Deletes a table.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteWatermark**
> DocumentResponse DeleteWatermark(ctx, name, optional)
Deletes watermark (for deleting last watermark from the document).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

[**DocumentResponse**](DocumentResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadFile**
> *os.File DownloadFile(ctx, path, optional)
Download file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **path** | **string**| Path of the file including the file name and extension e.g. /folder1/file.ext | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| Path of the file including the file name and extension e.g. /folder1/file.ext | 
 **storageName** | **string**| Storage name | 
 **versionId** | **string**| File version ID to download | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecuteMailMerge**
> DocumentResponse ExecuteMailMerge(ctx, name, optional)
Executes document mail merge operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **data** | **string**| Mail merge data | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **withRegions** | **bool**| With regions flag. | [default to false]
 **mailMergeDataFile** | **string**| Mail merge data. | 
 **cleanup** | **string**| Clean up options. | 
 **useWholeParagraphAsRegion** | **bool**| Gets or sets a value indicating whether paragraph with TableStart or             TableEnd field should be fully included into mail merge region or particular range between TableStart and TableEnd fields.             The default value is true. | [default to true]
 **destFileName** | **string**| Result name of the document after the operation. If this parameter is omitted then result of the operation will be saved with autogenerated name. | 

### Return type

[**DocumentResponse**](DocumentResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecuteMailMergeOnline**
> *os.File ExecuteMailMergeOnline(ctx, template, data, optional)
Executes document mail merge online.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **template** | ***os.File**| File with template | 
  **data** | ***os.File**| File with mailmerge data | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **template** | ***os.File**| File with template | 
 **data** | ***os.File**| File with mailmerge data | 
 **withRegions** | **bool**| With regions flag. | [default to false]
 **cleanup** | **string**| Clean up options. | 
 **documentFileName** | **string**| This file name will be used when resulting document has dynamic field for document file name {filename}. If it is not setted, \&quot;template\&quot; will be used instead.  | [default to template]

### Return type

[***os.File**](*os.File.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAvailableFonts**
> AvailableFontsResponse GetAvailableFonts(ctx, optional)
Gets the list of fonts, available for document processing.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fontsLocation** | **string**| Folder in filestorage with custom fonts. | 

### Return type

[**AvailableFontsResponse**](AvailableFontsResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBookmarkByName**
> BookmarkResponse GetBookmarkByName(ctx, name, bookmarkName, optional)
Reads document bookmark data by its name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **bookmarkName** | **string**| The bookmark name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **bookmarkName** | **string**| The bookmark name. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**BookmarkResponse**](BookmarkResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBookmarks**
> BookmarksResponse GetBookmarks(ctx, name, optional)
Reads document bookmarks common info.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**BookmarksResponse**](BookmarksResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBorder**
> BorderResponse GetBorder(ctx, name, nodePath, borderType, optional)
Returns a border.

'nodePath' should refer to paragraph, cell or row.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **nodePath** | **string**| Path to the node with border(node should be paragraph, cell or row). | 
  **borderType** | **string**| Border type. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **nodePath** | **string**| Path to the node with border(node should be paragraph, cell or row). | 
 **borderType** | **string**| Border type. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**BorderResponse**](BorderResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBorders**
> BordersResponse GetBorders(ctx, name, nodePath, optional)
Returns a collection of borders.

'nodePath' should refer to paragraph, cell or row.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **nodePath** | **string**| Path to the node with borders (node should be paragraph, cell or row). | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **nodePath** | **string**| Path to the node with borders (node should be paragraph, cell or row). | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**BordersResponse**](BordersResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetComment**
> CommentResponse GetComment(ctx, name, commentIndex, optional)
Gets comment from document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **commentIndex** | **int32**| The comment index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **commentIndex** | **int32**| The comment index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**CommentResponse**](CommentResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetComments**
> CommentsResponse GetComments(ctx, name, optional)
Gets comments from document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**CommentsResponse**](CommentsResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocument**
> DocumentResponse GetDocument(ctx, documentName, optional)
Reads document common info.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **documentName** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **documentName** | **string**| The document name. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**DocumentResponse**](DocumentResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentDrawingObjectByIndex**
> DrawingObjectResponse GetDocumentDrawingObjectByIndex(ctx, name, nodePath, index, optional)
Reads document drawing object common info by its index or convert to format specified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **nodePath** | **string**| Path to the node, which contains collection of drawing objects. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **nodePath** | **string**| Path to the node, which contains collection of drawing objects. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**DrawingObjectResponse**](DrawingObjectResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentDrawingObjectByIndexWithoutNodePath**
> DrawingObjectResponse GetDocumentDrawingObjectByIndexWithoutNodePath(ctx, name, index, optional)
Reads document drawing object common info by its index or convert to format specified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**DrawingObjectResponse**](DrawingObjectResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentDrawingObjectImageData**
> *os.File GetDocumentDrawingObjectImageData(ctx, name, nodePath, index, optional)
Reads drawing object image data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **nodePath** | **string**| Path to the node, which contains collection of drawing objects. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **nodePath** | **string**| Path to the node, which contains collection of drawing objects. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentDrawingObjectImageDataWithoutNodePath**
> *os.File GetDocumentDrawingObjectImageDataWithoutNodePath(ctx, name, index, optional)
Reads drawing object image data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentDrawingObjectOleData**
> *os.File GetDocumentDrawingObjectOleData(ctx, name, nodePath, index, optional)
Gets drawing object OLE data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **nodePath** | **string**| Path to the node, which contains collection of drawing objects. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **nodePath** | **string**| Path to the node, which contains collection of drawing objects. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentDrawingObjectOleDataWithoutNodePath**
> *os.File GetDocumentDrawingObjectOleDataWithoutNodePath(ctx, name, index, optional)
Gets drawing object OLE data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentDrawingObjects**
> DrawingObjectsResponse GetDocumentDrawingObjects(ctx, name, nodePath, optional)
Reads document drawing objects common info.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **nodePath** | **string**| Path to the node, which contains collection of drawing objects. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **nodePath** | **string**| Path to the node, which contains collection of drawing objects. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**DrawingObjectsResponse**](DrawingObjectsResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentDrawingObjectsWithoutNodePath**
> DrawingObjectsResponse GetDocumentDrawingObjectsWithoutNodePath(ctx, name, optional)
Reads document drawing objects common info.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**DrawingObjectsResponse**](DrawingObjectsResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentFieldNames**
> FieldNamesResponse GetDocumentFieldNames(ctx, name, optional)
Reads document field names.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **useNonMergeFields** | **bool**| If true, result includes \&quot;mustache\&quot; field names. | [default to false]

### Return type

[**FieldNamesResponse**](FieldNamesResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentFieldNamesOnline**
> FieldNamesResponse GetDocumentFieldNamesOnline(ctx, template, optional)
Reads document field names.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **template** | ***os.File**| File with template | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **template** | ***os.File**| File with template | 
 **useNonMergeFields** | **bool**| Use non merge fields or not. | [default to false]

### Return type

[**FieldNamesResponse**](FieldNamesResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentHyperlinkByIndex**
> HyperlinkResponse GetDocumentHyperlinkByIndex(ctx, name, hyperlinkIndex, optional)
Reads document hyperlink by its index.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **hyperlinkIndex** | **int32**| The hyperlink index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **hyperlinkIndex** | **int32**| The hyperlink index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**HyperlinkResponse**](HyperlinkResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentHyperlinks**
> HyperlinksResponse GetDocumentHyperlinks(ctx, name, optional)
Reads document hyperlinks common info.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**HyperlinksResponse**](HyperlinksResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentProperties**
> DocumentPropertiesResponse GetDocumentProperties(ctx, name, optional)
Reads document properties info.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document&#39;s name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document&#39;s name. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**DocumentPropertiesResponse**](DocumentPropertiesResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentProperty**
> DocumentPropertyResponse GetDocumentProperty(ctx, name, propertyName, optional)
Reads document property info by the property name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **propertyName** | **string**| The property name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **propertyName** | **string**| The property name. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**DocumentPropertyResponse**](DocumentPropertyResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentProtection**
> ProtectionDataResponse GetDocumentProtection(ctx, name, optional)
Reads document protection common info.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**ProtectionDataResponse**](ProtectionDataResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentStatistics**
> StatDataResponse GetDocumentStatistics(ctx, name, optional)
Reads document statistics.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **includeComments** | **bool**| Support including/excluding comments from the WordCount. Default value is \&quot;false\&quot;. | [default to false]
 **includeFootnotes** | **bool**| Support including/excluding footnotes from the WordCount. Default value is \&quot;false\&quot;. | [default to false]
 **includeTextInShapes** | **bool**| Support including/excluding shape&#39;s text from the WordCount. Default value is \&quot;false\&quot;. | [default to false]

### Return type

[**StatDataResponse**](StatDataResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentWithFormat**
> *os.File GetDocumentWithFormat(ctx, name, format, optional)
Exports the document into the specified format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **format** | **string**| The destination format. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **format** | **string**| The destination format. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **outPath** | **string**| Path to save the result. | 
 **fontsLocation** | **string**| Folder in filestorage with custom fonts. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetField**
> FieldResponse GetField(ctx, name, nodePath, index, optional)
Gets field from document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **nodePath** | **string**| Path to the node, which contains collection of fields. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **nodePath** | **string**| Path to the node, which contains collection of fields. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**FieldResponse**](FieldResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFieldWithoutNodePath**
> FieldResponse GetFieldWithoutNodePath(ctx, name, index, optional)
Gets field from document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**FieldResponse**](FieldResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFields**
> FieldsResponse GetFields(ctx, name, nodePath, optional)
Get fields from document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **nodePath** | **string**| Path to the node, which contains collection of fields. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **nodePath** | **string**| Path to the node, which contains collection of fields. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**FieldsResponse**](FieldsResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFieldsWithoutNodePath**
> FieldsResponse GetFieldsWithoutNodePath(ctx, name, optional)
Get fields from document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**FieldsResponse**](FieldsResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFilesList**
> FilesList GetFilesList(ctx, path, optional)
Get all files and folders within a folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **path** | **string**| Folder path e.g. /Folder1 | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| Folder path e.g. /Folder1 | 
 **storageName** | **string**| Storage name | 

### Return type

[**FilesList**](FilesList.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFootnote**
> FootnoteResponse GetFootnote(ctx, name, nodePath, index, optional)
Reads footnote by index.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **nodePath** | **string**| Path to the node, which contains collection of footnotes. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **nodePath** | **string**| Path to the node, which contains collection of footnotes. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**FootnoteResponse**](FootnoteResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFootnoteWithoutNodePath**
> FootnoteResponse GetFootnoteWithoutNodePath(ctx, name, index, optional)
Reads footnote by index.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**FootnoteResponse**](FootnoteResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFootnotes**
> FootnotesResponse GetFootnotes(ctx, name, nodePath, optional)
Gets footnotes from document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **nodePath** | **string**| Path to the node, which contains collection of footnotes. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **nodePath** | **string**| Path to the node, which contains collection of footnotes. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**FootnotesResponse**](FootnotesResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFootnotesWithoutNodePath**
> FootnotesResponse GetFootnotesWithoutNodePath(ctx, name, optional)
Gets footnotes from document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**FootnotesResponse**](FootnotesResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFormField**
> FormFieldResponse GetFormField(ctx, name, nodePath, index, optional)
Returns representation of an one of the form field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **nodePath** | **string**| Path to the node that contains collection of formfields. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **nodePath** | **string**| Path to the node that contains collection of formfields. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**FormFieldResponse**](FormFieldResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFormFieldWithoutNodePath**
> FormFieldResponse GetFormFieldWithoutNodePath(ctx, name, index, optional)
Returns representation of an one of the form field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**FormFieldResponse**](FormFieldResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFormFields**
> FormFieldsResponse GetFormFields(ctx, name, nodePath, optional)
Gets form fields from document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **nodePath** | **string**| Path to the node containing collection of form fields. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **nodePath** | **string**| Path to the node containing collection of form fields. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**FormFieldsResponse**](FormFieldsResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFormFieldsWithoutNodePath**
> FormFieldsResponse GetFormFieldsWithoutNodePath(ctx, name, optional)
Gets form fields from document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**FormFieldsResponse**](FormFieldsResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHeaderFooter**
> HeaderFooterResponse GetHeaderFooter(ctx, name, headerFooterIndex, optional)
Returns a header/footer from the document by index.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **headerFooterIndex** | **int32**| Header/footer index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **headerFooterIndex** | **int32**| Header/footer index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **filterByType** | **string**| List of types of headers and footers. | 

### Return type

[**HeaderFooterResponse**](HeaderFooterResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHeaderFooterOfSection**
> HeaderFooterResponse GetHeaderFooterOfSection(ctx, name, headerFooterIndex, sectionIndex, optional)
Returns a header/footer from the document section.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **headerFooterIndex** | **int32**| Header/footer index. | 
  **sectionIndex** | **int32**| Section index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **headerFooterIndex** | **int32**| Header/footer index. | 
 **sectionIndex** | **int32**| Section index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **filterByType** | **string**| List of types of headers and footers. | 

### Return type

[**HeaderFooterResponse**](HeaderFooterResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHeaderFooters**
> HeaderFootersResponse GetHeaderFooters(ctx, name, sectionPath, optional)
Returns a list of header/footers from the document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **sectionPath** | **string**| Path to parent section. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **sectionPath** | **string**| Path to parent section. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **filterByType** | **string**| List of types of headers and footers. | 

### Return type

[**HeaderFootersResponse**](HeaderFootersResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOfficeMathObject**
> OfficeMathObjectResponse GetOfficeMathObject(ctx, name, nodePath, index, optional)
Reads OfficeMath object by index.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **nodePath** | **string**| Path to the node, which contains collection of OfficeMath objects. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **nodePath** | **string**| Path to the node, which contains collection of OfficeMath objects. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**OfficeMathObjectResponse**](OfficeMathObjectResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOfficeMathObjectWithoutNodePath**
> OfficeMathObjectResponse GetOfficeMathObjectWithoutNodePath(ctx, name, index, optional)
Reads OfficeMath object by index.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**OfficeMathObjectResponse**](OfficeMathObjectResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOfficeMathObjects**
> OfficeMathObjectsResponse GetOfficeMathObjects(ctx, name, nodePath, optional)
Gets OfficeMath objects from document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **nodePath** | **string**| Path to the node, which contains collection of OfficeMath objects. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **nodePath** | **string**| Path to the node, which contains collection of OfficeMath objects. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**OfficeMathObjectsResponse**](OfficeMathObjectsResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOfficeMathObjectsWithoutNodePath**
> OfficeMathObjectsResponse GetOfficeMathObjectsWithoutNodePath(ctx, name, optional)
Gets OfficeMath objects from document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**OfficeMathObjectsResponse**](OfficeMathObjectsResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetParagraph**
> ParagraphResponse GetParagraph(ctx, name, nodePath, index, optional)
This resource represents one of the paragraphs contained in the document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **nodePath** | **string**| Path to the node which contains paragraphs. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **nodePath** | **string**| Path to the node which contains paragraphs. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**ParagraphResponse**](ParagraphResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetParagraphFormat**
> ParagraphFormatResponse GetParagraphFormat(ctx, name, nodePath, index, optional)
Represents all the formatting for a paragraph.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **nodePath** | **string**| Path to the node which contains paragraphs. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **nodePath** | **string**| Path to the node which contains paragraphs. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**ParagraphFormatResponse**](ParagraphFormatResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetParagraphFormatWithoutNodePath**
> ParagraphFormatResponse GetParagraphFormatWithoutNodePath(ctx, name, index, optional)
Represents all the formatting for a paragraph.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**ParagraphFormatResponse**](ParagraphFormatResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetParagraphWithoutNodePath**
> ParagraphResponse GetParagraphWithoutNodePath(ctx, name, index, optional)
This resource represents one of the paragraphs contained in the document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**ParagraphResponse**](ParagraphResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetParagraphs**
> ParagraphLinkCollectionResponse GetParagraphs(ctx, name, nodePath, optional)
Returns a list of paragraphs that are contained in the document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **nodePath** | **string**| Path to the node which contains paragraphs. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **nodePath** | **string**| Path to the node which contains paragraphs. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**ParagraphLinkCollectionResponse**](ParagraphLinkCollectionResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetParagraphsWithoutNodePath**
> ParagraphLinkCollectionResponse GetParagraphsWithoutNodePath(ctx, name, optional)
Returns a list of paragraphs that are contained in the document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**ParagraphLinkCollectionResponse**](ParagraphLinkCollectionResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRangeText**
> RangeTextResponse GetRangeText(ctx, name, rangeStartIdentifier, rangeEndIdentifier, optional)
Gets the text from the range.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document. | 
  **rangeStartIdentifier** | **string**| The range start identifier. Identifier is the value of the \&quot;nodeId\&quot; field, which every document node has, extended with the prefix \&quot;id\&quot;. It looks like \&quot;id0.0.7\&quot;. Also values like \&quot;image5\&quot; and \&quot;table3\&quot; can be used as an identifier for images and tables, where the number is an index of the image/table. | 
  **rangeEndIdentifier** | **string**| The range end identifier. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document. | 
 **rangeStartIdentifier** | **string**| The range start identifier. Identifier is the value of the \&quot;nodeId\&quot; field, which every document node has, extended with the prefix \&quot;id\&quot;. It looks like \&quot;id0.0.7\&quot;. Also values like \&quot;image5\&quot; and \&quot;table3\&quot; can be used as an identifier for images and tables, where the number is an index of the image/table. | 
 **rangeEndIdentifier** | **string**| The range end identifier. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**RangeTextResponse**](RangeTextResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRun**
> RunResponse GetRun(ctx, name, paragraphPath, index, optional)
This resource represents run of text contained in the document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **paragraphPath** | **string**| Path to parent paragraph. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **paragraphPath** | **string**| Path to parent paragraph. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**RunResponse**](RunResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRunFont**
> FontResponse GetRunFont(ctx, name, paragraphPath, index, optional)
This resource represents font of run.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **paragraphPath** | **string**| Path to parent paragraph. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **paragraphPath** | **string**| Path to parent paragraph. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**FontResponse**](FontResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRuns**
> RunsResponse GetRuns(ctx, name, paragraphPath, optional)
This resource represents collection of runs in the paragraph.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **paragraphPath** | **string**| Path to parent paragraph. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **paragraphPath** | **string**| Path to parent paragraph. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**RunsResponse**](RunsResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSection**
> SectionResponse GetSection(ctx, name, sectionIndex, optional)
Gets document section by index.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **sectionIndex** | **int32**| Section index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **sectionIndex** | **int32**| Section index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**SectionResponse**](SectionResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSectionPageSetup**
> SectionPageSetupResponse GetSectionPageSetup(ctx, name, sectionIndex, optional)
Gets page setup of section.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **sectionIndex** | **int32**| Section index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **sectionIndex** | **int32**| Section index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**SectionPageSetupResponse**](SectionPageSetupResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSections**
> SectionLinkCollectionResponse GetSections(ctx, name, optional)
Returns a list of sections that are contained in the document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**SectionLinkCollectionResponse**](SectionLinkCollectionResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTable**
> TableResponse GetTable(ctx, name, nodePath, index, optional)
Returns a table.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **nodePath** | **string**| Path to the node, which contains tables. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **nodePath** | **string**| Path to the node, which contains tables. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**TableResponse**](TableResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTableCell**
> TableCellResponse GetTableCell(ctx, name, tableRowPath, index, optional)
Returns a table cell.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **tableRowPath** | **string**| Path to table row. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **tableRowPath** | **string**| Path to table row. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**TableCellResponse**](TableCellResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTableCellFormat**
> TableCellFormatResponse GetTableCellFormat(ctx, name, tableRowPath, index, optional)
Returns a table cell format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **tableRowPath** | **string**| Path to table row. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **tableRowPath** | **string**| Path to table row. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**TableCellFormatResponse**](TableCellFormatResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTableProperties**
> TablePropertiesResponse GetTableProperties(ctx, name, nodePath, index, optional)
Returns a table properties.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **nodePath** | **string**| Path to the node, which contains tables. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **nodePath** | **string**| Path to the node, which contains tables. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**TablePropertiesResponse**](TablePropertiesResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTablePropertiesWithoutNodePath**
> TablePropertiesResponse GetTablePropertiesWithoutNodePath(ctx, name, index, optional)
Returns a table properties.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**TablePropertiesResponse**](TablePropertiesResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTableRow**
> TableRowResponse GetTableRow(ctx, name, tablePath, index, optional)
Returns a table row.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **tablePath** | **string**| Path to table. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **tablePath** | **string**| Path to table. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**TableRowResponse**](TableRowResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTableRowFormat**
> TableRowFormatResponse GetTableRowFormat(ctx, name, tablePath, index, optional)
Returns a table row format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **tablePath** | **string**| Path to table. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **tablePath** | **string**| Path to table. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**TableRowFormatResponse**](TableRowFormatResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTableWithoutNodePath**
> TableResponse GetTableWithoutNodePath(ctx, name, index, optional)
Returns a table.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**TableResponse**](TableResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTables**
> TableLinkCollectionResponse GetTables(ctx, name, nodePath, optional)
Returns a list of tables that are contained in the document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **nodePath** | **string**| Path to the node, which contains tables. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **nodePath** | **string**| Path to the node, which contains tables. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**TableLinkCollectionResponse**](TableLinkCollectionResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTablesWithoutNodePath**
> TableLinkCollectionResponse GetTablesWithoutNodePath(ctx, name, optional)
Returns a list of tables that are contained in the document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**TableLinkCollectionResponse**](TableLinkCollectionResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsertComment**
> CommentResponse InsertComment(ctx, name, comment, optional)
Adds comment to document, returns inserted comment data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **comment** | [**Comment**](Comment.md)| The comment data. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **comment** | [**Comment**](Comment.md)| The comment data. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

[**CommentResponse**](CommentResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsertDrawingObject**
> DrawingObjectResponse InsertDrawingObject(ctx, name, drawingObject, imageFile, nodePath, optional)
Adds drawing object to document, returns added  drawing object's data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **drawingObject** | **string**| Drawing object parameters | 
  **imageFile** | ***os.File**| File with image | 
  **nodePath** | **string**| Path to the node, which contains collection of drawing objects. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **drawingObject** | **string**| Drawing object parameters | 
 **imageFile** | ***os.File**| File with image | 
 **nodePath** | **string**| Path to the node, which contains collection of drawing objects. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

[**DrawingObjectResponse**](DrawingObjectResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsertDrawingObjectWithoutNodePath**
> DrawingObjectResponse InsertDrawingObjectWithoutNodePath(ctx, name, drawingObject, imageFile, optional)
Adds drawing object to document, returns added  drawing object's data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **drawingObject** | **string**| Drawing object parameters | 
  **imageFile** | ***os.File**| File with image | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **drawingObject** | **string**| Drawing object parameters | 
 **imageFile** | ***os.File**| File with image | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

[**DrawingObjectResponse**](DrawingObjectResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsertField**
> FieldResponse InsertField(ctx, name, field, nodePath, optional)
Adds field to document, returns inserted field's data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **field** | [**Field**](Field.md)| Field data. | 
  **nodePath** | **string**| Path to the node, which contains collection of fields. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **field** | [**Field**](Field.md)| Field data. | 
 **nodePath** | **string**| Path to the node, which contains collection of fields. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 
 **insertBeforeNode** | **string**| Field will be inserted before node with id&#x3D;\&quot;nodeId\&quot;. | 

### Return type

[**FieldResponse**](FieldResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsertFieldWithoutNodePath**
> FieldResponse InsertFieldWithoutNodePath(ctx, name, field, optional)
Adds field to document, returns inserted field's data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **field** | [**Field**](Field.md)| Field data. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **field** | [**Field**](Field.md)| Field data. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 
 **insertBeforeNode** | **string**| Field will be inserted before node with id&#x3D;\&quot;nodeId\&quot;. | 

### Return type

[**FieldResponse**](FieldResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsertFootnote**
> FootnoteResponse InsertFootnote(ctx, name, footnoteDto, nodePath, optional)
Adds footnote to document, returns added footnote's data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **footnoteDto** | [**Footnote**](Footnote.md)| Footnote data. | 
  **nodePath** | **string**| Path to the node, which contains collection of footnotes. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **footnoteDto** | [**Footnote**](Footnote.md)| Footnote data. | 
 **nodePath** | **string**| Path to the node, which contains collection of footnotes. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

[**FootnoteResponse**](FootnoteResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsertFootnoteWithoutNodePath**
> FootnoteResponse InsertFootnoteWithoutNodePath(ctx, name, footnoteDto, optional)
Adds footnote to document, returns added footnote's data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **footnoteDto** | [**Footnote**](Footnote.md)| Footnote data. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **footnoteDto** | [**Footnote**](Footnote.md)| Footnote data. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

[**FootnoteResponse**](FootnoteResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsertFormField**
> FormFieldResponse InsertFormField(ctx, name, formField, nodePath, optional)
Adds form field to paragraph, returns added form field's data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **formField** | [**FormField**](FormField.md)| From field data. | 
  **nodePath** | **string**| Path to the node that contains collection of formfields. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **formField** | [**FormField**](FormField.md)| From field data. | 
 **nodePath** | **string**| Path to the node that contains collection of formfields. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 
 **insertBeforeNode** | **string**| Form field will be inserted before node with index. | 

### Return type

[**FormFieldResponse**](FormFieldResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsertFormFieldWithoutNodePath**
> FormFieldResponse InsertFormFieldWithoutNodePath(ctx, name, formField, optional)
Adds form field to paragraph, returns added form field's data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **formField** | [**FormField**](FormField.md)| From field data. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **formField** | [**FormField**](FormField.md)| From field data. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 
 **insertBeforeNode** | **string**| Form field will be inserted before node with index. | 

### Return type

[**FormFieldResponse**](FormFieldResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsertHeaderFooter**
> HeaderFooterResponse InsertHeaderFooter(ctx, name, headerFooterType, sectionPath, optional)
Inserts to document header or footer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **headerFooterType** | **string**| Type of header/footer. | 
  **sectionPath** | **string**| Path to parent section. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **headerFooterType** | **string**| Type of header/footer. | 
 **sectionPath** | **string**| Path to parent section. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

[**HeaderFooterResponse**](HeaderFooterResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsertPageNumbers**
> DocumentResponse InsertPageNumbers(ctx, name, pageNumber, optional)
Inserts document page numbers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| A document name. | 
  **pageNumber** | [**PageNumber**](PageNumber.md)| with the page numbers settings. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| A document name. | 
 **pageNumber** | [**PageNumber**](PageNumber.md)| with the page numbers settings. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

[**DocumentResponse**](DocumentResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsertParagraph**
> ParagraphResponse InsertParagraph(ctx, name, paragraph, nodePath, optional)
Adds paragraph to document, returns added paragraph's data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **paragraph** | [**ParagraphInsert**](ParagraphInsert.md)| Paragraph data. | 
  **nodePath** | **string**| Path to the node which contains paragraphs. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **paragraph** | [**ParagraphInsert**](ParagraphInsert.md)| Paragraph data. | 
 **nodePath** | **string**| Path to the node which contains paragraphs. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 
 **insertBeforeNode** | **string**| Paragraph will be inserted before node with index. | 

### Return type

[**ParagraphResponse**](ParagraphResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsertRun**
> RunResponse InsertRun(ctx, name, paragraphPath, run, optional)
Adds run to document, returns added paragraph's data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **paragraphPath** | **string**| Path to parent paragraph. | 
  **run** | [**Run**](Run.md)| Run data. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **paragraphPath** | **string**| Path to parent paragraph. | 
 **run** | [**Run**](Run.md)| Run data. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 
 **insertBeforeNode** | **string**| Paragraph will be inserted before node with index. | 

### Return type

[**RunResponse**](RunResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsertTable**
> TableResponse InsertTable(ctx, name, nodePath, optional)
Adds table to document, returns added table's data.             

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **nodePath** | **string**| Path to the node, which contains tables. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **nodePath** | **string**| Path to the node, which contains tables. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 
 **table** | [**TableInsert**](TableInsert.md)| Table parameters/. | 

### Return type

[**TableResponse**](TableResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsertTableCell**
> TableCellResponse InsertTableCell(ctx, name, tableRowPath, optional)
Adds table cell to table, returns added cell's data.             

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **tableRowPath** | **string**| Path to table row. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **tableRowPath** | **string**| Path to table row. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 
 **cell** | [**TableCellInsert**](TableCellInsert.md)| Table cell parameters/. | 

### Return type

[**TableCellResponse**](TableCellResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsertTableRow**
> TableRowResponse InsertTableRow(ctx, name, tablePath, optional)
Adds table row to table, returns added row's data.             

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **tablePath** | **string**| Path to table. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **tablePath** | **string**| Path to table. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 
 **row** | [**TableRowInsert**](TableRowInsert.md)| Table row parameters/. | 

### Return type

[**TableRowResponse**](TableRowResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsertTableWithoutNodePath**
> TableResponse InsertTableWithoutNodePath(ctx, name, optional)
Adds table to document, returns added table's data.             

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 
 **table** | [**TableInsert**](TableInsert.md)| Table parameters/. | 

### Return type

[**TableResponse**](TableResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsertWatermarkImage**
> DocumentResponse InsertWatermarkImage(ctx, name, optional)
Inserts document watermark image.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **imageFile** | ***os.File**| File with image | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 
 **rotationAngle** | **float64**| The watermark rotation angle. | [default to 0.0]
 **image** | **string**| The image file server full name. If the name is empty the image is expected in request content. | 

### Return type

[**DocumentResponse**](DocumentResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsertWatermarkText**
> DocumentResponse InsertWatermarkText(ctx, name, watermarkText, optional)
Inserts document watermark text.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **watermarkText** | [**WatermarkText**](WatermarkText.md)| with the watermark data.             | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **watermarkText** | [**WatermarkText**](WatermarkText.md)| with the watermark data.             | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

[**DocumentResponse**](DocumentResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadWebDocument**
> SaveResponse LoadWebDocument(ctx, data, optional)
Loads new document from web into the file with any supported format of data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **data** | [**LoadWebDocumentData**](LoadWebDocumentData.md)| Parameters of loading. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**LoadWebDocumentData**](LoadWebDocumentData.md)| Parameters of loading. | 
 **storage** | **string**| Original document storage. | 

### Return type

[**SaveResponse**](SaveResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MoveFile**
> MoveFile(ctx, destPath, srcPath, optional)
Move file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **destPath** | **string**| Destination file path e.g. &#39;/dest.ext&#39; | 
  **srcPath** | **string**| Source file&#39;s path e.g. &#39;/Folder 1/file.ext&#39; or &#39;/Bucket/Folder 1/file.ext&#39; | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destPath** | **string**| Destination file path e.g. &#39;/dest.ext&#39; | 
 **srcPath** | **string**| Source file&#39;s path e.g. &#39;/Folder 1/file.ext&#39; or &#39;/Bucket/Folder 1/file.ext&#39; | 
 **srcStorageName** | **string**| Source storage name | 
 **destStorageName** | **string**| Destination storage name | 
 **versionId** | **string**| File version ID to move | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MoveFolder**
> MoveFolder(ctx, destPath, srcPath, optional)
Move folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **destPath** | **string**| Destination folder path to move to e.g &#39;/dst&#39; | 
  **srcPath** | **string**| Source folder path e.g. /Folder1 | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destPath** | **string**| Destination folder path to move to e.g &#39;/dst&#39; | 
 **srcPath** | **string**| Source folder path e.g. /Folder1 | 
 **srcStorageName** | **string**| Source storage name | 
 **destStorageName** | **string**| Destination storage name | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProtectDocument**
> ProtectionDataResponse ProtectDocument(ctx, name, protectionRequest, optional)
Protects document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **protectionRequest** | [**ProtectionRequest**](ProtectionRequest.md)| with protection settings.             | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **protectionRequest** | [**ProtectionRequest**](ProtectionRequest.md)| with protection settings.             | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 

### Return type

[**ProtectionDataResponse**](ProtectionDataResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RejectAllRevisions**
> RevisionsModificationResponse RejectAllRevisions(ctx, name, optional)
Rejects all revisions in document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 

### Return type

[**RevisionsModificationResponse**](RevisionsModificationResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveRange**
> DocumentResponse RemoveRange(ctx, name, rangeStartIdentifier, rangeEndIdentifier, optional)
Removes the range from the document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document. | 
  **rangeStartIdentifier** | **string**| The range start identifier. Identifier is the value of the \&quot;nodeId\&quot; field, which every document node has, extended with the prefix \&quot;id\&quot;. It looks like \&quot;id0.0.7\&quot;. Also values like \&quot;image5\&quot; and \&quot;table3\&quot; can be used as an identifier for images and tables, where the number is an index of the image/table. | 
  **rangeEndIdentifier** | **string**| The range end identifier. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document. | 
 **rangeStartIdentifier** | **string**| The range start identifier. Identifier is the value of the \&quot;nodeId\&quot; field, which every document node has, extended with the prefix \&quot;id\&quot;. It looks like \&quot;id0.0.7\&quot;. Also values like \&quot;image5\&quot; and \&quot;table3\&quot; can be used as an identifier for images and tables, where the number is an index of the image/table. | 
 **rangeEndIdentifier** | **string**| The range end identifier. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 

### Return type

[**DocumentResponse**](DocumentResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenderDrawingObject**
> *os.File RenderDrawingObject(ctx, name, format, nodePath, index, optional)
Renders drawing object to specified format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **format** | **string**| The destination format. | 
  **nodePath** | **string**| Path to the node, which contains drawing objects. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **format** | **string**| The destination format. | 
 **nodePath** | **string**| Path to the node, which contains drawing objects. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **fontsLocation** | **string**| Folder in filestorage with custom fonts. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenderDrawingObjectWithoutNodePath**
> *os.File RenderDrawingObjectWithoutNodePath(ctx, name, format, index, optional)
Renders drawing object to specified format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **format** | **string**| The destination format. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **format** | **string**| The destination format. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **fontsLocation** | **string**| Folder in filestorage with custom fonts. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenderMathObject**
> *os.File RenderMathObject(ctx, name, format, nodePath, index, optional)
Renders math object to specified format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **format** | **string**| The destination format. | 
  **nodePath** | **string**| Path to the node, which contains office math objects. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **format** | **string**| The destination format. | 
 **nodePath** | **string**| Path to the node, which contains office math objects. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **fontsLocation** | **string**| Folder in filestorage with custom fonts. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenderMathObjectWithoutNodePath**
> *os.File RenderMathObjectWithoutNodePath(ctx, name, format, index, optional)
Renders math object to specified format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **format** | **string**| The destination format. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **format** | **string**| The destination format. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **fontsLocation** | **string**| Folder in filestorage with custom fonts. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenderPage**
> *os.File RenderPage(ctx, name, pageIndex, format, optional)
Renders page to specified format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageIndex** | **int32**| Comment index. | 
  **format** | **string**| The destination format. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageIndex** | **int32**| Comment index. | 
 **format** | **string**| The destination format. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **fontsLocation** | **string**| Folder in filestorage with custom fonts. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenderParagraph**
> *os.File RenderParagraph(ctx, name, format, nodePath, index, optional)
Renders paragraph to specified format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **format** | **string**| The destination format. | 
  **nodePath** | **string**| Path to the node, which contains paragraphs. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **format** | **string**| The destination format. | 
 **nodePath** | **string**| Path to the node, which contains paragraphs. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **fontsLocation** | **string**| Folder in filestorage with custom fonts. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenderParagraphWithoutNodePath**
> *os.File RenderParagraphWithoutNodePath(ctx, name, format, index, optional)
Renders paragraph to specified format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **format** | **string**| The destination format. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **format** | **string**| The destination format. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **fontsLocation** | **string**| Folder in filestorage with custom fonts. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenderTable**
> *os.File RenderTable(ctx, name, format, nodePath, index, optional)
Renders table to specified format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **format** | **string**| The destination format. | 
  **nodePath** | **string**| Path to the node, which contains tables. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **format** | **string**| The destination format. | 
 **nodePath** | **string**| Path to the node, which contains tables. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **fontsLocation** | **string**| Folder in filestorage with custom fonts. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenderTableWithoutNodePath**
> *os.File RenderTableWithoutNodePath(ctx, name, format, index, optional)
Renders table to specified format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **format** | **string**| The destination format. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **format** | **string**| The destination format. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **fontsLocation** | **string**| Folder in filestorage with custom fonts. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceText**
> ReplaceTextResponse ReplaceText(ctx, name, replaceText, optional)
Replaces document text.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **replaceText** | [**ReplaceTextParameters**](ReplaceTextParameters.md)| with the replace operation settings.             | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **replaceText** | [**ReplaceTextParameters**](ReplaceTextParameters.md)| with the replace operation settings.             | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

[**ReplaceTextResponse**](ReplaceTextResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceWithText**
> DocumentResponse ReplaceWithText(ctx, name, rangeStartIdentifier, rangeText, rangeEndIdentifier, optional)
Replaces the content in the range.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document. | 
  **rangeStartIdentifier** | **string**| The range start identifier. Identifier is the value of the \&quot;nodeId\&quot; field, which every document node has, extended with the prefix \&quot;id\&quot;. It looks like \&quot;id0.0.7\&quot;. Also values like \&quot;image5\&quot; and \&quot;table3\&quot; can be used as an identifier for images and tables, where the number is an index of the image/table. | 
  **rangeText** | [**ReplaceRange**](ReplaceRange.md)| Model with text for replacement. | 
  **rangeEndIdentifier** | **string**| The range end identifier. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document. | 
 **rangeStartIdentifier** | **string**| The range start identifier. Identifier is the value of the \&quot;nodeId\&quot; field, which every document node has, extended with the prefix \&quot;id\&quot;. It looks like \&quot;id0.0.7\&quot;. Also values like \&quot;image5\&quot; and \&quot;table3\&quot; can be used as an identifier for images and tables, where the number is an index of the image/table. | 
 **rangeText** | [**ReplaceRange**](ReplaceRange.md)| Model with text for replacement. | 
 **rangeEndIdentifier** | **string**| The range end identifier. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 

### Return type

[**DocumentResponse**](DocumentResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetCache**
> ResetCache(ctx, )
Resets font's cache.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SaveAs**
> SaveResponse SaveAs(ctx, name, saveOptionsData, optional)
Converts document to destination format with detailed settings and saves result to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **saveOptionsData** | [**SaveOptionsData**](SaveOptionsData.md)| Save options. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **saveOptionsData** | [**SaveOptionsData**](SaveOptionsData.md)| Save options. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **fontsLocation** | **string**| Folder in filestorage with custom fonts. | 

### Return type

[**SaveResponse**](SaveResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SaveAsRange**
> DocumentResponse SaveAsRange(ctx, name, rangeStartIdentifier, documentParameters, rangeEndIdentifier, optional)
Saves the selected range as a new document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document. | 
  **rangeStartIdentifier** | **string**| The range start identifier. Identifier is the value of the \&quot;nodeId\&quot; field, which every document node has, extended with the prefix \&quot;id\&quot;. It looks like \&quot;id0.0.7\&quot;. Also values like \&quot;image5\&quot; and \&quot;table3\&quot; can be used as an identifier for images and tables, where the number is an index of the image/table. | 
  **documentParameters** | [**RangeDocument**](RangeDocument.md)| Parameters of a new document. | 
  **rangeEndIdentifier** | **string**| The range end identifier. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document. | 
 **rangeStartIdentifier** | **string**| The range start identifier. Identifier is the value of the \&quot;nodeId\&quot; field, which every document node has, extended with the prefix \&quot;id\&quot;. It looks like \&quot;id0.0.7\&quot;. Also values like \&quot;image5\&quot; and \&quot;table3\&quot; can be used as an identifier for images and tables, where the number is an index of the image/table. | 
 **documentParameters** | [**RangeDocument**](RangeDocument.md)| Parameters of a new document. | 
 **rangeEndIdentifier** | **string**| The range end identifier. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**DocumentResponse**](DocumentResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SaveAsTiff**
> SaveResponse SaveAsTiff(ctx, name, saveOptions, optional)
Converts document to tiff with detailed settings and saves result to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **saveOptions** | [**TiffSaveOptionsData**](TiffSaveOptionsData.md)| Tiff save options. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **saveOptions** | [**TiffSaveOptionsData**](TiffSaveOptionsData.md)| Tiff save options. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **useAntiAliasing** | **bool**| Use antialiasing flag. | 
 **useHighQualityRendering** | **bool**| Use high quality flag. | 
 **imageBrightness** | **float64**| Brightness for the generated images. | 
 **imageColorMode** | **string**| Color mode for the generated images. | 
 **imageContrast** | **float64**| The contrast for the generated images. | 
 **numeralFormat** | **string**| The images numeral format. | 
 **pageCount** | **int32**| Number of pages to render. | 
 **pageIndex** | **int32**| Page index to start rendering. | 
 **paperColor** | **string**| Background image color. | 
 **pixelFormat** | **string**| The pixel format of generated images. | 
 **resolution** | **float64**| The resolution of generated images. | 
 **scale** | **float64**| Zoom factor for generated images. | 
 **tiffCompression** | **string**| The compression tipe. | 
 **dmlRenderingMode** | **string**| Optional, default is Fallback. | 
 **dmlEffectsRenderingMode** | **string**| Optional, default is Simplified. | 
 **tiffBinarizationMethod** | **string**| Optional, Tiff binarization method, possible values are: FloydSteinbergDithering, Threshold. | 
 **zipOutput** | **bool**| Optional. A value determining zip output or not. | 
 **fontsLocation** | **string**| Folder in filestorage with custom fonts. | 

### Return type

[**SaveResponse**](SaveResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Search**
> SearchResponse Search(ctx, name, pattern, optional)
Searches text in document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pattern** | **string**| The regular expression used to find matches. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pattern** | **string**| The regular expression used to find matches. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SplitDocument**
> SplitDocumentResponse SplitDocument(ctx, name, optional)
Splits document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| Original document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Original document name. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **format** | **string**| Format to split. | 
 **from** | **int32**| Start page. | 
 **to** | **int32**| End page. | 
 **zipOutput** | **bool**| ZipOutput or not. | [default to false]
 **fontsLocation** | **string**| Folder in filestorage with custom fonts. | 

### Return type

[**SplitDocumentResponse**](SplitDocumentResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnprotectDocument**
> ProtectionDataResponse UnprotectDocument(ctx, name, protectionRequest, optional)
Unprotects document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **protectionRequest** | [**ProtectionRequest**](ProtectionRequest.md)| with protection settings.             | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **protectionRequest** | [**ProtectionRequest**](ProtectionRequest.md)| with protection settings.             | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 

### Return type

[**ProtectionDataResponse**](ProtectionDataResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBookmark**
> BookmarkResponse UpdateBookmark(ctx, name, bookmarkData, bookmarkName, optional)
Updates document bookmark.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **bookmarkData** | [**BookmarkData**](BookmarkData.md)| with new bookmark data.             | 
  **bookmarkName** | **string**| The bookmark name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **bookmarkData** | [**BookmarkData**](BookmarkData.md)| with new bookmark data.             | 
 **bookmarkName** | **string**| The bookmark name. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

[**BookmarkResponse**](BookmarkResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBorder**
> BorderResponse UpdateBorder(ctx, name, borderProperties, nodePath, borderType, optional)
Updates border properties.             

'nodePath' should refer to paragraph, cell or row.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **borderProperties** | [**Border**](Border.md)| Border properties. | 
  **nodePath** | **string**| Path to the node with border(node should be paragraph, cell or row). | 
  **borderType** | **string**| Border type. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **borderProperties** | [**Border**](Border.md)| Border properties. | 
 **nodePath** | **string**| Path to the node with border(node should be paragraph, cell or row). | 
 **borderType** | **string**| Border type. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

[**BorderResponse**](BorderResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateComment**
> CommentResponse UpdateComment(ctx, name, commentIndex, comment, optional)
Updates the comment, returns updated comment data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **commentIndex** | **int32**| The comment index. | 
  **comment** | [**Comment**](Comment.md)| The comment data. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **commentIndex** | **int32**| The comment index. | 
 **comment** | [**Comment**](Comment.md)| The comment data. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

[**CommentResponse**](CommentResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDrawingObject**
> DrawingObjectResponse UpdateDrawingObject(ctx, name, drawingObject, imageFile, nodePath, index, optional)
Updates drawing object, returns updated  drawing object's data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **drawingObject** | **string**| Drawing object parameters | 
  **imageFile** | ***os.File**| File with image | 
  **nodePath** | **string**| Path to the node, which contains collection of drawing objects. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **drawingObject** | **string**| Drawing object parameters | 
 **imageFile** | ***os.File**| File with image | 
 **nodePath** | **string**| Path to the node, which contains collection of drawing objects. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

[**DrawingObjectResponse**](DrawingObjectResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDrawingObjectWithoutNodePath**
> DrawingObjectResponse UpdateDrawingObjectWithoutNodePath(ctx, name, drawingObject, imageFile, index, optional)
Updates drawing object, returns updated  drawing object's data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **drawingObject** | **string**| Drawing object parameters | 
  **imageFile** | ***os.File**| File with image | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **drawingObject** | **string**| Drawing object parameters | 
 **imageFile** | ***os.File**| File with image | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

[**DrawingObjectResponse**](DrawingObjectResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateField**
> FieldResponse UpdateField(ctx, name, field, nodePath, index, optional)
Updates field's properties, returns updated field's data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **field** | [**Field**](Field.md)| Field data. | 
  **nodePath** | **string**| Path to the node, which contains collection of fields. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **field** | [**Field**](Field.md)| Field data. | 
 **nodePath** | **string**| Path to the node, which contains collection of fields. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

[**FieldResponse**](FieldResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFields**
> DocumentResponse UpdateFields(ctx, name, optional)
Updates (reevaluate) fields in document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 

### Return type

[**DocumentResponse**](DocumentResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFootnote**
> FootnoteResponse UpdateFootnote(ctx, name, footnoteDto, nodePath, index, optional)
Updates footnote's properties, returns updated run's data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **footnoteDto** | [**Footnote**](Footnote.md)| Footnote data. | 
  **nodePath** | **string**| Path to the node, which contains collection of footnotes. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **footnoteDto** | [**Footnote**](Footnote.md)| Footnote data. | 
 **nodePath** | **string**| Path to the node, which contains collection of footnotes. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

[**FootnoteResponse**](FootnoteResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFootnoteWithoutNodePath**
> FootnoteResponse UpdateFootnoteWithoutNodePath(ctx, name, footnoteDto, index, optional)
Updates footnote's properties, returns updated run's data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **footnoteDto** | [**Footnote**](Footnote.md)| Footnote data. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **footnoteDto** | [**Footnote**](Footnote.md)| Footnote data. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

[**FootnoteResponse**](FootnoteResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFormField**
> FormFieldResponse UpdateFormField(ctx, name, formField, nodePath, index, optional)
Updates properties of form field, returns updated form field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **formField** | [**FormField**](FormField.md)| From field data. | 
  **nodePath** | **string**| Path to the node that contains collection of formfields. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **formField** | [**FormField**](FormField.md)| From field data. | 
 **nodePath** | **string**| Path to the node that contains collection of formfields. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

[**FormFieldResponse**](FormFieldResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFormFieldWithoutNodePath**
> FormFieldResponse UpdateFormFieldWithoutNodePath(ctx, name, formField, index, optional)
Updates properties of form field, returns updated form field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **formField** | [**FormField**](FormField.md)| From field data. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **formField** | [**FormField**](FormField.md)| From field data. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

[**FormFieldResponse**](FormFieldResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateParagraphFormat**
> ParagraphFormatResponse UpdateParagraphFormat(ctx, name, dto, nodePath, index, optional)
Updates paragraph format properties, returns updated format properties.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **dto** | [**ParagraphFormat**](ParagraphFormat.md)| Paragraph format object. | 
  **nodePath** | **string**| Path to the node which contains paragraphs. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **dto** | [**ParagraphFormat**](ParagraphFormat.md)| Paragraph format object. | 
 **nodePath** | **string**| Path to the node which contains paragraphs. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

[**ParagraphFormatResponse**](ParagraphFormatResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRun**
> RunResponse UpdateRun(ctx, name, run, paragraphPath, index, optional)
Updates run's properties, returns updated run's data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **run** | [**Run**](Run.md)| Run data. | 
  **paragraphPath** | **string**| Path to parent paragraph. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **run** | [**Run**](Run.md)| Run data. | 
 **paragraphPath** | **string**| Path to parent paragraph. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

[**RunResponse**](RunResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRunFont**
> FontResponse UpdateRunFont(ctx, name, fontDto, paragraphPath, index, optional)
Updates font properties, returns updated font data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **fontDto** | [**Font**](Font.md)| Font dto object. | 
  **paragraphPath** | **string**| Path to parent paragraph. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **fontDto** | [**Font**](Font.md)| Font dto object. | 
 **paragraphPath** | **string**| Path to parent paragraph. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

[**FontResponse**](FontResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSectionPageSetup**
> SectionPageSetupResponse UpdateSectionPageSetup(ctx, name, sectionIndex, pageSetup, optional)
Updates page setup of section.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **sectionIndex** | **int32**| Section index. | 
  **pageSetup** | [**PageSetup**](PageSetup.md)| Page setup properties dto. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **sectionIndex** | **int32**| Section index. | 
 **pageSetup** | [**PageSetup**](PageSetup.md)| Page setup properties dto. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 

### Return type

[**SectionPageSetupResponse**](SectionPageSetupResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTableCellFormat**
> TableCellFormatResponse UpdateTableCellFormat(ctx, name, tableRowPath, index, optional)
Updates a table cell format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **tableRowPath** | **string**| Path to table row. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **tableRowPath** | **string**| Path to table row. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 
 **format** | [**TableCellFormat**](TableCellFormat.md)| The properties. | 

### Return type

[**TableCellFormatResponse**](TableCellFormatResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTableProperties**
> TablePropertiesResponse UpdateTableProperties(ctx, name, nodePath, index, optional)
Updates a table properties.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **nodePath** | **string**| Path to the node, which contains tables. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **nodePath** | **string**| Path to the node, which contains tables. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 
 **properties** | [**TableProperties**](TableProperties.md)| The properties. | 

### Return type

[**TablePropertiesResponse**](TablePropertiesResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTablePropertiesWithoutNodePath**
> TablePropertiesResponse UpdateTablePropertiesWithoutNodePath(ctx, name, index, optional)
Updates a table properties.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 
 **properties** | [**TableProperties**](TableProperties.md)| The properties. | 

### Return type

[**TablePropertiesResponse**](TablePropertiesResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTableRowFormat**
> TableRowFormatResponse UpdateTableRowFormat(ctx, name, tablePath, index, optional)
Updates a table row format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **tablePath** | **string**| Path to table. | 
  **index** | **int32**| Object index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **tablePath** | **string**| Path to table. | 
 **index** | **int32**| Object index. | 
 **folder** | **string**| Original document folder. | 
 **storage** | **string**| Original document storage. | 
 **loadEncoding** | **string**| Encoding that will be used to load an HTML (or TXT) document if the encoding is not specified in HTML. | 
 **password** | **string**| Password for opening an encrypted document. | 
 **destFileName** | **string**| Result path of the document after the operation. If this parameter is omitted then result of the operation will be saved as the source document. | 
 **revisionAuthor** | **string**| Initials of the author to use for revisions.If you set this parameter and then make some changes to the document programmatically, save the document and later open the document in MS Word you will see these changes as revisions. | 
 **revisionDateTime** | **string**| The date and time to use for revisions. | 
 **format** | [**TableRowFormat**](TableRowFormat.md)| Table row format. | 

### Return type

[**TableRowFormatResponse**](TableRowFormatResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/xml, application/json
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadFile**
> FilesUploadResult UploadFile(ctx, fileContent, path, optional)
Upload file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fileContent** | ***os.File**| File to upload | 
  **path** | **string**| Path where to upload including filename and extension e.g. /file.ext or /Folder 1/file.ext              If the content is multipart and path does not contains the file name it tries to get them from filename parameter              from Content-Disposition header. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileContent** | ***os.File**| File to upload | 
 **path** | **string**| Path where to upload including filename and extension e.g. /file.ext or /Folder 1/file.ext              If the content is multipart and path does not contains the file name it tries to get them from filename parameter              from Content-Disposition header. | 
 **storageName** | **string**| Storage name | 

### Return type

[**FilesUploadResult**](FilesUploadResult.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

