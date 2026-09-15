/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="job_info.go">
 *   Copyright (c) 2026 Aspose.Words for Cloud
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

// The REST response with a job result.

type IJobInfo interface {
    IsJobInfo() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetJobId() *string
    SetJobId(value *string)
    GetMessage() *string
    SetMessage(value *string)
    GetStatus() *string
    SetStatus(value *string)
}

type JobInfo struct {
    // The REST response with a job result.
    JobId *string `json:"JobId,omitempty"`

    // The REST response with a job result.
    Message *string `json:"Message,omitempty"`

    // The REST response with a job result.
    Status *string `json:"Status,omitempty"`
}

func (JobInfo) IsJobInfo() bool {
    return true
}


func (obj *JobInfo) Initialize() {
}

func (obj *JobInfo) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["JobId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.JobId = &parsedValue
        }

    } else if jsonValue, exists := json["jobId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.JobId = &parsedValue
        }

    }

    if jsonValue, exists := json["Message"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Message = &parsedValue
        }

    } else if jsonValue, exists := json["message"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Message = &parsedValue
        }

    }

    if jsonValue, exists := json["Status"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Status = &parsedValue
        }

    } else if jsonValue, exists := json["status"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Status = &parsedValue
        }

    }
}

func (obj *JobInfo) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *JobInfo) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Status == nil {
        return errors.New("Property Status in JobInfo is required.")
    }
    return nil;
}

func (obj *JobInfo) GetJobId() *string {
    return obj.JobId
}

func (obj *JobInfo) SetJobId(value *string) {
    obj.JobId = value
}

func (obj *JobInfo) GetMessage() *string {
    return obj.Message
}

func (obj *JobInfo) SetMessage(value *string) {
    obj.Message = value
}

func (obj *JobInfo) GetStatus() *string {
    return obj.Status
}

func (obj *JobInfo) SetStatus(value *string) {
    obj.Status = value
}

