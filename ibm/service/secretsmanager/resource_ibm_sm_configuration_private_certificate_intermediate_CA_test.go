// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/secrets-manager-go-sdk/secretsmanagerv2"
)

func TestAccIbmSmConfigurationPrivateCertificateIntermediateCABasic(t *testing.T) {
	var conf secretsmanagerv2.PrivateCertificateConfigurationIntermediateCA

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmConfigurationPrivateCertificateIntermediateCADestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmConfigurationPrivateCertificateIntermediateCAConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmConfigurationPrivateCertificateIntermediateCAExists("ibm_sm_configuration_private_certificate_intermediate_ca.sm_configuration_private_certificate_intermediate_ca", conf),
				),
			},
			//resource.TestStep{
			//	ResourceName:      "ibm_sm_configuration_private_certificate_intermediate_ca.sm_configuration_private_certificate_intermediate_ca",
			//	ImportState:       true,
			//	ImportStateVerify: true,
			//},
		},
	})
}

func testAccCheckIbmSmConfigurationPrivateCertificateIntermediateCAConfigBasic() string {
	return fmt.Sprintf(`

		resource "ibm_sm_configuration_private_certificate_root_ca" "ibm_sm_configuration_private_certificate_root_ca_instance" {
			instance_id   = "%s"
			region        = "%s"
			max_ttl = "180000"
			common_name = "ibm.com"
			crl_expiry = "10000h"
			name = "root-ca-terraform-private-cert-datasource-test"
		}
		resource "ibm_sm_configuration_private_certificate_intermediate_ca" "sm_configuration_private_certificate_intermediate_ca" {
  			instance_id   = "%s"
			region        = "%s"
			max_ttl = "180000"
			common_name = "ibm.com"
			issuer = ibm_sm_configuration_private_certificate_root_ca.ibm_sm_configuration_private_certificate_root_ca_instance.name
			signing_method = "internal"
			name = "intermediate-ca-terraform-private-cert-datasource-test"
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}

func testAccCheckIbmSmConfigurationPrivateCertificateIntermediateCAExists(n string, obj secretsmanagerv2.PrivateCertificateConfigurationIntermediateCA) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV2()
		if err != nil {
			return err
		}

		secretsManagerClient = getClientWithInstanceEndpointTest(secretsManagerClient)

		getConfigurationOptions := &secretsmanagerv2.GetConfigurationOptions{}

		getConfigurationOptions.SetName(rs.Primary.ID)

		privateCertificateConfigurationIntermediateCAIntf, _, err := secretsManagerClient.GetConfiguration(getConfigurationOptions)
		if err != nil {
			return err
		}

		privateCertificateConfigurationIntermediateCA := privateCertificateConfigurationIntermediateCAIntf.(*secretsmanagerv2.PrivateCertificateConfigurationIntermediateCA)
		obj = *privateCertificateConfigurationIntermediateCA
		return nil
	}
}

func testAccCheckIbmSmConfigurationPrivateCertificateIntermediateCADestroy(s *terraform.State) error {
	secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return err
	}

	secretsManagerClient = getClientWithInstanceEndpointTest(secretsManagerClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_sm_configuration_private_certificate_intermediate_ca" {
			continue
		}

		getConfigurationOptions := &secretsmanagerv2.GetConfigurationOptions{}

		getConfigurationOptions.SetName(rs.Primary.ID)

		// Try to find the key
		_, response, err := secretsManagerClient.GetConfiguration(getConfigurationOptions)

		if err == nil {
			return fmt.Errorf("PrivateCertificateConfigurationIntermediateCA still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for PrivateCertificateConfigurationIntermediateCA (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
