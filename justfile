dev:
    watchexec -r -e go "go run ." & disown

test:
    hurl test.hurl

clean:
    killall watchexec
