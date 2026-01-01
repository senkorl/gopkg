package main

import (
	"fmt"
	"os/exec"

	"gopkg.in/yaml.v3"
)

type Config struct {
	BandVersion string `yaml:"bandversion"`
	SimReady    int    `yaml:"sim_ready"`
	connect     int    `yaml:"connect"`
	iccid       int    `yaml:"iccid"`
	imsi        int    `yaml:"imsi"`
	rssid       int    `yaml:"rssid"`
}

func main() {
	// 打开YAML文件
	//file, err := os.Open("/Users/qiuxi/GolandProjects/gopkg/lib/yam/quectel_config.conf")
	//if err != nil {
	//	log.Fatalf("error opening file: %v", err)
	//}
	//defer file.Close()

	cmdIns := exec.Command("sh", "-c", "cat /Users/qiuxi/GolandProjects/gopkg/lib/yam/quectel_config.conf")
	out, _ := cmdIns.CombinedOutput()
	// 解析YAML文件
	var config Config
	_ = yaml.Unmarshal(out, &config)

	// 输出解析后的结果
	fmt.Printf("Name: %s\n", config.BandVersion)
	fmt.Printf("Version: %d\n", config.SimReady)
}

//
