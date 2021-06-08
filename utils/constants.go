package utils

// Golang can't have a constant string slice

func GetWebService() []string {
	// source: https://github.com/nmap/nmap/blob/4b46fa7097673f157e7b93e72f0c8b3249c54b4c/nselib/shortport.lua#L179
	return []string{"http", "https", "ipp", "http-alt", "https-alt", "vnc-http", "oem-agent",
		"soap", "http-proxy", "caldav", "carddav", "webdav"}
}
