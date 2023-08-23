/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="compare_options.go">
 *   Copyright (c) 2023 Aspose.Words for Cloud
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

// DTO container with compare documents options.

type ICompareOptions interface {
    IsCompareOptions() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetAcceptAllRevisionsBeforeComparison() *bool
    SetAcceptAllRevisionsBeforeComparison(value *bool)
    GetIgnoreCaseChanges() *bool
    SetIgnoreCaseChanges(value *bool)
    GetIgnoreComments() *bool
    SetIgnoreComments(value *bool)
    GetIgnoreFields() *bool
    SetIgnoreFields(value *bool)
    GetIgnoreFootnotes() *bool
    SetIgnoreFootnotes(value *bool)
    GetIgnoreFormatting() *bool
    SetIgnoreFormatting(value *bool)
    GetIgnoreHeadersAndFooters() *bool
    SetIgnoreHeadersAndFooters(value *bool)
    GetIgnoreTables() *bool
    SetIgnoreTables(value *bool)
    GetIgnoreTextboxes() *bool
    SetIgnoreTextboxes(value *bool)
    GetTarget() *string
    SetTarget(value *string)
}

type CompareOptions struct {
    // DTO container with compare documents options.
    AcceptAllRevisionsBeforeComparison *bool

    // DTO container with compare documents options.
    IgnoreCaseChanges *bool

    // DTO container with compare documents options.
    IgnoreComments *bool

    // DTO container with compare documents options.
    IgnoreFields *bool

    // DTO container with compare documents options.
    IgnoreFootnotes *bool

    // DTO container with compare documents options.
    IgnoreFormatting *bool

    // DTO container with compare documents options.
    IgnoreHeadersAndFooters *bool

    // DTO container with compare documents options.
    IgnoreTables *bool

    // DTO container with compare documents options.
    IgnoreTextboxes *bool

    // DTO container with compare documents options.
    Target *string
}

func (CompareOptions) IsCompareOptions() bool {
    return true
}


func (obj *CompareOptions) Initialize() {
}

func (obj *CompareOptions) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["AcceptAllRevisionsBeforeComparison"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.AcceptAllRevisionsBeforeComparison = &parsedValue
        }

    } else if jsonValue, exists := json["acceptAllRevisionsBeforeComparison"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.AcceptAllRevisionsBeforeComparison = &parsedValue
        }

    }

    if jsonValue, exists := json["IgnoreCaseChanges"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IgnoreCaseChanges = &parsedValue
        }

    } else if jsonValue, exists := json["ignoreCaseChanges"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IgnoreCaseChanges = &parsedValue
        }

    }

    if jsonValue, exists := json["IgnoreComments"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IgnoreComments = &parsedValue
        }

    } else if jsonValue, exists := json["ignoreComments"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IgnoreComments = &parsedValue
        }

    }

    if jsonValue, exists := json["IgnoreFields"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IgnoreFields = &parsedValue
        }

    } else if jsonValue, exists := json["ignoreFields"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IgnoreFields = &parsedValue
        }

    }

    if jsonValue, exists := json["IgnoreFootnotes"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IgnoreFootnotes = &parsedValue
        }

    } else if jsonValue, exists := json["ignoreFootnotes"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IgnoreFootnotes = &parsedValue
        }

    }

    if jsonValue, exists := json["IgnoreFormatting"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IgnoreFormatting = &parsedValue
        }

    } else if jsonValue, exists := json["ignoreFormatting"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IgnoreFormatting = &parsedValue
        }

    }

    if jsonValue, exists := json["IgnoreHeadersAndFooters"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IgnoreHeadersAndFooters = &parsedValue
        }

    } else if jsonValue, exists := json["ignoreHeadersAndFooters"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IgnoreHeadersAndFooters = &parsedValue
        }

    }

    if jsonValue, exists := json["IgnoreTables"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IgnoreTables = &parsedValue
        }

    } else if jsonValue, exists := json["ignoreTables"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IgnoreTables = &parsedValue
        }

    }

    if jsonValue, exists := json["IgnoreTextboxes"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IgnoreTextboxes = &parsedValue
        }

    } else if jsonValue, exists := json["ignoreTextboxes"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IgnoreTextboxes = &parsedValue
        }

    }

    if jsonValue, exists := json["Target"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Target = &parsedValue
        }

    } else if jsonValue, exists := json["target"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Target = &parsedValue
        }

    }
}

func (obj *CompareOptions) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *CompareOptions) GetAcceptAllRevisionsBeforeComparison() *bool {
    return obj.AcceptAllRevisionsBeforeComparison
}

func (obj *CompareOptions) SetAcceptAllRevisionsBeforeComparison(value *bool) {
    obj.AcceptAllRevisionsBeforeComparison = value
}

func (obj *CompareOptions) GetIgnoreCaseChanges() *bool {
    return obj.IgnoreCaseChanges
}

func (obj *CompareOptions) SetIgnoreCaseChanges(value *bool) {
    obj.IgnoreCaseChanges = value
}

func (obj *CompareOptions) GetIgnoreComments() *bool {
    return obj.IgnoreComments
}

func (obj *CompareOptions) SetIgnoreComments(value *bool) {
    obj.IgnoreComments = value
}

func (obj *CompareOptions) GetIgnoreFields() *bool {
    return obj.IgnoreFields
}

func (obj *CompareOptions) SetIgnoreFields(value *bool) {
    obj.IgnoreFields = value
}

func (obj *CompareOptions) GetIgnoreFootnotes() *bool {
    return obj.IgnoreFootnotes
}

func (obj *CompareOptions) SetIgnoreFootnotes(value *bool) {
    obj.IgnoreFootnotes = value
}

func (obj *CompareOptions) GetIgnoreFormatting() *bool {
    return obj.IgnoreFormatting
}

func (obj *CompareOptions) SetIgnoreFormatting(value *bool) {
    obj.IgnoreFormatting = value
}

func (obj *CompareOptions) GetIgnoreHeadersAndFooters() *bool {
    return obj.IgnoreHeadersAndFooters
}

func (obj *CompareOptions) SetIgnoreHeadersAndFooters(value *bool) {
    obj.IgnoreHeadersAndFooters = value
}

func (obj *CompareOptions) GetIgnoreTables() *bool {
    return obj.IgnoreTables
}

func (obj *CompareOptions) SetIgnoreTables(value *bool) {
    obj.IgnoreTables = value
}

func (obj *CompareOptions) GetIgnoreTextboxes() *bool {
    return obj.IgnoreTextboxes
}

func (obj *CompareOptions) SetIgnoreTextboxes(value *bool) {
    obj.IgnoreTextboxes = value
}

func (obj *CompareOptions) GetTarget() *string {
    return obj.Target
}

func (obj *CompareOptions) SetTarget(value *string) {
    obj.Target = value
}

