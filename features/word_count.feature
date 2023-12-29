Feature: Count bytes in a given file

  Scenario: "test/data/text.txt"
    When I run "go run ccwc -w test/data/text.txt"
    Then the result should contain "58164 test/data/text.txt"