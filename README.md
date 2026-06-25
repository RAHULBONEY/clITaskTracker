the core logic
1. Users types the task he wants to ,the name ,time
2. Can get the tasks
3. Can update the tasks he want
4. Can delete the tasks he wants
5. Can create the tasks he wants
6. Can  update the task status

task structure

{
   name string
   isCompleted boolean
   time string
   id int
   completedAt string
   notes string
}

Core functioning
1. User add task.
i. User writes task-tracker add  [task name]
ii. the task is added to the json with a key called isCompleted false with the current time
iii.This is stored in our json file with the task name ,task notes etc

APIs logic
1.Get api request return all the json data.
2.Post api ,accept json from the user ,add task and return response