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