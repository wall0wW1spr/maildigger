package main

import (
    "github.com/spf13/cobra"
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
               and/or SPF records

    ===> evil.djnn.sh/djnn/maildigger  <===
------------------------------------------------

`

var (
    outfilePath     string
    domainsFilepath string
    nameserver      string

    maxLenDKIM      int32
    maxLenSPF       int32
)

var rootCmd = &cobra.Command{
	Use:   "maildigger",
	Short: "simple cli to scrape DKIM or SPF records",
	Long:  ASCII_ART,
	Run: func(cmd *cobra.Command, args []string) {
        println(ASCII_ART) /* why make hacking CLIs if you cant print silly ascii art ? */


    },
}


func main() {

    rootCmd.Flags().StringVarP(&outfilePath, "outfile", "o", "", "outfile path (stdout if empty)")
    rootCmd.Flags().StringVarP(&nameserver, "nameserver", "n", "8.8.8.8", "DNS nameserver")
    rootCmd.Flags().StringVarP(&domainsFilepath, "domains", "d", "domains.txt", "file containing list of domains (line by line)")
    rootCmd.Flags().UintVarP(&maxLenDKIM, "dkim-max-len", "", 512, "DKIM key max size")
    rootCmd.Flags().UintVarP(&maxLenDKIM, "spf-max-len", "", 0, "SPF key max size")

    rootCmd.ExecuteC()
}
