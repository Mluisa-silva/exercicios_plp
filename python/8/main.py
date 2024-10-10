from calculadora import *
if __name__ == "__main__":
    calc = Calculadora()
    
    resultado1 = calc.somar(5.0, 3.0)
    print("Soma de 5.0 e 3.0:", resultado1)
    
    resultado2 = calc.somar_três(5.0, 3.0, 2.0)
    print("Soma de 5.0, 3.0 e 2.0:", resultado2)