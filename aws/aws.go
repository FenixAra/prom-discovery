package aws

const (
	TypeECS = "ecs"
	TypeASG = "asg"
	TypeEC2 = "ec2"
)

type AWS interface {
	GetTargets(cname, name string) ([]string, error)
}

func New(t string) AWS {
	switch t {
	case TypeECS:
		return newECS()
	case TypeASG:
		return nil
	case TypeEC2:
		return nil
	default:
		return nil
	}
}
