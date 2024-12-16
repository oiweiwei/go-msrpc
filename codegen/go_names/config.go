package go_names

import (
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"
)

var defaultConfig *Config

func init() { defaultConfig = LoadConfig() }

type Config struct {
	Rename     map[string][]string            `yaml:"rename"`
	Rotate     map[string]string              `yaml:"rotate"`
	Abbr       map[string]string              `yaml:"abbr"`
	Split      map[string][]string            `yaml:"split"`
	LookBehind map[string]map[string][]string `yaml:"lookBehind"`
	Trim       struct {
		Words     []string            `yaml:"word"`
		Word      map[string]struct{} `yaml:"-"`
		PrefixAll []string            `yaml:"prefixAll"`
		Prefixes  []string            `yaml:"prefix"`
		Prefix    map[string]struct{} `yaml:"-"`
		Suffix    []string            `yaml:"suffix"`
	}
}

func LoadConfig() *Config {

	_, p, _, ok := runtime.Caller(0)
	if !ok {
		panic("cannot determine current dir")
	}

	cc := Config{}

	cc.Trim.Prefix = make(map[string]struct{})
	cc.Trim.Word = make(map[string]struct{})

	fileNames, _ := filepath.Glob(filepath.Join(filepath.Dir(p), "config", "*.yaml"))

	for _, fileName := range fileNames {
		f, err := os.Open(fileName)
		if err != nil {
			panic(err.Error())
		}
		if err := yaml.NewDecoder(f).Decode(&cc); err != nil {
			panic(err.Error())
		}
		for _, pfx := range cc.Trim.Prefixes {
			cc.Trim.Prefix[pfx] = struct{}{}
		}
		for _, wrd := range cc.Trim.Words {
			cc.Trim.Word[wrd] = struct{}{}
		}
	}

	return &cc
}
