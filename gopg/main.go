package main

func main(){
	db := Connect()
	if db == nil{
		panic("np connection")
	}
	CreateSchema(db, false, true)

/*	b := Book{Id:1}
	err := db.Insert(&b)
	if err!= nil{
		panic(err)
	}

	g := Genre{Id:1}
	err = db.Insert(&g)
	if err!= nil{
		panic(err)
	}*/

	bg := BookGenre{BookId:1, GenreId:1}
	err := db.Insert(&bg)
	if err!= nil{
		panic(err)
	}
}
