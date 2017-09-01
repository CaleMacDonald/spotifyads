package main

import (
	"github.com/docopt/docopt-go"
	"github.com/calemacdonald/spotifyads"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	usage := `SpotifyAds - Spotify Ad Blocking.

	Usage:
	  spotifyads block
	  spotifyads add
	  spotifyads unblock
	  spotifyads remove
	  spotifyads -h | --help
	  spotifyads --version
	Options:
	  -h --help     Show this screen.
	  --version     Show the version.`

	args, _ := docopt.Parse(usage, nil, true, "SpotifyAds 1.0.0", false)

	hosts, err := spotifyads.NewHosts()
	check(err)
	
	if !hosts.IsWritable() {
		fmt.Fprintln(os.Stderr, "Hosts file is not writable. Try running with escalated privileges.")
		os.Exit(1)
	}

	const ip = "0.0.0.0"
	ads := []string {
		"pubads.g.doubleclick.net",
		"securepubads.g.doubleclick.net",
		"www.googletagservices.com",
		"gads.pubmatic.com",
		"ads.pubmatic.com",
		"spclient.wg.spotify.com",
	}

	if args["block"].(bool) || args["add"].(bool) {
		hosts.Add(ip, ads...)
		hosts.Flush()
	} 
	
	if args["unblock"].(bool) || args["remove"].(bool) {
		hosts.Remove(ip, ads...)
		hosts.Flush()
	}
}