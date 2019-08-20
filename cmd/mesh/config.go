package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"

	json "encoding/json"

	cli "gopkg.in/urfave/cli.v1"
)

type Config struct {
	ConfigMap map[string]interface{}
}

func NewConfig(ctx *cli.Context) *Config {
	c := &Config{
		ConfigMap: make(map[string]interface{}),
	}

	configFile := ctx.String(CONFIG_FILE)

	c.applyConfigFromFile(configFile, ctx)
	//c.applyConfigFromConsole(ctx)

	return c
}

func (c *Config) applyConfigFromFile(file string, ctx *cli.Context) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	parsedVal := make(map[string]string)

	err = json.Unmarshal(data, &parsedVal)
	if err != nil {
		return err
	}

	// TODO: Is passed key in our config or user passing some bogus key config.
	for k, v := range parsedVal {
		fmt.Println(k, v)
		c.ConfigMap[k] = v
	}

	return nil
}

func (c *Config) applyConfigFromConsole(ctx *cli.Context) {
	for _, n := range ctx.GlobalFlagNames() {
		g := ctx.Generic(n)
		if g == nil {
			continue
		}
		c.ConfigMap[n] = g
	}
}

func (c *Config) Bool(n string) bool {
	if val, ok := c.ConfigMap[n]; ok {
		parsed, err := strconv.ParseBool(val.(flag.Value).String())
		if err != nil {
			return false
		}
		return parsed
	}
	return false
}

func (c *Config) String(n string) string {
	if val, ok := c.ConfigMap[n]; ok {
		return val.(string)
	}
	return ""
}
