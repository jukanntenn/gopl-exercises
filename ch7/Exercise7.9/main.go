package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"
)

type TrackTable struct {
	t []*Track
	c []string
}

func (x TrackTable) Len() int { return len(x.t) }
func (x TrackTable) Less(i, j int) bool {
	for pos := len(x.c) - 1; pos >= 0; pos-- {
		switch x.c[pos] {
		case "title":
			if x.t[i].Title != x.t[j].Title {
				return x.t[i].Title < x.t[j].Title
			}
		case "artist":
			if x.t[i].Artist != x.t[j].Artist {
				return x.t[i].Artist < x.t[j].Artist
			}
		case "album":
			if x.t[i].Album != x.t[j].Album {
				return x.t[i].Album < x.t[j].Album
			}
		case "year":
			if x.t[i].Year != x.t[j].Year {
				return x.t[i].Year < x.t[j].Year
			}
		case "length":
			if x.t[i].Length != x.t[j].Length {
				return x.t[i].Length < x.t[j].Length
			}
		}
	}
	// default by Title
	return x.t[i].Title < x.t[j].Title
}
func (x TrackTable) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }
func (x *TrackTable) Click(head string) {
	for i, v := range x.c {
		if v == head {
			x.c = append(x.c[:i], x.c[i+1:]...)
		}
	}
	x.c = append(x.c, head)
	sort.Sort(x)
}

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

var html = template.Must(template.New("table").Parse(`
<html>
<body>
<table>
	<tr>
		<th><a href="?sort=title">Title</a></th>
		<th><a href="?sort=artist">Artist</a></th>
		<th><a href="?sort=album">Album</a></th>
		<th><a href="?sort=year">Year</a></th>
		<th><a href="?sort=length">Length</a></th>
	</tr>
{{range .}}
	<tr>
		<td>{{.Title}}</td>
		<td>{{.Artist}}</td>
		<td>{{.Album}}</td>
		<td>{{.Year}}</td>
		<td>{{.Length}}</td>
	</td>
{{end}}
</body>
</html>
`))

func main() {
	table := TrackTable{tracks, make([]string, 0)}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		head := r.FormValue("sort")
		if head == "" {
			head = "title"
		}
		table.Click(head)
		err := html.Execute(w, &tracks)
		if err != nil {
			log.Printf("template error: %s", err)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
