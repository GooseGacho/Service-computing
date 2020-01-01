package dao

import (
	"log"
	"github.com/boltdb/bolt"
)

func createDB() {
	db, err := bolt.Open("data.db", 0600, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
}

func initUserDB() {
	db,_ := bolt.Open("data.db", 0600, nil)
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("User"))
		if b == nil {
			_, err := tx.CreateBucket([]byte("User"))
			if err != nil {
				log.Fatal(err)
			}
		}
		//err := b.Put([]byte("name"), []byte("nam"))
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func initBlogDB() {
	db,_ := bolt.Open("data.db", 0600, nil)
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Blog"))
		if b == nil {
			_, err := tx.CreateBucket([]byte("Blog"))
			if err != nil {
				log.Fatal(err)
			}
		}
		//err := b.Put([]byte("bname1"), []byte("a"))
		//err = b.Put([]byte("bname1"), []byte("a"))
		// if err != nil {
		// 	log.Fatal(err)
		// }
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func initTagDB() {
	db,_ := bolt.Open("data.db", 0600, nil)
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tag"))
		if b == nil {
			_, err := tx.CreateBucket([]byte("Tag"))
			if err != nil {
				log.Fatal(err)
			}
		}
		//err := b.Put([]byte("id"), []byte("code"))
		//err = b.Put([]byte("na"), []byte("life"))
		// if err != nil {
		// 	log.Fatal(err)
		// }
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func initCommentDB() {
	db,_ := bolt.Open("data.db", 0600, nil)
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Comment"))
		if b == nil {
			_, err := tx.CreateBucket([]byte("Comment"))
			if err != nil {
				log.Fatal(err)
			}
		}
		//err := b.Put([]byte("name"), []byte("12"))
		//err = b.Put([]byte("name"), []byte("12"))
		// if err != nil {
		// 	log.Fatal(err)
		// }
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func InitDB() {
	//test
	// createDB()
	// initUserDB()
	// initBlogDB()
	// initCommentDB()
	// initTagDB()
}