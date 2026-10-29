package registry

import "testing"

func TestCodexTerraPaidTiersAndUpstreamPrecedence(t *testing.T) {
	for name, models := range map[string][]*ModelInfo{"plus": GetCodexPlusModels(), "team": GetCodexTeamModels(), "pro": GetCodexProModels()} {
		var count int
		for _, m := range models {
			if m.ID == "gpt-6.1-terra" {
				count++
				if m.Thinking == nil || len(m.Thinking.Levels) == 0 {
					t.Errorf("%s missing reasoning metadata", name)
				}
			}
		}
		if count != 1 {
			t.Errorf("%s: found %d terra definitions", name, count)
		}
	}
	for _, m := range GetCodexFreeModels() {
		if m.ID == "gpt-6.1-terra" {
			t.Error("terra unexpectedly enabled for free tier")
		}
	}
	upstream := &ModelInfo{ID: "gpt-6.1-terra", ContextLength: 12345}
	models := withCodexTerra([]*ModelInfo{upstream})
	if len(models) != 1 || models[0] != upstream || models[0].ContextLength != 12345 {
		t.Error("upstream metadata was replaced")
	}
}
