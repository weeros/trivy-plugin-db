package internal

import (
    "os"
    "path/filepath"
    "gopkg.in/yaml.v3"
)

type TrivyConfig struct {
    Cache struct {
        Dir string `yaml:"dir"`
    } `yaml:"cache"`
}

func GetTrivyDBPath() string {
    if envCacheDir := os.Getenv("TRIVY_CACHE_DIR"); envCacheDir != "" {
        return filepath.Join(envCacheDir, "db", "trivy.db")
    }

    configPaths := []string{
        "trivy.yaml",
        filepath.Join(os.Getenv("HOME"), ".trivy", "trivy.yaml"),
    }

    for _, path := range configPaths {
        if _, err := os.Stat(path); err == nil {
            content, err := os.ReadFile(path)
            if err != nil {
                continue
            }

            var cfg TrivyConfig
            if err := yaml.Unmarshal(content, &cfg); err != nil {
                continue
            }

            if cfg.Cache.Dir != "" {
                return filepath.Join(cfg.Cache.Dir, "db", "trivy.db")
            }
        }
    }

    return filepath.Join(os.Getenv("HOME"), ".cache", "trivy", "db", "trivy.db")
}