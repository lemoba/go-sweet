package framework

import "testing"

func Test_filterChildNodes(t *testing.T) {
	root := &node{
		isLast:   false,
		segment:  "",
		handlers: nil,
		childs: []*node{
			{
				isLast:   true,
				segment:  "FOO",
				handlers: nil,
				childs:   nil,
			},
			{
				isLast:   false,
				segment:  ":id",
				handlers: nil,
				childs:   nil,
			},
		},
	}

	{
		nodes := root.filterChildNodes("FOO")
		if len(nodes) != 2 {
			t.Error("foo error")
		}
	}

	{
		nodes := root.filterChildNodes(":foo")
		if len(nodes) != 2 {
			t.Error(":foo error")
		}
	}
}

func Test_matchNode(t *testing.T) {
	root := &node{
		isLast:   false,
		segment:  "",
		handlers: nil,
		childs: []*node{
			{
				isLast:   true,
				segment:  "FOO",
				handlers: nil,
				childs: []*node{
					&node{
						isLast:   true,
						segment:  "BAR",
						handlers: nil, //func(*Context) error { panic("not implemented") },
						childs:   []*node{},
					},
				},
			},
			{
				isLast:   true,
				segment:  ":id",
				handlers: nil,
				childs:   nil,
			},
		},
	}

	{
		node := root.matchNode("foo/bar")
		if node == nil {
			t.Error("match normal node error")
		}
	}

	{
		node := root.matchNode("test.go")
		if node == nil {
			t.Error("match test.go")
		}
	}
}
