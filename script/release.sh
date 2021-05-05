#!/bin/sh

if [ $# -ne 1 ]; then
  echo "usage: <version>"
  exit 1
fi

version=$1

cat <<__TEXT__ > cmd/version.go
package cmd

var VERSION = "$version"
__TEXT__


git add cmd/version.go
git commit -m ":up: bump up $version"
git tag $version

