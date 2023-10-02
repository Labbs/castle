// File: response.go
// Contains the responses of the bus messages
package bus

// GetUserByUsernameResponse is the response of GetUserByUsername
type GetUserByUsernameResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
