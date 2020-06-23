package drone

import "github.com/profzone/eden-framework/internal/project/drone/enums"

type PipelinePlatform struct {
	OS           enums.DroneCiPlatformOs   `yaml:"os"`
	Architecture enums.DroneCiPlatformArch `yaml:"arch"`
	Version      int                       `yaml:"version,omitempty"`
}

func (p *PipelinePlatform) WithOS(os enums.DroneCiPlatformOs) *PipelinePlatform {
	p.OS = os
	return p
}
