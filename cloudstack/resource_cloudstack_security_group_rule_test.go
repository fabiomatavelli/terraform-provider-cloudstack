//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package cloudstack

import (
	"fmt"
	"strings"
	"testing"

	"github.com/apache/cloudstack-go/v2/cloudstack"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudStackSecurityGroupRule_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudStackSecurityGroupRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudStackSecurityGroupRule_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudStackSecurityGroupRulesExist("cloudstack_security_group.foo"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.#", "2"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.0.cidr_list.0", "172.18.100.0/24"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.0.protocol", "tcp"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.0.ports.#", "1"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.0.ports.0", "80"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.0.traffic_type", "ingress"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.1.protocol", "tcp"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.1.ports.1", "80"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.1.ports.0", "443"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.1.traffic_type", "egress"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.1.user_security_group_list.0", "terraform-security-group-bar"),
				),
			},
		},
	})
}

func TestAccCloudStackSecurityGroupRule_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudStackSecurityGroupRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudStackSecurityGroupRule_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudStackSecurityGroupRulesExist("cloudstack_security_group.foo"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.#", "2"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.0.cidr_list.0", "172.18.100.0/24"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.0.protocol", "tcp"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.0.ports.#", "1"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.0.ports.0", "80"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.0.traffic_type", "ingress"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.1.protocol", "tcp"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.1.ports.1", "80"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.1.ports.0", "443"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.1.traffic_type", "egress"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.1.user_security_group_list.0", "terraform-security-group-bar"),
				),
			},

			{
				Config: testAccCloudStackSecurityGroupRule_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudStackSecurityGroupRulesExist("cloudstack_security_group.foo"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.#", "3"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.0.cidr_list.0", "172.18.100.0/24"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.0.cidr_list.1", "172.18.200.0/24"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.0.protocol", "tcp"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.0.ports.1", "80"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.0.ports.0", "443"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.1.cidr_list.#", "1"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.1.cidr_list.0", "172.18.100.0/24"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.1.icmp_code", "-1"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.1.icmp_type", "-1"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.2.protocol", "tcp"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.2.ports.#", "1"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.2.ports.0", "80"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.2.traffic_type", "egress"),
					resource.TestCheckResourceAttr(
						"cloudstack_security_group_rule.foo", "rule.2.user_security_group_list.0", "terraform-security-group-bar"),
				),
			},
		},
	})
}

func testAccCheckCloudStackSecurityGroupRulesExist(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No security group rule ID is set")
		}

		cs := testAccProvider.Meta().(*cloudstack.CloudStackClient)
		sg, count, err := cs.SecurityGroup.GetSecurityGroupByID(rs.Primary.ID)
		if err != nil {
			if count == 0 {
				return fmt.Errorf("Security group %s not found", rs.Primary.ID)
			}
			return err
		}

		// Make a map of all the rule indexes so we can easily find a rule
		sgRules := append(sg.Ingressrule, sg.Egressrule...)
		ruleIndex := make(map[string]int, len(sgRules))
		for idx, r := range sgRules {
			ruleIndex[r.Ruleid] = idx
		}

		for k, id := range rs.Primary.Attributes {
			if !strings.Contains(k, ".uuids.") || strings.HasSuffix(k, ".uuids.%") {
				continue
			}

			if _, ok := ruleIndex[id]; !ok {
				return fmt.Errorf("Security group rule %s not found", id)
			}
		}

		return nil
	}
}

func testAccCheckCloudStackSecurityGroupRuleDestroy(s *terraform.State) error {
	cs := testAccProvider.Meta().(*cloudstack.CloudStackClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudstack_security_group_rule" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No security group rule ID is set")
		}

		sg, count, err := cs.SecurityGroup.GetSecurityGroupByID(rs.Primary.ID)
		if err != nil {
			if count == 0 {
				continue
			}
			return err
		}

		// Make a map of all the rule indexes so we can easily find a rule
		sgRules := append(sg.Ingressrule, sg.Egressrule...)
		ruleIndex := make(map[string]int, len(sgRules))
		for idx, r := range sgRules {
			ruleIndex[r.Ruleid] = idx
		}

		for k, id := range rs.Primary.Attributes {
			if !strings.Contains(k, ".uuids.") || strings.HasSuffix(k, ".uuids.%") {
				continue
			}

			if _, ok := ruleIndex[id]; ok {
				return fmt.Errorf("Security group rule %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

const testAccCloudStackSecurityGroupRule_basic = `
resource "cloudstack_security_group" "foo" {
  name = "terraform-security-group-foo"
  description = "terraform-security-group-text"
}

resource "cloudstack_security_group" "bar" {
  name = "terraform-security-group-bar"
  description = "terraform-security-group-text"
}

resource "cloudstack_security_group_rule" "foo" {
  security_group_id = cloudstack_security_group.foo.id

  rule {
    cidr_list = ["172.18.100.0/24"]
    protocol = "tcp"
		ports = ["80"]
  }

  rule {
    protocol = "tcp"
    ports = ["80", "443"]
    traffic_type = "egress"
		user_security_group_list = ["terraform-security-group-bar"]
  }

	depends_on = ["cloudstack_security_group.bar"]
}`

const testAccCloudStackSecurityGroupRule_update = `
resource "cloudstack_security_group" "foo" {
  name = "terraform-security-group-foo"
  description = "terraform-security-group-text"
}

resource "cloudstack_security_group" "bar" {
  name = "terraform-security-group-bar"
  description = "terraform-security-group-text"
}

resource "cloudstack_security_group_rule" "foo" {
  security_group_id = cloudstack_security_group.foo.id

  rule {
    cidr_list = ["172.18.100.0/24", "172.18.200.0/24"]
    protocol = "tcp"
		ports = ["80", "443"]
  }

  rule {
    cidr_list = ["172.18.100.0/24"]
    protocol = "icmp"
    icmp_type = "-1"
    icmp_code = "-1"
    traffic_type = "ingress"
  }

  rule {
    protocol = "tcp"
    ports = ["80"]
    traffic_type = "egress"
		user_security_group_list = ["terraform-security-group-bar"]
  }

	depends_on = ["cloudstack_security_group.bar"]
}`
