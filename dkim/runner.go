package dkim

import (
	"github.com/miekg/dns"

	"fmt"
	"strings"
)

func (dk *DkimRunner) Run(maxLen int32) error {
	dkimPrefix := "k=rsa; p="

	m := new(dns.Client)
	for _, domain := range dk.Domains {
		q := dns.Question{
			Name:   domain + ".",
			Qtype:  dns.TypeTXT,
			Qclass: dns.ClassINET,
		}

		m1 := new(dns.Msg)
		m1.Id = dns.Id()
		m1.RecursionDesired = true
		m1.Question = make([]dns.Question, 1)
		m1.Question[0] = q

		/* make DNS request */
		msg, _, err := m.Exchange(m1, dk.DnsServer+":53")
		if err != nil {
			println("[+] Could not contact exchange server")
			return err
		}

		/* get TXT records from Answer */
		for _, rr := range msg.Answer {
			if rr.Header().Rrtype == dns.TypeTXT {
				txt := rr.(*dns.TXT).Txt

				for _, t := range txt {

					/* we have found something that looks like a DKIM record */
					if strings.HasPrefix(t, dkimPrefix) {

						key := strings.TrimPrefix(t, dkimPrefix)
						kLen := len(key)

						/* key is bigger than our max-size, we discard */
						if int32(kLen) > maxLen {
							continue
						}

						fmt.Printf("[+] => %s: %s (len: %d)\n", domain, key, kLen)
					}
				}
			}
		}
	}

	return nil
}
