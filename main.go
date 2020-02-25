package main

import (
	"fmt"

	"https://github.com/rsmnarts/port-scanner/port"

	"github.com/superhawk610/bar"
	"github.com/ttacon/chalk"
)

var listPort = make(map[string][]int)

func portScanned(protocol string, i int) {
	if ok := port.ScanPort(protocol, "localhost", i); ok {
		listPort[protocol] = append(listPort[protocol], i)
	}
}

func main() {
	var (
		n = 65535
		// n = 1000
		b = bar.NewWithOpts(
			bar.WithDimensions(n, 50),
			bar.WithFormat(fmt.Sprintf(" %sloading...%s :percent :bar %s:rate ops/s%s ", chalk.Yellow, chalk.Reset, chalk.Green, chalk.Reset)),
		)
	)

	fmt.Printf("\n%sPort Scanner in Go%s\n\n", chalk.Red, chalk.Reset)

	for i := 0; i < n; i++ {
		go portScanned("tcp", i)
		// go portScanned("udp", i)
		b.Tick()
	}

	b.Done()

	if listPort != nil {
		fmt.Println("List open ports: ", listPort)
	}
}
