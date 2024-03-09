package user_info

type User struct {
	UserID   string
	Password string
	PID      int
	UserName string
}

func NewUser(_userID string, _password string, _pID int, _userName string) *User {

	_user := User{
		UserID:   _userID,
		Password: _password,
		PID:      _pID,
		UserName: _userName,
	}
	return &_user
}
