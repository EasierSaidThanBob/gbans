package logparse_test

import (
	"errors"
	"github.com/leighmacdonald/gbans/pkg/logparse"
	"github.com/leighmacdonald/golib"
	"github.com/leighmacdonald/steamid/v2/steamid"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"os"
	"path"
	"strings"
	"testing"
	"time"
)

func TestMatch_Apply(t *testing.T) {
	// mock?
	testLogger, _ := zap.NewDevelopment()

	p := golib.FindFile(path.Join("test_data", "log_3124689.log"), "gbans")
	if p == "" {
		t.Skipf("Cant find test file: log_3124689.log")
		return
	}
	body, errRead := os.ReadFile(p)
	require.NoError(t, errRead)

	m := logparse.NewMatch(testLogger, 1, "test server")
	rows := strings.Split(string(body), "\n")
	for _, line := range rows {
		if line == "" {
			continue
		}
		result, errResult := logparse.Parse(line)
		require.NoError(t, errResult)
		if err := m.Apply(result); err != nil && !errors.Is(err, logparse.ErrIgnored) {
			t.Errorf("Failed to Apply: %v [%d] %v", err, result.EventType, line)
		}
	}

	match3124689, names := testMatch()
	getName := func(sid64 steamid.SID64) string {
		for name, sid := range names {
			if sid == sid64 {
				return name
			}
		}
		return "???"
	}
	getPS := func(m logparse.Match, sid steamid.SID64) *logparse.MatchPlayerSum {
		ps, err := m.PlayerSums.GetBySteamId(sid)
		if err != nil {
			t.Fatalf("Failed to fetch player sum [%d]: %v", sid, err)
		}
		return ps
	}
	// Player sum values
	for _, ps := range match3124689.PlayerSums {
		require.Equal(t, getPS(match3124689, ps.SteamID).Kills,
			getPS(m, ps.SteamID).Kills, "Kills incorrect %v", getName(ps.SteamID))
	}
	for _, ps := range match3124689.PlayerSums {
		require.Equal(t, getPS(match3124689, ps.SteamID).Deaths,
			getPS(m, ps.SteamID).Deaths, "Deaths incorrect %v", getName(ps.SteamID))
	}
	for _, ps := range match3124689.PlayerSums {
		require.Equal(t, getPS(match3124689, ps.SteamID).Damage,
			getPS(m, ps.SteamID).Damage, "Damage incorrect %v", getName(ps.SteamID))
	}
	for _, ps := range match3124689.PlayerSums {
		require.Equal(t, getPS(match3124689, ps.SteamID).Healing,
			getPS(m, ps.SteamID).Healing, "Healing incorrect %v", getName(ps.SteamID))
	}
	for _, ps := range match3124689.PlayerSums {
		require.Equal(t, getPS(match3124689, ps.SteamID).Dominations,
			getPS(m, ps.SteamID).Dominations, "Dominations incorrect %v", getName(ps.SteamID))
	}

	for _, ps := range match3124689.PlayerSums {
		require.Equal(t, getPS(match3124689, ps.SteamID).Revenges,
			getPS(m, ps.SteamID).Revenges, "Revenges incorrect %v", getName(ps.SteamID))
	}
	// for sid := range match3124689.playerSums {
	//	require.Equal(t, match3124689.playerSums[sid].Classes,
	//		m.playerSums[sid].Classes, "Classes incorrect %v", getName(sid))
	//}

	getMS := func(m logparse.Match, sid steamid.SID64) *logparse.MatchMedicSum {
		ms, err := m.MedicSums.GetBySteamId(sid)
		if err != nil {
			t.Fatalf("Failed to fetch player sum")
		}
		return ms
	}

	// Medic sums
	for _, ms := range match3124689.MedicSums {
		require.Equal(t, getMS(match3124689, ms.SteamID).Drops,
			getMS(m, ms.SteamID).Drops, "Drops incorrect %v", getName(ms.SteamID))
	}
	// for sid := range match3124689.medicSums {
	//	require.Equal(t, match3124689.medicSums[sid].NearFullChargeDeath,
	//		m.medicSums[sid].Drops, "NearFullChargeDeath incorrect %v", getName(sid))
	//}
	//for _, ms := range match3124689.MedicSums {
	//	require.Equal(t, map[logparse.Medigun]int{logparse.HadUber: 6, logparse.Kritzkrieg: 0, logparse.QuickFix: 0, logparse.Vaccinator: 0},
	//		getMS(m, ms.SteamID).Charges, "Charges incorrect %v", getName(ms.SteamID))
	//}

	// for team := range match3124689.teamSums {
	//	require.Equal(t, match3124689.teamSums[team].Kills,
	//		m.teamSums[team].Kills, "[Team] Kills incorrect %v", team)
	//}
}

// https://logs.tf/3124689
//
//nolint:funlen,maintidx
func testMatch() (logparse.Match, map[string]steamid.SID64) {
	match := logparse.Match{
		Title:   "Qixalite Booking: RED vs BLU",
		MapName: "koth_cascade_rc2",
		PlayerSums: []*logparse.MatchPlayerSum{
			{
				SteamID:           76561198164892406,
				Team:              logparse.BLU,
				TimeStart:         &time.Time{},
				TimeEnd:           &time.Time{},
				Dominations:       0,
				Revenges:          0,
				Dominated:         0,
				Shots:             0,
				Hits:              0,
				BuildingDestroyed: 0,
				Kills:             12,
				Assists:           10,
				Deaths:            14,
				Damage:            5078,
				DamageTaken:       5399,
				HealthPacks:       16,
				BackStabs:         0,
				HeadShots:         0,
				AirShots:          0,
				Captures:          nil,
				Classes:           []logparse.PlayerClass{logparse.Scout},
				Healing:           370,
			},
			{
				SteamID:     76561198057150173,
				Team:        logparse.BLU,
				TimeStart:   &time.Time{},
				TimeEnd:     &time.Time{},
				Kills:       9,
				Assists:     1,
				Deaths:      23,
				Revenges:    1,
				Damage:      4022,
				DamageTaken: 4616,
				HealthPacks: 27,
				BackStabs:   0,
				HeadShots:   0,
				AirShots:    1,
				Captures:    nil,
				Classes:     []logparse.PlayerClass{logparse.Soldier},
				Healing:     754,
			},
			{
				SteamID:     76561198126692772,
				Team:        logparse.BLU,
				TimeStart:   &time.Time{},
				TimeEnd:     &time.Time{},
				Kills:       13,
				Assists:     8,
				Deaths:      5,
				Dominations: 1,
				Damage:      3873,
				DamageTaken: 3537,
				HealthPacks: 16,
				BackStabs:   0,
				HeadShots:   0,
				AirShots:    0,
				Captures:    nil,
				Classes:     []logparse.PlayerClass{logparse.Pyro},
				Healing:     481,
			},
			{
				SteamID:     76561198036671190,
				Team:        logparse.BLU,
				TimeStart:   &time.Time{},
				TimeEnd:     &time.Time{},
				Kills:       13,
				Assists:     4,
				Deaths:      16,
				Dominations: 1,
				Damage:      7598,
				DamageTaken: 6936,
				HealthPacks: 10,
				BackStabs:   0,
				HeadShots:   0,
				AirShots:    0,
				Captures:    nil,
				Classes:     []logparse.PlayerClass{logparse.Demo},
				Healing:     436,
			},
			{
				SteamID:     76561198084686835,
				Team:        logparse.BLU,
				TimeStart:   &time.Time{},
				TimeEnd:     &time.Time{},
				Kills:       18,
				Dominations: 1,
				Assists:     1,
				Deaths:      9,
				Damage:      7250,
				DamageTaken: 8918,
				HealthPacks: 24,
				BackStabs:   0,
				HeadShots:   0,
				AirShots:    0,
				Captures:    nil,
				Classes:     []logparse.PlayerClass{logparse.Heavy},
				Healing:     1187,
			},
			{
				SteamID:     76561198061174192,
				Team:        logparse.BLU,
				TimeStart:   &time.Time{},
				TimeEnd:     &time.Time{},
				Kills:       3,
				Assists:     1,
				Deaths:      14,
				Revenges:    1,
				Damage:      3385,
				DamageTaken: 3246,
				HealthPacks: 16,
				BackStabs:   0,
				HeadShots:   8,
				AirShots:    0,
				Captures:    nil,
				Classes:     []logparse.PlayerClass{logparse.Engineer, logparse.Sniper},
				Healing:     686,
			},
			{
				SteamID:     76561198113244106,
				Team:        logparse.BLU,
				TimeStart:   &time.Time{},
				TimeEnd:     &time.Time{},
				Kills:       4,
				Assists:     8,
				Deaths:      9,
				Damage:      1098,
				DamageTaken: 3169,
				HealthPacks: 10,
				BackStabs:   0,
				HeadShots:   0,
				AirShots:    0,
				Captures:    nil,
				Classes:     []logparse.PlayerClass{logparse.Medic},
				Healing:     17708,
			},
			{
				SteamID:     76561198423392803,
				Team:        logparse.BLU,
				TimeStart:   &time.Time{},
				TimeEnd:     &time.Time{},
				Kills:       7,
				Assists:     0,
				Deaths:      19,
				Revenges:    1,
				Damage:      2809,
				DamageTaken: 4023,
				HealthPacks: 5,
				BackStabs:   0,
				HeadShots:   4,
				AirShots:    0,
				Captures:    nil,
				Classes:     []logparse.PlayerClass{logparse.Sniper, logparse.Engineer},
				Healing:     395,
			},
			{
				SteamID:     76561198051884373,
				Team:        logparse.BLU,
				TimeStart:   &time.Time{},
				TimeEnd:     &time.Time{},
				Kills:       13,
				Assists:     4,
				Deaths:      16,
				Revenges:    1,
				Damage:      9672,
				DamageTaken: 3450,
				HealthPacks: 23,
				BackStabs:   9,
				HeadShots:   0,
				AirShots:    0,
				Captures:    nil,
				Classes:     []logparse.PlayerClass{logparse.Spy, logparse.Pyro},
				Healing:     456,
			},
			{
				SteamID:     76561198087442614,
				Team:        logparse.BLU,
				TimeStart:   &time.Time{},
				TimeEnd:     &time.Time{},
				Kills:       12,
				Assists:     1,
				Deaths:      17,
				Revenges:    1,
				Damage:      8796,
				DamageTaken: 3517,
				HealthPacks: 22,
				BackStabs:   9,
				HeadShots:   0,
				AirShots:    0,
				Captures:    nil,
				Classes:     []logparse.PlayerClass{logparse.Spy},
				Healing:     546,
			},
			{
				SteamID:     76561198043171944,
				Team:        logparse.RED,
				TimeStart:   &time.Time{},
				TimeEnd:     &time.Time{},
				Kills:       28,
				Assists:     5,
				Deaths:      11,
				Dominations: 4,
				Damage:      6236,
				DamageTaken: 2932,
				HealthPacks: 46,
				BackStabs:   9,
				HeadShots:   0,
				AirShots:    0,
				Captures:    nil,
				Classes:     []logparse.PlayerClass{logparse.Scout},
				Healing:     1036,
			},
			{
				SteamID:     76561198809011070,
				Team:        logparse.RED,
				TimeStart:   &time.Time{},
				TimeEnd:     &time.Time{},
				Kills:       26,
				Assists:     2,
				Deaths:      9,
				Dominations: 4,
				Damage:      6107,
				DamageTaken: 4053,
				HealthPacks: 29,
				BackStabs:   0,
				HeadShots:   0,
				AirShots:    1,
				Captures:    nil,
				Classes:     []logparse.PlayerClass{logparse.Soldier},
				Healing:     1313,
			},
			{
				SteamID:     76561198271399587,
				Team:        logparse.RED,
				TimeStart:   &time.Time{},
				TimeEnd:     &time.Time{},
				Kills:       9,
				Assists:     13,
				Deaths:      8,
				Dominations: 1,
				Damage:      4613,
				DamageTaken: 5054,
				HealthPacks: 16,
				BackStabs:   0,
				HeadShots:   0,
				AirShots:    0,
				Captures:    nil,
				Classes:     []logparse.PlayerClass{logparse.Pyro},
				Healing:     553,
			},
			{
				SteamID:     76561198096251579,
				Team:        logparse.RED,
				TimeStart:   &time.Time{},
				TimeEnd:     &time.Time{},
				Kills:       7,
				Assists:     6,
				Deaths:      15,
				Damage:      3859,
				DamageTaken: 6871,
				HealthPacks: 0,
				BackStabs:   0,
				HeadShots:   0,
				AirShots:    0,
				Captures:    nil,
				Classes:     []logparse.PlayerClass{logparse.Demo},
				Healing:     0,
			},
			{
				SteamID:     76561198383642609,
				Team:        logparse.RED,
				TimeStart:   &time.Time{},
				TimeEnd:     &time.Time{},
				Kills:       15,
				Assists:     6,
				Deaths:      11,
				Revenges:    1,
				Damage:      6708,
				DamageTaken: 8299,
				HealthPacks: 20,
				BackStabs:   0,
				HeadShots:   0,
				AirShots:    0,
				Captures:    nil,
				Classes:     []logparse.PlayerClass{logparse.Heavy},
				Healing:     1216,
			},
			{
				SteamID:     76561198050517054,
				Team:        logparse.RED,
				TimeStart:   &time.Time{},
				TimeEnd:     &time.Time{},
				Kills:       18,
				Assists:     9,
				Deaths:      12,
				Dominations: 2,
				Damage:      6309,
				DamageTaken: 4655,
				HealthPacks: 23,
				BackStabs:   0,
				HeadShots:   0,
				AirShots:    0,
				Captures:    nil,
				Classes:     []logparse.PlayerClass{logparse.Engineer, logparse.Scout},
				Healing:     922,
			},
			{
				SteamID:     76561198082713023,
				Team:        logparse.RED,
				TimeStart:   &time.Time{},
				TimeEnd:     &time.Time{},
				Kills:       1,
				Assists:     13,
				Deaths:      5,
				Damage:      606,
				DamageTaken: 2733,
				HealthPacks: 6,
				BackStabs:   0,
				HeadShots:   0,
				AirShots:    0,
				Captures:    nil,
				Classes:     []logparse.PlayerClass{logparse.Medic},
				Healing:     19762,
			},
			{
				SteamID:     76561199050447792,
				Team:        logparse.RED,
				TimeStart:   &time.Time{},
				TimeEnd:     &time.Time{},
				Kills:       18,
				Assists:     1,
				Deaths:      12,
				Damage:      5723,
				DamageTaken: 3933,
				HealthPacks: 4,
				BackStabs:   0,
				HeadShots:   18,
				AirShots:    0,
				Captures:    nil,
				Classes:     []logparse.PlayerClass{logparse.Sniper},
				Healing:     101,
			},
			{
				SteamID:     76561198073709029,
				Team:        logparse.RED,
				TimeStart:   &time.Time{},
				TimeEnd:     &time.Time{},
				Kills:       0,
				Assists:     0,
				Deaths:      1,
				Damage:      40,
				DamageTaken: 372,
				HealthPacks: 2,
				BackStabs:   0,
				HeadShots:   0,
				AirShots:    0,
				Captures:    nil,
				Classes:     []logparse.PlayerClass{logparse.Spy},
				Healing:     57,
			},
		},
		MedicSums: []*logparse.MatchMedicSum{
			{
				SteamID: 76561198113244106,
				Healing: 17368,
				Charges: map[logparse.MedigunType]int{
					logparse.Uber: 4,
				},
				Drops:               2,
				AvgTimeToBuild:      42,
				AvgTimeBeforeUse:    28,
				NearFullChargeDeath: 1,
				AvgUberLength:       7,
				MajorAdvLost:        1,
				BiggestAdvLost:      39,
				DeathAfterCharge:    0,
				HealTargets:         []*logparse.MatchClassSums{},
			},
			{
				SteamID: 76561198082713023,
				Healing: 19545,
				Charges: map[logparse.MedigunType]int{
					logparse.Uber: 6,
				},
				Drops:               1,
				AvgTimeToBuild:      40,
				AvgTimeBeforeUse:    30,
				NearFullChargeDeath: 0,
				AvgUberLength:       7.3,
				MajorAdvLost:        0,
				BiggestAdvLost:      0,
				DeathAfterCharge:    1,
				HealTargets:         []*logparse.MatchClassSums{},
			},
		},
		TeamSums: []*logparse.MatchTeamSum{
			{
				Team:      logparse.RED,
				Kills:     122,
				Damage:    40201,
				Charges:   6,
				Drops:     1,
				Caps:      6,
				MidFights: 2,
			},
			{
				Team:      logparse.BLU,
				Kills:     104,
				Damage:    45512,
				Charges:   4,
				Drops:     2,
				Caps:      4,
				MidFights: 1,
			},
		},
		Rounds: []*logparse.MatchRoundSum{
			{
				Length: 313 * time.Second,
				Score: logparse.TeamScores{
					Red: 1,
					Blu: 0,
				},
				KillsBlu:  23,
				KillsRed:  33,
				UbersBlu:  1,
				UbersRed:  2,
				DamageBlu: 10605,
				DamageRed: 12805,
				MidFight:  logparse.RED,
			},
			{
				Length: 378 * time.Second,
				Score: logparse.TeamScores{
					Red: 2,
					Blu: 0,
				},
				KillsBlu:  38,
				KillsRed:  52,
				UbersBlu:  2,
				UbersRed:  3,
				DamageBlu: 15266,
				DamageRed: 18102,
				MidFight:  logparse.BLU,
			},
			{
				Length: 325 * time.Second,
				Score: logparse.TeamScores{
					Red: 3,
					Blu: 0,
				},
				KillsBlu:  34,
				KillsRed:  46,
				UbersBlu:  1,
				UbersRed:  1,
				DamageBlu: 15258,
				DamageRed: 12677,
				MidFight:  logparse.RED,
			},
		},
		ClassKills: []*logparse.MatchClassSums{
			{
				SteamId:  76561198113244106,
				Scout:    0,
				Soldier:  0,
				Pyro:     0,
				Demoman:  0,
				Heavy:    1,
				Engineer: 0,
				Medic:    0,
				Sniper:   0,
				Spy:      3,
			},
			{
				SteamId:  76561199050447792,
				Scout:    3,
				Soldier:  0,
				Pyro:     2,
				Demoman:  4,
				Heavy:    1,
				Engineer: 3,
				Medic:    2,
				Sniper:   2,
				Spy:      1,
			},
			{
				SteamId:  76561198383642609,
				Scout:    1,
				Soldier:  3,
				Pyro:     1,
				Demoman:  1,
				Heavy:    4,
				Engineer: 1,
				Medic:    1,
				Sniper:   0,
				Spy:      3,
			},
			{
				SteamId:  76561198423392803,
				Scout:    0,
				Soldier:  1,
				Pyro:     0,
				Demoman:  1,
				Heavy:    0,
				Engineer: 2,
				Medic:    0,
				Sniper:   2,
				Spy:      1,
			},
			{
				SteamId:  76561198061174192,
				Scout:    0,
				Soldier:  1,
				Pyro:     0,
				Demoman:  1,
				Heavy:    0,
				Engineer: 0,
				Medic:    0,
				Sniper:   1,
				Spy:      0,
			},
			{
				SteamId:  76561198082713023,
				Scout:    0,
				Soldier:  0,
				Pyro:     0,
				Demoman:  0,
				Heavy:    0,
				Engineer: 0,
				Medic:    0,
				Sniper:   0,
				Spy:      1,
			},
			{
				SteamId:  76561198096251579,
				Scout:    2,
				Soldier:  1,
				Pyro:     1,
				Demoman:  1,
				Heavy:    0,
				Engineer: 0,
				Medic:    0,
				Sniper:   0,
				Spy:      2,
			},
			{
				SteamId:  76561198084686835,
				Scout:    7,
				Soldier:  1,
				Pyro:     1,
				Demoman:  2,
				Heavy:    1,
				Engineer: 2,
				Medic:    1,
				Sniper:   1,
				Spy:      2,
			},
			{
				SteamId:  76561198043171944,
				Scout:    1,
				Soldier:  7,
				Pyro:     0,
				Demoman:  2,
				Heavy:    1,
				Engineer: 4,
				Medic:    2,
				Sniper:   7,
				Spy:      4,
			},
			{
				SteamId:  76561198087442614,
				Scout:    0,
				Soldier:  0,
				Pyro:     2,
				Demoman:  5,
				Heavy:    0,
				Engineer: 1,
				Medic:    2,
				Sniper:   2,
				Spy:      0,
			},
			{
				SteamId:  76561198050517054,
				Scout:    0,
				Soldier:  6,
				Pyro:     2,
				Demoman:  2,
				Heavy:    2,
				Engineer: 2,
				Medic:    0,
				Sniper:   2,
				Spy:      2,
			},
			{
				SteamId:  76561198051884373,
				Scout:    4,
				Soldier:  0,
				Pyro:     1,
				Demoman:  4,
				Heavy:    1,
				Engineer: 1,
				Medic:    1,
				Sniper:   1,
				Spy:      0,
			},
			{
				SteamId:  76561198057150173,
				Scout:    2,
				Soldier:  0,
				Pyro:     1,
				Demoman:  1,
				Heavy:    1,
				Engineer: 2,
				Medic:    0,
				Sniper:   2,
				Spy:      0,
			},
			{
				SteamId:  76561198036671190,
				Scout:    0,
				Soldier:  1,
				Pyro:     3,
				Demoman:  1,
				Heavy:    4,
				Engineer: 0,
				Medic:    1,
				Sniper:   2,
				Spy:      1,
			},
			{
				SteamId:  76561198126692772,
				Scout:    1,
				Soldier:  4,
				Pyro:     0,
				Demoman:  3,
				Heavy:    0,
				Engineer: 1,
				Medic:    1,
				Sniper:   1,
				Spy:      2,
			},
			{
				SteamId:  76561198271399587,
				Scout:    1,
				Soldier:  3,
				Pyro:     1,
				Demoman:  2,
				Heavy:    0,
				Engineer: 1,
				Medic:    0,
				Sniper:   0,
				Spy:      1,
			},
			{
				SteamId:  76561198809011070,
				Scout:    4,
				Soldier:  3,
				Pyro:     2,
				Demoman:  1,
				Heavy:    0,
				Engineer: 2,
				Medic:    3,
				Sniper:   8,
				Spy:      3,
			},
			{
				SteamId:  76561198164892406,
				Scout:    0,
				Soldier:  1,
				Pyro:     1,
				Demoman:  0,
				Heavy:    4,
				Engineer: 2,
				Medic:    0,
				Sniper:   1,
				Spy:      3,
			},
		},
		ClassKillsAssists: []*logparse.MatchClassSums{
			{
				SteamId:  76561198113244106,
				Scout:    1,
				Soldier:  1,
				Pyro:     1,
				Demoman:  1,
				Heavy:    3,
				Engineer: 1,
				Medic:    1,
				Sniper:   0,
				Spy:      3,
			},
			{
				SteamId:  76561199050447792,
				Scout:    3,
				Soldier:  0,
				Pyro:     2,
				Demoman:  4,
				Heavy:    1,
				Engineer: 3,
				Medic:    2,
				Sniper:   3,
				Spy:      1,
			},
			{
				SteamId:  76561198383642609,
				Scout:    2,
				Soldier:  6,
				Pyro:     1,
				Demoman:  1,
				Heavy:    4,
				Engineer: 2,
				Medic:    1,
				Sniper:   0,
				Spy:      4,
			},
			{
				SteamId:  76561198423392803,
				Scout:    0,
				Soldier:  1,
				Pyro:     0,
				Demoman:  1,
				Heavy:    0,
				Engineer: 2,
				Medic:    0,
				Sniper:   2,
				Spy:      1,
			},
			{
				SteamId:  76561198061174192,
				Scout:    0,
				Soldier:  1,
				Pyro:     1,
				Demoman:  1,
				Heavy:    0,
				Engineer: 0,
				Medic:    0,
				Sniper:   1,
				Spy:      0,
			},
			{
				SteamId:  76561198082713023,
				Scout:    1,
				Soldier:  3,
				Pyro:     2,
				Demoman:  4,
				Heavy:    1,
				Engineer: 0,
				Medic:    0,
				Sniper:   0,
				Spy:      3,
			},
			{
				SteamId:  76561198096251579,
				Scout:    3,
				Soldier:  2,
				Pyro:     1,
				Demoman:  1,
				Heavy:    1,
				Engineer: 0,
				Medic:    1,
				Sniper:   2,
				Spy:      2,
			},
			{
				SteamId:  76561198084686835,
				Scout:    7,
				Soldier:  1,
				Pyro:     1,
				Demoman:  2,
				Heavy:    2,
				Engineer: 2,
				Medic:    1,
				Sniper:   1,
				Spy:      2,
			},
			{
				SteamId:  76561198043171944,
				Scout:    2,
				Soldier:  8,
				Pyro:     1,
				Demoman:  2,
				Heavy:    2,
				Engineer: 4,
				Medic:    3,
				Sniper:   7,
				Spy:      4,
			},
			{
				SteamId:  76561198087442614,
				Scout:    0,
				Soldier:  0,
				Pyro:     2,
				Demoman:  5,
				Heavy:    1,
				Engineer: 1,
				Medic:    2,
				Sniper:   2,
				Spy:      0,
			},
			{
				SteamId:  76561198050517054,
				Scout:    1,
				Soldier:  6,
				Pyro:     6,
				Demoman:  3,
				Heavy:    2,
				Engineer: 2,
				Medic:    0,
				Sniper:   3,
				Spy:      4,
			},
			{
				SteamId:  76561198051884373,
				Scout:    4,
				Soldier:  1,
				Pyro:     1,
				Demoman:  5,
				Heavy:    1,
				Engineer: 1,
				Medic:    1,
				Sniper:   2,
				Spy:      1,
			},
			{
				SteamId:  76561198057150173,
				Scout:    3,
				Soldier:  0,
				Pyro:     1,
				Demoman:  1,
				Heavy:    1,
				Engineer: 2,
				Medic:    0,
				Sniper:   2,
				Spy:      0,
			},
			{
				SteamId:  76561198036671190,
				Scout:    0,
				Soldier:  1,
				Pyro:     4,
				Demoman:  2,
				Heavy:    5,
				Engineer: 1,
				Medic:    1,
				Sniper:   2,
				Spy:      1,
			},
			{
				SteamId:  76561198126692772,
				Scout:    3,
				Soldier:  5,
				Pyro:     0,
				Demoman:  3,
				Heavy:    1,
				Engineer: 1,
				Medic:    2,
				Sniper:   1,
				Spy:      5,
			},
			{
				SteamId:  76561198271399587,
				Scout:    3,
				Soldier:  3,
				Pyro:     2,
				Demoman:  3,
				Heavy:    2,
				Engineer: 2,
				Medic:    1,
				Sniper:   1,
				Spy:      5,
			},
			{
				SteamId:  76561198809011070,
				Scout:    4,
				Soldier:  4,
				Pyro:     2,
				Demoman:  1,
				Heavy:    1,
				Engineer: 2,
				Medic:    3,
				Sniper:   8,
				Spy:      3,
			},
			{
				SteamId:  76561198164892406,
				Scout:    1,
				Soldier:  1,
				Pyro:     3,
				Demoman:  1,
				Heavy:    7,
				Engineer: 5,
				Medic:    0,
				Sniper:   2,
				Spy:      3,
			},
		},
		ClassDeaths: []*logparse.MatchClassSums{
			{
				SteamId:  76561198113244106,
				Scout:    2,
				Soldier:  3,
				Pyro:     0,
				Demoman:  0,
				Heavy:    1,
				Engineer: 0,
				Medic:    0,
				Sniper:   2,
				Spy:      1,
			},
			{
				SteamId:  76561199050447792,
				Scout:    1,
				Soldier:  2,
				Pyro:     1,
				Demoman:  2,
				Heavy:    1,
				Engineer: 0,
				Medic:    0,
				Sniper:   3,
				Spy:      2,
			},
			{
				SteamId:  76561198383642609,
				Scout:    4,
				Soldier:  1,
				Pyro:     0,
				Demoman:  4,
				Heavy:    1,
				Engineer: 0,
				Medic:    1,
				Sniper:   0,
				Spy:      0,
			},
			{
				SteamId:  76561198423392803,
				Scout:    7,
				Soldier:  5,
				Pyro:     1,
				Demoman:  0,
				Heavy:    0,
				Engineer: 1,
				Medic:    0,
				Sniper:   4,
				Spy:      1,
			},
			{
				SteamId:  76561198061174192,
				Scout:    4,
				Soldier:  5,
				Pyro:     0,
				Demoman:  0,
				Heavy:    1,
				Engineer: 3,
				Medic:    0,
				Sniper:   1,
				Spy:      0,
			},
			{
				SteamId:  76561198082713023,
				Scout:    0,
				Soldier:  0,
				Pyro:     1,
				Demoman:  1,
				Heavy:    1,
				Engineer: 0,
				Medic:    0,
				Sniper:   0,
				Spy:      2,
			},
			{
				SteamId:  76561198096251579,
				Scout:    0,
				Soldier:  1,
				Pyro:     4,
				Demoman:  1,
				Heavy:    2,
				Engineer: 0,
				Medic:    0,
				Sniper:   2,
				Spy:      5,
			},
			{
				SteamId:  76561198084686835,
				Scout:    2,
				Soldier:  0,
				Pyro:     0,
				Demoman:  0,
				Heavy:    4,
				Engineer: 1,
				Medic:    0,
				Sniper:   1,
				Spy:      1,
			},
			{
				SteamId:  76561198043171944,
				Scout:    0,
				Soldier:  2,
				Pyro:     3,
				Demoman:  0,
				Heavy:    6,
				Engineer: 0,
				Medic:    0,
				Sniper:   0,
				Spy:      0,
			},
			{
				SteamId:  76561198087442614,
				Scout:    5,
				Soldier:  3,
				Pyro:     1,
				Demoman:  2,
				Heavy:    3,
				Engineer: 1,
				Medic:    1,
				Sniper:   1,
				Spy:      0,
			},
			{
				SteamId:  76561198050517054,
				Scout:    2,
				Soldier:  2,
				Pyro:     2,
				Demoman:  0,
				Heavy:    3,
				Engineer: 1,
				Medic:    0,
				Sniper:   1,
				Spy:      1,
			},
			{
				SteamId:  76561198051884373,
				Scout:    3,
				Soldier:  1,
				Pyro:     3,
				Demoman:  1,
				Heavy:    2,
				Engineer: 1,
				Medic:    3,
				Sniper:   2,
				Spy:      0,
			},
			{
				SteamId:  76561198057150173,
				Scout:    7,
				Soldier:  3,
				Pyro:     3,
				Demoman:  1,
				Heavy:    3,
				Engineer: 6,
				Medic:    0,
				Sniper:   0,
				Spy:      0,
			},
			{
				SteamId:  76561198036671190,
				Scout:    2,
				Soldier:  1,
				Pyro:     2,
				Demoman:  1,
				Heavy:    1,
				Engineer: 2,
				Medic:    0,
				Sniper:   4,
				Spy:      3,
			},
			{
				SteamId:  76561198126692772,
				Scout:    0,
				Soldier:  1,
				Pyro:     0,
				Demoman:  0,
				Heavy:    1,
				Engineer: 1,
				Medic:    0,
				Sniper:   1,
				Spy:      1,
			},
			{
				SteamId:  76561198271399587,
				Scout:    1,
				Soldier:  1,
				Pyro:     0,
				Demoman:  3,
				Heavy:    1,
				Engineer: 0,
				Medic:    0,
				Sniper:   0,
				Spy:      2,
			},
			{
				SteamId:  76561198809011070,
				Scout:    1,
				Soldier:  0,
				Pyro:     4,
				Demoman:  1,
				Heavy:    1,
				Engineer: 1,
				Medic:    0,
				Sniper:   1,
				Spy:      0,
			},
			{
				SteamId:  76561198164892406,
				Scout:    1,
				Soldier:  4,
				Pyro:     1,
				Demoman:  2,
				Heavy:    1,
				Engineer: 0,
				Medic:    0,
				Sniper:   3,
				Spy:      2,
			},
			{
				SteamId:  76561198073709029,
				Scout:    0,
				Soldier:  0,
				Pyro:     0,
				Demoman:  1,
				Heavy:    0,
				Engineer: 0,
				Medic:    0,
				Sniper:   0,
				Spy:      0,
			},
		},
	}
	return match, map[string]steamid.SID64{
		"var":                76561198164892406,
		"para":               76561198057150173,
		"sentar":             76561198126692772,
		"Pride (Pyro Main)":  76561198036671190,
		"jumbuck":            76561198084686835,
		"freakyjoy1.ttv":     76561198061174192,
		"avg Q enjoyer":      76561198113244106,
		"ExCalibre":          76561198423392803,
		"nomodick":           76561198051884373,
		"Lochlore":           76561198087442614,
		"Link":               76561198043171944,
		"Tunaaaaaa":          76561198809011070,
		"Tiger":              76561198271399587,
		"Invidia":            76561198096251579,
		"El Sur":             76561198383642609,
		"maz":                76561198050517054,
		"Golden Terrestrial": 76561198082713023,
		"Doctrine":           76561199050447792,
		"WitlessConnor":      76561198073709029,
	}
}
