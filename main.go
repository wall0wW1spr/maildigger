package main

import (
	"github.com/spf13/cobra"

	"evil.djnn.sh/djnn/maildigger/dkim"

	"bufio"
	"fmt"
	"os"
)

const version = "0.0.1"
const ASCII_ART = `

     .-.
    /'v'\           maildigger
   (/   \)           ~djnn.sh
==='="="===<
    |_|              v0.0.1

                                s/o vsim<3
                        hack the planet,
                        travel the world . . .
------------------------------------------------
       DNS scrapping tool to recover DKIM
                   records

    ===> evil.djnn.sh/djnn/maildigger  <===
------------------------------------------------

`

var (
	domainsFilepath string
	nameserver      string

	maxLenDKIM int32
	maxLenSPF  int32
)

var rootCmd = &cobra.Command{
	Use:     "maildigger",
	Short:   "simple cli to scrape DKIM or SPF records",
	Long:    ASCII_ART,
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {
		println(ASCII_ART) /* why make hacking CLIs if you cant print silly ascii art ? */

		if domainsFilepath == "" {
			println("[!] error: please specify domain.")
			os.Exit(1)
		}

		file, err := os.Open(domainsFilepath)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		domains := make([]string, 1)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			domains = append(domains, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}
		fmt.Printf("[+] Loaded %d domains...\n", len(domains))

		dkimRunner := dkim.DkimRunner{
			DnsServer: nameserver,
			Domains:   domains,
		}

		err = dkimRunner.Run(maxLenDKIM)
		if err != nil {
			println("[!] Error while recovering DKIM records")
			panic(err)
		}

		println("[+] DKIM check done.")

	},
}

func main() {

	rootCmd.Flags().StringVarP(&nameserver, "nameserver", "n", "8.8.8.8", "DNS nameserver")
	rootCmd.Flags().StringVarP(&domainsFilepath, "domains", "d", "domains.txt", "file containing list of domains (line by line)")
	rootCmd.Flags().Int32VarP(&maxLenDKIM, "dkim-max-len", "", 128, "DKIM key max size")

	rootCmd.ExecuteC()
}
