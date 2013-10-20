package msgcookie

import (
    "github.com/ugorji/go/codec"
    "encoding/base64"
    "net/url"
    "net/http"
    "strings"
    "github.com/yasuoza/msgcookie/utils"
)

var (
    mh codec.MsgpackHandle
)

func Decode(cookie *http.Cookie, secret string) (map[string]interface{}, error) {
    var v map[string]interface{}
    str, err := url.QueryUnescape(cookie.Value)
    if err != nil {
        return nil, err
    }

    session_data, digest := splitSessionDataAndDigest(str)
    if secret != "" {
        hmac, err := openssl.GenerateHmac(session_data, secret);
        if err != nil || digest != hmac {
            return nil, http.ErrNoCookie
        }
    }

    data, err := base64.StdEncoding.DecodeString(session_data)
    dec := codec.NewDecoderBytes(data, &mh)
    err = dec.Decode(&v)
    if err != nil {
        return nil, err
    }
    return v, nil
}

func splitSessionDataAndDigest(str string) (string, string) {
    s := strings.SplitN(str, "--", 2)[:2]
    return s[0], s[1]
}
