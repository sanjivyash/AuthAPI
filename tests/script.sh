#!usr/bin/bash
url=http://localhost:3000/api

# delete user check
curl -s -H "Content-Type: application/json" -X POST -d '{"username":"sanjivyash","password":"authapi"}' $url/delete
echo

# create user check
curl -s -H "Content-Type: application/json" -X POST -d '{"username":"sanjivyash","password":"authapi"}' $url/signup
echo

# valid login check
token=$(curl -s -H "Content-Type: application/json" -X POST -d '{"username":"sanjivyash","password":"authapi"}' $url/login | python3 -c 'import json,sys;output=json.load(sys.stdin);print(output["token"])')

# correct token check
echo $(curl -s -X GET $url/info?token=$token | python3 -c 'import json,sys;output=json.load(sys.stdin);print(output["message"])')

# invalid login check
error=$(curl -s -H "Content-Type: application/json" -X POST -d '{"username":"wronglol","password":"check"}' $url/login | python3 -c 'import json,sys;output=json.load(sys.stdin);print(output["error"])')
echo $error

# incorrect token check
echo $(curl -s -X GET $url/info?token=$error | python3 -c 'import json,sys;output=json.load(sys.stdin);print(output["error"])')