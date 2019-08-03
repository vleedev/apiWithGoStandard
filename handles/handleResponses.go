package handles
/*
*	The handle functions
*	Author: vlee.dev
 */
type ResponseResult struct {
	Message	string	`json:"message,omitempty"`
	SignInToken	string	`json:"signInToken,omitempty"`
}