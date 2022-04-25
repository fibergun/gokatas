// Sortbooks sorts a collection of books. To sort a collection of elements you
// need to define a type for this collection. This type needs to have the
// methods that satisfy the sort.Interface interface type.
//
// It's not necessary to use a pointer to book. However when swapping many
// elements it might be faster this way since pointer is always only a machine
// word in size (usually 32 or 64 bits).
//
// 	b := book{"The Go Programming Language", "Kernighan, Donovan", 2015}
// 	p := &b
// 	format := "type: %10T, size: %2d bytes\n"
// 	fmt.Printf(format, b, unsafe.Sizeof(b)) // type:  main.book, size: 40 bytes
// 	fmt.Printf(format, p, unsafe.Sizeof(p)) // type: *main.book, size:  8 bytes
//
// Based on github.com/adonovan/gopl.io/blob/master/ch7/sorting/.
package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
)

type Book struct {
	Title   string
	Authors Authors
	Year    int // when published
}

// START

type Authors []string

func (a Authors) String() string {
	return strings.Join(a, ", ")
}

var books = []*Book{
	{"The Lord of the Rings", Authors{"Tolkien"}, 1954},
	{"The Phoenix Project", Authors{"Kim", "Behr", "Spafford"}, 2013},
	{"The Go Programming Language", Authors{"Kernighan", "Donovan"}, 2015},
}

type byYear []*Book

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// STOP

func main() {
	sort.Sort(sort.Reverse(byYear(books)))
	printBooks(books)
}

func printBooks(books []*Book) {
	const format = "%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Authors", "Year")
	fmt.Fprintf(tw, format, "-----", "-------", "----")
	for _, b := range books {
		// you don't have to derefence here like (*b).title
		fmt.Fprintf(tw, format, b.Title, b.Authors, b.Year)
	}
	tw.Flush() // calculate column widths and print table
}
