package services

import (
	"attendance_uniapp/repositories"
	"attendance_uniapp/utils"
	"fmt"
	"strconv"
	"time"
)

type AttendanceService struct {
	AttendanceRepo *repositories.AttendanceRepository
	ClassReno      *repositories.ClassRepository
}

func NewAttendanceService(AttendanceRepo *repositories.AttendanceRepository, ClassReno *repositories.ClassRepository) *AttendanceService {
	return &AttendanceService{AttendanceRepo: AttendanceRepo, ClassReno: ClassReno}
}

// StartToCheckAndGetSignInCodeService 教师开始考勤服务,初始化相应数据
func (attendanceServ *AttendanceService) StartToCheckAndGetSignInCodeService(classId string, checkingSeconds int) (string, error) {
	endTime := time.Now().Add(time.Duration(checkingSeconds) * time.Second).Format("2006-01-02 15:04:05")
	totalCount := attendanceServ.ClassReno.GetStudentCountByClassId(classId)
	class, _ := attendanceServ.ClassReno.GetClassById(classId)
	signInCode := utils.GenerateSignInCode()
	class.TotalCheckingTimes++
	return signInCode, attendanceServ.AttendanceRepo.StartToCheck(classId, endTime, totalCount, signInCode, class.TotalCheckingTimes)
}

// GetSignedInCountService 获取当前签到成功人数服务，以“{{已签到人数}}/{{班级总人数}}”的形式返回
func (attendanceServ *AttendanceService) GetSignedInCountService(classId string) (string, error) {
	signedInCount := attendanceServ.AttendanceRepo.GetSignedInCountByClassId(classId)
	if signedInCount != "" {
		return signedInCount, nil
	}
	return "", fmt.Errorf("redis error")
}

// StudentSignInService 学生签到服务
func (attendanceServ *AttendanceService) StudentSignInService(studentId string, signInCode string, classId string) error {
	classHash := attendanceServ.AttendanceRepo.GetClassHashByClassId(classId)
	//获取考勤索引
	signInIndex, _ := strconv.Atoi(classHash["sign_index"])
	//验证签到条件
	correctCode := classHash["sign_code"]
	endTime, err := time.Parse("2006-01-02 15:04:05", classHash["end_time"])
	if err != nil {
		return err
	}

	if correctCode != signInCode || time.Now().After(endTime) {
		return fmt.Errorf("signInCode is incorrect or check-in has ended")
	}
	//验证是否已签到
	status := attendanceServ.AttendanceRepo.GetStudentSignInStatus(studentId, classId, signInIndex)
	//若未签则签
	if !status {
		err := attendanceServ.AttendanceRepo.StudentSignIn(studentId, classId, signInIndex)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetSignInSummaryService 教师获取班级当前签到汇总
func (attendanceServ *AttendanceService) GetSignInSummaryService(classId string) (string, error) {
	studentInfoList, _ := attendanceServ.ClassReno.GetStudentInfoListByClassId(classId)
	classHash := attendanceServ.AttendanceRepo.GetClassHashByClassId(classId)
	signInIndex, _ := strconv.Atoi(classHash["sign_index"])

	var studentIds []string
	for _, student := range studentInfoList {
		studentIds = append(studentIds, student.StudentId)
	}

	summary, err := attendanceServ.AttendanceRepo.GetSignInSummaryByStudentInfoList(studentInfoList, classId, signInIndex)
	if err != nil {
		return "", err
	}
	return utils.SummaryToComponents(summary), nil
}
