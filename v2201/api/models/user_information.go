/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="user_information.go">
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

// DTO for user information.
type UserInformationResult struct {
    // DTO for user information.
    Address string `json:"Address,omitempty"`

    // DTO for user information.
    Initials string `json:"Initials,omitempty"`

    // DTO for user information.
    Name string `json:"Name,omitempty"`
}

type UserInformation struct {
    // DTO for user information.
    Address *string `json:"Address,omitempty"`

    // DTO for user information.
    Initials *string `json:"Initials,omitempty"`

    // DTO for user information.
    Name *string `json:"Name,omitempty"`
}

type IUserInformation interface {
    IsUserInformation() bool
}
func (UserInformation) IsUserInformation() bool {
    return true
}


