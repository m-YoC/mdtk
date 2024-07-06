
Check task help
~~~
mdtk -a
~~~

# Basic Test

- The simplest task definition

~~~bash task::hello
echo hello world
~~~

## group

- Set group name

~~~bash task:group-test:group-test
echo group test
~~~

### group:task conflict

- If group and task name are same, a conflict occurs
    - Can be avoided by using **priority** attribute described below

~~~bash task:group-test:conflict-test
echo group:task conflict test
~~~
~~~bash task:group-test:conflict-test
echo group:task conflict test
~~~

## description & attributes

- Write description

~~~bash task:desc:desc-test -- description test
echo desc test
~~~

- Set **hidden** attribute

~~~bash task:attr:hidden-test -- [hidden] description hidden test
echo hidden test
~~~

- Set **b** (bottom) attribute
    - Be laid out at the bottom of the group to which it belongs 

~~~bash task:attr:bottom-test -- [b] This test is laid out at the bottom of this group 
echo bottom layout test
~~~

- Set **t** (top) attribute
    - Be laid out at the top of the group to which it belongs 

~~~bash task:attr:top-test -- [t] This test is laid out at the top of this group 
echo top layout test
~~~

- Set **priority**
    - Priority is integer, range of [-9, 9]

~~~bash task:attr:priority-test -- [priority:5] description
echo 'priority test (+5)'
~~~
~~~bash task:attr:priority-test -- [priority:-5] description
echo 'priority test (-5)'
~~~

- Set **weak**
    - A type of priority

~~~bash task:attr:weak-test -- [weak] description
echo 'weak test (weak)'
~~~
~~~bash task:attr:weak-test -- [] description
echo 'weak test (not weak)'
~~~

## other language

- Can write other languages.
    - ***Cannot run***
    - You can use some embedded comments

~~~go task:go:hello-go -- <-language
fmt.Println("Hello Go World")
~~~


# group order

- Can change display order of groups at task help 

~~~taskconfig:group-order
group-test: 9
desc: 8
attr: 7
~~~

# combine sub taskfile

- Can read and combine other taskfile

~~~taskfile
embed-sample.task.md
embed-replace-sample.task.md
~~~
