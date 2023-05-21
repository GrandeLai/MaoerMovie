package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/stores/redis"

	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/scrypt"
)

const (
	KC_RAND_KIND_NUM   = 0 // 纯数字
	KC_RAND_KIND_LOWER = 1 // 小写字母
	KC_RAND_KIND_UPPER = 2 // 大写字母
	KC_RAND_KIND_ALL   = 3 // 数字、大小写字母
)

// 随机字符串
func Krand(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}

func GenerateVerificationCode() string {
	str, verificationCode := "0123456789", ""
	rand.Seed(time.Now().Unix())
	for i := 0; i < VerificationCodeLength; i++ {
		verificationCode += fmt.Sprintf("%c", str[rand.Intn(10)])
	}
	return verificationCode
}

func GenerateUUID() string {
	return uuid.NewV4().String()
}

func StringToInt64(str string) int64 {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return num
}

func Int64ToString(num int64) string {
	return string([]rune(strconv.FormatInt(num, 10)))
}

func IntToString(num int) string {
	return strconv.Itoa(num)
}

func Float64ToString(num float64) string {
	return strconv.FormatFloat(num, 'f', -1, 64)
}

func StringToFloat64(str string) float64 {
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return num
}

func StringToTime(str string) time.Time {
	timeT, _ := time.ParseInLocation("2006-01-02", str, time.Local)
	return timeT
}

func StringToTimeAtLocal(str string) time.Time {
	timeLayout := "2006-01-02T00:00:00+08:00"
	theTime, _ := time.ParseInLocation(timeLayout, str, time.Local)
	return theTime
}

func TimeToString(time time.Time) string {
	return time.Format("2006-01-02")
}

func GenerateJwtToken(secret string, expire int64, userId int64) (string, int64, error) {
	iat := time.Now().Unix()
	exp := iat + expire
	claims := make(jwt.MapClaims)
	claims["exp"] = exp
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", 0, err
	}
	return tokenString, exp, nil
}

func PasswordEncrypt(salt, password string) string {
	dk, _ := scrypt.Key([]byte(password), []byte(salt), 32768, 8, 1, 32)
	return fmt.Sprintf("%x", string(dk))
}

func GenerateNewId(redis *redis.Redis, keyPrefix string) int64 {
	//获取当前时间戳
	nowStamp := time.Now().Unix() - BeginTimeStamp
	//调用lua脚本，获取当天累计序列号
	nowDate := time.Now().Format("2006:01:02")
	newKeyString := "cache:icr:" + keyPrefix + ":" + nowDate
	//L := lua.NewState()
	//defer L.Close()
	//L.SetGlobal("getKeyString", L.NewFunction(func(L *lua.LState) int {
	//	L.Push(lua.LString(newKeyString))
	//	return 1
	//}))
	//if err := L.DoFile("common/scripts/generateIncrCount.lua"); err != nil {
	//	panic(err)
	//}
	//res := L.Get(1)
	//count, err := strconv.ParseInt(res.String(), 10, 64)
	//if err != nil {
	//	fmt.Println("调用lua脚本错误！")
	//	return 0
	//}
	count, err := redis.Incr(newKeyString)
	if err != nil {
		fmt.Println("生成id错误！")
		return 0
	}
	//拼接结果
	return nowStamp<<IdCountBit | count
}
