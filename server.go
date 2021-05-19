package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
github.com/dgrijalva/jwt-go"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"go.uber.org/fx"
)

type User struct {
	Id       uint6
	Login    string
	Password string
}

func register(l fx.Lifecycle, db *pg.DB) {
	l.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				fmt.Println("Start")
				//http.HandleFunc
				return nil
			},
			OnStop: func(ctx context.Context) error {
				fmt.Println("Stop")
				return db.Close()
			},
		},
	)
}

func DB_Model() {
	db := pg.Connect(&pg.Options{
		User: "postgres",
	})
	defer db.Close()

	err := createSchema(db)
	if err != nil {
		panic(err)
	}

	user1 := &User{
		Login:    "admin",
		Password: "pass",
	}
	_, err = db.Model(user1).Insert()
	if err != nil {
		panic(err)
	}

}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*User)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func auth(req *http.Request, db *pg.DB) error {

	u := &User{
		Login:    `json:"Login"`,
		Password: `json:"Password"`,
	}

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&u)
	if err != nil {
		panic(err)
	}

	user := new(User)
	err2 := db.Model(user).
		Relation("Login").
		Where("user.Login = ?", u.Login)
	if err2 != nil {
		panic("User not found")
	}

	if user.Password != u.Password {
		panic("Wrong password")
	}

	token, err:= CreateToken(user.ID)
	f err != nil {
   c.JSON(http.StatusUnprocessableEntity, err.Error())
	   return
	}

return nil
}

func CreateToken(userId uint64) (string, error) {
	os.Setenv("CCESS_SECRET", "TestJwtSecret")
	atClaims : jwt.MapClaims{}
	atClais["authorized"] = true
	atClaims["user_id"] = uerId
at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, er := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	f err != nil {
	   return "", err
	}
	return token, nil
  }
  

// func addHandddler () {
// 	http.HandleFunc("/", auth)
// }

func main() {
	app := fx.New(fx.Provide(
		NewConnection,
		NewConfig,
		DB_Model,
		auth,
	), fx.Invoke(register))

	app.Run()
}
