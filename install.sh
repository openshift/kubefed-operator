#!/bin/bash
# This script will create a namespace and deploy all the crds within the same
# namespace
# usage ./install.sh <Namespace>

set -e

if [ "$#" -eq 0 ]; then
    NAMESPACE="default"
    LOCATION="local"
elif [ "$#" -eq 1 ]; then
    NAMESPACE=$1
    LOCATION="local"
else
    NAMESPACE=$1
    LOCATION=$2
fi

# create a namespace 
kubectl create ns ${NAMESPACE}

# Install crds 
for f in ./deploy/crds/*_crd.yaml ; do     
  
  kubectl apply -f "${f}" ; 

done

# Check if operator-sdk is installed or not and accordinlgy execute the command.
if [ "$LOCATION" = "local" ]; then
    operator-sdk &> /dev/null
    if [ $? == 0 ]; then
	operator-sdk up local  --namespace=${NAMESPACE}
    else
	echo "Operator SDK is not installed."
	exit 1
    fi
else
    #TODO: change the location in the container stanza of the operator yaml
    for f in ./deploy/*.yaml ; do     
	kubectl apply -f "${f}" --validate=false --namespace=${NAMESPACE} 
    done
    for f in ./deploy/resources/*.crd.yaml ; do     
	kubectl apply -f "${f}" --validate=false 
    done
    echo "Deployed all the operator yamls for kubefed-operator in the cluster"
fi

