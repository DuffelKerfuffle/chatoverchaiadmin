package documentsforadmin

//Doc is a document
type Doc struct {
	Title string
	Text  string
}

func (d *Doc) edit(NewTitle, newText string) {
	if NewTitle != "" {
		d.Title = NewTitle
	}
	if newText != "" {
		d.Text = newText
	}
}

//DocumentManager manages documents
type DocumentManager struct {
	Docs []Doc
}

//Add adds items
func (a *DocumentManager) Add(title, text string) bool {
	newDoc := Doc{title, text}
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
func (a *DocumentManager) Change(title, newtitle, newText string) bool {
	for i := 0; i < len(a.Docs); i++ {
		if a.Docs[i].Title == title {
			a.Docs[i].edit(newtitle, newText)
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

/*
//FileManager is for images/videos
type File struct {
	Title   string
	FileSrc string
}

func (m *File) editImg(newTitle, newImg string) {
	writeIfNotEmpty(&m.Title, newTitle)
	writeIfNotEmpty(&m.FileSrc, newImg)
}

func writeIfNotEmpty(s *string, s1 string) {
	if s1 != "" {
		*s = s1
	}
}

//DocumentManager manages documents
type FileManager struct {
	Files []File
}

//Add adds items
func (m *FileManager) Add(title, file string) bool {
	newFile := File{title, file}
	for _, file1 := range m.Files {
		if file1.Title == title {
			return false
		}
	}
	//MID: the title is exclusive
	m.Files = append(m.Files, newFile)
	return true
}

//Remove removes items
func (m *FileManager) Remove(title string) bool {
	newFiles := make([]File, len(m.Files)-1)
	j := 0
	for i := 0; i < len(m.Files); i++ {
		if m.Files[i].Title != title {
			if j == len(m.Files)-1 {
				m.Files = append(newFiles, m.Files[i])
				return false
			}
			newFiles[j] = m.Files[i]
			j++
		}
	}
	m.Files = newFiles
	return true
}

//Change PRE: title has either 1 or 0 occurences in the list
//changes docs
func (m *FileManager) Change(title, newtitle, newImg string) bool {
	for i := 0; i < len(m.Files); i++ {
		if m.Files[i].Title == title {
			m.Files[i].editImg(newtitle, newImg)
			return true
		}
	}
	return false
}

//GetDoc gets specific documents
func (m *FileManager) GetFile(title string) File {
	b := File{}
	for i := 0; i < len(m.Files); i++ {
		if m.Files[i].Title == title {
			b = m.Files[i]
		}
	}
	return b
}

//GetAllDocs gets all of the documents in the manager
func (m *FileManager) GetAllFiles() []File {
	b := make([]File, len(m.Files))
	for i := 0; i < len(m.Files); i++ {
		b[i] = m.Files[i]
	}
	return b
}*/
