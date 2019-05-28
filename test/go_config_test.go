package test

import (
	"microcosm/internal"
	"testing"
)

func TestConfig(t *testing.T) {
	cfg := internal.New("test").GetCfg()
	val, err := cfg.Section("db").GetKey("file")
	t.Log(val, err)
}
