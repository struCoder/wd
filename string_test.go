package wd

import (
  "reflect"
  "testing"
)

func TestAllWhitespace(t *testing.T) {
  type testCase struct {
    in   string
    want bool
  }
  testCases := []testCase{
    {"", true},
    {" ", true},
    {"a", false},
    {" a b c ", false},
    {"\u00a0", true},
    {" \u00a0 ", true},
  }
  for _, tc := range testCases {
    got := allWhitespace(tc.in)
    if got != tc.want {
      t.Errorf("allWhitespace(%q) = %v, want %v", tc.in, got, tc.want)
    }
  }
}

func TestTokenize(t *testing.T) {
  type testCase struct {
    in   string
    want []string
  }
  testCases := []testCase{
    {"", nil},
    {"a", []string{"a"}},
    {"abc", []string{"abc"}},
    {" ", []string{" "}},
    {"        ", []string{"        "}},
    {" abc", []string{" ", "abc"}},
    {"abc ", []string{"abc", " "}},
    {"!", []string{"!"}},
    {"!?!", []string{"!?!"}},
    {"<em>", []string{"<", "em", ">"}},
    {" ?", []string{" ", "?"}},
    {"? ", []string{"?", " "}},
    {"abc!123", []string{"abc", "!123"}},
    {"aaa        bbb", []string{"aaa", "        ", "bbb"}},
    {" \u00a0 ", []string{" \u00a0 "}},
    {
      "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
      []string{
        "Lorem", " ", "ipsum", " ", "dolor", " ", "sit", " ", "amet", ",", " ",
        "consectetur", " ", "adipiscing", " ", "elit",
      },
    },
    {
      "Иногда простым трёхчасовым ацетилированием простой домашней сметаны в фтороводородной среде с палладиевым катализом можно добиться удивительных результатов.",
      []string{
        "Иногда", " ", "простым", " ", "трёхчасовым", " ", "ацетилированием", " ",
        "простой", " ", "домашней", " ", "сметаны", " ", "в", " ", "фтороводородной",
        " ", "среде", " ", "с", " ", "палладиевым", " ", "катализом", " ", "можно", " ",
        "добиться", " ", "удивительных", " ", "результатов", ".",
      },
    },
  }
  for i, tc := range testCases {
    got := tokenize(tc.in)
    if !reflect.DeepEqual(got, tc.want) {
      t.Errorf("#%d: want %#v, got %#v", i+1, tc.want, got)
    }
  }
}

func TestNumberOfDigits(t *testing.T) {
  type testCase struct {
    in   int
    want int
  }
  testCases := []testCase{
    {0, 1}, {1, 1}, {2, 1}, {3, 1}, {4, 1}, {5, 1}, {6, 1}, {7, 1}, {8, 1}, {9, 1}, {10, 2},
    {99, 2}, {100, 3}, {999, 3}, {1000, 4}, {9999, 4}, {10000, 5}, {99999, 5}, {100000, 6},
  }
  for _, tc := range testCases {
    got := numberOfDigits(tc.in)
    if got != tc.want {
      t.Errorf("numberOfDigits(%d) = %d, want %d", tc.in, got, tc.want)
    }
  }
}

func TestNumberLines(t *testing.T) {
  type testCase struct {
    in   string
    want string
  }
  testCases := []testCase{
    {"", "1 | "},
    {"abc", "1 | abc"},
    {"abc\ndef\nghi", "1 | abc\n2 | def\n3 | ghi"},
    {"\n\n\n", "1 | \n2 | \n3 | \n4 | "},
    {"a\nb\nc\nd\ne\nf\ng\nh\ni\nj", " 1 | a\n 2 | b\n 3 | c\n 4 | d\n 5 | e\n 6 | f\n 7 | g\n 8 | h\n 9 | i\n10 | j"},
  }
  for _, tc := range testCases {
    got := NumberLines(tc.in)
    if got != tc.want {
      t.Errorf("NumberLines(%q) = %q, want %q", tc.in, got, tc.want)
    }
  }
}
