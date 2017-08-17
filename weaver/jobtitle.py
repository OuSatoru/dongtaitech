import requests

req_url = 'http://66.185.19.67/hrm/resource/HrmResourceSysOperation.jsp'
login_url = 'http://66.185.19.67/login/VerifyLogin.jsp'

user = {'loginfile': '/wui/theme/ecology8/page/login.jsp?templateId=3&logintype=1&gopage=',
        'logintype': '1', 'fontName': '微软雅黑', 'message': '', 'gopage': '', 'formmethod': 'get',
        'rnd': '', 'serial': '', 'username': '', 'isie': 'false', 'islanguid': '7',
        'loginid': 'sysadmin', 'userpassword': '1', 'submit': '登录'}

headers = {'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36',
'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8',
'Upgrade-Insecure-Requests': '1', 'Cache-Control': 'max-age=0', 
'Referer': 'http://66.185.19.67/hrm/resource/HrmResourceSys.jsp?subcompanyid1=&departmentid=&companyid=1&fromHrmTab=1',
'Origin': 'http://66.185.19.67'}

session = requests.Session()

session.post(login_url, user)

print(session.cookies)

session.headers.update(headers)

modify = {'ids': '', 'method': 'save', 'tableMax': '1', 'subcompanyid1': '', 'departmentid': '',
          'qname': '', 'resourceid': '', 'resourceid__': '', 'managerid_kwd': '', 'managerid_kwd__': '',
          'subcompanyid1_kwd': '', 'subcompanyid1_kwd__': '', 'departmentid_kwd': '',
          'departmentid_kwd__': '', 'accounttype': '',
          'pageId': 'Hrm:ResourceSys',
          'id': '385', 'accounttype': '', 'dsporder': '385', 'loginid': '09800891', 'password': 'qiwuien',
          'email': '',	
          'managerid': '381',	
          'managerid__': '381',	
          'jobtitle': '51', 'jobtitle__': '', 'seclevel': '20', 'passwordstate': '0',
          'pageSizeSel1inputText': '10',
          'pageSizeSel1': '10'}

r = session.post(req_url, modify)

print(r.content)
