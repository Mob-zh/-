package repositories

import (
	"attendance_uniapp/global"
	"attendance_uniapp/models"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"log"
)

type AttendanceRepository struct {
	DB     *gorm.DB
	Client *redis.Client
}

func NewAttendanceRepository() *AttendanceRepository {
	return &AttendanceRepository{DB: global.DB, Client: global.RedisClient}
}

func (attendanceReno *AttendanceRepository) GetSignInCodeById(classId string) (string, error) {
	//从redis中取得签到码
	classKey := fmt.Sprintf("class:%v", classId)
	classInfo, err := attendanceReno.Client.HGetAll(global.RedisCtx, classKey).Result()
	if err != nil {
		return "", err
	}
	return classInfo["sign_code"], err
}

// StartToCheck 创建对应classId的hash
func (attendanceReno *AttendanceRepository) StartToCheck(classId string, endTime string, totalCount int, signInCode string, signInIndex int) error {

	return attendanceReno.Client.HSet(global.RedisCtx, fmt.Sprintf("class:%v", classId), map[string]interface{}{
		"end_time":    endTime,
		"sign_code":   signInCode,
		"sign_index":  signInIndex,
		"sign_count":  0,
		"total_count": totalCount,
	}).Err()
}

// StudentSignIn 学生满足签到条件进行签到
func (attendanceReno *AttendanceRepository) StudentSignIn(classId string, studentId string, signInIndex int) error {
	bitKey := fmt.Sprintf("class:student:%v:%v", classId, studentId)
	//SET HASH
	err := attendanceReno.Client.HIncrBy(global.RedisCtx, fmt.Sprintf("class:%v", classId), "sign_count", 1).Err()
	if err != nil {
		return err
	}
	//SET BITMAP
	return attendanceReno.Client.SetBit(global.RedisCtx, bitKey, int64(signInIndex-1), 1).Err()
}

// GetStudentSignInStatus 获取学生签到状态
func (attendanceReno *AttendanceRepository) GetStudentSignInStatus(classId string, studentId string, signInIndex int) bool {
	bitKey := fmt.Sprintf("class:student:%v:%v", classId, studentId)
	status := attendanceReno.Client.GetBit(global.RedisCtx, bitKey, int64(signInIndex)-1).Val()
	return status == 1
}

// GetClassHashByClassId 获取redis中对应班级hash表
func (attendanceReno *AttendanceRepository) GetClassHashByClassId(classId string) map[string]string {
	return attendanceReno.Client.HGetAll(global.RedisCtx, fmt.Sprintf("class:%v", classId)).Val()
}

// GetSignedInCountByClassId 以“{{已签到人数}}/{{班级总人数}}”的形式返回
func (attendanceReno *AttendanceRepository) GetSignedInCountByClassId(classId string) string {
	classHash := attendanceReno.Client.HGetAll(global.RedisCtx, fmt.Sprintf("class:%v", classId))
	signedInCount := classHash.Val()["sign_count"]
	totalCount := classHash.Val()["total_count"]
	return signedInCount + "/" + totalCount
}

//// FromRedisSyncClassAttendanceToMysql 将某一班级的某一次签到情况从redis中存入mysql
//func (attendanceReno *AttendanceRepository) FromRedisSyncClassAttendanceToMysql(classId string, signIndex int) error {
//	//模糊查询redis中该班级所有的学生签到信息
//	var cursor *redis.ScanIterator
//	pattern := fmt.Sprintf("class:student:%v:*", classId)
//	iter := attendanceReno.Client.Scan(context.Background(), 0, pattern, 0).Iterator()
//	for iter.Next(context.Background()) {
//		keys = append(keys, iter.Val()) // 获取当前扫描的键并保存
//	}
//
//}

//// BatchInsertAttendanceToMysqlWithTransaction 事务批量插入签到情况
//func (attendanceReno *AttendanceRepository) BatchInsertAttendanceToMysqlWithTransaction(attendances *[]models.AttendanceRecord) error {
//
//}

type Summary struct {
	StudentId   string
	StudentName string
	Statistics  string
}

// GetSignInSummaryByStudentInfoList 获取这个班级的当前签到汇总
func (attendanceReno *AttendanceRepository) GetSignInSummaryByStudentInfoList(studentInfoList []models.Student, classId string, totalCheckTimes int) ([]Summary, error) {
	//创建流水线
	pipe := attendanceReno.Client.Pipeline()
	var key string
	for _, student := range studentInfoList {
		key = fmt.Sprintf("class:student:%v:%v", classId, student.StudentId)
		// 将命令加入流水线，若没有对应的值，bitcount会返回0
		pipe.BitCount(global.RedisCtx, key, &redis.BitCount{Start: 0, End: int64(totalCheckTimes - 1)})
	}
	cmds, err := pipe.Exec(global.RedisCtx)
	if err != nil {
		return nil, fmt.Errorf("failed to execute pipeline: %v", err)
	}
	var summaryList []Summary
	// 处理返回的结果
	for i, cmd := range cmds {
		key = fmt.Sprintf("class:student:%v:%v", classId, studentInfoList[i].StudentId)

		// 获取签到次数
		count, err := cmd.(*redis.IntCmd).Result()

		if err != nil {
			log.Printf("Error retrieving attendance count for %s: %v", key, err)
			continue
		}
		summaryList = append(summaryList, Summary{
			StudentId:   studentInfoList[i].StudentId,
			StudentName: studentInfoList[i].StudentName,
			Statistics:  fmt.Sprintf("%d/%d", count, totalCheckTimes),
		})
	}

	return summaryList, nil
}
