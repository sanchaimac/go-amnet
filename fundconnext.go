package fundconnext

import (
	"time"
)

type FundConnext struct {
	cfg   *FCConfiguration
	token string
}

type FCConfiguration struct {
	Username string
	Password string
	Timeout  time.Duration
	Env      string
}

func ToOptStr(str string) *string {
	r := &str
	return r
}

func ToOptBool(flag bool) *bool {
	r := &flag
	return r
}

func ToOptFloat(num float32) *float32 {
	r := &num
	return r
}

func ToOptFloat64(num float64) *float64 {
	r := &num
	return r
}

func ToOptInt(num int) *int {
	r := &num
	return r
}

func MakeAPICallerConfig(f *FundConnext) *APICallerConfig {
	return &APICallerConfig{
		Timeout: f.cfg.Timeout,
	}
}

func (f *FundConnext) Configure(cfg *FCConfiguration) {
	f.cfg = cfg
	f.cfg.Timeout = time.Duration(10 * time.Second)
	if cfg.Env == "" {
		f.cfg.Env = "production"
	}
}

func (f *FundConnext) Start() error {
	token, err := Login(f.cfg.Username, f.cfg.Password)
	if err != nil {
		return err
	}
	f.token = token.AccessToken
	return nil
}

func MakeFundconnext() *FundConnext {
	return &FundConnext{}
}
