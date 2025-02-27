// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package mysql

import (
	"github.com/oracle/oci-go-sdk/v55/common"
	"net/http"
)

// GetAnalyticsClusterMemoryEstimateRequest wrapper for the GetAnalyticsClusterMemoryEstimate operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/GetAnalyticsClusterMemoryEstimate.go.html to see an example of how to use GetAnalyticsClusterMemoryEstimateRequest.
type GetAnalyticsClusterMemoryEstimateRequest struct {

	// The DB System OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	DbSystemId *string `mandatory:"true" contributesTo:"path" name:"dbSystemId"`

	// Customer-defined unique identifier for the request. If you need to
	// contact Oracle about a specific request, please provide the request
	// ID that you supplied in this header with the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetAnalyticsClusterMemoryEstimateRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetAnalyticsClusterMemoryEstimateRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetAnalyticsClusterMemoryEstimateRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetAnalyticsClusterMemoryEstimateRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetAnalyticsClusterMemoryEstimateResponse wrapper for the GetAnalyticsClusterMemoryEstimate operation
type GetAnalyticsClusterMemoryEstimateResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The AnalyticsClusterMemoryEstimate instance
	AnalyticsClusterMemoryEstimate `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetAnalyticsClusterMemoryEstimateResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetAnalyticsClusterMemoryEstimateResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
