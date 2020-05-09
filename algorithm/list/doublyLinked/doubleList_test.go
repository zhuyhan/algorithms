package doublyLinked

import (
	"reflect"
	"testing"
)

func TestList_Node(t *testing.T) {
	type fields struct {
		root Element
		len  int
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Element
	}{
		// TODO: Add test cases.
		{
			name: "index所在的位置",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				root: tt.fields.root,
				len:  tt.fields.len,
			}
			if got := l.Node(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node() = %v, want %v", got, tt.want)
			}
		})
	}
}
