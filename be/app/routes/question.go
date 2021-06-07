package routes

import (
	"database/sql"
	"fmt"
	"github.com/dchest/captcha"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12/context"
	"neko-question-box-be/app"
	"neko-question-box-be/app/utils"
	"net/http"
	"time"
)

type SubmitData struct {
	Id       string `json:"captchaId"`
	Value    string `json:"captchaValue"`
	Question string `json:"question"`
}

type ResponseData struct {
	Time       int    `json:"time"`
	Question   string `json:"question"`
	Answer     string `json:"answer"`
	AnswerTime int    `json:"answerTime"`
}

func SubmitQuestion(app *app.NekoQuestionBoxApp, ctx *context.Context) {
	if !utils.ShouldAcceptJson(ctx) {
		return
	}
	var question = &SubmitData{}
	err := ctx.ReadJSON(question)
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		utils.Errorf(err.Error())
		return
	}

	success := captcha.VerifyString(question.Id, question.Value)
	if !success {
		ctx.StatusCode(http.StatusNotAcceptable)
		app.CaptchaFailedTimes++
		utils.Warnf("验证失败，已累计 %d 次", app.CaptchaFailedTimes)
		return
	}
	app.CaptchaSuccessTimes++
	utils.Infof("验证成功，已累计 %d 次", app.CaptchaSuccessTimes)

	ctx.Values().Set("question", question.Question)

	SaveQuestion(app, ctx)
}

func SaveQuestion(app *app.NekoQuestionBoxApp, ctx *context.Context) {
	question := ""
	if ctx.Path() == "/save-question" {
		questionData := &SubmitData{}
		err := ctx.ReadBody(questionData)
		if err != nil {
			utils.Errorf(err.Error())
			return
		}
		question = questionData.Question
	} else {
		question = ctx.Values().Get("question").(string)
	}

	db, err := sql.Open("mysql", app.DataSource)
	if err != nil {
		utils.Errorf(err.Error())
		ctx.StatusCode(http.StatusInternalServerError)
		return
	}

	sqlStr := "insert into questions(time, question) values (%d, '%s')"

	_, err = db.Exec(fmt.Sprintf(sqlStr, time.Now().UnixNano()/1000000, question))
	if err != nil {
		utils.Errorf(err.Error())
		ctx.StatusCode(http.StatusInternalServerError)
		return
	}

	err = db.Close()
	if err != nil {
		utils.Errorf(err.Error())
		ctx.StatusCode(http.StatusInternalServerError)
		return
	}
}

func GetQuestion(app *app.NekoQuestionBoxApp, ctx *context.Context) {
	var response []ResponseData
	rowData := &struct {
		time       int
		question   string
		answer     sql.NullString
		answerTime sql.NullInt64
	}{
		time:       0,
		question:   "",
		answer:     sql.NullString{String: ""},
		answerTime: sql.NullInt64{Int64: 0},
	}
	utils.ShouldAcceptJson(ctx)
	db, err := sql.Open("mysql", app.DataSource)
	if err != nil {
		utils.Errorf(err.Error())
		ctx.StatusCode(http.StatusInternalServerError)
		return
	}
	rows, err := db.Query("select time, question, answer, answer_time from questions order by time desc")
	if err != nil {
		utils.Errorf(err.Error())
		ctx.StatusCode(http.StatusInternalServerError)
		return
	}
	for rows.Next() {
		err = rows.Scan(&rowData.time, &rowData.question, &rowData.answer, &rowData.answerTime)
		if err != nil {
			utils.Errorf(err.Error())
			ctx.StatusCode(http.StatusInternalServerError)
			return
		}
		response = append(response, ResponseData{
			Question:   rowData.question,
			Time:       rowData.time,
			Answer:     rowData.answer.String,
			AnswerTime: int(rowData.answerTime.Int64),
		})
	}
	_, err = ctx.JSON(response)
	if err != nil {
		utils.Errorf(err.Error())
		ctx.StatusCode(http.StatusInternalServerError)
		return
	}
}
