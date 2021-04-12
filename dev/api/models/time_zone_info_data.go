/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="time_zone_info_data.go">
 *   Copyright (c) 2021 Aspose.Words for Cloud
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

// Class to specify TimeZoneInfo parameters.
type TimeZoneInfoDataResult struct {
    // Class to specify TimeZoneInfo parameters.
    BaseUtcOffset string `json:"BaseUtcOffset,omitempty"`

    // Class to specify TimeZoneInfo parameters.
    DisplayName string `json:"DisplayName,omitempty"`

    // Class to specify TimeZoneInfo parameters.
    Id string `json:"Id,omitempty"`

    // Class to specify TimeZoneInfo parameters.
    StandardDisplayName string `json:"StandardDisplayName,omitempty"`
}

type TimeZoneInfoData struct {
    // Class to specify TimeZoneInfo parameters.
    BaseUtcOffset *string `json:"BaseUtcOffset,omitempty"`

    // Class to specify TimeZoneInfo parameters.
    DisplayName *string `json:"DisplayName,omitempty"`

    // Class to specify TimeZoneInfo parameters.
    Id *string `json:"Id,omitempty"`

    // Class to specify TimeZoneInfo parameters.
    StandardDisplayName *string `json:"StandardDisplayName,omitempty"`
}

type ITimeZoneInfoData interface {
    IsTimeZoneInfoData() bool
}
func (TimeZoneInfoData) IsTimeZoneInfoData() bool {
    return true
}


