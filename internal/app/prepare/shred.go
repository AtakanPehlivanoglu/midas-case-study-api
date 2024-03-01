package prepare

import "github.com/lu4p/shred"

func ShredConfig(times int, zeros, remove bool) *shred.Conf {
	return &shred.Conf{Times: times, Zeros: zeros, Remove: remove}
}
