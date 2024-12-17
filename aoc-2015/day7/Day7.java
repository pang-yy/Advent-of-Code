import java.util.Scanner;
import java.util.function.BinaryOperator;
import java.util.Hashtable;
import java.util.List;
import java.util.stream.Stream;
import java.util.Optional;

class Instruction {
    private final String input1;
    private final Optional<String> input2;
    private final BinaryOperator<Integer> operation;
    private final String output;
    private final boolean canBeInterpret;

    Instruction(List<String> rawInput) {
        String[] rI = rawInput.get(0).split(" ");
        this.input1 = rI.length == 2 ? rI[1] : rI[0];
        this.input2 = rI.length == 3 ? Optional.of(rI[2]) : Optional.empty();
        this.operation = mapOp(rI.length == 1 ? "IDEN" : (rI.length == 2 ? rI[0] : rI[1]));
        this.output = rawInput.get(1);
        this.canBeInterpret = this.input1.matches("\\d+") && 
            this.input2.map(s -> s.matches("\\d+")).orElse(true);
    }

    Instruction(String input1, Optional<String> input2, BinaryOperator<Integer> operation, String output) {
        this.input1 = input1;
        this.input2 = input2;
        this.operation = operation;
        this.output = output;
        this.canBeInterpret = this.input1.matches("\\d+") && 
            this.input2.map(s -> s.matches("\\d+")).orElse(true);
    }

    private static BinaryOperator<Integer> mapOp(String opStr) {
        return opStr.equals("IDEN") ? (x, y) -> x :
            (opStr.equals("NOT") ? (x, y) -> (~x) & 0xFFFF :
                (opStr.equals("LSHIFT") ? (x, y) -> x << y :
                    (opStr.equals("RSHIFT") ? (x, y) -> x >> y :
                        (opStr.equals("AND") ? (x, y) -> x & y : (x, y) -> x | y))));
    }

    boolean canBeInterpret() {
        return this.canBeInterpret;
    }

    String output() {
        return this.output;
    }

    List<String> needed() {
        return this.canBeInterpret() ? List.of() :
            (this.input1.matches("\\d+") ? List.of(this.input2.map(s -> s).orElse("")) :
                this.input2.map(s -> s.matches("\\d+")).orElse(true) ? List.of(this.input1) :
                    List.of(this.input1, this.input2.map(s -> s).orElse("")));
    }

    Instruction replaceInput(String oldIn, String newIn) {
        return this.input1.equals(oldIn) ? 
            (new Instruction(newIn, this.input2, this.operation, this.output)) : 
            (this.input2.map(s -> s.equals(oldIn)).orElse(false) ? 
                (new Instruction(this.input1, Optional.of(newIn), this.operation, this.output)) : 
                this);
    }
    
    Instruction rewriteInputs(String newIn) {
        return new Instruction(newIn, Optional.empty(), this.operation, this.output);
    }

    int eval() {
        return this.operation.apply(
            Integer.parseInt(this.input1),
            this.input2.map(s -> Integer.parseInt(s)).orElse(0));
    }
}

class Day7 {
    private Hashtable<String, Integer> table;
    private final List<Instruction> instructions;
    private final boolean isPartTwo;

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        List<Instruction> instructions = sc.useDelimiter("\n")
            .tokens()
            .map(line -> {
                List<String> token = Stream.of(line.trim().split(" -> ")).toList();
                return new Instruction(token);
            })
            .toList();

        int a1 = new Day7(instructions, false).solveA();
        int a2 = new Day7(instructions, true).solveA();
        System.out.printf("Part One: %d\n", a1);
        System.out.printf("Part Two: %d\n", a2);

        sc.close();
    }

    Day7(List<Instruction> instructions, boolean isPartTwo) {
        this.instructions = instructions;
        this.table = new Hashtable<String, Integer>();
        this.isPartTwo = isPartTwo;
        this.instructions.stream()
            .filter(ins -> ins.canBeInterpret())
            .forEach(ins -> {
                table.put(ins.output(), ins.eval());
            });
        if (isPartTwo) {table.put("b", 956);}
    }

    int solveA() {
        boolean found = false;
        while (!found) {
            this.instructions.stream()
                .filter(ins -> !this.isPartTwo || !ins.output().equals("b"))
                .forEach(ins -> {
                    List<String> needed = ins.needed();
                    if (needed.size() == 1 && table.containsKey(needed.get(0))) {
                        Integer r = ins
                            .replaceInput(needed.get(0), table.get(needed.get(0)).toString())
                            .eval();
                        table.put(ins.output(), r);
                    } else if (needed.size() == 2 && table.containsKey(needed.get(0)) && 
                            table.containsKey(needed.get(1))) {
                        Integer r = ins
                            .replaceInput(needed.get(0), table.get(needed.get(0)).toString())
                            .replaceInput(needed.get(1), table.get(needed.get(1)).toString())
                            .eval();
                        table.put(ins.output(), r);
                    }
                });
            found = this.table.containsKey("a");
        }
        return table.get("a");
    }
}
