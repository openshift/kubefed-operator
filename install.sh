#!/bin/bash
# This script will create a namespace and deploy all the crds within the same
# namespace
# usage ./install.sh <location> <namespace>

set -e
 

NAMESPACE=""
LOCATION="local"
NAMESPACE_STR=""

while getopts “n:d:” opt; do
    case $opt in
	n) NAMESPACE=$OPTARG ;;
	d) LOCATION=$OPTARG ;;
    esac
done

echo "NS=$NAMESPACE"
echo "LOC=$LOCATION"

if test X"$NAMESPACE" != X; then
    # create a namespace 
    kubectl create ns ${NAMESPACE}
    NAMESPACE_STR="--namespace=${NAMESPACE}"
fi

# Install crds 
for f in ./deploy/crds/*_crd.yaml ; do     
  kubectl apply -f "${f}" ; 
done

# Install CR
for f in ./deploy/crds/*_cr.yaml; do
    kubectl apply -f "${f}" $NAMESPACE_STR
done

# Check if operator-sdk is installed or not and accordinlgy execute the command.
if test X"$LOCATION" = Xlocal; then
    operator-sdk &> /dev/null
    if [ $? == 0 ]; then
	operator-sdk up local $NAMESPACE_STR &
    else
	echo "Operator SDK is not installed."
	exit 1
    fi
else
    #TODO: change the location in the container stanza of the operator yaml
    for f in ./deploy/*.yaml ; do     
	kubectl apply -f "${f}" --validate=false $NAMESPACE_STR 
    done
    
    for f in ./deploy/resources/*.crd.yaml ; do     
	kubectl apply -f "${f}" --validate=false 
    done
    echo "Deployed all the operator yamls for kubefed-operator in the cluster"
fi



