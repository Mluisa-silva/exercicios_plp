public class ConfigMain {
    public static void main(String[] args) {
        Configuracao config1 = Configuracao.getInstance();
        config1.setConfig("tema", "escuro");

        Configuracao config2 = Configuracao.getInstance();
        System.out.println(config2.getConfig("tema"));  // Saída: exemplo
        System.out.println(config1 == config2);          // Saída: true (mesma instância)
    }
}
