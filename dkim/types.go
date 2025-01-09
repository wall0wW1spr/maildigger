package dkim

type DkimRunner struct {

    /* Nameserver target in order to perform the request */
    dnsServer string

    /* list of domains to hit */
    domains []string
}


type DkimDomains struct {

    /* related base domain */
    domain string

    /* DKIM key */
    key string

    /* algorithm TODO(djnn): algorithm enum */
    algorithm string
}
