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
}

func GetAccessToken(clientID string, clientSecret string) clientCredentialsAPI {
	const authenticationUrl string = "https://oauth.battle.net/token"
	resp, err := http.Post(authenticationUrl, "application/x-www-form-urlencoded", strings.NewReader(fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", clientID, clientSecret)))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	if resp.StatusCode != 200 {
		fmt.Println(resp.Status)
		panic("Failed to get access token")
	}
	var credentials clientCredentialsAPI
	err = json.Unmarshal(respBody, &credentials)
	if err != nil {
		fmt.Println(err)
		panic("Failed to parse access token response")
	}
	fmt.Println("Access Token: ", credentials.AccessToken)
	return credentials
}

func BattleNetAPI(params BattleNetAPIParams) []byte {
	var requestURL string
	if strings.HasPrefix(params.UrlOrEndpoint, "http") {
		requestURL = params.UrlOrEndpoint
	} else {
		var baseURL string = fmt.Sprintf("https://%s.api.blizzard.com", params.Region)
		requestURL = fmt.Sprintf("%s%s?namespace=%s-%s", baseURL, params.UrlOrEndpoint, params.Namespace, params.Region)
	}
	fmt.Println(requestURL)
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", params.Token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	if resp.StatusCode != 200 {
		fmt.Println(resp.Status)
		panic("Failed to process the request")
	}
	return (respBody)
}
