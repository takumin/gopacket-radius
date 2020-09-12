package radius

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

var LayerTypeRADIUS = gopacket.RegisterLayerType(1812, gopacket.LayerTypeMetadata{Name: "RADIUS", Decoder: gopacket.DecodeFunc(decodeRADIUS)})

func init() {
	layers.RegisterUDPPortLayerType(1812, LayerTypeRADIUS)
}
