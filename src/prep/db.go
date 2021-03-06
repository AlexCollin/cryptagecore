package prep

/*
	db - установка соединений с базами
*/

import (
	_ "pq"

	"config"
	"redis.v4"
	_ "database/sql"
	"sqlx"
	"amqp"
	"log"
)

// Он, главный. Инфа сотка
func InitDB(cfg config.Settings) (*sqlx.DB, *redis.Client) {
	// Хотим ПГ
	sqli := startPG(cfg)
	// И немножко редиса
	//r := RedisOn(cfg)

	return sqli, nil
}

func InitMQ(cfg config.Settings) *amqp.Connection {
	conn, err := amqp.Dial("amqp://" + cfg.MQ.Login + ":" + cfg.MQ.Password + "@" + cfg.MQ.Host + ":" + cfg.MQ.Port + "/")
	if err != nil {
		log.Fatalf("connection.open: %s", err)
	}
	return conn
}

func RedisOn(config config.Settings) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host+":"+config.Redis.Port,
		Password: config.Redis.Password, // нам нечего прятать
		DB:       config.Redis.Database,  // и база у нас по дефолту
	})

	return client
}

// Подключение к PG
func startPG(cfg config.Settings) *sqlx.DB {
	// Взлетаем
	db, err := sqlx.Open("postgres", "postgres://" + cfg.Postgre.User + ":" + cfg.Postgre.Passw + "@" + cfg.Postgre.Host + ":" + cfg.Postgre.Port + "/" + cfg.Postgre.Database + "?sslmode=disable")
	db.SetMaxIdleConns(500)
	db.SetMaxOpenConns(500)

	if err != nil{
		panic("m=GetPool,msg=connection has failed" + err.Error())
	}



	if err != nil {
		return nil
	} else {
		return db
	}
}
