package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/cyberark/conjur-api-go/conjurapi"
	"github.com/cyberark/conjur-api-go/conjurapi/authn"
)

func main() {
	secretID := os.Getenv("DEMO_SERVICE_KEY_ID")
	serviceURL := os.Getenv("DEMO_SERVICE_URL")
	conjurLogin := os.Getenv("CONJUR_AUTHN_LOGIN")
	conjurAPIKey := os.Getenv("CONJUR_AUTHN_API_KEY")

	useUnauthorized := flag.Bool("fail", false, "Sends an invalid token to helloworld")
	flag.Parse()

	// Determine which token value to use
	var serviceToken string
	if *useUnauthorized {
		serviceToken = "Invalid Token"
	} else {
		var err error
		serviceToken, err = fetchSecretFromConjur(secretID, conjurLogin, conjurAPIKey)

		if err != nil {
			panic(err)
		}
	}

	// Send the request to the remote service
	response, err := sendRequestToService(serviceURL, serviceToken)

	if err != nil {
		panic(err)
	}

	fmt.Println("Response from service: ", response)
}

func fetchSecretFromConjur(secretID string, conjurLogin string, conjurAPIKey string) (string, error) {
	fmt.Fprintf(os.Stderr, "Connecting to Conjur...\n")

	// Receive the consumer conjur login and key
	config := conjurapi.LoadConfig()
	conjur, err := conjurapi.NewClientFromKey(config,
		authn.LoginPair{
			Login:  conjurLogin,
			APIKey: conjurAPIKey,
		},
	)

	if err != nil {
		return "", err
	}

	// User Conjur login/key to fetch app secret
	secretValue, err := conjur.RetrieveSecret(secretID)
	if err != nil {
		return "", err
	}

	return string(secretValue), nil
}

func sendRequestToService(serviceURL string, serviceToken string) (string, error) {

	fmt.Fprintf(os.Stderr, "Connecting to %s with token: %s\n", serviceURL, serviceToken)

	client := &http.Client{}

	req, err := http.NewRequest("GET", serviceURL, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set(
		"Authorization",
		fmt.Sprintf("Token token=\"%s\"", serviceToken),
	)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}
