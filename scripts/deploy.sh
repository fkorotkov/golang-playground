#!/usr/bin/env bash

if [[ "$CIRRUS_RELEASE" == "" ]]; then
  exit 0
fi

file_content_type="application/octet-stream"
files_to_upload=(
  "golang-playground"
)

for fpath in $files_to_upload
do
  echo "Uploading $fpath"
  name=$(basename "$fpath")
  url_to_upload="https://uploads.github.com/repos/$CIRRUS_REPO_FULL_NAME/releases/$CIRRUS_RELEASE/assets?name=$name"
  curl -X POST \
    --data-binary @$fpath \
    --header "Authorization: Bearer $CIRRUS_GITHUB_TOKEN" \
    --header "Content-Type: $file_content_type" \
    $url_to_upload
done