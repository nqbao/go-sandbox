package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/nqbao/learn-go/dynonote/model"
	"github.com/nqbao/learn-go/dynonote/restapi"
	"github.com/nqbao/learn-go/dynonote/service"
)

func testHeavyWrite(user string) {
	// test heavy writing

	wg := sync.WaitGroup{}

	count := 100
	workers := 2

	nm := service.NewNoteManager(nil)

	for z := 0; z < workers; z++ {
		go func(starter int) {
			t := time.Now()

			for i := 0; i < count; i++ {
				note := &model.Note{}
				note.UserKey = user
				note.Title = fmt.Sprintf("my note %v", z*count+i)
				note.Content = "world"
				nm.CreateNote(note)

				if i > 0 && i%100 == 0 {
					fmt.Printf("Write %v items in %v seconds\n", i, (time.Now().Unix() - t.Unix()))

					t = time.Now()
				}
			}

			wg.Done()
		}(z)

		wg.Add(1)
	}

	wg.Wait()
}

func testHeavyRead() {
	t := time.Now()
	for i := 0; i < 100; i++ {
		l, err := service.GetAllNotes()

		if err != nil {
			fmt.Printf("%v\n", err)
		} else {
			fmt.Printf("Read %v in %v seconds\n", len(l), (time.Now().Unix() - t.Unix()))
		}

		t = time.Now()
	}
}

func main() {

	// note := service.GetNote("test", "01EB5GRBK31TZ37GA3KKC29SQG")
	// if note != nil {
	// 	fmt.Printf(note.Content)
	// 	note.Content = "again two three four"

	// 	service.UpdateNote(note)
	// }

	// service.DeleteNote("khanh", "01EB9C4K6PSR2JYHEJ5D39DEAQ")
	// service.StarNote("khanh", "01EB9C4K45AZ0PMDNFEC41YMWF", 0)
	// n := service.GetNote("khanh", "01EB9C4K45AZ0PMDNFEC41YMWF")
	// if n != nil {
	//	fmt.Printf("%v", n)
	//}

	// notes := service.GetStarNotes("khanh")
	// for _, n := range notes {
	// 	fmt.Printf("%v %v %v\n", n.ULID, n.UserKey, n.Title)
	// }

	// testHeavyRead()

	restapi.StartServer()
}
