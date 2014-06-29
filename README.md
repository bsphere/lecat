lecat
=====

Command line tool to dump stdin to Logentries

Installation
------------

```shell
go get github.com/bsphere/lecat
go install github.com/bsphere/lecat
```

Usage
-----
Add a new manual TCP token log at [logentries.com](https://logentries.com/quick-start/) and copy the [token](https://logentries.com/doc/input-token/).


Usage examples:

```shell
tail -F /var/log/some_logfile.log | lecat -token <logentries_token> -level <optional_log_level>
```