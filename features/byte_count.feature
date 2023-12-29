Feature: Count bytes in a given file

  Scenario: "test/data/text.txt"
    When I run "go run ccwc -c test/data/text.txt"
    Then the result should contain "342190 test/data/text.txt"