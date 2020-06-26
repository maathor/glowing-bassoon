package ssher

import (
	"github.com/elliotchance/sshtunnel"
)

func ConnectAndRunCommands(ips []string) {
	tunnel := sshtunnel.NewSSHTunnel(
		"ec2-user@jumpbox.us-east-1.mydomain.com",
		sshtunnel.PrivateKeyFile("path/to/private/key.pem")
	)
}
