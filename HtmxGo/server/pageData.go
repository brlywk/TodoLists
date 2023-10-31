package server

// Template data we need to provide
type PageData struct {
	Title string
}

// Struct holding all information needed to render template(s)
var pageData PageData = PageData{
	Title: "Just Another Todo List",
}
