package internal

import (
	"testing"
)

type stripHTMLCase struct {
	name           string
	html           string
	expectedOutput string
}

var stripHTMLCases = []stripHTMLCase{
	{"No tags in html text", "This is a simple text string.", "This is a simple text string."},
	{"Strips out tags that aren't replaced by anything", "<p>This is the <em>first</em> paragraph</p><p>This is the <hr>second<hr> paragraph</p>", "This is the first paragraph\nThis is the second paragraph\n"},
	{"Strips <p> tags and replaces </p> with new line character.", "<p>This is the first paragraph</p><p>This is the second paragraph</p>", "This is the first paragraph\nThis is the second paragraph\n"},
	{"Strips <strong> tags and replaces </strong> with ':' character.", "<p><strong>Special</strong> You can select this feat twice.</p>", "Special: You can select this feat twice.\n"},
	{"Strips <strong> tags and does not replace </strong> with ':' character if colon already exists next.", "<p><strong>Special</strong>: You can select this feat twice.</p>", "Special: You can select this feat twice.\n"},
	{"Formats a list with each item having a separate line and new lines around the list", "<p>This is the heading of the list</p>><ul><li><strong>First</strong>: This is the very first item.</li><li><strong>Second</strong>: This is the second item.</li><li><strong>Third</strong>: This is the last item.</li></ul>", "This is the heading of the list\n\n- First: This is the very first item.\n- Second: This is the second item.\n- Third: This is the last item.\n\n"},
}

func TestStripHTML(t *testing.T) {
	for _, tc := range stripHTMLCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := StripHTML(tc.html)
			if tc.expectedOutput != actual {
				t.Fatalf("Expected ouput text doesnt actual.\nActual: '%s'\nExpected: '%s'\n", actual, tc.expectedOutput)
			}
		})
	}
}

type kebabCaseTestCase struct {
	name          string
	input         string
	expected      string
	expectedError bool
}

var kebabCaseTestCases = []kebabCaseTestCase{
	{"Empty string returns an error", "", "", true},
	{"Single uppercase letter string returns lowercase string", "A", "a", false},
	{"Single lowercase letter string returns same string", "a", "a", false},
	{"Single non-letter string returns input as is", "1", "1", false},
	{"Single word string starting with uppercase returns word with first letter lowercase", "Test", "test", false},
	{"Single word string starting with lower case returns same string", "test", "test", false},
	{"Multiple words are returns with first letter of each word lower case and spaces changed to hyphens", "Test string 1", "test-string-1", false},
	{"Any non-letter characters are just concatenated to the output string", "Four Words With d1g17s", "four-words-with-d1g17s", false},
	{"All letters are lowercase in output string", "Four Words With d1g17s", "four-words-with-d1g17s", false},
	{"All letters are lowercase in already kebab case in string", "Test-StrING", "test-string", false},
}

func TestTitleToKebab(t *testing.T) {
	for _, tc := range kebabCaseTestCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := KebabCase(tc.input)

			if err != nil && !tc.expectedError {
				t.Fatalf("Expected to not get error. Got err: %s\n", err)

			} else if tc.expected != actual {
				t.Fatalf("Actual does not match expected. Actual: %s\nExpected: %s\n", actual, tc.expected)
			}
		})
	}
}

type titleCaseTestCase struct {
	name          string
	input         string
	expected      string
	expectedError bool
}

var titleCaseTestCases = []titleCaseTestCase{
	{"Empty string returns an error", "", "", true},
	{"Single lowercase letter string returns uppercase string", "a", "A", false},
	{"Single Uppercase letter string returns same string", "A", "A", false},
	{"Single non-letter string returns input as is", "1", "1", false},
	{"Single word string starting with lowercase returns word with first letter uppercase", "test", "Test", false},
	{"Single word string starting with upperr case returns same string", "Test", "Test", false},
	{"Multiple words are returns with first letter of each word upper case and hypens changes to spaces", "test-string-1", "Test String 1", false},
	{"Any non-letter characters are just concatenated to the output string", "Four Words-with d1g17s", "Four Words With D1g17s", false},
	{"All letters are lowercase in output string expect first of each word", "111four-words-with-digits111", "111four Words With Digits111", false},
}

func TestTitleCase(t *testing.T) {
	for _, tc := range titleCaseTestCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := TitleCase(tc.input)

			if err != nil && !tc.expectedError {
				t.Fatalf("Expected to not get error. Got err: %s\n", err)
			} else if tc.expected != actual {
				t.Fatalf("Actual does not match expected. Actual: %s\nExpected: %s\n", actual, tc.expected)
			}
		})
	}
}

type compendiumEntryFromStringTestCase struct {
	name          string
	input         string
	expected      CompendiumEntry
	expectedError bool
}

var compendiumEntryFromStringTestCases = []compendiumEntryFromStringTestCase{
	{"Malformed compendium string has no '.' characters", "CompendiumPf2eFeatsSrdBreathControl", CompendiumEntry{}, true},
	{"Malformed compendium string has too few '.' characters", "Compendium.pf2e.feats-srd.Breath Control", CompendiumEntry{}, true},
	{"Compendium with 5 parts sets type and value fields", "Compendium.pf2e.feats-srd.Item.Breath Control", CompendiumEntry{"feats-srd", "", "Breath Control"}, false},
	{"Compendium with 7 parts sets type, parentid, and value fields", "Compendium.pf2e.journals.JournalEntry.S55aqwWIzpQRFhcq.JournalEntryPage.pBS3DUjlzVuFgapv", CompendiumEntry{"journals", "S55aqwWIzpQRFhcq", "pBS3DUjlzVuFgapv"}, false},
}

func TestCompendiumEntryFromString(t *testing.T) {
	for _, tc := range compendiumEntryFromStringTestCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := CompendiumEntryFromString(tc.input)
			if tc.expectedError {
				if err == nil {
					t.Fatalf("Expected to get error but got no error\n")
				}
			} else {
				if err != nil {
					t.Fatalf("Expected to not get error. Got err: %s\n", err)
				}
				if tc.expected.Type != actual.Type || tc.expected.ParentID != actual.ParentID || tc.expected.Value != actual.Value {
					t.Fatalf("Actual does not match expected. Actual: %v\nExpected: %v\n", actual, tc.expected)
				}
			}
		})
	}
}

type compendiumEntryFromTagStringTestCase struct {
	name          string
	input         string
	expected      CompendiumEntry
	expectedError bool
}

var compendiumEntryFromTagStringTestCases = []compendiumEntryFromTagStringTestCase{
	{"Malformed tag compendium string has no '@' character at the start", "UUID[Compendium.pf2e.feats-srd.Item.Breath Control]", CompendiumEntry{}, true},
	{"Malformed tag compendium string has ']' and '[' characters in the wrong order", "@UUID]Compendiumpf2efeats-srdItemBreath Control[", CompendiumEntry{}, true},
	{"Malformed tag compendium string has missing ']' character", "@UUID[Compendium.pf2e.feats-srd.Breath Control", CompendiumEntry{}, true},
	{"Malformed tag compendium string has missing '[' character", "@UUIDCompendium.pf2e.feats-srd.Breath Control]", CompendiumEntry{}, true},
	{"Compendium with 5 parts sets type and value fields", "@UUID[Compendium.pf2e.feats-srd.Item.Breath Control]", CompendiumEntry{"feats-srd", "", "Breath Control"}, false},
	{"Compendium with 7 parts sets type, parentid, and value fields", "@UUID[Compendium.pf2e.journals.JournalEntry.S55aqwWIzpQRFhcq.JournalEntryPage.pBS3DUjlzVuFgapv]", CompendiumEntry{"journals", "S55aqwWIzpQRFhcq", "pBS3DUjlzVuFgapv"}, false},
}

func TestCompendiumEntryFromTagString(t *testing.T) {
	for _, tc := range compendiumEntryFromTagStringTestCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := CompendiumEntryFromTagString(tc.input)
			if tc.expectedError {
				if err == nil {
					t.Fatalf("Expected to get error but got no error\n")
				}
			} else {
				if err != nil {
					t.Fatalf("Expected to not get error. Got err: %s\n", err)
				}
				if tc.expected.Type != actual.Type || tc.expected.ParentID != actual.ParentID || tc.expected.Value != actual.Value {
					t.Fatalf("Actual does not match expected. Actual: %v\nExpected: %v\n", actual, tc.expected)
				}
			}
		})
	}
}
