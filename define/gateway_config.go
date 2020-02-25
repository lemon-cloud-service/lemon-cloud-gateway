package define

type GatewayConfig struct {
	Gateway struct {
		Server struct {
			HttpPort uint16 `json:"http_port" yaml:"http_port"`
		}
	}
}
