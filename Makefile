# This how we want to name the binary output
BINARY=stx-node-exporter

# VERSION=`git describe --tags`
VERSION=0.0
#VERSION=$(git log --pretty=format:'%h' -n 1)
#BUILD=$(date +%FT%T%z)
BUILD=TEST

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-w -s -X version.Version=${VERSION} -X version.BuildDate=${BUILD} -X version.BuildUser=jos -X version.Branch=fred -X version.Revision=2"

# Builds the project
build:
	go build -a ${LDFLAGS} -o ${BINARY} cmd/stx-node-exporter/stx-node-exporter.go

# Installs our project: copies binaries
install:
	go install ${LDFLAGS}

# Cleans our project: deletes binaries
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: clean install
