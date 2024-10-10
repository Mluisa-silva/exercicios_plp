import java.util.ArrayList;
import java.util.List;

public class Som {
    public static void main(String[] args) {
        List<Animal> animais = new ArrayList<>();

        animais.add(new Cachorro());
        animais.add(new Gato());

        emitirSomDosAnimais(animais);
    }

    public static void emitirSomDosAnimais(List<Animal> animais) {
        for (Animal animal : animais) {
            animal.som();
        }
    }
}
