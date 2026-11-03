#!/usr/bin/env python3
"""First installation on the inspected Ubuntu host; run with sudo after uploading the release."""
import os
from pathlib import Path
import pwd
import shutil
import subprocess
import time

stage=Path('/home/ubuntu/pledge-release')
root=Path('/etc/pledge')
assert os.geteuid()==0
assert not (root/'backend.env').exists(), 'Existing deployment: use a reviewed update procedure'
values=dict(line.split('=',1) for line in (stage/'backend.env').read_text().splitlines() if line)
password=values['PLEDGE_MYSQL_PASSWORD']
assert password.isalnum()
for filename in ['pledge-api','pledge-sync']:
 assert (stage/filename).is_file()
try: pwd.getpwnam('pledge')
except KeyError: subprocess.run(['useradd','--system','--home','/var/lib/pledge','--shell','/usr/sbin/nologin','pledge'],check=True)
root.mkdir(mode=0o700)
shutil.move(stage/'backend.env',root/'backend.env')
os.chmod(root/'backend.env',0o600)
os.chown(root/'backend.env',0,0)
subprocess.run(['mysql'],input=f"CREATE DATABASE pledge_study CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci; CREATE USER 'pledge_study'@'127.0.0.1' IDENTIFIED BY '{password}'; GRANT ALL PRIVILEGES ON pledge_study.* TO 'pledge_study'@'127.0.0.1';",text=True,check=True)
release=Path('/opt/pledge/releases')/time.strftime('%Y%m%d%H%M%S',time.gmtime())
release.mkdir(parents=True)
for filename in ['pledge-api','pledge-sync']:
 shutil.copy2(stage/filename,release/filename);os.chmod(release/filename,0o755)
Path('/opt/pledge/current').symlink_to(release)
Path('/var/lib/pledge/redis').mkdir(parents=True)
subprocess.run(['chown','-R','pledge:pledge','/var/lib/pledge'],check=True)
(root/'redis.conf').write_text('bind 127.0.0.1\nport 6381\nprotected-mode yes\ndaemonize no\nsupervised no\nlogfile ""\ndir /var/lib/pledge/redis\npidfile /var/lib/pledge/redis/redis.pid\nsave ""\nappendonly no\n')
os.chmod(root,0o755);os.chmod(root/'redis.conf',0o644)
common='User=pledge\nGroup=pledge\nWorkingDirectory=/opt/pledge/current\nRestart=on-failure\nRestartSec=5\nNoNewPrivileges=true\nPrivateTmp=true\nProtectSystem=strict\nProtectHome=true\nReadWritePaths=/var/lib/pledge\nUMask=0077\n'
units={
 'pledge-redis':'[Unit]\nDescription=Pledge isolated Redis cache\nAfter=network.target\n[Service]\n'+common+'ExecStart=/usr/bin/redis-server /etc/pledge/redis.conf\n[Install]\nWantedBy=multi-user.target\n',
 'pledge-api':'[Unit]\nDescription=Pledge Go API\nAfter=network-online.target mysql.service pledge-redis.service\nWants=network-online.target\nRequires=mysql.service pledge-redis.service\n[Service]\n'+common+'EnvironmentFile=/etc/pledge/backend.env\nExecStart=/opt/pledge/current/pledge-api\n[Install]\nWantedBy=multi-user.target\n',
 'pledge-sync':'[Unit]\nDescription=Pledge Sepolia read-only synchronizer\nAfter=network-online.target mysql.service pledge-redis.service pledge-api.service\nRequires=mysql.service pledge-redis.service\n[Service]\n'+common+'EnvironmentFile=/etc/pledge/backend.env\nExecStart=/opt/pledge/current/pledge-sync\n[Install]\nWantedBy=multi-user.target\n'}
for name,body in units.items():Path('/etc/systemd/system/'+name+'.service').write_text(body)
subprocess.run(['systemctl','daemon-reload'],check=True)
subprocess.run(['systemctl','enable','--now','pledge-redis','pledge-api'],check=True)
print('Installed Pledge release:',release)
print('Owner-only credentials: /etc/pledge/backend.env')
