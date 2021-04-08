// +build mage

// Manage the Docker volume plugin 'local-persist'.
package main

import (
	"fmt"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"log"
	"os"
	"strings"
)

type Docker mg.Namespace

var Aliases = map[string]interface{}{
	"i": Install,
}

var pwd = "."
var name = "local-persist"
var goOS = "linux"
var goArch = "amd64"

// Download the modules
func Build() error {
	if err := sh.Run("go", "mod", "download"); err != nil {
		return err
	}
	return sh.Run("go", "install", "./...")
}

// Install onto the current host.
func Install() error {
	return nil
}

// Show the current environment.
func Environment() error {
	fmt.Println()
	for ix, ep := range os.Environ() {
		pair := strings.SplitN(ep, "=", 2)
		fmt.Printf("[%d] %s = %s\n", ix, pair[0], pair[1])
	}
	return nil
}

// Determine the test coverage.
// GO_ENV=test
//  go test -v -coverprofile=coverage.out ./... &&
//  sed -i '' 's|'_$(PWD)'|.|g' coverage.out &&
//  go tool cover -html=coverage.out
func Coverage() error {
	log.Println("coverage running")

	os.Setenv("GO_ENV", "test")

	goTest, goTestErr := sh.Output("go", "test", "-v", "-coverprofile=coverage.out", "./...")
	fmt.Println("coverage: %s", goTest)
	if goTestErr != nil {
		return mg.Fatalf(1, "coverage error %s", goTestErr)
	}
	//f, err := os.Open("coverage.out")
	//if err != nil {
	//    return err
	//}
	//defer f.Close()
	//
	//// Splits on newlines by default.
	//scanner := bufio.NewScanner(f)
	//
	//line := 1
	//// https://golang.org/pkg/bufio/#Scanner.Scan
	//for scanner.Scan() {
	//    if strings.Contains(scanner.Text(), pwd) {
	//        return mg.Fatalf(1, "%s", pwd)
	//    }
	//    line++
	//}
	//if err := scanner.Err(); err != nil {
	//    return
	//}

	goTool, goToolErr := sh.Output("go", "tool", "cover", "-html=coverage.out")
	fmt.Println("coverage: %s", goTool)
	return goToolErr
}

// Run the test suite.
// test: export GO15VENDOREXPERIMENT=1
// test: export GO_ENV=test
// test:
//   go test -v .
func Test() error {
	log.Println("test running")

	os.Setenv("GO15VENDOREXPERIMENT", "1")
	os.Setenv("GO_ENV", "test")

	goTest, goTestErr := sh.Output("go", "test", "-v", ".")
	fmt.Println("test: %s", goTest)
	return goTestErr
}

// sudo -E go run main.go driver.go
func Run() error {
	log.Println("run running")
	goTest, goTestErr := sh.Output("go", "run", "main.go", "driver.go")
	fmt.Println("run: %s", goTest)
	return goTestErr
}

// Run the docker image in a container.
// docker run -d
//   -v /run/docker/plugins/:/run/docker/plugins/
//   -v ${DATA_VOLUME}:${DATA_VOLUME}
//   babeloff/docker-local-persist-volume-plugin
func (Docker) Run(dataVolume string) error {
	log.Println("docker-run running")
	if len(dataVolume) < 1 {
		return mg.Fatalf(1,
			"echo Missing required environment variable: %s", dataVolume)
	}
	goTest, goTestErr := sh.Output("docker", "run", "-d",
		"-v", "/run/docker/plugins/:/run/docker/plugins/",
		"-v", fmt.Sprint("%s:%s", dataVolume, dataVolume),
		"babeloff/docker-local-persist-volume-plugin")

	fmt.Println("docker:run: %s", goTest)
	return goTestErr
}

// Build the docker image.
//  LOCAL_PERSIST=$(docker build -q -f Dockerfile-build .)
//  docker run -it -v `pwd`/bin:/go/src/local-persist/bin $LOCAL_PERSIST
func (Docker) Build() error {
	mg.Deps(Clean)
	log.Println("docker-build running")
	localPersist, localPersistErr := sh.Output("docker", "build", "-q",
		"-f", "Dockerfile-build", ".")
	fmt.Println("docker:build: %s", localPersist)
	if localPersistErr != nil {
		return mg.Fatalf(1, "coverage error %s", localPersistErr)
	}

	goTest, goTestErr := sh.Output("docker", "run", "-id",
		"-v", fmt.Sprint("%s/bin:/go/src/local-persist/bin", pwd),
		localPersist)

	log.Printf("docker:build: %s\n", goTest)
	return goTestErr
}

// build for current architecture
//   go build -o bin/$(BIN_NAME) -v
func Compile() error {
	fmt.Println("compile running")
	os.Setenv("GO15VENDOREXPERIMENT", "1")
	goBuild, goBuildErr := sh.Output("go", "build",
		"-o", fmt.Sprint("bin/%s", name),
		"-v")
	fmt.Printf("compile: %s\n", goBuild)
	return goBuildErr
}

// build all the binaries:
// clean-bin compile-linux-amd64
func Binaries() error {
	mg.Deps(Clean, Compile_linux_amd64)
	log.Println("binaries running")
	return nil
}

// clean:
// 	rm -Rf bin/*
func Clean() error {
	log.Println("clean running")
	// return os.RemoveAll("./bin/")
	return sh.Rm("./bin/")
}

// build compile-linux-amd64:
//  go build -o bin/$(GOOS)/$(GOARCH)/$(BIN_NAME) -v
func Compile_linux_amd64() error {
	log.Println("compile-linux-amd64 running")

	os.Setenv("GOOS", goOS)
	os.Setenv("GOARCH", goArch)
	os.Setenv("GO15VENDOREXPERIMENT", "1")

	goTest, goTestErr := sh.Output("go", "build",
		"-o", fmt.Sprint("bin/%s/%s/%s", goOS, goArch, name),
		"-v")
	fmt.Println("compile linux amd64: %s", goTest)
	return goTestErr
}

// release:
// ./scripts/release.sh
func (Docker) Release() error {
	mg.Deps(Build)
	log.Println("release running")
	//    #!/usr/bin/env bash
	//
	//    black='\033[0;30m'        # Black
	//    red='\033[0;31m'          # Red
	//    green='\033[0;32m'        # Green
	//    yellow='\033[0;33m'       # Yellow
	//    blue='\033[0;34m'         # Blue
	//    purple='\033[0;35m'       # Purple
	//    cyan='\033[0;36m'         # Cyan
	//    white='\033[0;37m'        # White
	//    nocolor='\033[0m'         # No Color
	//
	//
	//    USER=cwspear
	//    REPO=local-persist
	//
	//    # check to make sure github-release is installed!
	//        github-release --version > /dev/null || exit
	//
	//    if [[ $RELEASE_NAME == "" ]]; then
	//echo -e ${cyan}Enter release name:${nocolor}
	//read RELEASE_NAME
	//echo ''
	//fi
	//
	//if [[ $RELEASE_DESCRIPTION == "" ]]; then
	//echo -e ${cyan}Enter release description:${nocolor}
	//read RELEASE_DESCRIPTION
	//echo ''
	//fi
	//
	//if [[ $RELEASE_TAG == "" ]]; then
	//printf "${cyan}Enter release tag:${nocolor} v"
	//read RELEASE_TAG
	//echo ''
	//
	//RELEASE_TAG="v${RELEASE_TAG}"
	//fi
	//
	//if [[ $PRERELEASE == "" ]]; then
	//printf "${cyan}Is this a prerelease? [yN]${nocolor} "
	//read PRERELEASE_REPLY
	//echo ''
	//
	//FIRST=${PRERELEASE_REPLY:0:1}
	//echo $FIRST
	//
	//PRERELEASE=false
	//[[ $FIRST == 'Y' || $FIRST == 'y' ]] && PRERELEASE=true
	//fi
	//
	//sed -i '' "s|VERSION=\".*\"|VERSION=\"${RELEASE_TAG}\"|" scripts/install.sh
	//sed -i '' "s|ENV VERSION .*|ENV VERSION ${RELEASE_TAG}|" Dockerfile
	//
	//git commit -am "Tagged ${RELEASE_TAG}"
	//git push
	//git tag -a $RELEASE_TAG -m "$RELEASE_NAME"
	//git push --tags
	//
	//echo ''
	//echo Releasing...
	//echo ''
	//echo USER=$USER
	//echo REPO=$REPO
	//echo RELEASE_NAME="'$RELEASE_NAME'"
	//echo RELEASE_DESCRIPTION="'$RELEASE_DESCRIPTION'"
	//echo RELEASE_TAG=$RELEASE_TAG
	//echo PRERELEASE=$PRERELEASE
	//echo ''
	//
	//if [[ "$PRERELEASE" == true ]]; then
	//github-release release \
	//--user $USER \
	//--repo $REPO \
	//--tag $RELEASE_TAG \
	//--name "$RELEASE_NAME" \
	//--description "$RELEASE_DESCRIPTION" \
	//--pre-release
	//else
	//github-release release \
	//--user $USER \
	//--repo $REPO \
	//--tag $RELEASE_TAG \
	//--name "$RELEASE_NAME" \
	//--description "$RELEASE_DESCRIPTION"
	//fi
	//
	//echo Uploading binaries...
	//for FILE in `find bin -type f`; do
	//NAME=${FILE/bin\//}
	//NAME=${NAME//\//-}
	//NAME=`echo $NAME | sed 's/\(.*\)-local-persist/local-persist-\1/'`
	//
	//echo Uploading ${NAME}...
	//
	//if [[ $NAME != "" ]]; then
	//github-release upload \
	//--user $USER \
	//--repo $REPO \
	//--tag $RELEASE_TAG \
	//--name $NAME \
	//--file $FILE
	//fi
	//done

	return nil
}
