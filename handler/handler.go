package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/rezairfanwijaya/go-autoscale.git/model"
	"github.com/rezairfanwijaya/go-autoscale.git/response"
	"github.com/sirupsen/logrus"
)

func Root(w http.ResponseWriter, r *http.Request) {
	logrus.Info("hit endpoint / ")
	resp := response.SuccessResp{
		Data:       "hi devops, go autocale will simulte the autoscaling with simple go app with current version v4",
		StatusCode: http.StatusOK,
	}

	respByte, err := resp.ChangToByte()
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(respByte)
	if err != nil {
		log.Printf("failed to create response, err: %s", err)
	}
}

func GetUserList(w http.ResponseWriter, r *http.Request) {
	logrus.Info("hit endpoint /users ")
	limitReq := r.URL.Query().Get("limit")
	if limitReq == "" {
		limitReq = "10"
	}

	limit, err := strconv.Atoi(limitReq)
	if err != nil {
		log.Printf("failed to create response getuserlist, err: %s", err)
		return
	}

	if limit <= 0 {
		limit = 10
	}

	users := getUsers(limit)

	resp := response.SuccessResp{
		Data:       users,
		StatusCode: http.StatusOK,
	}

	respByte, err := resp.ChangToByte()
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(respByte)
	if err != nil {
		log.Printf("failed to create response get userlist, err: %s", err)
	}
}

func getUsers(limit int) []model.User {
	users := []model.User{}

	for i := 1; i <= limit; i++ {
		user := model.User{
			ID:    i,
			Name:  fmt.Sprintf("example %d", i),
			Email: fmt.Sprintf("example%d@gmail.com", i),
		}
		users = append(users, user)
	}

	return users
}

func GetUserWithPanic(w http.ResponseWriter, r *http.Request) {
	logrus.Info("hit endpoint /users/panic")
	panic("panic")
}

func GetWorker(w http.ResponseWriter, r *http.Request) {
	logrus.Info("hit endpoint /workers ")
	workers := []model.Worker{
		{
			ID:   1,
			Name: "worker 1",
		},
		{
			ID:   2,
			Name: "worker 2",
		},
	}

	resp := response.SuccessResp{
		Data:       workers,
		StatusCode: http.StatusOK,
	}

	respByte, err := resp.ChangToByte()
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(respByte)
	if err != nil {
		log.Printf("failed to create response, err: %s", err)
	}
}
