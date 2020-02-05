package githubwrapper

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

//go:generate counterfeiter . WrapperInterface

type WrapperInterface interface {
	GetStars([]RepoStars) error
}

type Wrapper struct {
	ServerUrl string
	Token     string
}

func (w *Wrapper) GetStars(repoStarsMap map[string]int) error {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: w.Token},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	ghClient := githubv4.NewEnterpriseClient(w.ServerUrl, httpClient)

	var arr []reflect.StructField

	type ghReturn struct {
		NameWithOwner string
		Stargazers    struct {
			TotalCount int
		}
	}
	ghReturnReflection := reflect.TypeOf(ghReturn{})

	i := 0
	for k := range repoStarsMap {
		strs := strings.Split(k, "/")
		arr = append(arr, reflect.StructField{
			Name: fmt.Sprintf("I%d", i),
			Type: ghReturnReflection,
			Tag:  reflect.StructTag(fmt.Sprintf(`graphql:"i%d: repository(owner: \"%s\", name: \"%s\")"`, i, strs[0], strs[1])),
		})
		i++
	}

	gqlQuery := reflect.New(reflect.StructOf(arr)).Elem()
	err := ghClient.Query(context.TODO(), gqlQuery.Addr().Interface(), nil)

	if err != nil {
		return err
	}

	i = 0
	for range repoStarsMap {
		stars := reflect.ValueOf(gqlQuery.Interface()).FieldByIndex([]int{i}).FieldByName("Stargazers").FieldByIndex([]int{0}).Interface().(int)
		owner := reflect.ValueOf(gqlQuery.Interface()).FieldByIndex([]int{i}).FieldByName("NameWithOwner").Interface().(string)
		repoStarsMap[owner] = stars
		i++
	}
	return nil
}
