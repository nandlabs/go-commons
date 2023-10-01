package string

import "testing"

func TestToLower(t *testing.T) {
	input := "Hello, World!"
	expected := "hello, world!"
	result := ToLower(input)

	if result != expected {
		t.Errorf("ToLower(%s) = %s; expected %s", input, result, expected)
	}
}

func TestToUpper(t *testing.T) {
	input := "Hello, World!"
	expected := "HELLO, WORLD!"
	result := ToUpper(input)

	if result != expected {
		t.Errorf("ToUpper(%s) = %s; expected %s", input, result, expected)
	}
}

func TestConcat(t *testing.T) {
	s1 := "Hello, "
	s2 := "World!"
	expected := "Hello, World!"
	result := Concat(s1, s2)

	if result != expected {
		t.Errorf("Concat(%s, %s) = %s; expected %s", s1, s2, result, expected)
	}
}

func TestReverse(t *testing.T) {
	input := "Hello, World!"
	expected := "!dlroW ,olleH"
	result := Reverse(input)

	if result != expected {
		t.Errorf("Reverse(%s) = %s; expected %s", input, result, expected)
	}
}

func TestReverseEmptyString(t *testing.T) {
	input := ""
	expected := ""
	result := Reverse(input)

	if result != expected {
		t.Errorf("Reverse(%s) = %s; expected %s", input, result, expected)
	}
}

func TestReverseUnicode(t *testing.T) {
	input := "こんにちは"
	expected := "はちにんこ"
	result := Reverse(input)

	if result != expected {
		t.Errorf("Reverse(%s) = %s; expected %s", input, result, expected)
	}
}

func TestContains(t *testing.T) {
	input := "Hello, World!"
	substring := "World"
	expected := true
	result := Contains(input, substring)

	if result != expected {
		t.Errorf("Contains(%s, %s) = %v; expected %v", input, substring, result, expected)
	}
}

func TestCountSubstring(t *testing.T) {
	input := "Hello, Hello, Hello"
	substring := "Hello"
	expected := 3
	result := CountSubstring(input, substring)

	if result != expected {
		t.Errorf("CountSubstring(%s, %s) = %d; expected %d", input, substring, result, expected)
	}
}

func TestReplaceAll(t *testing.T) {
	input := "Hello, World!"
	old := "World"
	newStr := "Universe"
	expected := "Hello, Universe!"
	result := ReplaceAll(input, old, newStr)

	if result != expected {
		t.Errorf("ReplaceAll(%s, %s, %s) = %s; expected %s", input, old, newStr, result, expected)
	}
}

func TestTrimWhitespace(t *testing.T) {
	input := "   Hello, World!   "
	expected := "Hello, World!"
	result := TrimWhitespace(input)

	if result != expected {
		t.Errorf("TrimWhitespace(%s) = %s; expected %s", input, result, expected)
	}
}

func TestSplit(t *testing.T) {
	input := "apple,banana,cherry"
	delimiter := ","
	expected := []string{"apple", "banana", "cherry"}
	result := Split(input, delimiter)

	if len(result) != len(expected) {
		t.Errorf("Split(%s, %s) = %v; expected %v", input, delimiter, result, expected)
	} else {
		for i := range result {
			if result[i] != expected[i] {
				t.Errorf("Split(%s, %s) = %v; expected %v", input, delimiter, result, expected)
				break
			}
		}
	}
}

func TestJoin(t *testing.T) {
	strs := []string{"apple", "banana", "cherry"}
	separator := ","
	expected := "apple,banana,cherry"
	result := Join(strs, separator)

	if result != expected {
		t.Errorf("Join(%v, %s) = %s; expected %s", strs, separator, result, expected)
	}
}

func TestRepeat(t *testing.T) {
	s := "abc"
	count := 3
	expected := "abcabcabc"
	result := Repeat(s, count)

	if result != expected {
		t.Errorf("Repeat(%s, %d) = %s; expected %s", s, count, result, expected)
	}
}

func TestTrimLeft(t *testing.T) {
	s := "   Hello, World!   "
	cutset := " "
	expected := "Hello, World!   "
	result := TrimLeft(s, cutset)

	if result != expected {
		t.Errorf("TrimLeft(%s, %s) = %s; expected %s", s, cutset, result, expected)
	}
}

func TestTrimRight(t *testing.T) {
	s := "   Hello, World!   "
	cutset := " "
	expected := "   Hello, World!"
	result := TrimRight(s, cutset)

	if result != expected {
		t.Errorf("TrimRight(%s, %s) = %s; expected %s", s, cutset, result, expected)
	}
}

func TestFields(t *testing.T) {
	s := "apple banana cherry"
	expected := []string{"apple", "banana", "cherry"}
	result := Fields(s)

	if len(result) != len(expected) {
		t.Errorf("Fields(%s) = %v; expected %v", s, result, expected)
	} else {
		for i := range result {
			if result[i] != expected[i] {
				t.Errorf("Fields(%s) = %v; expected %v", s, result, expected)
				break
			}
		}
	}
}

func TestHasPrefix(t *testing.T) {
	s := "Hello, World!"
	prefix := "Hello"
	expected := true
	result := HasPrefix(s, prefix)

	if result != expected {
		t.Errorf("HasPrefix(%s, %s) = %v; expected %v", s, prefix, result, expected)
	}
}

func TestHasSuffix(t *testing.T) {
	s := "Hello, World!"
	suffix := "World!"
	expected := true
	result := HasSuffix(s, suffix)

	if result != expected {
		t.Errorf("HasSuffix(%s, %s) = %v; expected %v", s, suffix, result, expected)
	}
}

func TestMap(t *testing.T) {
	s := "abc"
	mapping := func(r rune) rune {
		return r + 1
	}
	expected := "bcd"
	result := Map(mapping, s)

	if result != expected {
		t.Errorf("Map(%s) = %s; expected %s", s, result, expected)
	}
}

func TestReplace(t *testing.T) {
	s := "Hello, World!"
	oldnew := []string{"Hello", "Hi"}
	expected := "Hi, World!"
	result := Replace(s, oldnew...)

	if result != expected {
		t.Errorf("Replace(%s, %v) = %s; expected %s", s, oldnew, result, expected)
	}
}

func TestCompare(t *testing.T) {
	s1 := "abc"
	s2 := "def"
	expected := -1
	result := Compare(s1, s2)

	if result != expected {
		t.Errorf("Compare(%s, %s) = %d; expected %d", s1, s2, result, expected)
	}
}

func TestToValidUTF8(t *testing.T) {
	s := "Hello, 世界"
	replacement := "?"
	expected := "Hello, ??"
	result := ToValidUTF8(s, replacement)

	if result != expected {
		t.Errorf("ToValidUTF8(%s, %s) = %s; expected %s", s, replacement, result, expected)
	}
}

func TestTrimPrefix(t *testing.T) {
	s := "Hello, World!"
	prefix := "Hello, "
	expected := "World!"
	result := TrimPrefix(s, prefix)

	if result != expected {
		t.Errorf("TrimPrefix(%s, %s) = %s; expected %s", s, prefix, result, expected)
	}
}

func TestTrimSuffix(t *testing.T) {
	s := "Hello, World!"
	suffix := ", World!"
	expected := "Hello"
	result := TrimSuffix(s, suffix)

	if result != expected {
		t.Errorf("TrimSuffix(%s, %s) = %s; expected %s", s, suffix, result, expected)
	}
}

func TestContainsAny(t *testing.T) {
	s := "Hello, World!"
	chars := "aeiou"
	expected := true
	result := ContainsAny(s, chars)

	if result != expected {
		t.Errorf("ContainsAny(%s, %s) = %v; expected %v", s, chars, result, expected)
	}
}

func TestCount(t *testing.T) {
	s := "Hello, Hello, World!"
	substr := "Hello"
	expected := 2
	result := Count(s, substr)

	if result != expected {
		t.Errorf("Count(%s, %s) = %d; expected %d", s, substr, result, expected)
	}
}

func TestIndex(t *testing.T) {
	s := "Hello, World!"
	substr := "World"
	expected := 7
	result := Index(s, substr)

	if result != expected {
		t.Errorf("Index(%s, %s) = %d; expected %d", s, substr, result, expected)
	}
}

func TestIndexAny(t *testing.T) {
	s := "Hello, World!"
	chars := "aeiou"
	expected := 1
	result := IndexAny(s, chars)

	if result != expected {
		t.Errorf("IndexAny(%s, %s) = %d; expected %d", s, chars, result, expected)
	}
}

func TestLastIndex(t *testing.T) {
	s := "Hello, Hello, World!"
	substr := "Hello"
	expected := 7
	result := LastIndex(s, substr)

	if result != expected {
		t.Errorf("LastIndex(%s, %s) = %d; expected %d", s, substr, result, expected)
	}
}

func TestLastIndexAny(t *testing.T) {
	s := "Hello, World!"
	chars := "aeiou"
	expected := 8
	result := LastIndexAny(s, chars)

	if result != expected {
		t.Errorf("LastIndexAny(%s, %s) = %d; expected %d", s, chars, result, expected)
	}
}
