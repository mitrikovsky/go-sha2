package db

type Job struct {
	Id            int
	Payload       string
	HashRoundsCnt int
	Status        int
	Hash          string
}

func Get(id int) (j Job, err error) {
	row := db.QueryRow("SELECT * FROM hash_data.jobs WHERE id = $1", id)
	err = row.Scan(&j.Id, &j.Payload, &j.HashRoundsCnt, &j.Status, &j.Hash)
	return
}

func Add(pl string, cnt int) (id int, err error) {
	err = db.QueryRow("INSERT INTO hash_data.jobs (payload, hash_rounds_cnt) VALUES($1, $2) RETURNING id", pl, cnt).Scan(&id)
	return
}

func Update(id int, hash string) (err error) {
	_, err = db.Exec("UPDATE hash_data.jobs SET hash = $1, status = 1 WHERE id = $2", hash, id)
	return
}
