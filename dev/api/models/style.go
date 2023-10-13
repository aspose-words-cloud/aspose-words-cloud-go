/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="style.go">
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

import (
    "errors"
)

// DTO container with a single document style.

type IStyle interface {
    IsStyle() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetFont() IFont
    SetFont(value IFont)
    GetBuiltIn() *bool
    SetBuiltIn(value *bool)
    GetNextParagraphStyleName() *string
    SetNextParagraphStyleName(value *string)
    GetBaseStyleName() *string
    SetBaseStyleName(value *string)
    GetIsQuickStyle() *bool
    SetIsQuickStyle(value *bool)
    GetLinkedStyleName() *string
    SetLinkedStyleName(value *string)
    GetType() *string
    SetType(value *string)
    GetIsHeading() *bool
    SetIsHeading(value *bool)
    GetAliases() []string
    SetAliases(value []string)
    GetStyleIdentifier() *string
    SetStyleIdentifier(value *string)
    GetName() *string
    SetName(value *string)
}

type Style struct {
    // DTO container with a single document style.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with a single document style.
    Font IFont `json:"Font,omitempty"`

    // DTO container with a single document style.
    BuiltIn *bool `json:"BuiltIn,omitempty"`

    // DTO container with a single document style.
    NextParagraphStyleName *string `json:"NextParagraphStyleName,omitempty"`

    // DTO container with a single document style.
    BaseStyleName *string `json:"BaseStyleName,omitempty"`

    // DTO container with a single document style.
    IsQuickStyle *bool `json:"IsQuickStyle,omitempty"`

    // DTO container with a single document style.
    LinkedStyleName *string `json:"LinkedStyleName,omitempty"`

    // DTO container with a single document style.
    Type *string `json:"Type,omitempty"`

    // DTO container with a single document style.
    IsHeading *bool `json:"IsHeading,omitempty"`

    // DTO container with a single document style.
    Aliases []string `json:"Aliases,omitempty"`

    // DTO container with a single document style.
    StyleIdentifier *string `json:"StyleIdentifier,omitempty"`

    // DTO container with a single document style.
    Name *string `json:"Name,omitempty"`
}

func (Style) IsStyle() bool {
    return true
}

func (Style) IsLinkElement() bool {
    return true
}

func (obj *Style) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.Font != nil) {
        obj.Font.Initialize()
    }


}

func (obj *Style) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Link"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IWordsApiLink = new(WordsApiLink)
            modelInstance.Deserialize(parsedValue)
            obj.Link = modelInstance
        }

    } else if jsonValue, exists := json["link"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IWordsApiLink = new(WordsApiLink)
            modelInstance.Deserialize(parsedValue)
            obj.Link = modelInstance
        }

    }

    if jsonValue, exists := json["Font"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFont = new(Font)
            modelInstance.Deserialize(parsedValue)
            obj.Font = modelInstance
        }

    } else if jsonValue, exists := json["font"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFont = new(Font)
            modelInstance.Deserialize(parsedValue)
            obj.Font = modelInstance
        }

    }

    if jsonValue, exists := json["BuiltIn"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.BuiltIn = &parsedValue
        }

    } else if jsonValue, exists := json["builtIn"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.BuiltIn = &parsedValue
        }

    }

    if jsonValue, exists := json["NextParagraphStyleName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NextParagraphStyleName = &parsedValue
        }

    } else if jsonValue, exists := json["nextParagraphStyleName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NextParagraphStyleName = &parsedValue
        }

    }

    if jsonValue, exists := json["BaseStyleName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.BaseStyleName = &parsedValue
        }

    } else if jsonValue, exists := json["baseStyleName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.BaseStyleName = &parsedValue
        }

    }

    if jsonValue, exists := json["IsQuickStyle"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsQuickStyle = &parsedValue
        }

    } else if jsonValue, exists := json["isQuickStyle"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsQuickStyle = &parsedValue
        }

    }

    if jsonValue, exists := json["LinkedStyleName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.LinkedStyleName = &parsedValue
        }

    } else if jsonValue, exists := json["linkedStyleName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.LinkedStyleName = &parsedValue
        }

    }

    if jsonValue, exists := json["Type"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Type = &parsedValue
        }

    } else if jsonValue, exists := json["type"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Type = &parsedValue
        }

    }

    if jsonValue, exists := json["IsHeading"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsHeading = &parsedValue
        }

    } else if jsonValue, exists := json["isHeading"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsHeading = &parsedValue
        }

    }

    if jsonValue, exists := json["Aliases"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.Aliases = make([]string, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(string); valid {
                    obj.Aliases = append(obj.Aliases, elementValue)
                }

            }
        }

    } else if jsonValue, exists := json["aliases"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.Aliases = make([]string, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(string); valid {
                    obj.Aliases = append(obj.Aliases, elementValue)
                }

            }
        }

    }

    if jsonValue, exists := json["StyleIdentifier"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.StyleIdentifier = &parsedValue
        }

    } else if jsonValue, exists := json["styleIdentifier"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.StyleIdentifier = &parsedValue
        }

    }

    if jsonValue, exists := json["Name"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Name = &parsedValue
        }

    } else if jsonValue, exists := json["name"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Name = &parsedValue
        }

    }
}

func (obj *Style) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *Style) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.BuiltIn == nil {
        return errors.New("Property BuiltIn in Style is required.")
    }

    if obj.IsQuickStyle == nil {
        return errors.New("Property IsQuickStyle in Style is required.")
    }

    if obj.Type == nil {
        return errors.New("Property Type in Style is required.")
    }

    if obj.IsHeading == nil {
        return errors.New("Property IsHeading in Style is required.")
    }

    if obj.StyleIdentifier == nil {
        return errors.New("Property StyleIdentifier in Style is required.")
    }

    return nil;
}

func (obj *Style) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *Style) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *Style) GetFont() IFont {
    return obj.Font
}

func (obj *Style) SetFont(value IFont) {
    obj.Font = value
}

func (obj *Style) GetBuiltIn() *bool {
    return obj.BuiltIn
}

func (obj *Style) SetBuiltIn(value *bool) {
    obj.BuiltIn = value
}

func (obj *Style) GetNextParagraphStyleName() *string {
    return obj.NextParagraphStyleName
}

func (obj *Style) SetNextParagraphStyleName(value *string) {
    obj.NextParagraphStyleName = value
}

func (obj *Style) GetBaseStyleName() *string {
    return obj.BaseStyleName
}

func (obj *Style) SetBaseStyleName(value *string) {
    obj.BaseStyleName = value
}

func (obj *Style) GetIsQuickStyle() *bool {
    return obj.IsQuickStyle
}

func (obj *Style) SetIsQuickStyle(value *bool) {
    obj.IsQuickStyle = value
}

func (obj *Style) GetLinkedStyleName() *string {
    return obj.LinkedStyleName
}

func (obj *Style) SetLinkedStyleName(value *string) {
    obj.LinkedStyleName = value
}

func (obj *Style) GetType() *string {
    return obj.Type
}

func (obj *Style) SetType(value *string) {
    obj.Type = value
}

func (obj *Style) GetIsHeading() *bool {
    return obj.IsHeading
}

func (obj *Style) SetIsHeading(value *bool) {
    obj.IsHeading = value
}

func (obj *Style) GetAliases() []string {
    return obj.Aliases
}

func (obj *Style) SetAliases(value []string) {
    obj.Aliases = value
}

func (obj *Style) GetStyleIdentifier() *string {
    return obj.StyleIdentifier
}

func (obj *Style) SetStyleIdentifier(value *string) {
    obj.StyleIdentifier = value
}

func (obj *Style) GetName() *string {
    return obj.Name
}

func (obj *Style) SetName(value *string) {
    obj.Name = value
}

