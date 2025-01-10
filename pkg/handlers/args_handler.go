package handlers

import (
	"flag"
	"fmt"
	"os"
)
//Args_handler() handles the cli arguments (-h and -p)
func Args_handler() string {
	showHelp := flag.Bool("h", false, "use -h to see help message")
	serverPort := flag.Int("p", 8080, "use -p PORT to provide a specific port number for a web server")

	flag.Usage = func() {
		fmt.Println("Use -h for help")
		os.Exit(0)
	}

	flag.Parse()
	if len(os.Args) > 2{
		fmt.Println("ERROR: To many arguments. Use -h for help")
		os.Exit(0)
	} else if *showHelp {
		printHelp()
	}
	return ":" + fmt.Sprint(*serverPort)


}

func printHelp() {

	fmt.Println(`Usage: [program_name] [p = PORT]
	
	Options:
	  -h	 Show this help message and exit
	  -p	 Specifies the port for the server to run on. If -p is not used the server will run on the default 8080 port
	  `)
	os.Exit(0)
}
