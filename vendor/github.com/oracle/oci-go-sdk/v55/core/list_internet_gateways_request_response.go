// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/v55/common"
	"net/http"
)

// ListInternetGatewaysRequest wrapper for the ListInternetGateways operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/ListInternetGateways.go.html to see an example of how to use ListInternetGatewaysRequest.
type ListInternetGatewaysRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN.
	VcnId *string `mandatory:"false" contributesTo:"query" name:"vcnId"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListInternetGatewaysSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListInternetGatewaysSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to only return resources that match the given lifecycle
	// state. The state value is case-insensitive.
	LifecycleState InternetGatewayLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListInternetGatewaysRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListInternetGatewaysRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListInternetGatewaysRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListInternetGatewaysRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListInternetGatewaysResponse wrapper for the ListInternetGateways operation
type ListInternetGatewaysResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []InternetGateway instances
	Items []InternetGateway `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListInternetGatewaysResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListInternetGatewaysResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListInternetGatewaysSortByEnum Enum with underlying type: string
type ListInternetGatewaysSortByEnum string

// Set of constants representing the allowable values for ListInternetGatewaysSortByEnum
const (
	ListInternetGatewaysSortByTimecreated ListInternetGatewaysSortByEnum = "TIMECREATED"
	ListInternetGatewaysSortByDisplayname ListInternetGatewaysSortByEnum = "DISPLAYNAME"
)

var mappingListInternetGatewaysSortBy = map[string]ListInternetGatewaysSortByEnum{
	"TIMECREATED": ListInternetGatewaysSortByTimecreated,
	"DISPLAYNAME": ListInternetGatewaysSortByDisplayname,
}

// GetListInternetGatewaysSortByEnumValues Enumerates the set of values for ListInternetGatewaysSortByEnum
func GetListInternetGatewaysSortByEnumValues() []ListInternetGatewaysSortByEnum {
	values := make([]ListInternetGatewaysSortByEnum, 0)
	for _, v := range mappingListInternetGatewaysSortBy {
		values = append(values, v)
	}
	return values
}

// ListInternetGatewaysSortOrderEnum Enum with underlying type: string
type ListInternetGatewaysSortOrderEnum string

// Set of constants representing the allowable values for ListInternetGatewaysSortOrderEnum
const (
	ListInternetGatewaysSortOrderAsc  ListInternetGatewaysSortOrderEnum = "ASC"
	ListInternetGatewaysSortOrderDesc ListInternetGatewaysSortOrderEnum = "DESC"
)

var mappingListInternetGatewaysSortOrder = map[string]ListInternetGatewaysSortOrderEnum{
	"ASC":  ListInternetGatewaysSortOrderAsc,
	"DESC": ListInternetGatewaysSortOrderDesc,
}

// GetListInternetGatewaysSortOrderEnumValues Enumerates the set of values for ListInternetGatewaysSortOrderEnum
func GetListInternetGatewaysSortOrderEnumValues() []ListInternetGatewaysSortOrderEnum {
	values := make([]ListInternetGatewaysSortOrderEnum, 0)
	for _, v := range mappingListInternetGatewaysSortOrder {
		values = append(values, v)
	}
	return values
}
