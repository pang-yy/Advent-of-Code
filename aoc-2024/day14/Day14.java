import java.util.Scanner;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.stream.IntStream;
import java.util.List;
import java.util.Objects;

class Pair {
    private final int x;
    private final int y;

    Pair(int x, int y) {
        this.x = x;
        this.y = y;
    }

    int x() {
        return this.x;
    }

    int y() {
        return this.y;
    }

    @Override
    public boolean equals(Object obj) {
        if (this == obj) {
            return true;
        }
        if (obj instanceof Pair other) {
            return other.x == this.x && other.y == this.y;
        }
        return false;
    }

    @Override
    public int hashCode() {
        return Objects.hash(this.x, this.y);
    }
}

class Robot {
    private final Pair pos;
    private final Pair vel;

    public Robot(Pair pos, Pair vel) {
        this.pos = pos;
        this.vel = vel;
    }

    public Pair position() {
        return this.pos;
    }

    public Pair velocity() {
        return this.vel;
    }
}

class Simulator {
    private final int width;
    private final int height;
    private final int count;
    private List<Robot> robots;

    public Simulator(int width, int height) {
        this.width = width;
        this.height = height;
        this.count = 0;
        this.robots = List.of();
    }

    public Simulator(int width, int height, int count, List<Robot> robots) {
        this.width = width;
        this.height = height;
        this.count = count;
        this.robots = robots;
    }

    private int calculatePos(int p, int v, int s, int m) {
        int r = (p + (v * s)) % m;
        if (r < 0) {
            r = m + r;
        }
        return r;
    }

    public List<Pair> calculateFor(int duration) {
        return this.robots.stream().parallel()
            .map(r -> new Pair(calculatePos(r.position().x(), r.velocity().x(), duration, this.width),
                                calculatePos(r.position().y(), r.velocity().y(), duration, this.height)))
            .toList();
    }

    public void displayFloor() {
        ConcurrentHashMap<Pair, AtomicInteger> map = new ConcurrentHashMap<>(this.width * this.height);

        this.robots.stream().parallel()
            .forEach(r -> {
                map.computeIfAbsent(r.position(), f -> new AtomicInteger(0)).incrementAndGet();
            });

        for (int y = 0; y < this.height; y += 1) {
            for (int x = 0; x < this.width; x += 1) {
                AtomicInteger c = map.getOrDefault(new Pair(x, y), new AtomicInteger(0));
                if (c.get() == 0) {
                    System.out.print(".");
                } else {
                    System.out.print(c);
                }
            }
            System.out.println("");
        }
        System.out.printf("Count: %d\n", this.count);
    }

    public Simulator nextState() {
        return new Simulator(this.width, this.height, this.count + 1, this.robots.stream().parallel()
            .map(r -> {
                int nx = this.calculatePos(r.position().x(), r.velocity().x(), 1, this.width);
                int ny = this.calculatePos(r.position().y(), r.velocity().y(), 1, this.height);
                return new Robot(new Pair(nx, ny), r.velocity());
            })
            .toList()
        );
    }
}

class Day14 {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        List<Robot> initRobots = sc.useDelimiter("\n")
            .tokens()
            .parallel()
            .map(line -> parseInput(line))
            .toList();
        sc.close();
        int width = 101;
        int height = 103;
        int duration = 10000;
        Simulator ns = new Simulator(width, height, 0, initRobots);

        System.out.printf("Part One: %d \n", calculatePartOne(ns.calculateFor(duration), width, height));

        runSimulator(ns, duration, width, height);
    }

    private static Robot parseInput(String input) { // p=x,y v=x,y
        String[] ss = input.trim().split(" ");
        String[] ps = ss[0].split("=")[1].split(",");
        String[] vs = ss[1].split("=")[1].split(",");
        return new Robot(
            new Pair(Integer.parseInt(ps[0]), Integer.parseInt(ps[1])),
            new Pair(Integer.parseInt(vs[0]), Integer.parseInt(vs[1]))
        );
    }

    private static int calculatePartOne(List<Pair> finalPos, int width, int height) {
        int q1 = 0; // x = 0 ~ (width / 2 - 1), y = 0 ~ (height / 2 - 1)
        int q2 = 0; // x = (width / 2 + 1) ~ (width - 1), y = 0 ~ (height / 2 - 1)
        int q3 = 0; // x = 0 ~ (width / 2 - 1), y = (height / 2 + 1) ~ (height - 1)
        int q4 = 0; // x = (width / 2 + 1) ~ (width - 1), y = (height / 2 + 1) ~ (height - 1)

        List<Integer> quadrants = finalPos.stream()
            .map(p -> belongsTo(p, width, height))
            .toList();

        q1 = (int)quadrants.stream().parallel().filter(q -> q == 1).count();
        q2 = (int)quadrants.stream().parallel().filter(q -> q == 2).count();
        q3 = (int)quadrants.stream().parallel().filter(q -> q == 3).count();
        q4 = (int)quadrants.stream().parallel().filter(q -> q == 4).count();
        return q1 * q2 * q3 * q4;
    }

    private static int belongsTo(Pair pos, int width, int height) {
        if ((pos.x() >= 0 && pos.x() <= (width / 2 - 1)) &&
            (pos.y() >= 0 && pos.y() <= (height / 2 - 1))) {
            return 1;
        } else if ((pos.x() >= (width / 2 + 1) && (pos.x() <= width - 1)) &&
                    (pos.y() >= 0 && pos.y() <= (height / 2 - 1))) {
            return 2;
        } else if ((pos.x() >= 0 && pos.x() <= (width / 2 - 1)) &&
                    (pos.y() >= (height / 2 + 1) && pos.y() <= height - 1)) {
            return 3;
        } else if ((pos.x() >= (width / 2 + 1) && (pos.x() <= width - 1)) &&
                    (pos.y() >= (height / 2 + 1) && (pos.y() <= height - 1))) {
            return 4;
        }
        return -1;
    }

    private static void runSimulator(Simulator ini, int maxRound, int width, int height) {
        ini.displayFloor();
        System.out.println("=".repeat(50));
        for (int i = 1; i <= maxRound; i += 1) {
            ini = ini.nextState();
            if ((i - height) % width == 0) {
                System.out.println(i);
                ini.displayFloor();
                System.out.println("=".repeat(50));
            }
        }
    }
}
