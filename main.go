package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"is/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println('b')

	defaultDir, err := os.Getwd()
	if err != nil {
		log.Error("Could not find the root directory")
	}

	defaultConfg := "golp.config.yaml"

	// take in flags for the filepath or default to rootDir/
	rootDir := flag.String("dir", defaultDir, "The root path to be watched")
	configFileName := flag.String("config", defaultConfig, "The name of the config file to be watched")

	flag.Parse()
	remainder := flag.Args()

	// parse config file
	err = parseConfig(rootDir, configFileName)

	// walk returns either a msg to changeChan or errChan
	go func() {
		change, err := walk(config)
	}()
}

type pipelineConfig struct {
	k8sDir    string `yaml:"k8sDir"`
	dockerDir string `yaml:"dockerDir"`
}

func parseConfig(rootDir string, configFilePath string) (config pipelineConfig, err error) {

	configBytes, err := ioutil.ReadFile(strings.Join(rootDir, configFilePath))
	if err != nil {
		log.Errorf("Could not read config file because: %v\n", err)
		return
	}

	err = yaml.Unmarshal(configBytes, config)
	if err != nil {
		log.Errorf("Trouble while trying to convert the confiuration into a go object: %v\n", err)
	}

	return
}

type errChan <-chan error
type changeChan <-chan map[string]bool

// walk directories and watch for changes
// run the watcher in a poll loop
// type WalkFunc(path string, info os.FileInfo, err error) error
func (pc pipelineConfig) walk(rootDir string, k8sDir string, k8sPattern string, dockerWalker, k8sWalker filepath.WalkFunc) (change changeChan, err errChan) {
	// filepath.Walk(rootdirectory, walkfunc())
	//filepath.Match func Match(pattern, name string) (matched bool, err error)
	// match for files in folders named k8
	isK8sMatch, err := filepath.Match(rootDir+k8sPattern, k8sDir)
	if err != nil {
		return err
	}

	// match file prefixes/subfixes for Dockerfile.
	isDockerMatch, err := filepath.Match(rootDir+DockerPattern, DockerDir)
	if err != nil {
		return err
	}
}

// this is one of the walkFunc possibilities
// triggers if the docker prefixed/subfixes files change
func (pc pipelineConfig) rebuildDocker(path string, info os.FileInfo, err error) error {
	return nil
}

// this is anothe one of the walkFunc possibilities
// triggers if k8s have changed
func (pc pipelineConfig) reapplyK8s(path string, info os.FileInfo, err error) error {
	return nil
}
