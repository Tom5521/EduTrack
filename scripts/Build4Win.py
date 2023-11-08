'''
  Copyright Tom5521(c) - All Rights Reserved.
 
  This project is licensed under the MIT License.
'''

import os
import platform
import shutil

def BuildForWindows():
    if not os.path.exists("builds"):
        os.mkdir("builds")

    # Check the platform (Windows or not).
    if platform.system() != "windows":
        # Cross-compile the application for Windows using fyne-cross.
        os.system("sudo fyne-cross windows -arch=amd64")

        # Check if the "builds" directory exists, create it if not.
        if not os.path.exists("builds"):
            os.mkdir("builds")

        # Copy the compiled EduTrack.exe to the "builds" directory.
        shutil.copy("./fyne-cross/bin/windows-amd64/EduTrack.exe", "./builds/EduTrack.exe")
    else:
        # Set the FYNE_THEME environment variable to 'dark' for Windows.

        # Package the application for Windows using fyne package.
        os.system("fyne package --os windows --src . --exe builds/EduTrack.exe")

        # Display a message and wait for user input before closing the command prompt.
        print("Press Enter to close this cmd...")
        input()

if __name__ == "__main__":
   BuildForWindows() 
