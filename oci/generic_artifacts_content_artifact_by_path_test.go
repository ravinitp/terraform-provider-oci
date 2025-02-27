// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ArtifactByPathResourceConfig = ArtifactByPathResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_generic_artifacts_content_artifact_by_path", "test_artifact_by_path", Required, Create, artifactByPathRepresentation)

	artifactByPathSingularDataSourceRepresentation = map[string]interface{}{
		"artifact_path": Representation{RepType: Required, Create: `artifactPath`},
		"repository_id": Representation{RepType: Required, Create: `${oci_artifacts_repository.test_repository.id}`},
		"version":       Representation{RepType: Required, Create: `1.0`},
	}

	artifactByPathRepresentation = map[string]interface{}{
		"artifact_path": Representation{RepType: Required, Create: `artifactPath`},
		"repository_id": Representation{RepType: Required, Create: `${oci_artifacts_repository.test_repository.id}`},
		"version":       Representation{RepType: Required, Create: `1.0`},
		"content":       Representation{RepType: Required, Create: `<a1>content</a1>`},
	}

	ArtifactByPathResourceDependencies = GenerateResourceFromRepresentationMap("oci_artifacts_repository", "test_repository", Required, Create, repositoryRepresentation)
	// the deletion of oci_generic_artifacts_content_artifact_by_path is done by oci_artifacts_generic_artifact
	GenericArtifactManager = GenerateResourceFromRepresentationMap("oci_artifacts_generic_artifact", "test_generic_artifact", Required, Create, genericArtifactRepresentation)
)

// issue-routing-tag: generic_artifacts_content/default
func TestGenericArtifactsContentArtifactByPathResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGenericArtifactsContentArtifactByPathResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_generic_artifacts_content_artifact_by_path.test_artifact_by_path"

	singularDatasourceName := "data.oci_generic_artifacts_content_artifact_by_path.test_artifact_by_path"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+ArtifactByPathResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_generic_artifacts_content_artifact_by_path", "test_artifact_by_path", Required, Create, artifactByPathRepresentation), "genericartifactscontent", "artifactByPath", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ArtifactByPathResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_generic_artifacts_content_artifact_by_path", "test_artifact_by_path", Required, Create, artifactByPathRepresentation) + GenericArtifactManager,
			Check: ComposeAggregateTestCheckFuncWrapper(

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ArtifactByPathResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_generic_artifacts_content_artifact_by_path", "test_artifact_by_path", Optional, Update, artifactByPathRepresentation) + GenericArtifactManager,
			Check: ComposeAggregateTestCheckFuncWrapper(

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_generic_artifacts_content_artifact_by_path", "test_artifact_by_path", Required, Create, artifactByPathSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ArtifactByPathResourceConfig + GenericArtifactManager,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "artifact_path", "artifactPath"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "repository_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "version", "1.0"),
			),
		},
	})
}

const (
	tempFilePrefix = "small-"
	tempFileSize   = 2e5
	tempFileSha256 = "4cbbd9be0cba685835755f827758705db5a413c5494c34262cd25946a73e7582"
)

func createTmpFile() (string, error) {
	tempFile, err := ioutil.TempFile(os.TempDir(), tempFilePrefix)
	if err != nil {
		return "", err
	}
	if err := tempFile.Truncate(tempFileSize); err != nil {
		return "", err
	}
	return tempFile.Name(), nil
}

var (
	artifactByPathSourceRepresentation = map[string]interface{}{
		"artifact_path": Representation{RepType: Required, Create: `artifactPath`},
		"repository_id": Representation{RepType: Required, Create: `${oci_artifacts_repository.test_repository.id}`},
		"version":       Representation{RepType: Required, Create: `1.0`},
		"source":        Representation{RepType: Required, Create: ``},
	}
)

// issue-routing-tag: generic_artifacts_content/default
func TestGenericArtifactsContentArtifactByPathResource_uploadFile(t *testing.T) {
	httpreplay.SetScenario("TestGenericArtifactsContentArtifactByPathResource_uploadFile")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_generic_artifacts_content_artifact_by_path.test_artifact_by_path"

	tempFilePath, err := createTmpFile()
	if err != nil {
		t.Fatalf("Unable to Create file to upload. Error: %q", err)
	}

	var resId, _ string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+ArtifactByPathResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_generic_artifacts_content_artifact_by_path", "test_artifact_by_path", Required, Create, artifactByPathRepresentation), "genericartifactscontent", "artifactByPath", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: config + compartmentIdVariableStr + ArtifactByPathResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_generic_artifacts_content_artifact_by_path", "test_artifact_by_path", Required, Create,
						GetUpdatedRepresentationCopy("source", Representation{RepType: Required, Create: tempFilePath}, artifactByPathSourceRepresentation)) + GenericArtifactManager,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "sha256", tempFileSha256),
					resource.TestCheckResourceAttr(resourceName, "size_in_bytes", strconv.Itoa(tempFileSize)),
					func(s *terraform.State) (err error) {
						resId, err = FromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},
		},
	})
}
