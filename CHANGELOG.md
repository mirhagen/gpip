# Changelog

## v1.1.1

### Fixes

* Correctly returns header `Content-Type: application/json; charset=UTF-8` when `Accept` is not provided or `Accept: */*`.

## v1.1.0

* Added package `ip`. This exports the function `Resolve()` which will check `r *http.Request`
for the headers: `Forwarded`, `X-Forwarded-For` and `X-Real-IP`. Replaces old `utils.go`.

## v1.0.1

### Fixes

* Added `syscall.SIGTERM` to signals in shutdown sequence. This allows for proper graceful handling when terminating the server alongside existing listener on `os.Interrupt`.

## v1.0.0

Initial release.
