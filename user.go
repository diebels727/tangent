package tangent

// User represents information about a user on IRC
type User struct {
	Nick, Ident, Host string // Host information
	Real              string // Realname(GECOS)
}
