#!/usr/bin/env bash

folder="$1"

if [[ "${folder}" == "" ]]; then echo "Missing folder"; exit 1;fi
folder="$(realpath "$folder")"
if [[ ! -d "${folder}" ]]; then echo "Not found ${folder}"; exit 1;fi

find ${folder} -type f -name "*.txt" | \
  while read -r file; do \
    name="$(basename "${file}")"
    ./search -action add -title "${name/.*}" -text "$(cat "${file}")" -tags "essai,demo"
  done
