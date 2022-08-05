Feature: Subtraction happens with addition

  Scenario:
    Given calculator is cleared
    When i press 5
    And i add 9
    And i subtract 4
    And i press 32
    And i add 2
    And i multiply by 2
    Then the result should be 68

 Scenario:
    Given the calculator is ready
    When i add 20
    And i subtract 3
    And i subtract 2
    And i multiply by 4
    Then the result should be 60