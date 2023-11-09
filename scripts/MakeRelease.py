'''
  Copyright Tom5521(c) - All Rights Reserved.
 
  This project is licensed under the MIT License.
'''


import Build
import MakeBundle
import MakeWinZip

def MakeRelease():
    print("Making bundle...")
    MakeBundle.MakeBundle()
    print("Compiling for linux...")
    Build.BuildForLinux()
    print("Compiling for windows...")
    Build.BuildForWindows()
    print("Making windows zip")
    MakeWinZip.MakeWinZip()

if __name__ == "__main__":
    MakeRelease()

