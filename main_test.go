package main

import "testing"

func TestReverse(t *testing.T) {
	exec := func(in, want []string) {
		in2 := in[:]
		t.Log("before reverse:", in2)
		reverse(in2)
		t.Log("after reverse:", in2)
		t.Log("want:", want)

		for i := range in2 {
			if in2[i] != want[i] {
				t.Error("配列が逆順になっていません。")
			}
		}
	}

	exec([]string{}, []string{})
	exec([]string{"1"}, []string{"1"})
	exec([]string{"1", "2"}, []string{"2", "1"})
	exec([]string{"1", "2", "3"}, []string{"3", "2", "1"})
	exec([]string{"1", "2", "3", "4"}, []string{"4", "3", "2", "1"})
}
