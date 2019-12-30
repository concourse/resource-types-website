package persistence

import "github.com/concourse/dutyfree/resource"

type Persistence interface {
	//GetResource(name string) (resource.Resource, error)
	GetAllResources() []resource.Resource

	//WriteResource(r resource.Resource) error
	//WriteBulkResources(r []resource.Resource) error
}
