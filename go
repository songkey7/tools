#!/usr/bin/expect

if {$argc < 1} {
    send_user "usage: $argv0 param.\n"
    exit
}

set timeout -1
set jump_host user@localhost
set cmd [lindex $argv 0]

# no need jump box
if { $cmd == "38" } {
    spawn ssh -t qsong@192.168.0.38
    interact
    exit
}

# via jump box
if { $cmd == "0" } {
    set host user1@host1
}

spawn ssh -t $jump_host

expect "$" { send "ssh -t $host\r" }

interact

