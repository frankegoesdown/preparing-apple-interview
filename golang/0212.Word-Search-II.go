package main

func main() {

}

type Trie struct {
	Nodes [26]*Trie
	Word  string
}

func findWords(board [][]byte, words []string) []string {
	var ans []string
	root := buildTrie(words)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			helper(board, root, i, j, &ans)
		}
	}
	return ans
}

func buildTrie(words []string) *Trie {
	root := &Trie{}
	for _, word := range words {
		tmp := root
		for _, c := range word {
			i := c - 'a'
			if tmp.Nodes[i] == nil {
				tmp.Nodes[i] = &Trie{}
			}
			tmp = tmp.Nodes[i]
		}
		tmp.Word = word
	}
	return root
}

func helper(board [][]byte, root *Trie, i, j int, ans *[]string) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return
	}
	tmp := board[i][j]
	if tmp == '$' || root.Nodes[tmp-'a'] == nil {
		return
	}
	root = root.Nodes[tmp-'a']
	if len(root.Word) > 0 {
		*ans = append(*ans, root.Word)
		root.Word = ""
	}
	board[i][j] = '$'
	helper(board, root, i+1, j, ans)
	helper(board, root, i-1, j, ans)
	helper(board, root, i, j+1, ans)
	helper(board, root, i, j-1, ans)
	board[i][j] = tmp
}
