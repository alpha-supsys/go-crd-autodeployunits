#!/bin/bash

$GOPATH/src/k8s.io/code-generator/generate-groups.sh all \
github.com/alpha-supsys/go-crd-autodeployunits/k8s/group/client \
github.com/alpha-supsys/go-crd-autodeployunits/k8s/group/apis \
"autodeploy:v1" -v 10