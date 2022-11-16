/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="field_options.go">
 *   Copyright (c) 2022 Aspose.Words for Cloud
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

package models

// DTO for field options.
type FieldOptionsResult struct {
    // DTO for field options.
    BuiltInTemplatesPaths []string `json:"BuiltInTemplatesPaths,omitempty"`

    // DTO for field options.
    CurrentUser UserInformationResult `json:"CurrentUser,omitempty"`

    // DTO for field options.
    CustomTocStyleSeparator string `json:"CustomTocStyleSeparator,omitempty"`

    // DTO for field options.
    DefaultDocumentAuthor string `json:"DefaultDocumentAuthor,omitempty"`

    // DTO for field options.
    FieldIndexFormat string `json:"FieldIndexFormat,omitempty"`

    // DTO for field options.
    FieldUpdateCultureName string `json:"FieldUpdateCultureName,omitempty"`

    // DTO for field options.
    FieldUpdateCultureSource string `json:"FieldUpdateCultureSource,omitempty"`

    // DTO for field options.
    FileName string `json:"FileName,omitempty"`

    // DTO for field options.
    IsBidiTextSupportedOnUpdate bool `json:"IsBidiTextSupportedOnUpdate,omitempty"`

    // DTO for field options.
    LegacyNumberFormat bool `json:"LegacyNumberFormat,omitempty"`

    // DTO for field options.
    PreProcessCultureName string `json:"PreProcessCultureName,omitempty"`

    // DTO for field options.
    TemplateName string `json:"TemplateName,omitempty"`

    // DTO for field options.
    UseInvariantCultureNumberFormat bool `json:"UseInvariantCultureNumberFormat,omitempty"`
}

type FieldOptions struct {
    // DTO for field options.
    BuiltInTemplatesPaths []string `json:"BuiltInTemplatesPaths,omitempty"`

    // DTO for field options.
    CurrentUser IUserInformation `json:"CurrentUser,omitempty"`

    // DTO for field options.
    CustomTocStyleSeparator *string `json:"CustomTocStyleSeparator,omitempty"`

    // DTO for field options.
    DefaultDocumentAuthor *string `json:"DefaultDocumentAuthor,omitempty"`

    // DTO for field options.
    FieldIndexFormat *string `json:"FieldIndexFormat,omitempty"`

    // DTO for field options.
    FieldUpdateCultureName *string `json:"FieldUpdateCultureName,omitempty"`

    // DTO for field options.
    FieldUpdateCultureSource *string `json:"FieldUpdateCultureSource,omitempty"`

    // DTO for field options.
    FileName *string `json:"FileName,omitempty"`

    // DTO for field options.
    IsBidiTextSupportedOnUpdate *bool `json:"IsBidiTextSupportedOnUpdate,omitempty"`

    // DTO for field options.
    LegacyNumberFormat *bool `json:"LegacyNumberFormat,omitempty"`

    // DTO for field options.
    PreProcessCultureName *string `json:"PreProcessCultureName,omitempty"`

    // DTO for field options.
    TemplateName *string `json:"TemplateName,omitempty"`

    // DTO for field options.
    UseInvariantCultureNumberFormat *bool `json:"UseInvariantCultureNumberFormat,omitempty"`
}

type IFieldOptions interface {
    IsFieldOptions() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
}

func (FieldOptions) IsFieldOptions() bool {
    return true
}


func (obj *FieldOptions) Initialize() {
    if (obj.CurrentUser != nil) {
        obj.CurrentUser.Initialize()
    }


}

func (obj *FieldOptions) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}


