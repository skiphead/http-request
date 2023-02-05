##How to use

import "github.com/SkipHead/http-request"

    tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: false,
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
	}
	r := request.Params{Ctx: context.Background(),
		Timeout:   10,
		Method:    http.MethodGet,
		URL:       "https://yandex.ru",
		Body:      nil,
		Transport: tr}
	fmt.Println(string(r.Request()))