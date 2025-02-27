// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v55/datasafe"
)

func init() {
	RegisterDatasource("oci_data_safe_user_assessment_comparison", DataSafeUserAssessmentComparisonDataSource())
}

func DataSafeUserAssessmentComparisonDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDataSafeUserAssessmentComparison,
		Schema: map[string]*schema.Schema{
			"comparison_user_assessment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_assessment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"summary": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"baseline": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"userAssessmentId": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"targetId": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"current": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"userAssessmentId": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"targetId": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDataSafeUserAssessmentComparison(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeUserAssessmentComparisonDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataSafeClient()

	return ReadResource(sync)
}

type DataSafeUserAssessmentComparisonDataSourceCrud struct {
	BaseCrud
	D                      *schema.ResourceData
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.GetUserAssessmentComparisonResponse
	DisableNotFoundRetries bool
}

func (s *DataSafeUserAssessmentComparisonDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeUserAssessmentComparisonDataSourceCrud) Get() error {
	request := oci_data_safe.GetUserAssessmentComparisonRequest{}

	if comparisonUserAssessmentId, ok := s.D.GetOkExists("comparison_user_assessment_id"); ok {
		tmp := comparisonUserAssessmentId.(string)
		request.ComparisonUserAssessmentId = &tmp
	}

	if userAssessmentId, ok := s.D.GetOkExists("user_assessment_id"); ok {
		tmp := userAssessmentId.(string)
		request.UserAssessmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetUserAssessmentComparison(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeUserAssessmentComparisonDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("DataSafeUserAssessmentComparisonDataSource-", DataSafeUserAssessmentComparisonDataSource(), s.D))

	s.D.Set("state", s.Res.LifecycleState)

	summary := []interface{}{}
	for _, item := range s.Res.Summary {
		summary = append(summary, item)
	}
	s.D.Set("summary", summary)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
