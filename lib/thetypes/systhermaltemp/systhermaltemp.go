package systhermaltemp

import (
	"github.com/ArchieT/gathering-data-into-rrdcached-neatly/lib/thetypes/afile"
	"github.com/ArchieT/gathering-data-into-rrdcached-neatly/lib/thetypes"
)

func ReadSysThermalTemp(path string) (int, error) {
	return afile.SimpleReadInt(path)
}

var SysThermalTempReadingBundle = thetypes.BundlePiece{thetypes.IsSysThermalTemp, ReadSysThermalTemp}
