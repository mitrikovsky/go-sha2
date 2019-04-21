package api

import (
	"../core"
	"../db"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttprouter"
	"strconv"
)

type jobView struct {
	Id            int    `json:"id"`
	Payload       string `json:"payload"`
	HashRoundsCnt int    `json:"hash_rounds_cnt"`
	Status        string `json:"status"`
	Hash          string `json:"hash"`
}

func index(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	log.Debug("index")
	fmt.Fprint(ctx, "sha2 hasher api")
}

func postJob(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	post := jobView{}
	err := json.Unmarshal(ctx.Request.Body(), &post)
	if err != nil {
		log.Errorf("json.Unmarshal: %v", err)
		outputJsonMessageResult(ctx, 400, "json.Unmarshal: "+err.Error())
		return
	}

	// add job to db
	log.Debug("post job", post.Payload, post.HashRoundsCnt)
	id, err := db.Add(post.Payload, post.HashRoundsCnt)
	if err != nil {
		log.Errorf("db.Add: %v", err)
		outputJsonMessageResult(ctx, 500, "db.Add: "+err.Error())
		return
	}

	// increase WaitGroup count
	core.WG.Add(1)
	// start hash job in background
	go core.ProcessHash(id, post.Payload, post.HashRoundsCnt)

	outputJSON(ctx, 200, map[string]int{"id": id})
}

func getJob(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
	req := ps.ByName("id")
	id, err := strconv.Atoi(req)
	if err != nil {
		outputJsonMessageResult(ctx, 400, "wrong id "+req)
		return
	}

	// get job from db
	log.Debug("get job", id)
	job, err := db.Get(id)
	if err == sql.ErrNoRows {
		outputJsonMessageResult(ctx, 404, "job "+req+" not found")
		return
	} else if err != nil {
		outputJsonMessageResult(ctx, 500, "db.Get: "+err.Error())
		return
	}

	// show view for user
	view := jobView{}
	copier.Copy(&view, job)
	if job.Status == 0 {
		view.Status = "in progress"
	} else {
		view.Status = "finished"
	}
	outputJSON(ctx, 200, view)
}
