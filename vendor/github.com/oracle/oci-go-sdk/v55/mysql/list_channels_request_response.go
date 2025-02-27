// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package mysql

import (
	"github.com/oracle/oci-go-sdk/v55/common"
	"net/http"
)

// ListChannelsRequest wrapper for the ListChannels operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/ListChannels.go.html to see an example of how to use ListChannelsRequest.
type ListChannelsRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Customer-defined unique identifier for the request. If you need to
	// contact Oracle about a specific request, please provide the request
	// ID that you supplied in this header with the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The DB System OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	DbSystemId *string `mandatory:"false" contributesTo:"query" name:"dbSystemId"`

	// The OCID of the Channel.
	ChannelId *string `mandatory:"false" contributesTo:"query" name:"channelId"`

	// A filter to return only the resource matching the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The LifecycleState of the Channel.
	LifecycleState ChannelLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// If true, returns only Channels that are enabled. If false, returns only
	// Channels that are disabled.
	IsEnabled *bool `mandatory:"false" contributesTo:"query" name:"isEnabled"`

	// The field to sort by. Only one sort order may be provided. Time fields are default ordered as descending. Display name is default ordered as ascending.
	SortBy ListChannelsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use (ASC or DESC).
	SortOrder ListChannelsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return in a paginated list call. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.comAPI/Concepts/usingapi.htm#List_Pagination).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` or `opc-prev-page` response header from
	// the previous list call. For information about pagination, see List
	// Pagination (https://docs.cloud.oracle.comAPI/Concepts/usingapi.htm#List_Pagination).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListChannelsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListChannelsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListChannelsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListChannelsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListChannelsResponse wrapper for the ListChannels operation
type ListChannelsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ChannelSummary instances
	Items []ChannelSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListChannelsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListChannelsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListChannelsSortByEnum Enum with underlying type: string
type ListChannelsSortByEnum string

// Set of constants representing the allowable values for ListChannelsSortByEnum
const (
	ListChannelsSortByDisplayname ListChannelsSortByEnum = "displayName"
	ListChannelsSortByTimecreated ListChannelsSortByEnum = "timeCreated"
)

var mappingListChannelsSortBy = map[string]ListChannelsSortByEnum{
	"displayName": ListChannelsSortByDisplayname,
	"timeCreated": ListChannelsSortByTimecreated,
}

// GetListChannelsSortByEnumValues Enumerates the set of values for ListChannelsSortByEnum
func GetListChannelsSortByEnumValues() []ListChannelsSortByEnum {
	values := make([]ListChannelsSortByEnum, 0)
	for _, v := range mappingListChannelsSortBy {
		values = append(values, v)
	}
	return values
}

// ListChannelsSortOrderEnum Enum with underlying type: string
type ListChannelsSortOrderEnum string

// Set of constants representing the allowable values for ListChannelsSortOrderEnum
const (
	ListChannelsSortOrderAsc  ListChannelsSortOrderEnum = "ASC"
	ListChannelsSortOrderDesc ListChannelsSortOrderEnum = "DESC"
)

var mappingListChannelsSortOrder = map[string]ListChannelsSortOrderEnum{
	"ASC":  ListChannelsSortOrderAsc,
	"DESC": ListChannelsSortOrderDesc,
}

// GetListChannelsSortOrderEnumValues Enumerates the set of values for ListChannelsSortOrderEnum
func GetListChannelsSortOrderEnumValues() []ListChannelsSortOrderEnum {
	values := make([]ListChannelsSortOrderEnum, 0)
	for _, v := range mappingListChannelsSortOrder {
		values = append(values, v)
	}
	return values
}
