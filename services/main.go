package main

import "golang.org/x/crypto/bcrypt"

var text string = "Pawan@123"

func main() {

	// hasedPass, err := gernateHashedPassword("Pawan@123")
	// if err != nil {
	// 	println(err)
	// }
	hasedPass := "$2a$10$8GtFHTUh847t/sDdvVEy4.GSe1GW999/uw.ic2vtanZvze5CuWsWK"
	println(hasedPass)
	isValid := comparePasswords(hasedPass, []byte(text))
	if isValid {
		print("same Password")
	} else {
		print("worng Password")
	}
}
func gernateHashedPassword(password string) (pass string, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil

}
func comparePasswords(hashedPassword string, plainPassword []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), plainPassword)
	return err == nil
}
