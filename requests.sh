#!/bin/bash

reset() {
  curl -X POST \
    "$URL/admin/reset"
}

validateChirp() {
  message=$1

  curl -X POST \
    "$URL/api/chirps" \
    --header 'Content-Type: application/json' \
    -d '
  {
    "body": "'"$message"'"
  }'
}

insertUser() {
  user=$1

  curl -X POST \
    "$URL/api/users" \
    --header 'Content-Type: application/json' \
    -d '
  {
    "email": "'"$user"'"
  }'
}

getChirps() {
  curl -X GET \
    "$URL/api/chirps"
}

# variables

URL="http://localhost:8080"

# examples

reset
insertUser 'cool2@world.com' | jq
validateChirp 'cool' | jq
validateChirp 'lorem ipsum cool too cool for school and thats all that there was and so it was told and all that blablabla hahahhdfanjbngfjnadfpnasf;lna;sdlngf'
validateChirp 'this is a kerfuffle situation' | jq
getChirps
