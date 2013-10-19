package msgcookie

import (
    "testing"
    "net/http"
)

const (
    expected_session_id = "1276af897be84812a0553eb94d207849ca12fc35e0201dabcb61abd110b4215d"
    correct_secret = "special_secret"
    invalid_secret = "invalid_secret"
    session_name = "rack.session"
)

func TestDecode(t *testing.T) {
    val := "gqpzZXNzaW9uX2lk2gBAMTI3NmFmODk3YmU4NDgxMmEwNTUzZWI5NGQyMDc4%0ANDljYTEyZmMzNWUwMjAxZGFiY2I2MWFiZDExMGI0MjE1ZKZ2aXNpdHMC%0A--2689ddd14b0c839fabc9faa3ab83cd51e4f299f2"

    cookie := &http.Cookie{
        Name: session_name,
        Value: val,
    }

    iv, err := Decode(cookie, invalid_secret)
    if err == nil || err != http.ErrNoCookie {
        t.Fatal("invalid secret should return http.ErrNoCookie")
    }
    if iv != nil {
        t.Fatal("invalid cookie should return nil interface")
    }

    cv, err := Decode(cookie, correct_secret)
    if err != nil {
        t.Fatal(err)
    }

    var s []uint8
    var ok bool
    if s, ok = cv["session_id"].([]uint8); !ok {
        t.Fatal("session_id must be string")
    }
    if string(s) != expected_session_id {
        t.Fatal("asertion failed")
    }
}
