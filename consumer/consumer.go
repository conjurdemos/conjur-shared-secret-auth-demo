package main

import (
	"fmt"
	"os"
	"github.com/cyberark/conjur-api-go/conjurapi"
	"github.com/cyberark/conjur-api-go/conjurapi/authn"
)

func main() {
	secretId := "helloworld-secret"

	fmt.Printf("Connecting to Conjur\n");

	// Receive the consumer conjur login and key
	config := conjurapi.LoadConfig()
	conjur, err := conjurapi.NewClientFromKey(config,
		authn.LoginPair{
			Login: os.Getenv("CONJUR_AUTHN_LOGIN"),
			APIKey: os.Getenv("CONJUR_AUTHN_API_KEY"),
		},
	)

	if err != nil {
	 	panic(err)
	}

	// User Conjur login/key to fetch app secret
	secretValue, err := conjur.RetrieveSecret(secretId)
	if err != nil {
		panic(err)
	}
	fmt.Println("The secret value is: ", string(secretValue))
	
	// Issue the GET request to helloworld
	// Print out the result
}