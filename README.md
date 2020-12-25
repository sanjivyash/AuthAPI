# Authenticated API for Blockchain Applications

The server can be run simply with 
```
go run main.go
```
All dependencies are mentioned in the ```go.mod``` file.

## Environment Variables
The development environment variables are stored in ```config/dev.env```. The default file is as follows:
```
PORT=:3000
ENDPOINT=api
BASE_DIR=$GOPATH/src/github.com/sanjivyash/AuthAPI
TOKEN_LENGTH=16
TIME_LIMIT=3600
```
Please enter your GOPATH environment variable in place of the ```$GOPATH``` variable. The file present here has my GOPATH variable (```/home/sanjivyash/go```) set as default. If your Golang settings are different from the standard one, just set the ```BASE_DIR``` variable to your project directory absolute path.

## Authentication mechanism
A new user is allowed to sign up with a unique username and password. Upon logging in, a token of fixed length, controlled by the environment variable ```TOKEN_LENGTH```, is issued with a certain time limit is issued. This time limit is controlled by the environment variable ```TIME_LIMIT```, set to this time in seconds. To request any information, the token must be sent as a query parameter. User can log in again upon expiry of the token. Sessions or JWTs have not been implemented because we do not use a web browser to make requests in blockchain networks. 

## Endpoints for the application
The ```config/dev.env``` file contains various environment variables set for the application, one of which is ```ENDPOINT```, the level at which the API operates. All requests are made to ```/$ENDPOINT/*```. This API supports the following requests:
```
POST /$ENDPOINT/signup              # new user signup
POST /$ENDPOINT/login               # new user login
GET  /$ENDPOINT/info?token=$token   # token authentication is done
```
You can either use ```curl``` or an application like ```POSTMAN``` to make these requests. All data received and transmitted is in ```JSON``` format. The signup and login requests require a file like the one shown below:
```json
{
    "username": "admin",
    "password": "admin"
}
```
Errors and messages are sent by the API in ```JSON``` format as well. 

