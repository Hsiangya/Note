package ugserver

import (
	"context"
	"growth/comm"
	"growth/models"
	"growth/pb"
	"growth/service"
	"log"
	"time"
)

type UgGradeServer struct {
	*pb.UnimplementedUserGradeServer
}

// ListGrades 获取所有的等级信息列表
func (s *UgGradeServer) ListGrades(ctx context.Context, in *pb.ListGradesRequest) (*pb.ListGradesReply, error) {
	log.Printf("UgGradeServer.ListGradesRequest=%+v\n", *in)
	//return nil, status.Errorf(codes.Unimplemented, "方法待实现")
	gradeInfoSvc := service.NewGradeInfoService(ctx)
	datalist, err := gradeInfoSvc.FindAll()
	if err != nil {
		return nil, err
	}
	dlist := make([]*pb.TbGradeInfo, len(datalist))
	for i := range datalist {
		dlist[i] = models.GradeInfoToMessage(&datalist[i])
	}
	out := &pb.ListGradesReply{
		Datalist: dlist,
	}
	return out, nil
}

// ListGradePrivileges 获取等级的特权列表
func (s *UgGradeServer) ListGradePrivileges(ctx context.Context, in *pb.ListGradePrivilegesRequest) (*pb.ListGradePrivilegesReply, error) {
	log.Printf("UgGradeServer.ListGradePrivilegesRequest=%+v\n", *in)
	//return nil, status.Errorf(codes.Unimplemented, "方法待实现")
	gradeId := int(in.GradeId)
	gradePrivilegeSvc := service.NewGradePrivilegeService(ctx)
	var datalist []models.TbGradePrivilege
	var err error
	if gradeId > 0 {
		datalist, err = gradePrivilegeSvc.FindByGrade(gradeId)
	} else {
		datalist, err = gradePrivilegeSvc.FindAll()
	}
	if err != nil {
		return nil, err
	}
	var dlist = make([]*pb.TbGradePrivilege, len(datalist))
	for i := range datalist {
		dlist[i] = models.GradePrivilegeToMessage(&datalist[i])
	}
	out := &pb.ListGradePrivilegesReply{
		Datalist: dlist,
	}
	return out, nil
}

// CheckUserPrivilege 检查用户是否有某个产品特权
func (s *UgGradeServer) CheckUserPrivilege(ctx context.Context, in *pb.CheckUserPrivilegeRequest) (*pb.CheckUserPrivilegeReply, error) {
	log.Printf("UgGradeServer.CheckUserPrivilegeRequest=%+v\n", *in)
	//return nil, status.Errorf(codes.Unimplemented, "方法待实现")
	uid := int(in.Uid)
	product := in.Product
	function := in.Function
	gradePrivilegeSvc := service.NewGradePrivilegeService(ctx)
	gradeUserSvc := service.NewGradeUserService(ctx)
	gradeUser, err := gradeUserSvc.GetByUid(uid)
	if err != nil {
		return nil, err
	}
	privilegeList, err := gradePrivilegeSvc.FindByGrade(gradeUser.GradeId)
	if err != nil {
		return nil, err
	}
	var isOk bool
	for _, privilege := range privilegeList {
		if privilege.Product == product && privilege.Function == function {
			isOk = true
			break
		}
	}
	out := &pb.CheckUserPrivilegeReply{
		Data: isOk,
	}
	return out, nil
}

// UserGradeInfo 获取用户的等级信息
func (s *UgGradeServer) UserGradeInfo(ctx context.Context, in *pb.UserGradeInfoRequest) (*pb.UserGradeInfoReply, error) {
	log.Printf("UgGradeServer.UserGradeInfoRequest=%+v\n", *in)
	//return nil, status.Errorf(codes.Unimplemented, "方法待实现")
	uid := int(in.Uid)
	gradeUserSvc := service.NewGradeUserService(ctx)
	gradeUser, err := gradeUserSvc.GetByUid(uid)
	if err != nil {
		return nil, err
	}
	data := models.GradeUserToMessage(gradeUser)
	out := &pb.UserGradeInfoReply{
		Data: data,
	}
	return out, nil
}

// UserGradeChange 调整用户的等级成长值
func (s *UgGradeServer) UserGradeChange(ctx context.Context, in *pb.UserGradeChangeRequest) (*pb.UserGradeChangeReply, error) {
	log.Printf("UgGradeServer.UserGradeChangeRequest=%+v\n", *in)
	//return nil, status.Errorf(codes.Unimplemented, "方法待实现")
	uid := int(in.Uid)
	score := int(in.Score)
	gradeUserSvc := service.NewGradeUserService(ctx)
	gradeInfoSvc := service.NewGradeInfoService(ctx)
	gradeUser, err := gradeUserSvc.GetByUid(uid)
	if err != nil {
		return nil, err
	}
	if gradeUser == nil {
		gradeUser = &models.TbGradeUser{
			Uid: uid,
		}
	}
	gradeUser.Score += score
	grade, err := gradeInfoSvc.NowGrade(gradeUser.Score)
	if err != nil {
		return nil, err
	}
	newData := models.TbGradeUser{
		Id:      gradeUser.Id,
		GradeId: 0,
		Expired: nil,
		Score:   gradeUser.Score,
	}
	if gradeUser.GradeId != grade.Id { // 等级更新
		newData.GradeId = grade.Id
		expired := comm.Now().AddDate(10, 0, 0)
		if grade.Expired > 0 {
			expired = comm.Now().Add(time.Hour * 24 * time.Duration(grade.Expired))
		}
		newData.Expired = &expired
	}
	err = gradeUserSvc.Save(&newData)
	if err != nil {
		return nil, err
	}
	out := &pb.UserGradeChangeReply{
		Data: models.GradeUserToMessage(&newData),
	}
	return out, nil
}
