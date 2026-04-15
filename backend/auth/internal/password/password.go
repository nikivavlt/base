package password

import (
    "fmt"

    "golang.org/x/crypto/bcrypt"
)

const cost = 12 // bcrypt work factor — higher = slower = safer
// cost=12 takes ~250ms per hash — fast enough for users, 
// slow enough to make brute-force attacks impractical.

func Hash(plain string) (string, error) {
    b, err := bcrypt.GenerateFromPassword([]byte(plain), cost)
    if err != nil {
        return "", fmt.Errorf("hashing password: %w", err)
    }
    return string(b), nil
}

func Verify(plain, hashed string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
    return err == nil
}