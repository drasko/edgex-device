package export

// Compression algorithm types
const (
	CompNone = iota
	CompGzip
	CompZip
)

// Data format types
const (
	FormatJSON = iota
	FormatXML
	FormatSerialized
	FormatIoTCoreJSON
	FormatAzureJSON
	FormatCSV
)

// Export destination types
const (
	DestMQTT = iota
	DestZMQ
	DestIotCoreMQTT
	DestAzureMQTT
	DestRest
)

// Registration - Defines the registration details
// on the part of north side export clients
type Registration struct {
	Name        string
	Addr        Addressable
	Format      int
	Filter      Filter
	Encryption  EncryptionDetails
	Compression int
	Enable      bool
	Destination int
}
