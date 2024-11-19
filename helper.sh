#!/bin/bash

# Variables
BASE_URL="https://localhost:2053"
USERNAME="admin"
PASSWORD="admin"
COOKIE_FILE="session.cookie"


if [ "$LOGIN" = true ]; then
    # Step 1: Login and save session cookie
    curl --location --request POST "$BASE_URL/login" \
        --header "Content-Type: application/x-www-form-urlencoded" \
        --data-urlencode "username=$USERNAME" \
        --data-urlencode "password=$PASSWORD" \
        --http1.1 --insecure \
        --cookie-jar $COOKIE_FILE
    echo "Login completed. Session cookie saved in $COOKIE_FILE."
fi


curl --location --request POST "$BASE_URL/panel/inbound/list" \
     --header "Accept: application/json" \
     --http1.1 --insecure \
     --cookie $COOKIE_FILE | jq

curl --location --request POST "$BASE_URL/panel/inbound/onlines" \
     --header "Accept: application/json" \
     --http1.1 --insecure \
     --cookie $COOKIE_FILE | jq
