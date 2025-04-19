package epulse

import (
	dtypes "git.n-hub.ru/neosy/npulse-watcher/domain/types"
)

type PulseState struct {
	IPAddress  string                  `json:"ip_address"`
	HostName   string                  `json:"host_hame"`
	LastTime   string                  `json:"last_time"`
	Status     dtypes.PulseStateStatus `json:"status"`
	IsNotified bool                    `json:"is_notified"`
}
