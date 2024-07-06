# Replace Sample

~~~bash task:ebrp:embed-replace-test
#embed> ebrp:base
~~~

- If you want to replace **'replaceable'** task, define new task as same group and task name.

~~~bash task:ebrp:replaceable -- [hidden]
echo 'Already replaced.'
~~~


## base file

- Base task

~~~bash task:ebrp:base -- [hidden]
echo hello
#embed> ebrp:replaceable
echo world
~~~

- Define new weak task that may be replaced, and embed to base task.

~~~bash task:ebrp:replaceable -- [hidden weak]
echo 'This text is replaceable.'
~~~
