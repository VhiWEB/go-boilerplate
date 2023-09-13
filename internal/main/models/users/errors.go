package users

type WrongUsernameOrPasswordError struct{}
type RegisterFailed struct{}

func (m *WrongUsernameOrPasswordError) Error() string {
	return "Wrong username or password"
}

func (m *RegisterFailed) Error() string {
	return "Error occured while registering your account"
}
