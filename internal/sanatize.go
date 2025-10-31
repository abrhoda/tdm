package internal

import (
	"fmt"
	"strings"
)

func KebabCase(in string) (string, error) {
	l := len(in)
	if l == 0 {
		return in, fmt.Errorf("Cannot convert string of length 0 to kebab case.")
	}

	
	var b strings.Builder
	b.Grow(l)

	current := 0
	for current < l {
		if in[current] == ' ' {
			b.WriteByte('-')
		} else {
			if 65 <= in[current] && in[current] <= 90 {
				b.WriteByte(in[current]+32)
			} else {
				b.WriteByte(in[current])
			}
		} 
		current++
	}

	return b.String(), nil
}

func TitleCase(in string) (string, error) {
	l := len(in)
	if l == 0 {
		return in, fmt.Errorf("Cannot convert string of length 0 to title case.")
	}
	
	var b strings.Builder
	b.Grow(l)
	
	current := 0
	beginning := true
	for current < l {
		if in[current] == '-' || in[current] == ' ' {
			b.WriteByte(' ')
			beginning = true
		} else if beginning {
			if 97 <= in[current] && in[current] <= 122 {
				b.WriteByte(in[current]-32)
			} else {
				b.WriteByte(in[current])
			}
			beginning = false
		} else if 65 <= in[current] && in[current] <= 90 {
			b.WriteByte(in[current]+32)
		} else {
			b.WriteByte(in[current])
		}
		current++
	}

	return b.String(), nil
}

type CompendiumEntry struct {
	Type string
	ParentID string
	Value string
}

func CompendiumEntryFromString(in string) (CompendiumEntry, error) {
	split := strings.Split(in, ".")
	out := CompendiumEntry{}
	if len(split) != 5 && len(split) != 7 {
		return out, fmt.Errorf("CompendiumCategoryAndValue split was not 5/7 parts. In value was \"%s\"\n", in)
	}
	
	out.Type = split[2]
	if len(split) == 5 {
		out.Value = split[4]
	} else {
		out.ParentID = split[4]
		out.Value = split[6]
	}

	return out, nil
}

// @UUID[...] could be followed by `{condition value}` token
func CompendiumEntryFromTagString(in string) (CompendiumEntry, error) {
	if in[0] != '@' {
		return CompendiumEntry{}, fmt.Errorf("In string does not start with '@'\n")
	}
	o := strings.Index(in, "[")
	c := strings.Index(in, "]")
	if o == -1 || c == -1 {
		return CompendiumEntry{}, fmt.Errorf("In string does not have opening and closing brackets.\n")
	}
	
	if c < o {
		return CompendiumEntry{}, fmt.Errorf("In string has closing bracket before opening bracket.\n")
	}
	
	return CompendiumEntryFromString(in[o+1:c])
}

// In order to clean up some text while stripping out html tags:
// replace </p> with \n.
// replace <ul> and <\ul> with \n
// replace <li> with -
// replace </li> with \n
// replace </strong> with :
func StripHTML(in string) string {
	inTag := false
	start := 0
	current := 0

	var b strings.Builder
	b.Grow(len(in))

	for current < len(in) {
		c := in[current]
		if c == '<' {
			start = current
			inTag = true
		} else if c == '>' {
			tag := in[start+1:current]
			switch tag {
			case "/p", "ul", "/ul", "/li":
				b.WriteByte('\n')
			case "li":
				b.WriteByte('-')
				b.WriteByte(' ')
			case "/strong":
				// could probably blindly assume that no closing strong tag is at the end of a string but bounds check because who knows with this insane dataset.
				if current + 1 < len(in) && in[current+1] != ':' {
					b.WriteByte(':')
				}
			}
			inTag = false
		} else if !inTag {
			b.WriteByte(c)
		}
		current++
	}

	return b.String()
}
