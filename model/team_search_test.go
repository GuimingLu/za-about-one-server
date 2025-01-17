


package model

import (
	"strings"
	"testing"
)

func TestTeamSearchJson(t *testing.T) {
	teamSearch := TeamSearch{Term: NewId()}
	json := teamSearch.ToJson()
	rteamSearch := ChannelSearchFromJson(strings.NewReader(json))

	if teamSearch.Term != rteamSearch.Term {
		t.Fatal("Terms do not match")
	}
}
