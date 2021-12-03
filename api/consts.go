package centreonapi

const (
	Disable = "0"
	Enable  = "1"
	Default = "2"
)

var clapiApiParams map[string]string = map[string]string{
	"action": "action",
	"object": "centreon_clapi",
}
