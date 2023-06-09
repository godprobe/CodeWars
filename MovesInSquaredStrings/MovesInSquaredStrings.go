package kata

func VertMirror(s string) string {
	// your code
	return "VertFail"
}

func HorMirror(s string) string {
	// your code
	return "HorFail"
}

type FParam func(string) string

func Oper(f FParam, x string) string {
	// s = "abcd\nefgh\nijkl\nmnop"
	// oper(vert_mirror, s) => "dcba\nhgfe\nlkji\nponm"
	// oper(hor_mirror, s) => "mnop\nijkl\nefgh\nabcd"
	// your code
	return "OperFail"
}
