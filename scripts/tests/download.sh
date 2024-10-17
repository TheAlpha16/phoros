#!/bin/sh

token="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImppdGhlbmRyYW5hZGgxNUBnbWFpbC5jb20iLCJleHAiOjE4MjA0NDg5OTUsInJhbmsiOjMsInRlYW1pZCI6LTEsInVzZXJpZCI6MX0.oH3c3WZy7Hw_nw-fGUOhxvp-wNSH_YX1DbBexp8OOmM"
port="9049"

rm welcome.txt
echo "[*] Sending request"

curl http://127.0.0.1:${port}/files/welcome.txt \
    -H "cookie: token=$token;" \
    --silent \
    --output welcome.txt

if ! cmp -s welcome.txt files/welcome.txt 
then
    echo "[-] Failed"
    cat welcome.txt
    exit 1
else
    echo "[+] Passed"
fi
