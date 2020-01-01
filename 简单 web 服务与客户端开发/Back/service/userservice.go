package service 

import (
	"github.com/cjf/smog/gen-go/rpc"
)

func GetUser(s string) (r *rpc.User, err error) {
	// fmt.Println("OHH!!")
	// db, err := bolt.Open("data.db", 0600, nil)
	// err = db.View(func(tx *bolt.Tx) error {
    //     b := tx.Bucket([]byte("MyBlocks"))
    //     if b != nil {
    //         data := b.Get([]byte("l"))
    //         fmt.Printf("%s\n", data)
    //         data := b.Get([]byte("l"))
    //         fmt.Printf("%s\n", data)
    //     }
    //     return nil
    // })
    // if err != nil {
    //     log.Fatal(err)
	// }
	//test
	return &rpc.User{ID: 1, Name: "周星星吖", Creacnt: 22, Fancnt: 67, Zancnt: 37, Commentcnt: 16, Visitcnt: 1550}, nil
}