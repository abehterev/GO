package main;

import(
	"errors"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"encoding/json"
	"os"
)

type Login interface {
	GetByLogin(l string) (bool, error)
}

type TConfig struct {
	Server string
	Login string
	Password string
	Database string
}

func (c *TConfig) ReadConfig(file string) error {

	f, err := os.Open(file)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(f)
	if decoder.Decode(&c) != nil {
		return err
	}

	return nil
	
}

type TAuth struct {
    id int
    login string
    md5 string
}

func (a *TAuth) GetByLogin(l string, C *TConfig) (bool, error){

	db, err := sql.Open("mysql", C.Login+":"+C.Password+"@"+C.Server+"/"+C.Database)
	if err != nil {
		return false, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return false, err
	}

	rows, err := db.Query("SELECT id,md5pass FROM users WHERE login = ?", l)
	/*if err != nil {
        	return false, err
	}*/

	var c int = 0
	if (err == sql.ErrNoRows){
		fmt.Printf("Err: Nothing to het\n")
	}
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

	C := new(TConfig)
	err := C.ReadConfig("conf.json")
	if (err != nil) {
		fmt.Printf("Config read error: %s\n",err.Error())
		os.Exit(1)
	}else{
		fmt.Println(*C)
	}

	r, err := Auth.GetByLogin("test", C)
	if (r){
		fmt.Printf("We get data: [%d][%s][%s]\n", Auth.id, Auth.login, Auth.md5)
	}else{
		fmt.Printf("Nothing to get (%s)\n",err.Error())
	}
	
}
