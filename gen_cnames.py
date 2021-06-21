import os
import re
import chevron

MAX_DOMAINS = 16

nginx_template = ""
nginx_ssl_template = ""
validate = re.compile("[a-zA-Z0-9\-\.]+\.[a-zA-Z]+").match

with open("nginx.mustache") as fp:
    nginx_template = fp.read()

with open("nginx.ssl.mustache") as fp:
    nginx_ssl_template = fp.read()

users = list()

for username in os.listdir("/home/"):
    cname = f"/home/{username}/.CNAME"
    if os.path.isfile(cname):
        with open(cname) as fp:
            domains = map(
                lambda domain, _: domain.strip() if validate(domain) else "",
                fp.readlines(),
                range(MAX_DOMAINS)
            )
            domains = tuple(domains)
            users.append(
                dict(username=username, domains=domains)
            )

with open("/tmp/result5.http", "w") as fp:
    fp.write(chevron.render(nginx_template, dict(users=users)))

with open("/tmp/result5.https", "w") as fp:
    fp.write(chevron.render(nginx_ssl_template, dict(users=users)))
