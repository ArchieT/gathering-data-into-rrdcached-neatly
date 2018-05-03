package abundleof

import (
	"github.com/ArchieT/gathering-data-into-rrdcached-neatly/lib/thetypes"
	"github.com/ArchieT/gathering-data-into-rrdcached-neatly/lib/thetypes/systhermaltemp"
	"github.com/ArchieT/gathering-data-into-rrdcached-neatly/lib/thetypes/ds18b20"
)

var components = []thetypes.BundlePiece{
	systhermaltemp.SysThermalTempReadingBundle,
	ds18b20.ReadDS18B20Bundle,
}
