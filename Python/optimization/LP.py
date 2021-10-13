from ortools.linear_solver import pywraplp

def main():
    # Select a solver according to the type of your problem.
    solver = pywraplp.Solver(name='SolveSimpleSystem', problem_type=pywraplp.Solver.GLOP_LINEAR_PROGRAMMING)
    # Create the variables
    x1 = solver.NumVar(0, 9, name='x1')
    x2 = solver.NumVar(0, 9, name='x2')
    x3 = solver.NumVar(0, 9, name='x3')
    # create the constraints
    constraint1 = solver.Constraint(-solver.infinity(), 9)
    constraint1.SetCoefficient(x1, 2)
    constraint1.SetCoefficient(x2, 3)
    constraint1.SetCoefficient(x3, 1)

    constraint2 = solver.Constraint(-solver.infinity(), 3)
    constraint2.SetCoefficient(x1, -3)
    constraint2.SetCoefficient(x2, -1)
    constraint2.SetCoefficient(x3, 4)

    # Create the objective function
    objective = solver.Objective()
    objective.SetCoefficient(x1, 4)
    objective.SetCoefficient(x2, -2)
    objective.SetCoefficient(x3, 3)
    objective.SetMaximization()

    # Call the solver.
    solver.Solve()
    print('Number of variables =', solver.NumVariables())
    print('Number of constraints =', solver.NumConstraints())
    # The value of each variable in the solution.
    print('Solution:')
    print('x1 = ', x1.solution_value())
    print('x2 = ', x2.solution_value())
    print('x3 = ', x3.solution_value())
    # The objective value of the solution.
    opt_solution = 4 * x1.solution_value() - 2 * x2.solution_value() + 3 * x3.solution_value()
    print('Optimal objective value =', opt_solution)

if __name__ == '__main__':
    main()
