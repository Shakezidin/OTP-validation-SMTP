package model

type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type Verify struct {
	Email string `json:"email"`
	OTP   int    `json:"otp"`
}
