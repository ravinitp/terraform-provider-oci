// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v55/databasemanagement"
)

func init() {
	RegisterDatasource("oci_database_management_managed_database_user_data_access_container", DatabaseManagementManagedDatabaseUserDataAccessContainerDataSource())
}

func DatabaseManagementManagedDatabaseUserDataAccessContainerDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseManagementManagedDatabaseUserDataAccessContainer,
		Schema: map[string]*schema.Schema{
			"managed_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularDatabaseManagementManagedDatabaseUserDataAccessContainer(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseUserDataAccessContainerDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dbManagementClient()

	return ReadResource(sync)
}

type DatabaseManagementManagedDatabaseUserDataAccessContainerDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListDataAccessContainersResponse
}

func (s *DatabaseManagementManagedDatabaseUserDataAccessContainerDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseUserDataAccessContainerDataSourceCrud) Get() error {
	request := oci_database_management.ListDataAccessContainersRequest{}

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if userName, ok := s.D.GetOkExists("user_name"); ok {
		tmp := userName.(string)
		request.UserName = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListDataAccessContainers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementManagedDatabaseUserDataAccessContainerDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("DatabaseManagementManagedDatabaseUserDataAccessContainerDataSource-", DatabaseManagementManagedDatabaseUserDataAccessContainerDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DataAccessContainerSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func DataAccessContainerSummaryToMap(obj oci_database_management.DataAccessContainerSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}
