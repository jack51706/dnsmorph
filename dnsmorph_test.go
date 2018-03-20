package main

import "testing"

/*
Helper functions tests
*/
func TestCountChar(t *testing.T) {
	count := countChar("test")
	if len(count) != 3 {
		t.Error("expected map keys lenght of 4, got", len(count))
	}
	if count['t'] != 2 {
		t.Error("expected count['t'] to be 2, got", count['t'])
	}
	if count['e'] != 1 {
		t.Error("expected count['t'] to be 1, got", count['t'])
	}
}

func TestProcessInput(t *testing.T) {
	sanitizedInput, tld := processInput("subdomain.test.co.uk")
	if sanitizedInput != "test" && tld != "co.uk" {
		t.Error("expected 'test' and 'co.uk', got", sanitizedInput, tld)
	}
}

/*
Permutation attack functions tests
*/

type testcase struct {
	testString string
	function func(string)[]string
	expectedResultLen int
	firstResult string
  }

  var tests = []testcase {
	  {"test", transpositionAttack, 3, "etst"},
	  {"test", additionAttack, 26, "testa"},
	  {"test", vowelswapAttack, 5, "tast"},
	  {"test", subdomainAttack, 3, "t.est"},
	  {"test", replacementAttack, 31, "6est"},
	  {"test", repetitionAttack, 4, "ttest"},
	  {"test", omissionAttack, 4, "est"},
	  {"test", hyphenationAttack, 3, "t-est"},
	  {"test", bitsquattingAttack, 31, "test"},
	  {"test", homographAttack, 27, "τest"},
  }

  func TestAttackResults(t *testing.T) {
	  for _, test := range tests {
		  results := test.function(test.testString)
		  if len(results) != test.expectedResultLen {
			  t.Errorf("expected array of lenght %d, got %d", test.expectedResultLen, len(results))
		  }
		  if results[0] != test.firstResult {
			t.Errorf("expected first element of array to be '%s', got %s", test.firstResult, results[0])
		  }
	  }
  }