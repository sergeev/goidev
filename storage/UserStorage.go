package storage

func UserIsValid(uName, pwd string) bool {
	// DB simulation
	_uName, _pwd, _isValid := "m", "1", false

	if uName == _uName && pwd == _pwd {
		_isValid = true
	} else {
		_isValid = false
	}

	return _isValid
}