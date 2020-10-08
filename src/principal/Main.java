/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package principal;

import java.util.Scanner;

/**
 *
 * @author enrique
 */
public class Main {

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        // TODO code application logic here
        Scanner scanner = new Scanner(System.in);
        String usuario= "User1";
        String contrasena = "1234";

        System.out.println("\nPor favor, introduzca su nombre de usuario:\n");
        String user_name = scanner.nextLine();
        if(user_name.equals(usuario)){
            System.out.println("\nPor favor, introduzca su contraseña:\n");
            String pass = scanner.nextLine();
                if(pass.equals(contrasena)){
                    System.out.println("\n Acceso concedido\n");
                }
                else{
                    System.out.println("\nContraseña incorrecto\n");
                }
        }
        else{
            System.out.println("\nUsuario incorrecto\n");
        }
    }
}
