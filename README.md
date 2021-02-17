# data-science-test
These are the instructions for the data science code test

## Problem definition:

The problem that we are trying to solve is to schedule working days for an agent according to some constraints.  
You have the following parameters:  
1. Number of days to schedule for - it should be a multiple of 7.
2. Maximum and minimum number of consecutive working days.
   So, longest working days series without a day off should not exceed maximum parameter and shortest series should not be less than a minimum parameter.
3. Maximum and minimum number of consecutive days off.
   Basically the same, but for days off.
4. Pre-defined days off (i.e. Tuesday and Friday are days off for this agent)

You can see an example of those parameters defined in `helpers` folder.
End result should be something like `[[5,3],[3,3],[5,2],[5,2]]`, meaning - agent works 5 days and has 3 days off, works 3 days and has 3 days off, etc.
(Note that we do not use this output in the test case, but you can convert the solution to this format using BlocksToSolution helper function, it can help you visualise and debug the problem)

First and foremost we need to make sure that this schedule does not violate the constraints (i.e. - evaluation penalty should be zero).  
Next step is optimisation - we want to make sure that agent works as much as possible(in real life it is not the case and each day has a different value, and we would try to schedule the agent to work in a more valuable days, but for the simplicity sake it is not the case in this test).
To achieve that we assign a bonus value for each working day.  
Lastly - we want to make sure that the algorithm can find the solution as fast as possible.

## Task

There are two ways you can accomplish this task:

1. Use the code from this repository
   1. Implement FindNeighborhood method(it's currently a stub)
   2. Make sure that tabu_test are passing (the only thing they check is that penalty is zero)
   3. Bonus: increase the total score (without changing the limitNotImproved and number of iterations)
   4. Bonus2: improve the algorithm performance(we have a benchmark ready in tabu_test)
2. If you are not comfortable with the golang or just want to use a completely different approach - feel free to do so.   
   It is okay to just take the constraints from the problem definition and solve the problem in any language you want, using any method you want.   
   (Mind you - in a simplified test case it is possible to just brute force a problem, but the actual problem have more constraints and more things to optimise for, so it would be best to have a more sophisticated solution)
