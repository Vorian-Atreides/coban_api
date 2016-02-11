package databases

const (
	// ClientScope provide an access right to the /clients
	ClientScope = 1
	// OfficeScope provide an access right to the /offices
	OfficeScope = 1 << 1
	// AdminScope provide an access right to the /administrations
	AdminScope = 1 << 2
)

// IsScope is a ptr on function for IsClient, IsOffice and IsAdmin
type IsScope func(byte) bool

// IsClient check if the scope has client access
func IsClient(scopes byte) bool {
	return scopes&ClientScope == ClientScope
}

// IsOffice check if the scope has office access
func IsOffice(scopes byte) bool {
	return scopes&OfficeScope == OfficeScope
}

// IsAdmin check if the scope has admin access
func IsAdmin(scopes byte) bool {
	return scopes&AdminScope == AdminScope
}

// BuildScope offers the ability to build a multi access right
func BuildScope(scopes ...byte) byte {
	var scope byte

	for _, val := range scopes {
		scope |= val
	}
	return scope
}
