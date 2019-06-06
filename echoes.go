package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/user"
	"time"
)

// Constants
const (
	connectionType = "tcp"
)

// Map all the connection.
var connections map[string]net.Conn
var currentUser, _ = user.Current()

func main() {

	lChannel := flag.String("lChannel", "server", "Left side connection model")
	rChannel := flag.String("rChannel", "server", "Right side connection model")

	lChannelPort := flag.String("lChannelPort", "25000", "Left side connection model")
	rChannelPort := flag.String("rChannelPort", "25001", "Right side connection model")

	lChannellAddress := flag.String("lChannellAddress", "192.168.200.10", "Left side ip address")
	rChannellAddress := flag.String("rChannellAddress", "192.168.200.20", "Right side ip address")
	flag.Parse()

	fmt.Println("lChannel", *lChannel)
	fmt.Println("rChannel", *rChannel)
	fmt.Println("lChannelPort", *lChannelPort)
	fmt.Println("rChannelPort", *rChannelPort)
	fmt.Println("lChannellAddress", *lChannellAddress)
	fmt.Println("rChannellAddress", *rChannellAddress)


	// Creates log file.
	logfile, err := os.Create(currentUser.HomeDir + "/server_" + time.Now().Format("2006-01-02") + ".log")

	if err != nil {
		fmt.Println(timestamp(), "\nAn error occurred creating the log file")
		fmt.Println(err)
		return
	}

	defer logfile.Close()

	// Create connections
	// Attempt connections
	connections = make(map[string]net.Conn)

	if err != nil {
		log(logfile, "An error occurred during server startup")
		log(logfile, err.Error())
		return
	}

	if *lChannel == "server" {
		// if the left channel is a server type, launch the go routine for creating connections
	} else {

	}

	if *rChannel == "server"{
		// if the right channel is a server type, launch the go routine for creating connections
	}


}

func createConnection(connType string, connPort string, connAddress string) {
	// this functions must act as a routine that launches multiples subroutines.
}

func handleConnection(conn net.Conn, logfile *os.File) {
	log(logfile, "Handling new connection made by: "+conn.RemoteAddr().String())
	// Close connection when this function ends
	defer func() {
		log(logfile, "Closing connection made by: "+conn.RemoteAddr().String())

		// Close the connection and remove ot from the map.
		conn.Close()
		delete(connections, conn.RemoteAddr().String())
	}()

	data_stream := make([]byte, 1024)

	for {

		// Read the data
		len, err := conn.Read(data_stream)
		if err != nil {
			if err.Error() != ("EOF") {
				log(logfile, "An error occurred reading buffer "+err.Error())
			}
			return
		}

		if len > 0 {
			log(logfile, conn.RemoteAddr().String(), data_stream[:len])

			for _, c := range connections {
				// send data to every single connection made but the sender.
				if c != conn {
					c.Write([]byte(data_stream[:len]))
				}

			}
		}

	}
}

// Transform special characters to be seen in the log file.
func byteToString(c []byte) string {
	str := ""

	for _, value := range c {
		switch value {
		case 4: //eot
			str += "<EOT>"
		case 5: //ENQ
			str += "<ENQ>"
		case 6: //ACK
			str += "<ACK>"
		case 10:
			str += "<LF>"
		case 13:
			str += "<CR>"
		case 21: //ENQ
			str += "<NAK>"
		default:
			str += string(value)
		}
	}

	return str
}

// Set timestamp format.
func timestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// log data into a file.
func log(logfile *os.File, text string, data_stream ...[]byte) {
	if data_stream != nil {
		parsedData := byteToString(data_stream[0])
		logfile.WriteString(fmt.Sprintf("%[1]s %[2]s - %[3]s %[4]s", timestamp(), text, parsedData, "\n"))
		//logfile.Write(bytes[0])
		fmt.Printf("%[1]s %[2]s - %[3]s %[4]s", timestamp(), text, parsedData, "\n")
	} else {
		logfile.WriteString(timestamp() + " " + text + "\r\n")
		fmt.Println(timestamp(), text, "\r")

	}
}
