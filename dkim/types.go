package dkim

type DkimRunner struct {

	/* Nameserver target in order to perform the request */
	DnsServer string

	/* list of domains to hit */
	Domains []string
}

type DkimDomain struct {

	/* related base domain */
	domain string

	/* DKIM key */
	key string

	/* DKIM key len */
	keyLen int32
}
