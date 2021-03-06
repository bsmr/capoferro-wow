package wow

import (
	. "launchpad.net/gocheck"
	"testing"
)

// GoCheck boilerplate
func Test(t *testing.T) { TestingT(t) }

type ApiClientSuite struct{}

var _ = Suite(&ApiClientSuite{})

func (s *ApiClientSuite) Test_signature(c *C) {
	client, _ := NewApiClient("US", "")
	c.Assert(client.signature("GET", "a/b/c"), Not(Equals), "")
}

func (s *ApiClientSuite) Test_NewApiClient_default(c *C) {
	client, _ := NewApiClient("US", "")
	c.Assert(client.Host, Equals, "us.battle.net")
	c.Assert(client.Locale, Equals, "en_US")
}

func (s *ApiClientSuite) Test_NewApiClient_specific(c *C) {
	client, _ := NewApiClient("EU", "fr_FR")
	c.Assert(client.Host, Equals, "eu.battle.net")
	c.Assert(client.Locale, Equals, "fr_FR")
}

func (s *ApiClientSuite) Test_NewApiClient_invalid(c *C) {
	_, err := NewApiClient("China", "it_IT")
	c.Assert(err.Error(), Equals, "Locale 'it_IT' is not valid for region 'China'")
}

func (s *ApiClientSuite) Test_NewApiClient_invalidRegion(c *C) {
	_, err := NewApiClient("Notaregion", "")
	c.Assert(err.Error(), Equals, "Region 'Notaregion' is not valid")
}

func (s *ApiClientSuite) Test_GetAchievement(c *C) {
	client, _ := NewApiClient("US", "")
	a, _ := client.GetAchievement(2144)

	c.Assert(a.AccountWide, Equals, true)
	c.Assert(a.Description, Equals, "Complete the world events achievements listed below.")
	c.Assert(len(a.Criteria), Equals, 9)
	c.Assert(len(a.RewardItems), Equals, 1)
}

func (s *ApiClientSuite) Test_GetAuctionData(c *C) {
	client, _ := NewApiClient("US", "")
	a, _ := client.GetAuctionData("Runetotem")
	c.Assert(len(a.Files), Equals, 1)
}

func (s *ApiClientSuite) Test_GetBattlePetAbility(c *C) {
	client, _ := NewApiClient("US", "")
	a, _ := client.GetBattlePetAbility(640)
	c.Assert(a.Id, Equals, 640)
	c.Assert(a.Cooldown, Equals, 0)
	c.Assert(a.Icon, Equals, "spell_shadow_plaguecloud")
	c.Assert(a.IsPassive, Equals, false)
	c.Assert(a.Name, Equals, "Toxic Smoke")
	c.Assert(a.PetTypeId, Equals, 9)
	c.Assert(a.Rounds, Equals, 1)
	c.Assert(a.ShowHints, Equals, false)
}

func (s *ApiClientSuite) Test_GetBattlePetSpecies(c *C) {
	client, _ := NewApiClient("US", "")
	a, _ := client.GetBattlePetSpecies(258)

	c.Assert(len(a.Abilities), Equals, 6)
	c.Assert(a.CanBattle, Equals, true)
	c.Assert(a.CreatureId, Equals, 42078)
	c.Assert(a.Description, Equals, "Powerful artillery of the terran army. The Thor is always the first one in and the last one out!")
	c.Assert(a.Icon, Equals, "t_roboticon")
	c.Assert(a.PetTypeId, Equals, 9)
	c.Assert(a.Source, Equals, "Promotion: StarCraft II: Wings of Liberty Collector's Edition")
	c.Assert(a.SpeciesId, Equals, 258)
}

func (s *ApiClientSuite) Test_GetBattlePetStats(c *C) {
	client, _ := NewApiClient("US", "")
	a, _ := client.GetBattlePet(258, 25, 5, 4)

	c.Assert(a.BreedId, Equals, 5)
	c.Assert(a.Health, Equals, 1587)
	c.Assert(a.Level, Equals, 25)
	c.Assert(a.PetQualityId, Equals, 4)
	c.Assert(a.Power, Equals, 315)
	c.Assert(a.SpeciesId, Equals, 258)
	c.Assert(a.Speed, Equals, 297)
}

func (s *ApiClientSuite) Test_GetChallenges(c *C) {
	client, _ := NewApiClient("US", "")
	a, _ := client.GetChallenges("Runetotem")
	c.Assert(len(a), Equals, 9)
	gate := a[0]
	expectedRealm := &Realm{"Runetotem", "runetotem", "Vengeance", "en_US", "America/Los_Angeles"}
	c.Assert(gate.Realm.Name, Equals, expectedRealm.Name)
	c.Assert(gate.Realm.Slug, Equals, expectedRealm.Slug)
	c.Assert(gate.Realm.Battlegroup, Equals, expectedRealm.Battlegroup)
	c.Assert(gate.Realm.Locale, Equals, expectedRealm.Locale)
	c.Assert(gate.Realm.Timezone, Equals, expectedRealm.Timezone)
	expectedMap := &Map{962, "Gate of the Setting Sun", "gate-of-the-setting-sun", true, &ChallengeTime{2700000, 0, 45, 0, 0, true}, &ChallengeTime{1320000, 0, 22, 0, 0, true}, &ChallengeTime{780000, 0, 13, 0, 0, true}}

	c.Assert(gate.Map.Id, Equals, expectedMap.Id)
	c.Assert(gate.Map.Name, Equals, expectedMap.Name)
	c.Assert(gate.Map.Slug, Equals, expectedMap.Slug)
	c.Assert(gate.Map.HasChallengeMode, Equals, expectedMap.HasChallengeMode)
	c.Assert(gate.Map.Bronze.Time, Equals, expectedMap.Bronze.Time)
	c.Assert(gate.Map.Bronze.Hours, Equals, expectedMap.Bronze.Hours)
	c.Assert(gate.Map.Silver.Minutes, Equals, expectedMap.Silver.Minutes)
	c.Assert(gate.Map.Silver.Seconds, Equals, expectedMap.Silver.Seconds)
	c.Assert(gate.Map.Gold.Milliseconds, Equals, expectedMap.Gold.Milliseconds)
	c.Assert(gate.Map.Gold.IsPositive, Equals, expectedMap.Gold.IsPositive)

	c.Assert(len(gate.Groups), Equals, 100)
}

func (s *ApiClientSuite) Test_GetCharacter(c *C) {
	client, _ := NewApiClient("US", "")
	a, _ := client.GetCharacter("Runetotem", "Capoferro")
	c.Assert(a.Race, Equals, 2)
	c.Assert(a.ClassId, Equals, 6)
	c.Assert(a.Gender, Equals, 0)
}

func (s *ApiClientSuite) Test_GetCharacterWithFields(c *C) {
	client, _ := NewApiClient("US", "")
	a, _ := client.GetCharacterWithFields(
		"Runetotem",
		"Capoferro",
		[]string{
			"achievements",
			"appearance",
			"feed",
			"guild",
			"hunterPets",
			"items",
			"mounts",
			"pets",
			"petSlots",
			"professions",
			"progression",
			"pvp",
			"quests",
			"reputation",
			"stats",
			"talents",
			"titles",
		})
	c.Assert(a.ClassId, Equals, 6)
	c.Assert(a.Race, Equals, 2)
	c.Assert(a.Gender, Equals, 0)
	c.Assert(len(a.Items.Head.Stats), Equals, 5)
	c.Assert(a.ApiClient, Equals, client)
}

func (s *ApiClientSuite) Test_GetItem(c *C) {
	client, _ := NewApiClient("US", "")

	a, _ := client.GetItem(18803)

	c.Assert(a.Armor, Equals, 0)
	c.Assert(a.BaseArmor, Equals, 0)
	c.Assert(len(a.Stats), Equals, 4)
}

func (s *ApiClientSuite) Test_GetItem2(c *C) {
	client, _ := NewApiClient("US", "")

	a, err := client.GetItem(104426)
	if err != nil {
		println(err.Error())
	}
	c.Assert(a.Armor, Equals, 0)
	c.Assert(a.BaseArmor, Equals, 0)
	c.Assert(len(a.Stats), Equals, 0)
}

func (s *ApiClientSuite) Test_GetItemSet(c *C) {
	client, _ := NewApiClient("US", "")

	a, _ := client.GetItemSet(1060)

	c.Assert(len(a.Items), Equals, 5)
	c.Assert(len(a.SetBonuses), Equals, 2)
}

func (s *ApiClientSuite) Test_GetGuild(c *C) {
	client, _ := NewApiClient("US", "")

	a, _ := client.GetGuildWithFields("Runetotem", "Reforged", []string{"achievements", "news", "challenge", "members"})
	c.Assert(a.Name, Equals, "Reforged")
	c.Assert(a.Level, Equals, 25)
	c.Assert(len(a.News) > 0, Equals, true)
	c.Assert(len(a.Members) > 0, Equals, true)
	c.Assert(len(a.Challenge) > 0, Equals, true)
	c.Assert(a.Achievements != nil, Equals, true)
}

func (s *ApiClientSuite) Test_GetPvPLeaderboard(c *C) {
	client, _ := NewApiClient("US", "")

	a, err := client.GetPvPLeaderboard("3v3")
	if err != nil {
		println(err.Error())
	}
	c.Assert(len(a) > 0, Equals, true)
}

func (s *ApiClientSuite) Test_GetQuest(c *C) {
	client, _ := NewApiClient("US", "")

	a, err := client.GetQuest(13146)
	if err != nil {
		println(err.Error())
	}
	c.Assert(a.Level, Equals, 80)
	c.Assert(a.SuggestedPartyMembers, Equals, 0)
}

func (s *ApiClientSuite) Test_GetRealmStatus(c *C) {
	client, _ := NewApiClient("US", "")

	a, err := client.GetRealmStatus()
	if err != nil {
		println(err.Error())
	}
	c.Assert(len(a) > 0, Equals, true)
}

func (s *ApiClientSuite) Test_GetRecipe(c *C) {
	client, _ := NewApiClient("US", "")

	a, err := client.GetRecipe(33994)
	if err != nil {
		println(err.Error())
	}
	c.Assert(a.Name, Equals, "Enchant Gloves - Precise Strikes")
	c.Assert(a.Profession, Equals, "Enchanting")
}

func (s *ApiClientSuite) Test_GetSpell(c *C) {
	client, _ := NewApiClient("US", "")

	a, err := client.GetSpell(8056)
	if err != nil {
		println(err.Error())
	}
	c.Assert(a.CastTime, Equals, "Instant")
	c.Assert(a.Cooldown, Equals, "6 sec cooldown")
}

func (s *ApiClientSuite) Test_GetBattlegroups(c *C) {
	client, _ := NewApiClient("US", "")

	a, err := client.GetBattlegroups()
	if err != nil {
		println(err.Error())
	}
	c.Assert(len(a) > 0, Equals, true)
}

func (s *ApiClientSuite) Test_GetRaces(c *C) {
	client, _ := NewApiClient("US", "")

	a, err := client.GetRaces()
	if err != nil {
		println(err.Error())
	}
	c.Assert(len(a) > 0, Equals, true)
}

func (s *ApiClientSuite) Test_GetClasss(c *C) {
	client, _ := NewApiClient("US", "")

	a, err := client.GetClasses()
	if err != nil {
		println(err.Error())
	}
	c.Assert(len(a) > 0, Equals, true)
}

func (s *ApiClientSuite) Test_GetAchievements(c *C) {
	client, _ := NewApiClient("US", "")

	a, err := client.GetAchievements()
	if err != nil {
		println(err.Error())
	}
	c.Assert(len(a) > 0, Equals, true)
}

func (s *ApiClientSuite) Test_GetGuildRewards(c *C) {
	client, _ := NewApiClient("US", "")

	a, err := client.GetGuildRewards()
	if err != nil {
		println(err.Error())
	}
	c.Assert(len(a) > 0, Equals, true)
}

func (s *ApiClientSuite) Test_GetGuildPerks(c *C) {
	client, _ := NewApiClient("US", "")

	a, err := client.GetGuildPerks()
	if err != nil {
		println(err.Error())
	}
	c.Assert(len(a) > 0, Equals, true)
}

func (s *ApiClientSuite) Test_GetGuildAchievements(c *C) {
	client, _ := NewApiClient("US", "")

	a, err := client.GetGuildAchievements()
	if err != nil {
		println(err.Error())
	}
	c.Assert(len(a) > 0, Equals, true)
}

func (s *ApiClientSuite) Test_GetItemClasses(c *C) {
	client, _ := NewApiClient("US", "")

	a, err := client.GetItemClasses()
	if err != nil {
		println(err.Error())
	}
	c.Assert(len(a) > 0, Equals, true)
	c.Assert(len(a[0].Subclasses) > 0, Equals, true)
}

func (s *ApiClientSuite) Test_GetTalents(c *C) {
	client, _ := NewApiClient("US", "")

	a, err := client.GetTalents()
	if err != nil {
		println(err.Error())
	}

	c.Assert(a.Warrior, Not(IsNil))
	c.Assert(len(a.Warrior.Glyphs) > 0, Equals, true)
	c.Assert(a.Warrior.Talents[0][0], Not(IsNil))
}

func (s *ApiClientSuite) Test_GetPetTypes(c *C) {
	client, _ := NewApiClient("US", "")

	a, err := client.GetPetTypes()
	if err != nil {
		println(err.Error())
	}
	c.Assert(len(a) > 0, Equals, true)
}

