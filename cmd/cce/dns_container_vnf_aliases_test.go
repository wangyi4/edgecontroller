// Copyright 2019 Smart-Edge.com, Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	cce "github.com/smartedgemec/controller-ce"
	"github.com/smartedgemec/controller-ce/uuid"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("/dns_container_vnf_aliases", func() {
	var (
		containerVNFID string
	)

	BeforeEach(func() {
		containerVNFID = postContainerVNFs()
	})

	Describe("POST /dns_container_vnf_aliases", func() {
		DescribeTable("201 Created",
			func(req string) {
				By("Sending a POST /dns_container_vnf_aliases request")
				resp, err := http.Post(
					"http://127.0.0.1:8080/dns_container_vnf_aliases",
					"application/json",
					strings.NewReader(fmt.Sprintf(req, containerVNFID)))
				Expect(err).ToNot(HaveOccurred())

				By("Verifying a 201 response")
				Expect(resp.StatusCode).To(Equal(http.StatusCreated))

				By("Reading the response body")
				body, err := ioutil.ReadAll(resp.Body)
				Expect(err).ToNot(HaveOccurred())

				var respBody struct {
					ID string
				}

				By("Unmarshalling the response")
				Expect(json.Unmarshal(body, &respBody)).To(Succeed())

				By("Verifying a UUID was returned")
				Expect(uuid.IsValid(respBody.ID)).To(BeTrue())
			},
			Entry(
				"POST /dns_container_vnf_aliases",
				`
				{
					"name": "dns container vnf alias 123",
					"description": "description 1",
					"container_vnf_id": "%s"
				}`),
		)

		DescribeTable("400 Bad Request",
			func(req, expectedResp string) {
				By("Sending a POST /dns_container_vnf_aliases request")
				resp, err := http.Post(
					"http://127.0.0.1:8080/dns_container_vnf_aliases",
					"application/json",
					strings.NewReader(req))
				Expect(err).ToNot(HaveOccurred())

				By("Verifying a 400 Bad Request response")
				Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))

				By("Reading the response body")
				body, err := ioutil.ReadAll(resp.Body)
				Expect(err).ToNot(HaveOccurred())

				By("Verifying the response body")
				Expect(string(body)).To(Equal(expectedResp))
			},
			Entry(
				"POST /dns_container_vnf_aliases with id",
				`
				{
					"id": "123"
				}`,
				"Validation failed: id cannot be specified in POST request"),
			Entry(
				"POST /dns_container_vnf_aliases without name",
				`
				{
					"description": "description 1",
					"container_vnf_id": "123"
				}`,
				"Validation failed: name cannot be empty"),
			Entry(
				"POST /dns_container_vnf_aliases without description",
				`
				{
					"name": "dns container vnf alias 123",
					"container_vnf_id": "123"
				}`,
				"Validation failed: description cannot be empty"),
			Entry(
				"POST /dns_container_vnf_aliases without container_vnf_id",
				`
				{
					"name": "dns container vnf alias 123",
					"description": "description 1"
				}`,
				"Validation failed: container_vnf_id not a valid uuid"),
		)
	})

	Describe("GET /dns_container_vnf_aliases", func() {
		var (
			dnsContainerVNFAliasID  string
			dnsContainerVNFAlias2ID string
		)

		BeforeEach(func() {
			dnsContainerVNFAliasID = postDNSContainerVNFAliases(containerVNFID)
			dnsContainerVNFAlias2ID = postDNSContainerVNFAliases(containerVNFID)
		})

		DescribeTable("200 OK",
			func() {
				By("Sending a GET /dns_container_vnf_aliases request")
				resp, err := http.Get(
					"http://127.0.0.1:8080/dns_container_vnf_aliases")

				By("Verifying a 200 OK response")
				Expect(err).ToNot(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(http.StatusOK))

				By("Reading the response body")
				body, err := ioutil.ReadAll(resp.Body)
				Expect(err).ToNot(HaveOccurred())

				var dnsContainerVNFAliases []cce.DNSContainerVNFAlias

				By("Unmarshalling the response")
				Expect(json.Unmarshal(body, &dnsContainerVNFAliases)).
					To(Succeed())

				By("Verifying the 2 created DNS container VNF aliases were returned") //nolint:lll
				Expect(dnsContainerVNFAliases).To(ContainElement(
					cce.DNSContainerVNFAlias{
						ID:             dnsContainerVNFAliasID,
						Name:           "dns container vnf alias 123",
						Description:    "description 1",
						ContainerVNFID: containerVNFID,
					}))
				Expect(dnsContainerVNFAliases).To(ContainElement(
					cce.DNSContainerVNFAlias{
						ID:             dnsContainerVNFAlias2ID,
						Name:           "dns container vnf alias 123",
						Description:    "description 1",
						ContainerVNFID: containerVNFID,
					}))
			},
			Entry("GET /dns_container_vnf_aliases"),
		)
	})

	Describe("GET /dns_container_vnf_aliases/{id}", func() {
		var (
			dnsContainerVNFAliasID string
		)

		BeforeEach(func() {
			dnsContainerVNFAliasID = postDNSContainerVNFAliases(containerVNFID)
		})

		DescribeTable("200 OK",
			func() {
				dnsContainerVNFAlias := getDNSContainerVNFAlias(
					dnsContainerVNFAliasID)

				By("Verifying the created DNS container VNF alias was returned")
				Expect(dnsContainerVNFAlias).To(Equal(
					&cce.DNSContainerVNFAlias{
						ID:             dnsContainerVNFAliasID,
						Name:           "dns container vnf alias 123",
						Description:    "description 1",
						ContainerVNFID: containerVNFID,
					},
				))
			},
			Entry("GET /dns_container_vnf_aliases/{id}"),
		)

		DescribeTable("404 Not Found",
			func() {
				By("Sending a GET /dns_container_vnf_aliases/{id} request")
				resp, err := http.Get(
					fmt.Sprintf(
						"http://127.0.0.1:8080/dns_container_vnf_aliases/%s",
						uuid.New()))

				By("Verifying a 404 Not Found response")
				Expect(err).ToNot(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(http.StatusNotFound))
			},
			Entry("GET /dns_container_vnf_aliases/{id} with nonexistent ID"),
		)
	})

	Describe("PATCH /dns_container_vnf_aliases", func() {
		var (
			dnsContainerVNFAliasID string
		)

		BeforeEach(func() {
			dnsContainerVNFAliasID = postDNSContainerVNFAliases(containerVNFID)
		})

		DescribeTable("204 No Content",
			func(reqStr string, expectedAlias *cce.DNSContainerVNFAlias) {
				By("Sending a PATCH /dns_container_vnf_aliases request")
				req, err := http.NewRequest(
					http.MethodPatch,
					"http://127.0.0.1:8080/dns_container_vnf_aliases",
					strings.NewReader(fmt.Sprintf(reqStr,
						dnsContainerVNFAliasID, containerVNFID)))
				Expect(err).ToNot(HaveOccurred())

				c := http.Client{}
				resp, err := c.Do(req)
				Expect(err).ToNot(HaveOccurred())

				By("Verifying a 204 No Content response")
				Expect(err).ToNot(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(http.StatusNoContent))

				By("Getting the updated VNF alias")
				updatedAlias := getDNSContainerVNFAlias(dnsContainerVNFAliasID)

				By("Verifying the DNS container VNF alias was updated")
				expectedAlias.SetID(dnsContainerVNFAliasID)
				expectedAlias.ContainerVNFID = containerVNFID
				Expect(updatedAlias).To(Equal(expectedAlias))
			},
			Entry(
				"PATCH /dns_container_vnf_aliases",
				`
				[
					{
						"id": "%s",
						"name": "dns container vnf alias 123456",
						"description": "description 1",
						"container_vnf_id": "%s"
					}
				]`,
				&cce.DNSContainerVNFAlias{
					Name:        "dns container vnf alias 123456",
					Description: "description 1",
				}),
		)

		DescribeTable("400 Bad Request",
			func(reqStr string, expectedResp string) {
				By("Sending a PATCH /dns_container_vnf_aliases request")
				if strings.Contains(reqStr, "%s") {
					reqStr = fmt.Sprintf(reqStr, dnsContainerVNFAliasID)
				}
				req, err := http.NewRequest(
					http.MethodPatch,
					"http://127.0.0.1:8080/dns_container_vnf_aliases",
					strings.NewReader(reqStr))
				Expect(err).ToNot(HaveOccurred())

				c := http.Client{}
				resp, err := c.Do(req)

				By("Verifying a 400 Bad Request")
				Expect(err).ToNot(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))

				By("Reading the response body")
				body, err := ioutil.ReadAll(resp.Body)
				Expect(err).ToNot(HaveOccurred())

				By("Verifying the response body")
				Expect(string(body)).To(Equal(expectedResp))
			},
			Entry(
				"PATCH /dns_container_vnf_aliases without id",
				`
				[{}]`,
				"Validation failed: id not a valid uuid"),
			Entry(
				"PATCH /dns_container_vnf_aliases without name",
				`
				[
					{
						"id": "%s",
						"description": "description 1",
						"container_vnf_id": "123"
					}
				]`,
				"Validation failed: name cannot be empty"),
			Entry(
				"PATCH /dns_container_vnf_aliases without description",
				`
				[
					{
						"id": "%s",
						"name": "dns container vnf alias 123",
						"container_vnf_id": "123"
					}
				]`,
				"Validation failed: description cannot be empty"),
			Entry(
				"PATCH /dns_container_vnf_aliases without container_vnf_id",
				`
				[
					{
						"id": "%s",
						"name": "dns container vnf alias 123",
						"description": "description 1"
					}
				]`,
				"Validation failed: container_vnf_id not a valid uuid"),
		)
	})

	Describe("DELETE /dns_container_vnf_aliases/{id}", func() {
		var (
			dnsContainerVNFAliasID string
		)

		BeforeEach(func() {
			dnsContainerVNFAliasID = postDNSContainerVNFAliases(containerVNFID)
		})

		DescribeTable("200 OK",
			func() {
				By("Sending a DELETE /dns_container_vnf_aliases/{id} request")
				req, err := http.NewRequest(
					http.MethodDelete,
					fmt.Sprintf(
						"http://127.0.0.1:8080/dns_container_vnf_aliases/%s",
						dnsContainerVNFAliasID),
					nil)
				Expect(err).ToNot(HaveOccurred())

				c := http.Client{}
				resp, err := c.Do(req)
				Expect(err).ToNot(HaveOccurred())

				By("Verifying a 200 OK response")
				Expect(err).ToNot(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(http.StatusOK))

				By("Verifying the DNS container VNF alias was deleted")

				By("Sending a GET /dns_container_vnf_aliases/{id} request")
				resp, err = http.Get(
					fmt.Sprintf(
						"http://127.0.0.1:8080/dns_container_vnf_aliases/%s",
						dnsContainerVNFAliasID))

				By("Verifying a 404 Not Found response")
				Expect(err).ToNot(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(http.StatusNotFound))
			},
			Entry("DELETE /dns_container_vnf_aliases/{id}"),
		)

		DescribeTable("404 Not Found",
			func(id string) {
				By("Sending a DELETE /dns_container_vnf_aliases/{id} request")
				req, err := http.NewRequest(
					http.MethodDelete,
					fmt.Sprintf(
						"http://127.0.0.1:8080/dns_container_vnf_aliases/%s",
						id),
					nil)
				Expect(err).ToNot(HaveOccurred())

				c := http.Client{}
				resp, err := c.Do(req)
				Expect(err).ToNot(HaveOccurred())

				By("Verifying a 404 Not Found response")
				Expect(err).ToNot(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(http.StatusNotFound))
			},
			Entry(
				"DELETE /dns_container_vnf_aliases/{id} with nonexistent ID",
				uuid.New()),
		)

		DescribeTable("422 Unprocessable Entity",
			func() {
				postDNSConfigsDNSContainerVNFAliases(
					postDNSConfigs(),
					dnsContainerVNFAliasID)

				By("Sending a DELETE /dns_container_vnf_aliases/{id} request")
				req, err := http.NewRequest(
					http.MethodDelete,
					fmt.Sprintf(
						"http://127.0.0.1:8080/dns_container_vnf_aliases/%s",
						dnsContainerVNFAliasID),
					nil)
				Expect(err).ToNot(HaveOccurred())

				c := http.Client{}
				resp, err := c.Do(req)
				Expect(err).ToNot(HaveOccurred())

				By("Verifying a 422 response")
				Expect(resp.StatusCode).To(Equal(
					http.StatusUnprocessableEntity))

				By("Reading the response body")
				body, err := ioutil.ReadAll(resp.Body)
				Expect(err).ToNot(HaveOccurred())

				By("Verifying the response body")
				Expect(string(body)).To(Equal(fmt.Sprintf(
					"cannot delete dns_container_vnf_alias_id %s: record in "+
						"use in dns_configs_dns_container_vnf_aliases",
					dnsContainerVNFAliasID)))
			},
			Entry("DELETE /dns_container_vnf_aliases/{id} with dns_configs_dns_container_vnf_aliases record"), //nolint:lll
		)
	})
})