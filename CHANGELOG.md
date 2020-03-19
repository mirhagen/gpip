# Changelog

## v1.0.1

### Fixes

* Added `syscall.SIGTERM` to signals in shutdown sequence. This allows for proper graceful handling when terminating the server alongside existing listener on `os.Interrupt`.

## v1.0.0

Initial release.
