#!/bin/bash

set -eux

export $(cat .env | xargs)

if [ -z ${PROJECT_ID+x} ]; then
  echo "PROJECT_ID is missing"
  exit 1
fi

if [ -z ${COLLECTION_ID+x} ]; then
  echo "COLLECTION_ID is missing"
  exit 1
fi

if [ -z ${DOC_ID+x} ]; then
  echo "DOC_ID is missing"
  exit 1
fi

echo '{"id":1}' | fso create $PROJECT_ID $COLLECTION_ID $DOC_ID

if [ "$(fso ls $PROJECT_ID | grep $COLLECTION_ID)" = "" ]; then
  echo "ERROR: fso ls $PROJECT_ID"
  exit 1
fi
if [ "$(fso ls $PROJECT_ID $COLLECTION_ID | grep $DOC_ID)" = "" ]; then
  echo "ERROR: fso ls $PROJECT_ID $COLLECTION_ID"
  exit 1
fi
if [ "$(fso get $PROJECT_ID $COLLECTION_ID)" != '{"sample":{"id":1}}' ]; then
  echo "ERROR: fso get $PROJECT_ID $COLLECTION_ID"
  exit 1
fi
if [ "$(fso get $PROJECT_ID $COLLECTION_ID $DOC_ID)" != '{"id":1}' ]; then
  echo "ERROR: fso get $PROJECT_ID $COLLECTION_ID $DOC_ID"
  exit 1
fi

echo '{"id":2}' | fso set $PROJECT_ID $COLLECTION_ID $DOC_ID
if [ "$(fso get $PROJECT_ID $COLLECTION_ID $DOC_ID)" != '{"id":2}' ]; then
  echo "ERROR: fso get $PROJECT_ID $COLLECTION_ID $DOC_ID"
  exit 1
fi

fso delete $PROJECT_ID $COLLECTION_ID $DOC_ID
if [ "$(fso get $PROJECT_ID $COLLECTION_ID)" != '{}' ]; then
  echo "ERROR: fso delete $PROJECT_ID $COLLECTION_ID $DOC_ID"
  exit 1
fi 
