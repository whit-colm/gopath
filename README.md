# gopath
A tool that will make a Go workspace for a project

---

This command line tool will make a go workspace for a new project. Just run it and it will make a path in your current working directory

What is created:
- `./src/`
- `./bin/`
- `setup`

It will only create any of those that are missing.

---

## USAGE:

How to:
1. Chmod executable perms onto the file.
2. Put the file wherever and link a bash alias to it.
3. Run.
  - It will create all missing directories and files and then source
    setup (`. setup`)

---

## IMPORTANT:

setup-colored will only run properly if you have defined the following
escape sequence colors:
  - $c_red
  - $c_green
  - $c_base01
  - $c_blue
  - $c_cyan
  - $c_reset
  - $c_orange
These are all defined on the SkilStak Colors module.

If you do not have the above defined and don't care enough to define
them, instead of making an alias to setup-colored, make it to
setup-blank.

---

# EOF
