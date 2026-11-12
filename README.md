# ci-sandbox

A sandbox for trying out pull request gates before they guard real repositories. The code is a
tiny Go library (`calc`) that exists only so that test pull requests have something to change.

Tested here: the merge gate (CI ends in one required check, `gate`, red unless every job passed)
and [famulus](https://github.com/daniil4545/famulus), a bot that reviews each open PR against its
issue and the rules in `CONTRIBUTING.md`. Some pull requests are broken on purpose.

Checks: `make ci-check`. License: MIT.
