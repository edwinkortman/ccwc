Feature: Count bytes in a given file

  Scenario: "test/data/text.txt"
    When I run "go run ccwc -m test/data/text.txt"
    Then the result should contain "339292 test/data/text.txt"