package awsec2

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	parseconf "../parseconf"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func GetPrivateIpFromTag(tag string, env *parseconf.Env) []string {
	toReturn := []string{}
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(env.AwsRegion)})
	ec2svc := ec2.New(sess)
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String(env.TagFilter),
				Values: []*string{aws.String(tag)},
			},
			{
				Name:   aws.String("instance-state-name"),
				Values: []*string{aws.String("running"), aws.String("pending")},
			},
		},
	}
	resp, err := ec2svc.DescribeInstances(params)
	if err != nil {
		log.Fatal("there was an error listing instances in", err.Error())
	}

	for idx, res := range resp.Reservations {
		fmt.Println("  > Reservation Id", *res.ReservationId, " Num Instances: ", len(res.Instances))
		for _, inst := range resp.Reservations[idx].Instances {
			fmt.Println("    - Instance ID: ", *inst.PrivateIpAddress)
			toReturn = append(toReturn, *inst.PrivateIpAddress)
		}
	}
	return toReturn
}
