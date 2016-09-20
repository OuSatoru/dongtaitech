import csv

with open('dest.csv', encoding='gbk') as f:
    dic = csv.DictReader(f)
    for i in dic:
        print(i)

