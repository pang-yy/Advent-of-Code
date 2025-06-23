import java.util.Set;
import java.util.HashSet;
import java.util.Scanner;
import java.util.stream.Stream;

class Day11 {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        String oldPassword = scanner.nextLine().trim();
        scanner.close();

        String newPassword = nextPassword(oldPassword);
        System.out.println("Part One: " + newPassword);
    }

    static String nextPassword(String oldPassword) {
        oldPassword = incrementString(oldPassword);
        while (!meetFirt(oldPassword) || !meetSecond(oldPassword) || !meetThird(oldPassword)) {
            oldPassword = incrementString(oldPassword);
        }
        return oldPassword;
    }

    static String incrementString(String str) {
        int lastPtr = str.length() - 1;
        boolean isWrap = true;
        while (isWrap) {
            isWrap = false;
            if (lastPtr == -1) {
                str = "a" + str;
            } else {
                char currChar = str.charAt(lastPtr);
                if (currChar == 'z') {
                    currChar = 'a';
                    isWrap = true;
                } else {
                    currChar += 1;
                }
                str = str.substring(0, lastPtr) + Character.toString(currChar) + str.substring(lastPtr + 1);
                lastPtr -= 1;
                
            }
        }
        return str;
    }

    static boolean meetFirt(String str) {
        if (str.length() < 3) {
            return false;
        }
        return Stream.iterate(0, i -> i < str.length() - 2, i -> i + 1)
            .map(i -> str.substring(i, i + 3))
            .filter(s -> (s.charAt(0) - s.charAt(1) == -1) && (s.charAt(1) - s.charAt(2) == -1))
            .findFirst()
            .map(s -> true)
            .orElse(false);
    }

    static boolean meetSecond(String str) {
        return str.chars()
            .boxed()
            .filter(x -> x == 'i' || x == 'o' ||x == 'l')
            .findFirst()
            .map(x -> false)
            .orElse(true);
    }

    static boolean meetThird(String str) {
        Set<String> ensureUnique = new HashSet<>();
        return Stream.iterate(0, i -> i < str.length() - 1, i -> i + 1)
            .map(i -> str.substring(i, i + 2))
            .filter(s -> s.charAt(0) == s.charAt(1))
            .filter(s -> ensureUnique.add(s))
            .map(s -> 1)
            .reduce(0, (x, y) -> x + 1) >= 2;
    }
}
