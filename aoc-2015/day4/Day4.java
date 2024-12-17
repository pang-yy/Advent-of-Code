import java.security.MessageDigest;
import java.util.concurrent.CompletableFuture;

class Day4 {
    public static void main(String[] args) {
        String key = args[0];
        //System.out.println(findHash(key, 0, 1));
       
        CompletableFuture<Integer> cf1 = CompletableFuture.supplyAsync(() -> findHash(key, 0, 4));
        CompletableFuture<Integer> cf2 = CompletableFuture.supplyAsync(() -> findHash(key, 1, 4));
        CompletableFuture<Integer> cf3 = CompletableFuture.supplyAsync(() -> findHash(key, 2, 4));
        CompletableFuture<Integer> cf4 = CompletableFuture.supplyAsync(() -> findHash(key, 3, 4));
        
        CompletableFuture<Object> answer = CompletableFuture.anyOf(cf1, cf2, cf3, cf4);
        answer.thenAccept(a -> System.out.println("Part Two: " + a)).join();
    }

    static int findHash(String key,int seed, int steps) {
        int answer = seed;
        try {
            MessageDigest md = MessageDigest.getInstance("MD5");
            byte[] bytesOfMessage;
            while (true) {
                bytesOfMessage = (key + answer).getBytes("UTF-8");
                String theMD5digest = toHexString(md.digest(bytesOfMessage));
                if (theMD5digest.substring(0, 6).equals("000000")) {
                    //System.out.println(answer);
                    return answer;
                } else {
                    answer += steps;
                }
            }

        } catch (Exception e) {
            System.out.println("Hold my poke ball!");
        }
        return -1;
    }

    // from StackOverFlow
    static String toHexString(byte[] bytes) {
        StringBuilder hexString = new StringBuilder();

        for (int i = 0; i < bytes.length; i++) {
            String hex = Integer.toHexString(0xFF & bytes[i]);
            if (hex.length() == 1) {
                hexString.append('0');
            }
            hexString.append(hex);
        }

        return hexString.toString();
    }
}
