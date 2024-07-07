

# Bash

~~~bash task::test -- [hidden]
#args> task:(group:task)
mdtk -f ./first-sample.taskrun.md $task > /dev/null 2>&1
if [ $? -ne 0 ]; then echo "failed: $task"; fi
~~~

~~~bash task::ntest -- [hidden]
#args> task:(group:task)
mdtk -f ./first-sample.taskrun.md $task > /dev/null 2>&1
if [ $? -ne 1 ]; then echo "failed: $task"; fi
~~~

~~~bash task:autotest:bash-test-run
#task> test -- task=hello
#task> test -- task=group-test:group-test
#task> ntest -- task=group-test:conflict-test
#task> ntest -- task=attr:hidden-test
#task> test -- task=eb:embed-test
#task> test -- task=eb:task-test
#task> test -- task=eb:taskarg-test
#task> test -- task=eb:taskarg2-test
#task> test -- task=eb:func-test
#task> test -- task=eb:funcarg-test
#task> test -- task=eb:funcarg2-test
#task> ntest -- task=eb:funcarg3-test
#task> test -- task=eb:funcarg4-test
#task> test -- task=eb:once1-test
#task> test -- task=eb:once2-test
#task> test -- task=eb:once3-test
~~~

# PowerShell

~~~powershell task::pstest -- [hidden]
#args> task:(group:task)
mdtk -f ./pwsh-sample.task.md $task > $null 2>&1
if (!$?) { echo "failed: $task" }
~~~

~~~powershell task::psntest -- [hidden]
#args> task:(group:task)
mdtk -f ./pwsh-sample.task.md $task > $null 2>&1
if ($?) { echo "failed: $task" }
~~~

~~~powershell task:autotest:pwsh-test-run
#task> pstest -- task=ps-hello
#task> pstest -- task=pseb:embed-test
#task> pstest -- task=pseb:task-test
#task> pstest -- task=pseb:taskarg-test
#task> pstest -- task=pseb:taskarg2-test
#task> pstest -- task=pseb:func-test
#task> pstest -- task=pseb:funcarg-test
#task> pstest -- task=pseb:funcarg2-test
#task> psntest -- task=pseb:funcarg3-test
#task> pstest -- task=pseb:funcarg4-test
#task> pstest -- task=pseb:once1-test
#task> pstest -- task=pseb:once2-test
#task> pstest -- task=pseb:once3-test
~~~
