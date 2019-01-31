package handler

import (
	"context"
	"github.com/kgoval/valueid-membership/protogen"
)

type MemberSvc struct{}

func (g *MemberSvc) Detail(ctx context.Context, in *member.MemberDetailRequest, out *member.MemberDetailResponse) error{
	out.Phone = "019119"
	out.Email = "aaaaaa@kaskdas.com"
	out.FirstName = "hahahah"
	return nil
}