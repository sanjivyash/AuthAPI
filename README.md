# Authenticated API for Blockchain Applications

Download and install the project with
```
go get github.com/sanjivyash/AuthAPI
```
You need to set up a storage pipeline before running the application. The dummy storage that the application uses is set up as follows after navigating to the main project directory:
```
mkdir storage
touch storage/users.json storage/tokens.json
```
After this, the server can be run simply with 
```
go run main.go
```
All dependencies are mentioned in the ```go.mod``` file. To test the server, run 
```
bash test.sh
```
from the project root directory, which will display the appropriate message.

## Environment Variables
The development environment variables are stored in ```config/dev.env```. The default file is as follows:
```
PORT=:3000
ENDPOINT=api
BASE_DIR=$GOPATH/src/github.com/sanjivyash/AuthAPI
TOKEN_LENGTH=16
TIME_LIMIT=3600
```
Please enter your ```GOPATH``` environment variable in place of the ```$GOPATH``` variable. The file present here has my ```GOPATH``` variable (```/home/sanjivyash/go```) set as default. If your Golang settings are different from the standard one, just set the ```BASE_DIR``` variable to the absolute path of project directory.

## Authentication mechanism
- A new user is allowed to sign up with a unique username and password, stored in the ```storage/users.json``` file. 
- Upon logging in, a token with a finite lifetime is issued controlled by the following environment varaibles:
    - ```TOKEN_LENGTH``` controls the length of the issued token.
    - ```TIME_LIMIT``` controls the  number of seconds after which the token expires. 
- To request any information, the token must be sent as a query parameter.

User can log in again upon expiry of the token. Sessions or JWTs have not been implemented because we do not use a web browser to make requests in blockchain networks. 

## Endpoints for the application
The ```config/dev.env``` file contains various environment variables set for the application, one of which is ```ENDPOINT```, the level at which the API operates. All requests are made to ```/$ENDPOINT/*```. This API supports the following requests:
```
POST /$ENDPOINT/signup              # new user signup
POST /$ENDPOINT/login               # new user login
GET  /$ENDPOINT/info?token=$token   # token authentication is done
```
You can either use ```curl``` or an application like ```POSTMAN``` to make these requests. All data received and transmitted is in ```JSON``` format. The signup and login requests require a file like the one shown below:
```javascript
{
    "username": "admin",
    "password": "admin"
}
```
Errors and messages are sent by the API in ```JSON``` format as well. 

