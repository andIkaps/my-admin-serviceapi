package security

import (
	"errors"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	hash := string(hashPassword)

	return hash, nil
}

func CheckPassword(hashed, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}

func WeaknessPassword(password string) (bool, error) {
	if len(password) < 8 {
		return false, errors.New("password harus terdiri dari minimal 8 karakter")
	}

	// Pengecekan huruf besar
	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return false, errors.New("password harus mengandung setidaknya satu huruf besar")
	}

	// Pengecekan huruf kecil
	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return false, errors.New("password harus mengandung setidaknya satu huruf kecil")
	}

	// Pengecekan angka
	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return false, errors.New("password harus mengandung setidaknya satu angka")
	}

	symbols := "!@#$%^&*()_+{}[]:;<>,.?/~`"
	hasSymbol := false
	for _, symbol := range symbols {
		if strings.ContainsRune(password, symbol) {
			hasSymbol = true
			break
		}
	}
	if !hasSymbol {
		return false, errors.New("password harus mengandung setidaknya satu simbol")
	}

	return true, nil
}
