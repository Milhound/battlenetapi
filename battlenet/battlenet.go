package battlenet

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	STATIC  string = "static"
	DYNAMIC string = "dynamic"
	PROFILE string = "profile"
)

const (
	US string = "us"
	EU string = "eu"
	KR string = "kr"
	TW string = "tw"
)

type clientCredentialsAPI struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Sub         string `json:"sub"`
}

type BattleNetAPIParams struct {
	UrlOrEndpoint string
	Namespace     string
	Region        string
	Token         string
	Options       interface{}
}

type URLFormatter interface {
	FormatURL(baseURL, endpoint, namespace, region string, options interface{}) string
}

func GetAccessToken(clientID string, clientSecret string) clientCredentialsAPI {
	const authenticationUrl string = "https://oauth.battle.net/token"
	resp, err := http.Post(authenticationUrl, "application/x-www-form-urlencoded", strings.NewReader(fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", clientID, clientSecret)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return clientCredentialsAPI{}
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return clientCredentialsAPI{}
	}
	if resp.StatusCode != 200 {
		fmt.Println("Error response from server:", resp.Status)
		return clientCredentialsAPI{}
	}
	var credentials clientCredentialsAPI
	err = json.Unmarshal(respBody, &credentials)
	if err != nil {
		fmt.Println("Error parsing response body:", err)
		return clientCredentialsAPI{}
	}
	fmt.Println("Access Token:", credentials.AccessToken)
	return credentials
}

func BattleNetAPI(params BattleNetAPIParams, formatter URLFormatter) []byte {
	if params.UrlOrEndpoint == "" || params.Namespace == "" || params.Region == "" || params.Token == "" {
		fmt.Println("Invalid parameters")
		return nil
	}

	var requestURL string
	if strings.HasPrefix(params.UrlOrEndpoint, "http") {
		requestURL = params.UrlOrEndpoint
	} else {
		baseURL := fmt.Sprintf("https://%s.api.blizzard.com", params.Region)
		if formatter != nil {
			requestURL = formatter.FormatURL(baseURL, params.UrlOrEndpoint, params.Namespace, params.Region, params.Options)
		} else {
			requestURL = fmt.Sprintf("%s%s?namespace=%s-%s", baseURL, params.UrlOrEndpoint, params.Namespace, params.Region)
		}
	}

	fmt.Println("Request URL:", requestURL)
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", params.Token))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}
	if resp.StatusCode != 200 {
		fmt.Printf("Error response from server: %s\n", resp.Status)
		return nil
	}

	return respBody
}
