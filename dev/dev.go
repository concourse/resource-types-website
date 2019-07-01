package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/concourse/dutyfree/sitegenerator"
	"github.com/onsi/gomega/gexec"
	"github.com/onsi/gomega/ghttp"
	"gopkg.in/yaml.v2"
)

var (
	server    *ghttp.Server
	outputDir string
	resources *os.File
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		log.Println(sig)
		cleanup()
		done <- true
	}()

	log.Println("===============================================================")
	log.Println("dev mode")

	pathToBin, err := gexec.Build("github.com/concourse/dutyfree")

	if err != nil {
		fmt.Printf("error: %s\n", err)
	}

	outputDir, err = ioutil.TempDir("", "dutyfree")

	if err != nil {
		fmt.Printf("error: %s\n", err)
	}

	resources, err = ioutil.TempFile("", "resources.yml")

	if err != nil {
		fmt.Printf("error: %s\n", err)
	}

	_, err = fmt.Fprint(resources, `---
- repository: https://github.com/concourse/git-resource
  name: git resource
  desc: git resource description
  categories: [Version Control, Deployment, category with space]
  get: yes
  put: yes
  verified: Yes
  example: |-
    ---
    jobs:
    - name: alert
      public: true
      plan:
      - put: errand-prod
        params:
          manifest: prod.yml
          errand: smoke_tests

    resources:
    - name: errand-prod
      type: bosh-errand
      source:
        target: {{bosh-target}}
        username: admin
        password: admin
        deployment: cf-prod
- repository: https://github.com/concourse/hg-resource
  name: hg resource
  desc: 
  get: 
  put: false
  verified: Yes
- repository: https://github.com/concourse/foo-resource
  name: foo resource
  categories: [Notifications, cat5, cat42, Deployment]
  put: yes
  verified: Yes
- repository: https://github.com/pivotal-cf/bosh-resource
  name: bosh resource
  desc: a resource for all the bosh
  get: yes
  verified: No
`)

	log.Println("starting mock server")
	server = ghttp.NewServer()
	githubMockServer(resources.Name(), server)

	if err != nil {
		fmt.Printf("error: %s\n", err)
	}

	cmd := exec.Command(pathToBin, outputDir, resources.Name())
	cmd.Env = append(cmd.Env, "GITHUB_API_ENDPOINT="+server.URL(), "GITHUB_TOKEN=SOMEGITHUBTOKEN")

	session, err := gexec.Start(cmd, os.Stdin, os.Stderr)

	if err != nil {
		fmt.Printf("error: %s\n", err)
	}

	<-session.Exited
	server.Close()

	srv := startHttpServer(outputDir, ":3000")
	log.Println("Listening... on port 3000")
	log.Println("http://localhost:3000/dutyfree")

	<-done

	srv.Shutdown(context.TODO())
}

func cleanup() {
	log.Println("cleaning up")

	gexec.CleanupBuildArtifacts()

	if _, err := os.Stat(outputDir); !os.IsNotExist(err) {
		os.RemoveAll(outputDir)
	}
	if resources != nil {
		if _, err := os.Stat(resources.Name()); !os.IsNotExist(err) {
			os.Remove(resources.Name())
		}
	}
}

func startHttpServer(outputDir, port string) *http.Server {
	srv := &http.Server{Addr: port}

	fs := http.StripPrefix("/dutyfree/", http.FileServer(http.Dir(outputDir)))
	http.Handle("/dutyfree/", fs)

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {

			log.Fatalf("ListenAndServe(): %s", err)
		}
	}()

	return srv
}

func githubMockServer(resourcesPath string, server *ghttp.Server) {
	var resources []sitegenerator.Resource
	resourceFile, err := os.Open(resourcesPath)
	if err != nil {
		panic(err)
	}

	decoder := yaml.NewDecoder(resourceFile)
	err = decoder.Decode(&resources)
	if err != nil {
		panic(err)
	}

	content, err := ioutil.ReadFile("dev/git-resource-readme.html")

	if err != nil {
		panic(err)
	}

	for range resources {
		server.AppendHandlers(
			ghttp.RespondWith(http.StatusOK, content),
		)
	}

}
