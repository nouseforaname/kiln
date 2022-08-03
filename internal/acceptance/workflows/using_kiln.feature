Feature: Using the Kiln CLI
  Scenario: I ask for the version
    When I invoke kiln version
    Then stdout contains substring: 1.0.0-dev
  Scenario: I ask for help
    When I invoke kiln help
    Then stdout contains substring: kiln helps you
    And stdout contains substring: Usage:
    And stdout contains substring: Commands:

  Scenario: I mess up my command name
    When I try to invoke kiln boo-boo
    # TODO: in this case we should expect output on stderr not stdout
    Then stderr contains substring: unknown command
    And the exit code is 1

  Scenario Outline: I mess up my command flags
    When I try to invoke kiln <command> --boo-boo
    Then the exit code is 1
    And stderr contains substring: flag provided but not defined

    Examples:
      | command                 |
      | bake                    |
      | cache-compiled-releases |
      | fetch                   |
      | find-release-version    |
      | find-stemcell-version   |
      | publish                 |
      | release-notes           |
      | sync-with-local         |
      | update-release          |
      | update-stemcell         |
      | upload-release          |
      | validate                |