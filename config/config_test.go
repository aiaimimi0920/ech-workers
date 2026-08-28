package config

import "testing"

func TestValidateRequiresServerAndToken(t *testing.T) {
	tests := []struct {
		name string
		cfg  Config
	}{
		{name: "server", cfg: Config{ListenAddr: "127.0.0.1:30000", Token: "token"}},
		{name: "token", cfg: Config{ListenAddr: "127.0.0.1:30000", ServerAddr: "worker.example:443"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cfg.Validate(); err == nil {
				t.Fatal("Validate() accepted incomplete configuration")
			}
		})
	}
}

func TestValidateNormalizesListenAddress(t *testing.T) {
	cfg := Config{ListenAddr: "127.0.0.1", ServerAddr: "worker.example:443", Token: "token"}
	if err := cfg.Validate(); err != nil {
		t.Fatal(err)
	}
	if cfg.ListenAddr != "127.0.0.1:30000" {
		t.Fatalf("ListenAddr = %q", cfg.ListenAddr)
	}
}
