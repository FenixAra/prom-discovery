package aws

import (
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/elbv2"
)

type ECS struct {
	ecs   *ecs.ECS
	elbv2 *elbv2.ELBV2
	ec2   *ec2.EC2
}

func newECS() *ECS {
	return &ECS{
		ecs:   ecs.New(session.New()),
		elbv2: elbv2.New(session.New()),
		ec2:   ec2.New(session.New()),
	}
}

func (e *ECS) GetTargets(cname, name string, isPrivate bool) ([]string, error) {
	out, err := e.ecs.DescribeServices(&ecs.DescribeServicesInput{
		Cluster:  aws.String(cname),
		Services: []*string{aws.String(name)},
	})
	if err != nil {
		log.Println("Unable to describe services. Err: ", err)
		return nil, err
	}

	var targets []string
	for _, service := range out.Services {
		tgOut, err := e.elbv2.DescribeTargetHealth(&elbv2.DescribeTargetHealthInput{
			TargetGroupArn: service.LoadBalancers[0].TargetGroupArn,
		})
		if err != nil {
			log.Println("Unable to get target group's target health. Err: ", err)
			return nil, err
		}

		for _, target := range tgOut.TargetHealthDescriptions {
			instance, err := e.ec2.DescribeInstances(&ec2.DescribeInstancesInput{
				InstanceIds: []*string{target.Target.Id},
			})
			if err != nil {
				log.Println("Unable to get instance details. Err: ", err)
				return nil, err
			}

			if len(instance.Reservations) > 0 && len(instance.Reservations[0].Instances) > 0 {
				if isPrivate {
					targets = append(targets, *instance.Reservations[0].Instances[0].PrivateIpAddress+":"+strconv.Itoa(int(*target.Target.Port)))
				} else {
					targets = append(targets, *instance.Reservations[0].Instances[0].PublicIpAddress+":"+strconv.Itoa(int(*target.Target.Port)))
				}
			}
		}
	}

	return targets, nil
}
