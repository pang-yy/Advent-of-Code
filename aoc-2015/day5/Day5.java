import java.util.Scanner;
import java.util.List;
import java.util.stream.Stream;

class Day5 {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        List<String> inputs = sc.useDelimiter("\n")
            .tokens()
            .map(line -> line.trim())
            .toList();
        sc.close();
       
        int partOne = inputs.stream()
            .map(input -> hasVowels(input) && hasRepeat1(input) && !hasContain1(input))
            .filter(b -> b)
            .map(b -> 1)
            .reduce(0, (x, y) -> x + 1);
        
        int partTwo = inputs.stream()
            .map(input -> hasRepeat2(input) && hasContain2(input))
            .filter(b -> b)
            .map(b -> 1)
            .reduce(0, (x, y) -> x + 1);

        System.out.println("Part One: " + partOne);
        System.out.println("Part Two: " + partTwo);
    }

    static boolean hasVowels(String input) {
        return input.chars()
            .mapToObj(c -> (char)c)
            .filter(c -> c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u')
            .map(c -> 1)
            .reduce(0, (x, y) -> x + 1) >= 3;
    }

    static boolean hasRepeat1(String input) {
        return Stream.iterate(0, i -> i < input.length() - 1, i -> i + 1)
            .map(i -> input.substring(i, i + 2))
            .filter(str -> str.substring(0, 1).equals(str.substring(1, 2)))
            .findFirst()
            .map(s -> true)
            .orElse(false);
    }
    
    static boolean hasRepeat2(String input) {
        return Stream.iterate(0, i -> i < input.length() - 2, i -> i + 1)
            .map(i -> input.substring(i, i + 3))
            .filter(str -> str.substring(0, 1).equals(str.substring(2, 3)))
            .findFirst()
            .map(s -> true)
            .orElse(false);
    }
    
    static boolean hasContain1(String input) {
        return Stream.iterate(0, i -> i < input.length() - 1, i -> i + 1)
            .map(i -> input.substring(i, i + 2))
            .filter(str -> 
                str.equals("ab") ||
                str.equals("cd") ||
                str.equals("pq") ||
                str.equals("xy"))
            .findFirst()
            .map(s -> true)
            .orElse(false);
    }
    
    static boolean hasContain2(String input) {
        return Stream.iterate(0, i -> i < input.length() - 1, i -> i + 1)
            .filter(i -> input.indexOf(input.substring(i, i + 2), i + 2) != -1)
            .findFirst()
            .map(s -> true)
            .orElse(false);
    }
}
