package index

import (
	db "xianyun/Databases"
)

type Photo struct {
	Id  int
	Url string
}

func (photo *Photo)GetOne()  (p *Photo ,err error) {
	res := db.SqlDB.QueryRow(`select id,url FROM td_index_url where id = 1`)
	var data Photo
	err = res.Scan(&data.Id, &data.Url)
	if err == nil{
		return &data,nil
	}
	return &data,err
}

func (photo *Photo)GetAll() [] Photo  {
	var data Photo
	var p []Photo
	rows, err := db.SqlDB.Query("select id,url from  td_index_url")
	if err != nil{
		return nil
	}
	defer rows.Close()
	for rows.Next(){
		err = rows.Scan(&data.Id, &data.Url)
		p = append(p, data)
		//fmt.Println(p)
	}
	return p

}