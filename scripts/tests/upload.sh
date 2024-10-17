#!/bin/sh

token="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImppdGhlbmRyYW5hZGgxNUBnbWFpbC5jb20iLCJleHAiOjE4MjA0NDg5OTUsInJhbmsiOjMsInRlYW1pZCI6LTEsInVzZXJpZCI6MX0.oH3c3WZy7Hw_nw-fGUOhxvp-wNSH_YX1DbBexp8OOmM"
port="9049"

curl http://127.0.0.1:${port}/admin/upload \
    -X POST \
    --silent \
    -H "Authorization: Bearer $token" \
    -H "Content-Type: multipart/form-data" \
    -F "file=@files/welcome.txt" 
