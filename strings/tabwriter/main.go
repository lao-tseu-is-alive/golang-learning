package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {

	w := tabwriter.NewWriter(os.Stdout, 15, 0, 1, ' ',
		tabwriter.AlignRight)
	fmt.Fprintln(w, "username\tfirstname\tlastname\t")
	fmt.Fprintln(w, "sohlich\tRadomir\tSohlich\t")
	fmt.Fprintln(w, "novak\tJohn\tSmith\t")
	w.Flush()
	// Observe how the b's and the d's, despite appearing in the
	// second cell of each line, belong to different columns.
	w = tabwriter.NewWriter(os.Stdout, 0, 0, 1, '.',
		tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "a\tb\tc")
	fmt.Fprintln(w, "aa\tbb\tcc")
	fmt.Fprintln(w, "aaa\t") // trailing tab
	fmt.Fprintln(w, "aaaa\tdddd\teeee")
	w.Flush()

}
