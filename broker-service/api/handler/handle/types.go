package handle

// RequestPayload describes the JSON that this service accepts as a request
type RequestPayload struct {
	Auth AuthPayload `json:"auth,omitempty"`
}

// AuthPayload is the embedded type (in RequestPayload) that describes an authentication request
type AuthPayload struct {
	Action   string `json:"action"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
