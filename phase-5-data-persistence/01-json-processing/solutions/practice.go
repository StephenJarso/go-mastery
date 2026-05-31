package solutions

import (
	"encoding/json"
	"time"
	"errors"
)


type Config struct {
	AppName string `json:"app_name"`
	Port    int    `json:"port"`
}

type PracticeCustomDate struct {
	time.Time
}

type PracticeCustomDateConfig struct {
	Created PracticeCustomDate `json:"created"`
}

func ParseConfig(data []byte) (Config, error) {
	var c Config
	if err := json.Unmarshal(data, &c); err != nil {
		return c, err
	}
	if c.AppName == "" {
		return c, errors.New("empty AppName")
	}
	return c, nil
}

func (cd *PracticeCustomDate) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	t, err := time.Parse("02-01-2006", s)
	if err != nil {
		return err
	}
	cd.Time = t
	return nil
}
