package model

//Status type definition
type Status int

const (
	//SENT : in case email is sent
	SENT Status = iota
	//FAILED : in case email is sent
	FAILED
	//SCHEDULED : in case email is sent
	SCHEDULED
)

func (status Status) String() string {
	return [...]string{"SENT", "FAILED", "SCHEDULED"}[status]
}
