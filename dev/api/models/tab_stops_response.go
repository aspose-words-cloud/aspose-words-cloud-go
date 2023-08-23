/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="tab_stops_response.go">
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

// The REST response with an array of tab stops.

type ITabStopsResponse interface {
    IsTabStopsResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetRequestId() *string
    SetRequestId(value *string)
    GetTabStops() []ITabStop
    SetTabStops(value []ITabStop)
}

type TabStopsResponse struct {
    // The REST response with an array of tab stops.
    RequestId *string

    // The REST response with an array of tab stops.
    TabStops []ITabStop
}

func (TabStopsResponse) IsTabStopsResponse() bool {
    return true
}

func (TabStopsResponse) IsWordsResponse() bool {
    return true
}

func (obj *TabStopsResponse) Initialize() {
    if (obj.TabStops != nil) {
        for _, objElementTabStops := range obj.TabStops {
            objElementTabStops.Initialize()
        }
    }

}

func (obj *TabStopsResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["TabStops"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance ITabStop = new(TabStop)
                    modelElementInstance.Deserialize(elementValue)
                    obj.TabStops = append(obj.TabStops, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["tabStops"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance ITabStop = new(TabStop)
                    modelElementInstance.Deserialize(elementValue)
                    obj.TabStops = append(obj.TabStops, modelElementInstance)
                }

            }
        }

    }
}

func (obj *TabStopsResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *TabStopsResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *TabStopsResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *TabStopsResponse) GetTabStops() []ITabStop {
    return obj.TabStops
}

func (obj *TabStopsResponse) SetTabStops(value []ITabStop) {
    obj.TabStops = value
}

