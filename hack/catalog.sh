#!/bin/bash

CSV_FILE=$1
OPERATOR_VERSION=$2
DIR=${DIR:-$(cd $(dirname "$0")/.. && pwd)}
NAME=${NAME:-$(ls $DIR/deploy/olm-catalog)}
# Use the latest version of operator for testing

x=( $(echo $NAME | tr '-' ' ') )
DISPLAYNAME=${DISPLAYNAME:=${x[*]^}}

indent() {
  INDENT="      "
  sed "s/^/$INDENT/" | sed "s/^${INDENT}\($1\)/${INDENT:0:-2}- \1/"
}

CRD=$(cat $(find $DIR/deploy/olm-catalog/$NAME/$OPERATOR_VERSION -name '*_crd.yaml' | sort -n) | grep -v -- "---" | indent apiVersion)
CSV=$(cat $CSV_FILE | indent apiVersion)

get_package_current_csv_name () {
    if [ X"${OPERATOR_VERSION}" == X"0.1.0" ]; then
	echo "${NAME}.v${OPERATOR_VERSION}"
    else
	echo "${NAME}.v${OPERATOR_VERSION}.0"
    fi
}

PKG_VERSION=$(get_package_current_csv_name)

replace_package_version() {
    sed "s,currentCSV: .*,currentCSV: ${PKG_VERSION},"
}

PKG=$(cat $DIR/deploy/olm-catalog/$NAME/*package.yaml | replace_package_version | indent packageName)

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
