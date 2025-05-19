package config

import "testing"

func TestLoadConfig(t *testing.T) {
	conf, err := LoadConfig("../config.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(conf)
}
