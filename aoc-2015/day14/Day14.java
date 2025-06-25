import java.util.Scanner;
import java.util.stream.IntStream;
import java.util.ArrayList;
import java.util.List;

class Reindeer {
    private final int speed;
    private final int sTime;
    private final int rTime;

    public Reindeer(int speed, int sTime, int rTime) {
        this.speed = speed;
        this.sTime = sTime;
        this.rTime = rTime;
    }

    // TODO: Think of better solution
    public int distanceAfter(int time) {
        int dist = 0;
        int timeNow = 0;
        int remainTime = this.sTime;
        boolean isRunning = true;
        while (timeNow <= time) {
            if (isRunning & remainTime > 0) {
                dist += this.speed;
            } else if (remainTime == 0) {
                if (!isRunning) {
                    remainTime = this.sTime;
                    dist += this.speed;
                } else {
                    remainTime = this.rTime;
                }
                isRunning = !isRunning;
            }
            remainTime -= 1;
            timeNow += 1;
        }
        return dist;
    }
}

class Day14 {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        List<Reindeer> reindeers = sc.useDelimiter("\n")
            .tokens()
            .parallel()
            .map(line -> parseInput(line))
            .toList();
        sc.close();

        int time = 1000;
        int maxDist = reindeers.stream().parallel()
            .map(r -> r.distanceAfter(time * 100000))
            .reduce(-1, (x, y) -> x > y ? x : y);
        System.out.printf("Part One: %d \n", maxDist);
        System.out.printf("Part Two: %d \n", partTwo(reindeers, time));
    }

    private static Reindeer parseInput(String line) {
        String[] temp = line.split(" ");
        int speed = Integer.parseInt(temp[3]);
        int sTime = Integer.parseInt(temp[6]);
        int rTime = Integer.parseInt(temp[13]);
        return new Reindeer(speed, sTime, rTime);
    }

    private static int partTwo(List<Reindeer> reindeers, int time) {
        List<Integer> points = new ArrayList<>(
            IntStream.range(0, reindeers.size()).map(i -> 1).boxed().toList()
        );

        // TODO: Replace for loop with intstream
        // TODO: Think of better solution
        for (int i = 0; i < time; i += 1) {
            final int j = i;
            List<Integer> distNow = reindeers.stream()
                .map(r -> r.distanceAfter(j))
                .toList();

            int maxNow = -1;
            for (int d = 0; d < distNow.size(); d += 1) {
                if (distNow.get(d) > maxNow) {
                    maxNow = distNow.get(d);
                }
            }

            for (int d = 0; d < distNow.size(); d += 1) {
                if (distNow.get(d).equals(maxNow)) {
                    points.set(d, points.get(d) + 1);
                }
            }
        }
        return points.stream()
            .reduce(0, (x, y) -> x > y ? x : y);
    }
}
