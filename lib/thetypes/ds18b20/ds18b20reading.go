package ds18b20

import (
	"github.com/ArchieT/gathering-data-into-rrdcached-neatly/lib/thetypes/afile"
	"github.com/ArchieT/gathering-data-into-rrdcached-neatly/lib/thetypes"
)

func ReadDS18B20(path string) (int, error) {
	return afile.ScanTheEndOfTheFilePrefixedBy([]byte("t="),path,8)
}

var ReadDS18B20Bundle = thetypes.BundlePiece{thetypes.IsDS18B20, ReadDS18B20}
