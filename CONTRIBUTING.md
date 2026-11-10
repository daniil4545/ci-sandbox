# Contributing

- Every exported name has a doc comment.
- Errors are compared with `errors.Is`, never by their text.
- Every behaviour change comes with a test.
- Standard library only. A new dependency needs the owner's consent in its issue.
- `calc` works on `int`: no floats, no generics.
