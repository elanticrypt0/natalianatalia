package nnconfig

import "github.com/k23dev/go4it"

type NNConfig struct {
	Tanga_fields_file string `json:"tanga_fields_file"`
	Scripts_folder    string `json:"scripts_folder"`
	Scripts_list      string `json:"scripts_list"`
}

func NewNNConfig() *NNConfig {
	nnconfig := &NNConfig{}
	// load config
	go4it.ReadAndParseToml("./config/nn.config.toml", &nnconfig)
	return nnconfig
}
