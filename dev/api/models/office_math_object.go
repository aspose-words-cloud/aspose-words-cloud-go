/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="office_math_object.go">
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

// DTO container with an OfficeMath object.

type IOfficeMathObject interface {
    IsOfficeMathObject() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetNodeId() *string
    SetNodeId(value *string)
    GetContent() IStoryChildNodes
    SetContent(value IStoryChildNodes)
    GetDisplayType() *string
    SetDisplayType(value *string)
    GetJustification() *string
    SetJustification(value *string)
    GetMathObjectType() *string
    SetMathObjectType(value *string)
}

type OfficeMathObject struct {
    // DTO container with an OfficeMath object.
    Link IWordsApiLink

    // DTO container with an OfficeMath object.
    NodeId *string

    // DTO container with an OfficeMath object.
    Content IStoryChildNodes

    // DTO container with an OfficeMath object.
    DisplayType *string

    // DTO container with an OfficeMath object.
    Justification *string

    // DTO container with an OfficeMath object.
    MathObjectType *string
}

func (OfficeMathObject) IsOfficeMathObject() bool {
    return true
}

func (OfficeMathObject) IsOfficeMathLink() bool {
    return true
}

func (OfficeMathObject) IsNodeLink() bool {
    return true
}

func (OfficeMathObject) IsLinkElement() bool {
    return true
}

func (obj *OfficeMathObject) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.Content != nil) {
        obj.Content.Initialize()
    }


}

func (obj *OfficeMathObject) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["NodeId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NodeId = &parsedValue
        }

    } else if jsonValue, exists := json["nodeId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NodeId = &parsedValue
        }

    }

    if jsonValue, exists := json["Content"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IStoryChildNodes = new(StoryChildNodes)
            modelInstance.Deserialize(parsedValue)
            obj.Content = modelInstance
        }

    } else if jsonValue, exists := json["content"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IStoryChildNodes = new(StoryChildNodes)
            modelInstance.Deserialize(parsedValue)
            obj.Content = modelInstance
        }

    }

    if jsonValue, exists := json["DisplayType"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DisplayType = &parsedValue
        }

    } else if jsonValue, exists := json["displayType"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DisplayType = &parsedValue
        }

    }

    if jsonValue, exists := json["Justification"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Justification = &parsedValue
        }

    } else if jsonValue, exists := json["justification"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Justification = &parsedValue
        }

    }

    if jsonValue, exists := json["MathObjectType"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.MathObjectType = &parsedValue
        }

    } else if jsonValue, exists := json["mathObjectType"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.MathObjectType = &parsedValue
        }

    }
}

func (obj *OfficeMathObject) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *OfficeMathObject) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *OfficeMathObject) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *OfficeMathObject) GetNodeId() *string {
    return obj.NodeId
}

func (obj *OfficeMathObject) SetNodeId(value *string) {
    obj.NodeId = value
}

func (obj *OfficeMathObject) GetContent() IStoryChildNodes {
    return obj.Content
}

func (obj *OfficeMathObject) SetContent(value IStoryChildNodes) {
    obj.Content = value
}

func (obj *OfficeMathObject) GetDisplayType() *string {
    return obj.DisplayType
}

func (obj *OfficeMathObject) SetDisplayType(value *string) {
    obj.DisplayType = value
}

func (obj *OfficeMathObject) GetJustification() *string {
    return obj.Justification
}

func (obj *OfficeMathObject) SetJustification(value *string) {
    obj.Justification = value
}

func (obj *OfficeMathObject) GetMathObjectType() *string {
    return obj.MathObjectType
}

func (obj *OfficeMathObject) SetMathObjectType(value *string) {
    obj.MathObjectType = value
}

