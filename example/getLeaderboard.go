package example

import (
	"encoding/json"
	"fmt"
	"os"

	"battlenetapi/battlenet"
	"battlenetapi/wow/gamedata"
)

func GetLeaderboard(params battlenet.BattleNetAPIParams, bracket string) {
	// Get the current PVP season
	response := battlenet.BattleNetAPI(params, nil)
	var pvpIndex gamedata.PvpSeasonIndexAPI
	json.Unmarshal(response, &pvpIndex)

	// Get the leaderboards for the current PVP season and Bracket

	params.UrlOrEndpoint = fmt.Sprintf(gamedata.PvpLeaderboardEndpoint, pvpIndex.CurrentSeason.ID, bracket)
	response = battlenet.BattleNetAPI(params, nil)
	file, err := os.Create(fmt.Sprintf("pvp_season_%d_leaderboard-bracket_%s.json", pvpIndex.CurrentSeason.ID, bracket))
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
