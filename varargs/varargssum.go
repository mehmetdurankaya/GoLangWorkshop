package varargs



func VarargsSum(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
		

	}
	return sum
}