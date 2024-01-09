package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

var sections = make(map[string]interface{})

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	if _, ok := sections[k]; !ok {
		sections[k] = v
	}

	return nil
}

func (s *Setting) reloadAllSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Setting) watchSettingChange() {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			_ = s.reloadAllSection()
		})
	}()
}

func NewSetting(configs string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigType("yaml")

	if configs != "" {
		vp.SetConfigFile(configs)
	} else {
		vp.AddConfigPath("config/")
		vp.SetConfigName("prod")
	}

	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}

	s := &Setting{vp}
	s.watchSettingChange()

	return s, nil
}
