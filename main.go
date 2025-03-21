package main

import (
	"battlenetapi/battlenet"
	"battlenetapi/example"
	"battlenetapi/wow/gamedata"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Authenticate with the BattleNet API
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	credentials := battlenet.GetAccessToken(clientID, clientSecret)

	// Get the leaderboard for the shuffle bracket
	params := battlenet.BattleNetAPIParams{
		UrlOrEndpoint: gamedata.PvpSeasonIndexEndpoint,
		Namespace:     battlenet.DYNAMIC,
		Region:        battlenet.US,
		Token:         credentials.AccessToken,
	}
	shuffleOrBlitz := fmt.Sprintf("blitz-%s", gamedata.ClassDemonHunterHavoc)
	example.GetLeaderboard(params, shuffleOrBlitz)

	// Get current realm status
	params = battlenet.BattleNetAPIParams{
		UrlOrEndpoint: gamedata.ConnectedRealmSearchEndpoint,
		Namespace:     battlenet.DYNAMIC,
		Region:        battlenet.US,
		Token:         credentials.AccessToken,
		Options: gamedata.RealmStatusParams{
			Status:  gamedata.UP,
			OrderBy: "id",
			Page:    1,
		},
	}
	formatter := gamedata.URLFormatterImpl{}
	example.GetRealmStatus(params, formatter)
}
