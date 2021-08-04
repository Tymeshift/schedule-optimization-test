# Tymeshift - Software engineer schedule optimization test
Welcome to Tymeshift's schedule optimization code test 🧪

## Problem description:

The problem that we are trying to solve is to schedule working days for an agent according to some constraints.  

You have the following parameters:  
1. Number of days to schedule for - it should be a multiple of 7.
2. Maximum and minimum number of consecutive working days.
   So, longest working days series without a day off should not exceed maximum parameter and shortest series should not be less than a minimum parameter.
3. Maximum and minimum number of consecutive days off.
   Basically the same, but for days off.
4. Pre-defined days off (i.e. Tuesday and Friday are days off for this agent).

You can see an example of those parameters defined in `helpers` folder.

End result should be something like `[[5,3],[3,3],[5,2],[5,2]]`, meaning - agent works 5 days and has 3 days off, works 3 days and has 3 days off, etc.
(Note that we do not use this output in the test case, but you can convert the solution to this format using BlocksToSolution helper function, it can help you visualise and debug the problem.)

1. First and foremost we need to make sure that this schedule does not violate the constraints (i.e. - evaluation penalty should be zero).  
2. Next step is optimisation - we want to make sure that agent works as much as possible(in real life it is not the case and each day has a different value, and we would try to schedule the agent to work in a more valuable days, but for the simplicity sake it is not the case in this test).
3. To achieve that we assign a bonus value for each working day.  
4. Lastly - we want to make sure that the algorithm can find the solution as fast as possible.

## Task

There are two ways you can accomplish this task:

1. Use the code from this repository
   1. Implement FindNeighborhood method(it's currently a stub)
   2. Make sure that tabu_test are passing (the only thing they check is that penalty is zero)
   3. Bonus: increase the total score (without changing the limitNotImproved and number of iterations)
   4. Bonus2: improve the algorithm performance(we have a benchmark ready in tabu_test)  
   You can improve anything you want, i.e. - base tabu search part, neighboor generation, how we choose the best solution, initial solution, even the evaluation process.   
   The only condition is that it should still adhere to the problem parameters. 
   
2. If you are not comfortable with the golang or just want to use a completely different approach - feel free to do so. It is okay to just take the constraints from the problem definition and solve the problem in any language you want, using any method you want as long as it respects the original rules.

(Mind you - in a simplified test case it is possible to just brute force a problem, but the actual problem has more constraints and more things to optimise for, so it would be better to have a more sophisticated solution.)

## Rules

* You can upload your solution to GitHub or send us a ZIP file with the solution
* Take the time you need to feel comfortable with your solution and do it at your own pace

## What we're looking for

* A clean solution that is easy to follow, respects all constraints and provides an optimal result
* How you approach and dissect the problem
* Bonus examples: test improvements, suggest optimizations and drawbacks of your implementation and how to improve it.

---

Thank you for applying and we can't wait to see your submission 🎉
