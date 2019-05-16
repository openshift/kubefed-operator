#!/bin/bash
# This script will create a namespace and deploy all the crds within the same
# namespace
# usage ./install.sh <Namespace>

set -e

NAMESPACE=$1

if [ -z "$1" ]; then
    echo "Enter the namespace after ./install.sh"
    exit 1
fi

# create a namespace 
kubectl create ns ${NAMESPACE}

# Install crds 
for f in ./deploy/crds/*_crd.yaml ; do     
  
  kubectl apply -f "${f}" ; 

done

for f in ./deploy/resources/*.yaml ; do     

  kubectl apply -f "${f}" --validate=false; 
  
done

# Check if operator-sdk is installed or not and accordinlgy execute the commad.
operator-sdk &> /dev/null
if [ $? == 0 ]; then
   operator-sdk up local  --namespace=${NAMESPACE}
else
   echo "Operator SDK is not installed."
   exit 1
fi
