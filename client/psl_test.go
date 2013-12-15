package client

type PSL struct{}

func (PSL) String() string {
	return "mock psl"
}

func (PSL) PublicSuffix(domain string) string {
	return domain
}
