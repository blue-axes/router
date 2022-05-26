package config

import "github.com/blue-axes/router/entity/route"

type (
	Config struct {
		BoxID     string     `yaml:"BoxID"`
		Logger    Logger     `yaml:"Logger"`
		Http      Http       `yaml:"Http"`
		GRpc      GRpc       `yaml:"GRpc"`
		Upstreams []Upstream `yaml:"Upstreams"`
		Routes    []Route    `yaml:"Routes"`
	}

	LoggerType   string
	LoggerFormat string
	Logger       struct {
		Debug        bool         `yaml:"Debug"`
		Format       LoggerFormat `yaml:"Format"`
		Level        string       `yaml:"Level"`
		Type         LoggerType   `yaml:"Type"`
		LogFile      string       `yaml:"LogFile"`
		LoggerRotate LoggerRotate `yaml:"LoggerRotate"`
	}

	LoggerRotatePolicy string
	LoggerRotate       struct {
		Policy               LoggerRotatePolicy `yaml:"Policy"`
		RotateSizeByte       uint64             `yaml:"MaxSizeByte"`
		RotateDurationSecond uint64             `yaml:"RotateDurationSecond"`
		MaxBackupCount       int                `yaml:"MaxBackupCount"`
	}

	Http struct {
		Host string `yaml:"Host"`
		Port uint16 `yaml:"Port"`
	}

	GRpc struct {
		Host     string `yaml:"Host"`
		Port     uint16 `yaml:"Port"`
		CertFile string `yaml:"CertFile"`
		KeyFile  string `yaml:"KeyFile"`
		RootCa   string `yaml:"RootCa"`
	}

	Upstream struct {
		Name                    string `yaml:"Name"`
		Host                    string `yaml:"Host"`
		Port                    uint16 `yaml:"Port"`
		CertFile                string `yaml:"CertFile"`
		KeyFile                 string `yaml:"KeyFile"`
		RootCa                  string `yaml:"RootCa"`
		ReconnectIntervalSecond int64  `yaml:"ReconnectIntervalSecond"`
		ReconnectTimes          int    `yaml:"ReconnectTimes"`
	}

	Route struct {
		PointOne route.Point `yaml:"PointOne"`
		PointTwo route.Point `yaml:"PointTwo"`
	}
)

const (
	LoggerFormatText LoggerFormat = "text"
	LoggerFormatJson LoggerFormat = "json"
	LoggerTypeStdout LoggerType   = "stdout"
	LoggerTypeFile   LoggerType   = "file"
)
