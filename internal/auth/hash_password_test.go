package auth

import "testing"

func TestHashPassword(t *testing.T) {
	password := "secretpassword"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Errorf("couldnt create hashword for password: %v, error: %s", password, err)
	}
	hashStatus, err := CheckPasswordHash(password, hashedPassword)
	if  !hashStatus || err != nil {
		t.Errorf("got wrong passwordhash")
	}

}
