tpl = '''SELECT ZJHM,
'''
for i in range(1, 42):
    tpl += f"sum(CASE WHEN zl = '{str(i).zfill(2)}' THEN SJ ELSE 0 END) sj{str(i).zfill(2)}, "

tpl += 'from ods.ygxdzbtmp GROUP BY ZJHM'

print(tpl)