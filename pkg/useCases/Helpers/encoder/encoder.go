package encoder

import "golang.org/x/crypto/bcrypt"

type Encoder interface {
	HashAndSalt(password string) ([]byte, error)
	ComparePasswords(hashedPwd string, plainPwd string) error
}

type EncoderImpl struct {
}

func (e *EncoderImpl) HashAndSalt(password string) ([]byte, error) {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.MinCost)
	if err != nil {
		return []byte{}, err
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return hash, nil
}

func (e *EncoderImpl) ComparePasswords(hashedPwd string, plainPwd string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	if err != nil {
		return err
	}
	return nil
}
