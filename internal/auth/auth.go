package auth

import (
	"errors"
	"net/http"
	"strings"
)

//this method exctracts the apikey from the headers
//eg : Authorixation: ApiKey "api_key"
func GetApiKey(headers http.Header)(string, error){
	val := headers.Get("Authorization")
	if val == ""{
		return "", errors.New("no Authentication Info Found")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 {
		return "", errors.New("malformed authorization header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed auth header for first part")
	}

	return vals[1], nil;
}	