package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"

	// import gin library
	"github.com/gin-gonic/gin"
	// import sample API packages

	"github.com/miraikeitai2020/backend-summer-vacation/pkg/db"
	"github.com/miraikeitai2020/backend-summer-vacation/pkg/server/model"
)

var (
	user     model.User
	datetime model.DateTime
	zeller   model.Zeller
	weeks    model.Weeks
	signUp   model.SignUp
	restoken model.Restoken
)

type Controller struct {
}

func (ctrl *Controller) HelloWorld(context *gin.Context) {
	context.JSON(200, gin.H{"message": "hello world"})
}

func (ctrl *Controller) SayHello(context *gin.Context) {
	err := context.BindJSON(&user)
	if err != nil {
		log.Println("[ERROR] Faild Bind JSON")
		context.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}
	context.JSON(200, gin.H{"message": "hello " + user.Name})
}

// 課題1
// 説明：
// 現在の日付と時間を返す.
// JSONの生成は gin.H を用いても良い
//
// リクエスト => なし
// レスポンス =>
// {
//   "timestamp": string,
//   "detail": {
//     "date": string, //例： 2020-09-02
//     "time": string, //例: 00:00:00
//   }
// }
func (ctrl *Controller) Task1(context *gin.Context) {
	nowTime := time.Now()
	const DateFormat = "2006-01-02"
	const TimeFormat = "15:04:05"

	datetime.Timestamp = strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	datetime.Details.Time = nowTime.Format(DateFormat)
	datetime.Details.Date = nowTime.Format(TimeFormat)
	context.JSON(200, datetime)
}

// 課題2
// 説明：
// ツェラーの公式でリクエストで投げた日付の曜日を返す
// JSONの生成は encoding/json を使用すること
//
// リクエスト =>
// {
//   "year": Int,
//   "month": Int,
//   "day": Int,
// }
// レスポンス =>
// {
//   "week": string //例： Monday
// }
func (ctrl *Controller) Task2(context *gin.Context) {
	err := context.BindJSON(&zeller)
	if err != nil {
		log.Println("[ERROR] Faild Bind JSON")
		context.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}
	if zeller.Month < 3 {
		zeller.Month += 12
		zeller.Year--
	}
	z := (zeller.Year + zeller.Year/4 - zeller.Year/100 + zeller.Year/400 + (13*zeller.Month+8)/5 + zeller.Day) % 7
	switch z {
	case 0:
		weeks.Week = "Sunday"
	case 1:
		weeks.Week = "Monday"
	case 2:
		weeks.Week = "Tuesday"
	case 3:
		weeks.Week = "Wednesday"
	case 4:
		weeks.Week = "Thursday"
	case 5:
		weeks.Week = "Fryday"
	case 6:
		weeks.Week = "Saturday"
	}
	jsonweek, err := json.MarshalIndent(weeks, "", "  ")
	if err != nil {
		context.JSON(500, gin.H{"JSON Marshal error": err})
		return
	}
	context.String(200, string(jsonweek))
}

// 課題3
// 説明：
// ユーザーIDとパスワードをデータベースに登録して, 発行したトークンを返す
// パスワードはハッシュ化したものをデータベースに登録する
// JSONの生成は encoding/json を使用すること
//
// リクエスト =>
// {
//   "id": string,
//   "password": string,
// }
// レスポンス =>
// {
//   "token": string
// }
func (ctrl *Controller) SignUp(context *gin.Context) {
	if err := context.BindJSON(&signUp); err != nil {
		context.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}

	hash := sha256.Sum256([]byte(signUp.Password))
	fmt.Println("ok")
	uuid, err := uuid.NewRandom()
	if err != nil {
		context.JSON(http.StatusInternalServerError, "AuthToken is error")
		return
	}
	restoken.Token = uuid.String()
	/*_, err = db.Con.Query("INSERT INTO `signUp` VALUES(?,?,?)", signUp.Id, hex.EncodeToString(hashed[:]), restoken.Token)
	if err != nil {
		context.JSON(http.StatusInternalServerError, "Internal Server Error")
	}*/
	if err := signUpData(signUp.Id, hex.EncodeToString(hash[:]), restoken.Token); err != nil {
		context.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	context.JSON(200, restoken)
}
func signUpData(id, Pass, token string) error {
	stmt, err := db.Con.Prepare("INSERT INTO signUp VALUES (?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id, Pass, token)
	return err
}

// 課題4
// 説明：
// ユーザーIDとパスワードをデータベースに登録されたものかを照合する
// 照合が終わったら結果を返す
// JSONの生成は encoding/json を使用すること
//
// リクエスト =>
// {
//   "id": string,
//   "password": string
// }
// レスポンス =>
// {
//   "certification": boolean
// }
func (ctrl *Controller) SignIn(context *gin.Context) {
}
