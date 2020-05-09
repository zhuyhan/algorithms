package sequnce

import (
	"testing"
)

func TestSequenceList_Insert(t *testing.T) {
	type fields struct {
		Len      int
		Capacity int
		Data     []interface{}
	}
	type args struct {
		index int
		elem  interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "插入元素",
			fields: fields{
				Len:      4,
				Capacity: 10,
				Data:     []interface{}{1, 4, 56, 10},
			},
			args: args{
				index: 2,
				elem:  3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			s := &SequenceList{
				Len:      tt.fields.Len,
				Capacity: tt.fields.Capacity,
				Data:     tt.fields.Data,
			}

			if got := s.Insert(tt.args.index, tt.args.elem); got != tt.want {
				t.Errorf("SequenceList.Insert() = %v, want %v", got, tt.want)
			}
			t.Logf("SequenceList %+v", s)

		})
	}
}
