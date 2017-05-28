lkml
====

`lkml` is a simple, stupid lkml (Linux Kernel Mailing List) viewer.


Install
=======

Installation is very simple.  Just use the `go get` tool if you already set up
Go development environment as below:

```
$ go get github.com/sjp38/lkml
```


Usage
=====

Usage is very simple.  Just execute the program from the shell.  Output of the
program will be something like below:

```
$ lkml
[PATCH 1/3] regmap: Add OneWire (W1) bus support
        "Alex A. Mihaylov" <invalid@email.com>
        http://lkml.org/lkml/2017/5/28/25

[PATCH] w1: Fix slave count on W1 bus
        "Alex A. Mihaylov" <invalid@email.com>
        http://lkml.org/lkml/2017/5/28/26

Re: [PATCH v2 20/20] ACPICA: Use designated initializers
        Christoph Hellwig <invalid@email.com>
        http://lkml.org/lkml/2017/5/28/27
```

The output is list of brief information about recent mails in LKML that
seperated with blank line.  Each item (brief information about a mail) is
configured with title, author, and link for the mail in each line.  If you have
interest to a specific mail, you may use the link to see the original full
content of the mail.


Options
-------

`lkml` provides few simple options as below.

```
  -count int
        Updates count (default 1)
  -delay int
        Delay between updates in seconds
  -keyword string
        Keyword to be exist in title
```

If `count` and `delay` option is given, the program repeats execution for
`count` times with `delay` seconds interval.  If `count` option value is given
as `-1`, it works as infinite.  You may use the options to see LKML in live.

`keyword` option is for filtering.  If the option is given, the program will
show items containing the keyword in title only.  Because LKML receives tons of
mails each day, it could be helpful for people who have interest in specific
topic only.


License
=======

GPL v3


Author
======

SeongJae Park (sj38.park@gmail.com)
