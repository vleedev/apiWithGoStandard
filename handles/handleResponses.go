package handles

type ResponseResult struct {
	Message	string	`json:"message,omitempty"`
	SignInToken	string	`json:"signInToken,omitempty"`
}