
	package main
	//
	func Create(){
	query := "INSERT INTO user(password,photo,latitude,longitude)VALUES(?,?,?,?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err, r.log

	}
	_, err = stmt.Exec(user.Password,user.Photo,user.Latitude,user.Longitude)
	if err != nil {
		r.log.Message += "|Exec|" + err.Error()
		return err, r.log
	}
	
	}