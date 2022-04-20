package impl

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/wuyadong1990/grpc-demo-proto/user"

	"github.com/xiaomeng79/go-log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	user.UnimplementedUserServiceServer
}

func (s *Server) UserInfo(ctx context.Context, in *user.UserBase) (out *user.UserBase, outerr error) {

	//log.Printf("[User] Create Req: %d", in.GetId())
	/*创建接口*/
	m := new(User)
	out = new(user.UserBase)
	err := copier.Copy(m, in)
	m.ID = in.GetId()
	//fmt.Printf("GormDB=%#v\n", cinit.GormDB)
	//fmt.Printf("Mysql=%#v\n", cinit.Mysql)
	log.Debugf("[User] Create Req in: %v", in)
	log.Debugf("[User] Create Req m: %v", m)
	if err != nil {
		//log.Printf(err.Error(), ctx)
		outerr = status.Error(codes.Internal, err.Error())
		return
	}
	if m.ID > 0 {
		log.Debugf("call update method")
		err = m.Update(ctx)
	} else {
		log.Debugf("call add method")
		err = m.Add(ctx)
	}

	if err != nil {
		outerr = status.Error(codes.InvalidArgument, err.Error())
		return
	}
	err = copier.Copy(out, m)
	if err != nil {
		//log.Error(err.Error(), ctx)
		outerr = status.Error(codes.Internal, err.Error())
		return
	}
	return
}

func (s *Server) UserQueryOne(ctx context.Context, in *user.UserID) (out *user.UserBase, outerr error) {
	m := new(User)
	out = new(user.UserBase)
	m.ID = in.GetId()
	var err error
	/*err := copier.Copy(m, in)
		if err != nil {
			log.Error(err.Error(), ctx)
			outerr = status.Error(codes.Internal, err.Error())
			return
	}*/
	log.Debugf("[UserQueryOne] in: %v", in)
	log.Debugf("[UserQueryOne] m: %v", m)

	err = m.QueryOne(ctx)
	if err != nil {
		outerr = status.Error(codes.InvalidArgument, err.Error())
		return
	}
	err = copier.Copy(out, m)
	if err != nil {
		log.Error(err.Error(), ctx)
		outerr = status.Error(codes.Internal, err.Error())
		return
	}
	return
}

func (s *Server) UserDelete(ctx context.Context, in *user.UserID) (out *user.UserID, outerr error) {
	m := new(User)
	out = new(user.UserID)
	m.ID = in.GetId()
	var err error

	/*err := copier.Copy(m, in)
	if err != nil {
		log.Error(err.Error(), ctx)
		outerr = status.Error(codes.Internal, err.Error())
		return
	}*/
	err = m.Delete(ctx)
	if err != nil {
		outerr = status.Error(codes.InvalidArgument, err.Error())
		return
	}
	err = copier.Copy(out, m)
	if err != nil {
		log.Error(err.Error(), ctx)
		outerr = status.Error(codes.Internal, err.Error())
		return
	}
	return
}

func (s *Server) UserQueryAll(ctx context.Context, in *user.UserAllOption) (*user.UserAll, error) {
	m := new(User)
	err := copier.Copy(m, in)
	if err != nil {
		log.Error(err.Error(), ctx)
		return &user.UserAll{}, status.Error(codes.Internal, err.Error())
	}
	//log.Debugf("[UserQueryAll] Create Req in: %v", in)
	//log.Debugf("[UserQueryAll] Create Req m: %v", m)
	ms, page, err := m.QueryAll(ctx)
	if err != nil {
		return &user.UserAll{}, status.Error(codes.InvalidArgument, err.Error())
	}
	var agt []*user.UserBase
	err = copier.Copy(&agt, ms)
	if err != nil {
		log.Error(err.Error(), ctx)
		return &user.UserAll{}, status.Error(codes.Internal, err.Error())
	}
	_page := new(user.Page)
	err = copier.Copy(_page, &page)
	if err != nil {
		log.Error(err.Error(), ctx)
		return &user.UserAll{}, status.Error(codes.Internal, err.Error())
	}
	return &user.UserAll{All: agt, Page: _page}, nil
}
