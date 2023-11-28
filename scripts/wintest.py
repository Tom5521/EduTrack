
import os
import Build

#To run this properly you must have wine installed
def Exec():
    print("Setting enviromevent values...")
    Build.SetEnvForWin()
    print("Compiling and running...")
    os.system("go run .")

if __name__ == "__main__":
   Exec() 
