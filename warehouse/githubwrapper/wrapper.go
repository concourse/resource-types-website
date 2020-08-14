package githubwrapper

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

//go:generate counterfeiter . Wrapper

type Wrapper interface {
	GetStars(map[string]bool) (map[string]int, error)
}

type wrapper struct {
	ServerUrl string
	Token     string
}

func NewWrapper(serverUrl, token string) wrapper {
	return wrapper{
		ServerUrl: serverUrl,
		Token:     token,
	}
}

// the function does a workaround the github graphql
// the grapphql expects a generic struct to be passed to it.
// in our case we have no prior knowledge of some fields of the struct, or their values
// that's why we used reflections to construct a struct object in the same way that is expected by the graphql.
func (w wrapper) GetStars(repoStarsMap map[string]bool) (map[string]int, error) {
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
	for k, isGithub := range repoStarsMap {
		if isGithub {
			strs := strings.Split(k, "/")
			arr = append(arr, reflect.StructField{
				Name: fmt.Sprintf("I%d", i),
				Type: ghReturnReflection,
				Tag:  reflect.StructTag(fmt.Sprintf(`graphql:"i%d: repository(owner: \"%s\", name: \"%s\")"`, i, strs[0], strs[1])),
			})
			i++
		}
	}

	gqlQuery := reflect.New(reflect.StructOf(arr)).Elem()
	err := ghClient.Query(context.TODO(), gqlQuery.Addr().Interface(), nil)

	if err != nil {
		return nil, err
	}

	i = 0

	returnMap := make(map[string]int)
	for _, isGithub := range repoStarsMap {
		if isGithub {
			stars := reflect.ValueOf(gqlQuery.Interface()).FieldByIndex([]int{i}).FieldByName("Stargazers").FieldByIndex([]int{0}).Interface().(int)
			owner := reflect.ValueOf(gqlQuery.Interface()).FieldByIndex([]int{i}).FieldByName("NameWithOwner").Interface().(string)
			returnMap[owner] = stars
			i++
		}
	}
	return returnMap, nil
}
