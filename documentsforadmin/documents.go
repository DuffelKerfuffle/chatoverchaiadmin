package documentsforadmin

import (
	"fmt"
)

//Doc is a document
type Doc struct {
	Title    string
	Text     string
	MediaSrc string
}

func (d *Doc) edit(NewTitle, newText, newimg string) {
	if NewTitle != "" {
		d.Title = NewTitle
	}
	if newText != "" {
		d.Text = newText
	}
	if newimg != "" {
		d.MediaSrc = newimg
	}
}

//DocumentManager manages documents
type DocumentManager struct {
	Docs []Doc
}

//Add adds items
func (a *DocumentManager) Add(title, text, imgSrc string) bool {
	newDoc := Doc{title, text, imgSrc}
	for _, doc := range a.Docs {
		if doc.Title == title {
			return false
		}
	}
	//MID: the title is exclusive
	a.Docs = append(a.Docs, newDoc)
	return true
}

//Remove removes items
func (a *DocumentManager) Remove(title string) bool {
	newDocs := make([]Doc, len(a.Docs)-1)
	j := 0
	for i := 0; i < len(a.Docs); i++ {
		if a.Docs[i].Title != title {
			if j == len(a.Docs)-1 {
				a.Docs = append(newDocs, a.Docs[i])
				return false
			}
			newDocs[j] = a.Docs[i]
			j++
		}
	}
	a.Docs = newDocs
	return true
}

//Change PRE: title has either 1 or 0 occurences in the list
//changes docs
func (a *DocumentManager) Change(title, newtitle, newText, newimg string) bool {
	for i := 0; i < len(a.Docs); i++ {
		if a.Docs[i].Title == title {
			a.Docs[i].edit(newtitle, newText, newimg)
			return true
		}
	}
	return false
}

//GetDoc gets specific documents
func (a *DocumentManager) GetDoc(title string) Doc {
	b := Doc{}
	for i := 0; i < len(a.Docs); i++ {
		if a.Docs[i].Title == title {
			b = a.Docs[i]
		}
	}
	return b
}

//GetAllDocs gets all of the documents in the manager
func (a *DocumentManager) GetAllDocs() []Doc {
	b := make([]Doc, len(a.Docs))
	for i := 0; i < len(a.Docs); i++ {
		b[i] = a.Docs[i]
	}
	return b
}

//MediaManager is for images/videos
type MediaManager struct {
	*DocumentManager
}

func (i *MediaManager) change() {
	fmt.Println("cannot change media")
}
