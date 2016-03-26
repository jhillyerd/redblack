package main

import (
	"testing"
)

func TestString(t *testing.T) {
	table := []struct {
		input  *RBNode
		expect string
	}{
		{&RBNode{}, "B:0(,)"},
		{&RBNode{
			10, false,
			&RBNode{5, true, nil, nil},
			&RBNode{
				20, true,
				&RBNode{17, false, nil, nil},
				&RBNode{25, false, nil, nil},
			},
		}, "B:10(R:5(,),R:20(B:17(,),B:25(,)))"},
	}

	for _, te := range table {
		if te.input.String() != te.expect {
			t.Errorf("String() => %q, want %q", te.input.String(), te.expect)
		}
	}
}

func TestRot1Left(t *testing.T) {
	table := []struct {
		input  *RBNode
		expect string
	}{
		{&RBNode{
			10, false,
			&RBNode{5, true, nil, nil},
			&RBNode{
				20, true,
				&RBNode{17, false, nil, nil},
				&RBNode{25, false, nil, nil},
			},
		}, "B:20(R:10(R:5(,),B:17(,)),B:25(,))"},
	}

	for _, te := range table {
		result := te.input.Rot1Left()
		if result.String() != te.expect {
			t.Errorf("Rot1Left() => %q, want %q", result.String(), te.expect)
		}
	}
}

func TestRot1Right(t *testing.T) {
	table := []struct {
		input  *RBNode
		expect string
	}{
		{&RBNode{
			10, false,
			&RBNode{5, true, nil, nil},
			&RBNode{
				20, true,
				&RBNode{17, false, nil, nil},
				&RBNode{25, false, nil, nil},
			},
		}, "B:5(,R:10(,R:20(B:17(,),B:25(,))))"},
	}

	for _, te := range table {
		result := te.input.Rot1Right()
		if result.String() != te.expect {
			t.Errorf("Rot1Right() => %q, want %q", result.String(), te.expect)
		}
	}
}

func TestRot2Left(t *testing.T) {
	table := []struct {
		input  *RBNode
		expect string
	}{
		{&RBNode{
			10, false,
			&RBNode{5, true, nil, nil},
			&RBNode{
				20, true,
				&RBNode{17, false, nil, nil},
				&RBNode{25, false, nil, nil},
			},
		}, "B:17(R:10(R:5(,),),R:20(,B:25(,)))"},
	}

	for _, te := range table {
		result := te.input.Rot2Left()
		if result.String() != te.expect {
			t.Errorf("Rot2Left() => %q, want %q", result.String(), te.expect)
		}
	}
}

func TestRot2Right(t *testing.T) {
	table := []struct {
		input  *RBNode
		expect string
	}{
		{&RBNode{
			10, false,
			&RBNode{
				5, true,
				&RBNode{1, false, nil, nil},
				&RBNode{7, false, nil, nil},
			},
			&RBNode{20, true, nil, nil},
		}, "B:7(R:5(B:1(,),),R:10(,R:20(,)))"},
	}

	for _, te := range table {
		result := te.input.Rot2Right()
		if result.String() != te.expect {
			t.Errorf("Rot2Right() => %q, want %q", result.String(), te.expect)
		}
	}
}

func TestValidate(t *testing.T) {
	table := []struct {
		input        *RBNode
		expectHeight int
		expectError  string
	}{
		{nil, 1, ""},
		{&RBNode{7, false, nil, nil}, 2, ""},
		{&RBNode{
			7, false,
			&RBNode{
				5, false,
				&RBNode{1, false, nil, nil},
				&RBNode{6, false, nil, nil},
			},
			&RBNode{
				10, false,
				&RBNode{9, false, nil, nil},
				&RBNode{20, false, nil, nil},
			},
		}, 4, ""},
		{&RBNode{
			7, false,
			&RBNode{
				5, false,
				&RBNode{1, false, nil, nil},
				&RBNode{6, false, nil, nil},
			},
			nil,
		}, 0, "Black height violation at value 7"},
		{&RBNode{
			7, false,
			&RBNode{
				5, true,
				&RBNode{1, true, nil, nil},
				&RBNode{6, true, nil, nil},
			},
			&RBNode{
				10, false,
				&RBNode{9, false, nil, nil},
				&RBNode{20, false, nil, nil},
			},
		}, 0, "Red violation at value 5"},
		{&RBNode{
			7, false,
			&RBNode{
				5, false,
				&RBNode{1, false, nil, nil},
				&RBNode{4, false, nil, nil},
			},
			&RBNode{
				10, false,
				&RBNode{9, false, nil, nil},
				&RBNode{20, false, nil, nil},
			},
		}, 0, "Binary tree violation at value 5"},
	}

	for _, te := range table {
		height, err := te.input.Validate()
		errStr := ""
		if err != nil {
			errStr = err.Error()
		}
		if height != te.expectHeight || errStr != te.expectError {
			t.Errorf("Validate(%s) => %v, %q, want %v, %q",
				te.input, height, err, te.expectHeight, te.expectError)
		}
	}
}

func TestInsert(t *testing.T) {
	table := []struct {
		vals   []int
		expect string
	}{
		{[]int{1}, "B:1(,)"},
		{[]int{5, 1, 10}, "B:5(B:1(,),B:10(,))"},
		{[]int{5, 1, 10, 7}, "B:5(B:1(,),B:10(R:7(,),))"},
		{[]int{5, 1, 10, 7, 6}, "B:5(B:1(,),B:7(R:6(,),R:10(,)))"},
		{[]int{1, 5, 6, 7, 10}, "B:5(B:1(,),B:7(R:6(,),R:10(,)))"},
	}

	for _, te := range table {
		var root *RBNode
		for _, val := range te.vals {
			root = root.Insert(val)
		}
		if root.String() != te.expect {
			t.Errorf("Insert() => %q, want %q", root, te.expect)
		}
		if _, err := root.Validate(); err != nil {
			t.Errorf("Insert() invalid: %s", err)
			t.Logf("  Tree: %s", root)
		}
	}
}
