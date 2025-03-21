package gamedata

const (
	PvpSeasonIndexEndpoint      string = "/data/wow/pvp-season/index"
	PvpSeasonEndpoint           string = "/data/wow/pvp-season/%d"
	PvpLeaderboardIndexEndpoint string = "/data/wow/pvp-season/%d/pvp-leaderboard/index"
	PvpLeaderboardEndpoint      string = "/data/wow/pvp-season/%d/pvp-leaderboard/%s"
	PvpRewardsIndexEndpoint     string = "/data/wow/pvp-season/%d/pvp-reward/index"
)

const (
	Bracket2v2                string = "2v2"
	Bracket3v3                string = "3v3"
	BracketRBG                string = "rbg"
	ClassDeathKnightBlood     string = "deathknight-blood"
	ClassDeathKnightFrost     string = "deathknight-frost"
	ClassDeathKnightUnholy    string = "deathknight-unholy"
	ClassDemonHunterHavoc     string = "demonhunter-havoc"
	ClassDemonHunterVengeance string = "demonhunter-vengeance"
	ClassDruidBalance         string = "druid-balance"
	ClassDruidFeral           string = "druid-feral"
	ClassDruidGuardian        string = "druid-guardian"
	ClassDruidRestoration     string = "druid-restoration"
	ClassEvokerAugmentation   string = "evoker-augmentation"
	ClassEvokerDevasation     string = "evoker-devastation"
	ClassEvokerPreservation   string = "evoker-preservation"
	ClassHunterBeastMastery   string = "hunter-beastmastery"
	ClassHunterMarksmanship   string = "hunter-marksmanship"
	ClassHunterSurvival       string = "hunter-survival"
	ClassMageArcane           string = "mage-arcane"
	ClassMageFire             string = "mage-fire"
	ClassMageFrost            string = "mage-frost"
	ClassMonkBrewmaster       string = "monk-brewmaster"
	ClassMonkMistweaver       string = "monk-mistweaver"
	ClassMonkWindwalker       string = "monk-windwalker"
	ClassPaladinHoly          string = "paladin-holy"
	ClassPaladinProtection    string = "paladin-protection"
	ClassPaladinRetribution   string = "paladin-retribution"
	ClassPriestDiscipline     string = "priest-discipline"
	ClassPriestHoly           string = "priest-holy"
	ClassPriestShadow         string = "priest-shadow"
	ClassRogueAssassination   string = "rogue-assassination"
	ClassRogueOutlaw          string = "rogue-outlaw"
	ClassRogueSubtlety        string = "rogue-subtlety"
	ClassShamanElemental      string = "shaman-elemental"
	ClassShamanEnhancement    string = "shaman-enhancement"
	ClassShamanRestoration    string = "shaman-restoration"
	ClassWarlockAffliction    string = "warlock-affliction"
	ClassWarlockDemonology    string = "warlock-demonology"
	ClassWarlockDestruction   string = "warlock-destruction"
	ClassWarriorArms          string = "warrior-arms"
	ClassWarriorFury          string = "warrior-fury"
	ClassWarriorProtection    string = "warrior-protection"
)

type PvpSeasonIndexAPI struct {
	Seasons       []idAndKey `json:"seasons"`
	CurrentSeason idAndKey   `json:"current_season"`
}

type PvpSeasonAPI struct {
	ID                   int  `json:"id"`
	Leaderboards         href `json:"leaderboards"`
	Rewards              href `json:"rewards"`
	SeasonStartTimestamp int  `json:"season_start_timestamp"`
	SeasonEndTimestamp   int  `json:"season_end_timestamp"`
}

type PvpLeaderboardAPI struct {
	Season  idAndKey  `json:"season"`
	Name    string    `json:"name"`
	Bracket idAndType `json:"bracket"`
	Entries []struct {
		Character struct {
			Name  string `json:"name"`
			ID    int    `json:"id"`
			Realm struct {
				Key  href   `json:"key"`
				ID   int    `json:"id"`
				Slug string `json:"slug"`
			} `json:"realm"`
		} `json:"character"`
		Faction struct {
			Type string `json:"type"`
		} `json:"faction"`
		Rank                  int `json:"rank"`
		Rating                int `json:"rating"`
		SeasonMatchStatistics struct {
			Played int `json:"played"`
			Won    int `json:"won"`
			Lost   int `json:"lost"`
		} `json:"season_match_statistics"`
		Tier idAndKey `json:"tier"`
	} `json:"entries"`
}

type PvpRewardsIndexAPI struct {
	Season  idAndKey `json:"season"`
	Rewards struct {
		Bracket     idAndType `json:"bracket"`
		Achievement struct {
			ID   int    `json:"id"`
			Key  href   `json:"key"`
			Name string `json:"name"`
		} `json:"achievement"`
		RatingCutOff int `json:"rating_cutoff"`
		Faction      struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"faction"`
	} `json:"rewards"`
}
