public class Configuracao {
    private static Configuracao instance;

    private Configuracao() {
    }

    public static Configuracao getInstance() {
        if (instance == null) {
            instance = new Configuracao();
        }
        return instance;
    }

    public void setConfig(String chave, String valor) {
    }

    public String getConfig(String chave) {
        return "exemplo";  // Retorno de exemplo
    }
}


