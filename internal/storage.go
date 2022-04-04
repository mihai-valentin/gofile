package internal

import "os"

type Storage struct {
	Config *Config
	Paths  []string
}

func NewStorage(config *Config) *Storage {
	return &Storage{
		Config: config,
	}
}

func (s *Storage) Init() error {
	root := s.Config.Get("storage.root")
	disks := s.Config.GetSlice("storage.disks")

	for _, disk := range disks {
		path := root + "/" + disk
		if _, err := os.Stat(path); os.IsExist(err) {
			continue
		}

		if err := os.MkdirAll(path, os.ModeDir); err != nil {
			return err
		}

		s.Paths = append(s.Paths, path)
	}

	return nil
}
