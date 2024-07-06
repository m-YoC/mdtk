
# Embedded Comment Test

- Terminal Tasks

~~~bash task:eb:t -- [hidden] embedded comment test termination
echo hello
~~~

~~~bash task:eb:tco -- [hidden] embedded comment test termination (config once)
#config> once
echo hello
~~~

~~~bash task:eb:targ -- [hidden] embedded comment test termination (with args)
#args> a:(string)
: ${a:='default'}
echo "hello / arg: $a"
~~~

## `#embed>`

- Simple Embedding

~~~bash task:eb:embed-test
#embed> eb:t
~~~

## `#task>`

- Embedding into Subshell

~~~bash task:eb:task-test
#task> eb:t
~~~

- with args

~~~bash task:eb:taskarg-test
#task> eb:targ -- a=task-arg
~~~

## `#func>`

- Embedding as a Function

~~~bash task:eb:func-test
#func> tt eb:t
tt
~~~

- with args

~~~bash task:eb:funcarg-test
#func> tt eb:targ -- a=func-arg
tt
~~~

- with special parameter (required positional parameter type) 
    - Set `a=$x`

~~~bash task:eb:funcarg2-test
#func> tt eb:targ -- a={$}
echo '* with positional parameter'
tt func-arg || true
echo '* no positional parameter'
tt || true
~~~

- with special parameter (optional positional parameter type) 
    - Set `a=${x-''}`

~~~bash task:eb:funcarg3-test
#func> tt eb:targ -- a={?}
echo '* with positional parameter'
tt func-arg || true
echo '* no positional parameter'
tt || true
~~~

## `#config> once`

- The second time is not embedded

~~~bash task:eb:once1-test
echo '* first'
#embed> eb:tco
echo '* second'
#embed> eb:tco
echo 'end'
~~~

- The second time is also embedded
    - Once flag is temporarily reset in `#task>`

~~~bash task:eb:once2-test
echo '* first'
#task> eb:tco
echo '* second'
#task> eb:tco
echo 'end'
~~~

- The second time is not embedded
    - Once flag is not reset in `#task> @`

~~~bash task:eb:once2b-test
echo '* first'
#task> @ eb:tco
echo '* second'
#task> @ eb:tco
echo 'end'
~~~

- The second time is also embedded
    - Once flag is temporarily reset in `#func>`

~~~bash task:eb:once3-test
echo '* first'
#func> tt eb:tco
tt
echo '* second'
#func> tt2 eb:tco
tt2
echo 'end'
~~~

## `#desc>`

- For display in task help

~~~bash task:eb:desc-test
#desc> This is a description of '#desc>'.
#desc> 2nd row.
echo hello
~~~

## `#args>`

- For display in task help

~~~bash task:eb:args-test
#args> a:(xxx) b:(xxx)
#desc> Text of '#args>' is simply a type of task help description.
echo hello
~~~
