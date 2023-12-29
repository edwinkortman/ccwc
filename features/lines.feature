Feature: Count bytes in a given file

  Scenario: "test/data/text.txt"
    When I run "go run ccwc -l test/data/text.txt"
    Then the result should contain "7145 test/data/text.txt"