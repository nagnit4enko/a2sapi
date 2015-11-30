package steam

// query.go - Used for querying individual game servers to retrieve their info
// for building a list to return to the API

import (
	"steamtest/src/models"
	"steamtest/src/steam/filters"
	"steamtest/src/util"
	"sync"
)

func batchInfoQuery(servers []string) map[string]*models.SteamServerInfo {
	m := make(map[string]*models.SteamServerInfo)
	var wg sync.WaitGroup
	var mut sync.Mutex
	var failed []string

	for _, h := range servers {
		wg.Add(1)
		go func(host string) {
			serverinfo, err := GetInfoForServer(host, QueryTimeout)
			if err != nil {
				mut.Lock()
				failed = append(failed, host)
				mut.Unlock()
				wg.Done()
				return
			}
			mut.Lock()
			m[host] = serverinfo
			mut.Unlock()
			wg.Done()
		}(h)
	}
	wg.Wait()
	retried := RetryFailedInfoReq(failed, 3)
	for k, v := range retried {
		m[k] = v
	}
	return m
}

func batchPlayerQuery(servers []string) map[string][]*models.SteamPlayerInfo {
	m := make(map[string][]*models.SteamPlayerInfo)
	var wg sync.WaitGroup
	var mut sync.Mutex
	var failed []string

	for _, h := range servers {
		wg.Add(1)
		go func(host string) {
			players, err := GetPlayersForServer(host, QueryTimeout)
			if err != nil {
				// server could just be empty
				if err != ErrNoPlayers {
					mut.Lock()
					failed = append(failed, host)
					mut.Unlock()
					wg.Done()
					return
				}
			}
			mut.Lock()
			m[host] = players
			mut.Unlock()
			wg.Done()
		}(h)
	}
	wg.Wait()
	retried := RetryFailedPlayersReq(failed, QueryRetryCount)
	for k, v := range retried {
		m[k] = v
	}
	return m
}

func batchRuleQuery(servers []string) map[string]map[string]string {
	m := make(map[string]map[string]string)
	var wg sync.WaitGroup
	var mut sync.Mutex
	var failed []string
	for _, h := range servers {
		wg.Add(1)
		go func(host string) {
			rules, err := GetRulesForServer(host, QueryTimeout)
			if err != nil {
				// server might have no rules
				if err != ErrNoRules {
					mut.Lock()
					failed = append(failed, host)
					mut.Unlock()
					wg.Done()
					return
				}
			}
			mut.Lock()
			m[host] = rules
			mut.Unlock()
			wg.Done()
		}(h)
	}
	wg.Wait()
	retried := RetryFailedRulesReq(failed, QueryRetryCount)
	for k, v := range retried {
		m[k] = v
	}
	return m
}

// func DirectQuery(hosts []string) (*models.APIServerList, error) {
// 	players := getPlayersForServers(hosts)
// 	rules := getRulesForServers(hosts)
// 	info := getInfoForServers(hosts)
// 	f := filters.NewFilter(filters.GameUnspecified, filters.SrAll, nil)
// 	sl, err := buildServerList(f, hosts, info, rules, players)
// 	if err != nil {
// 		return models.GetDefaultServerList(), util.LogAppError(err)
// 	}
// 	return sl, nil
// }

func Query(hostsgames map[string]string) (*models.APIServerList, error) {
	hg := make(map[string]*filters.Game, len(hostsgames))
	needsPlayers := make([]string, len(hostsgames))
	needsRules := make([]string, len(hostsgames))
	needsInfo := make([]string, len(hostsgames))

	for host, game := range hostsgames {
		fg := filters.GetGame(game)
		hg[host] = fg
		if !fg.IgnorePlayers {
			needsPlayers = append(needsPlayers, host)
		}
		if !fg.IgnoreRules {
			needsRules = append(needsRules, host)
		}
		if !fg.IgnoreInfo {
			needsInfo = append(needsInfo, host)
		}
	}
	players := batchPlayerQuery(needsPlayers)
	rules := batchRuleQuery(needsRules)
	info := batchInfoQuery(needsInfo)

	sl, err := buildQueryServerList(hg, info, rules, players)
	if err != nil {
		return models.GetDefaultServerList(), util.LogAppError(err)
	}
	return sl, nil
}
