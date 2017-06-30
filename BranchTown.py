import json


def toJSONfile():
    """with a source csv in directory"""
    dic = {}

    with open('SELECT_ZHMC__CMC_FROM_YGXD_ZHXXB_z_LEFT_.csv', encoding='utf8') as file:
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


# comment when not needed
# toJSONfile()
text = ""
d = json.loads(open('branchtown.json', encoding='gbk').read())
print(type(d))

with open('src.csv', encoding='gbk') as file:
    src = file.readlines()
for line in src:
    sr = line.strip('\n').split(',')
    text += sr[0] + '支行' + ','
    for cun in d[sr[0] + '支行']:
        if cun[:2] == sr[1]:
            text += cun + '\n'

with open('dest.csv', 'w', encoding='gbk') as file:
    file.write(text)

'''text = ""

with open('src.csv', encoding='gbk') as file:
    src = file.readlines()
for line in src:
    sr = line.strip('\n').split(',')
    text += sr[0] + '支行' + ','
    for cun in d[sr[0] + '支行']:
        if cun[:2] == sr[1]:
            text += cun + '\n'

    with open('dest.csv', 'w', encoding='gbk') as file:
        file.write(text)'''
