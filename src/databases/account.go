package databases

const (
	ClientScope 	= 1
	OfficeScope		= 1 << 1
	AdminScope 		= 1 << 2
)

type IsScope func(byte) bool

func IsClient(scopes byte) bool {
	return scopes & ClientScope == ClientScope
}

func IsOffice(scopes byte) bool {
	return scopes & OfficeScope == OfficeScope
}

func IsAdmin(scopes byte) bool {
	return scopes & AdminScope == AdminScope
}

func BuildScope(scopes ...byte) byte {
	var scope byte

	for _, val := range scopes {
		scope |= val
	}
	return scope
}