import os
import re

def exist():
    cwd = 'F:\\Work\\dongtaitech\\fileadd\\'
    
    files = filter(lambda f: f.startswith('Btrans20170712') and f.endswith('.u'), os.listdir(cwd))
    print(len(files))

exist()
