package main

import (
	"encoding/json"
	"fmt"
	"os"

	"battlenetapi/battlenet"
	"battlenetapi/wow/gamedata"

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

	// Get the current PVP season
	params := battlenet.BattleNetAPIParams{
		UrlOrEndpoint: gamedata.PvpSeasonIndexEndpoint,
		Namespace:     battlenet.DYNAMIC,
		Region:        battlenet.US,
		Token:         credentials.AccessToken,
	}
	response := battlenet.BattleNetAPI(params)
	var pvpIndex gamedata.PvpSeasonIndexAPI
	json.Unmarshal(response, &pvpIndex)

	// Get the leaderboards for the current PVP season and Bracket
	shuffleOrBlitz := fmt.Sprintf("blitz-%s", gamedata.ClassDemonHunterHavoc)
	params.UrlOrEndpoint = fmt.Sprintf(gamedata.PvpLeaderboardEndpoint, pvpIndex.CurrentSeason.ID, shuffleOrBlitz)
	response = battlenet.BattleNetAPI(params)
	file, err := os.Create(fmt.Sprintf("pvp_season_%d_leaderboard-bracket_%s.json", pvpIndex.CurrentSeason.ID, shuffleOrBlitz))
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	_, err = file.Write(response)
	if err != nil {
		fmt.Println("Error writing data to file:", err)
		return
	}
}
