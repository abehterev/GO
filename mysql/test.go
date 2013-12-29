package main;

import(
	"errors"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type Login interface {
	GetByLogin(l string) (bool, error)
}

type TAuth struct {
    id int
    login string
    md5 string
}

func (a *TAuth) GetByLogin(l string) (bool, error){

	db, err := sql.Open("mysql", "musicbox:musicboxpass@/musicbox")
	if err != nil {
		return false, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return false, err
	}

	rows, err := db.Query("SELECT id,md5pass FROM users WHERE login = ?", l)
	if err != nil {
        	return false, err
	}

	var c int = 0;
	for rows.Next() {
		c++;
	        if err := rows.Scan(&a.id, &a.md5); err != nil {
        	        return false, err
        	}
		a.login = l
		return true, err
	}
	if c == 0 {
		err := errors.New("Nothing load");
		return false, err
	}
	if err := rows.Err(); err != nil {
        	return false, err
	}

	fmt.Printf("It is not normal, be here!\n")
	return false, err
}

func main(){
	Auth := new(TAuth)

	r, err := Auth.GetByLogin("test123")
	if (r){
		fmt.Printf("We get data: [%d][%s][%s]\n", Auth.id, Auth.login, Auth.md5)
	}else{
		fmt.Printf("Nothing to get (%s)\n",err.Error())
	}
	
}
