package user_info

type User struct {
	userID   string
	password string
	pID      int
	userName string
}

func NewUser(_userID string, _password string, _pID int, _userName string) *User {

	_user := User{
		userID:   _userID,
		password: _password,
		pID:      _pID,
		userName: _userName,
	}
	return &_user
}
