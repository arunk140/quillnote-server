package quilltypes

type OCSNotes struct {
	APIVersion []string `json:"api_version"`
	Version    string   `json:"version"`
}
type OCSCapabilities struct {
	Notes OCSNotes `json:"notes"`
}
type OCSData struct {
	Capabilities OCSCapabilities `json:"capabilities"`
}
type OCS struct {
	Data OCSData `json:"data"`
}
type Capabilities struct {
	OCS OCS `json:"ocs"`
}
