package openssl

import (
    "bytes"
    "os/exec"
    "strings"
)

func init() {
    c := exec.Command("type", "openssl")
    err := c.Run()
    if err != nil {
        panic(exec.ErrNotFound)
    }
}

func GenerateHmac(data, secret string) (string, error) {
    c1 := exec.Command("echo", "-n", data)
    c2 := exec.Command("openssl",  "dgst", "-sha1", "-hmac", secret)

    var out bytes.Buffer
    c2.Stdin, _ = c1.StdoutPipe()
    c2.Stdout = &out
    var err error
    err = c2.Start()
    if err != nil {
        return "", err
    }
    err = c1.Run()
    if err != nil {
        return "", err
    }
    err = c2.Wait()
    if err != nil {
        return "", err
    }
    return strings.TrimRight(out.String(), "\n"), nil
}
