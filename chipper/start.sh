#!/bin/bash

if [[ ! -f ./chipper ]]; then
   if [[ -f ./go.mod ]]; then
     echo "You need to build chipper first. This can be done with the setup.sh script."
   else
     echo "You must be in the chipper directory."
   fi
   exit 0
fi

if [[ ! -f ./source.sh ]]; then
  echo "You need to make a source.sh file. This can be done with the setup.sh script."
  exit 0
fi

source source.sh

./chipper
