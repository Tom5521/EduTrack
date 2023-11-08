'''
  Copyright Tom5521(c) - All Rights Reserved.
 
  This project is licensed under the MIT License.
'''

import os
import platform

def BuildForWindows():
    if not os.path.exists("builds"):
        os.mkdir("builds")

    # Check the platform (Windows or not).
    if platform.system() != "windows":
        os.environ["GGO_ENABLED"] = "1"
        os.environ["CC"] = "/usr/bin/x86_64-w64-mingw32-gcc"
        os.environ["CXX"] = "/usr/bin/x86_64-w64-mingw32-c++"
        os.environ["GOOS"] = "windows"

    # Package the application for Windows using fyne package.
    os.system("fyne package --os windows --exe builds/EduTrack.exe")

if __name__ == "__main__":
   BuildForWindows() 
