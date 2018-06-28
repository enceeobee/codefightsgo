package regularhell

import "testing"

func TestRepetitionEncryption(t *testing.T) {
	var actual int
	tests := []struct {
		letter string
		x      int
	}{
		{"Hi, hi Jane! I'm so. So glad to to finally be able to write - WRITE!! - to you!", 4},
		{"Yo-yo, how's s it going going for ya? Ya is okay, okay'nt ya?", 5},
		{"Hi Jane, how are you doing today?", 0},
		{"My friend, friend of mine, I am really-really happy for you! You are amazing, please write me back when you can.", 3},
		{"Everything is fine, fine perfectly here. Here I have fun (fun all the day!) days. Although I miss you, so please please, Jane, write, write me whenever you can! Can you do that? That is the only (!!ONLY!!) thing I ask from you, you sunshine.", 9},
		{"Do not notify me about this in the future", 0},
		{"ho-ho--he-he", 2},
		{"nb-nb--he-_he", 2},
	}

	for _, test := range tests {
		actual = repetitionEncryption(test.letter)

		if actual != test.x {
			t.Errorf("repetitionEncryption(%s) = %d; expected %d", test.letter, actual, test.x)
		}
	}
}
