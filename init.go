package onebox

func Init(host string, bearer string) Context {
	return &context{
		bearer: bearer,
		host:   host,
	}
}
