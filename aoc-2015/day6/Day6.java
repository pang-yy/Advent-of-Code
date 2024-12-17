import java.util.Scanner;
import java.util.List;
import java.util.stream.Stream;

class Pair<T, R> {
    private final T x;
    private final R y;

    Pair(T x, R y) {
        this.x = x;
        this.y = y;
    }

    T x() {
        return this.x;
    }

    R y() {
        return this.y;
    }
}

class Point extends Pair<Integer, Integer> {
    Point(int x, int y) {
        super(x, y);
    }
}

class Day6 {
    private final List<List<Integer>> grid;
    private final List<Pair<Integer, Pair<Point, Point>>> instructions;
    private static final int x_side = 1000;
    private static final int y_side = 1000;

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        List<Pair<Integer, Pair<Point, Point>>> instructions = sc.useDelimiter("\n")
            .tokens()
            .map(line -> line.trim().split(" "))
            .map(strArr -> new Pair<Integer, Pair<Point, Point>>(parseInstruction(strArr[0], strArr[1]),
                new Pair<Point, Point>(
                    new Point(Integer.parseInt(strArr[strArr.length - 3].split(",")[0]),
                            Integer.parseInt(strArr[strArr.length - 3].split(",")[1])),
                    new Point(Integer.parseInt(strArr[strArr.length - 1].split(",")[0]),
                            Integer.parseInt(strArr[strArr.length - 1].split(",")[1]))
                )))
            .toList();
        sc.close();

        System.out.printf("Part One: %d \n", new Day6(instructions).partOne());
        System.out.printf("Part Two: %d \n", new Day6(instructions).partTwo());
    }

    static int parseInstruction(String a, String b) { // off: 0, on: 1, toggle: 2
        return a.equals("toggle") ? 2 : 
            (b.equals("on") ? 1 : 0);
    }
    
    Day6(List<Pair<Integer, Pair<Point, Point>>> instructions) {
        this.instructions = instructions;
        this.grid = Stream.iterate(0, i -> i < Day6.y_side, i -> i += 1)
            .map(y -> Stream.iterate(0, i -> i < Day6.x_side, i -> i += 1)
                .map(x -> -1)
                .toList())
            .toList();
    }
    
    private Day6(List<Pair<Integer, Pair<Point, Point>>> instructions,
                List<List<Integer>> grid) {
        this.instructions = instructions;
        this.grid = grid;
    }

    int partOne() {
        return this.instructions.stream()
            .map(ins -> new Day6(this.instructions, toGrid(ins.x(), ins.y())))
            .reduce(this, (x, y) -> x.compareGrid(y))
            .countOn();
    }
    
    int partTwo() {
        return this.instructions.stream()
            .map(ins -> new Day6(this.instructions, toGrid(ins.x(), ins.y())))
            .reduce(this, (x, y) -> x.compareGrid2(y))
            .countTotal();
    }

    List<List<Integer>> toGrid(int instruction, Pair<Point, Point> points) {
        int startX = points.x().x();
        int startY = points.x().y();
        int endX = points.y().x();
        int endY = points.y().y();
        List<List<Integer>> newGrid = Stream.iterate(0, i -> i < Day6.y_side, i -> i += 1)
            .map(y -> Stream.iterate(0, i -> i < Day6.x_side, i -> i += 1)
                .map(x -> x >= startX && x <= endX && y >= startY && y <= endY ? instruction : -1)
                .toList())
            .toList();
        return newGrid;
    }

    Day6 compareGrid(Day6 other) {
        List<List<Integer>> newGrid = Stream.iterate(0, i -> i < Day6.y_side, i -> i += 1)
            .map(y -> Stream.iterate(0, i -> i < Day6.x_side, i -> i += 1)
                .map(x -> other.grid.get(y).get(x) == 2 ? (this.grid.get(y).get(x) == 1 ? 0 : 1) :
                            (other.grid.get(y).get(x) == -1 ? this.grid.get(y).get(x) : 
                                other.grid.get(y).get(x)))
                .toList())
            .toList();

        return new Day6(this.instructions, newGrid);
    }
    
    Day6 compareGrid2(Day6 other) {
        List<List<Integer>> newGrid = Stream.iterate(0, i -> i < Day6.y_side, i -> i += 1)
            .map(y -> Stream.iterate(0, i -> i < Day6.x_side, i -> i += 1)
                .map(x -> this.grid.get(y).get(x) == -1 ? other.grid.get(y).get(x) :
                            (other.grid.get(y).get(x) == 0 ? this.grid.get(y).get(x) - 1 :
                                ((other.grid.get(y).get(x) == -1) ? this.grid.get(y).get(x) : 
                                    this.grid.get(y).get(x) + other.grid.get(y).get(x))))
                .toList())
            .toList();

        return new Day6(this.instructions, newGrid);
    }

    int countOn() {
        return this.grid.stream()
            .map(y -> y.stream()
                .filter(x -> x == 1)
                .reduce(0, (left, right) -> left + 1))
            .reduce(0, (x, y) -> x + y);
    }

    int countTotal() {
        return this.grid.stream()
            .map(y -> y.stream()
                .filter(x -> x >= 0)
                .reduce(0, (left, right) -> left + right))
            .reduce(0, (x, y) -> x + y);
    }
}
