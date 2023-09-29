# run-external-cmd
Tired of `system`, `popen`, `fork`, `exec`, and `posix_spawn`. Starting process should be made simple.

E.g.
```
run-external-cmd gvim file
```
```
run-external-cmd -hex 6776696D 66696C65
```

INSTALLATION COMMAND: `go install 'github.com/cshu/run-external-cmd@latest'`

WHEN IS -hex USEFUL: e.g. when writing a c program, a simple `system()` might have some issue with special chars. In this case `system` can call `run-external-cmd` with all args written in hex instead.

WHEN IS -wait-stdin USEFUL: e.g. a parent process uses this program to launch a child process, and later parent process decides to kill the child. It can do so by close the stdin pipe.
