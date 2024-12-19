import java.util.Scanner;
import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.CompletableFuture;

class Day10 {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        String input = scanner.nextLine().trim();
        scanner.close();

        int iteration = 50;
        int res = lookAndSayAsync(input, iteration).length();
        System.out.printf("Length after %d iterations: %d\n", iteration, res);
    }

    static String lookAndSay(String initial, int iteration) {
        int t = 0;
        String temp = "";
        for (int i = 0; i < iteration; i += 1) {
            t = 0;
            char c = initial.charAt(0);
            for (char chr : initial.toCharArray()) {
                if (chr == c) {
                    t += 1;
                } else {
                    temp += (Integer.toString(t) + Character.toString(c));
                    t = 1;
                    c = chr;
                }
            }
            temp += (Integer.toString(t) + Character.toString(c));
            initial = temp;
            temp = "";
        }
        return initial;
    }

    static String lookAndSayAsync(String initial, int iteration) {
        final int threshold = 1000;
        for (int i = 0; i < iteration; i += 1) {
            if (initial.length() <= threshold) {
                initial = lookAndSay(initial, 1);
            } else {
                List<CompletableFuture<String>> futures = new ArrayList<CompletableFuture<String>>();
                int startPtr = 0;
                final String temp = initial;
                while (startPtr < initial.length()) {
                    final int start = startPtr;
                    final int end = (startPtr + threshold >= initial.length()) ? 
                        initial.length() : splitAt(initial, startPtr + threshold);
                    futures.add(CompletableFuture.<String>supplyAsync(() -> 
                        lookAndSay(temp.substring(start, end), 1)));
                    startPtr = end;
                }
                initial = futures.stream()
                    .map(cf -> cf.join())
                    .reduce("", (left, right) -> left + right);
            }
        }
        return initial;
    }

    private static int splitAt(String str, int beginAt) {
        int sA = beginAt;
        while (sA < str.length() - 1) {
            if (str.charAt(sA) != str.charAt(sA + 1)) {
                return sA + 1;
            }
            sA += 1;
        }
        return beginAt;
    }
}
