package persistence

import "github.com/concourse/dutyfree/resource"

type persistence interface {
	GetResource(name string) (resource.Resource, error)
	GetAllResources() ([]resource.Resource, error)

	WriteResource(r resource.Resource) error
	WriteBulkResources(r []resource.Resource) error
}
