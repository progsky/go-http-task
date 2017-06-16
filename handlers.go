package main

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type md5HandlerInputJSON struct {
	ID   uint   `json:"id"   binding:"required"`
	Text string `json:"text" binding:"required"`
}

func md5Handler(c *gin.Context) {
	var params md5HandlerInputJSON
	if c.BindJSON(&params) == nil {
		if len(params.Text) > 100 {
			c.String(http.StatusBadRequest, "bad text")
			return
		}
	} else {
		c.String(http.StatusBadRequest, "bad json")
		return
	}
	md5sum := md5.Sum([]byte(strconv.Itoa(int(params.ID)) + params.Text + strconv.Itoa(int(params.ID%2))))
	c.String(http.StatusOK, hex.EncodeToString(md5sum[:]))
}
