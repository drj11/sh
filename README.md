# Key Challenges

Can we build a pipeline? Note that in a command like:
`ls -s | sort -n | head` the shell never sees the output of any
of the commands. `pipe()` is used to connect `ls` stdout to 
`sort` stdin, and so on.

Can we correctly handle variables for _special built-ins_? For
example: `a=foo : ; echo $a` should display `foo` (it doesn't in
`bash` but don't be caught out by that).

traps just seem tricky. Wonder if they really are.

Job control. Again sounds tricky. But optional.
