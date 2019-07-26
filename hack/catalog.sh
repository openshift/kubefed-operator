#!/bin/bash

CSV_FILE=$1
DIR=${DIR:-$(cd $(dirname "$0")/.. && pwd)}
NAME=${NAME:-$(ls $DIR/deploy/olm-catalog)}
# Use the latest version of operator for testing
OPERATOR_VERSION=${OPERATOR_VERSION:-0.1.0}

x=( $(echo $NAME | tr '-' ' ') )
DISPLAYNAME=${DISPLAYNAME:=${x[*]^}}

indent() {
  INDENT="      "
  sed "s/^/$INDENT/" | sed "s/^${INDENT}\($1\)/${INDENT:0:-2}- \1/"
}

CRD=$(cat $(find $DIR/deploy/olm-catalog/$NAME/$OPERATOR_VERSION -name '*_crd.yaml' | sort -n) | grep -v -- "---" | indent apiVersion)
CSV=$(cat $CSV_FILE | indent apiVersion)
PKG=$(cat $DIR/deploy/olm-catalog/$NAME/*package.yaml | indent packageName)

cat <<EOF | sed 's/^  *$//'
kind: ConfigMap
apiVersion: v1
metadata:
  name: $NAME

data:
  customResourceDefinitions: |-
$CRD
  clusterServiceVersions: |-
$CSV
  packages: |-
$PKG
---
apiVersion: operators.coreos.com/v1alpha1
kind: CatalogSource
metadata:
  name: $NAME
spec:
  configMap: $NAME
  displayName: $DISPLAYNAME
  publisher: Red Hat
  sourceType: internal
EOF
