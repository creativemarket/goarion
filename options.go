package goarion

type Gravity int

const (
	CENTER Gravity = iota
	NORTH
	SOUTH
	WEST
	EAST
	NORTH_WEST
	NORTH_EAST
	SOUTH_WEST
	SOUTH_EAST
)

func GravtiyToString(g Gravity) string {
	switch g {
	case NORTH:
		return "n"
    case SOUTH:
		return "s"
    case WEST:
		return "w"
	case EAST:
		return "e"
    case NORTH_WEST:
		return "nw"
    case NORTH_EAST:
		return "ne"
    case SOUTH_WEST:
		return "sw"
	case SOUTH_EAST:
		return "se"
	default:
		return "c"
	}
}

type Algo int

const (
	WIDTH Algo = iota
    HEIGHT
    SQUARE
	FILL
)

type Options struct {
	Width         int
	Height        int
	Algo          Algo
	Gravity       Gravity
	Quality       int
	SharpenRadius float64
	SharpenAmount int
}

func AlgoToString(a Algo) string {
	switch a {
	case WIDTH:
		return "width"
    case HEIGHT:
		return "height"
    case SQUARE:
		return "square"
	case FILL:
		return "fill"
	default:
		return "invalid"
	}
}
