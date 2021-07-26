package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func HandleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}

		result := strconv.Itoa(random()) + "\n"
		c.Write([]byte(string(result)))
	}
	c.Close()

}

// ChannelEndpointChanged()
// SetGroupChannel()
// CloseChannel()
// SetExtranoncePrefix()
// NewMiningJob()
// NewExtendedMiningJob()
// SetNewPrevHash()
// SetTarget()
// Reconnect()
// SetGroupChannel()
