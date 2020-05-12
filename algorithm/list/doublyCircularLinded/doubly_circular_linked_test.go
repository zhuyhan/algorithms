package doublyCircularLinded

import (
	"reflect"
	"testing"
)

func testLinked(list []int) *List {
	linked := New()
	for _, v := range list {
		linked.Add(v)
	}

	return linked
}

func TestList_GetAll(t *testing.T) {
	nums := []int{1, 4, 7, 9}
	list := testLinked(nums)
	list.Insert(3, 99)
	list.Insert(0, 199)
	list.Insert(6, 299)

	nums2 := []int{1, 4, 7, 9, 199, 299}
	list2 := testLinked(nums2)
	list2.Remove(6)
	list2.Remove(5)
	list2.Remove(0)
	tests := []struct {
		name   string
		fields *List
		want   []int
	}{
		// TODO: Add test cases.
		{
			name:   "查询",
			fields: list,
			want:   []int{199, 1, 4, 7, 99, 9, 299},
		},
		{
			name:   "查询",
			fields: list2,
			want:   []int{4, 7, 9, 199},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				first: tt.fields.first,
				last:  tt.fields.last,
				len:   tt.fields.len,
			}
			if got := l.GetAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
