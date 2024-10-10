from abc import ABC, abstractmethod

class Animal(ABC):
    @abstractmethod
    def som(self):
        pass

def emitir_som_dos_animais(animais):
    for animal in animais:
        animal.som()  


