import java.util.Scanner;
import java.util.List;
import java.util.stream.Stream;

class Box {
    private final int length;
    private final int width;
    private final int height;

    Box(int length, int width, int height) {
        this.length = length;
        this.width = width;
        this.height = height;
    }

    int smallestArea() {
        int lw = this.length * this.width;
        int wh = this.height * this.width;
        int lh = this.height * this.length;
        return lw < wh ? (lw < lh ? lw : lh) : (wh < lh ? wh : lh);
    }

    int smallestPerimeter() {
        int lw = (2 * this.length) + (2 * this.width);
        int wh = (2 * this.height) + (2 * this.width);
        int lh = (2 * this.length) + (2 * this.height);
        return lw < wh ? (lw < lh ? lw : lh) : (wh < lh ? wh : lh);
    }

    int surfaceArea() {
        return (2 * this.length * this.width) + 
            (2 * this.width * this.height) + 
            (2 * this.height * this.length);
    }

    int volume() {
        return this.length * this.width * this.height;
    }
}

class Day2 {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        List<Box> presents = sc.useDelimiter("\n")
            .tokens()
            .map(line -> {
                List<String> token = Stream.of(line.trim().split("x")).toList();
                return new Box(Integer.parseInt(token.get(0)), 
                    Integer.parseInt(token.get(1)), Integer.parseInt(token.get(2)));
            })
            .toList();
        sc.close();

        System.out.println("Part One: " + wrapperNeeded(presents));
        System.out.println("Part Two: " + ribbonNeeded(presents));
    }

    static int wrapperNeeded(List<Box> presents) {
        return presents.stream()
            .map(box -> box.surfaceArea() + box.smallestArea())
            .reduce(0, (x, y) -> x + y);
    }

    static int ribbonNeeded(List<Box> presents) {
        return presents.stream()
            .map(box -> box.volume() + box.smallestPerimeter())
            .reduce(0, (x, y) -> x + y);
    }
}
