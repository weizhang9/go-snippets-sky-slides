package main

import "fmt"

type ByteSize float64 // declare a new type called ByteSize with underlying type of float64

const (
	// iota is an enumerated constant, starting at 0 from each declaration
	_ = iota // ignore first value (0) by assigning to blank identifier

	// establishing a formula for constants in this block
	// the rest of the constants will follow the same formula
	// while incrementing iota each line
	KB ByteSize = 1 << (10 * iota) // iota = 1, i.e. shifting 10 bit to the left in binary
	MB                             // which is 10,000,000,000 in binary, which is 1024 in decimal
	GB                             // same kind calculations on MB, GB, TB, PB, EB, ZB, YB
	TB
	PB
	EB
	ZB
	YB
)

// attached a method String who return a string to type ByteSize
func (b ByteSize) String() string {
	// dividing the size of ByteSize by largest file size available under the size of ByteSize
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB) // expecting a number with 2 floating points to be printed
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}

func main() {
	fmt.Println("Binary\t\t\tDecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%f\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%f\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%f\n", GB)
	fmt.Printf("%b\t", TB)
	fmt.Printf("%f\n", TB)
	fmt.Printf("%b\t", PB)
	fmt.Printf("%f\n", PB)
	fmt.Printf("%b\t", EB)
	fmt.Printf("%f\n", EB)
	fmt.Printf("%b\t", ZB)
	fmt.Printf("%f\n", ZB)
	fmt.Printf("%b\t", YB)
	fmt.Printf("%f\n", YB)
	s := ByteSize(2379473.73)
	fmt.Println(s.String())
}
