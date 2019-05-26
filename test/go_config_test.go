package test

import (
	"microcosm/conf"
	"testing"
)

func TestConfig(t *testing.T) {
	cfg := conf.New("test").GetCfg()
	val, err := cfg.Section("db").GetKey("file")
	t.Log(val, err)
}
