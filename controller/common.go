// Package controller contains all controllers.
// Test cases in this package are integral tests.
package controller

import (
	"log"
	"strconv"
	"strings"

	"github.com/Patrolavia/mdpadgo/common"
	"github.com/Patrolavia/mdpadgo/model"
)

func Me(sess common.Session) (u *model.User, err error) {
	uidStr := sess.Get("uid")
	if err = sess.Err(); err != nil {
		log.Printf("Failed to read session data: %s", err)
		return
	}

	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		log.Printf("uid (%s) not integer: %s", uidStr, err)
		return
	}

	if u, err = model.LoadUser(uid); err != nil {
		log.Printf("Failed to load user#%d from db: %s", uid, err)
	}
	return
}

func PathArg(url, base string) (args []string) {
	if url[:len(base)] != base {
		return
	}
	url = url[len(base):]
	return strings.Split(url, "/")
}

type CreateRequest struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Coops   []int    `json:"cooperator"`
	Tags    []string `json:"tags"`
}

type EditRequest struct {
	*CreateRequest
	Version int `json:"version"`
}