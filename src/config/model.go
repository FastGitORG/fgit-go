package config

import (
	"encoding/json"
	"fgit-go/shared"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type ConfModel struct {
	GithubProxy   string `json:"github_proxy"`
	RawProxy      string `json:"raw_proxy"`
	DownloadProxy string `json:"download_proxy"`
}

func (c *ConfModel) Apply() {
	if c.GithubProxy != "" {
		shared.GitMirror = c.GithubProxy
	}
	if c.RawProxy != "" {
		shared.RawMirror = c.RawProxy
	}
	if c.DownloadProxy != "" {
		shared.DownloadMirror = c.DownloadProxy
	}
}

func (c *ConfModel) SaveConfig() {
	target, err := ensureConfigPath()
	if err != nil {
		fmt.Println("Construct path failed")
		return
	}
	s, err := json.Marshal(c)
	if err != nil {
		fmt.Println("Parse conf to json failed")
		return
	}
	err = ioutil.WriteFile(filepath.Join(target, "config.json"), s, 0644)
	if err != nil {
		fmt.Println("Write to file failed")
		return
	}
}

func ensureConfigPath() (path string, err error) {
	path, err = os.UserHomeDir()
	if err != nil {
		fmt.Println("Get Home Directory Failed")
		return
	}
	path = filepath.Join(path, ".config", "fgit-go")
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Println("Create Configuation Path Failed")
	}
	return
}

func ReadConfig() {
	target, err := ensureConfigPath()
	shared.CheckErr(err, "Construct path failed", 1)
	if err != nil {
		fmt.Println()
		return
	}
	target = filepath.Join(target, "config.json")
	if !shared.IsExists(target) {
		conf := ConfModel{
			GithubProxy:   shared.GitMirror,
			RawProxy:      shared.RawMirror,
			DownloadProxy: shared.DownloadMirror,
		}
		conf.SaveConfig()
		return
	}
	jsonFile, err := os.Open(target)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var conf ConfModel
	err = json.Unmarshal(byteValue, &conf)
	if err != nil {
		return
	}
	conf.Apply()
}
