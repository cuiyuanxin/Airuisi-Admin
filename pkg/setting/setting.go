package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

var sections = make(map[string]interface{})

func (s *Setting) readSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	fmt.Println(k, v, sections[k])
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
		err := s.readSection(k, v)
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

func NewSetting(configs ...string) error {
	vp := viper.New()
	vp.AddConfigPath("config/")
	vp.SetConfigName("prod")
	vp.SetConfigType("yaml")

	if len(configs) > 0 {
		for _, config := range configs {
			if config != "" {
				vp.AddConfigPath(config)
			}
		}
	}

	if err := vp.ReadInConfig(); err != nil {
		return err
	}

	s := &Setting{vp}
	s.watchSettingChange()

	return nil
}
