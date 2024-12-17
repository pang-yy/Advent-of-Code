import java.util.Scanner;

class Day1 {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        String input = scanner.nextLine();
        scanner.close();

        System.out.println("Part One: " + countFloor(input));
        System.out.println("Part Two: " + enterBasementAt(input, 0, 0));
    }

    static int countFloor(String input) {
        return input.chars()
            .parallel()
            .map(c -> c == 40 ? 1 : -1)
            .reduce(0, (x, y) -> x + y);
    }
    
    static int enterBasementAt(String in, int at, int now) {
        if (now == -1) {
            return at;
        }
        int i = in.substring(0, 1).equals("(") ? 1 : -1;
        return enterBasementAt(in.substring(1), at + 1, now + i);
    }
}
