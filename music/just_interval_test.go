package music

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterval_String(t *testing.T) {
	type fields struct {
		numerator   uint
		denominator uint
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "String representation of perfect fifth",
			fields: fields{
				numerator:   3,
				denominator: 2,
			},
			want: "3:2",
		},
		{
			name: "String representation of septimal minor seventh",
			fields: fields{
				numerator:   7,
				denominator: 6,
			},
			want: "7:6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := JustInterval{
				numerator:   tt.fields.numerator,
				denominator: tt.fields.denominator,
			}
			if got := i.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterval_add(t *testing.T) {
	type fields struct {
		numerator   uint
		denominator uint
	}
	type args struct {
		other JustInterval
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   JustInterval
	}{
		{
			name: "Adding a Synthonic Comma to a lesser major second produces a greater major second",
			fields: fields{
				numerator:   10,
				denominator: 9,
			},
			args: args{
				other: JustInterval{
					numerator:   81,
					denominator: 80,
				},
			},
			want: JustInterval{
				numerator:   9,
				denominator: 8,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := JustInterval{
				numerator:   tt.fields.numerator,
				denominator: tt.fields.denominator,
			}
			if got := i.Add(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterval_Diff(t *testing.T) {
	type fields struct {
		numerator   uint
		denominator uint
	}
	type args struct {
		other JustInterval
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   JustInterval
	}{
		{
			name: "Difference between a lesser major second and a greater major second is a syntonic comma",
			fields: fields{
				numerator:   10,
				denominator: 9,
			},
			args: args{
				other: JustInterval{
					numerator:   9,
					denominator: 8,
				},
			},
			want: JustInterval{
				numerator:   81,
				denominator: 80,
			},
		},
		{
			name: "Difference between a greater major second and a lesser major second is a syntonic comma",
			fields: fields{
				numerator:   9,
				denominator: 8,
			},
			args: args{
				other: JustInterval{
					numerator:   10,
					denominator: 9,
				},
			},
			want: JustInterval{
				numerator:   81,
				denominator: 80,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := JustInterval{
				numerator:   tt.fields.numerator,
				denominator: tt.fields.denominator,
			}
			if got := i.Diff(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterval_greaterThan(t *testing.T) {
	type fields struct {
		numerator   uint
		denominator uint
	}
	type args struct {
		other JustInterval
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Greater than test for greater major second and lesser major second",
			fields: fields{
				numerator:   9,
				denominator: 8,
			},
			args: args{
				other: JustInterval{
					numerator:   10,
					denominator: 9,
				},
			},
			want: true,
		},
		{
			name: "Greater than test for lesser major second and greater major second",
			fields: fields{
				numerator:   10,
				denominator: 9,
			},
			args: args{
				other: JustInterval{
					numerator:   9,
					denominator: 8,
				},
			},
			want: false,
		},
		{
			name: "Greater than test for equal intervals",
			fields: fields{
				numerator:   9,
				denominator: 8,
			},
			args: args{
				other: JustInterval{
					numerator:   9,
					denominator: 8,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := JustInterval{
				numerator:   tt.fields.numerator,
				denominator: tt.fields.denominator,
			}
			if got := i.GreaterThan(tt.args.other); got != tt.want {
				t.Errorf("greaterThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterval_isDiminishedFifth(t *testing.T) {
	type fields struct {
		numerator   uint
		denominator uint
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Diminished fifth test for diminished fifth interval",
			fields: fields{
				numerator:   64,
				denominator: 45,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := JustInterval{
				numerator:   tt.fields.numerator,
				denominator: tt.fields.denominator,
			}
			if got := i.IsDiminishedFifth(); got != tt.want {
				t.Errorf("isDiminishedFifth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterval_isEqualTo(t *testing.T) {
	type fields struct {
		numerator   uint
		denominator uint
	}
	type args struct {
		other JustInterval
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Equality test for equal intervals",
			fields: fields{
				numerator:   3,
				denominator: 2,
			},
			args: args{
				other: JustInterval{
					numerator:   3,
					denominator: 2,
				},
			},
			want: true,
		},
		{
			name: "Equality test for unequal intervals",
			fields: fields{
				numerator:   3,
				denominator: 2,
			},
			args: args{
				other: JustInterval{
					numerator:   4,
					denominator: 3,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := JustInterval{
				numerator:   tt.fields.numerator,
				denominator: tt.fields.denominator,
			}
			if got := i.IsEqualTo(tt.args.other); got != tt.want {
				t.Errorf("isEqualTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterval_isOctave(t *testing.T) {
	type fields struct {
		numerator   uint
		denominator uint
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Octave test for octave interval",
			fields: fields{
				numerator:   2,
				denominator: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := JustInterval{
				numerator:   tt.fields.numerator,
				denominator: tt.fields.denominator,
			}
			if got := i.IsOctave(); got != tt.want {
				t.Errorf("isOctave() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterval_isPerfect(t *testing.T) {
	type fields struct {
		numerator   uint
		denominator uint
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Perfect interval test for perfect unison",
			fields: fields{
				numerator:   1,
				denominator: 1,
			},
			want: true,
		},
		{
			name: "Perfect interval test for perfect forth",
			fields: fields{
				numerator:   4,
				denominator: 3,
			},
			want: true,
		},
		{
			name: "Perfect interval test for perfect fifth",
			fields: fields{
				numerator:   3,
				denominator: 2,
			},
			want: true,
		},
		{
			name: "Perfect interval test for perfect octave",
			fields: fields{
				numerator:   2,
				denominator: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := JustInterval{
				numerator:   tt.fields.numerator,
				denominator: tt.fields.denominator,
			}
			if got := i.IsPerfect(); got != tt.want {
				t.Errorf("isPerfect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterval_isPerfectFifth(t *testing.T) {
	type fields struct {
		numerator   uint
		denominator uint
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Perfect fifth test for perfect fifth interval",
			fields: fields{
				numerator:   3,
				denominator: 2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := JustInterval{
				numerator:   tt.fields.numerator,
				denominator: tt.fields.denominator,
			}
			if got := i.IsPerfectFifth(); got != tt.want {
				t.Errorf("isPerfectFifth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterval_lessThan(t *testing.T) {
	type fields struct {
		numerator   uint
		denominator uint
	}
	type args struct {
		other JustInterval
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Return true if interval is less than another interval",
			fields: fields{
				numerator:   10,
				denominator: 9,
			},
			args: args{
				other: JustInterval{
					numerator:   9,
					denominator: 8,
				},
			},
			want: true,
		},
		{
			name: "Return false if interval is greater than another interval",
			fields: fields{
				numerator:   9,
				denominator: 8,
			},
			args: args{
				other: JustInterval{
					numerator:   10,
					denominator: 9,
				},
			},
			want: false,
		},
		{
			name: "Return false if interval is equal to another interval",
			fields: fields{
				numerator:   10,
				denominator: 9,
			},
			args: args{
				other: JustInterval{
					numerator:   10,
					denominator: 9,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := JustInterval{
				numerator:   tt.fields.numerator,
				denominator: tt.fields.denominator,
			}
			if got := i.LessThan(tt.args.other); got != tt.want {
				t.Errorf("lessThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterval_octaveReduce(t *testing.T) {
	type fields struct {
		numerator   uint
		denominator uint
	}
	tests := []struct {
		name   string
		fields fields
		want   JustInterval
	}{
		{
			name: "Octave reduce a major ninth to a major second",
			fields: fields{
				numerator:   9,
				denominator: 4,
			},
			want: JustInterval{
				numerator:   9,
				denominator: 8,
			},
		},
		{
			name: "Octave reduce a reciprocal prefect fourth to a perfect fourth",
			fields: fields{
				numerator:   2,
				denominator: 3,
			},
			want: JustInterval{
				numerator:   4,
				denominator: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := JustInterval{
				numerator:   tt.fields.numerator,
				denominator: tt.fields.denominator,
			}
			if got := i.OctaveReduce(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("octaveReduce() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestInterval_reciprocal(t *testing.T) {
	type fields struct {
		numerator   uint
		denominator uint
	}
	tests := []struct {
		name   string
		fields fields
		want   JustInterval
	}{
		{
			name: "Reciprocal of an  perfect fifth is perfect fourth",
			fields: fields{
				numerator:   81,
				denominator: 80,
			},
			want: JustInterval{
				numerator:   80,
				denominator: 81,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := JustInterval{
				numerator:   tt.fields.numerator,
				denominator: tt.fields.denominator,
			}
			if got := i.Reciprocal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reciprocal() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestInterval_simplify(t *testing.T) {
	type fields struct {
		numerator   uint
		denominator uint
	}
	tests := []struct {
		name   string
		fields fields
		want   JustInterval
	}{
		{
			name: "Simplify 10:8 to 5:4",
			fields: fields{
				numerator:   10,
				denominator: 8,
			},
			want: JustInterval{
				numerator:   5,
				denominator: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := JustInterval{
				numerator:   tt.fields.numerator,
				denominator: tt.fields.denominator,
			}
			if got := i.Simplify(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("simplify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterval_sortWith(t *testing.T) {
	type fields struct {
		numerator   uint
		denominator uint
	}
	type args struct {
		j JustInterval
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "Sorting a lesser major second with a greater major second",
			fields: fields{
				numerator:   10,
				denominator: 9,
			},
			args: args{
				j: JustInterval{
					numerator:   9,
					denominator: 8,
				},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := JustInterval{
				numerator:   tt.fields.numerator,
				denominator: tt.fields.denominator,
			}
			if got := i.sortWith(tt.args.j); got != tt.want {
				t.Errorf("sortWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterval_subtract(t *testing.T) {
	type fields struct {
		numerator   uint
		denominator uint
	}
	type args struct {
		other JustInterval
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   JustInterval
	}{
		{
			name: "Subtracting a lesser major second from a greater major second produces a syntonic comma",
			fields: fields{
				numerator:   9,
				denominator: 8,
			},
			args: args{
				other: JustInterval{
					numerator:   10,
					denominator: 9,
				},
			},
			want: JustInterval{
				numerator:   81,
				denominator: 80,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := JustInterval{
				numerator:   tt.fields.numerator,
				denominator: tt.fields.denominator,
			}
			if got := i.Subtract(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subtract() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestInterval_toCents(t *testing.T) {
	type fields struct {
		numerator   uint
		denominator uint
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Cents value of a perfect fifth",
			fields: fields{
				numerator:   3,
				denominator: 2,
			},
			want: 701.9550008653874,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := JustInterval{
				numerator:   tt.fields.numerator,
				denominator: tt.fields.denominator,
			}
			if got := i.ToCents(); got != tt.want {
				t.Errorf("toCents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterval_toFloat(t *testing.T) {
	type fields struct {
		numerator   uint
		denominator uint
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Float value of a perfect fifth",
			fields: fields{
				numerator:   3,
				denominator: 2,
			},
			want: 1.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := JustInterval{
				numerator:   tt.fields.numerator,
				denominator: tt.fields.denominator,
			}
			if got := i.ToFloat(); got != tt.want {
				t.Errorf("toFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterval_toPowerOf(t *testing.T) {
	type fields struct {
		numerator   uint
		denominator uint
	}
	type args struct {
		p int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   JustInterval
	}{
		{
			name: "Raising a perfect fifth to the power of 2",
			fields: fields{
				numerator:   3,
				denominator: 2,
			},
			args: args{
				p: 2,
			},
			want: JustInterval{
				numerator:   9,
				denominator: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := JustInterval{
				numerator:   tt.fields.numerator,
				denominator: tt.fields.denominator,
			}
			if got := i.ToPowerOf(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toPowerOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createMultiplierTableOf(t *testing.T) {
	type args struct {
		multiplierListA [][]uint
		multiplierListB [][]uint
	}
	tests := []struct {
		name string
		args args
		want [][]uint
	}{
		{
			name: "Create multiplier table from two lists",
			args: args{
				multiplierListA: [][]uint{{5, 1}, {1, 1}, {1, 5}},
				multiplierListB: [][]uint{{3, 1}, {1, 1}, {1, 3}},
			},
			want: [][]uint{
				{15, 1}, {5, 1}, {5, 3},
				{3, 1}, {1, 1}, {1, 3},
				{3, 5}, {1, 5}, {1, 15},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createMultiplierTableOf(tt.args.multiplierListA, tt.args.multiplierListB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createMultiplierTableOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fromIntArray(t *testing.T) {
	type args struct {
		i []uint
	}
	tests := []struct {
		name string
		args args
		want JustInterval
	}{
		{
			name: "Create interval from integer array representing a perfect fifth",
			args: args{
				i: []uint{3, 2},
			},
			want: JustInterval{
				numerator:   3,
				denominator: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromIntArray(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fromIntArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intervalsFromIntegers(t *testing.T) {
	type args struct {
		integers [][]uint
	}
	tests := []struct {
		name string
		args args
		want []JustInterval
	}{
		{
			name: "Create intervals from integer pairs",
			args: args{
				integers: [][]uint{{3, 2}, {5, 4}, {4, 3}},
			},
			want: []JustInterval{
				{numerator: 3, denominator: 2},
				{numerator: 5, denominator: 4},
				{numerator: 4, denominator: 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntervalsFromIntegers(tt.args.integers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intervalsFromIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_justIntervalsFromMultipliers(t *testing.T) {
	type args struct {
		multiplierList [][]uint
		filter         intervalFilterFunction
	}
	tests := []struct {
		name string
		args args
		want []JustInterval
	}{
		{
			name: "Create just intervals from multipliers without filtering",
			args: args{
				multiplierList: [][]uint{{3, 1}, {1, 1}, {1, 3}},
				filter:         func(ratio JustInterval) bool { return false },
			},
			want: []JustInterval{
				{numerator: 1, denominator: 1},
				{numerator: 4, denominator: 3},
				{numerator: 3, denominator: 2},
				{numerator: 2, denominator: 1},
			},
		},
		{
			name: "Filter out intervals from multipliers",
			args: args{
				multiplierList: [][]uint{{3, 1}, {1, 1}, {1, 3}},
				filter:         func(ratio JustInterval) bool { return ratio == PerfectFifth() },
			},
			want: []JustInterval{
				{numerator: 1, denominator: 1},
				{numerator: 4, denominator: 3},
				{numerator: 2, denominator: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := justIntervalsFromMultipliers(tt.args.multiplierList, tt.args.filter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("justIntervalsFromMultipliers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multipliers(t *testing.T) {
	type args struct {
		base uint
	}
	tests := []struct {
		name string
		args args
		want [][]uint
	}{
		{
			name: "Generate multipliers for base 5",
			args: args{
				base: 5,
			},
			want: [][]uint{{5, 1}, {1, 1}, {1, 5}},
		},
		{
			name: "3 is a special case and generates more multipliers in order for there to be enough notes in the scale",
			args: args{
				base: 3,
			},
			want: [][]uint{{9, 1}, {3, 1}, {1, 1}, {1, 3}, {1, 9}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multipliers(tt.args.base); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("multipliers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newInterval(t *testing.T) {
	type args struct {
		numerator   uint
		denominator uint
	}
	tests := []struct {
		name string
		args args
		want JustInterval
	}{
		{
			name: "Create new interval representing a perfect fifth",
			args: args{
				numerator:   3,
				denominator: 2,
			},
			want: JustInterval{
				numerator:   3,
				denominator: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInterval(tt.args.numerator, tt.args.denominator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortIntervals(t *testing.T) {
	type args struct {
		intervals []JustInterval
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Sort a list of just intervals",
			args: args{
				intervals: []JustInterval{
					{numerator: 10, denominator: 9},
					{numerator: 3, denominator: 2},
					{numerator: 5, denominator: 4},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortIntervals(tt.args.intervals)
		})
	}
}

func TestPerfectFourth(t *testing.T) {
	tests := []struct {
		name string
		want JustInterval
	}{
		{
			name: "Returns a perfect fourth interval",
			want: JustInterval{
				numerator:   4,
				denominator: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, PerfectFourth(), "PerfectFourth()")
		})
	}
}

func TestDieses(t *testing.T) {
	tests := []struct {
		name string
		want JustInterval
	}{
		{
			name: "Returns a dieses interval",
			want: JustInterval{
				numerator:   128,
				denominator: 125,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Dieses(), "Dieses()")
		})
	}
}

func TestJustChromaticSemitone(t *testing.T) {
	tests := []struct {
		name string
		want JustInterval
	}{
		{
			name: "Returns a just chromatic semitone interval",
			want: JustInterval{
				numerator:   25,
				denominator: 24,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, JustChromaticSemitone(), "JustChromaticSemitone()")
		})
	}
}

func TestJustInterval_ToTemperedInterval(t *testing.T) {
	type fields struct {
		numerator   uint
		denominator uint
		name        string
	}
	tests := []struct {
		name   string
		fields fields
		want   TemperedInterval
	}{
		{
			name: "Convert a perfect fifth to a tempered interval",
			fields: fields{
				numerator:   3,
				denominator: 2,
			},
			want: TemperedInterval(1.5),
		},
		{
			name: "Any interval with zero on the denominator converts to a tempered interval of zero",
			fields: fields{
				numerator:   3,
				denominator: 0,
			},
			want: TemperedInterval(0),
		},
		{
			name: "Any interval with zero on the numerator converts to a tempered interval of zero",
			fields: fields{
				numerator:   0,
				denominator: 2,
			},
			want: TemperedInterval(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := JustInterval{
				numerator:   tt.fields.numerator,
				denominator: tt.fields.denominator,
				name:        tt.fields.name,
			}
			assert.Equalf(t, tt.want, i.ToTemperedInterval(), "ToTemperedInterval()")
		})
	}
}

func TestJustInterval_Subtract(t *testing.T) {
	type fields struct {
		numerator   uint
		denominator uint
		name        string
	}
	type args struct {
		other JustInterval
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   JustInterval
	}{
		{
			name: "Subtract a perfect fourth from a perfect fifth to get a major second",
			fields: fields{
				numerator:   3,
				denominator: 2,
			},
			args: args{
				other: JustInterval{
					numerator:   4,
					denominator: 3,
				},
			},
			want: JustInterval{
				numerator:   9,
				denominator: 8,
			},
		},
		{
			name: "Subtract a major fifth from a perfect second to get a perfect fourth",
			fields: fields{
				numerator:   9,
				denominator: 8,
			},
			args: args{
				other: JustInterval{
					numerator:   3,
					denominator: 2,
				},
			},
			want: JustInterval{
				numerator:   4,
				denominator: 3,
			},
		},
		{
			name: "Subtracting an interval from itself results in a unison",
			fields: fields{
				numerator:   5,
				denominator: 4,
			},
			args: args{
				other: JustInterval{
					numerator:   5,
					denominator: 4,
				},
			},
			want: JustInterval{
				numerator:   1,
				denominator: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := JustInterval{
				numerator:   tt.fields.numerator,
				denominator: tt.fields.denominator,
				name:        tt.fields.name,
			}
			assert.Equalf(t, tt.want, i.Subtract(tt.args.other), "Subtract(%v)", tt.args.other)
		})
	}
}

func TestJustInterval_Simplify(t *testing.T) {
	type fields struct {
		numerator   uint
		denominator uint
		name        string
	}
	tests := []struct {
		name   string
		fields fields
		want   JustInterval
	}{
		{
			name: "Simplify 10:8 to 5:4",
			fields: fields{
				numerator:   10,
				denominator: 8,
			},
			want: JustInterval{
				numerator:   5,
				denominator: 4,
			},
		},
		{
			name: "Simplify 100:80 to 5:4",
			fields: fields{
				numerator:   100,
				denominator: 80,
			},
			want: JustInterval{
				numerator:   5,
				denominator: 4,
			},
		},
		{
			name: "Simplifying an interval with a zero in the denominator returns a unison",
			fields: fields{
				numerator:   5,
				denominator: 0,
			},
			want: JustInterval{
				numerator:   1,
				denominator: 1,
			},
		},
		{
			name: "Simplifying an interval with a zero in the numerator returns a unison",
			fields: fields{
				numerator:   0,
				denominator: 5,
			},
			want: JustInterval{
				numerator:   1,
				denominator: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := JustInterval{
				numerator:   tt.fields.numerator,
				denominator: tt.fields.denominator,
				name:        tt.fields.name,
			}
			assert.Equalf(t, tt.want, i.Simplify(), "Simplify()")
		})
	}
}

func TestJustInterval_Name(t *testing.T) {
	type fields struct {
		numerator   uint
		denominator uint
		name        string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get the name of an interval",
			fields: fields{
				numerator:   3,
				denominator: 2,
			},
			want: "Perfect Fifth",
		},
		{
			name: "An unnamed interval returns an empty string",
			fields: fields{
				numerator:   1381,
				denominator: 2112,
				name:        "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := JustInterval{
				numerator:   tt.fields.numerator,
				denominator: tt.fields.denominator,
				name:        tt.fields.name,
			}
			assert.Equalf(t, tt.want, i.Name(), "Name()")
		})
	}
}

func TestJustInterval_Invert(t *testing.T) {
	tests := []struct {
		name     string
		interval JustInterval
		want     JustInterval
	}{
		{
			name:     "Invert of a major third is a minor sixth",
			interval: JustInterval{numerator: 5, denominator: 4},
			want:     JustInterval{numerator: 8, denominator: 5},
		},
		{
			name:     "Invert of a perfect fifth is a perfect fourth",
			interval: JustInterval{numerator: 3, denominator: 2},
			want:     JustInterval{numerator: 4, denominator: 3},
		},
		{
			name:     "Invert of a minor second (16/15) is a major seventh (15/8)",
			interval: JustInterval{numerator: 16, denominator: 15},
			want:     JustInterval{numerator: 15, denominator: 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.interval.Invert())
		})
	}
}

func TestJustInterval_DeviationFromEqualTemperament(t *testing.T) {
	tests := []struct {
		name     string
		interval JustInterval
		want     float64
	}{
		{
			name:     "Perfect fifth (3:2) deviates by ~+1.955 cents from 700 cents",
			interval: JustInterval{numerator: 3, denominator: 2},
			want:     1.9550008653873953,
		},
		{
			name:     "Major third (5:4) deviates by ~-13.686 cents from 400 cents",
			interval: JustInterval{numerator: 5, denominator: 4},
			want:     -13.686286135567685,
		},
		{
			name:     "Unison (1:1) has zero deviation",
			interval: JustInterval{numerator: 1, denominator: 1},
			want:     0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.InDelta(t, tt.want, tt.interval.DeviationFromEqualTemperament(), 1e-9)
		})
	}
}

func TestJustInterval_ToPowerOf_NegativeExponents(t *testing.T) {
	tests := []struct {
		name     string
		interval JustInterval
		power    int
		want     JustInterval
	}{
		{
			name:     "Fifth to the power of -1 is a fourth (after octave reducing the reciprocal)",
			interval: JustInterval{numerator: 3, denominator: 2},
			power:    -1,
			want:     JustInterval{numerator: 2, denominator: 3},
		},
		{
			name:     "Fifth to the power of 0 is unison",
			interval: JustInterval{numerator: 3, denominator: 2},
			power:    0,
			want:     JustInterval{numerator: 1, denominator: 1},
		},
		{
			name:     "Fifth to the power of -2 is a minor seventh (2/3)^2 = 4/9",
			interval: JustInterval{numerator: 3, denominator: 2},
			power:    -2,
			want:     JustInterval{numerator: 4, denominator: 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.interval.ToPowerOf(tt.power))
		})
	}
}

func TestJustInterval_OctaveReduce_ZeroGuard(t *testing.T) {
	tests := []struct {
		name     string
		interval JustInterval
		want     JustInterval
	}{
		{
			name:     "OctaveReduce with zero numerator returns unison",
			interval: JustInterval{numerator: 0, denominator: 2},
			want:     JustInterval{numerator: 1, denominator: 1},
		},
		{
			name:     "OctaveReduce with zero denominator returns unison",
			interval: JustInterval{numerator: 3, denominator: 0},
			want:     JustInterval{numerator: 1, denominator: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.interval.OctaveReduce())
		})
	}
}

func TestSyntonicComma_IsAliasForAcuteUnison(t *testing.T) {
	assert.Equal(t, AcuteUnison(), SyntonicComma(), "SyntonicComma and AcuteUnison should return the same ratio (81/80)")
}
