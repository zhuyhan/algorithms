package trie

import "testing"

func testTrie(word []string) *Trie {
	trie := NewTrie()
	for _, v := range word {
		trie.Add(v)
	}
	return trie
}

func TestTrie_Search(t1 *testing.T) {
	data := []string{"abc", "dog", "dogger", "abb", "a"}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields *Trie
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name:   "查找",
			fields: testTrie(data),
			args: args{
				word: "dog",
			},
			want: true,
		},
		{
			name:   "查找",
			fields: testTrie(data),
			args: args{
				word: "a",
			},
			want: true,
		},
		{
			name:   "查找",
			fields: testTrie(data),
			args: args{
				word: "aaa",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Trie{
				root: tt.fields.root,
			}
			if got := t.Search(tt.args.word); got != tt.want {
				t1.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrie_Remove(t1 *testing.T) {
	data := []string{"abc", "dog", "dogger", "abb", "a", "汉语"}
	trie := testTrie(data)
	trie.Remove("abc")
	if trie.Search("abc") == true {
		t1.Error("删除abc失败")
	}
	if trie.Search("abb") == false {
		t1.Error("删除abc失败")
	}
	if trie.Search("a") == false {
		t1.Error("删除abc失败")
	}
}

func TestTrie_StartWith(t1 *testing.T) {
	data := []string{"abc", "dog", "dogger", "abb", "a"}

	type args struct {
		prefix string
	}
	tests := []struct {
		name   string
		fields *Trie
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name:   "查找",
			fields: testTrie(data),
			args: args{
				prefix: "dog",
			},
			want: true,
		},
		{
			name:   "查找",
			fields: testTrie(data),
			args: args{
				prefix: "doggg",
			},
			want: false,
		},
		{
			name:   "查找",
			fields: testTrie(data),
			args: args{
				prefix: "sd",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Trie{
				root: tt.fields.root,
			}
			if got := t.StartWith(tt.args.prefix); got != tt.want {
				t1.Errorf("StartWith() = %v, want %v", got, tt.want)
			}
		})
	}
}
