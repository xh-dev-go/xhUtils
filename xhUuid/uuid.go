package xhUuid

import "github.com/google/uuid"

func UUIDString() (string,error) {
	if uuid, err :=uuid.NewUUID(); err!= nil{
		return "",err
	} else {
		return uuid.String(), nil
	}
}
