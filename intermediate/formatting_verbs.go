package intermediate

import "fmt"

func main() {
	// --- General formatting verbs
	// %v  Prints the value in the default format
	// %#v Prints the value in the Go-syntax format
	// %T  Prints the type of the value
	// %%  Prints the % sign

	i := 143_505.5
	str := "Hello, World!"

	fmt.Println("--- General formatting verbs (float) ---")
	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%v%%\n", i)

	fmt.Println("--- General formatting verbs (string) ---")
	fmt.Printf("%v\n", str)
	fmt.Printf("%#v\n", str)
	fmt.Printf("%T\n", str)

	// --- Integer Formatting Verbs
	// %b   Base 2
	// %d   Base 10
	// %+d  Base 10 and always show sign
	// %o   Base 8
	// %#o  Base 8 with leading 0
	// %O   Base 8 with leading 0o
	// %x   Base 16 lowercase
	// %X   Base 16 uppercase
	// %#x  Base 16 with leading 0x
	// %4d  Pad with spaces (width 4, right justified)
	// %-4d Pad with spaces (width 4, left justified)
	// %04d Pad with zeros (width 4, right justified)

	n := 255

	fmt.Println("--- Integer Formatting Verbs ---")
	fmt.Printf("%b\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%+d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%#o\n", n)
	fmt.Printf("%O\n", n)
	fmt.Printf("%x\n", n)
	fmt.Printf("%X\n", n)
	fmt.Printf("%#x\n", n)
	fmt.Printf("%4d\n", n)
	fmt.Printf("%-4d\n", n)
	fmt.Printf("%04d\n", n)

	// --- String formatting verbs
	// %s   Prints the value as plain string
	// %q   Prints the value as a double quoted string
	// %8s  Prints the value as plain string (width 8, right justified)
	// %-8s Prints the value as plain string (width 8, left justified)
	// %x   Prints the value as hex dump of byte values
	// % x  Prints the value as hex dump with spaces

	txt := "World"

	fmt.Println("--- String Formatting Verbs ---")
	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%8s\n", txt)
	fmt.Printf("%-8s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("% x\n", txt)

	// --- Boolean formatting verbs
	// %t   Value of the boolean operator in true or false format (same as using %v)

	t := true
	f := false

	fmt.Println("--- Boolean Formatting Verbs ---")
	fmt.Printf("%t\n", t)
	fmt.Printf("%t\n", f)
	fmt.Printf("%v\n", t)
	fmt.Printf("%v\n", f)

	// --- Float formatting verbs
	// %e     Scientific notation with 'e' as exponent
	// %f     Decimal point, no exponent
	// %.2f   Default width, precision 2
	// %6.2f  Width 6, precision 2
	// %g     Exponent as needed, only necessary digits

	fl := 9.18

	fmt.Println("--- Float Formatting Verbs ---")
	fmt.Printf("%e\n", fl)
	fmt.Printf("%f\n", fl)
	fmt.Printf("%.2f\n", fl)
	fmt.Printf("%6.2f\n", fl)
	fmt.Printf("%g\n", fl)
}
