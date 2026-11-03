#!/usr/bin/env python3
"""Add the Pledge read-only route without replacing existing server locations."""
from pathlib import Path
import shutil
import subprocess
import time

site=Path('/etc/nginx/sites-enabled/default').resolve()
original=site.read_text()
assert 'location /db/' in original and 'proxy_pass http://127.0.0.1:8080' in original
snippet=Path('/etc/nginx/snippets/pledge.conf')
assert not snippet.exists(),'Existing Pledge route requires an update review'
backup=Path('/etc/pledge/nginx-backups');backup.mkdir(mode=0o700)
shutil.copy2(site,backup/('default-'+time.strftime('%Y%m%d%H%M%S',time.gmtime())))
snippet.write_text('''location = /pledge { return 301 /pledge/; }
location ^~ /pledge/ {
    limit_except GET { deny all; }
    proxy_pass http://127.0.0.1:8081/;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_read_timeout 120s;
}
''')
site.write_text(original.replace('location /db/ {','include /etc/nginx/snippets/pledge.conf;\n\n\tlocation /db/ {',1))
try:subprocess.run(['nginx','-t'],check=True)
except Exception:
 site.write_text(original)
 raise
subprocess.run(['systemctl','reload','nginx'],check=True)
print('Pledge read-only route installed; existing locations preserved.')
