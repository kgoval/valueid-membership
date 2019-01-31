package handler

import (
	member "github.com/kgoval/valueid-membership/protogen"
	"github.com/micro/go-micro/server"
)

// register handler
func Register(s server.Server){
	member.RegisterMemberServiceHandler(s, new(MemberSvc))
}

