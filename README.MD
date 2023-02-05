##How to use

    tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: false,
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
	}

	req := Params{
		Ctx:       context.Background(),
		Timeout:   10,
		Method:    http.MethodGet,
		URL:       "https://example.com",
		Body:      nil,
		Transport: tr,
	}
	fmt.Println(string(req.Request()))