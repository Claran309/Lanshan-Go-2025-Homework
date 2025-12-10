package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Character struct {
	Name         string `redis:"Name"`
	Profession   string `redis:"Profession"`
	Cost         int    `redis:"Cost"`
	Favorability int    `redis:"Favorability"`
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "",
		DB:           0,               // DB_ID
		Protocol:     2,               // 使用 Redis 协议版本 2（默认，最常用） - Protocol: 0 表示自动选择
		PoolSize:     100,             // 连接池大小
		MinIdleConns: 10,              // 最小空闲连接数
		DialTimeout:  5 * time.Second, // 连接超时
		ReadTimeout:  3 * time.Second, // 读取超时
		WriteTimeout: 3 * time.Second, // 写入超时
	})
	defer client.Close()

	ctx := context.Background()

	// 连接测试
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("连接 Redis 失败: %v", err))
	}
	fmt.Println("Redis 连接成功:", pong)

	fmt.Println("==================================== Lab1 =======================================")
	err1 := client.Set(ctx, "Elaina", "The ashen witch", 0).Err() // 没有.Err()就不返回err
	if err1 != nil {
		fmt.Printf("set failed: %v\n", err1)
	}

	val, err := client.Get(ctx, "Elaina").Result() // .Result()可以返回值，也可用于Set，用于Set时返回操作结果（Ok）
	if err != nil {
		fmt.Printf("get failed: %v\n", err)
	}

	fmt.Printf("%s\n", val)

	fmt.Println("==================================== Lab2 =======================================")
	hashField := []string{
		"Name", "Amiya",
		"Profession", "Healthier",
		"Cost", "15",
		"Favorability", "200",
	}

	resSet, err := client.HSet(ctx, "Character:1", hashField).Result()
	if err != nil {
		fmt.Printf("hset failed: %v\n", err)
	}
	fmt.Printf("%v\n", resSet) // -> 4 , k-v量

	resName, err := client.HGet(ctx, "Character:1", "Name").Result()
	if err != nil {
		fmt.Printf("hget failed: %v\n", err)
	}
	fmt.Printf("%s\n", resName) // -> Amiya

	resProf, err := client.HGet(ctx, "Character:1", "Profession").Result()
	if err != nil {
		fmt.Printf("hget failed: %v\n", err)
	}
	fmt.Printf("%s\n", resProf) // -> Healthier

	resCost, err := client.HGet(ctx, "Character:1", "Cost").Result()
	if err != nil {
		fmt.Printf("hget failed: %v\n", err)
	}
	fmt.Printf("%s\n", resCost) // -> 15

	resFav, err := client.HGet(ctx, "Character:1", "Favorability").Result()
	if err != nil {
		fmt.Printf("hget failed: %v\n", err)
	}
	fmt.Printf("%s\n", resFav) // -> 200

	resCharacter, err := client.HGetAll(ctx, "Character:1").Result()
	if err != nil {
		fmt.Printf("hget failed: %v\n", err)
	}
	fmt.Printf("%s\n", resCharacter) // -> map[Cost:15 Favorability:200 Name:Amiya Profession:Healthier]

	fmt.Println("==================================== Lab3 =======================================")
	var Info Character
	err = client.HGetAll(ctx, "Character:1").Scan(&Info) // 绑定Redis-Hash到结构体
	if err != nil {
		fmt.Printf("hget failed: %v\n", err)
	}
	fmt.Printf("Name: %13s\nProfession:   %8s\nCost: %10d\nFavorability: %d\n", Info.Name, Info.Profession, Info.Cost, Info.Favorability)
}
