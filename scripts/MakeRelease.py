'''
  Copyright Tom5521(c) - All Rights Reserved.
 
  This project is licensed under the MIT License.
'''


import Build4Win
import Build4linux
import MakeBundle
import MakeWinZip

def MakeRelease():
    print("Making bundle...")
    MakeBundle.MakeBundle()
    print("Compiling for linux...")
    Build4linux.BuildForLinux()
    print("Compiling for windows...")
    Build4Win.BuildForWindows()
    print("Making windows zip")
    MakeWinZip.MakeWinZip()

if __name__ == "__main__":
    MakeRelease()

