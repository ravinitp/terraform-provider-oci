// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"io/ioutil"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_opsi "github.com/oracle/oci-go-sdk/v55/opsi"
)

func init() {
	RegisterResource("oci_opsi_operations_insights_warehouse_download_warehouse_wallet", OpsiOperationsInsightsWarehouseDownloadWarehouseWalletResource())
}

func OpsiOperationsInsightsWarehouseDownloadWarehouseWalletResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createOpsiOperationsInsightsWarehouseDownloadWarehouseWallet,
		Read:     readOpsiOperationsInsightsWarehouseDownloadWarehouseWallet,
		Delete:   deleteOpsiOperationsInsightsWarehouseDownloadWarehouseWallet,
		Schema: map[string]*schema.Schema{
			// Required
			"operations_insights_warehouse_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"operations_insights_warehouse_wallet_password": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},

			// Optional

			// Computed
		},
	}
}

func createOpsiOperationsInsightsWarehouseDownloadWarehouseWallet(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiOperationsInsightsWarehouseDownloadWarehouseWalletResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).operationsInsightsClient()

	return CreateResource(d, sync)
}

func readOpsiOperationsInsightsWarehouseDownloadWarehouseWallet(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteOpsiOperationsInsightsWarehouseDownloadWarehouseWallet(d *schema.ResourceData, m interface{}) error {
	return nil
}

type OpsiOperationsInsightsWarehouseDownloadWarehouseWalletResourceCrud struct {
	BaseCrud
	Client                 *oci_opsi.OperationsInsightsClient
	Res                    *string
	DisableNotFoundRetries bool
}

func (s *OpsiOperationsInsightsWarehouseDownloadWarehouseWalletResourceCrud) ID() string {
	return GenerateDataSourceHashID("OpsiOperationsInsightsWarehouseDownloadWarehouseWalletResource-", OpsiOperationsInsightsWarehouseDownloadWarehouseWalletResource(), s.D)
}

func (s *OpsiOperationsInsightsWarehouseDownloadWarehouseWalletResourceCrud) Create() error {
	request := oci_opsi.DownloadOperationsInsightsWarehouseWalletRequest{}

	if operationsInsightsWarehouseId, ok := s.D.GetOkExists("operations_insights_warehouse_id"); ok {
		tmp := operationsInsightsWarehouseId.(string)
		request.OperationsInsightsWarehouseId = &tmp
	}

	if operationsInsightsWarehouseWalletPassword, ok := s.D.GetOkExists("operations_insights_warehouse_wallet_password"); ok {
		tmp := operationsInsightsWarehouseWalletPassword.(string)
		request.OperationsInsightsWarehouseWalletPassword = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.DownloadOperationsInsightsWarehouseWallet(context.Background(), request)
	if err != nil {
		return err
	}

	if response.Content != nil {
		defer response.Content.Close()
		if contentBytes, err := ioutil.ReadAll(response.Content); err == nil {
			contentInStr := string(contentBytes)
			//                        s.Res = &([]byte(contentBytes)).String()
			s.Res = &contentInStr
		} else {
			return err
		}
	}
	return nil
}

func (s *OpsiOperationsInsightsWarehouseDownloadWarehouseWalletResourceCrud) SetData() error {
	return nil
}
