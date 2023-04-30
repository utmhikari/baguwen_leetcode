package others;

import java.util.*;
import java.util.stream.Collectors;

/**
 * Simple 24 point solver
 * Created by 36191 on 2019/8/23
 */
public class Point24 extends Object {

    private static final double TARGET = 24d;
    private static final int LEN = 4;

    private static class Order {
        List<Double> order;

        Order(List<Double> order) {
            this.order = new ArrayList<>(order);
        }

        @Override
        public int hashCode() {
            int hash = 17;
            for (double d : order) {
                hash = 31 * hash + new Double(d).hashCode();
            }
            return hash;
        }

        @Override
        public boolean equals(Object obj) {
            return obj instanceof Order && (this == obj || obj.hashCode() == this.hashCode());
        }
    }

    private static class Solution {
        double first;
        double second;
        double ret;
        char operator;

        Solution(double first, double second, double ret, char operator) {
            this.first = first;
            this.second = second;
            this.ret = ret;
            this.operator = operator;
        }

        @Override
        public int hashCode() {
            int hash = 17;
            hash = 31 * hash + operator;
            hash = 31 * hash + new Double(ret).hashCode();
            switch (operator) {
                case '-':
                case '/':
                    hash = 31 * hash + new Double(first).hashCode();
                    hash = 31 * hash + new Double(second).hashCode();
                    break;
                case '+':
                case '*':
                    hash = 31 * hash + new Double(first).hashCode() + new Double(second).hashCode();
                    hash = 31 * hash + new Double(first).hashCode() * new Double(second).hashCode();
                    break;
                default:
                    break;
            }
            return hash;
        }

        @Override
        public boolean equals(Object obj) {
            return obj instanceof Solution && (this == obj || obj.hashCode() == this.hashCode());
        }

        @Override
        public String toString() {
            return String.format("%s %c %s = %s", dStr(first), operator, dStr(second), dStr(ret));
        }

        private String dStr(double d) {
            return d == Math.floor(d) ? String.format("%.0f", d) : String.format("%.4f", d);
        }
    }

    private static class Solutions {
        List<Solution> solutions;

        @Override
        public int hashCode() {
            int hash = 17;
            for (Solution r : solutions) {
                hash = 31 * hash + r.hashCode();
            }
            return hash;
        }

        @Override
        public boolean equals(Object obj) {
            return obj instanceof Solutions && (this == obj || obj.hashCode() == this.hashCode()) &&
                    ((Solutions) obj).solutions.size() == this.solutions.size();
        }

        Solutions(List<Solution> solutions) {
            this.solutions = new ArrayList<>(solutions);
        }
    }

    private static void dfsOrders(List<List<Double>> orders, List<Double> order, boolean[] visited,
                                  double[] input) {
        if (order.size() == LEN) {
            orders.add(new ArrayList<>(order));
            return;
        }
        for (int i = 0; i < LEN; ++i) {
            if (!visited[i]) {
                visited[i] = true;
                order.add(input[i]);
                dfsOrders(orders, order, visited, input);
                order.remove(order.size() - 1);
                visited[i] = false;
            }
        }
    }

    private static void dfsCalculate(List<Double> order, List<List<Solution>> solutions, List<Solution> solution,
                                     boolean[] visited, double current) {
        if (solution.size() == LEN - 1) {
            if (Double.compare(current, TARGET) == 0) {
                solutions.add(new ArrayList<>(solution));
            }
            return;
        }
        for (int i = 0; i < LEN; ++i) {
            if (!visited[i]) {
                visited[i] = true;
                // +
                double plus = current + order.get(i);
                solution.add(new Solution(current, order.get(i), plus, '+'));
                dfsCalculate(order, solutions, solution, visited, plus);
                solution.remove(solution.size() - 1);

                // -
                double minus = current - order.get(i);
                solution.add(new Solution(current, order.get(i), minus, '-'));
                dfsCalculate(order, solutions, solution, visited, minus);
                solution.remove(solution.size() - 1);

                // - reverse
                double minusRev = order.get(i) - current;
                solution.add(new Solution(order.get(i), current, minusRev, '-'));
                dfsCalculate(order, solutions, solution, visited, minusRev);
                solution.remove(solution.size() - 1);

                // *
                double multiple = current * order.get(i);
                solution.add(new Solution(current, order.get(i), multiple, '*'));
                dfsCalculate(order, solutions, solution, visited, multiple);
                solution.remove(solution.size() - 1);

                // /
                double divide = current / order.get(i);
                solution.add(new Solution(current, order.get(i), divide, '/'));
                dfsCalculate(order, solutions, solution, visited, divide);
                solution.remove(solution.size() - 1);

                // / reverse
                double divideRev = order.get(i) / current;
                solution.add(new Solution(order.get(i), current, divideRev, '/'));
                dfsCalculate(order, solutions, solution, visited, divideRev);
                solution.remove(solution.size() - 1);

                visited[i] = false;
            }
        }
    }

    private static void calculate(Set<Order> orderSet, List<List<Solution>> solutions,
                                  List<Solution> solution, boolean[] visited) {
        for (Order o : orderSet) {
            for (int i = 0; i < LEN; ++i) {
                visited[i] = true;
                double first = o.order.get(i);
                dfsCalculate(o.order, solutions, solution, visited, first);
                visited[i] = false;
            }
        }
    }

    public static void point24(double[] input) {
        assert input.length == LEN;
        // traverse orders
        List<List<Double>> orders = new ArrayList<>();
        List<Double> order = new ArrayList<>();
        boolean[] visited = new boolean[LEN];
        dfsOrders(orders, order, visited, input);
        Set<Order> orderSet = orders.stream().map(Order::new).collect(Collectors.toSet());
        List<List<Solution>> solutions = new ArrayList<>();
        List<Solution> solution = new ArrayList<>();
        calculate(orderSet, solutions, solution, visited);
        System.out.printf("The solution of %.0f, %.0f, %.0f, %.0f is:\n",
                input[0], input[1], input[2], input[3]);
        if (solutions.size() == 0) {
            System.out.println("\tNo solution!");
        } else {
            Set<Solutions> solutionsSet =  solutions.stream()
                    .map(Solutions::new)
                    .collect(Collectors.toSet());
            for (Solutions s : solutionsSet) {
                System.out.println(String.join("; ", s.solutions.toString()));
            }
            System.out.printf("Overall %s solutions!\n", solutionsSet.size());
        }
    }

    public static void main(String[] args) {
        double[][] inputs = new double[][] {
                new double[]{2, 3, 4, 5},
                new double[]{5, 5, 5, 1},
                new double[]{4, 5, 6, 7},
                new double[]{3, 3, 7, 7},
                new double[]{3, 4, 9, 10}
        };
        for (double[] input : inputs) {
            point24(input);
            System.out.println("\n----------------------------------------------\n");
        }
    }
}
