package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

func main() {
	var tracks = []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
	table := TrackTable{tracks, make([]string, 0)}

	table.Click("artist")
	printTracks(tracks)
	println()

	table.Click("title")
	printTracks(tracks)
	println()

	table.Click("length")
	printTracks(tracks)
	println()

	table.Click("year")
	printTracks(tracks)
	println()

	table.Click("album")
	printTracks(tracks)
}

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

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}
