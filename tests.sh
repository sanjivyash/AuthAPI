#!usr/bin/bash
url=http://localhost:3000/api

curl -H "Content-Type: application/json" -X POST -d '{"username":"sanjivyash","password":"authapi"}' $url/delete
echo 
sleep 1

curl -H "Content-Type: application/json" -X POST -d '{"username":"sanjivyash","password":"authapi"}' $url/signup
echo
sleep 1

curl -s -H "Content-Type: application/json" -X POST -d '{"username":"sanjivyash","password":"authapi"}' $url/login
echo
sleep 1

curl -X GET $url/info?token=c%%?�s�ש���!\u0011�\u0016
echo
sleep 1