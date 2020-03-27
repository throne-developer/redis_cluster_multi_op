package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/pihao/redis-go-cluster"
)

type ConfigData struct {
	CommandFile string   `json:"command_file"`
	RedisAddrs  []string `json:"redis_addrs"`
	RedisPwd    string   `json:"redis_pwd"`
}

var conf = new(ConfigData)

func main() {
	if err := loadConfig(); err != nil {
		LogError(err.Error())
		return
	}

	cluster, err := redis.NewCluster(
		&redis.Options{
			StartNodes:   conf.RedisAddrs,
			ConnTimeout:  500 * time.Millisecond,
			ReadTimeout:  500 * time.Millisecond,
			WriteTimeout: 500 * time.Millisecond,
			KeepAlive:    16,
			AliveTime:    60 * time.Second,
			Password:     conf.RedisPwd,
		})
	if err != nil {
		LogError(err.Error())
		return
	}

	var Commands [][]string
	for _, line := range LoadFile(conf.CommandFile) {
		if fields := strings.Split(line, ","); len(fields) > 1 {
			Commands = append(Commands, fields)
		}
	}

	batch := cluster.NewBatch()
	for _, cmd := range Commands {
		batch.Put(cmd[0], string2Interface(cmd[1:])...)
	}
	if replys, err := cluster.RunBatch(batch); err != nil {
		LogError("run err " + err.Error())
	} else {
		for i, reply := range replys {
			fmt.Println("reply ", i, reply)
		}
	}
	fmt.Println("done count ", len(Commands))
}

func loadConfig() error {
	f, err := os.Open("config.json")
	if err != nil {
		return err
	}

	defer f.Close()
	if err := json.NewDecoder(f).Decode(conf); err != nil {
		return err
	}
	return nil
}

func string2Interface(data []string) []interface{} {
	results := make([]interface{}, 0, len(data))
	for _, item := range data {
		results = append(results, item)
	}
	return results
}
