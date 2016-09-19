import json

dic = {}

with open('branchtown.csv', encoding='gbk') as file:
    ls = file.readlines()
for line in ls:
    bt = line.strip('\n').split(',')
    if bt[0] not in dic:
        dic[bt[0]] = []
        dic[bt[0]].append(bt[1])
    else:
        dic[bt[0]].append(bt[1])


out = json.dumps(dic, ensure_ascii=False)

with open('branchtown.json', 'w', encoding='gbk') as file:
    file.write(out)

text = ""

with open('src.csv', encoding='gbk') as file:
    src = file.readlines()
for line in src:
    sr = line.strip('\n').split(',')
    text += sr[0]+'支行'+','
    for cun in dic[sr[0]+'支行']:
        if cun[:2] == sr[1]:
            text += cun + '\n'

with open('dest.csv', 'w', encoding='gbk') as file:
    file.write(text)