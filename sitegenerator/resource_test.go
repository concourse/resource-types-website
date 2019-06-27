package sitegenerator_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/concourse/dutyfree/sitegenerator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"gopkg.in/yaml.v2"
)

var _ = Describe("Resources", func() {
	It("extracts all the values", func() {
		var r sitegenerator.Resource
		err := yaml.Unmarshal([]byte(`---
name: time resource
repository: https://github.com/concourse/time-resource
desc: time resource description
categories: [Category1, Category3]
get: 
put: yes
verified: Yes
`), &r)

		Expect(err).ToNot(HaveOccurred())
		Expect(r).To(MatchAllFields(Fields{
			"Name":        Equal("time resource"),
			"Repository":  Equal("https://github.com/concourse/time-resource"),
			"Description": Equal("time resource description"),
			"Categories":  Equal([]string{"Category1", "Category3"}),
			"Get":         Equal(false),
			"Put":         Equal(true),
			"Verified":    Equal(true),
		}))
	})

	It("fails if repository is not part of github and has less than 5 components", func() {
		resources := []sitegenerator.Resource{
			{Name: "time resource", Repository: "https://github.com/concourse"},
		}
		_, err := sitegenerator.Enrich(resources, sitegenerator.HttpReadmeClient{})

		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring(fmt.Sprintf("invalid repository for the resource (%s)", "time resource")))

	})

	It("calculates the additional fields", func() {
		resources := []sitegenerator.Resource{
			{Name: "time resource", Repository: "https://github.com/concourse/time-resource"},
		}

		i, err := sitegenerator.Enrich(resources, sitegenerator.HttpReadmeClient{
			GetReadme: func(url string) (*http.Response, error) {
				Expect(url).To(Equal("https://api.github.com/repos/concourse/time-resource/readme"))
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(strings.NewReader("foobar readme")),
				}, nil
			},
		})

		Expect(err).ToNot(HaveOccurred())

		Expect(i[0]).To(Equal(sitegenerator.ResourceModel{
			Resource: sitegenerator.Resource{
				Name:       "time resource",
				Repository: "https://github.com/concourse/time-resource",
			},
			Repo:              "time-resource",
			Identifier:        "concourse-time-resource",
			AuthorHandle:      "concourse",
			AuthorProfileLink: "https://github.com/concourse",
			Readme:            "foobar readme",
		}))
	})

	It("fails to retrieve readme", func() {
		resources := []sitegenerator.Resource{
			{Name: "time resource", Repository: "https://github.com/concourse/time-resource"},
		}

		_, err := sitegenerator.Enrich(resources, sitegenerator.HttpReadmeClient{
			GetReadme: func(url string) (*http.Response, error) {
				Expect(url).To(Equal("https://api.github.com/repos/concourse/time-resource/readme"))
				return &http.Response{
					StatusCode: 404,
					Body:       ioutil.NopCloser(strings.NewReader("foobar readme")),
				}, nil
			},
		})

		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("Unable to access readme for concourse/time-resource due to error: 404, reason: foobar readme"))

	})
})
