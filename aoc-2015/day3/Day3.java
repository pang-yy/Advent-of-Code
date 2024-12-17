import java.util.Scanner;
import java.util.function.Predicate;
import java.util.stream.Stream;
import java.util.Hashtable;
import java.util.Objects;

class Point {
    private final int x;
    private final int y;
    
    Point(int x, int y) {
        this.x = x;
        this.y = y;
    }

    Point addWith(Point dc) {
        return new Point(this.x + dc.x, this.y + dc.y);
    }

    @Override
    public boolean equals(Object obj) {
        if (this == obj) {
            return true;
        }
        if (obj instanceof Point other) {
            return (this.x == other.x) && (this.y == other.y);
        }
        return false;
    }

    @Override
    public int hashCode() {
        return Objects.hash(this.x, this.y);
    }
}

class Day3 {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        String input = scanner.nextLine();
        scanner.close();
        
        System.out.println("Part One: " + partOne(input));
        System.out.println("Part Two: " + partTwo(input));
    }

    static int partOne(String routes) {
        Hashtable<Point, Boolean> table = new Hashtable<Point, Boolean>();
        table.put(new Point(0, 0), true);
        table = countDistinct(routes, table, x -> true);
        return table.size();
    }
    
    static int partTwo(String routes) {
        Hashtable<Point, Boolean> table = new Hashtable<Point, Boolean>();
        table.put(new Point(0, 0), true);
        table = countDistinct(routes, table, x -> x % 2 == 0);
        table = countDistinct(routes, table, x -> x % 2 != 0);
        return table.size();
    }

    static Hashtable<Point, Boolean> countDistinct(String routes, 
        Hashtable<Point, Boolean> table, Predicate<Integer> pred) {
        Stream.<Integer>iterate(0, i -> i < routes.length(), i -> i += 1)
            .filter(pred)
            .map(i -> routes.chars().boxed().toList().get(i))
            .map(c -> c.equals((int)'^') ? new Point(0, -1) :
                    (c.equals((int)'>') ? new Point(1, 0) :
                    (c.equals((int)'<') ? new Point(-1, 0) : new Point(0, 1))))
            .reduce(new Point(0, 0), (x, y) -> {
                Point newC = x.addWith(y);
                table.put(newC, true);
                return newC;
            });
        return table;
    }
}
