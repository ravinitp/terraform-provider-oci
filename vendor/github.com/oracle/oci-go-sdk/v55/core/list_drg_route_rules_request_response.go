// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/v55/common"
	"net/http"
)

// ListDrgRouteRulesRequest wrapper for the ListDrgRouteRules operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/ListDrgRouteRules.go.html to see an example of how to use ListDrgRouteRulesRequest.
type ListDrgRouteRulesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG route table.
	DrgRouteTableId *string `mandatory:"true" contributesTo:"path" name:"drgRouteTableId"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Static routes are specified through the DRG route table API.
	// Dynamic routes are learned by the DRG from the DRG attachments through various routing protocols.
	RouteType ListDrgRouteRulesRouteTypeEnum `mandatory:"false" contributesTo:"query" name:"routeType" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDrgRouteRulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDrgRouteRulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDrgRouteRulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDrgRouteRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListDrgRouteRulesResponse wrapper for the ListDrgRouteRules operation
type ListDrgRouteRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []DrgRouteRule instances
	Items []DrgRouteRule `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListDrgRouteRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDrgRouteRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDrgRouteRulesRouteTypeEnum Enum with underlying type: string
type ListDrgRouteRulesRouteTypeEnum string

// Set of constants representing the allowable values for ListDrgRouteRulesRouteTypeEnum
const (
	ListDrgRouteRulesRouteTypeStatic  ListDrgRouteRulesRouteTypeEnum = "STATIC"
	ListDrgRouteRulesRouteTypeDynamic ListDrgRouteRulesRouteTypeEnum = "DYNAMIC"
)

var mappingListDrgRouteRulesRouteType = map[string]ListDrgRouteRulesRouteTypeEnum{
	"STATIC":  ListDrgRouteRulesRouteTypeStatic,
	"DYNAMIC": ListDrgRouteRulesRouteTypeDynamic,
}

// GetListDrgRouteRulesRouteTypeEnumValues Enumerates the set of values for ListDrgRouteRulesRouteTypeEnum
func GetListDrgRouteRulesRouteTypeEnumValues() []ListDrgRouteRulesRouteTypeEnum {
	values := make([]ListDrgRouteRulesRouteTypeEnum, 0)
	for _, v := range mappingListDrgRouteRulesRouteType {
		values = append(values, v)
	}
	return values
}
