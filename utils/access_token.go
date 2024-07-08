package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func RequestToken(appid, secret, code string) (string, string, error) {
	u, err := url.Parse("https://api.weixin.qq.com/sns/jscode2session")

	if err != nil {
		log.Fatal(err)
	}

	parse := &url.Values{}
	parse.Set("appid", appid)
	parse.Set("secret", secret)
	parse.Set("js_code", code)
	parse.Set("grant_type", "authorization_code")

	u.RawQuery = parse.Encode()

	resp, err := http.Get(u.String())

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return "", "", errors.New("request token error :" + err.Error())
	}

	jMap := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&jMap)
	fmt.Sprintln("jmap uis :", jMap)
	if err != nil {
		return "", "", errors.New("request token response json parse err :" + err.Error())
	}
	if jMap["errcode"] == nil || jMap["errcode"] == 0 {

		openid, _ := jMap["openid"].(string)
		session_key, _ := jMap["session_key"].(string)
		return openid, session_key, nil
	} else {
		// 返回错误信息
		// errcode := string(jMap["errcode"].(float64))
		errmsg := jMap["errmsg"].(string)
		err = errors.New(errmsg)
		return "", "", err
	}

}
