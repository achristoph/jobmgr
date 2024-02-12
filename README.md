# Job Manager

Job Manager package provides the mechanism to manage job such as running shell commands

Job Manager package provides the mechanism to manage job such as running shell commands, http requests , etc
In the first iteration, we want to define the main methods of the package.
A job can have 3 main actions: run, retry, cancel
A workflow is a set of commands that run in sequence (parallel command would have different complexity - saved for later improvements)

- Run is running the actual command specified
- A job is associated with a set of commands, which would also have a set of coresponding outputs
- Retry is used when run resulted in failure, hence require another retry. This action becomes useful when previous successful commands are skipped to the point of failed command that can be retried
- Cancel is when the workflow is completely canceled

- In this first iteration, these are the implementations:
  - Job implements an interface called Manager, this allow us to mock the exec command implementation later on
  - We only implement the Run method where it does execute actual shell commands
  - We add a test method so that we can compare it with the expected outputs using shell commands the result in constant values such as echo
  - Two test are defined, one for a successful run with expected output and one for the failure run where the command throws an error
