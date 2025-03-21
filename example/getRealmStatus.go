package example

import (
	"battlenetapi/battlenet"
	"battlenetapi/wow/gamedata"
	"encoding/json"
	"fmt"
	"os"
)

func GetRealmStatus(params battlenet.BattleNetAPIParams, formatter battlenet.URLFormatter) {
	// Get the realm status
	response := battlenet.BattleNetAPI(params, formatter)
	var searchResults gamedata.ConnectedRealmSearchAPI
	var data []gamedata.RealmSearchResult
	err := json.Unmarshal(response, &searchResults)
	if err != nil {
		fmt.Println("Unable to parse ConnectedRealmSearchAPI")
		return
	}

	// Handle more than one page of data
	currentPage := params.Options.(gamedata.RealmStatusParams).Page
	if currentPage > 1 {
		for currentPage <= searchResults.PageCount {
			currentPage += 1

			// Update page to current page
			realmStatusParams := params.Options.(*gamedata.RealmStatusParams)
			realmStatusParams.Page = currentPage
			params.Options = realmStatusParams

			// Call next page
			response := battlenet.BattleNetAPI(params, formatter)
			json.Unmarshal(response, &searchResults)
			data = append(data, searchResults.Results...)
		}
	} else {
		data = searchResults.Results
	}

	file, err := os.Create("realm_status.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Print out select information from the data
	for _, realm := range data {
		info := realm.Data.Realms[0]
		fmt.Printf(
			"%s-%s (%s - %s): %s (%s) \n",
			info.Name.US,
			info.Type.Name.US,
			info.Category.US,
			info.Timezone,
			realm.Data.Status.Name.US,
			realm.Data.Population.Name.US,
		)
	}

	// Save complete data
	searchResults.Results = data
	results, err := json.Marshal(searchResults)
	if err != nil {
		fmt.Println("Unable to convert search results to a seralized object")
	}
	_, err = file.Write(results)
	if err != nil {
		fmt.Println("Error writing data to file:", err)
		return
	}
}
