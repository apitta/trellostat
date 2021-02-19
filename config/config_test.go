package config_test

import (
	"testing"
	"trellostat/config"
)

func TestConfigFile(t *testing.T) {
	t.Parallel()
	c, err := config.ConfigFile("../trellostat.json")
	if err != nil {
		t.Error(err)
	}
	t.Log(c)
}
