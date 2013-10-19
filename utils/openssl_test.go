package openssl

import (
    "testing"
)

const (
    expected_digest = "2689ddd14b0c839fabc9faa3ab83cd51e4f299f2"
)

func TestGenerateHmac(t *testing.T) {
    session_val := "gqpzZXNzaW9uX2lk2gBAMTI3NmFmODk3YmU4NDgxMmEwNTUzZWI5NGQyMDc4\nNDljYTEyZmMzNWUwMjAxZGFiY2I2MWFiZDExMGI0MjE1ZKZ2aXNpdHMC\n"
    secret := "special_secret"

    digest, err := GenerateHmac(session_val, secret)
    if err != nil  {
        t.Fatal(err)
    }

    if digest != expected_digest {
        t.Fatal("generated digest does not match")
    }
}
