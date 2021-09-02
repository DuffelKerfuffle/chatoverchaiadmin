package storage

import (
	documents "chatoverchaiadmin/documentsforadmin"
	"context"
	"fmt"
	"strconv"

	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

//Adder allows both mediamanagers and documentmanagers to add
type Adder interface {
	Add(title, text, imgSrc string) bool
}

// Getter gets the documents
type Getter interface {
	GetAllDocs() []documents.Doc
}

//Store updates the database
func Store(g Getter, c string) {
	Remove(c)
	opt := option.WithCredentialsFile("chatoverchai-a0df7-firebase-adminsdk-kwyvn-e3f8d1f06d.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Println("store")
		panic(err)
	}

	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		fmt.Println("store1")
		panic(err)
	}

	for i := 0; i < len(g.GetAllDocs()); i++ {
		y := strconv.Itoa(i)
		fmt.Println(y)
		result, err := client.Collection(c).Doc("Doc"+y).Set(ctx, g.GetAllDocs()[i])
		if err != nil {
			fmt.Println("store2")
			panic(err)
		}
		fmt.Println(result)
	}
}

//Load loads
func Load(a Adder, c string) {
	opt := option.WithCredentialsFile("chatoverchai-a0df7-firebase-adminsdk-kwyvn-e3f8d1f06d.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Println("store4")
		panic(err)
	}

	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		fmt.Println("store5")
		panic(err)
	}
	iter := client.Collection(c).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			panic(err)
		}
		a.Add(doc.Data()["Title"].(string), doc.Data()["Text"].(string), doc.Data()["MediaSrc"].(string))
	}
}

//Remove removes
func Remove(c string) {
	count := 0
	opt := option.WithCredentialsFile("chatoverchai-a0df7-firebase-adminsdk-kwyvn-e3f8d1f06d.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		fmt.Println("store4")
		panic(err)
	}

	ctx := context.Background()
	client, err := app.Firestore(ctx)

	if err != nil {
		fmt.Println("store5")
		panic(err)
	}

	iter := client.Collection(c).Documents(ctx)

	for {
		_, err := iter.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			panic(err)
		}
		fmt.Println("Doc" + strconv.Itoa(count))
		_, err = client.Collection(c).Doc("Doc" + strconv.Itoa(count)).Delete(ctx)
		if err != nil {
			panic(err)
		}
		count++
	}

}
