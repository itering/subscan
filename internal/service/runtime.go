package service

import (
	"github.com/itering/subscan/internal/substrate/metadata"
	"github.com/itering/subscan/internal/substrate/rpc"
	"strings"
)

func (s *Service) regRuntimeVersion(name string, spec int) error {
	if affected := s.Dao.CreateRuntimeVersion(name, spec); affected > 0 {
		if coded := s.regCodecMetadata(); coded != "" {
			runtime := metadata.RegNewMetadataType(spec, coded)
			s.setRuntimeData(spec, runtime, coded)
		}
	}
	return nil
}

func (s *Service) regCodecMetadata() string {
	if coded, err := rpc.GetMetadataByHash(); err == nil {
		return coded
	}
	return ""
}

func (s *Service) MigrateRuntimeVersion(spec int, runtime *metadata.MetadataType, rawData string) {
	s.setRuntimeData(spec, runtime, rawData)
}

func (s *Service) setRuntimeData(spec int, runtime *metadata.MetadataType, rawData string) {
	var modules []string
	for _, value := range runtime.Metadata.Modules {
		modules = append(modules, value.Name)
	}
	s.Dao.SetRuntimeData(spec, strings.Join(modules, "|"), rawData)
}
