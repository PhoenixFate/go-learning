package main

//bolt 数据库
//记得要配置goPath
//go get github.com/boltdb/bolt/...
import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	//1.打开数据库  0600 读写
	db, openDbError := bolt.Open("test.db", 0600, nil)
	defer db.Close()
	if openDbError != nil {
		log.Panic("开始数据库失败")
	}
	//2.找到抽屉bucket，如果没有则创建
	db.Update(func(tx *bolt.Tx) error {
		//tx 事务
		//// Returns nil if the bucket does not exist.
		bucket := tx.Bucket([]byte("test"))
		if bucket == nil {
			//创建bucket
			var createBucketError error
			bucket, createBucketError = tx.CreateBucket([]byte("test"))
			if createBucketError != nil {
				log.Panic("创建bucket test失败")
			}
		}
		//已经有bucket了
		//3.写数据
		//Put插入数据
		//Get查询数据
		putError := bucket.Put([]byte("test"), []byte("test abc"))
		if putError != nil {
			log.Panic(putError)
		}
		bucket.Get([]byte("test"))
		return nil
	})

	//4.查询数据
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("test"))
		if bucket == nil {
			log.Panic("bucket 不存在")
		}
		result := bucket.Get([]byte("test"))
		fmt.Println("bolt查询到的数据:", string(result))
		fmt.Printf("bolt查询到的数据:%s\n", result)
		return nil
	})

}
