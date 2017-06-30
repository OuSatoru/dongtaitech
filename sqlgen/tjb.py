tpl = '''CREATE TABLE ODS.khzjtjb
(
    zjhm VARCHAR(50),
'''

for i in range(1, 42):
    tpl += f'sj{str(i).zfill(2)} DECIMAL(16,2),'

tpl += ');'

print(tpl)

