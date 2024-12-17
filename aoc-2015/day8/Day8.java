import java.util.Scanner;
import java.util.List;

class Literal {
    private final String seq;
    private final int prevCodeLength;

    Literal(String seq) {
        this.seq = seq;
        this.prevCodeLength = -1;
    }

    private Literal(String seq, int prev) {
        this.seq = seq;
        this.prevCodeLength = prev;
    }

    Literal encode() {
        String newStr = "\"";
        newStr += e(this.seq);
        newStr += "\"";
        return new Literal(newStr, this.lengthInCode());
    }

    private String e(String s) {
        if (s.length() == 0) {
            return "";
        }
        if (s.substring(0, 1).equals("\"")) {
            return "\\" + "\"" + e(s.substring(1));
        }
        if (s.substring(0, 1).equals("\\")) {
            return "\\\\" + e(s.substring(1));
        }
        return s.substring(0, 1) + e(s.substring(1));
    }

    int lengthInCode() {
        return this.seq.length();
    }

    int prevLengthInCode() {
        return this.prevCodeLength;
    }

    int lengthInMem() {
        return this.r(this.seq) - 2;
    }

    private int r(String s) {
        if (s.length() == 0) {
            return 0;
        }
        if (s.substring(0, 1).equals("\\")) {
            if (s.substring(1,2).equals("\"") || s.substring(1, 2).equals("\\")) {
                s = s.substring(2);
            } else {
                s = s.substring(4);
            }
        } else {
            s = s.substring(1);
        }
        return 1 + r(s);
    }
}

class Day8 {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        List<Literal> inputs = sc.useDelimiter("\n")
            .tokens()
            .map(line -> new Literal(line.trim()))
            .toList();
        sc.close();

        int partOne = inputs.stream()
            .map(li -> li.lengthInCode() - li.lengthInMem())
            .reduce(0, (x, y) -> x + y);
        int partTwo = inputs.stream()
            .map(li -> li.encode())
            .map(li -> li.lengthInCode() - li.prevLengthInCode())
            .reduce(0, (x, y) -> x + y);

        System.out.printf("Part One: %d\n", partOne);
        System.out.printf("Part Two: %d\n", partTwo);

    }
}
