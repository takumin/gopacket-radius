package radius

import "github.com/google/gopacket/layers"

func init() {
	layers.RegisterUDPPortLayerType(1812, LayerTypeRADIUS)
}
