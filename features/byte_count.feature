Feature: Count bytes in a given file

  Scenario: "test/data/bytecount.txt"
    When I run "go run ccwc -c test/data/bytecount.txt"
    Then the result should contain "342190 test/data/bytecount.txt"