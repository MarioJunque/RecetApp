public class Hello {

    public static void main(String[] args) {
        javax.swing.SwingUtilities.invokeLater(new Runnable() {
            public void run() {
                // your logic here
                System.out.println("Hello, world!");
                System.out.println(args[0]);
            }
        });
    }
}
